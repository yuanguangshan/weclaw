package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fastclaw-ai/weclaw/ilink"
	"github.com/fastclaw-ai/weclaw/messaging"
	"github.com/fastclaw-ai/weclaw/web"
)

// Server provides an HTTP API for sending messages and admin management.
type Server struct {
	clients  []*ilink.Client
	addr     string
	todosMu  sync.Mutex // protects todos.json read/write
	timersMu sync.Mutex // protects timers.json read/write
}

// NewServer creates an API server.
func NewServer(clients []*ilink.Client, addr string) *Server {
	if addr == "" {
		addr = "127.0.0.1:18011"
	}
	return &Server{clients: clients, addr: addr}
}

// SendRequest is the JSON body for POST /api/send.
type SendRequest struct {
	To       string `json:"to"`
	Text     string `json:"text,omitempty"`
	MediaURL string `json:"media_url,omitempty"` // image/video/file URL
}

// Run starts the HTTP server. Blocks until ctx is cancelled.
func (s *Server) Run(ctx context.Context) error {
	mux := http.NewServeMux()

	// Existing endpoints
	mux.HandleFunc("/api/send", s.handleSend)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

	// Admin API - Config & Agents
	mux.HandleFunc("GET /api/config", s.handleGetConfig)
	mux.HandleFunc("PUT /api/config", s.handleUpdateConfig)
	mux.HandleFunc("GET /api/agents", s.handleListAgents)
	mux.HandleFunc("POST /api/agents", s.handleAddAgent)
	mux.HandleFunc("PUT /api/agents/{name}", s.handleUpdateAgent)
	mux.HandleFunc("DELETE /api/agents/{name}", s.handleDeleteAgent)
	mux.HandleFunc("POST /api/agents/detect", s.handleDetectAgents)

	// Admin API - Status & Accounts
	mux.HandleFunc("GET /api/status", s.handleStatus)
	mux.HandleFunc("GET /api/accounts", s.handleListAccounts)
	mux.HandleFunc("DELETE /api/accounts/{id}", s.handleDeleteAccount)

	// Admin API - Login
	mux.HandleFunc("GET /api/login/qrcode", s.handleLoginQRCode)
	mux.HandleFunc("GET /api/login/status", s.handleLoginStatus)

	// Admin API - Service Control
	mux.HandleFunc("POST /api/service/restart", s.handleServiceRestart)
	mux.HandleFunc("POST /api/service/update", s.handleServiceUpdate)

	// Admin API - Logs
	mux.HandleFunc("GET /api/logs", s.handleLogs)

	// Admin API - Hub
	mux.HandleFunc("GET /api/hub", s.handleListHub)
	mux.HandleFunc("GET /api/hub/{name}", s.handleReadHubFile)
	mux.HandleFunc("DELETE /api/hub/{name}", s.handleDeleteHubFile)
	mux.HandleFunc("POST /api/hub/clear", s.handleClearHub)

	// Admin API - Todos
	mux.HandleFunc("GET /api/todos", s.handleListTodos)
	mux.HandleFunc("POST /api/todos", s.handleAddTodo)
	mux.HandleFunc("PUT /api/todos/{id}/done", s.handleDoneTodo)
	mux.HandleFunc("DELETE /api/todos/{id}", s.handleDeleteTodo)

	// Admin API - Timers
	mux.HandleFunc("GET /api/timers", s.handleListTimers)
	mux.HandleFunc("POST /api/timers", s.handleAddTimer)
	mux.HandleFunc("PUT /api/timers/{id}/cancel", s.handleCancelTimer)

	// Admin API - Cron Jobs
	mux.HandleFunc("GET /api/cron", s.handleListCronJobs)
	mux.HandleFunc("POST /api/cron", s.handleAddCronJob)
	mux.HandleFunc("DELETE /api/cron/{id}", s.handleDeleteCronJob)
	mux.HandleFunc("PUT /api/cron/{id}/enable", s.handleEnableCronJob)
	mux.HandleFunc("PUT /api/cron/{id}/disable", s.handleDisableCronJob)

	// Admin UI
	mux.HandleFunc("/admin", s.handleAdminUI)
	mux.HandleFunc("/admin/", s.handleAdminUI)

	srv := &http.Server{Addr: s.addr, Handler: mux}

	go func() {
		<-ctx.Done()
		srv.Shutdown(context.Background())
	}()

	log.Printf("[api] listening on %s", s.addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) handleAdminUI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(web.AdminHTML)
}

