package messaging

import (
	"bytes"
	"context"
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fastclaw-ai/weclaw/agent"
	"github.com/fastclaw-ai/weclaw/hub"
	"github.com/fastclaw-ai/weclaw/ilink"
	"github.com/google/uuid"
)

// AgentFactory creates an agent by config name. Returns nil if the name is unknown.
type AgentFactory func(ctx context.Context, name string) agent.Agent

// SaveDefaultFunc persists the default agent name to config file.
type SaveDefaultFunc func(name string) error

// AgentMeta holds static config info about an agent (for /status display).
type AgentMeta struct {
	Name    string
	Type    string // "acp", "cli", "http"
	Command string // binary path or endpoint
	Model   string
}

// Handler processes incoming WeChat messages and dispatches replies.
type Handler struct {
	mu              sync.RWMutex
	defaultName     string
	agents          map[string]agent.Agent // name -> running agent
	agentMetas      []AgentMeta            // all configured agents (for /status)
	agentWorkDirs   map[string]string      // agent name -> configured/runtime cwd
	customAliases   map[string]string      // custom alias -> agent name (from config)
	factory         AgentFactory
	saveDefault     SaveDefaultFunc
	hub             *hub.Hub         // shared context for cross-agent collaboration
	contextTokens   sync.Map         // map[userID]contextToken
	saveDir         string           // directory to save images/files to
	seenMsgs        sync.Map         // map[int64]time.Time — dedup by message_id
	progressCtx     *progressContext // current request context for progress notifications
	lastReplies     sync.Map         // map[userID]string — last agent reply per user (for /save without message)
	shellModeStates sync.Map         // map[userID]*shellModeState — per-user shell mode state
	todoStore       *TodoStore
	timerStore      *TimerStore
	cronManager     *CronManager
	clients         []*ilink.Client
	// Remote clipboard configuration for sending AI replies to remote endpoint
	remoteClipboardURL string
	remoteClipboardKey string
	// Relay configuration for relaying Q&A pairs and disconnection notices
	relayURL     string
	relayAuthKey string
}

// progressContext holds context for sending progress notifications.
type progressContext struct {
	client   *ilink.Client
	userID   string
	token    string
	cancel   context.CancelFunc
	lastTime time.Time // last progress notification time
	mu       sync.Mutex
}

// shellModeState holds per-user shell mode state.
type shellModeState struct {
	enabled bool   // whether shell mode is active
	cwd     string // current working directory
	baseDir string // base directory for path sandboxing (empty = no restriction)
}

// NewHandler creates a new message handler.
func NewHandler(factory AgentFactory, saveDefault SaveDefaultFunc) *Handler {
	return &Handler{
		agents:        make(map[string]agent.Agent),
		agentWorkDirs: make(map[string]string),
		factory:       factory,
		saveDefault:   saveDefault,
		hub:           hub.New(hub.DefaultDir()),
		todoStore:     NewTodoStore(hub.DefaultDir()),
		timerStore:    NewTimerStore(hub.DefaultDir()),
	}
}

// SetHub sets a custom Hub instance (for testing or custom paths).
func (h *Handler) SetHub(hu *hub.Hub) {
	h.hub = hu
}

// SetSaveDir sets the directory for saving images and files.
func (h *Handler) SetSaveDir(dir string) {
	h.saveDir = dir
}

// SetRemoteClipboard sets the remote clipboard endpoint configuration.
func (h *Handler) SetRemoteClipboard(url, key string) {
	h.remoteClipboardURL = url
	h.remoteClipboardKey = key
}

// SetRelay sets the relay endpoint configuration for Q&A pairs.
func (h *Handler) SetRelay(url, key string) {
	h.relayURL = url
	h.relayAuthKey = key
}

// cleanSeenMsgs removes entries older than 5 minutes from the dedup cache.
func (h *Handler) cleanSeenMsgs() {
	cutoff := time.Now().Add(-5 * time.Minute)
	h.seenMsgs.Range(func(key, value any) bool {
		if t, ok := value.(time.Time); ok && t.Before(cutoff) {
			h.seenMsgs.Delete(key)
		}
		return true
	})
}

// SetCustomAliases sets custom alias mappings from config.
func (h *Handler) SetCustomAliases(aliases map[string]string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.customAliases = aliases
}

func (h *Handler) SetClients(clients []*ilink.Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients = clients
}

// SetAgentMetas sets the list of all configured agents (for /status).
func (h *Handler) SetAgentMetas(metas []AgentMeta) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.agentMetas = metas
}

// SetAgentWorkDirs sets the configured working directory for each agent.
func (h *Handler) SetAgentWorkDirs(workDirs map[string]string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.agentWorkDirs = make(map[string]string, len(workDirs))
	for name, dir := range workDirs {
		h.agentWorkDirs[name] = dir
	}
}

// SetCronManager sets the cron manager for scheduled tasks.
func (h *Handler) SetCronManager(cm *CronManager) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.cronManager = cm
}

// SetDefaultAgent sets the default agent (already started).
func (h *Handler) SetDefaultAgent(name string, ag agent.Agent) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.defaultName = name
	h.agents[name] = ag
	log.Printf("[handler] default agent ready: %s (%s)", name, ag.Info())
}

// getAgent returns a running agent by name, or starts it on demand via factory.
func (h *Handler) getAgent(ctx context.Context, name string) (agent.Agent, error) {
	// Fast path: already running
	h.mu.RLock()
	ag, ok := h.agents[name]
	h.mu.RUnlock()
	if ok {
		return ag, nil
	}

	// Slow path: create on demand
	if h.factory == nil {
		return nil, fmt.Errorf("agent %q not found and no factory configured", name)
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	// Double-check after acquiring write lock
	if ag, ok := h.agents[name]; ok {
		return ag, nil
	}

	log.Printf("[handler] starting agent %q on demand...", name)
	ag = h.factory(ctx, name)
	if ag == nil {
		return nil, fmt.Errorf("agent %q not available", name)
	}

	h.agents[name] = ag
	log.Printf("[handler] agent started on demand: %s (%s)", name, ag.Info())
	return ag, nil
}

// getDefaultAgent returns the default agent (may be nil if not ready yet).
func (h *Handler) getDefaultAgent() agent.Agent {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.defaultName == "" {
		return nil
	}
	return h.agents[h.defaultName]
}

// isKnownAgent checks if a name corresponds to a configured agent.
func (h *Handler) isKnownAgent(name string) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	// Check running agents
	if _, ok := h.agents[name]; ok {
		return true
	}
	// Check configured agents (metas)
	for _, meta := range h.agentMetas {
		if meta.Name == name {
			return true
		}
	}
	return false
}

// agentAliases maps short aliases to agent config names.
var agentAliases = map[string]string{
	"cc":  "claude",
	"cx":  "codex",
	"oc":  "openclaw",
	"cs":  "cursor",
	"km":  "kimi",
	"gm":  "gemini",
	"ocd": "opencode",
	"pi":  "pi",
	"cp":  "copilot",
	"dr":  "droid",
	"if":  "iflow",
	"kr":  "kiro",
	"qw":  "qwen",
}

// resolveAlias returns the full agent name for an alias, or the original name if no alias matches.
// Checks custom aliases (from config) first, then built-in aliases.
// Matching is case-insensitive for aliases, and also falls back to case-insensitive agent name matching.
func (h *Handler) resolveAlias(name string) string {
	lower := strings.ToLower(name)
	h.mu.RLock()
	custom := h.customAliases
	h.mu.RUnlock()
	if custom != nil {
		for alias, full := range custom {
			if strings.ToLower(alias) == lower {
				return full
			}
		}
	}
	for alias, full := range agentAliases {
		if strings.ToLower(alias) == lower {
			return full
		}
	}
	return name
}

// isBuiltinCommand returns true if the text starts with a built-in weclaw command.
// These should NOT be parsed as agent name prefixes.
func isBuiltinCommand(text string) bool {
	for _, cmd := range []string{"/help", "/info", "/new", "/clear", "/cwd", "/save", "/hub", "/sh", "/$", "/q", "/podcast", "/debate", "/todo", "/timer", "/workflow", "/cron"} {
		if strings.HasPrefix(text, cmd) {
			// Make sure it's the command itself, not an agent name that starts with "help" etc.
			// e.g. "/helpful stuff" should not match, but "/help", "/help " and "/help\n" should
			rest := strings.TrimPrefix(text, cmd)
			if rest == "" {
				return true
			}
			// Check if next character is whitespace (space, tab, newline, etc.)
			// This handles both single-line and multi-line command formats
			nextChar := rest[0]
			return nextChar == ' ' || nextChar == '\t' || nextChar == '\n' || nextChar == '\r'
		}
	}
	return false
}

// parseCommand checks if text starts with "/" or "@" followed by agent name(s).
// Supports multiple agents: "@cc @cx hello" returns (["claude","codex"], "hello").
// Returns (agentNames, actualMessage). Aliases are resolved automatically.
// If no command prefix, returns (nil, originalText).
// Built-in commands (/help, /save, /hub, etc.) are NOT parsed as agent names.
func (h *Handler) parseCommand(text string) ([]string, string) {
	if !strings.HasPrefix(text, "/") && !strings.HasPrefix(text, "@") {
		return nil, text
	}

	// Don't parse built-in commands as agent prefixes
	trimmed := strings.TrimSpace(text)
	if isBuiltinCommand(trimmed) {
		return nil, text
	}

	// Parse consecutive @name or /name tokens from the start
	var names []string
	rest := text
	for {
		rest = strings.TrimSpace(rest)
		if !strings.HasPrefix(rest, "/") && !strings.HasPrefix(rest, "@") {
			break
		}

		// Save original rest before parsing this token (needed if it's a builtin command)
		originalRest := rest

		// Strip prefix
		after := rest[1:]
		idx := strings.IndexAny(after, " /@")
		var token string
		if idx < 0 {
			// Rest is just the name, no message
			token = after
			rest = ""
		} else if after[idx] == '/' || after[idx] == '@' {
			// Next token is another @name or /name
			token = after[:idx]
			rest = after[idx:]
		} else {
			// Space — name ends here
			token = after[:idx]
			rest = strings.TrimSpace(after[idx+1:])
		}

		if token != "" {
			// Don't parse built-in commands as agent names
			if isBuiltinCommand("/" + token) {
				// Keep the built-in command in rest so it can be matched by the router
				rest = originalRest
				break
			}
			names = append(names, h.resolveAlias(token))
		}

		if rest == "" {
			break
		}
	}

	// Deduplicate names preserving order
	seen := make(map[string]bool)
	unique := names[:0]
	for _, n := range names {
		if !seen[n] {
			seen[n] = true
			unique = append(unique, n)
		}
	}

	return unique, rest
}

// HandleMessage processes a single incoming message.
func (h *Handler) HandleMessage(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage) {
	// Only process user messages that are finished
	if msg.MessageType != ilink.MessageTypeUser {
		return
	}
	if msg.MessageState != ilink.MessageStateFinish {
		return
	}

	// Deduplicate by message_id to avoid processing the same message multiple times
	// (voice messages may trigger multiple finish-state updates)
	if msg.MessageID != 0 {
		if _, loaded := h.seenMsgs.LoadOrStore(msg.MessageID, time.Now()); loaded {
			return
		}
		// Clean up old entries periodically (fire-and-forget)
		go h.cleanSeenMsgs()
	}

	// Extract text from item list (text message or voice transcription)
	text := extractText(msg)
	if text == "" {
		if voiceText := extractVoiceText(msg); voiceText != "" {
			text = voiceText
			log.Printf("[handler] voice transcription from %s: %q", msg.FromUserID, truncate(text, 80))
		}
	}

	// Check for media attachments (image, file, video) — regardless of whether text exists
	media := h.extractAllMedia(ctx, client, msg)
	if len(media) > 0 {
		log.Printf("[handler] extracted %d media items from message (text=%q)", len(media), truncate(text, 40))
		h.sendMediaToAgent(ctx, client, msg, text, media)
		return
	}

	if text == "" {
		log.Printf("[handler] received non-text message from %s, skipping", msg.FromUserID)
		return
	}

	log.Printf("[handler] received from %s: %q", msg.FromUserID, truncate(text, 80))

	// Store context token for this user
	h.contextTokens.Store(msg.FromUserID, msg.ContextToken)

	// Generate a clientID for this reply (used to correlate typing → finish)
	clientID := NewClientID()

	// Check if user is in shell mode
	if state, ok := h.shellModeStates.Load(msg.FromUserID); ok && state != nil {
		sm := state.(*shellModeState)
		if sm.enabled {
			trimmed := strings.TrimSpace(text)
			// Exit shell mode
			if trimmed == "/q" {
				sm.enabled = false
				reply := "已退出命令行模式"
				if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
					log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
				}
				return
			}
			// Execute command in shell mode
			reply := h.handleShellWithState(ctx, sm, trimmed)
			if reply != "" {
				if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
					log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
				}
			}
			return
		}
	}

	// Intercept URLs: save to Linkhoard directly without AI agent
	trimmed := strings.TrimSpace(text)
	if h.saveDir != "" && IsURL(trimmed) {
		rawURL := ExtractURL(trimmed)
		if rawURL != "" {
			log.Printf("[handler] saving URL to linkhoard: %s", rawURL)
			meta, err := SaveLinkToLinkhoard(ctx, h.saveDir, rawURL)
			var reply string
			if err != nil {
				log.Printf("[handler] link save failed: %v", err)
				reply = fmt.Sprintf("保存失败: %v", err)
			} else {
				reply = fmt.Sprintf("已保存: %s", meta.Title)
				// If it's a WeChat article, send to nanobot for analysis
				if isWeChatURL(rawURL) {
					go h.analyzeWithNanobot(ctx, client, msg, meta)
				}
			}
			if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
				log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
			}
			return
		}
	}

	// Pre-parse agent prefix so "@agent /hub ..." and "@agent /save ..." work correctly.
	// Without this, "/hub" check on trimmed (which starts with "@agent") would fail,
	// causing the command to be forwarded raw to the agent instead of being handled by weclaw.
	parsedAgentNames, parsedMessage := h.parseCommand(text)

	// Build effective trimmed (strip agent prefix if present)
	effectiveTrimmed := trimmed
	if len(parsedAgentNames) > 0 {
		effectiveTrimmed = strings.TrimSpace(parsedMessage)
	}

	// Built-in commands (no typing needed)
