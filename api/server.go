package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fastclaw-ai/weclaw/ilink"
	"github.com/fastclaw-ai/weclaw/messaging"
	"github.com/fastclaw-ai/weclaw/web"
)

// Server provides an HTTP API for sending messages and admin management.
type Server struct {
	clients []*ilink.Client
	addr    string
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
	mux.HandleFunc("DELETE /api/hub/clear", s.handleClearHub)

	// Admin API - Todos
	mux.HandleFunc("GET /api/todos", s.handleListTodos)
	mux.HandleFunc("POST /api/todos", s.handleAddTodo)
	mux.HandleFunc("PUT /api/todos/{id}/done", s.handleDoneTodo)
	mux.HandleFunc("DELETE /api/todos/{id}", s.handleDeleteTodo)

	// Admin API - Timers
	mux.HandleFunc("GET /api/timers", s.handleListTimers)
	mux.HandleFunc("POST /api/timers", s.handleAddTimer)
	mux.HandleFunc("PUT /api/timers/{id}/cancel", s.handleCancelTimer)

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
