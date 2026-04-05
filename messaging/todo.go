package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fastclaw-ai/weclaw/ilink"
)

// TodoItem represents a single todo task.
type TodoItem struct {
	ID        int    `json:"id"`
	UserID    string `json:"user_id"`
	Title     string `json:"title"`
	DueTime   int64  `json:"due_time"` // Unix timestamp, 0 = no deadline
	Status    int    `json:"status"`   // 0=pending, 1=done
	CreatedAt int64  `json:"created_at"`
	Reminded  bool   `json:"reminded"` // whether reminder was sent
}

// TodoStore manages todos in a JSON file.
type TodoStore struct {
	mu       sync.Mutex
	filePath string
	items    []TodoItem
	nextID   int
}

// NewTodoStore creates a todo store backed by a JSON file.
func NewTodoStore(dir string) *TodoStore {
	ts := &TodoStore{
		filePath: filepath.Join(dir, "todos.json"),
	}
	ts.load()
	return ts
}

func (ts *TodoStore) load() {
	data, err := os.ReadFile(ts.filePath)
	if err != nil {
		ts.items = []TodoItem{}
		ts.nextID = 1
		return
	}
	var items []TodoItem
	if err := json.Unmarshal(data, &items); err != nil {
		log.Printf("[todo] failed to parse %s: %v", ts.filePath, err)
		ts.items = []TodoItem{}
		ts.nextID = 1
		return
	}
	ts.items = items
	ts.nextID = 1
	for _, item := range items {
		if item.ID >= ts.nextID {
			ts.nextID = item.ID + 1
		}
	}
}

func (ts *TodoStore) save() error {
	data, err := json.MarshalIndent(ts.items, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(ts.filePath, data, 0644)
}

// Add creates a new todo and returns its ID.
func (ts *TodoStore) Add(userID, title string, dueTime int64) int {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	item := TodoItem{
		ID:        ts.nextID,
		UserID:    userID,
		Title:     title,
		DueTime:   dueTime,
		Status:    0,
		CreatedAt: time.Now().Unix(),
	}
	ts.nextID++
	ts.items = append(ts.items, item)
	ts.save()
	return item.ID
}

// List returns pending todos for a user, sorted by due time.
func (ts *TodoStore) List(userID string) []TodoItem {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	var result []TodoItem
	for _, item := range ts.items {
		if item.UserID == userID && item.Status == 0 {
			result = append(result, item)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].DueTime == 0 && result[j].DueTime == 0 {
			return result[i].CreatedAt < result[j].CreatedAt
		}
		if result[i].DueTime == 0 {
			return false
		}
		if result[j].DueTime == 0 {
			return true
		}
		return result[i].DueTime < result[j].DueTime
	})
	return result
}

// Done marks a todo as completed.
func (ts *TodoStore) Done(userID string, id int) (string, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	for i := range ts.items {
		if ts.items[i].UserID == userID && ts.items[i].ID == id {
			if ts.items[i].Status == 1 {
				return "", fmt.Errorf(" #%d 已经完成了", id)
			}
			ts.items[i].Status = 1
			title := ts.items[i].Title
			ts.save()
			return title, nil
		}
	}
	return "", fmt.Errorf("没有找到 #%d", id)
}

// Delete removes a todo.
func (ts *TodoStore) Delete(userID string, id int) (string, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	for i := range ts.items {
		if ts.items[i].UserID == userID && ts.items[i].ID == id {
			title := ts.items[i].Title
			ts.items = append(ts.items[:i], ts.items[i+1:]...)
			ts.save()
			return title, nil
		}
	}
	return "", fmt.Errorf("没有找到 #%d", id)
}

// Clear removes all todos for a user and returns the count removed.
func (ts *TodoStore) Clear(userID string) int {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	n := 0
	remaining := ts.items[:0]
	for _, item := range ts.items {
		if item.UserID == userID && item.Status == 0 {
			n++
		} else {
			remaining = append(remaining, item)
		}
	}
	ts.items = remaining
	ts.save()
	return n
}

// GetDueReminders returns pending todos that are due within the next minute
// and haven't been reminded yet.
func (ts *TodoStore) GetDueReminders() []TodoItem {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	now := time.Now().Unix()
	var result []TodoItem
	for _, item := range ts.items {
		if item.Status == 0 && !item.Reminded && item.DueTime > 0 && item.DueTime <= now+60 {
			result = append(result, item)
		}
	}
	return result
}

// MarkReminded marks a todo as reminded.
func (ts *TodoStore) MarkReminded(id int) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	for i := range ts.items {
		if ts.items[i].ID == id {
			ts.items[i].Reminded = true
			ts.save()
			return
		}
	}
}

// handleTodo processes /todo commands.
func (h *Handler) handleTodo(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage, trimmed, clientID string) string {
	var reply string
	rest := strings.TrimSpace(strings.TrimPrefix(trimmed, "/todo"))

	if rest == "" || rest == "list" || rest == "ls" {
		return h.formatTodoList(msg.FromUserID)
	}

	parts := strings.Fields(rest)
	sub := parts[0]

	switch sub {
	case "done", "ok", "finish":
		if len(parts) < 2 {
			reply = "用法: /todo done <编号>"
		}
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			reply = "编号必须是数字"
		}
		title, err := h.todoStore.Done(msg.FromUserID, id)
		if err != nil {
			return err.Error()
		}
		return fmt.Sprintf("✅ 已完成 #%d: %s", id, title)

	case "del", "rm", "delete":
		if len(parts) < 2 {
			reply = "用法: /todo del <编号>"
		}
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			reply = "编号必须是数字"
		}
		title, err := h.todoStore.Delete(msg.FromUserID, id)
		if err != nil {
			return err.Error()
		}
		return fmt.Sprintf("🗑 已删除 #%d: %s", id, title)

	case "clear":
		n := h.todoStore.Clear(msg.FromUserID)
		if n == 0 {
			return "📋 没有待办事项"
		}
		return fmt.Sprintf("🗑 已清空 %d 条待办事项", n)

	default:
		// Create new todo: try to parse time from the text using agent
		reply = h.createTodo(ctx, msg.FromUserID, rest)
	}

	return reply
}