handleBuiltinCommand:
	if effectiveTrimmed == "/info" {
		reply := h.buildStatus()
		if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
			log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
		}
		return
	} else if effectiveTrimmed == "/help" {
		reply := buildHelpText()
		if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
			log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
		}
		return
	} else if effectiveTrimmed == "/new" || effectiveTrimmed == "/clear" {
		reply := h.resetDefaultSession(ctx, msg.FromUserID)
		if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
			log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
		}
		return
	} else if strings.HasPrefix(effectiveTrimmed, "/cwd") {
		reply := h.handleCwd(effectiveTrimmed)
		if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
			log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
		}
		return
	} else if strings.HasPrefix(effectiveTrimmed, "/save") {
		// Reconstruct trimmed to include agent prefix for handleSave parsing
		// handleSave expects "/save @agent filename message" or "/save filename message"
		saveTrimmed := effectiveTrimmed
		if len(parsedAgentNames) > 0 {
			saveTrimmed = "/save @" + parsedAgentNames[0] + " " + strings.TrimPrefix(effectiveTrimmed, "/save")
		}
		reply := h.handleSave(ctx, client, msg, strings.TrimSpace(saveTrimmed), clientID)
		if reply != "" {
			if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
				log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
			}
		}
		return
	} else if strings.HasPrefix(effectiveTrimmed, "/hub") {
		// Reconstruct trimmed to include agent prefix for handleHub parsing
		// handleHub expects "/hub @agent filename message" or "/hub filename message"
		hubTrimmed := effectiveTrimmed
		if len(parsedAgentNames) > 0 {
			hubTrimmed = "/hub @" + parsedAgentNames[0] + " " + strings.TrimPrefix(effectiveTrimmed, "/hub")
		}
		reply := h.handleHub(ctx, client, msg, strings.TrimSpace(hubTrimmed), clientID)
		if reply != "" {
			if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
				log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
			}
		}
		return
	} else if strings.HasPrefix(effectiveTrimmed, "/podcast") {
		reply := h.handlePodcast(ctx, client, msg, effectiveTrimmed, clientID)
		if reply != "" {
			if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
				log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
			}
		}
		return
	} else if strings.HasPrefix(effectiveTrimmed, "/todo") {
		reply := h.handleTodo(ctx, client, msg, effectiveTrimmed, clientID)
		if reply != "" {
			if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
				log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
			}
		}
		return
	} else if strings.HasPrefix(effectiveTrimmed, "/timer") {
		reply := h.handleTimer(ctx, client, msg, effectiveTrimmed, clientID)
		if reply != "" {
			if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
				log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
			}
		}
		return
	} else if strings.HasPrefix(effectiveTrimmed, "/debate") {
		reply := h.handleDebate(ctx, client, msg, effectiveTrimmed, clientID)
		if reply != "" {
			if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
				log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
			}
		}
		return
	} else if strings.HasPrefix(effectiveTrimmed, "/workflow") {
		reply := h.handleWorkflow(ctx, client, msg, effectiveTrimmed, clientID)
		if reply != "" {
			if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
				log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
			}
		}
		return
	} else if effectiveTrimmed == "/sh" || effectiveTrimmed == "/$" {
		// Enter shell mode
		reply := h.enterShellMode(ctx, msg.FromUserID)
		if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
			log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
		}
		return
	} else if strings.HasPrefix(effectiveTrimmed, "/sh ") || strings.HasPrefix(effectiveTrimmed, "/$ ") {
		// Execute single command without entering shell mode
		reply := h.handleShell(ctx, effectiveTrimmed)
		if reply != "" {
			if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
				log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
			}
		}
		return
	} else if strings.HasPrefix(effectiveTrimmed, "/cron") {
		// Handle cron commands
		reply := h.handleCron(ctx, client, msg, effectiveTrimmed, clientID)
		if reply != "" {
			if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
				log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
			}
		}
		return
	}

	// Route: "/agentname message" or "@agent1 @agent2 message" -> specific agent(s)
	// Reuse pre-parsed values from above
	agentNames := parsedAgentNames
	message := parsedMessage

	// No command prefix -> send to default agent
	if len(agentNames) == 0 {
		h.sendToDefaultAgent(ctx, client, msg, text, clientID)
		return
	}

	// No message -> switch default agent (only first name)
	if message == "" {
		if len(agentNames) == 1 && h.isKnownAgent(agentNames[0]) {
			reply := h.switchDefault(ctx, agentNames[0])
			if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
				log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
			}
		} else if len(agentNames) == 1 && !h.isKnownAgent(agentNames[0]) {
			// Unknown agent -> forward to default
			h.sendToDefaultAgent(ctx, client, msg, text, clientID)
		} else {
			reply := "Usage: specify one agent to switch, or add a message to broadcast"
			if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
				log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
			}
		}
		return
	}

	// Filter to known agents; if single unknown agent -> forward to default
	var knownNames []string
	for _, name := range agentNames {
		if h.isKnownAgent(name) {
			knownNames = append(knownNames, name)
		}
	}
	if len(knownNames) == 0 {
		// No known agents — check if the remaining message is a built-in command
		// e.g. "@gpt /hub ..." should be treated as "/hub ..."
		restMsg := strings.TrimSpace(parsedMessage)
		if isBuiltinCommand(restMsg) {
			effectiveTrimmed = restMsg
			goto handleBuiltinCommand
		}
		// Forward entire text to default agent
		h.sendToDefaultAgent(ctx, client, msg, text, clientID)
		return
	}

	// Send typing indicator
	go func() {
		if typingErr := SendTypingState(ctx, client, msg.FromUserID, msg.ContextToken); typingErr != nil {
			log.Printf("[handler] failed to send typing state: %v", typingErr)
		}
	}()

	if len(knownNames) == 1 {
		// Single agent
		h.sendToNamedAgent(ctx, client, msg, knownNames[0], message, clientID)
	} else {
		// Multi-agent broadcast: parallel dispatch, send replies as they arrive
		h.broadcastToAgents(ctx, client, msg, knownNames, message)
	}
}

// sendToDefaultAgent sends the message to the default agent and replies.
func (h *Handler) sendToDefaultAgent(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage, text, clientID string) {
	go func() {
		if typingErr := SendTypingState(ctx, client, msg.FromUserID, msg.ContextToken); typingErr != nil {
			log.Printf("[handler] failed to send typing state: %v", typingErr)
		}
	}()

	h.mu.RLock()
	defaultName := h.defaultName
	h.mu.RUnlock()

	ag := h.getDefaultAgent()
	var reply string
	if ag != nil {
		var err error
		reply, err = h.chatWithAgent(ctx, ag, msg.FromUserID, text, client, msg.ContextToken)
		if err != nil {
			reply = fmt.Sprintf("Error: %v", err)
		}
	} else {
		log.Printf("[handler] agent not ready, using echo mode for %s", msg.FromUserID)
		reply = "[echo] " + text
	}

	if h.relayURL != "" {
		go h.sendToRelay(ctx, defaultName, text, reply)
	}

	h.sendReplyWithMedia(ctx, client, msg, defaultName, reply, clientID)
}

// sendToNamedAgent sends the message to a specific agent and replies.
func (h *Handler) sendToNamedAgent(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage, name, message, clientID string) {
	ag, agErr := h.getAgent(ctx, name)
	if agErr != nil {
		log.Printf("[handler] agent %q not available: %v", name, agErr)
		reply := fmt.Sprintf("Agent %q is not available: %v", name, agErr)
		SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID)
		return
	}

	reply, err := h.chatWithAgent(ctx, ag, msg.FromUserID, message, client, msg.ContextToken)
	if err != nil {
		reply = fmt.Sprintf("Error: %v", err)
	}
	if h.relayURL != "" {
		go h.sendToRelay(ctx, name, message, reply)
	}
	h.sendReplyWithMedia(ctx, client, msg, name, reply, clientID)
}

// broadcastToAgents sends the message to multiple agents in parallel.
// Each reply is sent as a separate message with the agent name prefix.
func (h *Handler) broadcastToAgents(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage, names []string, message string) {
	type result struct {
		name  string
		reply string
	}

	ch := make(chan result, len(names))

	for _, name := range names {
		go func(n string) {
			ag, err := h.getAgent(ctx, n)
			if err != nil {
				ch <- result{name: n, reply: fmt.Sprintf("Error: %v", err)}
				return
			}
			reply, err := h.chatWithAgent(ctx, ag, msg.FromUserID, message, client, msg.ContextToken)
			if err != nil {
				ch <- result{name: n, reply: fmt.Sprintf("Error: %v", err)}
				return
			}
			ch <- result{name: n, reply: reply}
		}(name)
	}

	// Send replies as they arrive
	for range names {
		r := <-ch
		reply := fmt.Sprintf("[%s] %s", r.name, r.reply)
		clientID := NewClientID()
		if h.relayURL != "" {
			go h.sendToRelay(ctx, r.name, message, r.reply)
		}
		h.sendReplyWithMedia(ctx, client, msg, r.name, reply, clientID)
	}
}

// sendReplyWithMedia sends a text reply and any extracted image URLs.
func (h *Handler) sendReplyWithMedia(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage, agentName, reply, clientID string) {
	imageURLs := ExtractImageURLs(reply)
	attachmentPaths := extractLocalAttachmentPaths(reply)
	allowedRoots := h.allowedAttachmentRoots(agentName)

	var sentPaths []string
	var failedPaths []string
	for _, attachmentPath := range attachmentPaths {
		if !isAllowedAttachmentPath(attachmentPath, allowedRoots) {
			log.Printf("[handler] rejected attachment outside allowed roots for agent %q: %s", agentName, attachmentPath)
			failedPaths = append(failedPaths, attachmentPath)
			continue
		}
		if err := SendMediaFromPath(ctx, client, msg.FromUserID, attachmentPath, msg.ContextToken); err != nil {
			log.Printf("[handler] failed to send attachment to %s: %v", msg.FromUserID, err)
			failedPaths = append(failedPaths, attachmentPath)
			continue
		}
		sentPaths = append(sentPaths, attachmentPath)
	}

	reply = rewriteReplyWithAttachmentResults(reply, sentPaths, failedPaths)

	// Cache last reply for /save without message
	h.lastReplies.Store(msg.FromUserID, reply)

	if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
		log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
	}

	for _, imgURL := range imageURLs {
		if err := SendMediaFromURL(ctx, client, msg.FromUserID, imgURL, msg.ContextToken); err != nil {
			log.Printf("[handler] failed to send image to %s: %v", msg.FromUserID, err)
		}
	}
}

func (h *Handler) allowedAttachmentRoots(agentName string) []string {
	roots := []string{defaultAttachmentWorkspace()}

	h.mu.RLock()
	agentDir := h.agentWorkDirs[agentName]
	h.mu.RUnlock()

	if agentDir != "" {
		roots = append(roots, agentDir)
	}

	return roots
}

// chatWithAgent sends a message to an agent and returns the reply, with logging.
// Optional client and token can be provided for progress notifications.
func (h *Handler) chatWithAgent(ctx context.Context, ag agent.Agent, userID, message string, clientAndToken ...interface{}) (string, error) {
	info := ag.Info()
	log.Printf("[handler] dispatching to agent (%s) for %s", info, userID)

	// Set up progress callback if client and token are provided
	if len(clientAndToken) >= 2 {
		if client, ok := clientAndToken[0].(*ilink.Client); ok && client != nil {
			if token, ok := clientAndToken[1].(string); ok && token != "" {
				// Get existing context token for this user
				if contextTokenVal, ok := h.contextTokens.Load(userID); ok && contextTokenVal != nil {
					if contextToken, ok := contextTokenVal.(string); ok {
						// Create progress context
						pCtx := &progressContext{
							client:   client,
							userID:   userID,
							token:    contextToken,
							lastTime: time.Time{}, // zero time means no notification sent yet
						}

						// Set progress callback on the agent
						ag.SetProgressCallback(func(ctx context.Context, event agent.ProgressEvent) {
							h.handleProgressEvent(ctx, pCtx, event)
						})

						// Clean up progress context after chat completes
						defer func() {
							h.setProgressContext(nil)
						}()
					}
				}
			}
		}
	}

	start := time.Now()
	reply, err := ag.Chat(ctx, userID, message)
	elapsed := time.Since(start)

	if err != nil {
		log.Printf("[handler] agent error (%s, elapsed=%s): %v", info, elapsed, err)
		return "", err
	}

	log.Printf("[handler] agent replied (%s, elapsed=%s): %q", info, elapsed, truncate(reply, 100))
	return reply, nil
}

// setProgressContext sets the current progress context.
func (h *Handler) setProgressContext(ctx *progressContext) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.progressCtx = ctx
}

// handleProgressEvent handles a progress event from an agent.
func (h *Handler) handleProgressEvent(ctx context.Context, pCtx *progressContext, event agent.ProgressEvent) {
	// Check if we should send this notification (rate limit: at most 1 per 3 seconds)
	pCtx.mu.Lock()
	now := time.Now()
	if !pCtx.lastTime.IsZero() && now.Sub(pCtx.lastTime) < 3*time.Second {
		pCtx.mu.Unlock()
		return
	}
	pCtx.lastTime = now
	pCtx.mu.Unlock()

	// Send progress notification to WeChat
	clientID := NewClientID()
	message := fmt.Sprintf("⏳ %s", event.Message)
	if err := SendTextReply(ctx, pCtx.client, pCtx.userID, message, pCtx.token, clientID); err != nil {
		log.Printf("[handler] failed to send progress notification: %v", err)
	} else {
		log.Printf("[handler] sent progress notification: %s", event.Message)
	}
}

// switchDefault switches the default agent. Starts it on demand if needed.
// The change is persisted to config file.
func (h *Handler) switchDefault(ctx context.Context, name string) string {
	ag, err := h.getAgent(ctx, name)
	if err != nil {
		log.Printf("[handler] failed to switch default to %q: %v", name, err)
		return fmt.Sprintf("Failed to switch to %q: %v", name, err)
	}

	h.mu.Lock()
	old := h.defaultName
	h.defaultName = name
	h.agents[name] = ag
	h.mu.Unlock()

	// Persist to config file
	if h.saveDefault != nil {
		if err := h.saveDefault(name); err != nil {
			log.Printf("[handler] failed to save default agent to config: %v", err)
		} else {
			log.Printf("[handler] saved default agent %q to config", name)
		}
	}

	info := ag.Info()
	log.Printf("[handler] switched default agent: %s -> %s (%s)", old, name, info)
	return fmt.Sprintf("switch to %s", name)
}

// resetDefaultSession resets the session for the given userID on the default agent.
func (h *Handler) resetDefaultSession(ctx context.Context, userID string) string {
	ag := h.getDefaultAgent()
	if ag == nil {
		return "No agent running."
	}
	name := ag.Info().Name
	sessionID, err := ag.ResetSession(ctx, userID)
	if err != nil {
		log.Printf("[handler] reset session failed for %s: %v", userID, err)
		return fmt.Sprintf("Failed to reset session: %v", err)
	}
	if sessionID != "" {
		return fmt.Sprintf("已创建新的%s会话\n%s", name, sessionID)
	}
	return fmt.Sprintf("已创建新的%s会话", name)
}

// handleCwd handles the /cwd command. It updates the working directory for all running agents.
func (h *Handler) handleCwd(trimmed string) string {
	arg := strings.TrimSpace(strings.TrimPrefix(trimmed, "/cwd"))
	if arg == "" {
		// No path provided — show current cwd of default agent
		ag := h.getDefaultAgent()
		if ag == nil {
			return "No agent running."
		}
		info := ag.Info()
		return fmt.Sprintf("cwd: (check agent config)\nagent: %s", info.Name)
	}

	// Expand ~ to home directory
	if arg == "~" {
		home, err := os.UserHomeDir()
		if err == nil {
			arg = home
		}
	} else if strings.HasPrefix(arg, "~/") {
		home, err := os.UserHomeDir()
		if err == nil {
			arg = filepath.Join(home, arg[2:])
		}
	}

	// Resolve to absolute path
	absPath, err := filepath.Abs(arg)
	if err != nil {
		return fmt.Sprintf("Invalid path: %v", err)
	}

	// Verify directory exists
	info, err := os.Stat(absPath)
	if err != nil {
		return fmt.Sprintf("Path not found: %s", absPath)
	}
	if !info.IsDir() {
		return fmt.Sprintf("Not a directory: %s", absPath)
	}

	// Update cwd on all running agents
	h.mu.RLock()
	agents := make(map[string]agent.Agent, len(h.agents))
	for name, ag := range h.agents {
		agents[name] = ag
	}
	h.mu.RUnlock()

	for name, ag := range agents {
		ag.SetCwd(absPath)
		log.Printf("[handler] updated cwd for agent %s: %s", name, absPath)
	}

	h.mu.Lock()
	for name := range agents {
		h.agentWorkDirs[name] = absPath
	}
	h.mu.Unlock()

	return fmt.Sprintf("cwd: %s", absPath)
}

// hubReplyHint is prepended to /save messages to instruct the agent to return full content directly.
const hubReplyHint = "[系统指令] 你只需要直接回复文本内容。不要创建、写入或保存任何文件。不要请求授权。直接输出你的完整回复即可。\n\n"

