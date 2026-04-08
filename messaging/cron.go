package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fastclaw-ai/weclaw/ilink"
	"github.com/robfig/cron/v3"
)

// CronJob represents a scheduled task.
type CronJob struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	CronExpr  string    `json:"cron_expr"`
	Command   CronCommand `json:"command"`
	Enabled   bool      `json:"enabled"`
	CreatedAt int64     `json:"created_at"`
	NextRun   int64     `json:"next_run,omitempty"`
}

// CronCommand represents the command to execute when cron triggers.
type CronCommand struct {
	Type    string `json:"type"`    // text, workflow, agent
	Content string `json:"content"` // message text or workflow DSL
	Agent   string `json:"agent,omitempty"`   // optional agent for text type
}

// CronStore manages cron job persistence.
type CronStore struct {
	filePath string
	mu       sync.RWMutex
}

// NewCronStore creates a new CronStore.
func NewCronStore(dataDir string) *CronStore {
	return &CronStore{
		filePath: filepath.Join(dataDir, "cron_jobs.json"),
	}
}

// LoadAll loads all cron jobs from disk.
func (s *CronStore) LoadAll() ([]*CronJob, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []*CronJob{}, nil
		}
		return nil, err
	}

	var result struct {
		Jobs []*CronJob `json:"jobs"`
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result.Jobs, nil
}

// SaveAll saves all cron jobs to disk.
func (s *CronStore) SaveAll(jobs []*CronJob) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data := struct {
		Jobs []*CronJob `json:"jobs"`
	}{
		Jobs: jobs,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Write to temp file first, then rename for atomicity
	tmpPath := s.filePath + ".tmp"
	if err := os.WriteFile(tmpPath, jsonData, 0600); err != nil {
		return err
	}

	return os.Rename(tmpPath, s.filePath)
}

// Add adds a new cron job.
func (s *CronStore) Add(job *CronJob) error {
	jobs, err := s.LoadAll()
	if err != nil {
		return err
	}

	jobs = append(jobs, job)
	return s.SaveAll(jobs)
}

// Remove removes a cron job by ID.
func (s *CronStore) Remove(id string) error {
	jobs, err := s.LoadAll()
	if err != nil {
		return err
	}

	var filtered []*CronJob
	for _, j := range jobs {
		if j.ID != id {
			filtered = append(filtered, j)
		}
	}

	return s.SaveAll(filtered)
}

// Update updates a cron job.
func (s *CronStore) Update(job *CronJob) error {
	jobs, err := s.LoadAll()
	if err != nil {
		return err
	}

	for i, j := range jobs {
		if j.ID == job.ID {
			jobs[i] = job
			return s.SaveAll(jobs)
		}
	}

	return fmt.Errorf("job %q not found", job.ID)
}

// GetByUserID returns all jobs for a specific user.
func (s *CronStore) GetByUserID(userID string) ([]*CronJob, error) {
	jobs, err := s.LoadAll()
	if err != nil {
		return nil, err
	}

	var result []*CronJob
	for _, j := range jobs {
		if j.UserID == userID {
			result = append(result, j)
		}
	}

	return result, nil
}

// CronManager manages scheduled cron jobs.
type CronManager struct {
	cron    *cron.Cron
	jobs    map[string]cron.EntryID
	store   *CronStore
	client  *ilink.Client
	handler *Handler
	mu      sync.RWMutex
}

// NewCronManager creates a new CronManager.
func NewCronManager(dataDir string, client *ilink.Client, handler *Handler) *CronManager {
	// Use cron with seconds precision and local timezone
	return &CronManager{
		cron:    cron.New(cron.WithSeconds()),
		jobs:    make(map[string]cron.EntryID),
		store:   NewCronStore(dataDir),
		client:  client,
		handler: handler,
	}
}

// Start starts the cron scheduler and loads existing jobs.
func (cm *CronManager) Start() error {
	cm.cron.Start()

	// Load and schedule existing jobs
	jobs, err := cm.store.LoadAll()
	if err != nil {
		log.Printf("[cron] failed to load jobs: %v", err)
		return err
	}

	for _, job := range jobs {
		if job.Enabled {
			if err := cm.scheduleJob(job); err != nil {
				log.Printf("[cron] failed to schedule job %s: %v", job.ID, err)
			}
		}
	}

	log.Printf("[cron] started with %d jobs", len(cm.jobs))
	return nil
}

// Stop stops the cron scheduler.
func (cm *CronManager) Stop() {
	cm.cron.Stop()
	log.Printf("[cron] stopped")
}

// AddJob adds a new cron job.
func (cm *CronManager) AddJob(job *CronJob) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	// Save to store first
	if err := cm.store.Add(job); err != nil {
		return err
	}

	// Schedule if enabled
	if job.Enabled {
		return cm.scheduleJob(job)
	}

	return nil
}

