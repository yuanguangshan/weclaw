package api

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/fastclaw-ai/weclaw/config"
	"github.com/fastclaw-ai/weclaw/hub"
	"github.com/fastclaw-ai/weclaw/ilink"
)

// --- Config & Agents ---

func (s *Server) handleGetConfig(w http.ResponseWriter, r *http.Request) {
	cfg, err := config.Load()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, cfg)
}

func (s *Server) handleUpdateConfig(w http.ResponseWriter, r *http.Request) {
	var req struct {
		DefaultAgent string `json:"default_agent"`
		APIAddr      string `json:"api_addr"`
		SaveDir      string `json:"save_dir"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	cfg, err := config.Load()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if req.DefaultAgent != "" {
		cfg.DefaultAgent = req.DefaultAgent
	}
	if req.APIAddr != "" {
		cfg.APIAddr = req.APIAddr
	}
	if req.SaveDir != "" {
		cfg.SaveDir = req.SaveDir
	}
	if err := config.Save(cfg); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, map[string]string{"status": "ok"})
}

func (s *Server) handleListAgents(w http.ResponseWriter, r *http.Request) {
	cfg, err := config.Load()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, cfg.Agents)
}

func (s *Server) handleAddAgent(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
		config.AgentConfig
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	if req.Name == "" {
		writeError(w, http.StatusBadRequest, "name is required")
		return
	}
	cfg, err := config.Load()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if _, exists := cfg.Agents[req.Name]; exists {
		writeError(w, http.StatusConflict, "agent already exists: "+req.Name)
		return
	}
	cfg.Agents[req.Name] = req.AgentConfig
	if err := config.Save(cfg); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, map[string]string{"status": "ok"})
}

func (s *Server) handleUpdateAgent(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	if name == "" {
		writeError(w, http.StatusBadRequest, "agent name is required")
		return
	}
	var agentCfg config.AgentConfig
	if err := json.NewDecoder(r.Body).Decode(&agentCfg); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	cfg, err := config.Load()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if _, exists := cfg.Agents[name]; !exists {
		writeError(w, http.StatusNotFound, "agent not found: "+name)
		return
	}
	cfg.Agents[name] = agentCfg
	if err := config.Save(cfg); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, map[string]string{"status": "ok"})
}

func (s *Server) handleDeleteAgent(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	if name == "" {
		writeError(w, http.StatusBadRequest, "agent name is required")
		return
	}
	cfg, err := config.Load()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if _, exists := cfg.Agents[name]; !exists {
		writeError(w, http.StatusNotFound, "agent not found: "+name)
		return
	}
	delete(cfg.Agents, name)
	if cfg.DefaultAgent == name {
		cfg.DefaultAgent = ""
	}
	if err := config.Save(cfg); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, map[string]string{"status": "ok"})
}

func (s *Server) handleDetectAgents(w http.ResponseWriter, r *http.Request) {
	cfg, err := config.Load()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	before := len(cfg.Agents)
	modified := config.DetectAndConfigure(cfg)
	if modified {
		if err := config.Save(cfg); err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	after := len(cfg.Agents)
	writeJSON(w, map[string]interface{}{
		"status":      "ok",
		"modified":    modified,
		"agents_before": before,
		"agents_after":  after,
		"new_agents":   after - before,
	})
}

// --- Status & Accounts ---

func (s *Server) handleStatus(w http.ResponseWriter, r *http.Request) {
	home, _ := os.UserHomeDir()
	weclawDir := filepath.Join(home, ".weclaw")

	// Check PID file (daemon mode)
	pidFile := filepath.Join(weclawDir, "weclaw.pid")
	pid := 0
	running := false
	uptime := ""
	if data, err := os.ReadFile(pidFile); err == nil {
		if p, err := strconv.Atoi(strings.TrimSpace(string(data))); err == nil {
			pid = p
			if proc, err := os.FindProcess(p); err == nil {
				if proc.Signal(nil) == nil {
					running = true
					if stat, err := os.Stat(filepath.Join(weclawDir, "weclaw.log")); err == nil {
						uptime = time.Since(stat.ModTime()).Truncate(time.Second).String()
					}
				}
			}
		}
	}

	// If PID file check failed, we are running (e.g. under systemd)
	if !running {
		pid = os.Getpid()
		running = true
		// Get process start time from /proc/self/stat
		if statData, err := os.ReadFile("/proc/self/stat"); err == nil {
			fields := strings.Fields(string(statData))
			if len(fields) > 21 {
				if starttime, err := strconv.ParseUint(fields[21], 10, 64); err == nil {
					const clocksPerSec = 100
					secsSinceBoot := starttime / clocksPerSec
					if bstat, err := os.ReadFile("/proc/stat"); err == nil {
						for _, line := range strings.Split(string(bstat), "\n") {
							if strings.HasPrefix(line, "btime ") {
								parts := strings.Fields(line)
								if len(parts) >= 2 {
									if btime, err := strconv.ParseInt(parts[1], 10, 64); err == nil {
										startTime := time.Unix(btime+int64(secsSinceBoot), 0)
										uptime = time.Since(startTime).Truncate(time.Second).String()
									}
								}
								break
							}
						}
					}
				}
			}
		}
	}

	// Config info
	cfg, _ := config.Load()
	agentCount := 0
	if cfg != nil {
		agentCount = len(cfg.Agents)
	}

	// Accounts
	accounts, _ := ilink.LoadAllCredentials()
	accountCount := 0
	if accounts != nil {
		accountCount = len(accounts)
	}

	// Hub files
	hubDir := hub.DefaultDir()
	hubFileCount := 0
	if entries, err := os.ReadDir(hubDir); err == nil {
		for _, e := range entries {
			if !e.IsDir() {
				hubFileCount++
			}
		}
	}

	writeJSON(w, map[string]interface{}{
		"running":       running,
		"pid":           pid,
		"uptime":        uptime,
		"agent_count":   agentCount,
		"account_count": accountCount,
		"hub_files":     hubFileCount,
		"default_agent": func() string {
			if cfg != nil {
				return cfg.DefaultAgent
			}
			return ""
		}(),
		"api_addr": func() string {
			if cfg != nil {
				return cfg.APIAddr
			}
			return ""
		}(),
		"version": "dev",
	})
}

func (s *Server) handleListAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := ilink.LoadAllCredentials()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// Don't expose sensitive tokens
	type accountInfo struct {
		ID         string `json:"id"`
		ILinkBotID string `json:"ilink_bot_id"`
	}
	var result []accountInfo
	for _, a := range accounts {
		id := normalizeAccountID(a.ILinkBotID)
		result = append(result, accountInfo{
			ID:         id,
			ILinkBotID: a.ILinkBotID,
		})
	}
	if result == nil {
		result = []accountInfo{}
	}
	writeJSON(w, result)
}

func (s *Server) handleDeleteAccount(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		writeError(w, http.StatusBadRequest, "account id is required")
		return
	}
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".weclaw", "accounts", id+".json")
	if err := os.Remove(path); err != nil {
		writeError(w, http.StatusNotFound, "account not found: "+err.Error())
		return
	}
	writeJSON(w, map[string]string{"status": "ok"})
}

// --- Logs ---

func (s *Server) handleLogs(w http.ResponseWriter, r *http.Request) {
	home, _ := os.UserHomeDir()
	logPath := filepath.Join(home, ".weclaw", "weclaw.log")

	lines := 200
	if l := r.URL.Query().Get("lines"); l != "" {
		if n, err := strconv.Atoi(l); err == nil && n > 0 && n <= 2000 {
			lines = n
		}
	}

	f, err := os.Open(logPath)
	if err != nil {
		writeJSON(w, []string{})
		return
	}
	defer f.Close()

	// Read last N lines efficiently
	var result []string
	scanner := bufio.NewScanner(f)
	buf := make([]string, 0, lines)
	for scanner.Scan() {
		buf = append(buf, scanner.Text())
		if len(buf) > lines {
			buf = buf[1:]
		}
	}
	result = buf
	if result == nil {
		result = []string{}
	}
	writeJSON(w, result)
}

// --- Hub ---

func (s *Server) handleListHub(w http.ResponseWriter, r *http.Request) {
	h := hub.New(hub.DefaultDir())
	files, err := h.ListWithInfo()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	type fileInfo struct {
		Name    string `json:"name"`
		ModTime string `json:"mod_time"`
	}
	var result []fileInfo
	for _, f := range files {
		result = append(result, fileInfo{
			Name:    f.Name,
			ModTime: f.ModTime.Format("2006-01-02 15:04:05"),
		})
	}
	if result == nil {
		result = []fileInfo{}
	}
	writeJSON(w, result)
}

func (s *Server) handleReadHubFile(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	if name == "" {
		writeError(w, http.StatusBadRequest, "file name is required")
		return
	}
	h := hub.New(hub.DefaultDir())
	content, err := h.ReadFile(name)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, map[string]string{"name": name, "content": content})
}

func (s *Server) handleDeleteHubFile(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	if name == "" {
		writeError(w, http.StatusBadRequest, "file name is required")
		return
	}
	h := hub.New(hub.DefaultDir())
	sharedDir := h.SharedDir()
	path := filepath.Join(sharedDir, filepath.Base(name))
	if err := os.Remove(path); err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, map[string]string{"status": "ok"})
}

func (s *Server) handleClearHub(w http.ResponseWriter, r *http.Request) {
	h := hub.New(hub.DefaultDir())
	count, err := h.Clear()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, map[string]interface{}{"status": "ok", "deleted": count})
}

// --- Todos ---

func (s *Server) handleListTodos(w http.ResponseWriter, r *http.Request) {
	dataDir := hub.DefaultDir()
	path := filepath.Join(dataDir, "todos.json")
	data, err := os.ReadFile(path)
	if err != nil {
		writeJSON(w, []interface{}{})
		return
	}
	var items []interface{}
	if err := json.Unmarshal(data, &items); err != nil {
		writeJSON(w, []interface{}{})
		return
	}
	writeJSON(w, items)
}

func (s *Server) handleAddTodo(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	if req.Title == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}

	dataDir := hub.DefaultDir()
	path := filepath.Join(dataDir, "todos.json")

	// Read existing
	var items []map[string]interface{}
	if data, err := os.ReadFile(path); err == nil {
		json.Unmarshal(data, &items)
	}
	if items == nil {
		items = []map[string]interface{}{}
	}

	// Find next ID
	nextID := 1
	for _, item := range items {
		if id, ok := item["id"].(float64); ok && int(id) >= nextID {
			nextID = int(id) + 1
		}
	}

	item := map[string]interface{}{
		"id":         nextID,
		"user_id":    "admin",
		"title":      req.Title,
		"due_time":   0,
		"status":     0,
		"created_at": time.Now().Unix(),
		"reminded":   false,
	}
	items = append(items, item)

	data, _ := json.MarshalIndent(items, "", "  ")
	os.WriteFile(path, data, 0644)

	writeJSON(w, item)
}

func (s *Server) handleDoneTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	dataDir := hub.DefaultDir()
	path := filepath.Join(dataDir, "todos.json")
	data, err := os.ReadFile(path)
	if err != nil {
		writeError(w, http.StatusNotFound, "no todos")
		return
	}

	var items []map[string]interface{}
	json.Unmarshal(data, &items)

	found := false
	for _, item := range items {
		if itemID, ok := item["id"].(float64); ok && int(itemID) == id {
			item["status"] = 1
			found = true
			break
		}
	}
	if !found {
		writeError(w, http.StatusNotFound, "todo not found")
		return
	}

	data, _ = json.MarshalIndent(items, "", "  ")
	os.WriteFile(path, data, 0644)
	writeJSON(w, map[string]string{"status": "ok"})
}

func (s *Server) handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	dataDir := hub.DefaultDir()
	path := filepath.Join(dataDir, "todos.json")
	data, err := os.ReadFile(path)
	if err != nil {
		writeError(w, http.StatusNotFound, "no todos")
		return
	}

	var items []map[string]interface{}
	json.Unmarshal(data, &items)

	var remaining []map[string]interface{}
	found := false
	for _, item := range items {
		if itemID, ok := item["id"].(float64); ok && int(itemID) == id {
			found = true
			continue
		}
		remaining = append(remaining, item)
	}
	if !found {
		writeError(w, http.StatusNotFound, "todo not found")
		return
	}

	if remaining == nil {
		remaining = []map[string]interface{}{}
	}
	data, _ = json.MarshalIndent(remaining, "", "  ")
	os.WriteFile(path, data, 0644)
	writeJSON(w, map[string]string{"status": "ok"})
}

// --- Timers ---

func (s *Server) handleListTimers(w http.ResponseWriter, r *http.Request) {
	dataDir := hub.DefaultDir()
	path := filepath.Join(dataDir, "timers.json")
	data, err := os.ReadFile(path)
	if err != nil {
		writeJSON(w, []interface{}{})
		return
	}
	var items []interface{}
	if err := json.Unmarshal(data, &items); err != nil {
		writeJSON(w, []interface{}{})
		return
	}
	writeJSON(w, items)
}

func (s *Server) handleAddTimer(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Duration int64  `json:"duration"` // seconds
		Label    string `json:"label"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	if req.Duration <= 0 {
		writeError(w, http.StatusBadRequest, "duration must be > 0")
		return
	}
	if req.Duration > 86400 {
		writeError(w, http.StatusBadRequest, "timer cannot exceed 24 hours")
		return
	}
	if req.Label == "" {
		req.Label = "Timer"
	}

	dataDir := hub.DefaultDir()
	path := filepath.Join(dataDir, "timers.json")

	var items []map[string]interface{}
	if data, err := os.ReadFile(path); err == nil {
		json.Unmarshal(data, &items)
	}
	if items == nil {
		items = []map[string]interface{}{}
	}

	nextID := 1
	for _, item := range items {
		if id, ok := item["id"].(float64); ok && int(id) >= nextID {
			nextID = int(id) + 1
		}
	}

	now := time.Now().Unix()
	item := map[string]interface{}{
		"id":         nextID,
		"user_id":    "admin",
		"label":      req.Label,
		"duration":   req.Duration,
		"end_time":   now + req.Duration,
		"status":     0,
		"created_at": now,
		"reminded":   false,
	}
	items = append(items, item)

	data, _ := json.MarshalIndent(items, "", "  ")
	os.WriteFile(path, data, 0644)

	writeJSON(w, item)
}

func (s *Server) handleCancelTimer(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	dataDir := hub.DefaultDir()
	path := filepath.Join(dataDir, "timers.json")
	data, err := os.ReadFile(path)
	if err != nil {
		writeError(w, http.StatusNotFound, "no timers")
		return
	}

	var items []map[string]interface{}
	json.Unmarshal(data, &items)

	found := false
	for _, item := range items {
		if itemID, ok := item["id"].(float64); ok && int(itemID) == id {
			item["status"] = 2
			found = true
			break
		}
	}
	if !found {
		writeError(w, http.StatusNotFound, "timer not found")
		return
	}

	data, _ = json.MarshalIndent(items, "", "  ")
	os.WriteFile(path, data, 0644)
	writeJSON(w, map[string]string{"status": "ok"})
}

// --- Login ---

func (s *Server) handleLoginQRCode(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	qr, err := ilink.FetchQRCode(ctx)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "fetch QR failed: "+err.Error())
		return
	}
	writeJSON(w, map[string]string{
		"qrcode": qr.QRCode,
		"url":    qr.QRCodeImgContent,
	})
}