// handleSave processes the /save command: sends message to agent, saves reply to hub.
// Usage: /save {filename} {message} — or just /save {filename} when replying to context
func (h *Handler) handleSave(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage, trimmed, clientID string) string {
	// Parse: /save filename [message]
	// Also handles: /save @agent filename message
	parts := strings.Fields(trimmed)
	if len(parts) < 2 {
		return "用法: /save 文件名 消息内容\n例: /save round1.md 分析AI未来"
	}

	// Check if next token is an agent reference (@name or /name)
	var agentName string
	var filenameIdx int

	if (strings.HasPrefix(parts[1], "@") || strings.HasPrefix(parts[1], "/")) && !strings.Contains(parts[1], ".") {
		// parts[1] looks like an agent reference, not a filename
		resolved := h.resolveAlias(parts[1][1:])
		if h.isKnownAgent(resolved) {
			agentName = resolved
			filenameIdx = 2
		} else {
			filenameIdx = 1
		}
	} else {
		filenameIdx = 1
	}

	if len(parts) < filenameIdx+1 {
		return "用法: /save 文件名 消息内容\n例: /save round1.md 分析AI未来"
	}

	filename := parts[filenameIdx]
	message := strings.Join(parts[filenameIdx+1:], " ")

	// No message content → save last agent reply directly
	if message == "" {
		lastReply, ok := h.lastReplies.Load(msg.FromUserID)
		if !ok {
			return "没有找到上一条回复。请先与 agent 对话，或使用 /save 文件名 消息内容。"
		}
		content := lastReply.(string)
		savePath, err := h.hub.Save(filename, content, "user")
		if err != nil {
			log.Printf("[handler] hub save failed: %v", err)
			return "⚠️ 保存到 Hub 失败: " + err.Error()
		}
		log.Printf("[handler] saved last reply to hub: %s", savePath)
		return fmt.Sprintf("✅ 已保存上一条回复到 Hub: %s", filename)
	}

	// Has message content → send to agent, save agent's reply
	// Determine which agent to use
	var ag agent.Agent
	var useName string
	if agentName != "" {
		var err error
		ag, err = h.getAgent(ctx, agentName)
		if err != nil {
			return fmt.Sprintf("Agent %q 不可用: %v", agentName, err)
		}
		useName = agentName
	} else {
		ag = h.getDefaultAgent()
		if ag == nil {
			return "没有可用的 agent"
		}
		h.mu.RLock()
		useName = h.defaultName
		h.mu.RUnlock()
	}

	// Send typing
	go func() {
		if typingErr := SendTypingState(ctx, client, msg.FromUserID, msg.ContextToken); typingErr != nil {
			log.Printf("[handler] failed to send typing state: %v", typingErr)
		}
	}()

	// Use agent-specific conversationID to avoid polluting default session
	conversationID := msg.FromUserID
	if agentName != "" {
		conversationID = "hub:" + agentName + ":" + msg.FromUserID
	}

	// Prepend hint so agent returns full content instead of writing to local files
	fullMessage := hubReplyHint + message

	// Send to agent
	reply, err := h.chatWithAgent(ctx, ag, conversationID, fullMessage, client, msg.ContextToken)
	if err != nil {
		return fmt.Sprintf("Agent 错误: %v", err)
	}

	// Save reply to hub
	savePath, err := h.hub.Save(filename, reply, useName)
	if err != nil {
		log.Printf("[handler] hub save failed: %v", err)
		return reply + "\n\n⚠️ 保存到 Hub 失败: " + err.Error()
	}

	log.Printf("[handler] saved agent reply to hub: %s (agent=%s)", savePath, useName)
	return reply + fmt.Sprintf("\n\n✅ 已保存到 Hub: %s", filename)
}

// handleHub processes the /hub command: reads shared context and optionally sends to agent.
// Usage:
//
//	/hub {message}              — read all shared files, inject context, send to default agent
//	/hub {filename} {msg}       — read specific file, inject, send to agent
//	/hub {filename} {msg}       — if filename ends with .md, save reply to hub
//	/hub ls                     — list files in hub
//	/hub clear                  — clear all hub files
func (h *Handler) handleHub(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage, trimmed, clientID string) string {
	// Parse: /hub [filename] [message] | /hub ls | /hub clear
	rest := strings.TrimSpace(strings.TrimPrefix(trimmed, "/hub"))

	// No arguments → list files
	if rest == "" {
		files, err := h.hub.List()
		if err != nil {
			return fmt.Sprintf("读取 Hub 失败: %v", err)
		}
		if len(files) == 0 {
			return "Hub 是空的。使用 /save 文件名 消息 来保存内容。"
		}
		var sb strings.Builder
		sb.WriteString("📁 Hub 文件列表:\n")
		for _, f := range files {
			sb.WriteString(fmt.Sprintf("  • %s\n", f))
		}
		return sb.String()
	}

	// Sub-commands
	switch {
	case rest == "ls" || rest == "list":
		files, err := h.hub.ListWithInfo()
		if err != nil {
			return fmt.Sprintf("读取 Hub 失败: %v", err)
		}
		if len(files) == 0 {
			return "Hub 是空的。"
		}
		var sb strings.Builder
		sb.WriteString("📁 Hub 文件列表 (最新优先):\n")
		for i, f := range files {
			// Format: [1] filename (时间)
			timeStr := f.ModTime.Format("01-02 15:04")
			sb.WriteString(fmt.Sprintf("  [%d] %s (%s)\n", i+1, f.Name, timeStr))
		}
		sb.WriteString("\n💡 使用 /hub cat <编号> 读取文件")
		return sb.String()

	case strings.HasPrefix(rest, "cat "):
		// /hub cat <number>
		parts := strings.Fields(rest)
		if len(parts) != 2 {
			return "用法: /hub cat <编号>\n示例: /hub cat 1"
		}
		var num int
		_, err := fmt.Sscanf(parts[1], "%d", &num)
		if err != nil || num < 1 {
			return fmt.Sprintf("无效的编号: %q，请使用数字", parts[1])
		}
		files, err := h.hub.ListWithInfo()
		if err != nil {
			return fmt.Sprintf("读取 Hub 失败: %v", err)
		}
		if num > len(files) {
			return fmt.Sprintf("编号超出范围，Hub 只有 %d 个文件", len(files))
		}
		// num is 1-indexed, array is 0-indexed
		targetFile := files[num-1].Name
		content, err := h.hub.ReadFile(targetFile)
		if err != nil {
			return fmt.Sprintf("读取文件失败: %v", err)
		}
		return fmt.Sprintf("📄 %s\n\n%s", targetFile, content)

	case rest == "clear":
		count, err := h.hub.Clear()
		if err != nil {
			return fmt.Sprintf("清空 Hub 失败: %v", err)
		}
		return fmt.Sprintf("🗑️ 已清空 Hub (%d 个文件)", count)

	case strings.HasPrefix(rest, "pipe "):
		// /hub pipe <target_agent> <message>
		// /hub pipe <target_agent> @<编号> <message>  (使用 Hub 文件编号引用)
		// /hub pipe <target_agent> @-1 <message>    (使用最新文件)
		// /hub pipe <target_agent> @<文件名> <消息>  (直接引用文件名，支持部分匹配)
		parts := strings.Fields(rest)
		if len(parts) < 2 {
			return "用法: /hub pipe <目标agent> <消息>\n       /hub pipe <目标agent> @<编号> <消息>\n       /hub pipe <目标agent> @-1 <消息>\n       /hub pipe <目标agent> @<文件名> <消息>\n\n示例: /hub pipe gemini 分析量子计算\n      /hub pipe claude @1 继续分析\n      /hub pipe claude @-1 补充说明\n      /hub pipe claude @gemini 继续分析 (部分匹配)\n      /hub pipe claude @gem 继续分析 (简写)"
		}
		targetAgent := h.resolveAlias(parts[1])
		var message string
		// 处理引用语法: @<编号>、@-1、@<文件名>
		if len(parts) >= 3 && strings.HasPrefix(parts[2], "@") {
			// 引用模式: /hub pipe <agent> @<ref> <message>
			message = strings.Join(parts[2:], " ") // 包含 @<ref> 和后续消息
		} else {
			// 普通模式: /hub pipe <agent> <message>
			message = strings.Join(parts[2:], " ")
			if message == "" {
				return "用法: /hub pipe <目标agent> <消息>\n       /hub pipe <目标agent> @<编号> <消息>\n       /hub pipe <目标agent> @-1 <消息>\n       /hub pipe <目标agent> @<文件名> <消息>\n\n示例: /hub pipe gemini 分析量子计算\n      /hub pipe claude @1 继续分析\n      /hub pipe claude @-1 补充说明\n      /hub pipe claude @gemini 继续分析 (部分匹配)\n      /hub pipe claude @gem 继续分析 (简写)"
			}
		}
		return h.handlePipe(ctx, client, msg, targetAgent, message, clientID)
	}

	// Parse: could be "/hub filename message" or "/hub message"
	// Check if first word is a known hub file
	words := strings.Fields(rest)
	if len(words) == 0 {
		return "用法: /hub {文件名} {消息} 或 /hub {消息}"
	}

	var hubContext string
	var message string
	var targetAgentName string
	var saveFilename string // if set, auto-save reply to hub

	// Check if first word is an agent reference
	wordIdx := 0
	if (strings.HasPrefix(words[0], "@") || strings.HasPrefix(words[0], "/")) && !strings.Contains(words[0], ".") {
		resolved := h.resolveAlias(words[0][1:])
		if h.isKnownAgent(resolved) {
			targetAgentName = resolved
			wordIdx = 1
		}
	}

	if wordIdx >= len(words) {
		return "用法: /hub {文件名} {消息} 或 /hub {消息}"
	}

	// Check if current first word is a known hub file
	if h.hub.Exists(words[wordIdx]) {
		// Read specific file
		ctx2, err := h.hub.ReadSpecific([]string{words[wordIdx]})
		if err != nil {
			return fmt.Sprintf("读取文件失败: %v", err)
		}
		hubContext = ctx2
		// If message follows and the hub file name looks like a save target (.md),
		// use it as save filename for the reply
		if len(words) > wordIdx+1 {
			message = strings.Join(words[wordIdx+1:], " ")
		} else {
			message = ""
		}
	} else {
		// Read all shared files
		ctx2, err := h.hub.ReadAll()
		if err != nil {
			return fmt.Sprintf("读取 Hub 失败: %v", err)
		}
		hubContext = ctx2
		message = strings.Join(words[wordIdx:], " ")
	}

	if message == "" {
		if hubContext == "" {
			return "Hub 是空的，没有可注入的上下文。"
		}
		// Just show the hub content
		return hubContext
	}

	// Determine target agent
	var ag agent.Agent
	var resolvedAgentName string
	if targetAgentName != "" {
		var err error
		ag, err = h.getAgent(ctx, targetAgentName)
		if err != nil {
			return fmt.Sprintf("Agent %q 不可用: %v", targetAgentName, err)
		}
		resolvedAgentName = targetAgentName
	} else {
		ag = h.getDefaultAgent()
		if ag == nil {
			return "没有可用的 agent"
		}
		h.mu.RLock()
		resolvedAgentName = h.defaultName
		h.mu.RUnlock()
	}

	// Send typing
	go func() {
		if typingErr := SendTypingState(ctx, client, msg.FromUserID, msg.ContextToken); typingErr != nil {
			log.Printf("[handler] failed to send typing state: %v", typingErr)
		}
	}()

	// Always use agent-specific conversationID to avoid polluting default session
	conversationID := "hub:" + resolvedAgentName + ":" + msg.FromUserID

	// Build prompt: put hub context as user message (not system) to reduce tool-use tendency.
	// Explicitly forbid file/search tools so agents use the injected context directly.
	wrappedMessage := fmt.Sprintf(
		"【重要】请直接基于下方提供的材料回答问题。禁止使用任何工具（搜索、读文件、写文件等），不要访问文件系统，不要搜索网络。材料已完整提供给你，直接分析即可。\n\n---\n共享材料：\n%s\n---\n\n问题：%s",
		hubContext, message,
	)

	reply, err := h.chatWithAgent(ctx, ag, conversationID, wrappedMessage, client, msg.ContextToken)
	if err != nil {
		return fmt.Sprintf("Agent 错误: %v", err)
	}

	// Auto-detect save filename from conversation flow:
	// If the injected file was round1.md and this is round2, suggest saving as round2
	// But only save if user explicitly used a .md filename as the hub file reference
	if saveFilename != "" {
		savePath, err := h.hub.Save(saveFilename, reply, resolvedAgentName)
		if err != nil {
			log.Printf("[handler] hub auto-save failed: %v", err)
		} else {
			log.Printf("[handler] auto-saved hub reply to: %s (agent=%s)", savePath, resolvedAgentName)
			return reply + fmt.Sprintf("\n\n✅ 已保存到 Hub: %s", saveFilename)
		}
	}

	return reply
}