// RemoveJob removes a cron job.
func (cm *CronManager) RemoveJob(id string) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	// Remove from cron
	if entryID, ok := cm.jobs[id]; ok {
		cm.cron.Remove(entryID)
		delete(cm.jobs, id)
	}

	// Remove from store
	return cm.store.Remove(id)
}

// UpdateJob updates a cron job.
func (cm *CronManager) UpdateJob(job *CronJob) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	// Update in store
	if err := cm.store.Update(job); err != nil {
		return err
	}

	// Reschedule
	if entryID, ok := cm.jobs[job.ID]; ok {
		cm.cron.Remove(entryID)
		delete(cm.jobs, job.ID)
	}

	if job.Enabled {
		return cm.scheduleJob(job)
	}

	return nil
}

// ListJobs returns all jobs for a user.
func (cm *CronManager) ListJobs(userID string) ([]*CronJob, error) {
	return cm.store.GetByUserID(userID)
}

// scheduleJob schedules a single job.
func (cm *CronManager) scheduleJob(job *CronJob) error {
	entryID, err := cm.cron.AddFunc(job.CronExpr, func() {
		cm.executeJob(job)
	})
	if err != nil {
		return err
	}

	cm.jobs[job.ID] = entryID

	// Update next run time
	job.NextRun = time.Now().Add(time.Minute).Unix() // approximate
	cm.store.Update(job)

	return nil
}

// executeJob executes a cron job.
func (cm *CronManager) executeJob(job *CronJob) {
	log.Printf("[cron] executing job %s for user %s", job.ID, job.UserID)

	ctx := context.Background()

	switch job.Command.Type {
	case "text":
		cm.executeTextJob(ctx, job)
	case "workflow":
		cm.executeWorkflowJob(ctx, job)
	case "agent":
		cm.executeAgentJob(ctx, job)
	default:
		log.Printf("[cron] unknown command type: %s", job.Command.Type)
	}
}

// executeTextJob sends a text message to the user.
func (cm *CronManager) executeTextJob(ctx context.Context, job *CronJob) {
	clientID := NewClientID()
	content := job.Command.Content

	req := &ilink.SendMessageRequest{
		Msg: ilink.SendMsg{
			FromUserID:   cm.client.BotID(),
			ToUserID:     job.UserID,
			ClientID:     clientID,
			MessageType:  ilink.MessageTypeBot,
			MessageState: ilink.MessageStateFinish,
			ItemList: []ilink.MessageItem{
				{Type: ilink.ItemTypeText, TextItem: &ilink.TextItem{Text: content}},
			},
		},
	}

	resp, err := cm.client.SendMessage(ctx, req)
	if err != nil || resp.Ret != 0 {
		log.Printf("[cron] failed to send text message: err=%v ret=%d", err, resp.Ret)
	} else {
		log.Printf("[cron] sent text message to %s", job.UserID)
	}
}

// executeWorkflowJob executes a workflow DSL.
func (cm *CronManager) executeWorkflowJob(ctx context.Context, job *CronJob) {
	// Construct a fake WeixinMessage for workflow execution
	msg := ilink.WeixinMessage{
		FromUserID: job.UserID,
		// Other fields are not used by workflow handler
	}

	clientID := NewClientID()

	// Call workflow handler with the DSL content
	reply := cm.handler.handleWorkflow(ctx, cm.client, msg, job.Command.Content, clientID)

	// Send reply if any
	if reply != "" {
		req := &ilink.SendMessageRequest{
			Msg: ilink.SendMsg{
				FromUserID:   cm.client.BotID(),
				ToUserID:     job.UserID,
				ClientID:     clientID,
				MessageType:  ilink.MessageTypeBot,
				MessageState: ilink.MessageStateFinish,
				ItemList: []ilink.MessageItem{
					{Type: ilink.ItemTypeText, TextItem: &ilink.TextItem{Text: reply}},
				},
			},
		}

		resp, err := cm.client.SendMessage(ctx, req)
		if err != nil || resp.Ret != 0 {
			log.Printf("[cron] failed to send workflow reply: err=%v ret=%d", err, resp.Ret)
		}
	}

	log.Printf("[cron] executed workflow for %s", job.UserID)
}

// executeAgentJob sends a message to a specific agent.
func (cm *CronManager) executeAgentJob(ctx context.Context, job *CronJob) {
	agentName := job.Command.Agent
	if agentName == "" {
		agentName = cm.handler.getDefaultAgentName()
	}

	// This would require calling the agent directly
	// For now, just send the text message
	log.Printf("[cron] agent job not fully implemented, sending as text")
	cm.executeTextJob(ctx, job)
}