func (h *Handler) formatTodoList(userID string) string {
	items := h.todoStore.List(userID)
	if len(items) == 0 {
		return "📋 待办清单为空"
	}

	var sb strings.Builder
	sb.WriteString("📋 **待办清单**\n\n")
	for _, item := range items {
		dueStr := ""
		if item.DueTime > 0 {
			due := time.Unix(item.DueTime, 0)
			diff := time.Until(due)
			if diff < 0 {
				dueStr = fmt.Sprintf(" (⚠️ 已超时 %s)", formatDuration(-diff))
			} else if diff < time.Hour {
				dueStr = fmt.Sprintf(" (%d分钟后)", int(diff.Minutes()))
			} else {
				dueStr = fmt.Sprintf(" (%s)", due.Format("01-02 15:04"))
			}
		}
		sb.WriteString(fmt.Sprintf("#%d. %s%s\n", item.ID, item.Title, dueStr))
	}
	return sb.String()
}

func (h *Handler) createTodo(ctx context.Context, userID, text string) string {
	// Try to extract time using the default agent
	var dueTime int64
	var title string

	ag := h.getDefaultAgent()
	if ag != nil {
		prompt := fmt.Sprintf(`从下面这句话中提取时间和待办事项。只返回JSON，格式：{"time": "YYYY-MM-DD HH:MM:SS", "title": "事项"}。
如果没有明确时间，time设为空字符串。只输出JSON，不要其他文字。
句子：%s`, text)

		reply, err := ag.Chat(ctx, userID+"_todo", prompt)
		if err == nil {
			var parsed struct {
				Time  string `json:"time"`
				Title string `json:"title"`
			}
			// Try to extract JSON from reply
			reply = strings.TrimSpace(reply)
			if idx := strings.Index(reply, "{"); idx >= 0 {
				reply = reply[idx:]
			}
			if err := json.Unmarshal([]byte(reply), &parsed); err == nil && parsed.Title != "" {
				title = parsed.Title
				if parsed.Time != "" {
					if t, err := time.ParseInLocation("2006-01-02 15:04:05", parsed.Time, time.Local); err == nil {
						dueTime = t.Unix()
					}
				}
			}
		}
	}

	// Fallback: use raw text as title
	if title == "" {
		title = text
	}

	id := h.todoStore.Add(userID, title, dueTime)
	if dueTime > 0 {
		due := time.Unix(dueTime, 0)
		return fmt.Sprintf("✅ 已添加 #%d: %s (截止 %s)", id, title, due.Format("01-02 15:04"))
	}
	return fmt.Sprintf("✅ 已添加 #%d: %s", id, title)
}

// formatDuration returns a human-readable duration string.
func formatDuration(d time.Duration) string {
	if d < time.Minute {
		return "刚刚"
	}
	if d < time.Hour {
		return fmt.Sprintf("%d分钟", int(d.Minutes()))
	}
	if d < 24*time.Hour {
		return fmt.Sprintf("%d小时", int(d.Hours()))
	}
	return fmt.Sprintf("%d天", int(d.Hours()/24))
}

// StartTodoScheduler starts a background goroutine that checks for due reminders.
func (h *Handler) StartTodoScheduler(ctx context.Context) {
	if h.todoStore == nil {
		return
	}
	h.mu.RLock()
	hasClients := len(h.clients) > 0
	h.mu.RUnlock()
	if !hasClients {
		return
	}

	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				h.checkReminders(ctx)
			}
		}
	}()
}

func (h *Handler) checkReminders(ctx context.Context) {
	items := h.todoStore.GetDueReminders()
	if len(items) == 0 {
		return
	}

	h.mu.RLock()
	clients := h.clients
	h.mu.RUnlock()

	for _, item := range items {
		text := fmt.Sprintf("⏰ 提醒: %s", item.Title)
		if item.DueTime > 0 {
			due := time.Unix(item.DueTime, 0)
			text += fmt.Sprintf(" (截止 %s)", due.Format("15:04"))
		}
		cid := NewClientID()
		plainText := MarkdownToPlainText(text)

		for _, client := range clients {
			req := &ilink.SendMessageRequest{
				Msg: ilink.SendMsg{
					FromUserID:   client.BotID(),
					ToUserID:     item.UserID,
					ClientID:     cid,
					MessageType:  ilink.MessageTypeBot,
					MessageState: ilink.MessageStateFinish,
					ItemList: []ilink.MessageItem{
						{Type: ilink.ItemTypeText, TextItem: &ilink.TextItem{Text: plainText}},
					},
				},
			}
			resp, err := client.SendMessage(ctx, req)
			if err != nil || resp.Ret != 0 {
				log.Printf("[todo] failed to send reminder to %s: err=%v ret=%d", item.UserID, err, resp.Ret)
			} else {
				log.Printf("[todo] sent reminder #%d to %s: %s", item.ID, item.UserID, item.Title)
				h.todoStore.MarkReminded(item.ID)
			}
		}
	}
}