// handlePipe 实现自动链式调用: 先将消息发送给默认 agent，然后将回复保存并发送给目标 agent
// 支持引用语法：
//
//	/hub pipe <agent> @<编号> <消息> - 直接使用 Hub 中编号对应的文件作为源内容
//	/hub pipe <agent> @-1 <消息> - 使用最新文件（-1=最新，-2=第二新）
//	/hub pipe <agent> @<文件名> <消息> - 直接使用文件名引用
func (h *Handler) handlePipe(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage, targetAgent, message, clientID string) string {
	log.Printf("[hub/pipe] starting pipe: target=%s, message=%q", targetAgent, truncate(message, 50))

	timestamp := time.Now().Format("20060102-150405")

	var reply1 string
	var filename string
	var sourceAgentName string

	// 检测是否使用 @ 引用语法
	trimmedMsg := strings.TrimSpace(message)
	if strings.HasPrefix(trimmedMsg, "@") {
		// 解析引用语法
		refStr := trimmedMsg[1:] // 去掉 @

		// 尝试解析为相对编号 (@-1, @-2) 或绝对编号 (@1, @2)
		var refNum int
		n, err := fmt.Sscanf(refStr, "%d", &refNum)

		if n == 1 && err == nil {
			// 数字引用模式
			files, ferr := h.hub.ListWithInfo()
			if ferr != nil {
				return fmt.Sprintf("❌ 读取 Hub 失败: %v", ferr)
			}
			if len(files) == 0 {
				return "❌ Hub 是空的，没有可引用的文件"
			}

			var targetFile string
			if refNum < 0 {
				// 相对编号: @-1=最新, @-2=第二新
				idx := -refNum - 1
				if idx >= len(files) {
					return fmt.Sprintf("❌ 相对编号超出范围，Hub 只有 %d 个文件", len(files))
				}
				targetFile = files[idx].Name
				sourceAgentName = fmt.Sprintf("Hub[@%d=最新]", refNum)
			} else {
				// 绝对编号: @1=最新, @2=第二新
				if refNum > len(files) {
					return fmt.Sprintf("❌ 编号超出范围，Hub 只有 %d 个文件", len(files))
				}
				targetFile = files[refNum-1].Name
				sourceAgentName = fmt.Sprintf("Hub[@%d]", refNum)
			}

			content, cerr := h.hub.ReadFile(targetFile)
			if cerr != nil {
				return fmt.Sprintf("❌ 读取文件 %s 失败: %v", targetFile, cerr)
			}
			reply1 = content
			filename = targetFile
			log.Printf("[hub/pipe] using hub reference @%s -> file %s", refStr, targetFile)
		} else {
			// 尝试作为文件名引用 @filename.md
			refFilename := refStr
			// 如果引用后没有空格或消息，整个 trimmedMsg 就是 @filename
			// 否则需要提取文件名部分（遇到空格为止）
			if spaceIdx := strings.Index(refStr, " "); spaceIdx > 0 {
				refFilename = refStr[:spaceIdx]
			} else {
				refFilename = refStr
			}

			// 先尝试完全匹配
			if h.hub.Exists(refFilename) {
				content, cerr := h.hub.ReadFile(refFilename)
				if cerr != nil {
					return fmt.Sprintf("❌ 读取文件 %s 失败: %v", refFilename, cerr)
				}
				reply1 = content
				filename = refFilename
				sourceAgentName = fmt.Sprintf("Hub[%s]", refFilename)
				log.Printf("[hub/pipe] using hub file reference @%s", refFilename)
			} else {
				// 尝试部分匹配
				matchedFile, err := h.hub.FindByPartialName(refFilename)
				if err == nil {
					content, cerr := h.hub.ReadFile(matchedFile)
					if cerr != nil {
						return fmt.Sprintf("❌ 读取文件 %s 失败: %v", matchedFile, cerr)
					}
					reply1 = content
					filename = matchedFile
					sourceAgentName = fmt.Sprintf("Hub[%s]", refFilename)
					log.Printf("[hub/pipe] using hub partial match @%s -> file %s", refFilename, matchedFile)
				} else {
					return fmt.Sprintf("❌ 找不到匹配 %q 的文件\n\n💡 提示:\n- 使用 @<编号> 引用: @1、@-1\n- 使用 @<部分文件名>: @gemini、@gem\n- 查看文件: /hub list\n\n示例: /hub pipe claude @1 继续分析", refFilename)
				}
			}
		}
	}

	// 如果没有使用引用语法，则走正常的 pipe 流程
	if reply1 == "" {
		// 1. 获取默认 agent（作为 source）
		sourceAgent := h.getDefaultAgent()
		if sourceAgent == nil {
			return "❌ 没有可用的默认 agent，请先设置默认 agent（如 /claude）"
		}

		// 使用配置名称而不是 Info().Name（后者可能返回进程路径）
		h.mu.RLock()
		sourceAgentName = h.defaultName
		h.mu.RUnlock()

		// 2. 发送消息给 source agent，得到第一轮回复
		log.Printf("[hub/pipe] step1: sending to default agent (%s)", sourceAgentName)
		var err error
		reply1, err = h.chatWithAgent(ctx, sourceAgent, msg.FromUserID, message, client, msg.ContextToken)
		if err != nil {
			return fmt.Sprintf("❌ 第一步（默认 agent %s）失败: %v", sourceAgentName, err)
		}

		// 3. 自动保存第一轮回复到 Hub
		// 使用简洁的文件名：pipe_<timestamp>_<agent>.md
		shortAgentName := sourceAgentName
		if idx := strings.LastIndex(sourceAgentName, "/"); idx >= 0 {
			shortAgentName = sourceAgentName[idx+1:]
		}
		filename = fmt.Sprintf("pipe_%s_%s.md", timestamp, shortAgentName)
		savePath, err := h.hub.Save(filename, reply1, sourceAgentName)
		if err != nil {
			log.Printf("[hub/pipe] save failed: %v", err)
			// 即使保存失败，仍继续执行第二步（降级）
			filename = ""
		} else {
			log.Printf("[hub/pipe] saved step1 reply to %s", savePath)
		}
	}

	// 4. 获取目标 agent
	targetAg, err := h.getAgent(ctx, targetAgent)
	if err != nil {
		return fmt.Sprintf("❌ 目标 agent %q 不可用: %v", targetAgent, err)
	}

	// 5. 构造第二步的 prompt：让目标 agent 基于刚保存的文件进行分析
	var hubContext string
	if filename != "" {
		hubContext, err = h.hub.ReadSpecific([]string{filename})
		if err != nil {
			log.Printf("[hub/pipe] read saved file failed: %v", err)
			hubContext = ""
		}
	}

	if hubContext == "" {
		// 若读取失败，降级为直接传递 reply1
		hubContext = fmt.Sprintf("上一步的回复：\n%s", reply1)
	}

	secondPrompt := fmt.Sprintf(
		"请基于以下内容，继续进行分析或给出你的观点：\n\n---\n%s\n---\n\n要求：直接输出分析结果，不要重复原文。",
		hubContext,
	)

	// 6. 发送给目标 agent（使用独立 conversationID 避免污染）
	convID := "hub:" + targetAgent + ":" + msg.FromUserID
	log.Printf("[hub/pipe] step2: sending to target agent (%s)", targetAgent)
	reply2, err := h.chatWithAgent(ctx, targetAg, convID, secondPrompt, client, msg.ContextToken)
	if err != nil {
		return fmt.Sprintf("❌ 第二步（目标 agent %s）失败: %v", targetAgent, err)
	}

	// 7. 自动保存最终结果
	finalFilename := fmt.Sprintf("pipe_%s_%s_final.md", timestamp, targetAgent)
	finalSaved := false
	if _, err := h.hub.Save(finalFilename, reply2, targetAgent); err != nil {
		log.Printf("[hub/pipe] failed to save final reply: %v", err)
	} else {
		finalSaved = true
	}

	// 8. 返回最终回复（附加保存路径信息和文件编号）
	result := reply2
	if filename != "" || finalSaved {
		// 获取当前 Hub 文件列表以显示编号
		files, _ := h.hub.ListWithInfo()

		// 查找源文件和目标文件的编号
		var sourceNum, finalNum int
		for i, f := range files {
			if f.Name == filename {
				sourceNum = i + 1
			}
			if f.Name == finalFilename {
				finalNum = i + 1
			}
		}

		var fileInfo strings.Builder
		fileInfo.WriteString(fmt.Sprintf("\n\n📁 Pipe 流程: %s → %s", sourceAgentName, targetAgent))

		if filename != "" && sourceNum > 0 {
			fileInfo.WriteString(fmt.Sprintf("\n💾 源文件: [@%d] %s", sourceNum, filename))
		}
		if finalSaved && finalNum > 0 {
			fileInfo.WriteString(fmt.Sprintf("\n💾 结果: [@%d] %s", finalNum, finalFilename))
		}

		// 提示用户如何继续
		if finalNum > 0 {
			fileInfo.WriteString(fmt.Sprintf("\n\n💡 继续分析: /hub pipe <agent> @%d <消息>", finalNum))
		}

		result += fileInfo.String()
	}
	return result
}

// buildStatus returns a short status string showing the current default agent.
func (h *Handler) buildStatus() string {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if h.defaultName == "" {
		return "agent: none (echo mode)"
	}

	ag, ok := h.agents[h.defaultName]
	if !ok {
		return fmt.Sprintf("agent: %s (not started)", h.defaultName)
	}

	info := ag.Info()
	return fmt.Sprintf("agent: %s\ntype: %s\nmodel: %s", h.defaultName, info.Type, info.Model)
}

// analyzeWithNanobot sends a WeChat article to nanobot for analysis.
func (h *Handler) analyzeWithNanobot(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage, meta *LinkMetadata) {
	// Get nanobot agent
	ag, err := h.getAgent(ctx, "nanobot")
	if err != nil {
		log.Printf("[handler] failed to get nanobot for analysis: %v", err)
		return
	}

	// Build analysis prompt
	prompt := fmt.Sprintf("请分析这篇微信文章，给出摘要和关键观点：\n\n标题：%s\n\n文章内容：\n%s",
		meta.Title, meta.Body)

	// Send typing indicator
	go func() {
		if typingErr := SendTypingState(ctx, client, msg.FromUserID, msg.ContextToken); typingErr != nil {
			log.Printf("[handler] failed to send typing state: %v", typingErr)
		}
	}()

	// Get analysis from nanobot
	reply, err := h.chatWithAgent(ctx, ag, msg.FromUserID, prompt, client, msg.ContextToken)
	if err != nil {
		log.Printf("[handler] nanobot analysis failed: %v", err)
		reply = fmt.Sprintf("分析失败: %v", err)
	}

	// Send analysis result
	clientID := NewClientID()
	if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
		log.Printf("[handler] failed to send analysis reply to %s: %v", msg.FromUserID, err)
	}

	// Optionally send to remote clipboard endpoint (include original article + AI analysis)
	if h.remoteClipboardURL != "" {
		go func() {
			clipboardCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			var clipboardContent string
			if meta.Body != "" {
				clipboardContent = fmt.Sprintf("【原文】\n标题：%s\n\n%s\n\n━━━━━━━━━━━━━━━━\n\n【AI 解读】\n%s",
					meta.Title, meta.Body, reply)
			} else {
				clipboardContent = reply
			}
			if err := h.sendToClipboard(clipboardCtx, clipboardContent); err != nil {
				log.Printf("[handler] failed to send to clipboard: %v", err)
			}
		}()
	}
}

func buildHelpText() string {
	return `🤖 WeClaw Agent Hub

📌 基本指令
  @agent msg       发给指定 agent
  @a @b msg        广播给多个 agent
  @agent           切换默认 agent
  /new /clear      新会话
  /cwd /path       切换工作目录
  /info /help      信息 / 帮助

🎙️ 播客生成
  /podcast         使用上一条回复生成播客
  /podcast <内容>   指定内容生成播客
  （无论当前处于哪个 agent，均会自动拦截并发送）

🎭 多 Agent 辩论
  /debate <话题>              默认两个 agent 辩论
  /debate @a @b <话题>        指定 agent 辩论
  示例: /debate AI 会取代人类决策吗

🖥️ 终端模拟
  /sh              进入命令行模式（支持持久化目录、免前缀）
  /sh <命令>       执行单次命令（不进入模式）
  命令行模式下: cd /q 退出/切换目录，ls cat pwd 等

 📋 待办事项
   /todo <事项>        添加待办（支持自然语言时间）
   /todo list          查看待办列表
   /todo done <编号>   完成待办
   /todo del <编号>    删除待办
   /todo clear         清空所有待办

 ⏱ 计时器
   /timer 25           25分钟倒计时
   /timer 2h 写报告    设定时间+标签（支持 AI 解析）
   /timer list         查看进行中的计时器
   /timer cancel <编号> 取消计时器
   /timer clear        取消所有计时器

📂 Agent（默认: nanobot）
  nanobot(nb,n,bot)  claude(c)  gemini(g)  deepseek(ds)
  pa(p)  ps  po  pg  zhipu(glm,z)

🔗 Hub · 跨 Agent 上下文共享
  /hub              列出共享文件（显示编号）
  /hub {msg}        注入所有共享文件后发给 agent
  /hub {file} {msg} 注入指定文件后发给 agent
  /hub ls /clear    列出 / 清空
  /hub cat {编号}   查看指定编号的文件内容

🔄 Pipe · Agent 链式协作
  /hub pipe <agent> <消息>           默认 agent → 目标 agent
  /hub pipe <agent> @1 <消息>        引用 Hub 编号 1 的文件
  /hub pipe <agent> @-1 <消息>       引用最新文件
  /hub pipe <agent> @file.md <消息>  引用指定文件名

  示例:
  /hub pipe gemini 量子计算原理          # nanobot → gemini
  /hub pipe claude @2 商业应用前景        # 继续分析结果 2
  /hub pipe deepseek @-1 投资建议         # 引用最新结果

💾 /save {file} {msg}          发给 agent 并保存回复
     /save {file} @agent {msg}  指定 agent 并保存回复

💡 多 Agent 辩论示例
  /hub pipe gemini AI应该替代人类决策
  /hub pipe claude @1 反驳以上观点
  /hub pipe deepseek @2 总结双方观点

🔄 工作流
  /workflow           查看工作流语法帮助
  /workflow DSL...    执行多步骤 Agent 编排
  支持: 顺序链式 + 并行分支 + 自动保存 + @N 引用

📅 定时任务
  /cron add 每天早上9点提醒我喝水            添加定时任务（支持自然语言）
  /cron add 每周一早上8点生成周报            添加定时工作流
  /cron add "0 9 * * *" 提醒喝水              传统 cron 表达式格式
  /cron list                                  查看定时任务
  /cron remove <id>                           删除定时任务
  /cron enable/disable <id>                  启用/禁用任务`
}

func extractText(msg ilink.WeixinMessage) string {
	for _, item := range msg.ItemList {
		if item.Type == ilink.ItemTypeText && item.TextItem != nil {
			return item.TextItem.Text
		}
	}
	return ""
}

func extractImage(msg ilink.WeixinMessage) *ilink.ImageItem {
	for _, item := range msg.ItemList {
		if item.Type == ilink.ItemTypeImage && item.ImageItem != nil {
			return item.ImageItem
		}
	}
	return nil
}

func extractVoiceText(msg ilink.WeixinMessage) string {
	for _, item := range msg.ItemList {
		if item.Type == ilink.ItemTypeVoice && item.VoiceItem != nil && item.VoiceItem.Text != "" {
			return item.VoiceItem.Text
		}
	}
	return ""
}

func (h *Handler) handleImageSave(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage, img *ilink.ImageItem) {
	clientID := NewClientID()
	log.Printf("[handler] received image from %s, saving to %s", msg.FromUserID, h.saveDir)

	// Download image data
	var data []byte
	var err error

	if img.URL != "" {
		// Direct URL download
		data, _, err = downloadFile(ctx, img.URL)
	} else if img.Media != nil && img.Media.EncryptQueryParam != "" {
		// CDN encrypted download
		data, err = DownloadFileFromCDN(ctx, img.Media.EncryptQueryParam, img.Media.AESKey)
	} else {
		log.Printf("[handler] image has no URL or media info from %s", msg.FromUserID)
		return
	}

	if err != nil {
		log.Printf("[handler] failed to download image from %s: %v", msg.FromUserID, err)
		reply := fmt.Sprintf("Failed to save image: %v", err)
		_ = SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID)
		return
	}

	// Detect extension from content
	ext := detectImageExt(data)

	// Generate filename with timestamp
	ts := time.Now().Format("20060102-150405")
	fileName := fmt.Sprintf("%s%s", ts, ext)
	filePath := filepath.Join(h.saveDir, fileName)

	// Ensure save directory exists
	if err := os.MkdirAll(h.saveDir, 0o755); err != nil {
		log.Printf("[handler] failed to create save dir: %v", err)
		return
	}

	// Write image file
	if err := os.WriteFile(filePath, data, 0o644); err != nil {
		log.Printf("[handler] failed to write image: %v", err)
		reply := fmt.Sprintf("Failed to save image: %v", err)
		_ = SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID)
		return
	}

	// Write sidecar file
	sidecarPath := filePath + ".sidecar.md"
	sidecarContent := fmt.Sprintf("---\nid: %s\n---\n", uuid.New().String())
	if err := os.WriteFile(sidecarPath, []byte(sidecarContent), 0o644); err != nil {
		log.Printf("[handler] failed to write sidecar: %v", err)
	}

	log.Printf("[handler] saved image to %s (%d bytes)", filePath, len(data))
	reply := fmt.Sprintf("Saved: %s", fileName)
	if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
		log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
	}
}

func detectImageExt(data []byte) string {
	if len(data) < 4 {
		return ".bin"
	}
	// PNG: 89 50 4E 47
	if data[0] == 0x89 && data[1] == 0x50 && data[2] == 0x4E && data[3] == 0x47 {
		return ".png"
	}
	// JPEG: FF D8 FF
	if data[0] == 0xFF && data[1] == 0xD8 && data[2] == 0xFF {
		return ".jpg"
	}
	// GIF: 47 49 46
	if data[0] == 0x47 && data[1] == 0x49 && data[2] == 0x46 {
		return ".gif"
	}
	// WebP: 52 49 46 46 ... 57 45 42 50
	if len(data) >= 12 && data[0] == 0x52 && data[1] == 0x49 && data[8] == 0x57 && data[9] == 0x45 {
		return ".webp"
	}
	// BMP: 42 4D
	if data[0] == 0x42 && data[1] == 0x4D {
		return ".bmp"
	}
	return ".jpg" // default to jpg for WeChat images
}