func (s *Server) handleLoginStatus(w http.ResponseWriter, r *http.Request) {
	qrcode := r.URL.Query().Get("qrcode")
	if qrcode == "" {
		writeError(w, http.StatusBadRequest, "qrcode parameter required")
		return
	}

	// Single poll (not infinite loop like CLI) — frontend will repeat
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()

	c := ilink.NewUnauthenticatedClient()

	var resp ilink.QRStatusResponse
	if err := c.DoGet(ctx, "https://ilinkai.weixin.qq.com/ilink/bot/get_qrcode_status?qrcode="+qrcode, &resp); err != nil {
		writeJSON(w, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	result := map[string]interface{}{"status": resp.Status}
	if resp.Status == "confirmed" && resp.BotToken != "" {
		creds := &ilink.Credentials{
			BotToken:    resp.BotToken,
			ILinkBotID:  resp.ILinkBotID,
			BaseURL:     resp.BaseURL,
			ILinkUserID: resp.ILinkUserID,
		}
		if err := ilink.SaveCredentials(creds); err != nil {
			writeError(w, http.StatusInternalServerError, "save failed: "+err.Error())
			return
		}
		result["bot_id"] = creds.ILinkBotID
		log.Printf("[admin] new account logged in: %s", creds.ILinkBotID)
	}
	writeJSON(w, result)
}

// --- Service Control ---

func (s *Server) handleServiceRestart(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]string{"status": "restarting"})
	// Flush response before restart kills us
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
	go func() {
		time.Sleep(500 * time.Millisecond)
		exec.Command("systemctl", "restart", "weclaw").Run()
	}()
}

