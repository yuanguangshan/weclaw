package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fastclaw-ai/weclaw/ilink"
)

// TimerItem represents a single timer.
type TimerItem struct {
	ID        int    `json:"id"`
	UserID    string `json:"user_id"`
	Label     string `json:"label"`      // user-facing label
	Duration  int64  `json:"duration"`   // requested duration in seconds
	EndTime   int64  `json:"end_time"`   // Unix timestamp when timer expires
	Status    int    `json:"status"`     // 0=running, 1=done, 2=cancelled
	CreatedAt int64  `json:"created_at"`
	Reminded  bool   `json:"reminded"`   // whether expiration notification was sent
}

// TimerStore manages timers in a JSON file.
type TimerStore struct {
	mu       sync.Mutex
	filePath string
	items    []TimerItem
	nextID   int
}

// NewTimerStore creates a timer store backed by a JSON file.
func NewTimerStore(dir string) *TimerStore {
	ts := &TimerStore{
		filePath: filepath.Join(dir, "timers.json"),
	}
	ts.load()
	return ts
}

func (ts *TimerStore) load() {
	data, err := os.ReadFile(ts.filePath)
	if err != nil {
		ts.items = []TimerItem{}
		ts.nextID = 1
		return
	}
	var items []TimerItem
	if err := json.Unmarshal(data, &items); err != nil {
		log.Printf("[timer] failed to parse %s: %v", ts.filePath, err)
		ts.items = []TimerItem{}
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

func (ts *TimerStore) save() error {
	data, err := json.MarshalIndent(ts.items, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(ts.filePath, data, 0644)
}

// Add creates a new timer and returns its ID.
func (ts *TimerStore) Add(userID, label string, duration, endTime int64) int {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	item := TimerItem{
		ID:        ts.nextID,
		UserID:    userID,
		Label:     label,
		Duration:  duration,
		EndTime:   endTime,
		Status:    0,
		CreatedAt: time.Now().Unix(),
	}
	ts.nextID++
	ts.items = append(ts.items, item)
	ts.save()
	return item.ID
}

// List returns running timers for a user, sorted by end time.
func (ts *TimerStore) List(userID string) []TimerItem {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	var result []TimerItem
	for _, item := range ts.items {
		if item.UserID == userID && item.Status == 0 {
			result = append(result, item)
		}
	}
	// Sort by end time (soonest first)
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if result[j].EndTime < result[i].EndTime {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	return result
}

// Cancel marks a timer as cancelled.
func (ts *TimerStore) Cancel(userID string, id int) (string, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	for i := range ts.items {
		if ts.items[i].UserID == userID && ts.items[i].ID == id {
			if ts.items[i].Status != 0 {
				return "", fmt.Errorf(" #%d 已经结束或取消了", id)
			}
			ts.items[i].Status = 2
			label := ts.items[i].Label
			ts.save()
			return label, nil
		}
	}
	return "", fmt.Errorf("没有找到 #%d", id)
}

// Clear removes all running timers for a user.
func (ts *TimerStore) Clear(userID string) int {
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

// GetExpiredTimers returns running timers that have passed their end time
// and haven't been reminded yet.
func (ts *TimerStore) GetExpiredTimers() []TimerItem {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	now := time.Now().Unix()
	var result []TimerItem
	for _, item := range ts.items {
		if item.Status == 0 && !item.Reminded && item.EndTime > 0 && item.EndTime <= now {
			result = append(result, item)
		}
	}
	return result
}

// MarkReminded marks a timer as reminded.
func (ts *TimerStore) MarkReminded(id int) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	for i := range ts.items {
		if ts.items[i].ID == id {
			ts.items[i].Reminded = true
			ts.items[i].Status = 1 // mark as done
			ts.save()
			return
		}
	}
}

// parseTimerDuration parses a natural language duration string.
// Supports: pure numbers (minutes), "Xm", "Xh", "XhXm", "Xmin", "Xhour(s)".
// Returns duration in seconds and remaining text (label).
func parseTimerDuration(text string) (int64, string) {
	text = strings.TrimSpace(text)
	if text == "" {
		return 0, ""
	}

	// Check if it starts with a number
	var numStr strings.Builder
	i := 0
	for i < len(text) && (text[i] >= '0' && text[i] <= '9' || text[i] == '.') {
		numStr.WriteByte(text[i])
		i++
	}

	if numStr.Len() == 0 {
		return 0, text // no number found, return as label for AI parsing
	}

	num, err := strconv.ParseFloat(numStr.String(), 64)
	if err != nil {
		return 0, text
	}

	rest := strings.TrimSpace(text[i:])

	// Parse unit
	lower := strings.ToLower(rest)
	var seconds int64

	if strings.HasPrefix(lower, "h") {
		seconds = int64(num * 3600)
		// Skip the "h" and any trailing characters like "our" / "ours"
		label := skipUnit(rest, "h", "hour", "hours")
		return seconds, label
	} else if strings.HasPrefix(lower, "m") {
		seconds = int64(num * 60)
		label := skipUnit(rest, "m", "min", "mins", "minute", "minutes")
		return seconds, label
	} else if strings.HasPrefix(lower, "s") {
		seconds = int64(num)
		label := skipUnit(rest, "s", "sec", "secs", "second", "seconds")
		return seconds, label
	}

	// No unit — default to minutes
	seconds = int64(num * 60)
	return seconds, rest
}

// skipUnit strips a known unit prefix from text and returns the remaining label.
// Units are tried in descending length order (longest match first).
func skipUnit(text string, units ...string) string {
	// Sort by length descending for longest match
	for longest := 20; longest >= 1; longest-- {
		for _, u := range units {
			if len(u) == longest && strings.HasPrefix(strings.ToLower(text), u) {
				return strings.TrimSpace(text[len(u):])
			}
		}
	}
	return text
}

// handleTimer processes /timer commands.
func (h *Handler) handleTimer(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage, trimmed, clientID string) string {
	rest := strings.TrimSpace(strings.TrimPrefix(trimmed, "/timer"))

	if rest == "" || rest == "list" || rest == "ls" {
		return h.formatTimerList(msg.FromUserID)
	}

	parts := strings.Fields(rest)
	sub := parts[0]

	switch sub {
	case "cancel", "stop", "rm":
		if len(parts) < 2 {
			return "用法: /timer cancel <编号>"
		}
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			return "编号必须是数字"
		}
		label, err := h.timerStore.Cancel(msg.FromUserID, id)
		if err != nil {
			return err.Error()
		}
		return fmt.Sprintf("🛑 已取消 #%d: %s", id, label)

	case "clear":
		n := h.timerStore.Clear(msg.FromUserID)
		if n == 0 {
			return "⏱ 没有进行中的计时器"
		}
		return fmt.Sprintf("🛑 已取消 %d 个计时器", n)

	default:
		return h.createTimer(ctx, msg.FromUserID, rest)
	}
}

func (h *Handler) formatTimerList(userID string) string {
	items := h.timerStore.List(userID)
	if len(items) == 0 {
		return "⏱ 没有进行中的计时器"
	}

	var sb strings.Builder
	sb.WriteString("⏱ **计时器**\n\n")
	for _, item := range items {
		remaining := time.Until(time.Unix(item.EndTime, 0))
		var remainStr string
		if remaining <= 0 {
			remainStr = "已到期"
		} else {
			remainStr = formatDuration(remaining)
		}
		endTime := time.Unix(item.EndTime, 0).Format("15:04")
		sb.WriteString(fmt.Sprintf("#%d. %s — 剩余 %s (截止 %s)\n", item.ID, item.Label, remainStr, endTime))
	}
	return sb.String()
}

func (h *Handler) createTimer(ctx context.Context, userID, text string) string {
	var seconds int64
	var label string

	// Try direct parsing first (fast path)
	seconds, label = parseTimerDuration(text)

	// If direct parsing found a duration, use it
	if seconds > 0 {
		if label == "" {
			label = "计时器"
		}
	} else {
		// Slow path: use AI to parse natural language time
		ag := h.getDefaultAgent()
		if ag != nil {
			prompt := fmt.Sprintf(`从下面这句话中提取倒计时时间和标签。只返回JSON，格式：{"seconds": 300, "label": "标签"}。
seconds 是总倒计时秒数（整数）。如果没有明确时间，seconds设为0。只输出JSON，不要其他文字。
句子：%s`, text)

			reply, err := ag.Chat(ctx, userID+"_timer", prompt)
			if err == nil {
				var parsed struct {
					Seconds int64  `json:"seconds"`
					Label   string `json:"label"`
				}
				reply = strings.TrimSpace(reply)
				if idx := strings.Index(reply, "{"); idx >= 0 {
					reply = reply[idx:]
				}
				if err := json.Unmarshal([]byte(reply), &parsed); err == nil && parsed.Seconds > 0 {
					seconds = parsed.Seconds
					if parsed.Label != "" {
						label = parsed.Label
					}
				}
			}
		}

		// Fallback: couldn't parse
		if seconds <= 0 {
			return "❓ 无法解析时间。用法:\n  /timer 25\n  /timer 2h 写报告\n  /timer 30m 休息"
		}
	}

	// Validate: don't allow timers longer than 24 hours
	if seconds > 86400 {
		return "❓ 计时器不能超过 24 小时"
	}

	endTime := time.Now().Unix() + seconds
	id := h.timerStore.Add(userID, label, seconds, endTime)

	durationStr := formatDuration(time.Duration(seconds) * time.Second)
	endTimeStr := time.Unix(endTime, 0).Format("15:04")
	return fmt.Sprintf("⏱ 已设置 #%d: %s (%s, 截止 %s)", id, label, durationStr, endTimeStr)
}

// StartTimerScheduler starts a background goroutine that checks for expired timers.
func (h *Handler) StartTimerScheduler(ctx context.Context) {
	if h.timerStore == nil {
		return
	}
	h.mu.RLock()
	hasClients := len(h.clients) > 0
	h.mu.RUnlock()
	if !hasClients {
		return
	}

	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				h.checkTimerExpirations(ctx)
			}
		}
	}()
}

func (h *Handler) checkTimerExpirations(ctx context.Context) {
	items := h.timerStore.GetExpiredTimers()
	if len(items) == 0 {
		return
	}

	h.mu.RLock()
	clients := h.clients
	h.mu.RUnlock()

	for _, item := range items {
		elapsed := time.Duration(item.Duration) * time.Second
		text := fmt.Sprintf("⏰ 计时器 #%d 到期: %s (共 %s)", item.ID, item.Label, formatDuration(elapsed))
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
				log.Printf("[timer] failed to send expiration to %s: err=%v ret=%d", item.UserID, err, resp.Ret)
			} else {
				log.Printf("[timer] sent expiration #%d to %s: %s", item.ID, item.UserID, item.Label)
				h.timerStore.MarkReminded(item.ID)
			}
		}
	}
}