// extractAllMedia extracts all media items (image, file, video) from a message.
// Downloads CDN media to local files if necessary.
func (h *Handler) extractAllMedia(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage) []agent.MediaEntry {
	var media []agent.MediaEntry

	for _, item := range msg.ItemList {
		switch item.Type {
		case ilink.ItemTypeImage:
			if item.ImageItem != nil {
				entry := agent.MediaEntry{Type: "image"}
				log.Printf("[handler] image item: URL=%q, Media=%v, MidSize=%d", item.ImageItem.URL, item.ImageItem.Media != nil, item.ImageItem.MidSize)
				// Check if URL is a valid HTTP URL
				if item.ImageItem.URL != "" && strings.HasPrefix(item.ImageItem.URL, "http") {
					entry.URL = item.ImageItem.URL
					log.Printf("[handler] image HTTP URL: %s", entry.URL)
				} else if item.ImageItem.Media != nil && h.saveDir != "" {
					// CDN media - download and decrypt
					log.Printf("[handler] image has CDN media: encrypt_param=%s", item.ImageItem.Media.EncryptQueryParam)
					localPath, err := downloadCDNMedia(ctx, client, item.ImageItem.Media, h.saveDir, ".jpg")
					if err != nil {
						log.Printf("[handler] failed to download CDN image: %v", err)
					} else {
						entry.Path = localPath
						log.Printf("[handler] downloaded CDN image to: %s", localPath)
					}
				} else if item.ImageItem.URL != "" && h.saveDir != "" {
					// URL is actually encrypt_query_param, download from CDN
					log.Printf("[handler] image URL is encrypt_param: %s (MidSize=%d)", item.ImageItem.URL, item.ImageItem.MidSize)
					mediaInfo := &ilink.MediaInfo{
						EncryptQueryParam: item.ImageItem.URL,
						AESKey:            "",
						EncryptType:       0,
					}
					localPath, err := downloadCDNMedia(ctx, client, mediaInfo, h.saveDir, ".jpg")
					if err != nil {
						log.Printf("[handler] failed to download CDN image from encrypt_param: %v", err)
					} else {
						entry.Path = localPath
						log.Printf("[handler] downloaded CDN image to: %s", localPath)
					}
				} else {
					log.Printf("[handler] image has no valid URL or CDN media, skipping")
				}
				media = append(media, entry)
			}
		case ilink.ItemTypeFile:
			if item.FileItem != nil {
				entry := agent.MediaEntry{
					Type:     "file",
					FileName: item.FileItem.FileName,
				}
				// Debug: print full FileItem JSON
				if fileJSON, err := json.Marshal(item.FileItem); err == nil {
					log.Printf("[handler] FileItem JSON: %s", string(fileJSON))
				}
				if item.FileItem.Media != nil {
					log.Printf("[handler] file MediaInfo: EncryptQueryParam=%q AESKey=%q EncryptType=%d Len=%s",
						item.FileItem.Media.EncryptQueryParam[:min(40, len(item.FileItem.Media.EncryptQueryParam))]+"...",
						item.FileItem.Media.AESKey,
						item.FileItem.Media.EncryptType,
						item.FileItem.Len)
				}
				if item.FileItem.Media != nil && h.saveDir != "" {
					// CDN file - download and decrypt
					ext := filepath.Ext(item.FileItem.FileName)
					if ext == "" {
						ext = ".bin"
					}
					localPath, err := downloadCDNMedia(ctx, client, item.FileItem.Media, h.saveDir, ext)
					if err != nil {
						log.Printf("[handler] failed to download CDN file: %v", err)
					} else {
						entry.Path = localPath
						log.Printf("[handler] downloaded CDN file to: %s", localPath)
					}
				}
				log.Printf("[handler] file: name=%s, path=%s", entry.FileName, entry.Path)
				media = append(media, entry)
			}
		case ilink.ItemTypeVideo:
			if item.VideoItem != nil {
				entry := agent.MediaEntry{Type: "video"}
				if item.VideoItem.Media != nil && h.saveDir != "" {
					// CDN video - download and decrypt
					localPath, err := downloadCDNMedia(ctx, client, item.VideoItem.Media, h.saveDir, ".mp4")
					if err != nil {
						log.Printf("[handler] failed to download CDN video: %v", err)
					} else {
						entry.Path = localPath
						log.Printf("[handler] downloaded CDN video to: %s", localPath)
					}
				}
				log.Printf("[handler] video item found, path=%s", entry.Path)
				media = append(media, entry)
			}
		}
	}

	return media
}

// sendMediaToAgent sends a message with media attachments to the default agent.
func (h *Handler) sendMediaToAgent(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage, text string, media []agent.MediaEntry) {
	// Store context token for this user
	h.contextTokens.Store(msg.FromUserID, msg.ContextToken)

	clientID := NewClientID()

	// Send typing indicator
	go func() {
		if typingErr := SendTypingState(ctx, client, msg.FromUserID, msg.ContextToken); typingErr != nil {
			log.Printf("[handler] failed to send typing state: %v", typingErr)
		}
	}()

	h.mu.RLock()
	defaultName := h.defaultName
	h.mu.RUnlock()

	ag := h.getDefaultAgent()
	var reply string
	if ag != nil {
		var err error
		log.Printf("[handler] sending %d media items to agent for %s", len(media), msg.FromUserID)
		reply, err = h.chatWithAgentAndMedia(ctx, ag, msg.FromUserID, text, media)
		if err != nil {
			reply = fmt.Sprintf("Error: %v", err)
		}
	} else {
		log.Printf("[handler] agent not ready, using echo mode for %s", msg.FromUserID)
		reply = "[echo] received media"
	}

	h.sendReplyWithMedia(ctx, client, msg, defaultName, reply, clientID)
}

// chatWithAgentAndMedia sends a message with media attachments to an agent and returns the reply.
func (h *Handler) chatWithAgentAndMedia(ctx context.Context, ag agent.Agent, userID, message string, media []agent.MediaEntry) (string, error) {
	info := ag.Info()
	log.Printf("[handler] dispatching to agent (%s) for %s with %d media items", info, userID, len(media))

	start := time.Now()
	reply, err := ag.ChatWithMedia(ctx, userID, message, media)
	elapsed := time.Since(start)

	if err != nil {
		log.Printf("[handler] agent error (%s, elapsed=%s): %v", info, elapsed, err)
		return "", err
	}

	log.Printf("[handler] agent replied (%s, elapsed=%s): %q", info, elapsed, truncate(reply, 100))
	return reply, nil
}

// downloadCDNMedia downloads and decrypts media from WeChat CDN.
// Returns the local file path where the decrypted media is saved.
func downloadCDNMedia(ctx context.Context, client *ilink.Client, media *ilink.MediaInfo, saveDir string, ext string) (string, error) {
	if media == nil || media.EncryptQueryParam == "" {
		return "", fmt.Errorf("invalid media info")
	}

	// Build CDN download URL using the correct CDN endpoint
	cdnURL := fmt.Sprintf("https://novac2c.cdn.weixin.qq.com/c2c/download?encrypted_query_param=%s",
		url.QueryEscape(media.EncryptQueryParam))
	log.Printf("[handler] downloading CDN media from: %s", cdnURL)

	// Download encrypted data
	req, err := http.NewRequestWithContext(ctx, "GET", cdnURL, nil)
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("download failed: status %d", resp.StatusCode)
	}

	encryptedData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response: %w", err)
	}

	log.Printf("[handler] downloaded %d bytes of data", len(encryptedData))

	var fileData []byte
	log.Printf("[handler] file MediaInfo: EncryptQueryParam=%q AESKey=%q EncryptType=%d Len=%d",
		media.EncryptQueryParam[:20], media.AESKey, media.EncryptType, len(encryptedData))
	if media.AESKey != "" {
		// Decrypt using AES-128-ECB
		// AES key format: base64 -> hex string -> raw bytes
		aesKeyHexBytes, err := base64.StdEncoding.DecodeString(media.AESKey)
		if err != nil {
			return "", fmt.Errorf("decode aes key base64: %w", err)
		}
		aesKey, err := hex.DecodeString(string(aesKeyHexBytes))
		if err != nil {
			return "", fmt.Errorf("decode aes key hex: %w", err)
		}

		fileData, err = decryptAES128ECB(encryptedData, aesKey)
		if err != nil {
			return "", fmt.Errorf("decrypt: %w", err)
		}
		log.Printf("[handler] decrypted %d bytes", len(fileData))
	} else {
		// No encryption key or EncryptType != 1 — data is plaintext
		fileData = encryptedData
		if media.AESKey == "" {
			log.Printf("[handler] no AES key, using raw data (no decryption)")
		} else {
			log.Printf("[handler] EncryptType=%d (not AES-128-ECB), using raw data", media.EncryptType)
		}
	}

	// Save to local file
	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	filePath := filepath.Join(saveDir, filename)

	if err := os.WriteFile(filePath, fileData, 0644); err != nil {
		return "", fmt.Errorf("save file: %w", err)
	}

	log.Printf("[handler] saved decrypted media to: %s", filePath)
	return filePath, nil
}

// decryptAES128ECB decrypts data using AES-128-ECB mode.
func decryptAES128ECB(encrypted, key []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, fmt.Errorf("invalid key length: %d (expected 16)", len(key))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("create cipher: %w", err)
	}

	if len(encrypted)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("encrypted data length %d is not a multiple of block size", len(encrypted))
	}

	decrypted := make([]byte, len(encrypted))
	for i := 0; i < len(encrypted); i += aes.BlockSize {
		block.Decrypt(decrypted[i:i+aes.BlockSize], encrypted[i:i+aes.BlockSize])
	}

	// Remove PKCS7 padding
	padding := int(decrypted[len(decrypted)-1])
	if padding > 0 && padding <= aes.BlockSize {
		decrypted = decrypted[:len(decrypted)-padding]
	}

	return decrypted, nil
}

const podcastAPIURL = "https://api.yuangs.cc/api/publish"

// generatePodcastTitle extracts the first line, removes markdown markers, and truncates.
func generatePodcastTitle(text string) string {
	// Take first line
	lines := strings.Split(text, "\n")
	if len(lines) == 0 {
		return "[Read] 无标题"
	}
	firstLine := lines[0]

	// Remove common markdown markers: #, *, >, -, `, [, ], etc.
	re := regexp.MustCompile(`[#*>\-\[\]` + "`" + `]`)
	cleaned := re.ReplaceAllString(firstLine, "")
	cleaned = strings.TrimSpace(cleaned)
	if cleaned == "" {
		cleaned = "无标题"
	}

	// Add prefix and truncate to 50 chars (using rune to safely handle Chinese)
	title := "[Read] " + cleaned
	runes := []rune(title)
	if len(runes) > 50 {
		title = string(runes[:50])
	}
	return title
}

// sendToPodcast sends text to the remote podcast API.
func (h *Handler) sendToPodcast(ctx context.Context, text string) error {
	title := generatePodcastTitle(text)

	payload := map[string]interface{}{
		"title":      title,
		"content":    text,
		"content_md": text,
		"targets":    []string{"nas"},
		"transform":  "read",
	}
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, podcastAPIURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-App-ID", "taio-quick-read")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	// Read and log response body for debugging
	respBody, _ := io.ReadAll(resp.Body)
	log.Printf("[podcast] API response status=%d, body=%s", resp.StatusCode, string(respBody))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	return nil
}

// sendToClipboard sends content to the remote clipboard endpoint.
func (h *Handler) sendToClipboard(ctx context.Context, content string) error {
	if h.remoteClipboardURL == "" {
		return fmt.Errorf("remote clipboard URL not configured")
	}

	// Build payload - wrap content in JSON as per Cloudflare Worker spec
	payload := map[string]string{
		"content": content,
	}
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, h.remoteClipboardURL+"/push", bytes.NewReader(bodyBytes))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if h.remoteClipboardKey != "" {
		req.Header.Set("X-Auth-Key", h.remoteClipboardKey)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(respBody))
	}

	log.Printf("[clipboard] sent content to remote endpoint, status=%d", resp.StatusCode)
	return nil
}

// sendToRelay sends a Q&A pair or notification to the relay endpoint.
func (h *Handler) sendToRelay(ctx context.Context, agentName, question, reply string) {
	if h.relayURL == "" {
		return
	}

	ts := time.Now().Format("2006-01-02 15:04:05")
	content := fmt.Sprintf("## %s (%s)\n\n**Q:** %s\n\n**A:** %s", agentName, ts, question, reply)

	payload := map[string]string{"content": content}
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[relay] marshal payload: %v", err)
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, h.relayURL, bytes.NewReader(bodyBytes))
	if err != nil {
		log.Printf("[relay] create request: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	if h.relayAuthKey != "" {
		req.Header.Set("X-Auth-Key", h.relayAuthKey)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[relay] push failed: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Printf("[relay] pushed to %s (status=%d)", h.relayURL, resp.StatusCode)
	} else {
		log.Printf("[relay] push returned non-2xx: status=%d", resp.StatusCode)
	}
}

// SendDisconnectNotice sends a disconnection notification to the relay endpoint.
func SendDisconnectNotice(ctx context.Context, relayURL, relayAuthKey, botID, reason string) {
	if relayURL == "" {
		return
	}

	ts := time.Now().Format("2006-01-02 15:04:05")
	content := fmt.Sprintf("## 断线通知 (%s)\n\n**Bot:** %s\n\n**原因:** %s", ts, botID, reason)

	payload := map[string]string{"content": content}
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[relay] marshal payload: %v", err)
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, relayURL, bytes.NewReader(bodyBytes))
	if err != nil {
		log.Printf("[relay] create request: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	if relayAuthKey != "" {
		req.Header.Set("X-Auth-Key", relayAuthKey)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[relay] disconnect push failed: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Printf("[relay] disconnect push to %s (status=%d)", relayURL, resp.StatusCode)
	} else {
		log.Printf("[relay] disconnect push returned non-2xx: status=%d", resp.StatusCode)
	}
}

// handlePodcast processes /podcast command.
func (h *Handler) handlePodcast(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage, trimmed, clientID string) string {
	parts := strings.Fields(trimmed)
	var text string

	if len(parts) == 1 {
		// No argument: use last agent reply
		lastReply, ok := h.lastReplies.Load(msg.FromUserID)
		if !ok {
			return "没有找到上一条回复。请先与 agent 对话，或使用 /podcast <消息> 指定内容。"
		}
		text = lastReply.(string)
	} else {
		// Has argument: join remaining parts
		text = strings.Join(parts[1:], " ")
	}

	if strings.TrimSpace(text) == "" {
		return "消息内容为空，无法生成播客。"
	}

	// Send to podcast API
	if err := h.sendToPodcast(ctx, text); err != nil {
		log.Printf("[handler] podcast error: %v", err)
		return "❌ 播客生成失败，请稍后重试。"
	}

	return "✅ 已加入 NAS 直读队列，请稍后查看播客。"
}

