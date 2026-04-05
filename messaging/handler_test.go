package messaging

import (
	"strings"
	"testing"

	"github.com/fastclaw-ai/weclaw/agent"
)

func newTestHandler() *Handler {
	return &Handler{agents: make(map[string]agent.Agent)}
}

func TestParseCommand_NoPrefix(t *testing.T) {
	h := newTestHandler()
	names, msg := h.parseCommand("hello world")
	if len(names) != 0 {
		t.Errorf("expected nil names, got %v", names)
	}
	if msg != "hello world" {
		t.Errorf("expected full text, got %q", msg)
	}
}

func TestParseCommand_SlashWithAgent(t *testing.T) {
	h := newTestHandler()
	names, msg := h.parseCommand("/claude explain this code")
	if len(names) != 1 || names[0] != "claude" {
		t.Errorf("expected [claude], got %v", names)
	}
	if msg != "explain this code" {
		t.Errorf("expected 'explain this code', got %q", msg)
	}
}

func TestParseCommand_AtPrefix(t *testing.T) {
	h := newTestHandler()
	names, msg := h.parseCommand("@claude explain this code")
	if len(names) != 1 || names[0] != "claude" {
		t.Errorf("expected [claude], got %v", names)
	}
	if msg != "explain this code" {
		t.Errorf("expected 'explain this code', got %q", msg)
	}
}

func TestParseCommand_MultiAgent(t *testing.T) {
	h := newTestHandler()
	names, msg := h.parseCommand("@cc @cx hello")
	if len(names) != 2 || names[0] != "claude" || names[1] != "codex" {
		t.Errorf("expected [claude codex], got %v", names)
	}
	if msg != "hello" {
		t.Errorf("expected 'hello', got %q", msg)
	}
}

func TestParseCommand_MultiAgentDedup(t *testing.T) {
	h := newTestHandler()
	names, msg := h.parseCommand("@cc @cc hello")
	if len(names) != 1 || names[0] != "claude" {
		t.Errorf("expected [claude] (deduped), got %v", names)
	}
	if msg != "hello" {
		t.Errorf("expected 'hello', got %q", msg)
	}
}

func TestParseCommand_SwitchOnly(t *testing.T) {
	h := newTestHandler()
	names, msg := h.parseCommand("/claude")
	if len(names) != 1 || names[0] != "claude" {
		t.Errorf("expected [claude], got %v", names)
	}
	if msg != "" {
		t.Errorf("expected empty message, got %q", msg)
	}
}

func TestParseCommand_Alias(t *testing.T) {
	h := newTestHandler()
	names, msg := h.parseCommand("/cc write a function")
	if len(names) != 1 || names[0] != "claude" {
		t.Errorf("expected [claude] from /cc alias, got %v", names)
	}
	if msg != "write a function" {
		t.Errorf("expected 'write a function', got %q", msg)
	}
}

func TestParseCommand_CustomAlias(t *testing.T) {
	h := newTestHandler()
	h.customAliases = map[string]string{"ai": "claude", "c": "claude"}
	names, msg := h.parseCommand("/ai hello")
	if len(names) != 1 || names[0] != "claude" {
		t.Errorf("expected [claude] from custom alias, got %v", names)
	}
	if msg != "hello" {
		t.Errorf("expected 'hello', got %q", msg)
	}
}

func TestResolveAlias(t *testing.T) {
	h := newTestHandler()
	tests := map[string]string{
		"cc":  "claude",
		"cx":  "codex",
		"oc":  "openclaw",
		"cs":  "cursor",
		"km":  "kimi",
		"gm":  "gemini",
		"ocd": "opencode",
	}
	for alias, want := range tests {
		got := h.resolveAlias(alias)
		if got != want {
			t.Errorf("resolveAlias(%q) = %q, want %q", alias, got, want)
		}
	}
	if got := h.resolveAlias("unknown"); got != "unknown" {
		t.Errorf("resolveAlias(unknown) = %q, want %q", got, "unknown")
	}
	h.customAliases = map[string]string{"cc": "custom-claude"}
	if got := h.resolveAlias("cc"); got != "custom-claude" {
		t.Errorf("resolveAlias(cc) with custom = %q, want custom-claude", got)
	}
}