func (s *Server) handleServiceUpdate(w http.ResponseWriter, r *http.Request) {
	projectDir := "/home/nanobot/.nanobot/weclaw"

	// Build
	cmd := exec.Command("go", "build", "-o", filepath.Join(projectDir, "weclaw"), projectDir)
	cmd.Dir = projectDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "build failed: "+string(output))
		return
	}

	// Copy
	src := filepath.Join(projectDir, "weclaw")
	dst := "/usr/local/bin/weclaw"
	if err := CopyFile(dst, src); err != nil {
		writeError(w, http.StatusInternalServerError, "copy failed: "+err.Error())
		return
	}
	// Make executable
	os.Chmod(dst, 0755)

	writeJSON(w, map[string]string{"status": "built, restarting"})
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
	go func() {
		time.Sleep(500 * time.Millisecond)
		exec.Command("systemctl", "restart", "weclaw").Run()
	}()
}

// --- Helpers ---

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

// CopyFile copies a file from src to dst (used internally if needed).
func CopyFile(dst, src string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

// ensure admin.go uses log and fmt
var _ = log.Printf
var _ = fmt.Sprintf
var _ = io.Copy

// normalizeAccountID creates a safe filename from a bot ID.
func normalizeAccountID(raw string) string {
	s := raw
	for _, ch := range []string{"@", ".", ":"} {
		s = strings.ReplaceAll(s, ch, "-")
	}
	return s
}