// handleDebate orchestrates a multi-round debate between two agents on a topic.
// Usage:
//
//	/debate <话题>                  — 使用默认两个 agent 辩论
//	/debate @agent1 @agent2 <话题>  — 指定 agent 辩论
func (h *Handler) handleDebate(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage, trimmed, clientID string) string {
	rest := strings.TrimSpace(strings.TrimPrefix(trimmed, "/debate"))
	if rest == "" {
		return "用法:\n/debate <话题> — 使用默认 agent 辩论\n/debate @agent1 @agent2 <话题> — 指定 agent 辩论\n\n示例:\n/debate AI 会取代人类决策吗\n/debate @cc @gm 远程办公是否更高效"
	}

	// Parse optional agent prefixes from rest
	parsedNames, parsedMsg := h.parseCommand(rest)
	topic := strings.TrimSpace(parsedMsg)

	// If no topic after parsing, the entire rest is the topic
	if topic == "" {
		topic = rest
		parsedNames = nil
	}

	topic = strings.TrimSpace(topic)
	if topic == "" {
		return "辩论话题不能为空。示例: /debate AI 会取代人类决策吗"
	}

	// Determine debate participants
	var agentNames []string
	if len(parsedNames) >= 2 {
		// User specified agents
		agentNames = parsedNames[:2] // Take first two only
	} else {
		// Use default agents: try to get first two configured agents
		h.mu.RLock()
		metas := h.agentMetas
		h.mu.RUnlock()

		if len(metas) >= 2 {
			agentNames = []string{metas[0].Name, metas[1].Name}
		} else {
			// Fallback: use default + try to find any other agent
			defaultAg := h.getDefaultAgent()
			if defaultAg == nil {
				return "❌ 默认 agent 未就绪，请稍后重试。"
			}
			// Try common agents
			candidates := []string{"claude", "codex", "gemini", "deepseek", "qwen"}
			for _, c := range candidates {
				if c != h.defaultName {
					agentNames = []string{h.defaultName, c}
					break
				}
			}
			if len(agentNames) < 2 {
				return "❌ 可用 agent 不足，至少需要两个 agent 才能辩论。"
			}
		}
	}

	// Validate agents are available
	for _, name := range agentNames {
		if _, err := h.getAgent(ctx, name); err != nil {
			return fmt.Sprintf("❌ agent %q 不可用: %v", name, err)
		}
	}

	// Start debate asynchronously
	go h.runDebate(ctx, client, msg, agentNames[0], agentNames[1], topic, clientID)

	return fmt.Sprintf("🎭 辩论开始！\n话题: %s\n正方: %s\n反方: %s\n\n辩论进行中，结果将陆续发送给你...", topic, agentNames[0], agentNames[1])
}

// runDebate executes the debate rounds and sends results to the user.
func (h *Handler) runDebate(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage, proAgent, conAgent, topic, clientID string) {
	const rounds = 3
	var prevConReply string
	var allProReplies []string
	var allConReplies []string

	// Helper: send a standalone message (no contextToken dependency)
	sendMsg := func(text string) {
		cid := NewClientID()
		plainText := MarkdownToPlainText(text)
		req := &ilink.SendMessageRequest{
			Msg: ilink.SendMsg{
				FromUserID:   client.BotID(),
				ToUserID:     msg.FromUserID,
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
			log.Printf("[debate] failed to send: err=%v ret=%d", err, resp.Ret)
		}
	}

	// Send debate header
	sendMsg(fmt.Sprintf("🎭 **辩论: %s**\n正方: %s | 反方: %s", topic, proAgent, conAgent))

	for round := 1; round <= rounds; round++ {
		// Build pro prompt
		var proPrompt string
		if round == 1 {
			proPrompt = fmt.Sprintf(`你现在是辩论赛的正方。请针对以下话题，提出你的核心论点和论据（3-5个要点），立场鲜明地展开论述。

话题: %s
你的立场: 正方（支持/赞同）

请用清晰的逻辑、具体的例子来论证。控制在 500 字以内。`, topic)
		} else {
			proPrompt = fmt.Sprintf(`辩论继续。以下是反方上一轮的发言。请针对反方的论点进行回应和反驳，并进一步强化你的观点。

话题: %s
反方的观点:
%s

请继续你的论述。控制在 400 字以内。`, topic, prevConReply)
		}

		// Pro speaks
		proAg, proErr := h.getAgent(ctx, proAgent)
		var proReply string
		if proErr != nil {
			log.Printf("[debate] pro round %d error: %v", round, proErr)
		} else {
			proReply, proErr = proAg.Chat(ctx, msg.FromUserID+"_debate_pro", proPrompt)
			if proErr != nil {
				log.Printf("[debate] pro round %d error: %v", round, proErr)
			} else {
				log.Printf("[debate] pro round %d: %s", round, truncate(proReply, 80))
			}
		}

		// Build con prompt with pro's reply
		var conPrompt string
		if proReply != "" {
			if round == 1 {
				conPrompt = fmt.Sprintf(`你现在是辩论赛的反方。以下是正方的观点，请逐一反驳，并提出你自己的核心论点。

话题: %s
正方的观点:
%s

你的立场: 反方（反对/不赞同）

请有理有据地反驳并提出自己的观点。控制在 500 字以内。`, topic, proReply)
			} else {
				conPrompt = fmt.Sprintf(`辩论继续。以下是正方上一轮的发言。请针对正方的论点进行回应和反驳，并进一步强化你的观点。

话题: %s
正方的观点:
%s

请继续你的论述。控制在 400 字以内。`, topic, proReply)
			}
		}

		// Con speaks
		conAg, conErr := h.getAgent(ctx, conAgent)
		var conReply string
		if conErr != nil {
			log.Printf("[debate] con round %d error: %v", round, conErr)
		} else {
			conReply, conErr = conAg.Chat(ctx, msg.FromUserID+"_debate_con", conPrompt)
			if conErr != nil {
				log.Printf("[debate] con round %d error: %v", round, conErr)
			} else {
				log.Printf("[debate] con round %d: %s", round, truncate(conReply, 80))
			}
		}

		// Save replies
		prevConReply = conReply
		allProReplies = append(allProReplies, proReply)
		allConReplies = append(allConReplies, conReply)

		// Combine pro + con into one message
		var roundText string
		if proReply != "" && conReply != "" {
			roundText = fmt.Sprintf("📢 **第 %d/%d 轮**\n\n🟢 **正方 (%s):** %s\n\n🔴 **反方 (%s):** %s", round, rounds, proAgent, proReply, conAgent, conReply)
		} else if proReply != "" {
			roundText = fmt.Sprintf("📢 **第 %d/%d 轮**\n\n🟢 **正方 (%s):** %s\n\n🔴 **反方 (%s):** [出错]", round, rounds, proAgent, proReply, conAgent)
		} else {
			roundText = fmt.Sprintf("📢 **第 %d/%d 轮**\n\n🟢 **正方 (%s):** [出错]", round, rounds, proAgent)
		}
		sendMsg(roundText)
		time.Sleep(3 * time.Second)
	}

			time.Sleep(5 * time.Second)
		sendMsg("✅ 辩论结束！正在整理完整文档...")
		time.Sleep(3 * time.Second)

	// Build full markdown document
	var md strings.Builder
	md.WriteString(fmt.Sprintf("# 🎭 辩论记录：%s\n\n", topic))
	md.WriteString(fmt.Sprintf("> 正方：**%s** | 反方：**%s**\n\n---\n\n", proAgent, conAgent))
	for i := 0; i < rounds; i++ {
		md.WriteString(fmt.Sprintf("## 第 %d 轮\n\n", i+1))
		if i < len(allProReplies) && allProReplies[i] != "" {
			md.WriteString(fmt.Sprintf("### 🟢 正方 (%s)\n\n%s\n\n", proAgent, allProReplies[i]))
		}
		if i < len(allConReplies) && allConReplies[i] != "" {
			md.WriteString(fmt.Sprintf("### 🔴 反方 (%s)\n\n%s\n\n", conAgent, allConReplies[i]))
		}
		md.WriteString("---\n\n")
	}

	// Send markdown doc (split if too long, split on rune boundary)
	docText := md.String()
	runes := []rune(docText)
	const maxRuneLen = 3500
	if len(runes) <= maxRuneLen {
		sendMsg(docText)
	} else {
		for i := 0; i < len(runes); i += maxRuneLen {
			end := i + maxRuneLen
			if end > len(runes) {
				end = len(runes)
			}
			sendMsg(string(runes[i:end]))
			time.Sleep(3 * time.Second)
		}
	}

			sendMsg("💡 使用 /podcast 可以将辩论内容生成播客.")
}

// handleShell processes /sh or /$ command to execute shell commands.
func (h *Handler) handleShell(ctx context.Context, trimmed string) string {
	// Extract command: "/sh ls -la" -> "ls -la" or "/$ ls -la" -> "ls -la"
	var cmdStr string
	if strings.HasPrefix(trimmed, "/sh ") {
		cmdStr = strings.TrimPrefix(trimmed, "/sh ")
	} else {
		cmdStr = strings.TrimPrefix(trimmed, "/$ ")
	}
	cmdStr = strings.TrimSpace(cmdStr)

	if cmdStr == "" {
		return "用法: /sh <命令> 或 /$ <命令>\n示例: /sh ls -la\n可用命令: ls, cat, pwd, find, grep, head, tail 等"
	}

	// === Shortcut aliases ===
	if cmdStr == "ll" {
		cmdStr = "ls -lh"
	} else if cmdStr == ".." {
		cmdStr = "cd .."
	} else if cmdStr == "..." {
		cmdStr = "cd ../.."
	}

	// === Security: Check for dangerous operators ===
	dangerousPatterns := []string{">", ">>", "<", "|", "&&", "||", ";", "`", "$("}
	for _, pattern := range dangerousPatterns {
		if strings.Contains(cmdStr, pattern) {
			return fmt.Sprintf("❌ 出于安全考虑，不允许使用特殊字符: %s\n如需复杂操作，请在本地终端执行", pattern)
		}
	}

	// === Command whitelist for security ===
	allowedCommands := map[string]bool{
		"ls": true, "pwd": true, "cd": true, "cat": true, "head": true, "tail": true,
		"grep": true, "find": true, "wc": true, "du": true, "df": true,
		"file": true, "stat": true, "date": true, "echo": true, "basename": true,
		"dirname": true, "realpath": true, "readlink": true, "which": true,
		"tree": true, "nl": true, "sort": true, "uniq": true, "cut": true,
		"awk": true, "sed": true, "tr": true, "xargs": true,
	}

	// Extract the base command
	parts := strings.Fields(cmdStr)
	if len(parts) > 0 {
		baseCmd := parts[0]
		if !allowedCommands[baseCmd] {
			return fmt.Sprintf("❌ 命令不在白名单中: %s\n允许的命令: ls pwd cd cat head tail grep find wc du df file stat date echo basename dirname realpath readlink which tree nl sort uniq cut awk sed tr xargs\n快捷指令: ll(=ls -lh) ..(=cd ..) ...(=cd ../..)", baseCmd)
		}
	}

	// === Auto-add flags to ls for better output ===
	if strings.HasPrefix(cmdStr, "ls") {
		if !strings.Contains(cmdStr, "-C") && !strings.Contains(cmdStr, "-l") && !strings.Contains(cmdStr, "-1") {
			if cmdStr == "ls" {
				cmdStr = "ls -C"
				if isLinux() {
					cmdStr += " --group-directories-first"
				}
			} else {
				// Save the original args before modifying cmdStr
				rest := strings.TrimPrefix(cmdStr, "ls")
				cmdStr = "ls -C"
				if isLinux() {
					cmdStr += " --group-directories-first"
				}
				cmdStr += rest
			}
		} else if strings.Contains(cmdStr, "-C") && isLinux() && !strings.Contains(cmdStr, "--group-directories-first") {
			cmdStr += " --group-directories-first"
		}
	}

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Sprintf("无法获取当前目录: %v", err)
	}

	// === Large file protection for cat command ===
	if len(parts) >= 1 && parts[0] == "cat" && len(parts) >= 2 {
		filePath := parts[1]
		if !filepath.IsAbs(filePath) {
			filePath = filepath.Join(cwd, filePath)
		}
		if info, err := os.Stat(filePath); err == nil {
			// Check file size (limit to 50KB)
			const maxSize = 50 * 1024
			if info.Size() > maxSize {
				return fmt.Sprintf("⚠️ 文件过大 (%.1f MB)\n💡 建议使用:\n   tail -n 100 %s  # 查看末尾\n   head -n 100 %s  # 查看开头\n   grep \"关键词\" %s  # 搜索内容",
					float64(info.Size())/(1024*1024), filepath.Base(filePath), filepath.Base(filePath), filepath.Base(filePath))
			}
		}
	}

	// Execute command
	cmd := exec.CommandContext(ctx, "sh", "-c", cmdStr)
	cmd.Dir = cwd

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		output := stderr.String()
		if output == "" {
			output = stdout.String()
		}
		// Truncate long output
		if len(output) > 3000 {
			output = output[:3000] + "\n... (输出已截断)"
		}
		return fmt.Sprintf("❌ 命令执行失败:\n%s", output)
	}

	output := stdout.String()
	// Combine stderr if there's any stdout output
	if stderr.String() != "" {
		if output != "" {
			output += "\n"
		}
		output += stderr.String()
	}

	// Clean ANSI escape codes for WeChat
	output = cleanANSI(output)

	// Truncate long output (WeChat message has length limit)
	if len(output) > 4000 {
		output = output[:4000] + "\n... (输出已截断)"
	}

	if output == "" {
		return "✅ 命令执行成功，无输出"
	}

	return formatOutput(output)
}

// enterShellMode enters shell mode for the user.
func (h *Handler) enterShellMode(ctx context.Context, userID string) string {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Sprintf("无法获取当前目录: %v", err)
	}

	// Create or update shell mode state
	state := &shellModeState{
		enabled: true,
		cwd:     cwd,
	}
	h.shellModeStates.Store(userID, state)

	prompt := shellPrompt(cwd)
	return `--- 当前为命令行模式 (/q 退出) ---
当前目录: ` + cwd + ` (` + prompt + `)
提示: 直接输入命令即可，无需 /sh 前缀
支持 cd 切换目录，目录会持久化保存`
}