func TestBuildHelpText(t *testing.T) {
	text := buildHelpText()
	if text == "" {
		t.Error("help text is empty")
	}
	if !strings.Contains(text, "/info") {
		t.Error("help text should mention /info")
	}
	if !strings.Contains(text, "/help") {
		t.Error("help text should mention /help")
	}
}

func TestParseCommand_PodcastBuiltin(t *testing.T) {
	h := newTestHandler()

	// Test /podcast alone - should not be parsed as agent name
	names, msg := h.parseCommand("/podcast some text")
	if len(names) != 0 {
		t.Errorf("expected no agent names, got %v", names)
	}
	if msg != "/podcast some text" {
		t.Errorf("expected '/podcast some text', got %q", msg)
	}

	// Test /podcast alone
	names, msg = h.parseCommand("/podcast")
	if len(names) != 0 {
		t.Errorf("expected no agent names, got %v", names)
	}
	if msg != "/podcast" {
		t.Errorf("expected '/podcast', got %q", msg)
	}
}

func TestParseCommand_PodcastWithAgentPrefix(t *testing.T) {
	h := newTestHandler()

	// Test @cc /podcast - should intercept /podcast and not treat as agent command
	names, msg := h.parseCommand("@cc /podcast some text")
	// The parser should recognize @cc as agent, but then /podcast as builtin command
	// So it returns the original text starting from /podcast
	if len(names) != 1 || names[0] != "claude" {
		t.Errorf("expected [claude], got %v", names)
	}
	if msg != "/podcast some text" {
		t.Errorf("expected '/podcast some text', got %q", msg)
	}

	// Test /claude /podcast - similar behavior
	names, msg = h.parseCommand("/claude /podcast some text")
	if len(names) != 1 || names[0] != "claude" {
		t.Errorf("expected [claude], got %v", names)
	}
	if msg != "/podcast some text" {
		t.Errorf("expected '/podcast some text', got %q", msg)
	}
}

func TestIsBuiltinCommand_Podcast(t *testing.T) {
	// Test /podcast variants
	if !isBuiltinCommand("/podcast") {
		t.Error("/podcast should be a builtin command")
	}
	if !isBuiltinCommand("/podcast some text") {
		t.Error("/podcast with text should be a builtin command")
	}
	if isBuiltinCommand("/podcasting") {
		t.Error("/podcasting should NOT be a builtin command")
	}
}

func TestParseCommand_DebateBuiltin(t *testing.T) {
	h := newTestHandler()

	// Test /debate alone - should not be parsed as agent name
	names, msg := h.parseCommand("/debate AI 会取代人类吗")
	if len(names) != 0 {
		t.Errorf("expected no agent names, got %v", names)
	}
	if msg != "/debate AI 会取代人类吗" {
		t.Errorf("expected '/debate AI 会取代人类吗', got %q", msg)
	}

	// Test /debate with agent prefix
	names, msg = h.parseCommand("@cc /debate AI 会取代人类吗")
	if len(names) != 1 || names[0] != "claude" {
		t.Errorf("expected [claude], got %v", names)
	}
	if msg != "/debate AI 会取代人类吗" {
		t.Errorf("expected '/debate AI 会取代人类吗', got %q", msg)
	}
}

func TestIsBuiltinCommand_Debate(t *testing.T) {
	if !isBuiltinCommand("/debate") {
		t.Error("/debate should be a builtin command")
	}
	if !isBuiltinCommand("/debate some topic") {
		t.Error("/debate with text should be a builtin command")
	}
	if isBuiltinCommand("/debating") {
		t.Error("/debating should NOT be a builtin command")
	}
}