func (s *Server) handleSend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}

	var req SendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if req.To == "" {
		http.Error(w, `"to" is required`, http.StatusBadRequest)
		return
	}
	if req.Text == "" && req.MediaURL == "" {
		http.Error(w, `"text" or "media_url" is required`, http.StatusBadRequest)
		return
	}

	if len(s.clients) == 0 {
		http.Error(w, "no accounts configured", http.StatusServiceUnavailable)
		return
	}

	// Use the first client
	client := s.clients[0]
	ctx := r.Context()

	// Send text if provided
	if req.Text != "" {
		if err := messaging.SendTextReply(ctx, client, req.To, req.Text, "", ""); err != nil {
			log.Printf("[api] send text failed: %v", err)
			http.Error(w, "send text failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("[api] sent text to %s: %q", req.To, req.Text)

		// Extract and send any markdown images embedded in text
		for _, imgURL := range messaging.ExtractImageURLs(req.Text) {
			if err := messaging.SendMediaFromURL(ctx, client, req.To, imgURL, ""); err != nil {
				log.Printf("[api] send extracted image failed: %v", err)
			} else {
				log.Printf("[api] sent extracted image to %s: %s", req.To, imgURL)
			}
		}
	}

	// Send media if provided
	if req.MediaURL != "" {
		if err := messaging.SendMediaFromURL(ctx, client, req.To, req.MediaURL, ""); err != nil {
			log.Printf("[api] send media failed: %v", err)
			http.Error(w, "send media failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("[api] sent media to %s: %s", req.To, req.MediaURL)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// Cron Job types
type CronJob struct {
	ID        string       `json:"id"`
	UserID    string       `json:"user_id"`
	CronExpr  string       `json:"cron_expr"`
	Command   CronCommand  `json:"command"`
	Enabled   bool         `json:"enabled"`
	CreatedAt int64        `json:"created_at"`
	NextRun   int64        `json:"next_run,omitempty"`
}

type CronCommand struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	Agent   string `json:"agent,omitempty"`
}

type CronJobsFile struct {
	Jobs []*CronJob `json:"jobs"`
}

// cronDataDir returns the directory where cron jobs are stored
func cronDataDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "/tmp/.weclaw"
	}
	return filepath.Join(home, ".weclaw")
}

// loadCronJobs loads all cron jobs from disk
func loadCronJobs() (*CronJobsFile, error) {
	dataDir := cronDataDir()
	data, err := os.ReadFile(filepath.Join(dataDir, "cron_jobs.json"))
	if err != nil {
		if os.IsNotExist(err) {
			return &CronJobsFile{Jobs: []*CronJob{}}, nil
		}
		return nil, err
	}

	var result CronJobsFile
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// saveCronJobs saves all cron jobs to disk
func saveCronJobs(jobs *CronJobsFile) error {
	dataDir := cronDataDir()
	data, err := json.MarshalIndent(jobs, "", "  ")
	if err != nil {
		return err
	}

	filePath := filepath.Join(dataDir, "cron_jobs.json")
	tmpPath := filePath + ".tmp"
	if err := os.WriteFile(tmpPath, data, 0600); err != nil {
		return err
	}

	return os.Rename(tmpPath, filePath)
}

// handleListCronJobs returns all cron jobs
func (s *Server) handleListCronJobs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "GET only", http.StatusMethodNotAllowed)
		return
	}

	jobs, err := loadCronJobs()
	if err != nil {
		http.Error(w, "Failed to load cron jobs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs.Jobs)
}

// handleAddCronJob adds a new cron job
func (s *Server) handleAddCronJob(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}

	var job CronJob
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if job.CronExpr == "" {
		http.Error(w, `"cron_expr" is required`, http.StatusBadRequest)
		return
	}
	if job.Command.Content == "" {
		http.Error(w, `"command.content" is required`, http.StatusBadRequest)
		return
	}
	if job.Command.Type == "" {
		job.Command.Type = "text"
	}

	// Generate ID if not provided
	if job.ID == "" {
		job.ID = fmt.Sprintf("cron_%d", time.Now().UnixNano())
	}

	// Set created at if not provided
	if job.CreatedAt == 0 {
		job.CreatedAt = time.Now().Unix()
	}

	// Load existing jobs
	jobs, err := loadCronJobs()
	if err != nil {
		http.Error(w, "Failed to load cron jobs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Add new job
	jobs.Jobs = append(jobs.Jobs, &job)

	// Save
	if err := saveCronJobs(jobs); err != nil {
		http.Error(w, "Failed to save cron jobs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "ok",
		"job":    job,
	})
}

// handleDeleteCronJob deletes a cron job
func (s *Server) handleDeleteCronJob(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "DELETE only", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path
	id := strings.TrimPrefix(r.URL.Path, "/api/cron/")
	if id == "" {
		http.Error(w, "job ID is required", http.StatusBadRequest)
		return
	}

	// Load existing jobs
	jobs, err := loadCronJobs()
	if err != nil {
		http.Error(w, "Failed to load cron jobs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Find and remove job
	found := false
	var filtered []*CronJob
	for _, j := range jobs.Jobs {
		if j.ID == id {
			found = true
		} else {
			filtered = append(filtered, j)
		}
	}

	if !found {
		http.Error(w, "job not found", http.StatusNotFound)
		return
	}

	jobs.Jobs = filtered

	// Save
	if err := saveCronJobs(jobs); err != nil {
		http.Error(w, "Failed to save cron jobs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// handleEnableCronJob enables a cron job
func (s *Server) handleEnableCronJob(w http.ResponseWriter, r *http.Request) {
	s.updateCronJobEnabled(w, r, true)
}

// handleDisableCronJob disables a cron job
func (s *Server) handleDisableCronJob(w http.ResponseWriter, r *http.Request) {
	s.updateCronJobEnabled(w, r, false)
}

// updateCronJobEnabled updates the enabled status of a cron job
func (s *Server) updateCronJobEnabled(w http.ResponseWriter, r *http.Request, enabled bool) {
	if r.Method != http.MethodPut {
		http.Error(w, "PUT only", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path
	id := strings.TrimPrefix(r.URL.Path, "/api/cron/")
	id = strings.TrimSuffix(id, "/enable")
	id = strings.TrimSuffix(id, "/disable")
	if id == "" {
		http.Error(w, "job ID is required", http.StatusBadRequest)
		return
	}

	// Load existing jobs
	jobs, err := loadCronJobs()
	if err != nil {
		http.Error(w, "Failed to load cron jobs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Find and update job
	found := false
	for _, j := range jobs.Jobs {
		if j.ID == id {
			j.Enabled = enabled
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "job not found", http.StatusNotFound)
		return
	}

	// Save
	if err := saveCronJobs(jobs); err != nil {
		http.Error(w, "Failed to save cron jobs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	status := "disabled"
	if enabled {
		status = "enabled"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"job":    status,
	})
}