// handleShellWithState executes a command in shell mode with persistent state.
func (h *Handler) handleShellWithState(ctx context.Context, state *shellModeState, cmdStr string) string {
	cmdStr = strings.TrimSpace(cmdStr)
	if cmdStr == "" {
		return ""
	}

	// === Shortcut aliases ===
	if cmdStr == "ll" {
		cmdStr = "ls -lh"
	} else if cmdStr == ".." {
		cmdStr = "cd .."
	} else if cmdStr == "..." {
		cmdStr = "cd ../.."
	}

	// === Security: Check for dangerous operators ===
	dangerousPatterns := []string{">", ">>", "<", "|", "&&", "||", ";", "`", "$("}
	for _, pattern := range dangerousPatterns {
		if strings.Contains(cmdStr, pattern) {
			return fmt.Sprintf("❌ 出于安全考虑，不允许使用特殊字符: %s\n如需复杂操作，请在本地终端执行", pattern)
		}
	}

	// === Command whitelist for security ===
	allowedCommands := map[string]bool{
		"ls": true, "pwd": true, "cd": true, "cat": true, "head": true, "tail": true,
		"grep": true, "find": true, "wc": true, "du": true, "df": true,
		"file": true, "stat": true, "date": true, "echo": true, "basename": true,
		"dirname": true, "realpath": true, "readlink": true, "which": true,
		"tree": true, "nl": true, "sort": true, "uniq": true, "cut": true,
		"awk": true, "sed": true, "tr": true, "xargs": true,
	}

	// Extract the base command
	parts := strings.Fields(cmdStr)
	if len(parts) == 0 {
		return ""
	}
	baseCmd := parts[0]

	// Check if command is allowed
	if !allowedCommands[baseCmd] {
		return fmt.Sprintf("❌ 命令不在白名单中: %s\n允许的命令: ls pwd cd cat head tail grep find wc du df file stat date echo basename dirname realpath readlink which tree nl sort uniq cut awk sed tr xargs\n快捷指令: ll(=ls -lh) ..(=cd ..) ...(=cd ../..)", baseCmd)
	}

	// === Auto-add flags to ls for better output ===
	if baseCmd == "ls" {
		if !strings.Contains(cmdStr, "-C") && !strings.Contains(cmdStr, "-l") && !strings.Contains(cmdStr, "-1") {
			cmdStr = "ls -C"
			if isLinux() {
				cmdStr += " --group-directories-first"
			}
			cmdStr += strings.TrimPrefix(cmdStr, "ls")
		} else if strings.Contains(cmdStr, "-C") && isLinux() && !strings.Contains(cmdStr, "--group-directories-first") {
			cmdStr += " --group-directories-first"
		}
	}

	// === Handle cd command specially to update state ===
	if strings.HasPrefix(cmdStr, "cd ") || cmdStr == "cd" {
		var newDir string
		if cmdStr == "cd" {
			newDir = "~"
		} else {
			newDir = strings.TrimSpace(strings.TrimPrefix(cmdStr, "cd "))
		}

		var targetPath string
		if newDir == "" || newDir == "~" {
			home, _ := os.UserHomeDir()
			if home != "" {
				targetPath = home
			}
		} else if filepath.IsAbs(newDir) {
			targetPath = newDir
		} else {
			targetPath = filepath.Join(state.cwd, newDir)
		}

		// Resolve to absolute path
		absPath, err := filepath.Abs(targetPath)
		if err != nil {
			return fmt.Sprintf("❌ 路径解析失败: %v", err)
		}

		// Resolve symlinks to get real path (security: prevent symlink escape)
		realPath, err := filepath.EvalSymlinks(absPath)
		if err != nil {
			return fmt.Sprintf("❌ 路径解析失败: %v", err)
		}

		// Path sandboxing: check if real path is within baseDir (if set)
		if state.baseDir != "" {
			baseRealPath, err := filepath.EvalSymlinks(state.baseDir)
			if err == nil {
				relPath, err := filepath.Rel(baseRealPath, realPath)
				if err != nil || strings.HasPrefix(relPath, "..") {
					return fmt.Sprintf("❌ 路径超出允许范围: %s", newDir)
				}
			}
		}

		if info, err := os.Stat(realPath); err != nil || !info.IsDir() {
			return fmt.Sprintf("❌ 目录不存在: %s", newDir)
		}

		state.cwd = realPath

		// Auto ls after cd for better UX
		lsArgs := []string{"-C"}
		if isLinux() {
			lsArgs = append(lsArgs, "--group-directories-first")
		}
		lsCmd := exec.CommandContext(ctx, "ls", lsArgs...)
		lsCmd.Dir = state.cwd
		var lsOut bytes.Buffer
		lsCmd.Stdout = &lsOut
		if err := lsCmd.Run(); err == nil {
			lsOutput := strings.TrimSpace(lsOut.String())
			if lsOutput != "" {
				prompt := shellPrompt(state.cwd)
				return fmt.Sprintf("✅ 已切换到: %s\n%s\n```\n%s\n```", state.cwd, prompt, cleanANSI(lsOutput))
			}
		}
		return fmt.Sprintf("✅ 已切换到: %s", state.cwd)
	}

	// === Handle pwd command ===
	if cmdStr == "pwd" {
		return state.cwd
	}

	// === Large file protection for cat command ===
	if baseCmd == "cat" && len(parts) >= 2 {
		filePath := parts[1]
		if !filepath.IsAbs(filePath) {
			filePath = filepath.Join(state.cwd, filePath)
		}
		if info, err := os.Stat(filePath); err == nil {
			// Check file size (limit to 50KB)
			const maxSize = 50 * 1024
			if info.Size() > maxSize {
				return fmt.Sprintf("⚠️ 文件过大 (%.1f MB)\n💡 建议使用:\n   tail -n 100 %s  # 查看末尾\n   head -n 100 %s  # 查看开头\n   grep \"关键词\" %s  # 搜索内容",
					float64(info.Size())/(1024*1024), filepath.Base(filePath), filepath.Base(filePath), filepath.Base(filePath))
			}
		}
	}

	// === Execute command ===
	cmd := exec.CommandContext(ctx, "sh", "-c", cmdStr)
	cmd.Dir = state.cwd

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		output := stderr.String()
		if output == "" {
			output = stdout.String()
		}
		// Truncate long output
		if len(output) > 3000 {
			output = output[:3000] + "\n... (输出已截断)"
		}
		return fmt.Sprintf("❌ 命令执行失败:\n%s", output)
	}

	output := stdout.String()
	// Combine stderr if there's any stdout output
	if stderr.String() != "" {
		if output != "" {
			output += "\n"
		}
		output += stderr.String()
	}

	// Clean ANSI escape codes for WeChat
	output = cleanANSI(output)

	// Truncate long output (WeChat message has length limit)
	if len(output) > 4000 {
		output = output[:4000] + "\n... (输出已截断)"
	}

	if output == "" {
		return "✅ 命令执行成功，无输出"
	}

	return formatShellOutput(state.cwd, output)
}

// cleanANSI removes ANSI escape codes from output.
func cleanANSI(s string) string {
	// ANSI escape sequence pattern: \x1b[...m
	re := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	return re.ReplaceAllString(s, "")
}

// isLinux returns true if the OS is Linux.
func isLinux() bool {
	return runtime.GOOS == "linux"
}

// formatOutput wraps output in markdown code block for better display in WeChat.
func formatOutput(output string) string {
	if output == "" {
		return ""
	}
	// Remove trailing newlines before closing code block
	output = strings.TrimRight(output, "\n")
	return fmt.Sprintf("```\n%s\n```", output)
}

// shellPrompt generates a shell prompt string for the given directory.
func shellPrompt(cwd string) string {
	return fmt.Sprintf("%s:#", cwd)
}

// formatShellOutput wraps output with shell prompt prefix.
func formatShellOutput(cwd string, output string) string {
	if output == "" {
		return ""
	}
	prompt := shellPrompt(cwd)
	// Remove trailing newlines before closing code block
	output = strings.TrimRight(output, "\n")
	return fmt.Sprintf("%s\n```\n%s\n```", prompt, output)
}

// handleCron processes /cron commands.
// Usage:
//   /cron add "0 9 * * *" message          - add a text cron job
//   /cron add "0 9 * * *" workflow DSL...   - add a workflow cron job
//   /cron list                             - list all cron jobs
//   /cron remove <id>                      - remove a cron job
//   /cron enable <id>                      - enable a cron job
//   /cron disable <id>                     - disable a cron job
func (h *Handler) handleCron(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage, trimmed, clientID string) string {
	if h.cronManager == nil {
		return "❌ Cron 管理器未初始化"
	}

	rest := strings.TrimSpace(strings.TrimPrefix(trimmed, "/cron"))

	if rest == "" || rest == "list" {
		return h.handleCronList(ctx, msg.FromUserID)
	}

	parts := strings.Fields(rest)
	if len(parts) == 0 {
		return h.handleCronList(ctx, msg.FromUserID)
	}

	command := parts[0]
	args := parts[1:]

	switch command {
	case "add":
		return h.handleCronAdd(ctx, msg.FromUserID, args)
	case "remove", "rm", "delete":
		if len(args) < 1 {
			return "用法: /cron remove <id>"
		}
		return h.handleCronRemove(ctx, msg.FromUserID, args[0])
	case "enable":
		if len(args) < 1 {
			return "用法: /cron enable <id>"
		}
		return h.handleCronEnable(ctx, msg.FromUserID, args[0], true)
	case "disable":
		if len(args) < 1 {
			return "用法: /cron disable <id>"
		}
		return h.handleCronEnable(ctx, msg.FromUserID, args[0], false)
	default:
		// Try to parse as "add" command with implicit add
		return h.handleCronAdd(ctx, msg.FromUserID, parts)
	}
}

// handleCronList lists all cron jobs for the user.
func (h *Handler) handleCronList(ctx context.Context, userID string) string {
	jobs, err := h.cronManager.ListJobs(userID)
	if err != nil {
		return fmt.Sprintf("❌ 获取任务列表失败: %v", err)
	}

	if len(jobs) == 0 {
		return "📅 你还没有定时任务\n\n用法:\n/cron add \"0 9 * * *\" 提醒喝水\n/cron add \"0 9 * * 1\" workflow step1 @claude 生成周报"
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("📅 定时任务列表 (共 %d 个):\n\n", len(jobs)))

	for i, job := range jobs {
		status := "✅ 启用"
		if !job.Enabled {
			status = "❌ 禁用"
		}

		cmdType := "文本"
		if job.Command.Type == "workflow" {
			cmdType = "工作流"
		} else if job.Command.Type == "agent" {
			cmdType = "Agent"
		}

		content := job.Command.Content
		if len(content) > 30 {
			content = content[:30] + "..."
		}

		sb.WriteString(fmt.Sprintf("%d. [%s] %s\n", i+1, job.ID, status))
		sb.WriteString(fmt.Sprintf("   表达式: %s\n", job.CronExpr))
		sb.WriteString(fmt.Sprintf("   类型: %s\n", cmdType))
		sb.WriteString(fmt.Sprintf("   内容: %s\n\n", content))
	}

	sb.WriteString("💡 管理命令:\n")
	sb.WriteString("   /cron remove <id>   - 删除任务\n")
	sb.WriteString("   /cron enable <id>   - 启用任务\n")
	sb.WriteString("   /cron disable <id>  - 禁用任务")

	return sb.String()
}

// handleCronAdd adds a new cron job.
func (h *Handler) handleCronAdd(ctx context.Context, userID string, args []string) string {
	if len(args) < 1 {
		return "用法: /cron add <定时描述>\n   /cron add \"<cron表达式>\" <消息内容>\n\n示例:\n   /cron add 每天早上9点提醒我喝水\n   /cron add 每周一早上8点生成周报\n   /cron add \"0 9 * * *\" workflow step1 @claude 生成周报"
	}

	// Check if args[0] looks like a cron expression
	// A cron expression should have space-separated parts with special characters
	possibleCronExpr := args[0]
	possibleCronExpr = strings.Trim(possibleCronExpr, "\"")

	// Better detection: check if it looks like a cron expression (space-separated with special chars)
	// rather than natural language which might contain numbers
	isCronExpr := looksLikeCronExpr(possibleCronExpr)

	var cronExpr, cmdType, cmdContent, cmdAgent string

	if isCronExpr {
		// User provided cron expression directly
		cronExpr = possibleCronExpr

		// Validate cron expression
		if err := validateCronExpr(cronExpr); err != nil {
			return fmt.Sprintf("❌ Cron 表达式无效: %v\n\n格式: 秒 分 时 日 月 周\n示例: \"0 9 * * *\" - 每天 9:00", err)
		}

		// Parse the rest of args
		if len(args) < 2 {
			return "❌ 请提供消息内容\n用法: /cron add \"0 9 * * *\" 消息内容"
		}

		// Check if second arg is a command type
		switch args[1] {
		case "workflow", "wf":
			if len(args) < 3 {
				return "❌ 工作流内容不能为空\n用法: /cron add \"0 9 * * *\" workflow step1 @claude 分析..."
			}
			cmdType = "workflow"
			cmdContent = strings.Join(args[2:], " ")
		case "agent":
			if len(args) < 4 {
				return "❌ Agent 任务格式错误\n用法: /cron add \"0 9 * * *\" agent @claude 消息内容"
			}
			cmdType = "agent"
			cmdAgent = args[2]
			cmdContent = strings.Join(args[3:], " ")
		default:
			// Default to text type
			cmdType = "text"
			cmdContent = strings.Join(args[1:], " ")
		}
	} else {
		// User provided natural language - use AI to parse
		naturalLang := strings.Join(args, " ")
		return h.handleNaturalLanguageCron(ctx, userID, naturalLang)
	}

	// Generate job ID
	jobID := fmt.Sprintf("cron_%d", time.Now().UnixNano())

	job := &CronJob{
		ID:       jobID,
		UserID:   userID,
		CronExpr: cronExpr,
		Command: CronCommand{
			Type:    cmdType,
			Content: cmdContent,
			Agent:   cmdAgent,
		},
		Enabled:   true,
		CreatedAt: time.Now().Unix(),
	}

	if err := h.cronManager.AddJob(job); err != nil {
		return fmt.Sprintf("❌ 添加任务失败: %v", err)
	}

	return fmt.Sprintf("✅ 定时任务已添加\n\nID: %s\n表达式: %s\n类型: %s\n内容: %s", jobID, cronExpr, cmdType, truncate(cmdContent, 50))
}

// handleNaturalLanguageCron parses natural language and creates a cron job.
// Uses the "Assistant" HTTP agent to avoid function calling issues.
// Falls back to rule-based parsing if AI fails or returns invalid result.
func (h *Handler) handleNaturalLanguageCron(ctx context.Context, userID, naturalLang string) string {
	log.Printf("[cron] Attempting AI parsing for: %q", naturalLang)

	// Use getAgent to start Assistant on demand if not already running
	assistantAgent, err := h.getAgent(ctx, "Assistant")
	if err != nil {
		log.Printf("[cron] Failed to get Assistant agent, falling back to rule-based parsing: %v", err)
		return h.handleNaturalLanguageCronRuleBased(ctx, userID, naturalLang)
	}

	log.Printf("[cron] Using Assistant agent for parsing")

	// Use AI to parse natural language
	prompt := fmt.Sprintf(`你是一个 cron 表达式生成器。请将用户的时间描述转换为标准的 cron 格式。

用户描述：%s

请直接输出 JSON 格式（不要使用代码块，不要使用 markdown）：

{
  "cron_expr": "0 8 2 * * *",
  "message": "根据用户描述生成的消息内容",
  "type": "text"
}

cron 格式说明（6位）：秒 分 时 日 月 周
- 每天9点 → 0 0 9 * * *
- 每天2:08 → 0 8 2 * * *
- 每周一8点 → 0 0 8 * * 1
- 每30分钟 → 0 */30 * * * *
- 每天早上9点提醒开会 → {"cron_expr": "0 0 9 * * *", "message": "提醒开会", "type": "text"}

注意：小时范围0-23，分钟范围0-59。

现在请为用户描述生成 JSON：`, naturalLang)

	reply, err := assistantAgent.Chat(ctx, userID+"_cron_parse", prompt)
	if err != nil {
		// Fallback to rule-based parsing on error
		log.Printf("[cron] AI parsing failed, falling back to rule-based: %v", err)
		return h.handleNaturalLanguageCronRuleBased(ctx, userID, naturalLang)
	}

	log.Printf("[cron] AI response: %q", truncate(reply, 200))

	// Extract JSON from reply
	jsonStart := strings.Index(reply, "{")
	jsonEnd := strings.LastIndex(reply, "}")
	if jsonStart == -1 || jsonEnd == -1 || jsonEnd <= jsonStart {
		// Fallback to rule-based parsing on JSON extraction error
		log.Printf("[cron] AI response doesn't contain valid JSON, falling back to rule-based")
		return h.handleNaturalLanguageCronRuleBased(ctx, userID, naturalLang)
	}

	jsonStr := reply[jsonStart : jsonEnd+1]

	// Parse JSON
	var result struct {
		CronExpr string `json:"cron_expr"`
		Message  string `json:"message"`
		Type     string `json:"type"`
	}

	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		// Fallback to rule-based parsing on JSON parse error
		log.Printf("[cron] Failed to parse AI response JSON, falling back to rule-based: %v", err)
		return h.handleNaturalLanguageCronRuleBased(ctx, userID, naturalLang)
	}

	log.Printf("[cron] AI parsed: cron_expr=%q message=%q type=%q", result.CronExpr, result.Message, result.Type)

	// Strictly validate cron expression
	if err := validateCronExpr(result.CronExpr); err != nil {
		log.Printf("[cron] AI returned invalid cron expression: %v, falling back to rule-based", err)
		return h.handleNaturalLanguageCronRuleBased(ctx, userID, naturalLang)
	}

	// Additional validation: check hour and minute ranges
	parts := strings.Fields(result.CronExpr)
	if len(parts) >= 3 {
		// Check minute field (index 1)
		if parts[1] != "*" {
			minStr := strings.TrimPrefix(parts[1], "*/")
			if min, err := strconv.Atoi(minStr); err == nil {
				if min < 0 || min > 59 {
					log.Printf("[cron] AI returned invalid minute value: %d, falling back to rule-based", min)
					return h.handleNaturalLanguageCronRuleBased(ctx, userID, naturalLang)
				}
			}
		}
		// Check hour field (index 2)
		if parts[2] != "*" {
			hourStr := strings.TrimPrefix(parts[2], "*/")
			if hour, err := strconv.Atoi(hourStr); err == nil {
				if hour < 0 || hour > 23 {
					log.Printf("[cron] AI returned invalid hour value: %d, falling back to rule-based", hour)
					return h.handleNaturalLanguageCronRuleBased(ctx, userID, naturalLang)
				}
			}
		}
	}

	// Validate type
	if result.Type != "text" && result.Type != "workflow" && result.Type != "agent" {
		log.Printf("[cron] AI returned invalid type: %q, falling back to rule-based", result.Type)
		return h.handleNaturalLanguageCronRuleBased(ctx, userID, naturalLang)
	}

	log.Printf("[cron] AI parsing successful, creating job")

	// Generate job ID
	jobID := fmt.Sprintf("cron_%d", time.Now().UnixNano())

	var cmdAgent string
	if result.Type == "agent" {
		// Extract agent from message
		parts := strings.Fields(result.Message)
		if len(parts) > 0 && strings.HasPrefix(parts[0], "@") {
			cmdAgent = strings.TrimPrefix(parts[0], "@")
			result.Message = strings.Join(parts[1:], " ")
		}
	}

	job := &CronJob{
		ID:       jobID,
		UserID:   userID,
		CronExpr: result.CronExpr,
		Command: CronCommand{
			Type:    result.Type,
			Content: result.Message,
			Agent:   cmdAgent,
		},
		Enabled:   true,
		CreatedAt: time.Now().Unix(),
	}

	if err := h.cronManager.AddJob(job); err != nil {
		return fmt.Sprintf("❌ 添加任务失败: %v", err)
	}

	return fmt.Sprintf("✅ 定时任务已添加（AI 解析）\n\nID: %s\n表达式: %s\n类型: %s\n内容: %s", jobID, result.CronExpr, result.Type, truncate(result.Message, 50))
}

// handleNaturalLanguageCronRuleBased is a fallback rule-based parser.
func (h *Handler) handleNaturalLanguageCronRuleBased(ctx context.Context, userID, naturalLang string) string {
	// Parse natural language using simple pattern matching
	cronExpr, message, cmdType := parseNaturalLanguageCron(naturalLang)

	if cronExpr == "" {
		return "❌ 无法解析时间描述\n\n支持的格式：\n" +
			"  每天9点提醒我开会\n" +
			"  每周一早上8点生成周报\n" +
			"  每30分钟检查状态\n" +
			"  每天早上9点\n\n" +
			"或使用标准 cron 表达式：\n" +
			"  /cron add \"0 9 * * *\" 消息内容"
	}

	// Generate job ID
	jobID := fmt.Sprintf("cron_%d", time.Now().UnixNano())

	var cmdAgent string
	if cmdType == "agent" {
		// Extract agent from message
		parts := strings.Fields(message)
		if len(parts) > 0 && strings.HasPrefix(parts[0], "@") {
			cmdAgent = strings.TrimPrefix(parts[0], "@")
			message = strings.Join(parts[1:], " ")
		}
	}

	job := &CronJob{
		ID:       jobID,
		UserID:   userID,
		CronExpr: cronExpr,
		Command: CronCommand{
			Type:    cmdType,
			Content: message,
			Agent:   cmdAgent,
		},
		Enabled:   true,
		CreatedAt: time.Now().Unix(),
	}

	if err := h.cronManager.AddJob(job); err != nil {
		return fmt.Sprintf("❌ 添加任务失败: %v", err)
	}

	return fmt.Sprintf("✅ 定时任务已添加\n\nID: %s\n表达式: %s\n类型: %s\n内容: %s", jobID, cronExpr, cmdType, truncate(message, 50))
}

// parseNaturalLanguageCron parses natural language time descriptions into cron expressions.
// Returns (cronExpr, message, cmdType). If parsing fails, returns ("", "", "").
func parseNaturalLanguageCron(input string) (string, string, string) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", "", ""
	}

	// Check for workflow keyword
	cmdType := "text"
	if strings.Contains(input, "workflow") || strings.Contains(input, "工作流") {
		cmdType = "workflow"
	}

	// Extract time part and message part
	// Common patterns: "每天9点提醒", "每周一早上8点", "每30分钟检查"
	var timePart, messagePart string

	// Try to split by common action verbs
	actionVerbs := []string{"提醒", "发送", "推送", "通知", "检查", "生成", "汇报", "开会", "消息"}
	for _, verb := range actionVerbs {
		idx := strings.Index(input, verb)
		if idx > 0 {
			timePart = strings.TrimSpace(input[:idx])
			messagePart = strings.TrimSpace(input[idx:])
			break
		}
	}

	// If no action verb found, try other patterns
	if timePart == "" {
		// Pattern: "每天9点" (no message) or "每天9点 "
		if strings.Contains(input, "每天") || strings.Contains(input, "每日") {
			idx := strings.IndexAny(input, "0123456789每")
			if idx >= 0 {
				// Find where the time description ends
				endIdx := len(input)
				for _, sep := range []string{" ", "，", ","} {
					if j := strings.Index(input[idx:], sep); j > 0 && idx+j < endIdx {
						endIdx = idx + j
					}
				}
				timePart = strings.TrimSpace(input[:endIdx])
				messagePart = strings.TrimSpace(input[endIdx:])
			}
		} else if strings.Contains(input, "每周") {
			idx := strings.Index(input, "每周")
			if idx >= 0 {
				// Find where it ends
				endIdx := len(input)
				for _, sep := range []string{" ", "，", ","} {
					if j := strings.Index(input[idx:], sep); j > 0 && idx+j < endIdx {
						endIdx = idx + j
					}
				}
				timePart = strings.TrimSpace(input[:endIdx])
				messagePart = strings.TrimSpace(input[endIdx:])
			}
		} else if strings.Contains(input, "每") && (strings.Contains(input, "分钟") || strings.Contains(input, "小时")) {
			// Every N minutes/hours
			idx := strings.Index(input, "每")
			if idx >= 0 {
				endIdx := len(input)
				for _, sep := range []string{" ", "，", ","} {
					if j := strings.Index(input[idx:], sep); j > 0 && idx+j < endIdx {
						endIdx = idx + j
					}
				}
				timePart = strings.TrimSpace(input[:endIdx])
				messagePart = strings.TrimSpace(input[endIdx:])
			}
		} else {
			// Whole thing is time part, no message
			timePart = input
		}
	}

	// If still no time part, use whole input as time
	if timePart == "" {
		timePart = input
	}

	// Default message if none extracted
	if messagePart == "" {
		messagePart = "定时提醒"
	}

	// Parse time patterns
	// Pattern: 每天9点 → 0 0 9 * * *
	// Pattern: 每天2:08 → 0 8 2 * * *
	if strings.Contains(timePart, "每天") || strings.Contains(timePart, "每日") {
		hour, minute := extractTime(timePart)
		if hour >= 0 {
			return fmt.Sprintf("0 %d %d * * *", minute, hour), messagePart, cmdType
		}
	}

	// Pattern: 每周一/二/三/四/五/六/日 X点
	weekdayMap := map[string]int{
		"一": 1, "二": 2, "三": 3, "四": 4, "五": 5, "六": 6, "日": 0, "天": 0,
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "0": 0, "7": 0,
	}
	for day, num := range weekdayMap {
		pattern := "每周" + day
		if strings.Contains(timePart, pattern) {
			hour, minute := extractTime(timePart)
			if hour >= 0 {
				return fmt.Sprintf("0 %d %d * * %d", minute, hour, num), messagePart, cmdType
			}
		}
	}

	// Pattern: 每30分钟、每1小时
	if strings.Contains(timePart, "每") && strings.Contains(timePart, "分钟") {
		// Extract number
		num := extractNumber(timePart)
		if num > 0 && num <= 60 {
			return fmt.Sprintf("0 */%d * * * *", num), messagePart, cmdType
		}
	}

	if strings.Contains(timePart, "每") && strings.Contains(timePart, "小时") {
		num := extractNumber(timePart)
		if num > 0 && num <= 24 {
			return fmt.Sprintf("0 0 */%d * * *", num), messagePart, cmdType
		}
	}

	// Pattern: 早上/中午/晚上 X点 or X:Y
	hour, minute := extractTime(timePart)
	if hour >= 0 {
		return fmt.Sprintf("0 %d %d * * *", minute, hour), messagePart, cmdType
	}

	return "", "", ""
}

// extractHour extracts the hour (0-23) from a time description.
// Returns -1 if not found.
func extractHour(s string) int {
	hour, _ := extractTime(s)
	return hour
}

// extractTime extracts both hour and minute from a time description.
// Returns (hour, minute). If minute is not found, returns 0.
// If hour is not found, returns (-1, 0).
func extractTime(s string) (int, int) {
	// Look for patterns like "2:08", "9:30"
	re := regexp.MustCompile(`(\d{1,2}):(\d{2})`)
	matches := re.FindStringSubmatch(s)
	if len(matches) >= 3 {
		hour, err1 := strconv.Atoi(matches[1])
		minute, err2 := strconv.Atoi(matches[2])
		if err1 == nil && err2 == nil && hour >= 0 && hour <= 23 && minute >= 0 && minute <= 59 {
			return hour, minute
		}
	}

	// Look for patterns like "9点8分", "9点08分", "23点30分"
	re2 := regexp.MustCompile(`(\d{1,2})点(\d{1,2})分`)
	matches = re2.FindStringSubmatch(s)
	if len(matches) >= 3 {
		hour, err1 := strconv.Atoi(matches[1])
		minute, err2 := strconv.Atoi(matches[2])
		if err1 == nil && err2 == nil && hour >= 0 && hour <= 23 && minute >= 0 && minute <= 59 {
			return hour, minute
		}
	}

	// Look for patterns like "9点", "09点", "23点" (no minute specified)
	re3 := regexp.MustCompile(`(\d{1,2})点`)
	matches = re3.FindStringSubmatch(s)
	if len(matches) >= 2 {
		hour, err := strconv.Atoi(matches[1])
		if err == nil && hour >= 0 && hour <= 23 {
			return hour, 0
		}
	}

	// Time words
	timeWords := map[string]int{
		"凌晨": 0, "午夜": 0, "零点": 0,
		"早上": 8, "早晨": 8, "上午": 9,
		"中午": 12, "下午": 14,
		"晚上": 18, "傍晚": 18,
	}
	for word, hour := range timeWords {
		if strings.Contains(s, word) {
			// Check if a specific time follows
			h, m := extractTime(strings.ReplaceAll(s, word, " "))
			if h >= 0 {
				return h, m
			}
			return hour, 0
		}
	}

	return -1, 0
}

// extractNumber extracts the first number found in the string.
func extractNumber(s string) int {
	re := regexp.MustCompile(`(\d+)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) >= 2 {
		num, err := strconv.Atoi(matches[1])
		if err == nil {
			return num
		}
	}
	return -1
}

// handleCronRemove removes a cron job.
func (h *Handler) handleCronRemove(ctx context.Context, userID, id string) string {
	// First check if job belongs to user
	jobs, err := h.cronManager.ListJobs(userID)
	if err != nil {
		return fmt.Sprintf("❌ 获取任务列表失败: %v", err)
	}

	found := false
	for _, job := range jobs {
		if job.ID == id || fmt.Sprintf("%d", extractIDFromInput(id, len(jobs))) == job.ID {
			found = true
			break
		}
	}

	if !found {
		return fmt.Sprintf("❌ 任务 %q 不存在或无权删除", id)
	}

	// Get actual job ID
	jobID := id
	for _, job := range jobs {
		if job.ID == id {
			jobID = job.ID
			break
		}
	}

	if err := h.cronManager.RemoveJob(jobID); err != nil {
		return fmt.Sprintf("❌ 删除任务失败: %v", err)
	}

	return fmt.Sprintf("✅ 任务 %s 已删除", id)
}

// handleCronEnable enables or disables a cron job.
func (h *Handler) handleCronEnable(ctx context.Context, userID, id string, enable bool) string {
	// Find the job
	jobs, err := h.cronManager.ListJobs(userID)
	if err != nil {
		return fmt.Sprintf("❌ 获取任务列表失败: %v", err)
	}

	var targetJob *CronJob
	for _, job := range jobs {
		if job.ID == id || job.ID == fmt.Sprintf("cron_%s", id) {
			targetJob = job
			break
		}
	}

	if targetJob == nil {
		return fmt.Sprintf("❌ 任务 %q 不存在", id)
	}

	targetJob.Enabled = enable

	if err := h.cronManager.UpdateJob(targetJob); err != nil {
		return fmt.Sprintf("❌ 更新任务失败: %v", err)
	}

	status := "启用"
	if !enable {
		status = "禁用"
	}

	return fmt.Sprintf("✅ 任务 %s 已%s", id, status)
}

// looksLikeCronExpr checks if a string looks like a cron expression
// rather than natural language. Cron expressions have specific patterns:
// - Space-separated parts (5 or 6)
// - Contains special characters: * , / - ?
// - Each part is mostly numeric/special chars
func looksLikeCronExpr(s string) bool {
	// Empty string is not a cron expr
	if s == "" {
		return false
	}

	// Check if it's wrapped in quotes - user is explicitly providing a cron expr
	if (strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"")) ||
		(strings.HasPrefix(s, "'") && strings.HasSuffix(s, "'")) {
		return true
	}

	// Split by spaces and check the parts
	parts := strings.Fields(s)

	// Cron expressions have 5 or 6 parts
	if len(parts) < 5 || len(parts) > 6 {
		return false
	}

	// Check each part looks like a cron field
	cronSpecialChars := "*,-/?"
	for _, part := range parts {
		// Each part should contain mostly special chars or digits
		hasDigit := false
		hasSpecial := false
		alphaCount := 0

		for _, ch := range part {
			if ch >= '0' && ch <= '9' {
				hasDigit = true
			} else if strings.ContainsRune(cronSpecialChars, ch) {
				hasSpecial = true
			} else if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
				alphaCount++
			}
		}

		// Valid cron part: has special chars OR (has digits and minimal letters)
		// Allow things like "MON" or "JAN" but reject longer natural language
		if !hasDigit && !hasSpecial {
			return false
		}

		// Too many letters suggest natural language, not cron
		if alphaCount > 3 {
			return false
		}
	}

	return true
}

// validateCronExpr validates a cron expression.
func validateCronExpr(expr string) error {
	// Basic validation: check if it has 5 or 6 parts
	parts := strings.Fields(expr)
	if len(parts) < 5 || len(parts) > 6 {
		return fmt.Errorf("cron 表达式应有 5 或 6 个部分，实际有 %d 个", len(parts))
	}

	// For now, just check format - cron library will validate further
	return nil
}

// extractIDFromInput extracts numeric ID from user input like "1", "2", etc.
func extractIDFromInput(input string, maxJobs int) int {
	// Try to parse as number
	var id int
	if _, err := fmt.Sscanf(input, "%d", &id); err == nil && id > 0 && id <= maxJobs {
		return id
	}
	return -1
}

// getDefaultAgentName returns the current default agent name.
func (h *Handler) getDefaultAgentName() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.defaultName
}
