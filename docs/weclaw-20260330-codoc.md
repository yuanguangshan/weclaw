# Project Documentation

- **Generated at:** 2026-03-30 16:40:10
- **Root Dir:** `/Users/ygs/ygs/weclaw`
- **File Count:** 50
- **Total Size:** 488.15 KB

<a name="toc"></a>
## 📂 扫描目录
- [📄 .air.toml](#airtoml) (52 lines, 0.91 KB)
- [📄 .dockerignore](#dockerignore) (13 lines, 0.09 KB)
- [📄 .gitignore](#gitignore) (23 lines, 0.16 KB)
- [📄 Dockerfile](#dockerfile) (16 lines, 0.35 KB)
- [📄 LICENSE](#license) (21 lines, 1.04 KB)
- [📄 Makefile](#makefile) (2 lines, 0.03 KB)
- [📄 README.md](#readmemd) (343 lines, 8.63 KB)
- [📄 agent/acp_agent.go](#agentacp_agentgo) (1318 lines, 33.34 KB)
- [📄 agent/agent.go](#agentagentgo) (136 lines, 4.45 KB)
- [📄 agent/cli_agent.go](#agentcli_agentgo) (304 lines, 8.77 KB)
- [📄 agent/env_test.go](#agentenv_testgo) (62 lines, 1.50 KB)
- [📄 agent/http_agent.go](#agenthttp_agentgo) (194 lines, 5.23 KB)
- [📄 api/server.go](#apiservergo) (119 lines, 3.14 KB)
- [📄 cmd/login.go](#cmdlogingo) (30 lines, 0.56 KB)
- [📄 cmd/proc_unix.go](#cmdproc_unixgo) (12 lines, 0.16 KB)
- [📄 cmd/proc_windows.go](#cmdproc_windowsgo) (9 lines, 0.15 KB)
- [📄 cmd/restart.go](#cmdrestartgo) (40 lines, 0.72 KB)
- [📄 cmd/root.go](#cmdrootgo) (27 lines, 0.50 KB)
- [📄 cmd/send.go](#cmdsendgo) (68 lines, 1.84 KB)
- [📄 cmd/start.go](#cmdstartgo) (435 lines, 11.48 KB)
- [📄 cmd/status.go](#cmdstatusgo) (31 lines, 0.56 KB)
- [📄 cmd/stop.go](#cmdstopgo) (21 lines, 0.31 KB)
- [📄 cmd/update.go](#cmdupdatego) (207 lines, 4.63 KB)
- [📄 config/config.go](#configconfiggo) (141 lines, 4.21 KB)
- [📄 config/config_test.go](#configconfig_testgo) (119 lines, 2.53 KB)
- [📄 config/detect.go](#configdetectgo) (281 lines, 9.21 KB)
- [📄 config/detect_test.go](#configdetect_testgo) (82 lines, 2.50 KB)
- [📄 docs/README_CN.md](#docsreadme_cnmd) (345 lines, 9.32 KB)
- [📄 docs/weclaw-20260330-docs.md](#docsweclaw-20260330-docsmd) (8906 lines, 233.33 KB)
- [📄 docs/项目学习.md](#docsmd) (1384 lines, 39.32 KB)
- [📄 go.mod](#gomod) (15 lines, 0.43 KB)
- [📄 go.sum](#gosum) (26 lines, 2.09 KB)
- [📄 ilink/auth.go](#ilinkauthgo) (177 lines, 3.96 KB)
- [📄 ilink/client.go](#ilinkclientgo) (224 lines, 5.66 KB)
- [📄 ilink/monitor.go](#ilinkmonitorgo) (181 lines, 4.60 KB)
- [📄 ilink/types.go](#ilinktypesgo) (219 lines, 6.62 KB)
- [📄 install.sh](#installsh) (64 lines, 1.60 KB)
- [📄 main.go](#maingo) (7 lines, 0.09 KB)
- [📄 messaging/attachment.go](#messagingattachmentgo) (127 lines, 2.90 KB)
- [📄 messaging/attachment_test.go](#messagingattachment_testgo) (100 lines, 2.96 KB)
- [📄 messaging/cdn.go](#messagingcdngo) (232 lines, 6.56 KB)
- [📄 messaging/handler.go](#messaginghandlergo) (1175 lines, 36.32 KB)
- [📄 messaging/handler_test.go](#messaginghandler_testgo) (140 lines, 3.60 KB)
- [📄 messaging/linkhoard.go](#messaginglinkhoardgo) (326 lines, 8.66 KB)
- [📄 messaging/markdown.go](#messagingmarkdowngo) (103 lines, 3.01 KB)
- [📄 messaging/media.go](#messagingmediago) (213 lines, 5.31 KB)
- [📄 messaging/media_test.go](#messagingmedia_testgo) (73 lines, 1.81 KB)
- [📄 messaging/sender.go](#messagingsendergo) (86 lines, 2.21 KB)
- [📄 service/com.fastclaw.weclaw.plist](#servicecomfastclawweclawplist) (21 lines, 0.58 KB)
- [📄 service/weclaw.service](#serviceweclawservice) (14 lines, 0.24 KB)

---

## .air.toml

```toml
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = ["start", "-f"]
  bin = "./weclaw"
  cmd = "go build -o ./weclaw ."
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata", "debug"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html"]
  include_file = []
  kill_delay = "0s"
  log = "build-errors.log"
  poll = false
  poll_interval = 0
  post_cmd = []
  pre_cmd = []
  rerun = false
  rerun_delay = 500
  send_interrupt = false
  stop_on_error = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  main_only = false
  silent = false
  time = false

[misc]
  clean_on_exit = false

[proxy]
  app_port = 0
  enabled = false
  proxy_port = 0

[screen]
  clear_on_rebuild = false
  keep_scroll = true

```

[⬆ 回到目录](#toc)

## .dockerignore

```text
weclaw
tmp/
.git/
.idea/
.vscode/
.claude/
.env
*.local
.DS_Store
Thumbs.db
*.swp
*.swo
*~

```

[⬆ 回到目录](#toc)

## .gitignore

```text
# Binary
weclaw

# Air hot reload
tmp/

# IDE
.idea/
.vscode/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db

# Environment & config
.env
*.local

# Claude Code
.claude/

```

[⬆ 回到目录](#toc)

## Dockerfile

```text
FROM golang:1.25-alpine AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /usr/local/bin/weclaw .

FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder /usr/local/bin/weclaw /usr/local/bin/weclaw

VOLUME /root/.weclaw
ENTRYPOINT ["weclaw"]
CMD ["start"]

```

[⬆ 回到目录](#toc)

## LICENSE

```text
MIT License

Copyright (c) 2026 fastclaw-ai

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```

[⬆ 回到目录](#toc)

## Makefile

```text
dev:
	air -c .air.toml start
```

[⬆ 回到目录](#toc)

## README.md

```markdown
# WeClaw

[中文文档](README_CN.md)

WeChat AI Agent Bridge — connect WeChat to AI agents (Claude, Codex, Gemini, Kimi, etc.).

> This project is inspired by [@tencent-weixin/openclaw-weixin](https://npmx.dev/package/@tencent-weixin/openclaw-weixin). For personal learning only, not for commercial use.

| | | |
|:---:|:---:|:---:|
| <img src="previews/preview1.png" width="280" /> | <img src="previews/preview2.png" width="280" /> | <img src="previews/preview3.png" width="280" /> |

## Quick Start

```bash
# One-line install
curl -sSL https://raw.githubusercontent.com/fastclaw-ai/weclaw/main/install.sh | sh

# Start (first run will prompt QR code login)
weclaw start
```

That's it. On first start, WeClaw will:
1. Show a QR code — scan with WeChat to login
2. Auto-detect installed AI agents (Claude, Codex, Gemini, etc.)
3. Save config to `~/.weclaw/config.json`
4. Start receiving and replying to WeChat messages

Use `weclaw login` to add additional WeChat accounts.

### Other install methods

```bash
# Via Go
go install github.com/fastclaw-ai/weclaw@latest

# Via Docker
docker run -it -v ~/.weclaw:/root/.weclaw ghcr.io/fastclaw-ai/weclaw start
```

## How It Works

<p align="center">
  <img src="previews/architecture.png" width="600" />
</p>

**Agent modes:**

| Mode | How it works | Examples |
|------|-------------|----------|
| ACP  | Long-running subprocess, JSON-RPC over stdio. Fastest — reuses process and sessions. | Claude, Codex, Kimi, Gemini, Cursor, OpenCode, OpenClaw |
| CLI  | Spawns a new process per message. Supports session resume via `--resume`. | Claude (`claude -p`), Codex (`codex exec`) |
| HTTP | OpenAI-compatible chat completions API. | OpenClaw (HTTP fallback) |

Auto-detection picks ACP over CLI when both are available.

## Chat Commands

Send these as WeChat messages:

| Command | Description |
|---------|-------------|
| `hello` | Send to default agent |
| `/codex write a function` | Send to a specific agent |
| `/cc explain this code` | Send to agent by alias |
| `/claude` | Switch default agent to Claude |
| `/cwd /path/to/project` | Switch workspace directory |
| `/new` | Start a new conversation (clear session) |
| `/info` | Show current agent info |
| `/help` | Show help message |

### Aliases

| Alias | Agent |
|-------|-------|
| `/cc` | claude |
| `/cx` | codex |
| `/cs` | cursor |
| `/km` | kimi |
| `/gm` | gemini |
| `/ocd` | opencode |
| `/oc` | openclaw |

You can also define custom aliases per agent in config:

```json
{
  "agents": {
    "claude": {
      "type": "acp",
      "aliases": ["ai", "c"]
    }
  }
}
```

Then `/ai hello` or `/c hello` will route to claude.

Switching default agent is persisted to config — survives restarts.

## Media Messages

WeClaw supports sending images, videos, files, and voice messages to/from WeChat.

**Voice messages:** When you send a voice message in WeChat, WeClaw automatically uses WeChat's speech-to-text transcription and forwards the text to the AI agent. Duplicate voice message events are automatically deduplicated.

**From agent replies:** When an AI agent returns markdown with images (`![](url)`), WeClaw automatically extracts the image URLs, downloads them, uploads to WeChat CDN (AES-128-ECB encrypted), and sends them as image messages.

**Markdown handling:** Agent responses are automatically converted from markdown to plain text for WeChat display — code fences are stripped, links show display text only, bold/italic markers are removed, etc.

## Proactive Messaging

Send messages to WeChat users without waiting for them to message first.

**CLI:**

```bash
# Send text
weclaw send --to "user_id@im.wechat" --text "Hello from weclaw"

# Send image
weclaw send --to "user_id@im.wechat" --media "https://example.com/photo.png"

# Send text + image
weclaw send --to "user_id@im.wechat" --text "Check this out" --media "https://example.com/photo.png"

# Send file
weclaw send --to "user_id@im.wechat" --media "https://example.com/report.pdf"
```

**HTTP API** (runs on `127.0.0.1:18011` when `weclaw start` is running):

```bash
# Send text
curl -X POST http://127.0.0.1:18011/api/send \
  -H "Content-Type: application/json" \
  -d '{"to": "user_id@im.wechat", "text": "Hello from weclaw"}'

# Send image
curl -X POST http://127.0.0.1:18011/api/send \
  -H "Content-Type: application/json" \
  -d '{"to": "user_id@im.wechat", "media_url": "https://example.com/photo.png"}'

# Send text + media
curl -X POST http://127.0.0.1:18011/api/send \
  -H "Content-Type: application/json" \
  -d '{"to": "user_id@im.wechat", "text": "See this", "media_url": "https://example.com/photo.png"}'
```

Supported media types: images (png, jpg, gif, webp), videos (mp4, mov), files (pdf, doc, zip, etc.).

Set `WECLAW_API_ADDR` to change the listen address (e.g. `0.0.0.0:18011`).

## Configuration

Config file: `~/.weclaw/config.json`

```json
{
  "default_agent": "claude",
  "agents": {
    "claude": {
      "type": "acp",
      "command": "/usr/local/bin/claude-agent-acp",
      "env": {
        "ANTHROPIC_API_KEY": "sk-ant-xxx"
      },
      "model": "sonnet"
    },
    "codex": {
      "type": "acp",
      "command": "/usr/local/bin/codex-acp",
      "env": {
        "OPENAI_API_KEY": "sk-xxx"
      }
    },
    "openclaw": {
      "type": "http",
      "endpoint": "https://api.example.com/v1/chat/completions",
      "api_key": "sk-xxx",
      "model": "openclaw:main"
    }
  }
}
```

Environment variables:
- `WECLAW_DEFAULT_AGENT` — override default agent
- `OPENCLAW_GATEWAY_URL` — OpenClaw HTTP fallback endpoint
- `OPENCLAW_GATEWAY_TOKEN` — OpenClaw API token

Custom agent CLI environment variables:

```json
{
  "default_agent": "...",
  "agents": {
    "...": {
      ...
      "env": {
        "ENV_NAME": "ENV_VALUE"
      }
    },
  }
}
```

### Permission bypass

By default, some agents require interactive permission approval which doesn't work in WeChat. Add `args` to your agent config to bypass:

| Agent | Flag | What it does |
|-------|------|-------------|
| Claude (CLI) | `--dangerously-skip-permissions` | Skip all tool permission prompts |
| Codex (CLI) | `--skip-git-repo-check` | Allow running outside git repos |

Example:

```json
{
  "claude": {
    "type": "cli",
    "command": "/usr/local/bin/claude",
    "cwd": "/home/user/my-project",
    "args": ["--dangerously-skip-permissions"]
  },
  "codex": {
    "type": "cli",
    "command": "/usr/local/bin/codex",
    "cwd": "/home/user/my-project",
    "args": ["--skip-git-repo-check"]
  }
}
```

Set `cwd` to specify the agent's working directory (workspace). If omitted, defaults to `~/.weclaw/workspace`.

> **Warning:** These flags disable safety checks. Only enable them if you understand the risks. ACP agents handle permissions automatically and don't need these flags.

## Background Mode

```bash
# Start (runs in background by default)
weclaw start

# Check if running
weclaw status

# Stop
weclaw stop

# Run in foreground (for debugging)
weclaw start -f
```

Logs are written to `~/.weclaw/weclaw.log`.

### System service (auto-start on boot)

**macOS (launchd):**

```bash
cp service/com.fastclaw.weclaw.plist ~/Library/LaunchAgents/
launchctl load ~/Library/LaunchAgents/com.fastclaw.weclaw.plist
```

**Linux (systemd):**

```bash
sudo cp service/weclaw.service /etc/systemd/system/
sudo systemctl enable --now weclaw
```

## Docker

```bash
# Build
docker build -t weclaw .

# Login (interactive — scan QR code)
docker run -it -v ~/.weclaw:/root/.weclaw weclaw login

# Start with HTTP agent
docker run -d --name weclaw \
  -v ~/.weclaw:/root/.weclaw \
  -e OPENCLAW_GATEWAY_URL=https://api.example.com \
  -e OPENCLAW_GATEWAY_TOKEN=sk-xxx \
  weclaw

# View logs
docker logs -f weclaw
```

> Note: ACP and CLI agents require the agent binary inside the container.
> The Docker image ships only WeClaw itself. For ACP/CLI agents, mount
> the binary or build a custom image. HTTP agents work out of the box.

## Release

```bash
# Tag a new version to trigger GitHub Actions build & release
git tag v0.1.0
git push origin v0.1.0
```

The workflow builds binaries for `darwin/linux/windows` x `amd64/arm64`, creates a GitHub Release, and uploads all artifacts with checksums.

## Update

```bash
# Update to the latest version (auto-restarts if running)
weclaw update

# Check current version
weclaw version
```

## Development

```bash
# Hot reload
make dev

# Build
go build -o weclaw .

# Run
./weclaw start
```

## Contributors

<a href="https://github.com/fastclaw-ai/weclaw/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=fastclaw-ai/weclaw" />
</a>

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=fastclaw-ai/weclaw&type=Timeline)](https://star-history.com/#fastclaw-ai/weclaw&Timeline)

## License

[MIT](LICENSE)

```

[⬆ 回到目录](#toc)

## agent/acp_agent.go

```go
package agent

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// ACPAgent communicates with ACP-compatible agents (claude-agent-acp, codex-acp, cursor agent, etc.) via stdio JSON-RPC 2.0.
type ACPAgent struct {
	command      string
	args         []string
	model        string
	systemPrompt string
	cwd          string
	env          map[string]string
	protocol     string // "legacy_acp" or "codex_app_server"

	mu       sync.Mutex
	cmd      *exec.Cmd
	stdin    io.WriteCloser
	scanner  *bufio.Scanner
	started  bool
	nextID   atomic.Int64
	sessions map[string]string // conversationID -> sessionID (legacy ACP)
	threads  map[string]string // conversationID -> threadID (codex app-server)

	// pending tracks in-flight JSON-RPC requests
	pendingMu sync.Mutex
	pending   map[int64]chan *rpcResponse

	// notifications channel for session/update events
	notifyMu sync.Mutex
	notifyCh map[string]chan *sessionUpdate // sessionID -> channel
	turnCh   map[string]chan *codexTurnEvent

	stderr          *acpStderrWriter // captures stderr for error reporting
	progressCallback ProgressCallback // progress notification callback
}

// ACPAgentConfig holds configuration for the ACP agent.
type ACPAgentConfig struct {
	Command      string   // path to ACP agent binary (claude-agent-acp, codex-acp, cursor agent, etc.)
	Args         []string // extra args for command (e.g. ["acp"] for cursor)
	Model        string
	SystemPrompt string
	Cwd          string            // working directory
	Env          map[string]string // extra environment variables
}

// --- JSON-RPC types ---

type rpcRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      int64       `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type rpcResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      *int64          `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *rpcError       `json:"error,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// --- ACP protocol types ---

type initParams struct {
	ProtocolVersion    int                `json:"protocolVersion"`
	ClientCapabilities clientCapabilities `json:"clientCapabilities"`
}

type clientCapabilities struct {
	FS *fsCapabilities `json:"fs,omitempty"`
}

type fsCapabilities struct {
	ReadTextFile  bool `json:"readTextFile"`
	WriteTextFile bool `json:"writeTextFile"`
}

type newSessionParams struct {
	Cwd        string        `json:"cwd"`
	McpServers []interface{} `json:"mcpServers"`
}

type newSessionResult struct {
	SessionID string `json:"sessionId"`
}

type promptParams struct {
	SessionID string        `json:"sessionId"`
	Prompt    []promptEntry `json:"prompt"`
}

type promptEntry struct {
	Type     string `json:"type"`
	Text     string `json:"text,omitempty"`
	URL      string `json:"url,omitempty"`
	Path     string `json:"path,omitempty"`
	Data     string `json:"data,omitempty"`
	MimeType string `json:"mimeType,omitempty"`
}

type promptResult struct {
	StopReason string `json:"stopReason"`
}

type sessionUpdateParams struct {
	SessionID string        `json:"sessionId"`
	Update    sessionUpdate `json:"update"`
}

type sessionUpdate struct {
	SessionUpdate string          `json:"sessionUpdate"`
	Content       json.RawMessage `json:"content,omitempty"`
	// For agent_message_chunk
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

type permissionRequestParams struct {
	ToolCall json.RawMessage    `json:"toolCall"`
	Options  []permissionOption `json:"options"`
}

type permissionOption struct {
	OptionID string `json:"optionId"`
	Name     string `json:"name"`
	Kind     string `json:"kind"`
}

// Codex app-server protocol constants and types.
const (
	protocolLegacyACP      = "legacy_acp"
	protocolCodexAppServer = "codex_app_server"
)

type codexTurnStartParams struct {
	ThreadID       string           `json:"threadId"`
	ApprovalPolicy string           `json:"approvalPolicy,omitempty"`
	Input          []codexUserInput `json:"input"`
	SandboxPolicy  interface{}      `json:"sandboxPolicy,omitempty"`
	Model          string           `json:"model,omitempty"`
	Cwd            string           `json:"cwd,omitempty"`
}

type codexUserInput struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

type codexTurnEvent struct {
	Kind  string
	Delta string
	Text  string
}

func detectACPProtocol(command string, args []string) string {
	base := strings.ToLower(filepath.Base(command))
	// codex-acp is a standard ACP wrapper, NOT codex app-server
	// Only `codex app-server` uses the codex-native protocol
	if base == "codex" || base == "codex.exe" {
		for _, arg := range args {
			if arg == "app-server" {
				return protocolCodexAppServer
			}
		}
	}
	return protocolLegacyACP
}

// NewACPAgent creates a new ACP agent.
func NewACPAgent(cfg ACPAgentConfig) *ACPAgent {
	if cfg.Command == "" {
		cfg.Command = "claude-agent-acp"
	}
	if cfg.Cwd == "" {
		cfg.Cwd = defaultWorkspace()
	}
	protocol := detectACPProtocol(cfg.Command, cfg.Args)
	return &ACPAgent{
		command:      cfg.Command,
		args:         cfg.Args,
		model:        cfg.Model,
		systemPrompt: cfg.SystemPrompt,
		cwd:          cfg.Cwd,
		env:          cfg.Env,
		protocol:     protocol,
		sessions:     make(map[string]string),
		threads:      make(map[string]string),
		pending:      make(map[int64]chan *rpcResponse),
		notifyCh:     make(map[string]chan *sessionUpdate),
		turnCh:       make(map[string]chan *codexTurnEvent),
	}
}

// Start launches the claude-agent-acp subprocess and initializes the connection.
func (a *ACPAgent) Start(ctx context.Context) error {
	a.mu.Lock()
	if a.started {
		a.mu.Unlock()
		return nil
	}

	a.cmd = exec.CommandContext(ctx, a.command, a.args...)
	a.cmd.Dir = a.cwd
	if len(a.env) > 0 {
		cmdEnv, err := mergeEnv(os.Environ(), a.env)
		if err != nil {
			a.mu.Unlock()
			return fmt.Errorf("build acp env: %w", err)
		}
		a.cmd.Env = cmdEnv
	}
	// Capture stderr for debugging and error reporting
	a.stderr = &acpStderrWriter{prefix: "[acp-stderr]"}
	a.cmd.Stderr = a.stderr

	var err error
	a.stdin, err = a.cmd.StdinPipe()
	if err != nil {
		a.mu.Unlock()
		return fmt.Errorf("create stdin pipe: %w", err)
	}

	stdout, err := a.cmd.StdoutPipe()
	if err != nil {
		a.mu.Unlock()
		return fmt.Errorf("create stdout pipe: %w", err)
	}

	if err := a.cmd.Start(); err != nil {
		a.mu.Unlock()
		return fmt.Errorf("start acp agent %s: %w", a.command, err)
	}

	pid := a.cmd.Process.Pid
	log.Printf("[acp] started subprocess (command=%s, pid=%d)", a.command, pid)

	a.scanner = bufio.NewScanner(stdout)
	a.scanner.Buffer(make([]byte, 0, 4*1024*1024), 4*1024*1024) // 4MB
	a.started = true

	// Start reading loop
	go a.readLoop()

	// Release lock before calling initialize — call() needs a.mu to write to stdin
	a.mu.Unlock()

	// Initialize handshake with timeout
	initCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	log.Printf("[acp] sending initialize handshake (pid=%d, protocol=%s)...", pid, a.protocol)
	var result json.RawMessage
	if a.protocol == protocolCodexAppServer {
		result, err = a.call(initCtx, "initialize", map[string]interface{}{
			"clientInfo": map[string]string{"name": "weclaw", "version": "0.3.0"},
		})
		if err == nil {
			// codex app-server expects an "initialized" notification after initialize response
			err = a.notify("initialized", nil)
		}
	} else {
		result, err = a.call(initCtx, "initialize", initParams{
			ProtocolVersion: 1,
			ClientCapabilities: clientCapabilities{
				FS: &fsCapabilities{ReadTextFile: true, WriteTextFile: true},
			},
		})
	}
	if err != nil {
		a.mu.Lock()
		a.started = false
		a.mu.Unlock()
		a.stdin.Close()
		a.cmd.Process.Kill()
		a.cmd.Wait()
		// Use stderr detail if available (e.g. "connect ECONNREFUSED")
		if detail := a.stderr.LastError(); detail != "" {
			return fmt.Errorf("agent startup failed: %s", detail)
		}
		// Provide a helpful hint when the binary looks like a Claude CLI that doesn't support ACP
		base := strings.ToLower(filepath.Base(a.command))
		if base == "claude" || base == "claude.exe" {
			return fmt.Errorf("agent startup failed (pid=%d): %w\n\nHint: the 'claude' CLI does not support ACP protocol directly.\nSet type to \"cli\" in your config, or install claude-agent-acp and set command to \"claude-agent-acp\".", pid, err)
		}
		return fmt.Errorf("agent startup failed (pid=%d): %w", pid, err)
	}

	log.Printf("[acp] initialized (pid=%d): %s", pid, string(result))
	return nil
}

// Stop terminates the subprocess.
func (a *ACPAgent) Stop() {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.started {
		return
	}
	a.stdin.Close()
	a.cmd.Process.Kill()
	a.cmd.Wait()
	a.started = false
}

// SetCwd changes the working directory for subsequent sessions.
func (a *ACPAgent) SetCwd(cwd string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.cwd = cwd
}

// SetProgressCallback sets a callback for progress notifications.
func (a *ACPAgent) SetProgressCallback(callback ProgressCallback) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.progressCallback = callback
}

// sendProgress sends a progress event if a callback is registered.
func (a *ACPAgent) sendProgress(ctx context.Context, event ProgressEvent) {
	a.mu.Lock()
	callback := a.progressCallback
	a.mu.Unlock()

	if callback != nil {
		// Call callback in goroutine to avoid blocking
		go callback(ctx, event)
	}
}

// ResetSession clears the existing session for the given conversationID and
// immediately creates a new one, returning the new session ID.
func (a *ACPAgent) ResetSession(ctx context.Context, conversationID string) (string, error) {
	a.mu.Lock()
	delete(a.sessions, conversationID)
	a.mu.Unlock()
	log.Printf("[acp] session reset (conversation=%s), creating new session", conversationID)

	sessionID, _, err := a.getOrCreateSession(ctx, conversationID)
	if err != nil {
		return "", fmt.Errorf("create new session: %w", err)
	}
	return sessionID, nil
}

// Chat sends a message and returns the full response.
func (a *ACPAgent) Chat(ctx context.Context, conversationID string, message string) (string, error) {
	if !a.started {
		if err := a.Start(ctx); err != nil {
			return "", err
		}
	}

	// Route to codex app-server protocol if applicable
	if a.protocol == protocolCodexAppServer {
		return a.chatCodexAppServer(ctx, conversationID, message)
	}

	// Get or create session
	sessionID, isNew, err := a.getOrCreateSession(ctx, conversationID)
	if err != nil {
		return "", fmt.Errorf("session error: %w", err)
	}

	pid := a.cmd.Process.Pid
	if isNew {
		log.Printf("[acp] new session created (pid=%d, session=%s, conversation=%s)", pid, sessionID, conversationID)
	} else {
		log.Printf("[acp] reusing session (pid=%d, session=%s, conversation=%s)", pid, sessionID, conversationID)
	}

	// Register notification channel for this session
	notifyCh := make(chan *sessionUpdate, 256)
	a.notifyMu.Lock()
	a.notifyCh[sessionID] = notifyCh
	a.notifyMu.Unlock()

	defer func() {
		a.notifyMu.Lock()
		delete(a.notifyCh, sessionID)
		a.notifyMu.Unlock()
	}()

	// Send prompt (this blocks until the prompt completes)
	type promptDoneMsg struct {
		result json.RawMessage
		err    error
	}
	promptDone := make(chan promptDoneMsg, 1)
	go func() {
		result, err := a.call(ctx, "session/prompt", promptParams{
			SessionID: sessionID,
			Prompt:    []promptEntry{{Type: "text", Text: message}},
		})
		if result != nil {
			log.Printf("[acp] prompt result (session=%s): %s", sessionID, string(result))
		}
		promptDone <- promptDoneMsg{result: result, err: err}
	}()

	// Collect text chunks from notifications
	var textParts []string

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case update := <-notifyCh:
			if update.SessionUpdate == "agent_message_chunk" {
				text := extractChunkText(update)
				if text != "" {
					textParts = append(textParts, text)
				}
			}
		case done := <-promptDone:
			// Drain remaining notifications
			for {
				select {
				case update := <-notifyCh:
					if update.SessionUpdate == "agent_message_chunk" {
						text := extractChunkText(update)
						if text != "" {
							textParts = append(textParts, text)
						}
					}
				default:
					goto drained
				}
			}
		drained:
			if done.err != nil {
				return "", fmt.Errorf("prompt error: %w", done.err)
			}
			result := strings.TrimSpace(strings.Join(textParts, ""))
			if result == "" {
				// Try extracting from prompt result (some agents return content here)
				result = extractPromptResultText(done.result)
			}
			if result == "" {
				return "", fmt.Errorf("agent returned empty response")
			}
			return result, nil
		}
	}
}

// ChatWithMedia sends a message with media attachments and returns the full response.
func (a *ACPAgent) ChatWithMedia(ctx context.Context, conversationID string, message string, media []MediaEntry) (string, error) {
	if !a.started {
		if err := a.Start(ctx); err != nil {
			return "", err
		}
	}

	// Route to codex app-server protocol if applicable
	if a.protocol == protocolCodexAppServer {
		return a.chatCodexAppServerWithMedia(ctx, conversationID, message, media)
	}

	// Get or create session
	sessionID, isNew, err := a.getOrCreateSession(ctx, conversationID)
	if err != nil {
		return "", fmt.Errorf("session error: %w", err)
	}

	pid := a.cmd.Process.Pid
	if isNew {
		log.Printf("[acp] new session created (pid=%d, session=%s, conversation=%s)", pid, sessionID, conversationID)
	} else {
		log.Printf("[acp] reusing session (pid=%d, session=%s, conversation=%s)", pid, sessionID, conversationID)
	}

	// Register notification channel for this session
	notifyCh := make(chan *sessionUpdate, 256)
	a.notifyMu.Lock()
	a.notifyCh[sessionID] = notifyCh
	a.notifyMu.Unlock()

	defer func() {
		a.notifyMu.Lock()
		delete(a.notifyCh, sessionID)
		a.notifyMu.Unlock()
	}()

	// Build prompt entries with media
	prompt := buildPromptEntries(message, media)

	// Send prompt (this blocks until the prompt completes)
	type promptDoneMsg struct {
		result json.RawMessage
		err    error
	}
	promptDone := make(chan promptDoneMsg, 1)
	go func() {
		result, err := a.call(ctx, "session/prompt", promptParams{
			SessionID: sessionID,
			Prompt:    prompt,
		})
		if result != nil {
			log.Printf("[acp] prompt result (session=%s): %s", sessionID, string(result))
		}
		promptDone <- promptDoneMsg{result: result, err: err}
	}()

	// Collect text chunks from notifications
	var textParts []string

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case update := <-notifyCh:
			if update.SessionUpdate == "agent_message_chunk" {
				text := extractChunkText(update)
				if text != "" {
					textParts = append(textParts, text)
				}
			}
		case done := <-promptDone:
			// Drain remaining notifications
			for {
				select {
				case update := <-notifyCh:
					if update.SessionUpdate == "agent_message_chunk" {
						text := extractChunkText(update)
						if text != "" {
							textParts = append(textParts, text)
						}
					}
				default:
					goto drainedMedia
				}
			}
		drainedMedia:
			if done.err != nil {
				return "", fmt.Errorf("prompt error: %w", done.err)
			}
			result := strings.TrimSpace(strings.Join(textParts, ""))
			if result == "" {
				// Try extracting from prompt result (some agents return content here)
				result = extractPromptResultText(done.result)
			}
			if result == "" {
				return "", fmt.Errorf("agent returned empty response")
			}
			return result, nil
		}
	}
}

// buildPromptEntries builds prompt entries from message and media.
func buildPromptEntries(message string, media []MediaEntry) []promptEntry {
	var entries []promptEntry

	// Add media entries first
	for _, m := range media {
		entry := promptEntry{Type: m.Type}
		switch m.Type {
		case "image":
			if m.URL != "" {
				entry.URL = m.URL
			} else if m.Path != "" {
				entry.Path = m.Path
			}
		case "file":
			entry.Type = "file"
			if m.URL != "" {
				entry.URL = m.URL
			} else if m.Path != "" {
				entry.Path = m.Path
			}
			entry.MimeType = m.MIMEType
		case "video":
			entry.Type = "video"
			if m.URL != "" {
				entry.URL = m.URL
			} else if m.Path != "" {
				entry.Path = m.Path
			}
		}
		entries = append(entries, entry)
	}

	// Add text entry
	if message != "" {
		entries = append(entries, promptEntry{Type: "text", Text: message})
	}

	return entries
}

// chatCodexAppServerWithMedia handles media for codex app-server protocol.
func (a *ACPAgent) chatCodexAppServerWithMedia(ctx context.Context, conversationID string, message string, media []MediaEntry) (string, error) {
	threadID, isNew, err := a.getOrCreateThread(ctx, conversationID)
	if err != nil {
		return "", fmt.Errorf("thread error: %w", err)
	}

	pid := 0
	a.mu.Lock()
	if a.cmd != nil && a.cmd.Process != nil {
		pid = a.cmd.Process.Pid
	}
	a.mu.Unlock()

	if isNew {
		log.Printf("[acp] new thread created (pid=%d, thread=%s, conversation=%s)", pid, threadID, conversationID)
	} else {
		log.Printf("[acp] reusing thread (pid=%d, thread=%s, conversation=%s)", pid, threadID, conversationID)
	}

	// Register turn event channel
	turnCh := make(chan *codexTurnEvent, 256)
	a.notifyMu.Lock()
	a.turnCh[threadID] = turnCh
	a.notifyMu.Unlock()

	defer func() {
		a.notifyMu.Lock()
		delete(a.turnCh, threadID)
		a.notifyMu.Unlock()
	}()

	// Build input entries
	var input []codexUserInput
	for _, m := range media {
		input = append(input, codexUserInput{Type: m.Type, Text: m.URL})
	}
	if message != "" {
		input = append(input, codexUserInput{Type: "text", Text: message})
	}

	// Start turn (call returns quickly with turn info, actual content comes via events)
	go func() {
		_, err := a.call(ctx, "turn/start", codexTurnStartParams{
			ThreadID:       threadID,
			ApprovalPolicy: "never",
			Input:          input,
			SandboxPolicy:  map[string]interface{}{"type": "dangerFullAccess"},
			Model:          a.model,
			Cwd:            a.cwd,
		})
		if err != nil {
			// If call itself fails, signal via turn channel
			turnCh <- &codexTurnEvent{Kind: "error", Text: err.Error()}
		}
	}()

	// Collect text from events until turn/completed
	var textParts []string
	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case evt := <-turnCh:
			if evt.Kind == "error" {
				return "", fmt.Errorf("turn error: %s", evt.Text)
			}
			if evt.Delta != "" {
				textParts = append(textParts, evt.Delta)
			}
			if evt.Text != "" {
				textParts = append(textParts, evt.Text)
			}
			if evt.Kind == "completed" {
				result := strings.TrimSpace(strings.Join(textParts, ""))
				if result == "" {
					return "", fmt.Errorf("agent returned empty response")
				}
				return result, nil
			}
		}
	}
}

func (a *ACPAgent) getOrCreateSession(ctx context.Context, conversationID string) (string, bool, error) {
	a.mu.Lock()
	sid, exists := a.sessions[conversationID]
	a.mu.Unlock()

	if exists {
		return sid, false, nil
	}

	result, err := a.call(ctx, "session/new", newSessionParams{
		Cwd:        a.cwd,
		McpServers: []interface{}{},
	})
	if err != nil {
		return "", false, err
	}

	var sessionResult newSessionResult
	if err := json.Unmarshal(result, &sessionResult); err != nil {
		return "", false, fmt.Errorf("parse session result: %w", err)
	}

	a.mu.Lock()
	a.sessions[conversationID] = sessionResult.SessionID
	a.mu.Unlock()

	return sessionResult.SessionID, true, nil
}

// --- Codex app-server protocol ---

func (a *ACPAgent) getOrCreateThread(ctx context.Context, conversationID string) (string, bool, error) {
	a.mu.Lock()
	tid, exists := a.threads[conversationID]
	a.mu.Unlock()

	if exists {
		return tid, false, nil
	}

	params := map[string]interface{}{
		"approvalPolicy": "never",
		"cwd":            a.cwd,
		"sandbox":        "danger-full-access",
	}
	if a.model != "" {
		params["model"] = a.model
	}
	result, err := a.call(ctx, "thread/start", params)
	if err != nil {
		return "", false, err
	}

	var threadResult struct {
		Thread struct {
			ID string `json:"id"`
		} `json:"thread"`
	}
	if err := json.Unmarshal(result, &threadResult); err != nil {
		return "", false, fmt.Errorf("parse thread/start result: %w", err)
	}
	if threadResult.Thread.ID == "" {
		return "", false, fmt.Errorf("thread/start returned empty thread id")
	}

	a.mu.Lock()
	a.threads[conversationID] = threadResult.Thread.ID
	a.mu.Unlock()

	return threadResult.Thread.ID, true, nil
}

func (a *ACPAgent) chatCodexAppServer(ctx context.Context, conversationID string, message string) (string, error) {
	threadID, isNew, err := a.getOrCreateThread(ctx, conversationID)
	if err != nil {
		return "", fmt.Errorf("thread error: %w", err)
	}

	pid := 0
	a.mu.Lock()
	if a.cmd != nil && a.cmd.Process != nil {
		pid = a.cmd.Process.Pid
	}
	a.mu.Unlock()

	if isNew {
		log.Printf("[acp] new thread created (pid=%d, thread=%s, conversation=%s)", pid, threadID, conversationID)
	} else {
		log.Printf("[acp] reusing thread (pid=%d, thread=%s, conversation=%s)", pid, threadID, conversationID)
	}

	// Register turn event channel
	turnCh := make(chan *codexTurnEvent, 256)
	a.notifyMu.Lock()
	a.turnCh[threadID] = turnCh
	a.notifyMu.Unlock()

	defer func() {
		a.notifyMu.Lock()
		delete(a.turnCh, threadID)
		a.notifyMu.Unlock()
	}()

	// Start turn (call returns quickly with turn info, actual content comes via events)
	go func() {
		_, err := a.call(ctx, "turn/start", codexTurnStartParams{
			ThreadID:       threadID,
			ApprovalPolicy: "never",
			Input:          []codexUserInput{{Type: "text", Text: message}},
			SandboxPolicy:  map[string]interface{}{"type": "dangerFullAccess"},
			Model:          a.model,
			Cwd:            a.cwd,
		})
		if err != nil {
			// If call itself fails, signal via turn channel
			turnCh <- &codexTurnEvent{Kind: "error", Text: err.Error()}
		}
	}()

	// Collect text from events until turn/completed
	var textParts []string
	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case evt := <-turnCh:
			if evt.Kind == "error" {
				return "", fmt.Errorf("turn error: %s", evt.Text)
			}
			if evt.Delta != "" {
				textParts = append(textParts, evt.Delta)
			}
			if evt.Text != "" {
				textParts = append(textParts, evt.Text)
			}
			if evt.Kind == "completed" {
				result := strings.TrimSpace(strings.Join(textParts, ""))
				if result == "" {
					return "", fmt.Errorf("agent returned empty response")
				}
				return result, nil
			}
		}
	}
}

// notify sends a JSON-RPC notification (no id, no response expected).
func (a *ACPAgent) notify(method string, params interface{}) error {
	msg := struct {
		JSONRPC string      `json:"jsonrpc"`
		Method  string      `json:"method"`
		Params  interface{} `json:"params,omitempty"`
	}{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal notification: %w", err)
	}

	a.mu.Lock()
	_, err = fmt.Fprintf(a.stdin, "%s\n", data)
	a.mu.Unlock()
	return err
}

// call sends a JSON-RPC request and waits for the response.
func (a *ACPAgent) call(ctx context.Context, method string, params interface{}) (json.RawMessage, error) {
	id := a.nextID.Add(1)

	ch := make(chan *rpcResponse, 1)
	a.pendingMu.Lock()
	a.pending[id] = ch
	a.pendingMu.Unlock()

	defer func() {
		a.pendingMu.Lock()
		delete(a.pending, id)
		a.pendingMu.Unlock()
	}()

	req := rpcRequest{
		JSONRPC: "2.0",
		ID:      id,
		Method:  method,
		Params:  params,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	a.mu.Lock()
	_, err = fmt.Fprintf(a.stdin, "%s\n", data)
	a.mu.Unlock()
	if err != nil {
		return nil, fmt.Errorf("write to stdin: %w", err)
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case resp := <-ch:
		if resp.Error != nil {
			msg := resp.Error.Message
			// Enrich with stderr context if available
			if a.stderr != nil {
				if detail := a.stderr.LastError(); detail != "" {
					msg = detail
				}
			}
			return nil, fmt.Errorf("agent error: %s", msg)
		}
		return resp.Result, nil
	}
}

// readLoop reads NDJSON lines from stdout and dispatches to pending requests or notification channels.
func (a *ACPAgent) readLoop() {
	for a.scanner.Scan() {
		line := a.scanner.Text()
		if line == "" {
			continue
		}

		var msg rpcResponse
		if err := json.Unmarshal([]byte(line), &msg); err != nil {
			log.Printf("[acp] failed to parse message: %v", err)
			continue
		}

		// Response to a request we made (has id, no method)
		if msg.ID != nil && msg.Method == "" {
			a.pendingMu.Lock()
			ch, ok := a.pending[*msg.ID]
			a.pendingMu.Unlock()
			if ok {
				ch <- &msg
			}
			continue
		}

		// Request from agent or notification
		switch msg.Method {
		case "session/update":
			a.handleSessionUpdate(msg.Params)

		case "session/request_permission":
			// Auto-allow all permissions
			a.handlePermissionRequest(line)

		// Codex app-server events (multiple protocol versions)
		case "codex/event/agent_message_delta":
			a.handleCodexDelta(msg.Params)
		case "item/agentMessage/delta":
			a.handleCodexItemDelta(msg.Params)
		case "item/started":
			a.handleCodexItemStarted(msg.Params)
		case "turn/started", "turn/completed":
			a.handleCodexTurnEvent(msg.Method, msg.Params)
		case "codex/event/agent_message", "codex/event/task_complete",
			"codex/event/item_completed", "codex/event/token_count",
			"item/completed", "thread/tokenUsage/updated",
			"account/rateLimits/updated", "thread/status/changed":
			// Known events we don't need to act on
		case "turn/approval/request":
			a.handlePermissionRequest(line)

		default:
			if msg.Method != "" {
				log.Printf("[acp] unhandled method: %s (raw: %.200s)", msg.Method, line)
			}
		}
	}

	if err := a.scanner.Err(); err != nil {
		log.Printf("[acp] read loop error: %v", err)
	}
	log.Println("[acp] read loop ended")
}

func (a *ACPAgent) handleSessionUpdate(params json.RawMessage) {
	var p sessionUpdateParams
	if err := json.Unmarshal(params, &p); err != nil {
		log.Printf("[acp] failed to parse session/update: %v (raw: %s)", err, string(params))
		return
	}

	// Only log non-streaming events (skip chunks to reduce noise)
	switch p.Update.SessionUpdate {
	case "agent_message_chunk", "agent_thought_chunk":
		// skip — too noisy
	default:
		log.Printf("[acp] session/update (session=%s, type=%s)", p.SessionID, p.Update.SessionUpdate)
	}

	a.notifyMu.Lock()
	ch, ok := a.notifyCh[p.SessionID]
	a.notifyMu.Unlock()

	if ok {
		select {
		case ch <- &p.Update:
		default:
			log.Printf("[acp] notification channel full, dropping update (session=%s)", p.SessionID)
		}
	}
}

func (a *ACPAgent) handleCodexDelta(params json.RawMessage) {
	var p struct {
		Msg struct {
			Type  string `json:"type"`
			Delta string `json:"delta"`
		} `json:"msg"`
		ConversationID string `json:"conversationId"`
		ThreadID       string `json:"threadId"` // some versions use threadId
	}
	if err := json.Unmarshal(params, &p); err != nil {
		return
	}

	// Try conversationId first (codex uses this), fallback to threadId
	key := p.ConversationID
	if key == "" {
		key = p.ThreadID
	}

	delta := p.Msg.Delta
	if delta == "" {
		return
	}

	// Find the turn channel by thread ID — we need to match against stored threads
	a.notifyMu.Lock()
	ch, ok := a.turnCh[key]
	if !ok {
		// Try matching by iterating all turn channels (codex uses conversationId, not threadId)
		for _, c := range a.turnCh {
			ch = c
			ok = true
			break
		}
	}
	a.notifyMu.Unlock()

	if ok {
		select {
		case ch <- &codexTurnEvent{Delta: delta}:
		default:
		}
	}
}

// handleCodexItemDelta handles "item/agentMessage/delta" events.
// These contain incremental text deltas for the agent's response.
func (a *ACPAgent) handleCodexItemDelta(params json.RawMessage) {
	var p struct {
		ThreadID string `json:"threadId"`
		ItemID   string `json:"itemId"`
		Delta    string `json:"delta"`
	}
	if err := json.Unmarshal(params, &p); err != nil {
		return
	}

	if p.Delta == "" {
		return
	}

	a.dispatchToTurnCh(p.ThreadID, &codexTurnEvent{Delta: p.Delta})
}

// handleCodexItemStarted handles "item/started" events.
// When type=agentMessage, extracts text from content array.
func (a *ACPAgent) handleCodexItemStarted(params json.RawMessage) {
	var p struct {
		ThreadID string `json:"threadId"`
		Item     struct {
			Type    string `json:"type"`
			Content []struct {
				Type string `json:"type"`
				Text string `json:"text"`
			} `json:"content"`
		} `json:"item"`
	}
	if err := json.Unmarshal(params, &p); err != nil {
		return
	}

	// Send progress notification for non-agentMessage items
	if p.Item.Type != "agentMessage" {
		// Map item types to user-friendly messages
		var message string
		switch p.Item.Type {
		case "tool_use":
			message = "正在执行工具..."
		case "thinking":
			message = "正在思考..."
		default:
			message = fmt.Sprintf("正在处理: %s", p.Item.Type)
		}
		a.sendProgress(context.Background(), ProgressEvent{
			Type:    ProgressTypeProcessing,
			Message: message,
		})
		return
	}

	for _, c := range p.Item.Content {
		if c.Type == "text" && c.Text != "" {
			a.dispatchToTurnCh(p.ThreadID, &codexTurnEvent{Text: c.Text})
		}
	}
}

// handleCodexTurnEvent handles "turn/started" and "turn/completed" notifications.
func (a *ACPAgent) handleCodexTurnEvent(method string, params json.RawMessage) {
	var p struct {
		ThreadID string `json:"threadId"`
		Status   string `json:"status"`
	}
	if err := json.Unmarshal(params, &p); err != nil {
		return
	}

	if method == "turn/completed" {
		a.dispatchToTurnCh(p.ThreadID, &codexTurnEvent{Kind: "completed"})
	}
}

// dispatchToTurnCh sends an event to the turn channel for a thread.
func (a *ACPAgent) dispatchToTurnCh(threadID string, evt *codexTurnEvent) {
	a.notifyMu.Lock()
	ch, ok := a.turnCh[threadID]
	if !ok {
		// Fallback: try any active turn channel
		for _, c := range a.turnCh {
			ch = c
			ok = true
			break
		}
	}
	a.notifyMu.Unlock()

	if ok {
		select {
		case ch <- evt:
		default:
		}
	}
}

func (a *ACPAgent) handlePermissionRequest(raw string) {
	// Parse the request to get the ID and auto-allow
	var req struct {
		ID     json.RawMessage         `json:"id"`
		Params permissionRequestParams `json:"params"`
	}
	if err := json.Unmarshal([]byte(raw), &req); err != nil {
		log.Printf("[acp] failed to parse permission request: %v", err)
		return
	}

	// Extract tool name for progress notification
	var toolName string
	if req.Params.ToolCall != nil {
		var toolCall struct {
			Name string `json:"name"`
		}
		if err := json.Unmarshal(req.Params.ToolCall, &toolCall); err == nil && toolCall.Name != "" {
			toolName = toolCall.Name
			// Send progress notification
			a.sendProgress(context.Background(), ProgressEvent{
				Type:     ProgressTypeToolStart,
				Message:  fmt.Sprintf("正在调用工具: %s", toolName),
				ToolName: toolName,
			})
		}
	}

	// Find the "allow" option
	optionID := "allow"
	for _, opt := range req.Params.Options {
		if opt.Kind == "allow" {
			optionID = opt.OptionID
			break
		}
	}

	// Send response
	resp := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      req.ID,
		"result": map[string]interface{}{
			"outcome": map[string]interface{}{
				"outcome":  "selected",
				"optionId": optionID,
			},
		},
	}

	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("[acp] failed to marshal permission response: %v", err)
		return
	}

	a.mu.Lock()
	fmt.Fprintf(a.stdin, "%s\n", data)
	a.mu.Unlock()

	log.Printf("[acp] auto-allowed permission request (tool=%s)", toolName)
}

// Info returns metadata about this agent.
func (a *ACPAgent) Info() AgentInfo {
	info := AgentInfo{
		Name:    a.command,
		Type:    "acp",
		Model:   a.model,
		Command: a.command,
	}
	a.mu.Lock()
	if a.cmd != nil && a.cmd.Process != nil {
		info.PID = a.cmd.Process.Pid
	}
	a.mu.Unlock()
	return info
}

func extractChunkText(update *sessionUpdate) string {
	// The content field in agent_message_chunk can be a text content block
	if update.Text != "" {
		return update.Text
	}

	// Try to extract from content JSON
	if update.Content != nil {
		var content struct {
			Type string `json:"type"`
			Text string `json:"text"`
		}
		if err := json.Unmarshal(update.Content, &content); err == nil && content.Text != "" {
			return content.Text
		}
	}

	return ""
}

// extractPromptResultText tries to extract text from the session/prompt response.
// Some ACP agents include response content in the result alongside stopReason.
func extractPromptResultText(result json.RawMessage) string {
	if result == nil {
		return ""
	}

	// Try to extract content array from result
	var r struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
		// Some agents use a flat text field
		Text string `json:"text"`
	}
	if err := json.Unmarshal(result, &r); err != nil {
		return ""
	}

	if r.Text != "" {
		return r.Text
	}

	var parts []string
	for _, c := range r.Content {
		if c.Type == "text" && c.Text != "" {
			parts = append(parts, c.Text)
		}
	}
	return strings.Join(parts, "")
}

// acpStderrWriter forwards the ACP subprocess stderr to the application log
// and captures the last meaningful error line.
type acpStderrWriter struct {
	prefix string
	mu     sync.Mutex
	last   string // last non-empty, non-traceback line
}

func (w *acpStderrWriter) Write(p []byte) (int, error) {
	lines := strings.Split(strings.TrimRight(string(p), "\n"), "\n")
	w.mu.Lock()
	for _, line := range lines {
		if line != "" {
			log.Printf("%s %s", w.prefix, line)
			// Capture lines that look like actual error messages (not traceback frames)
			if !strings.HasPrefix(line, "  ") && !strings.HasPrefix(line, "Traceback") && !strings.HasPrefix(line, "...") {
				w.last = line
			}
		}
	}
	w.mu.Unlock()
	return len(p), nil
}

// LastError returns the last captured error line and resets it.
func (w *acpStderrWriter) LastError() string {
	w.mu.Lock()
	defer w.mu.Unlock()
	s := w.last
	w.last = ""
	return s
}

```

[⬆ 回到目录](#toc)

## agent/agent.go

```go
package agent

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// MediaEntry represents a media item (image, file, video) in a message.
type MediaEntry struct {
	Type     string // "image", "file", "video"
	URL      string // download URL (if available)
	Path     string // local file path (after download)
	MIMEType string // MIME type (if known)
	FileName string // original filename (for files)
}

// AgentInfo holds metadata about an agent for logging/debugging.
type AgentInfo struct {
	Name    string // e.g. "claude-acp", "claude", "gpt-4o"
	Type    string // e.g. "acp", "cli", "http"
	Model   string // e.g. "sonnet", "gpt-4o-mini"
	Command string // binary path, e.g. "/usr/local/bin/claude-agent-acp"
	PID     int    // subprocess PID (0 if not applicable, e.g. http agent)
}

// ProgressType represents the type of progress event.
type ProgressType string

const (
	ProgressTypeToolStart   ProgressType = "tool_start"   // Tool execution started
	ProgressTypeToolEnd     ProgressType = "tool_end"     // Tool execution ended
	ProgressTypeThought     ProgressType = "thought"      // Agent thinking/reasoning
	ProgressTypeFileRead    ProgressType = "file_read"    // Reading file
	ProgressTypeFileWrite   ProgressType = "file_write"   // Writing file
	ProgressTypeProcessing  ProgressType = "processing"   // General processing
	ProgressTypeSearching   ProgressType = "searching"    // Searching/analyzing
)

// ProgressEvent represents a progress notification from an agent.
type ProgressEvent struct {
	Type    ProgressType // Type of progress event
	Message string       // Human-readable progress message
	ToolName string      // Name of the tool being used (optional)
}

// ProgressCallback is called when an agent reports progress.
// The callback receives the context and the progress event.
type ProgressCallback func(ctx context.Context, event ProgressEvent)

// String returns a human-readable summary for logging.
func (i AgentInfo) String() string {
	s := fmt.Sprintf("name=%s, type=%s, model=%s, command=%s", i.Name, i.Type, i.Model, i.Command)
	if i.PID > 0 {
		s += fmt.Sprintf(", pid=%d", i.PID)
	}
	return s
}

// defaultWorkspace returns ~/.weclaw/workspace as the default working directory.
func defaultWorkspace() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return os.TempDir()
	}
	dir := filepath.Join(home, ".weclaw", "workspace")
	os.MkdirAll(dir, 0o755)
	return dir
}

// mergeEnv merges extra environment variables into the base environment.
func mergeEnv(base []string, extra map[string]string) ([]string, error) {
	if len(extra) == 0 {
		return base, nil
	}

	merged := append([]string(nil), base...)
	indexByKey := make(map[string]int, len(base))
	for i, entry := range merged {
		key, _, found := strings.Cut(entry, "=")
		if !found || key == "" {
			continue
		}
		indexByKey[key] = i
	}

	newKeys := make([]string, 0, len(extra))
	for key, value := range extra {
		if key == "" || strings.Contains(key, "=") {
			return nil, fmt.Errorf("invalid env key %q", key)
		}
		entry := key + "=" + value
		if idx, ok := indexByKey[key]; ok {
			merged[idx] = entry
			continue
		}
		newKeys = append(newKeys, key)
	}

	sort.Strings(newKeys)
	for _, key := range newKeys {
		merged = append(merged, key+"="+extra[key])
	}

	return merged, nil
}

// Agent is the interface for AI chat agents.
type Agent interface {
	// Chat sends a message to the agent and returns the response.
	// conversationID is used to maintain conversation history per user.
	Chat(ctx context.Context, conversationID string, message string) (string, error)

	// ChatWithMedia sends a message with media attachments to the agent.
	// media can contain images, files, videos, etc.
	ChatWithMedia(ctx context.Context, conversationID string, message string, media []MediaEntry) (string, error)

	// ResetSession clears the existing session for the given conversationID and
	// starts a new one. Returns the new session ID if immediately available
	// (ACP mode), or an empty string if the ID will be assigned on next Chat
	// (CLI mode) or is not applicable (HTTP mode).
	ResetSession(ctx context.Context, conversationID string) (string, error)

	// Info returns metadata about this agent.
	Info() AgentInfo

	// SetCwd changes the working directory for subsequent operations.
	SetCwd(cwd string)

	// SetProgressCallback sets a callback for progress notifications.
	// The callback will be invoked when the agent reports progress during long-running operations.
	SetProgressCallback(callback ProgressCallback)
}

```

[⬆ 回到目录](#toc)

## agent/cli_agent.go

```go
package agent

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

// CLIAgent invokes a local CLI agent (claude, codex, etc.) via streaming JSON.
type CLIAgent struct {
	name         string
	command      string
	args         []string          // extra args from config
	cwd          string            // working directory
	env          map[string]string // extra environment variables
	model        string
	systemPrompt string
	mu           sync.Mutex
	sessions     map[string]string // conversationID -> session ID for multi-turn
}

// CLIAgentConfig holds configuration for a CLI agent.
type CLIAgentConfig struct {
	Name         string            // agent name for logging, e.g. "claude", "codex"
	Command      string            // path to binary
	Args         []string          // extra args (e.g. ["--dangerously-skip-permissions"])
	Cwd          string            // working directory (workspace)
	Env          map[string]string // extra environment variables
	Model        string
	SystemPrompt string
}

// NewCLIAgent creates a new CLI agent.
func NewCLIAgent(cfg CLIAgentConfig) *CLIAgent {
	cwd := cfg.Cwd
	if cwd == "" {
		cwd = defaultWorkspace()
	}
	return &CLIAgent{
		name:         cfg.Name,
		command:      cfg.Command,
		args:         cfg.Args,
		cwd:          cwd,
		env:          cfg.Env,
		model:        cfg.Model,
		systemPrompt: cfg.SystemPrompt,
		sessions:     make(map[string]string),
	}
}

// streamEvent represents a single event from claude's stream-json output.
type streamEvent struct {
	Type      string         `json:"type"`
	SessionID string         `json:"session_id"`
	Result    string         `json:"result"`
	IsError   bool           `json:"is_error"`
	Message   *streamMessage `json:"message,omitempty"`
}

// streamMessage represents the message field in an assistant event.
type streamMessage struct {
	Content []streamContent `json:"content"`
}

// streamContent represents a content block in an assistant message.
type streamContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// Info returns metadata about this agent.
func (a *CLIAgent) Info() AgentInfo {
	return AgentInfo{
		Name:    a.name,
		Type:    "cli",
		Model:   a.model,
		Command: a.command,
	}
}

// ResetSession clears the existing session for the given conversationID.
// Returns an empty string because the new session ID is only known after the
// next Chat call (claude assigns it during the conversation).
func (a *CLIAgent) ResetSession(_ context.Context, conversationID string) (string, error) {
	a.mu.Lock()
	delete(a.sessions, conversationID)
	a.mu.Unlock()
	log.Printf("[cli] session reset (command=%s, conversation=%s)", a.command, conversationID)
	return "", nil
}

// SetCwd changes the working directory for subsequent CLI invocations.
func (a *CLIAgent) SetCwd(cwd string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.cwd = cwd
}

// SetProgressCallback sets a callback for progress notifications.
// CLI agent doesn't support progress notifications, so this is a no-op.
func (a *CLIAgent) SetProgressCallback(callback ProgressCallback) {
	// CLI agent runs in separate processes, can't report progress
}

// Chat sends a message to the CLI agent and returns the response.
func (a *CLIAgent) Chat(ctx context.Context, conversationID string, message string) (string, error) {
	switch a.name {
	case "codex":
		return a.chatCodex(ctx, message)
	default:
		return a.chatClaude(ctx, conversationID, message)
	}
}

// ChatWithMedia sends a message with media attachments.
// CLI agents currently don't support media natively, so we add media info to the message.
func (a *CLIAgent) ChatWithMedia(ctx context.Context, conversationID string, message string, media []MediaEntry) (string, error) {
	// Build enhanced message with media descriptions
	enhancedMessage := message
	for _, m := range media {
		switch m.Type {
		case "image":
			enhancedMessage += fmt.Sprintf("\n[图片: %s]", m.URL)
		case "file":
			enhancedMessage += fmt.Sprintf("\n[文件: %s]", m.FileName)
		case "video":
			enhancedMessage += fmt.Sprintf("\n[视频: %s]", m.URL)
		}
	}
	return a.Chat(ctx, conversationID, enhancedMessage)
}

// chatClaude uses claude -p with stream-json to get structured output and session persistence.
func (a *CLIAgent) chatClaude(ctx context.Context, conversationID string, message string) (string, error) {
	args := []string{"-p", message, "--output-format", "stream-json", "--verbose"}

	if a.model != "" {
		args = append(args, "--model", a.model)
	}
	if a.systemPrompt != "" {
		args = append(args, "--append-system-prompt", a.systemPrompt)
	}
	// Append extra args from config (e.g. --dangerously-skip-permissions)
	args = append(args, a.args...)

	// Resume existing session for multi-turn conversation
	a.mu.Lock()
	sessionID, hasSession := a.sessions[conversationID]
	a.mu.Unlock()

	if hasSession {
		args = append(args, "--resume", sessionID)
		log.Printf("[cli] resuming session (command=%s, session=%s, conversation=%s)", a.command, sessionID, conversationID)
	} else {
		log.Printf("[cli] starting new conversation (command=%s, conversation=%s)", a.command, conversationID)
	}

	cmd := exec.CommandContext(ctx, a.command, args...)
	if a.cwd != "" {
		cmd.Dir = a.cwd
	}
	if len(a.env) > 0 {
		cmdEnv, err := mergeEnv(os.Environ(), a.env)
		if err != nil {
			return "", fmt.Errorf("build %s env: %w", a.name, err)
		}
		cmd.Env = cmdEnv
	}
	var stderr strings.Builder
	cmd.Stderr = &stderr

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("create stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("start %s: %w", a.name, err)
	}

	log.Printf("[cli] spawned process (command=%s, pid=%d, conversation=%s)", a.command, cmd.Process.Pid, conversationID)

	// Parse streaming JSON events
	var result string
	var newSessionID string
	var assistantTexts []string

	scanner := bufio.NewScanner(stdout)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024) // 1MB buffer for large responses

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var event streamEvent
		if err := json.Unmarshal([]byte(line), &event); err != nil {
			continue
		}

		// Capture session ID from any event
		if event.SessionID != "" {
			newSessionID = event.SessionID
		}

		switch event.Type {
		case "result":
			if event.IsError {
				return "", fmt.Errorf("%s returned error: %s", a.name, event.Result)
			}
			result = event.Result
		case "assistant":
			// Newer claude CLI versions send text in assistant events
			// instead of the result event's result field.
			if event.Message != nil {
				for _, c := range event.Message.Content {
					if c.Type == "text" && c.Text != "" {
						assistantTexts = append(assistantTexts, c.Text)
					}
				}
			}
		}
	}

	// If the result event had an empty result, fall back to accumulated assistant texts.
	if result == "" && len(assistantTexts) > 0 {
		result = strings.Join(assistantTexts, "")
	}

	if err := cmd.Wait(); err != nil {
		if result == "" {
			errMsg := strings.TrimSpace(stderr.String())
			if errMsg != "" {
				return "", fmt.Errorf("%s exited with error: %w, stderr: %s", a.name, err, errMsg)
			}
			return "", fmt.Errorf("%s exited with error: %w", a.name, err)
		}
		// If we got a result but exit code is non-zero (e.g. hook failures), still return the result
	}

	log.Printf("[cli] process exited (command=%s, pid=%d)", a.command, cmd.Process.Pid)

	// Save session ID for multi-turn conversation
	if newSessionID != "" {
		a.mu.Lock()
		a.sessions[conversationID] = newSessionID
		a.mu.Unlock()
		log.Printf("[cli] saved session (session=%s, conversation=%s)", newSessionID, conversationID)
	}

	result = strings.TrimSpace(result)
	if result == "" {
		return "", fmt.Errorf("%s returned empty response", a.name)
	}

	return result, nil
}

// chatCodex handles codex CLI invocation using "codex exec".
func (a *CLIAgent) chatCodex(ctx context.Context, message string) (string, error) {
	args := []string{"exec", message}
	if a.model != "" {
		args = append(args, "--model", a.model)
	}
	// Append extra args from config (e.g. --skip-git-repo-check)
	args = append(args, a.args...)

	log.Printf("[cli] running codex exec (command=%s)", a.command)
	cmd := exec.CommandContext(ctx, a.command, args...)
	if a.cwd != "" {
		cmd.Dir = a.cwd
	}
	if len(a.env) > 0 {
		cmdEnv, err := mergeEnv(os.Environ(), a.env)
		if err != nil {
			return "", fmt.Errorf("build %s env: %w", a.name, err)
		}
		cmd.Env = cmdEnv
	}
	var stderr strings.Builder
	cmd.Stderr = &stderr

	out, err := cmd.Output()
	if err != nil {
		errMsg := strings.TrimSpace(stderr.String())
		if errMsg != "" {
			return "", fmt.Errorf("codex error: %w, stderr: %s", err, errMsg)
		}
		return "", fmt.Errorf("codex error: %w", err)
	}

	result := strings.TrimSpace(string(out))
	if result == "" {
		return "", fmt.Errorf("codex returned empty response")
	}
	return result, nil
}

```

[⬆ 回到目录](#toc)

## agent/env_test.go

```go
package agent

import (
	"reflect"
	"testing"
)

func TestMergeEnvOverridesAndAppends(t *testing.T) {
	base := []string{"PATH=/usr/bin", "KEEP=1", "DUP=old"}
	extra := map[string]string{
		"NEW":   "2",
		"DUP":   "new",
		"EMPTY": "",
	}

	got, err := mergeEnv(base, extra)
	if err != nil {
		t.Fatalf("mergeEnv returned error: %v", err)
	}

	want := []string{"PATH=/usr/bin", "KEEP=1", "DUP=new", "EMPTY=", "NEW=2"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("mergeEnv() = %#v, want %#v", got, want)
	}
}

func TestMergeEnvRejectsInvalidKey(t *testing.T) {
	_, err := mergeEnv(nil, map[string]string{"BAD=KEY": "value"})
	if err == nil {
		t.Fatal("mergeEnv() error = nil, want invalid env key error")
	}
}

func TestMergeEnvPreservesBaseWhenNoExtra(t *testing.T) {
	base := []string{"PATH=/usr/bin", "KEEP=1"}

	got, err := mergeEnv(base, nil)
	if err != nil {
		t.Fatalf("mergeEnv returned error: %v", err)
	}
	if !reflect.DeepEqual(got, base) {
		t.Fatalf("mergeEnv() = %#v, want %#v", got, base)
	}
}

func TestMergeEnvRejectsEmptyKey(t *testing.T) {
	_, err := mergeEnv(nil, map[string]string{"": "value"})
	if err == nil {
		t.Fatal("mergeEnv() error = nil, want empty env key error")
	}
}

func TestMergeEnvOverridesExistingKeyWithEmptyValue(t *testing.T) {
	got, err := mergeEnv([]string{"EMPTY=old"}, map[string]string{"EMPTY": ""})
	if err != nil {
		t.Fatalf("mergeEnv returned error: %v", err)
	}
	want := []string{"EMPTY="}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("mergeEnv() = %#v, want %#v", got, want)
	}
}

```

[⬆ 回到目录](#toc)

## agent/http_agent.go

```go
package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// ChatMessage represents a single message in a conversation.
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// HTTPAgent is an OpenAI-compatible chat completions API client.
type HTTPAgent struct {
	endpoint     string
	apiKey       string
	headers      map[string]string
	model        string
	systemPrompt string
	httpClient   *http.Client
	mu           sync.Mutex
	history      map[string][]ChatMessage // conversationID -> messages
	maxHistory   int
}

// HTTPAgentConfig holds configuration for the HTTP agent.
type HTTPAgentConfig struct {
	Endpoint     string
	APIKey       string
	Headers      map[string]string
	Model        string
	SystemPrompt string
	MaxHistory   int
}

// NewHTTPAgent creates a new OpenAI-compatible HTTP agent.
func NewHTTPAgent(cfg HTTPAgentConfig) *HTTPAgent {
	if cfg.MaxHistory == 0 {
		cfg.MaxHistory = 20
	}
	if cfg.Model == "" {
		cfg.Model = "gpt-4o-mini"
	}
	return &HTTPAgent{
		endpoint:     cfg.Endpoint,
		apiKey:       cfg.APIKey,
		headers:      cfg.Headers,
		model:        cfg.Model,
		systemPrompt: cfg.SystemPrompt,
		httpClient:   &http.Client{Timeout: 120 * time.Second},
		history:      make(map[string][]ChatMessage),
		maxHistory:   cfg.MaxHistory,
	}
}

// Info returns metadata about this agent.
func (a *HTTPAgent) Info() AgentInfo {
	return AgentInfo{
		Name:    "http",
		Type:    "http",
		Model:   a.model,
		Command: a.endpoint,
	}
}

// SetCwd is a no-op for HTTP agents (they have no working directory).
func (a *HTTPAgent) SetCwd(_ string) {}

// SetProgressCallback sets a callback for progress notifications.
// HTTP agents don't support progress notifications, so this is a no-op.
func (a *HTTPAgent) SetProgressCallback(callback ProgressCallback) {
	// HTTP agents use standard OpenAI API with no progress reporting
}

// ResetSession clears the conversation history for the given conversationID.
// HTTP agents have no server-side session ID, so an empty string is returned.
func (a *HTTPAgent) ResetSession(_ context.Context, conversationID string) (string, error) {
	a.mu.Lock()
	delete(a.history, conversationID)
	a.mu.Unlock()
	return "", nil
}

// Chat sends a message to the OpenAI-compatible API and returns the response.
func (a *HTTPAgent) Chat(ctx context.Context, conversationID string, message string) (string, error) {
	a.mu.Lock()
	messages := a.buildMessages(conversationID, message)
	a.mu.Unlock()

	reqBody := map[string]interface{}{
		"model":    a.model,
		"messages": messages,
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, a.endpoint, bytes.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if a.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+a.apiKey)
	}
	for k, v := range a.headers {
		req.Header.Set(k, v)
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error HTTP %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("parse response: %w", err)
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}

	reply := result.Choices[0].Message.Content

	// Save to history
	a.mu.Lock()
	a.history[conversationID] = append(a.history[conversationID],
		ChatMessage{Role: "user", Content: message},
		ChatMessage{Role: "assistant", Content: reply},
	)
	// Trim history
	if len(a.history[conversationID]) > a.maxHistory*2 {
		a.history[conversationID] = a.history[conversationID][len(a.history[conversationID])-a.maxHistory*2:]
	}
	a.mu.Unlock()

	return reply, nil
}

// ChatWithMedia sends a message with media attachments.
// For HTTP agents, media is converted to text description (limited support).
func (a *HTTPAgent) ChatWithMedia(ctx context.Context, conversationID string, message string, media []MediaEntry) (string, error) {
	// Build enhanced message with media descriptions
	enhancedMessage := message
	for _, m := range media {
		switch m.Type {
		case "image":
			enhancedMessage += fmt.Sprintf("\n[图片: %s]", m.URL)
		case "file":
			enhancedMessage += fmt.Sprintf("\n[文件: %s (%s)]", m.FileName, m.URL)
		case "video":
			enhancedMessage += fmt.Sprintf("\n[视频: %s]", m.URL)
		}
	}
	return a.Chat(ctx, conversationID, enhancedMessage)
}

func (a *HTTPAgent) buildMessages(conversationID string, message string) []ChatMessage {
	var messages []ChatMessage
	if a.systemPrompt != "" {
		messages = append(messages, ChatMessage{Role: "system", Content: a.systemPrompt})
	}
	if hist, ok := a.history[conversationID]; ok {
		messages = append(messages, hist...)
	}
	messages = append(messages, ChatMessage{Role: "user", Content: message})
	return messages
}

```

[⬆ 回到目录](#toc)

## api/server.go

```go
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fastclaw-ai/weclaw/ilink"
	"github.com/fastclaw-ai/weclaw/messaging"
)

// Server provides an HTTP API for sending messages.
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
	mux.HandleFunc("/api/send", s.handleSend)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

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

```

[⬆ 回到目录](#toc)

## cmd/login.go

```go
package cmd

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Add a WeChat account via QR code scan",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()

		creds, err := doLogin(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Account %s added. Run 'weclaw start' to begin.\n", creds.ILinkBotID)
		return nil
	},
}

```

[⬆ 回到目录](#toc)

## cmd/proc_unix.go

```go
//go:build !windows

package cmd

import (
	"os/exec"
	"syscall"
)

func setSysProcAttr(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
}

```

[⬆ 回到目录](#toc)

## cmd/proc_windows.go

```go
//go:build windows

package cmd

import "os/exec"

func setSysProcAttr(_ *exec.Cmd) {
	// No Setsid on Windows — process is already detached via Start()
}

```

[⬆ 回到目录](#toc)

## cmd/restart.go

```go
package cmd

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(restartCmd)
}

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the background weclaw process",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Stop if running
		pid, err := readPid()
		if err == nil && processExists(pid) {
			fmt.Printf("Stopping weclaw (pid=%d)...\n", pid)
			if p, err := os.FindProcess(pid); err == nil {
				p.Signal(syscall.SIGTERM)
			}
			for i := 0; i < 20; i++ {
				if !processExists(pid) {
					break
				}
				time.Sleep(500 * time.Millisecond)
			}
			os.Remove(pidFile())
		}

		// Start
		fmt.Println("Starting weclaw...")
		return runDaemon()
	},
}

```

[⬆ 回到目录](#toc)

## cmd/root.go

```go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Version is set at build time via -ldflags.
var Version = "dev"

var rootCmd = &cobra.Command{
	Use:     "weclaw",
	Short:   "WeChat AI agent bridge",
	Long:    "weclaw bridges WeChat messages to AI agents via the iLink API.",
	Version: Version,
	RunE:    runStart, // default command is start
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

```

[⬆ 回到目录](#toc)

## cmd/send.go

```go
package cmd

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/fastclaw-ai/weclaw/ilink"
	"github.com/fastclaw-ai/weclaw/messaging"
	"github.com/spf13/cobra"
)

var (
	sendTo       string
	sendText     string
	sendMediaURL string
)

func init() {
	sendCmd.Flags().StringVar(&sendTo, "to", "", "Target user ID (ilink user ID)")
	sendCmd.Flags().StringVar(&sendText, "text", "", "Message text to send")
	sendCmd.Flags().StringVar(&sendMediaURL, "media", "", "Media URL to send (image/video/file)")
	sendCmd.MarkFlagRequired("to")
	rootCmd.AddCommand(sendCmd)
}

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message to a WeChat user",
	Example: `  weclaw send --to "user_id@im.wechat" --text "Hello"
  weclaw send --to "user_id@im.wechat" --media "https://example.com/image.png"
  weclaw send --to "user_id@im.wechat" --text "See this" --media "https://example.com/image.png"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if sendText == "" && sendMediaURL == "" {
			return fmt.Errorf("at least one of --text or --media is required")
		}

		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()

		accounts, err := ilink.LoadAllCredentials()
		if err != nil {
			return fmt.Errorf("load credentials: %w", err)
		}
		if len(accounts) == 0 {
			return fmt.Errorf("no accounts found, run 'weclaw start' first")
		}

		client := ilink.NewClient(accounts[0])

		if sendText != "" {
			if err := messaging.SendTextReply(ctx, client, sendTo, sendText, "", ""); err != nil {
				return fmt.Errorf("send text failed: %w", err)
			}
			fmt.Println("Text sent")
		}

		if sendMediaURL != "" {
			if err := messaging.SendMediaFromURL(ctx, client, sendTo, sendMediaURL, ""); err != nil {
				return fmt.Errorf("send media failed: %w", err)
			}
			fmt.Println("Media sent")
		}

		return nil
	},
}

```

[⬆ 回到目录](#toc)

## cmd/start.go

```go
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/fastclaw-ai/weclaw/agent"
	"github.com/fastclaw-ai/weclaw/api"
	"github.com/fastclaw-ai/weclaw/config"
	"github.com/fastclaw-ai/weclaw/ilink"
	"github.com/fastclaw-ai/weclaw/messaging"
	"github.com/mdp/qrterminal/v3"
	"github.com/spf13/cobra"
)

var (
	foregroundFlag bool
	apiAddrFlag    string
)

func init() {
	startCmd.Flags().BoolVarP(&foregroundFlag, "foreground", "f", false, "Run in foreground (default is background)")
	startCmd.Flags().StringVar(&apiAddrFlag, "api-addr", "", "API server listen address (default 127.0.0.1:18011)")
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the WeChat message bridge (auto-login if needed)",
	RunE:  runStart,
}

func runStart(cmd *cobra.Command, args []string) error {
	if !foregroundFlag {
		// Check if login is needed — if so, do it in foreground first, then daemon
		accounts, _ := ilink.LoadAllCredentials()
		if len(accounts) == 0 {
			fmt.Println("No WeChat accounts found, starting login...")
			ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
			_, err := doLogin(ctx)
			cancel()
			if err != nil {
				return fmt.Errorf("login failed: %w", err)
			}
		}
		return runDaemon()
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// Load all accounts
	accounts, err := ilink.LoadAllCredentials()
	if err != nil {
		return fmt.Errorf("failed to load credentials: %w", err)
	}

	// No accounts — trigger login
	if len(accounts) == 0 {
		log.Println("No WeChat accounts found, starting login...")
		creds, err := doLogin(ctx)
		if err != nil {
			return fmt.Errorf("login failed: %w", err)
		}
		accounts = append(accounts, creds)
	}

	// Load config and auto-detect agents
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if config.DetectAndConfigure(cfg) {
		if err := config.Save(cfg); err != nil {
			log.Printf("Warning: failed to save auto-detected config: %v", err)
		} else {
			path, _ := config.ConfigPath()
			log.Printf("Auto-detected agents saved to %s", path)
		}
	}

	// Log all available agents
	if len(cfg.Agents) > 0 {
		names := make([]string, 0, len(cfg.Agents))
		for name := range cfg.Agents {
			names = append(names, name)
		}
		log.Printf("Available agents: %v (default: %s)", names, cfg.DefaultAgent)
	}

	// Create handler with an agent factory for on-demand agent creation
	handler := messaging.NewHandler(
		func(ctx context.Context, name string) agent.Agent {
			return createAgentByName(ctx, cfg, name)
		},
		func(name string) error {
			cfg.DefaultAgent = name
			return config.Save(cfg)
		},
	)

	// Populate agent metas for /status
	var metas []messaging.AgentMeta
	workDirs := make(map[string]string, len(cfg.Agents))
	for name, agCfg := range cfg.Agents {
		command := agCfg.Command
		if agCfg.Type == "http" {
			command = agCfg.Endpoint
		}
		metas = append(metas, messaging.AgentMeta{
			Name:    name,
			Type:    agCfg.Type,
			Command: command,
			Model:   agCfg.Model,
		})
		if agCfg.Cwd != "" {
			workDirs[name] = agCfg.Cwd
		}
	}
	handler.SetAgentMetas(metas)
	handler.SetAgentWorkDirs(workDirs)

	// Load custom aliases from agent configs
	handler.SetCustomAliases(config.BuildAliasMap(cfg.Agents))

	// Set save directory for images/files if configured
	if cfg.SaveDir != "" {
		handler.SetSaveDir(cfg.SaveDir)
		log.Printf("Image save directory: %s", cfg.SaveDir)
	}

	// Start default agent initialization in background so monitors can start immediately
	go func() {
		if cfg.DefaultAgent == "" {
			log.Println("No default agent configured, staying in echo mode")
			return
		}
		log.Printf("Initializing default agent %q in background...", cfg.DefaultAgent)
		ag := createAgentByName(ctx, cfg, cfg.DefaultAgent)
		if ag == nil {
			log.Printf("Failed to initialize default agent %q, staying in echo mode", cfg.DefaultAgent)
		} else {
			handler.SetDefaultAgent(cfg.DefaultAgent, ag)
		}
	}()

	// Start HTTP API server for sending messages
	var clients []*ilink.Client
	for _, c := range accounts {
		clients = append(clients, ilink.NewClient(c))
	}
	// Resolve API addr: flag > env/config > default
	apiAddr := cfg.APIAddr // already includes env override from loadEnv
	if apiAddrFlag != "" {
		apiAddr = apiAddrFlag
	}
	apiServer := api.NewServer(clients, apiAddr)
	go func() {
		if err := apiServer.Run(ctx); err != nil {
			log.Printf("API server error: %v", err)
		}
	}()

	// Start monitors immediately — they will use echo mode until agent is ready
	log.Printf("Starting message bridge for %d account(s)...", len(accounts))

	var wg sync.WaitGroup
	for _, creds := range accounts {
		wg.Add(1)
		go func(c *ilink.Credentials) {
			defer wg.Done()
			runMonitorWithRestart(ctx, c, handler)
		}(creds)
	}

	wg.Wait()
	log.Println("All monitors stopped")
	return nil
}

// runMonitorWithRestart runs a monitor with automatic restart on failure.
func runMonitorWithRestart(ctx context.Context, creds *ilink.Credentials, handler *messaging.Handler) {
	const maxRestartDelay = 30 * time.Second
	restartDelay := 3 * time.Second

	for {
		log.Printf("[%s] Starting monitor...", creds.ILinkBotID)

		client := ilink.NewClient(creds)
		monitor, err := ilink.NewMonitor(client, handler.HandleMessage)
		if err != nil {
			log.Printf("[%s] Failed to create monitor: %v", creds.ILinkBotID, err)
		} else {
			err = monitor.Run(ctx)
		}

		// If context is cancelled, exit
		if ctx.Err() != nil {
			return
		}

		log.Printf("[%s] Monitor stopped: %v, restarting in %s", creds.ILinkBotID, err, restartDelay)
		select {
		case <-time.After(restartDelay):
		case <-ctx.Done():
			return
		}

		// Exponential backoff for restarts, capped
		restartDelay *= 2
		if restartDelay > maxRestartDelay {
			restartDelay = maxRestartDelay
		}
	}
}

// createAgentByName creates and starts an agent by its config name.
// Returns nil if the agent is not configured or fails to start.
func createAgentByName(ctx context.Context, cfg *config.Config, name string) agent.Agent {
	agCfg, ok := cfg.Agents[name]
	if !ok {
		log.Printf("[agent] %q not found in config", name)
		return nil
	}

	switch agCfg.Type {
	case "acp":
		ag := agent.NewACPAgent(agent.ACPAgentConfig{
			Command:      agCfg.Command,
			Args:         agCfg.Args,
			Cwd:          agCfg.Cwd,
			Env:          agCfg.Env,
			Model:        agCfg.Model,
			SystemPrompt: agCfg.SystemPrompt,
		})
		if err := ag.Start(ctx); err != nil {
			log.Printf("[agent] failed to start ACP agent %q: %v", name, err)
			return nil
		}
		log.Printf("[agent] started ACP agent: %s (command=%s, type=%s, model=%s)", name, agCfg.Command, agCfg.Type, agCfg.Model)
		return ag
	case "cli":
		ag := agent.NewCLIAgent(agent.CLIAgentConfig{
			Name:         name,
			Command:      agCfg.Command,
			Args:         agCfg.Args,
			Cwd:          agCfg.Cwd,
			Env:          agCfg.Env,
			Model:        agCfg.Model,
			SystemPrompt: agCfg.SystemPrompt,
		})
		log.Printf("[agent] created CLI agent: %s (command=%s, type=%s, model=%s)", name, agCfg.Command, agCfg.Type, agCfg.Model)
		return ag
	case "http":
		if agCfg.Endpoint == "" {
			log.Printf("[agent] HTTP agent %q has no endpoint", name)
			return nil
		}
		ag := agent.NewHTTPAgent(agent.HTTPAgentConfig{
			Endpoint:     agCfg.Endpoint,
			APIKey:       agCfg.APIKey,
			Headers:      agCfg.Headers,
			Model:        agCfg.Model,
			SystemPrompt: agCfg.SystemPrompt,
			MaxHistory:   agCfg.MaxHistory,
		})
		log.Printf("[agent] created HTTP agent: %s (endpoint=%s, model=%s)", name, agCfg.Endpoint, agCfg.Model)
		return ag
	default:
		log.Printf("[agent] unknown type %q for %q", agCfg.Type, name)
		return nil
	}
}

// doLogin runs the interactive QR login flow and returns credentials.
func doLogin(ctx context.Context) (*ilink.Credentials, error) {
	fmt.Println("Fetching QR code...")
	qr, err := ilink.FetchQRCode(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch QR code: %w", err)
	}

	fmt.Println("\nScan this QR code with WeChat:")
	fmt.Println()
	qrterminal.GenerateWithConfig(qr.QRCodeImgContent, qrterminal.Config{
		Level:          qrterminal.L,
		Writer:         os.Stdout,
		HalfBlocks:     true,
		BlackChar:      qrterminal.BLACK_BLACK,
		WhiteBlackChar: qrterminal.WHITE_BLACK,
		WhiteChar:      qrterminal.WHITE_WHITE,
		BlackWhiteChar: qrterminal.BLACK_WHITE,
		QuietZone:      1,
	})
	fmt.Printf("\nQR URL: %s\n", qr.QRCodeImgContent)
	fmt.Println("\nWaiting for scan...")

	lastStatus := ""
	creds, err := ilink.PollQRStatus(ctx, qr.QRCode, func(status string) {
		if status != lastStatus {
			lastStatus = status
			switch status {
			case "scaned":
				fmt.Println("QR code scanned! Please confirm on your phone.")
			case "confirmed":
				fmt.Println("Login confirmed!")
			case "expired":
				fmt.Println("QR code expired.")
			}
		}
	})
	if err != nil {
		return nil, err
	}

	if err := ilink.SaveCredentials(creds); err != nil {
		return nil, fmt.Errorf("failed to save credentials: %w", err)
	}

	dir, _ := ilink.CredentialsPath()
	fmt.Printf("\nLogin successful! Credentials saved to %s\n", dir)
	fmt.Printf("Bot ID: %s\n\n", creds.ILinkBotID)
	return creds, nil
}

// --- Daemon mode ---

func weclawDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".weclaw")
}

func pidFile() string {
	return filepath.Join(weclawDir(), "weclaw.pid")
}

func logFile() string {
	return filepath.Join(weclawDir(), "weclaw.log")
}

// runDaemon spawns weclaw start (without --daemon) as a background process.
func runDaemon() error {
	// Kill any existing weclaw processes before starting a new one
	stopAllWeclaw()

	// Ensure log directory exists
	if err := os.MkdirAll(weclawDir(), 0o700); err != nil {
		return fmt.Errorf("create weclaw dir: %w", err)
	}

	// Open log file
	lf, err := os.OpenFile(logFile(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return fmt.Errorf("open log file: %w", err)
	}

	// Re-exec ourselves without --daemon
	exe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("find executable: %w", err)
	}

	cmd := exec.Command(exe, "start", "-f")
	cmd.Stdout = lf
	cmd.Stderr = lf
	setSysProcAttr(cmd)

	if err := cmd.Start(); err != nil {
		lf.Close()
		return fmt.Errorf("start daemon: %w", err)
	}

	// Save PID
	pid := cmd.Process.Pid
	os.WriteFile(pidFile(), []byte(fmt.Sprintf("%d", pid)), 0o644)

	// Detach — don't wait
	cmd.Process.Release()
	lf.Close()

	fmt.Printf("weclaw started in background (pid=%d)\n", pid)
	fmt.Printf("Log: %s\n", logFile())
	fmt.Printf("Stop: weclaw stop\n")
	return nil
}

func readPid() (int, error) {
	data, err := os.ReadFile(pidFile())
	if err != nil {
		return 0, err
	}
	var pid int
	if _, err := fmt.Sscanf(string(data), "%d", &pid); err != nil {
		return 0, err
	}
	return pid, nil
}

func processExists(pid int) bool {
	p, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	// Signal 0 checks if process exists without killing it
	return p.Signal(syscall.Signal(0)) == nil
}

// stopAllWeclaw kills all running weclaw processes (by PID file and by process scan).
func stopAllWeclaw() {
	// 1. Kill by PID file
	if pid, err := readPid(); err == nil && processExists(pid) {
		if p, err := os.FindProcess(pid); err == nil {
			_ = p.Signal(syscall.SIGTERM)
		}
	}
	os.Remove(pidFile())

	// 2. Kill any remaining weclaw processes by scanning
	exe, err := os.Executable()
	if err != nil {
		return
	}
	// Use pkill to kill all processes matching the executable path
	_ = exec.Command("pkill", "-f", exe+" start").Run()
	time.Sleep(500 * time.Millisecond)
}

```

[⬆ 回到目录](#toc)

## cmd/status.go

```go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check if weclaw is running in background",
	RunE: func(cmd *cobra.Command, args []string) error {
		pid, err := readPid()
		if err != nil {
			fmt.Println("weclaw is not running")
			return nil
		}

		if processExists(pid) {
			fmt.Printf("weclaw is running (pid=%d)\n", pid)
			fmt.Printf("Log: %s\n", logFile())
		} else {
			fmt.Println("weclaw is not running (stale pid file)")
		}
		return nil
	},
}

```

[⬆ 回到目录](#toc)

## cmd/stop.go

```go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the background weclaw process",
	RunE: func(cmd *cobra.Command, args []string) error {
		stopAllWeclaw()
		fmt.Println("weclaw stopped")
		return nil
	},
}

```

[⬆ 回到目录](#toc)

## cmd/update.go

```go
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const githubRepo = "fastclaw-ai/weclaw"

func init() {
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(upgradeCmd)
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("weclaw %s (%s/%s)\n", Version, runtime.GOOS, runtime.GOARCH)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update weclaw to the latest version and restart",
	RunE:  runUpdate,
}

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Update weclaw to the latest version and restart (alias for update)",
	RunE:  runUpdate,
}

func runUpdate(cmd *cobra.Command, args []string) error {
	// 1. Get latest version
	fmt.Println("Checking for updates...")
	latest, err := getLatestVersion()
	if err != nil {
		return fmt.Errorf("failed to check latest version: %w", err)
	}

	if latest == Version {
		fmt.Printf("Already up to date (%s)\n", Version)
		return nil
	}

	fmt.Printf("Current: %s -> Latest: %s\n", Version, latest)

	// 2. Download new binary
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	filename := fmt.Sprintf("weclaw_%s_%s", goos, goarch)
	url := fmt.Sprintf("https://github.com/%s/releases/download/%s/%s", githubRepo, latest, filename)

	fmt.Printf("Downloading %s...\n", url)
	tmpFile, err := downloadFile(url)
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}
	defer os.Remove(tmpFile)

	// 3. Replace current binary
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("find executable: %w", err)
	}
	// Resolve symlinks
	if resolved, err := resolveSymlink(exePath); err == nil {
		exePath = resolved
	}

	if err := replaceBinary(tmpFile, exePath); err != nil {
		return fmt.Errorf("replace binary: %w", err)
	}

	// Clear macOS quarantine/provenance attributes to avoid Gatekeeper killing the binary
	if runtime.GOOS == "darwin" {
		exec.Command("xattr", "-d", "com.apple.quarantine", exePath).Run()
		exec.Command("xattr", "-d", "com.apple.provenance", exePath).Run()
	}

	fmt.Printf("Updated to %s\n", latest)

	// 4. Restart if running in background
	pid, pidErr := readPid()
	if pidErr == nil && processExists(pid) {
		fmt.Println("Stopping old process...")
		if p, err := os.FindProcess(pid); err == nil {
			p.Signal(os.Interrupt)
		}
		// Wait for old process to exit
		for i := 0; i < 20; i++ {
			if !processExists(pid) {
				break
			}
			time.Sleep(500 * time.Millisecond)
		}
		os.Remove(pidFile())

		fmt.Println("Starting new version...")
		if err := runDaemon(); err != nil {
			log.Printf("Failed to restart: %v", err)
			fmt.Println("Update complete. Please run 'weclaw start' manually.")
		}
	} else {
		fmt.Println("Update complete. Run 'weclaw start' to start.")
	}

	return nil
}

func getLatestVersion() (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", githubRepo))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API returned %d", resp.StatusCode)
	}

	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err
	}
	return release.TagName, nil
}

func downloadFile(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	tmp, err := os.CreateTemp("", "weclaw-update-*")
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(tmp, resp.Body); err != nil {
		tmp.Close()
		os.Remove(tmp.Name())
		return "", err
	}
	tmp.Close()

	if err := os.Chmod(tmp.Name(), 0o755); err != nil {
		os.Remove(tmp.Name())
		return "", err
	}

	return tmp.Name(), nil
}

func replaceBinary(src, dst string) error {
	// Check if we can write directly
	if err := os.Rename(src, dst); err == nil {
		return nil
	}

	// Try with sudo on Unix
	if runtime.GOOS != "windows" {
		fmt.Printf("Installing to %s (requires sudo)...\n", dst)
		cmd := exec.Command("sudo", "cp", src, dst)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	return fmt.Errorf("cannot write to %s", dst)
}

func resolveSymlink(path string) (string, error) {
	for {
		target, err := os.Readlink(path)
		if err != nil {
			return path, nil
		}
		if !strings.HasPrefix(target, "/") {
			// Relative symlink
			dir := path[:strings.LastIndex(path, "/")+1]
			target = dir + target
		}
		path = target
	}
}

```

[⬆ 回到目录](#toc)

## config/config.go

```go
package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Config holds the application configuration.
type Config struct {
	DefaultAgent string                 `json:"default_agent"`
	APIAddr      string                 `json:"api_addr,omitempty"`
	SaveDir      string                 `json:"save_dir,omitempty"`
	Agents       map[string]AgentConfig `json:"agents"`
}

// AgentConfig holds configuration for a single agent.
type AgentConfig struct {
	Type         string            `json:"type"`                    // "acp", "cli", or "http"
	Command      string            `json:"command,omitempty"`       // binary path (cli/acp type)
	Args         []string          `json:"args,omitempty"`          // extra args for command (e.g. ["acp"] for cursor)
	Aliases      []string          `json:"aliases,omitempty"`       // custom trigger commands (e.g. ["gpt", "4o"])
	Cwd          string            `json:"cwd,omitempty"`           // working directory (workspace)
	Env          map[string]string `json:"env,omitempty"`           // extra environment variables (cli/acp type)
	Model        string            `json:"model,omitempty"`         // model name
	SystemPrompt string            `json:"system_prompt,omitempty"` // system prompt
	Endpoint     string            `json:"endpoint,omitempty"`      // API endpoint (http type)
	APIKey       string            `json:"api_key,omitempty"`       // API key (http type)
	Headers      map[string]string `json:"headers,omitempty"`       // extra HTTP headers (http type)
	MaxHistory   int               `json:"max_history,omitempty"`   // max history (http type)
}

// BuildAliasMap builds a map from custom alias to agent name from all agent configs.
// It logs warnings for conflicts: duplicate aliases and aliases shadowing agent keys.
func BuildAliasMap(agents map[string]AgentConfig) map[string]string {
	// Built-in commands that cannot be overridden
	reserved := map[string]bool{
		"info": true, "help": true, "new": true, "clear": true, "cwd": true,
	}

	m := make(map[string]string)
	for name, cfg := range agents {
		for _, alias := range cfg.Aliases {
			if reserved[alias] {
				log.Printf("[config] WARNING: alias %q for agent %q conflicts with built-in command, ignored", alias, name)
				continue
			}
			if existing, ok := m[alias]; ok {
				log.Printf("[config] WARNING: alias %q is defined by both %q and %q, using %q", alias, existing, name, name)
			}
			m[alias] = name
		}
	}

	// Warn if a custom alias shadows an agent key
	for alias, target := range m {
		if _, isAgent := agents[alias]; isAgent && alias != target {
			log.Printf("[config] WARNING: alias %q (-> %q) shadows agent key %q", alias, target, alias)
		}
	}

	return m
}

// DefaultConfig returns an empty configuration.
func DefaultConfig() *Config {
	return &Config{
		Agents: make(map[string]AgentConfig),
	}
}

// ConfigPath returns the path to the config file.
func ConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".weclaw", "config.json"), nil
}

// Load loads configuration from disk and environment variables.
func Load() (*Config, error) {
	cfg := DefaultConfig()

	path, err := ConfigPath()
	if err != nil {
		return cfg, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			loadEnv(cfg)
			return cfg, nil
		}
		return nil, fmt.Errorf("read config: %w", err)
	}

	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	if cfg.Agents == nil {
		cfg.Agents = make(map[string]AgentConfig)
	}

	loadEnv(cfg)
	return cfg, nil
}

func loadEnv(cfg *Config) {
	if v := os.Getenv("WECLAW_DEFAULT_AGENT"); v != "" {
		cfg.DefaultAgent = v
	}
	if v := os.Getenv("WECLAW_API_ADDR"); v != "" {
		cfg.APIAddr = v
	}
	if v := os.Getenv("WECLAW_SAVE_DIR"); v != "" {
		cfg.SaveDir = v
	}
}

// Save saves the configuration to disk.
func Save(cfg *Config) error {
	path, err := ConfigPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return fmt.Errorf("create config dir: %w", err)
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}

	return os.WriteFile(path, data, 0o600)
}

```

[⬆ 回到目录](#toc)

## config/config_test.go

```go
package config

import (
	"encoding/json"
	"testing"
)

func TestAgentConfigUnmarshalEnv(t *testing.T) {
	var cfg Config
	data := []byte(`{
		"agents": {
			"claude": {
				"type": "cli",
				"command": "claude",
				"env": {
					"ANTHROPIC_API_KEY": "test-key",
					"EMPTY": ""
				}
			}
		}
	}`)

	if err := json.Unmarshal(data, &cfg); err != nil {
		t.Fatalf("unmarshal config: %v", err)
	}

	ag, ok := cfg.Agents["claude"]
	if !ok {
		t.Fatalf("expected claude agent config")
	}
	if got := ag.Env["ANTHROPIC_API_KEY"]; got != "test-key" {
		t.Fatalf("ANTHROPIC_API_KEY = %q, want %q", got, "test-key")
	}
	if got, ok := ag.Env["EMPTY"]; !ok || got != "" {
		t.Fatalf("EMPTY = %q, present=%v; want empty string present", got, ok)
	}
}

func TestAgentConfigMarshalEnvRoundTrip(t *testing.T) {
	cfg := Config{
		Agents: map[string]AgentConfig{
			"claude": {
				Type:    "cli",
				Command: "claude",
				Env: map[string]string{
					"ANTHROPIC_API_KEY": "test-key",
					"EMPTY":             "",
				},
			},
		},
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("marshal config: %v", err)
	}

	var decoded Config
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("round-trip unmarshal: %v", err)
	}

	got := decoded.Agents["claude"].Env
	if got["ANTHROPIC_API_KEY"] != "test-key" || got["EMPTY"] != "" {
		t.Fatalf("round-trip env = %#v", got)
	}
}

func TestAgentConfigWithoutEnvStillLoads(t *testing.T) {
	var cfg Config
	data := []byte(`{
		"agents": {
			"claude": {
				"type": "cli",
				"command": "claude"
			}
		}
	}`)

	if err := json.Unmarshal(data, &cfg); err != nil {
		t.Fatalf("unmarshal config without env: %v", err)
	}

	if cfg.Agents["claude"].Env != nil {
		t.Fatalf("Env = %#v, want nil", cfg.Agents["claude"].Env)
	}
}

func TestDefaultConfigInitializesAgentsMap(t *testing.T) {
	cfg := DefaultConfig()
	if cfg.Agents == nil {
		t.Fatal("DefaultConfig() Agents = nil, want initialized map")
	}
}

func TestLoadEnvOverridesTopLevelOnly(t *testing.T) {
	t.Setenv("WECLAW_DEFAULT_AGENT", "codex")
	t.Setenv("WECLAW_API_ADDR", "127.0.0.1:18011")

	cfg := DefaultConfig()
	cfg.Agents["claude"] = AgentConfig{
		Type: "cli",
		Env: map[string]string{
			"KEEP": "value",
		},
	}

	loadEnv(cfg)

	if cfg.DefaultAgent != "codex" {
		t.Fatalf("DefaultAgent = %q, want %q", cfg.DefaultAgent, "codex")
	}
	if cfg.APIAddr != "127.0.0.1:18011" {
		t.Fatalf("APIAddr = %q, want %q", cfg.APIAddr, "127.0.0.1:18011")
	}
	if got := cfg.Agents["claude"].Env["KEEP"]; got != "value" {
		t.Fatalf("agent env = %q, want preserved value", got)
	}
}

```

[⬆ 回到目录](#toc)

## config/detect.go

```go
package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// agentCandidate defines one way to run an agent.
// Multiple candidates can map to the same agent name; the first detected wins.
type agentCandidate struct {
	Name      string   // agent name (e.g. "claude", "codex")
	Binary    string   // binary to look up in PATH
	Args      []string // extra args (e.g. ["acp"] for cursor)
	CheckArgs []string // optional capability probe args (must exit 0)
	Type      string   // "acp", "cli"
	Model     string   // default model
}

// agentCandidates is ordered by priority: for each agent name, earlier entries
// are preferred. E.g. claude ACP is tried before claude CLI.
var agentCandidates = []agentCandidate{
	// claude: prefer ACP, fallback to CLI
	{Name: "claude", Binary: "claude-agent-acp", Type: "acp", Model: "sonnet"},
	{Name: "claude", Binary: "claude", Type: "cli", Model: "sonnet"},
	// codex: prefer ACP, fallback to CLI
	{Name: "codex", Binary: "codex-acp", Type: "acp", Model: ""},
	{Name: "codex", Binary: "codex", Args: []string{"app-server", "--listen", "stdio://"}, CheckArgs: []string{"app-server", "--help"}, Type: "acp", Model: ""},
	{Name: "codex", Binary: "codex", Type: "cli", Model: ""},
	// ACP-only agents
	{Name: "cursor", Binary: "agent", Args: []string{"acp"}, Type: "acp", Model: ""},
	{Name: "kimi", Binary: "kimi", Args: []string{"acp"}, Type: "acp", Model: ""},
	{Name: "gemini", Binary: "gemini", Args: []string{"--acp"}, Type: "acp", Model: ""},
	{Name: "opencode", Binary: "opencode", Args: []string{"acp"}, Type: "acp", Model: ""},
	{Name: "openclaw", Binary: "openclaw", Type: "acp", Model: "openclaw:main"}, // args built dynamically
	{Name: "pi", Binary: "pi-acp", Type: "acp", Model: ""},
	{Name: "copilot", Binary: "copilot", Args: []string{"--acp", "--stdio"}, Type: "acp", Model: ""},
	{Name: "droid", Binary: "droid", Args: []string{"exec", "--output-format", "acp"}, Type: "acp", Model: ""},
	{Name: "iflow", Binary: "iflow", Args: []string{"--experimental-acp"}, Type: "acp", Model: ""},
	{Name: "kiro", Binary: "kiro-cli", Args: []string{"acp"}, Type: "acp", Model: ""},
	{Name: "qwen", Binary: "qwen", Args: []string{"--acp"}, Type: "acp", Model: ""},
}

// defaultOrder defines the priority for choosing the default agent.
// Lower index = higher priority.
var defaultOrder = []string{
	"claude", "codex", "cursor", "kimi", "gemini", "opencode", "openclaw",
	"pi", "copilot", "droid", "iflow", "kiro", "qwen",
}

// DetectAndConfigure auto-detects local agents and populates the config.
// For each agent name, it picks the highest-priority candidate (acp > cli).
// Returns true if the config was modified.
func DetectAndConfigure(cfg *Config) bool {
	modified := false

	for _, candidate := range agentCandidates {
		// Skip if this agent name is already configured
		if _, exists := cfg.Agents[candidate.Name]; exists {
			continue
		}

		path, err := lookPath(candidate.Binary)
		if err != nil {
			continue
		}

		// Run capability probe if specified
		if len(candidate.CheckArgs) > 0 && !commandProbe(path, candidate.CheckArgs) {
			log.Printf("[config] skipping %s at %s (type=%s): probe failed (%v)", candidate.Name, path, candidate.Type, candidate.CheckArgs)
			continue
		}

		log.Printf("[config] auto-detected %s at %s (type=%s)", candidate.Name, path, candidate.Type)
		cfg.Agents[candidate.Name] = AgentConfig{
			Type:    candidate.Type,
			Command: path,
			Args:    candidate.Args,
			Model:   candidate.Model,
		}
		modified = true
	}

	// Special handling for openclaw: prefer HTTP mode over ACP to avoid
	// session routing conflicts with openclaw-weixin plugin (see #9).
	// Priority: HTTP (gateway) > ACP (with user-configured --session) > skip.
	if agCfg, exists := cfg.Agents["openclaw"]; exists && agCfg.Type == "acp" && len(agCfg.Args) == 0 {
		gwURL, gwToken, gwPassword := loadOpenclawGateway()
		if gwURL != "" {
			// Prefer HTTP mode — no session routing issues
			httpURL := gwURL
			httpURL = strings.Replace(httpURL, "wss://", "https://", 1)
			httpURL = strings.Replace(httpURL, "ws://", "http://", 1)
			endpoint := strings.TrimRight(httpURL, "/") + "/v1/chat/completions"
			log.Printf("[config] openclaw using HTTP mode: %s", endpoint)
			cfg.Agents["openclaw"] = AgentConfig{
				Type:     "http",
				Endpoint: endpoint,
				APIKey:   gwToken,
				Headers:  map[string]string{"x-openclaw-scopes": "operator.write"},
				Model:    "openclaw:main",
			}
			modified = true

			// Also register openclaw-acp as a separate agent for users who want ACP
			if _, apcExists := cfg.Agents["openclaw-acp"]; !apcExists {
				args := []string{"acp", "--url", gwURL}
				if gwToken != "" {
					args = append(args, "--token", gwToken)
				} else if gwPassword != "" {
					args = append(args, "--password", gwPassword)
				}
				cfg.Agents["openclaw-acp"] = AgentConfig{
					Type:    "acp",
					Command: agCfg.Command,
					Args:    args,
					Model:   "openclaw:main",
				}
				log.Printf("[config] openclaw ACP also available as 'openclaw-acp' (use /openclaw-acp to switch)")
			}
		} else {
			log.Printf("[config] openclaw binary found but no gateway config, skipping")
			delete(cfg.Agents, "openclaw")
			modified = true
		}
	}

	// Fallback: if openclaw still not configured, try HTTP via gateway config.
	if _, exists := cfg.Agents["openclaw"]; !exists {
		gwURL, gwToken, _ := loadOpenclawGateway()
		if gwURL != "" {
			httpURL := gwURL
			httpURL = strings.Replace(httpURL, "wss://", "https://", 1)
			httpURL = strings.Replace(httpURL, "ws://", "http://", 1)
			endpoint := strings.TrimRight(httpURL, "/") + "/v1/chat/completions"
			log.Printf("[config] using openclaw HTTP: %s", endpoint)
			cfg.Agents["openclaw"] = AgentConfig{
				Type:     "http",
				Endpoint: endpoint,
				APIKey:   gwToken,
				Headers:  map[string]string{"x-openclaw-scopes": "operator.write"},
				Model:    "openclaw:main",
			}
			modified = true
		}
	}

	// Pick the highest-priority default agent.
	if cfg.DefaultAgent == "" || !agentExists(cfg, cfg.DefaultAgent) {
		for _, name := range defaultOrder {
			if _, ok := cfg.Agents[name]; ok {
				if cfg.DefaultAgent != name {
					log.Printf("[config] setting default agent: %s", name)
					cfg.DefaultAgent = name
					modified = true
				}
				break
			}
		}
	}

	return modified
}

// loadOpenclawGateway resolves openclaw gateway connection info.
// Priority: env vars > ~/.openclaw/openclaw.json.
// Returns (url, token, password). url="" means not configured.
func loadOpenclawGateway() (gwURL, gwToken, gwPassword string) {
	// 1. Environment variables take priority
	gwURL = os.Getenv("OPENCLAW_GATEWAY_URL")
	gwToken = os.Getenv("OPENCLAW_GATEWAY_TOKEN")
	gwPassword = os.Getenv("OPENCLAW_GATEWAY_PASSWORD")
	if gwURL != "" {
		return
	}

	// 2. Read from ~/.openclaw/openclaw.json
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}

	data, err := os.ReadFile(filepath.Join(home, ".openclaw", "openclaw.json"))
	if err != nil {
		return
	}

	var ocCfg struct {
		Gateway struct {
			Port int    `json:"port"`
			Mode string `json:"mode"`
			Auth struct {
				Mode     string `json:"mode"`
				Token    string `json:"token"`
				Password string `json:"password"`
			} `json:"auth"`
			Remote struct {
				URL   string `json:"url"`
				Token string `json:"token"`
			} `json:"remote"`
		} `json:"gateway"`
	}
	if err := json.Unmarshal(data, &ocCfg); err != nil {
		log.Printf("[config] failed to parse openclaw config: %v", err)
		return
	}

	gw := ocCfg.Gateway

	// Remote gateway (gateway.remote.url)
	if gw.Remote.URL != "" {
		gwURL = gw.Remote.URL
		gwToken = gw.Remote.Token
		return
	}

	// Local gateway (gateway.port + gateway.auth)
	if gw.Port > 0 {
		gwURL = fmt.Sprintf("ws://127.0.0.1:%d", gw.Port)
		switch gw.Auth.Mode {
		case "token":
			gwToken = gw.Auth.Token
		case "password":
			gwPassword = gw.Auth.Password
		}
		return
	}

	return
}

// commandProbe runs a binary with args and returns true if it exits 0.
func commandProbe(binary string, args []string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, binary, args...)
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	return cmd.Run() == nil
}

func agentExists(cfg *Config, name string) bool {
	_, ok := cfg.Agents[name]
	return ok
}

// lookPath finds a binary by name. It first tries exec.LookPath (fast, uses
// current PATH). If that fails, it falls back to resolving via a login shell
// which sources the user's profile (~/.zshrc, ~/.bashrc) — this picks up
// binaries installed through version managers like nvm, mise, etc. that only
// add their paths in interactive shells.
func lookPath(binary string) (string, error) {
	// Fast path: binary is in current PATH
	if p, err := exec.LookPath(binary); err == nil {
		return p, nil
	}

	// Fallback: resolve via login interactive shell (sources .zshrc/.bashrc)
	shell := "zsh"
	if runtime.GOOS != "darwin" {
		shell = "bash"
	}
	out, err := exec.Command(shell, "-lic", "which "+binary).Output()
	if err != nil {
		return "", fmt.Errorf("not found: %s", binary)
	}
	p := strings.TrimSpace(string(out))
	if p == "" || strings.Contains(p, "not found") {
		return "", fmt.Errorf("not found: %s", binary)
	}
	log.Printf("[config] resolved %s via login shell: %s", binary, p)
	return p, nil
}

```

[⬆ 回到目录](#toc)

## config/detect_test.go

```go
package config

import (
	"os"
	"os/exec"
	"testing"
)

// TestLookPath_InPath verifies that lookPath finds binaries already in PATH.
func TestLookPath_InPath(t *testing.T) {
	p, err := lookPath("ls")
	if err != nil {
		t.Fatalf("expected to find ls, got error: %v", err)
	}
	if p == "" {
		t.Fatal("expected non-empty path for ls")
	}
}

// TestLookPath_NotExist verifies that lookPath returns an error for missing binaries.
func TestLookPath_NotExist(t *testing.T) {
	_, err := lookPath("nonexistent-binary-xyz-12345")
	if err == nil {
		t.Fatal("expected error for nonexistent binary")
	}
}

// TestLookPath_LoginShellFallback reproduces the daemon scenario:
// PATH is stripped to system-only dirs (no nvm), so exec.LookPath fails,
// but lookPath resolves claude via login shell fallback.
func TestLookPath_LoginShellFallback(t *testing.T) {
	// Precondition: claude must be discoverable via login shell (i.e. nvm in .zshrc)
	fullPath, err := exec.LookPath("claude")
	if err != nil {
		t.Skip("claude not installed, skipping login shell fallback test")
	}

	// Simulate daemon environment: strip PATH to system-only dirs
	origPath := os.Getenv("PATH")
	os.Setenv("PATH", "/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin")
	defer os.Setenv("PATH", origPath)

	// Reproduce the bug: exec.LookPath must fail under stripped PATH
	_, err = exec.LookPath("claude")
	if err == nil {
		t.Skip("claude found in minimal PATH, cannot reproduce nvm issue")
	}

	// Verify fix: lookPath should find claude via login shell
	p, err := lookPath("claude")
	if err != nil {
		t.Fatalf("lookPath should find claude via login shell, got: %v", err)
	}
	if p != fullPath {
		t.Logf("resolved path differs: direct=%s, login-shell=%s (acceptable)", fullPath, p)
	}
	t.Logf("lookPath resolved claude via login shell: %s", p)
}

// TestDetectAndConfigure_StrippedPath is an end-to-end test:
// empty config + stripped PATH → DetectAndConfigure should still find claude.
func TestDetectAndConfigure_StrippedPath(t *testing.T) {
	if _, err := exec.LookPath("claude"); err != nil {
		t.Skip("claude not installed, skipping")
	}

	origPath := os.Getenv("PATH")
	os.Setenv("PATH", "/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin")
	defer os.Setenv("PATH", origPath)

	cfg := DefaultConfig()
	DetectAndConfigure(cfg)

	agent, ok := cfg.Agents["claude"]
	if !ok {
		t.Fatal("expected claude to be detected via login shell fallback")
	}
	if agent.Type != "cli" {
		t.Fatalf("expected type=cli, got %s", agent.Type)
	}
	t.Logf("detected claude: type=%s, command=%s", agent.Type, agent.Command)
}

```

[⬆ 回到目录](#toc)

## docs/README_CN.md

```markdown
# WeClaw

[English](README.md)

微信 AI Agent 桥接器 — 将微信消息接入 AI Agent（Claude、Codex、Gemini、Kimi 等）。

> 本项目参考 [@tencent-weixin/openclaw-weixin](https://npmx.dev/package/@tencent-weixin/openclaw-weixin) 实现，仅限个人学习，勿做他用。

|                                                 |                                                 |                                                 |
| :---------------------------------------------: | :---------------------------------------------: | :---------------------------------------------: |
| <img src="previews/preview1.png" width="280" /> | <img src="previews/preview2.png" width="280" /> | <img src="previews/preview3.png" width="280" /> |

## 快速开始

```bash
# 一键安装
curl -sSL https://raw.githubusercontent.com/fastclaw-ai/weclaw/main/install.sh | sh

# 启动（首次运行会弹出微信扫码登录）
weclaw start
```

就这么简单。首次启动时，WeClaw 会：

1. 显示二维码 — 用微信扫码登录
2. 自动检测已安装的 AI Agent（Claude、Codex、Gemini 等）
3. 保存配置到 `~/.weclaw/config.json`
4. 开始接收和回复微信消息

使用 `weclaw login` 可以添加更多微信账号。

### 其他安装方式

```bash
# 通过 Go 安装
go install github.com/fastclaw-ai/weclaw@latest

# 通过 Docker
docker run -it -v ~/.weclaw:/root/.weclaw ghcr.io/fastclaw-ai/weclaw start
```

## 架构

<p align="center">
  <img src="previews/architecture.png" width="600" />
</p>

**Agent 接入模式：**

| 模式 | 工作方式                                                         | 支持的 Agent                                            |
| ---- | ---------------------------------------------------------------- | ------------------------------------------------------- |
| ACP  | 长驻子进程，通过 stdio JSON-RPC 通信。速度最快，复用进程和会话。 | Claude, Codex, Kimi, Gemini, Cursor, OpenCode, OpenClaw |
| CLI  | 每条消息启动一个新进程，支持通过 `--resume` 恢复会话。           | Claude (`claude -p`)、Codex (`codex exec`)              |
| HTTP | OpenAI 兼容的 Chat Completions API。                             | OpenClaw（HTTP 回退）                                   |

同时存在 ACP 和 CLI 时，自动优先选择 ACP。

## 聊天命令

在微信中发送以下命令：

| 命令                    | 说明                     |
| ----------------------- | ------------------------ |
| `你好`                  | 发送给默认 Agent         |
| `/codex 写一个排序函数` | 发送给指定 Agent         |
| `/cc 解释一下这段代码`  | 通过别名发送             |
| `/claude`               | 切换默认 Agent 为 Claude |
| `/cwd /path/to/project` | 切换工作目录             |
| `/new`                  | 开始新对话（清除会话）   |
| `/info`                 | 查看当前 Agent 信息      |
| `/help`                 | 查看帮助信息             |

### 快捷别名

| 别名   | Agent    |
| ------ | -------- |
| `/cc`  | Claude   |
| `/cx`  | Codex    |
| `/cs`  | Cursor   |
| `/km`  | Kimi     |
| `/gm`  | Gemini   |
| `/ocd` | OpenCode |
| `/oc`  | OpenClaw |

也可以在配置文件中为每个 Agent 自定义触发命令：

```json
{
  "agents": {
    "claude": {
      "type": "acp",
      "aliases": ["ai", "c"]
    }
  }
}
```

然后 `/ai 你好` 或 `/c 你好` 就会路由到 claude。

切换默认 Agent 会写入配置文件，重启后仍然生效。

## 富媒体消息

WeClaw 支持收发图片、视频、文件和语音消息。

**语音消息：** 在微信中发送语音消息时，WeClaw 会自动使用微信的语音转文字功能，将转写后的文本发送给 AI Agent。重复的语音消息事件会自动去重。

**Agent 回复自动处理：** 当 AI Agent 返回包含图片的 markdown（`![](url)`）时，WeClaw 会自动提取图片 URL，下载文件，上传到微信 CDN（AES-128-ECB 加密），然后作为图片消息发送。

**Markdown 转换：** Agent 的回复会自动从 markdown 转为纯文本再发送 — 代码块去掉围栏、链接只保留文字、加粗斜体标记去除等。

## 主动推送消息

无需等待用户发消息，主动向微信用户推送消息。

**命令行：**

```bash
# 发送文本
weclaw send --to "user_id@im.wechat" --text "你好，来自 weclaw"

# 发送图片
weclaw send --to "user_id@im.wechat" --media "https://example.com/photo.png"

# 发送文本 + 图片
weclaw send --to "user_id@im.wechat" --text "看看这个" --media "https://example.com/photo.png"

# 发送文件
weclaw send --to "user_id@im.wechat" --media "https://example.com/report.pdf"
```

**HTTP API**（`weclaw start` 运行时，默认监听 `127.0.0.1:18011`）：

```bash
# 发送文本
curl -X POST http://127.0.0.1:18011/api/send \
  -H "Content-Type: application/json" \
  -d '{"to": "user_id@im.wechat", "text": "你好，来自 weclaw"}'

# 发送图片
curl -X POST http://127.0.0.1:18011/api/send \
  -H "Content-Type: application/json" \
  -d '{"to": "user_id@im.wechat", "media_url": "https://example.com/photo.png"}'

# 发送文本 + 媒体
curl -X POST http://127.0.0.1:18011/api/send \
  -H "Content-Type: application/json" \
  -d '{"to": "user_id@im.wechat", "text": "看看这个", "media_url": "https://example.com/photo.png"}'
```

支持的媒体类型：图片（png、jpg、gif、webp）、视频（mp4、mov）、文件（pdf、doc、zip 等）。

设置 `WECLAW_API_ADDR` 环境变量可更改监听地址（如 `0.0.0.0:18011`）。

## 配置

配置文件路径：`~/.weclaw/config.json`

```json
{
  "default_agent": "claude",
  "agents": {
    "claude": {
      "type": "acp",
      "command": "/usr/local/bin/claude-agent-acp",
      "env": {
        "ANTHROPIC_API_KEY": "sk-ant-xxx"
      },
      "model": "sonnet"
    },
    "codex": {
      "type": "acp",
      "command": "/usr/local/bin/codex-acp",
      "env": {
        "OPENAI_API_KEY": "sk-xxx"
      }
    },
    "openclaw": {
      "type": "http",
      "endpoint": "https://api.example.com/v1/chat/completions",
      "api_key": "sk-xxx",
      "model": "openclaw:main"
    }
  }
}
```

环境变量：

- `WECLAW_DEFAULT_AGENT` — 覆盖默认 Agent
- `OPENCLAW_GATEWAY_URL` — OpenClaw HTTP 回退地址
- `OPENCLAW_GATEWAY_TOKEN` — OpenClaw API Token

自定义 agent cli 环境变量

```json
{
  "default_agent": "...",
  "agents": {
    "...": {
      ...
      "env": {
        "ENV_NAME": "ENV_VALUE"
      }
    },
  }
}
```

### 权限配置

部分 Agent 默认需要交互式权限确认，在微信场景下无法操作会导致卡住。可通过 `args` 配置跳过：

| Agent | 参数 | 说明 |
|-------|------|------|
| Claude (CLI) | `--dangerously-skip-permissions` | 跳过所有工具权限确认 |
| Codex (CLI) | `--skip-git-repo-check` | 允许在非 git 仓库目录运行 |

配置示例：

```json
{
  "claude": {
    "type": "cli",
    "command": "/usr/local/bin/claude",
    "cwd": "/home/user/my-project",
    "args": ["--dangerously-skip-permissions"]
  },
  "codex": {
    "type": "cli",
    "command": "/usr/local/bin/codex",
    "cwd": "/home/user/my-project",
    "args": ["--skip-git-repo-check"]
  }
}
```

通过 `cwd` 指定 Agent 的工作目录（workspace）。不设置则默认为 `~/.weclaw/workspace`。

> **注意：** 这些参数会跳过安全检查，请了解风险后再启用。ACP 模式的 Agent 会自动处理权限，无需配置。

## 后台运行

```bash
# 启动（默认后台运行）
weclaw start

# 查看状态
weclaw status

# 停止
weclaw stop

# 前台运行（调试用）
weclaw start -f
```

日志输出到 `~/.weclaw/weclaw.log`。

### 系统服务（开机自启）

**macOS (launchd)：**

```bash
cp service/com.fastclaw.weclaw.plist ~/Library/LaunchAgents/
launchctl load ~/Library/LaunchAgents/com.fastclaw.weclaw.plist
```

**Linux (systemd)：**

```bash
sudo cp service/weclaw.service /etc/systemd/system/
sudo systemctl enable --now weclaw
```

## Docker

```bash
# 构建
docker build -t weclaw .

# 登录（交互式，扫描二维码）
docker run -it -v ~/.weclaw:/root/.weclaw weclaw login

# 使用 HTTP Agent 启动
docker run -d --name weclaw \
  -v ~/.weclaw:/root/.weclaw \
  -e OPENCLAW_GATEWAY_URL=https://api.example.com \
  -e OPENCLAW_GATEWAY_TOKEN=sk-xxx \
  weclaw

# 查看日志
docker logs -f weclaw
```

> 注意：ACP 和 CLI 模式需要容器内有对应的 Agent 二进制文件。
> 默认镜像只包含 WeClaw 本体。如需使用 ACP/CLI Agent，请挂载二进制文件或构建自定义镜像。
> HTTP 模式开箱即用。

## 发版

```bash
# 打 tag 触发 GitHub Actions 自动构建发版
git tag v0.1.0
git push origin v0.1.0
```

自动构建 `darwin/linux/windows` x `amd64/arm64` 的二进制，创建 GitHub Release 并上传所有产物和校验文件。

## 更新

```bash
# 更新到最新版本（运行中会自动重启）
weclaw update

# 查看当前版本
weclaw version
```

## 开发

```bash
# 热重载
make dev

# 编译
go build -o weclaw .

# 运行
./weclaw start
```

## 贡献者

<a href="https://github.com/fastclaw-ai/weclaw/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=fastclaw-ai/weclaw" />
</a>

## Star 趋势

[![Star History Chart](https://api.star-history.com/svg?repos=fastclaw-ai/weclaw&type=Timeline)](https://star-history.com/#fastclaw-ai/weclaw&Timeline)

## 许可证

[MIT](LICENSE)

```

[⬆ 回到目录](#toc)

## docs/weclaw-20260330-docs.md

```markdown
# Project Documentation

- **Generated at:** 2026-03-30 14:06:55
- **Root Dir:** `.`
- **File Count:** 48
- **Total Lines:** 8409
- **Total Size:** 225.10 KB

<a name="toc"></a>
## 📂 扫描目录
- [.air.toml](#file-.air.toml) (52 lines, 0.91 KB)
- [.dockerignore](#file-.dockerignore) (13 lines, 0.09 KB)
- [.gitignore](#file-.gitignore) (23 lines, 0.16 KB)
- [Dockerfile](#file-dockerfile) (16 lines, 0.35 KB)
- [LICENSE](#file-license) (21 lines, 1.04 KB)
- [Makefile](#file-makefile) (2 lines, 0.03 KB)
- [README.md](#file-readme.md) (343 lines, 8.63 KB)
- [agent/acp_agent.go](#file-agent-acp_agent.go) (1266 lines, 31.82 KB)
- [agent/agent.go](#file-agent-agent.go) (108 lines, 3.17 KB)
- [agent/cli_agent.go](#file-agent-cli_agent.go) (298 lines, 8.50 KB)
- [agent/env_test.go](#file-agent-env_test.go) (62 lines, 1.50 KB)
- [agent/http_agent.go](#file-agent-http_agent.go) (188 lines, 4.96 KB)
- [api/server.go](#file-api-server.go) (119 lines, 3.14 KB)
- [cmd/login.go](#file-cmd-login.go) (30 lines, 0.56 KB)
- [cmd/proc_unix.go](#file-cmd-proc_unix.go) (12 lines, 0.16 KB)
- [cmd/proc_windows.go](#file-cmd-proc_windows.go) (9 lines, 0.15 KB)
- [cmd/restart.go](#file-cmd-restart.go) (40 lines, 0.72 KB)
- [cmd/root.go](#file-cmd-root.go) (27 lines, 0.50 KB)
- [cmd/send.go](#file-cmd-send.go) (68 lines, 1.84 KB)
- [cmd/start.go](#file-cmd-start.go) (435 lines, 11.48 KB)
- [cmd/status.go](#file-cmd-status.go) (31 lines, 0.56 KB)
- [cmd/stop.go](#file-cmd-stop.go) (21 lines, 0.31 KB)
- [cmd/update.go](#file-cmd-update.go) (207 lines, 4.63 KB)
- [config/config.go](#file-config-config.go) (141 lines, 4.21 KB)
- [config/config_test.go](#file-config-config_test.go) (119 lines, 2.53 KB)
- [config/detect.go](#file-config-detect.go) (281 lines, 9.21 KB)
- [config/detect_test.go](#file-config-detect_test.go) (82 lines, 2.50 KB)
- [docs/README_CN.md](#file-docs-readme_cn.md) (345 lines, 9.32 KB)
- [docs/项目学习.md](#file-docs-项目学习.md) (663 lines, 18.94 KB)
- [go.mod](#file-go.mod) (15 lines, 0.43 KB)
- [ilink/auth.go](#file-ilink-auth.go) (177 lines, 3.96 KB)
- [ilink/client.go](#file-ilink-client.go) (224 lines, 5.66 KB)
- [ilink/monitor.go](#file-ilink-monitor.go) (181 lines, 4.60 KB)
- [ilink/types.go](#file-ilink-types.go) (219 lines, 6.62 KB)
- [install.sh](#file-install.sh) (64 lines, 1.60 KB)
- [main.go](#file-main.go) (7 lines, 0.09 KB)
- [messaging/attachment.go](#file-messaging-attachment.go) (127 lines, 2.90 KB)
- [messaging/attachment_test.go](#file-messaging-attachment_test.go) (100 lines, 2.96 KB)
- [messaging/cdn.go](#file-messaging-cdn.go) (232 lines, 6.56 KB)
- [messaging/handler.go](#file-messaging-handler.go) (1066 lines, 32.49 KB)
- [messaging/handler_test.go](#file-messaging-handler_test.go) (140 lines, 3.60 KB)
- [messaging/linkhoard.go](#file-messaging-linkhoard.go) (325 lines, 8.58 KB)
- [messaging/markdown.go](#file-messaging-markdown.go) (103 lines, 3.01 KB)
- [messaging/media.go](#file-messaging-media.go) (213 lines, 5.31 KB)
- [messaging/media_test.go](#file-messaging-media_test.go) (73 lines, 1.81 KB)
- [messaging/sender.go](#file-messaging-sender.go) (86 lines, 2.21 KB)
- [service/com.fastclaw.weclaw.plist](#file-service-com.fastclaw.weclaw.plist) (21 lines, 0.58 KB)
- [service/weclaw.service](#file-service-weclaw.service) (14 lines, 0.24 KB)

---

<a name="file-.air.toml"></a>
## 📄 .air.toml

````toml
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = ["start", "-f"]
  bin = "./weclaw"
  cmd = "go build -o ./weclaw ."
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata", "debug"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html"]
  include_file = []
  kill_delay = "0s"
  log = "build-errors.log"
  poll = false
  poll_interval = 0
  post_cmd = []
  pre_cmd = []
  rerun = false
  rerun_delay = 500
  send_interrupt = false
  stop_on_error = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  main_only = false
  silent = false
  time = false

[misc]
  clean_on_exit = false

[proxy]
  app_port = 0
  enabled = false
  proxy_port = 0

[screen]
  clear_on_rebuild = false
  keep_scroll = true

````

[⬆ 回到目录](#toc)

<a name="file-.dockerignore"></a>
## 📄 .dockerignore

````text
weclaw
tmp/
.git/
.idea/
.vscode/
.claude/
.env
*.local
.DS_Store
Thumbs.db
*.swp
*.swo
*~

````

[⬆ 回到目录](#toc)

<a name="file-.gitignore"></a>
## 📄 .gitignore

````text
# Binary
weclaw

# Air hot reload
tmp/

# IDE
.idea/
.vscode/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db

# Environment & config
.env
*.local

# Claude Code
.claude/

````

[⬆ 回到目录](#toc)

<a name="file-dockerfile"></a>
## 📄 Dockerfile

````dockerfile
FROM golang:1.25-alpine AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /usr/local/bin/weclaw .

FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder /usr/local/bin/weclaw /usr/local/bin/weclaw

VOLUME /root/.weclaw
ENTRYPOINT ["weclaw"]
CMD ["start"]

````

[⬆ 回到目录](#toc)

<a name="file-license"></a>
## 📄 LICENSE

````text
MIT License

Copyright (c) 2026 fastclaw-ai

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

````

[⬆ 回到目录](#toc)

<a name="file-makefile"></a>
## 📄 Makefile

````makefile
dev:
	air -c .air.toml start
````

[⬆ 回到目录](#toc)

<a name="file-readme.md"></a>
## 📄 README.md

````markdown
# WeClaw

[中文文档](README_CN.md)

WeChat AI Agent Bridge — connect WeChat to AI agents (Claude, Codex, Gemini, Kimi, etc.).

> This project is inspired by [@tencent-weixin/openclaw-weixin](https://npmx.dev/package/@tencent-weixin/openclaw-weixin). For personal learning only, not for commercial use.

| | | |
|:---:|:---:|:---:|
| <img src="previews/preview1.png" width="280" /> | <img src="previews/preview2.png" width="280" /> | <img src="previews/preview3.png" width="280" /> |

## Quick Start

```bash
# One-line install
curl -sSL https://raw.githubusercontent.com/fastclaw-ai/weclaw/main/install.sh | sh

# Start (first run will prompt QR code login)
weclaw start
```

That's it. On first start, WeClaw will:
1. Show a QR code — scan with WeChat to login
2. Auto-detect installed AI agents (Claude, Codex, Gemini, etc.)
3. Save config to `~/.weclaw/config.json`
4. Start receiving and replying to WeChat messages

Use `weclaw login` to add additional WeChat accounts.

### Other install methods

```bash
# Via Go
go install github.com/fastclaw-ai/weclaw@latest

# Via Docker
docker run -it -v ~/.weclaw:/root/.weclaw ghcr.io/fastclaw-ai/weclaw start
```

## How It Works

<p align="center">
  <img src="previews/architecture.png" width="600" />
</p>

**Agent modes:**

| Mode | How it works | Examples |
|------|-------------|----------|
| ACP  | Long-running subprocess, JSON-RPC over stdio. Fastest — reuses process and sessions. | Claude, Codex, Kimi, Gemini, Cursor, OpenCode, OpenClaw |
| CLI  | Spawns a new process per message. Supports session resume via `--resume`. | Claude (`claude -p`), Codex (`codex exec`) |
| HTTP | OpenAI-compatible chat completions API. | OpenClaw (HTTP fallback) |

Auto-detection picks ACP over CLI when both are available.

## Chat Commands

Send these as WeChat messages:

| Command | Description |
|---------|-------------|
| `hello` | Send to default agent |
| `/codex write a function` | Send to a specific agent |
| `/cc explain this code` | Send to agent by alias |
| `/claude` | Switch default agent to Claude |
| `/cwd /path/to/project` | Switch workspace directory |
| `/new` | Start a new conversation (clear session) |
| `/info` | Show current agent info |
| `/help` | Show help message |

### Aliases

| Alias | Agent |
|-------|-------|
| `/cc` | claude |
| `/cx` | codex |
| `/cs` | cursor |
| `/km` | kimi |
| `/gm` | gemini |
| `/ocd` | opencode |
| `/oc` | openclaw |

You can also define custom aliases per agent in config:

```json
{
  "agents": {
    "claude": {
      "type": "acp",
      "aliases": ["ai", "c"]
    }
  }
}
```

Then `/ai hello` or `/c hello` will route to claude.

Switching default agent is persisted to config — survives restarts.

## Media Messages

WeClaw supports sending images, videos, files, and voice messages to/from WeChat.

**Voice messages:** When you send a voice message in WeChat, WeClaw automatically uses WeChat's speech-to-text transcription and forwards the text to the AI agent. Duplicate voice message events are automatically deduplicated.

**From agent replies:** When an AI agent returns markdown with images (`![](url)`), WeClaw automatically extracts the image URLs, downloads them, uploads to WeChat CDN (AES-128-ECB encrypted), and sends them as image messages.

**Markdown handling:** Agent responses are automatically converted from markdown to plain text for WeChat display — code fences are stripped, links show display text only, bold/italic markers are removed, etc.

## Proactive Messaging

Send messages to WeChat users without waiting for them to message first.

**CLI:**

```bash
# Send text
weclaw send --to "user_id@im.wechat" --text "Hello from weclaw"

# Send image
weclaw send --to "user_id@im.wechat" --media "https://example.com/photo.png"

# Send text + image
weclaw send --to "user_id@im.wechat" --text "Check this out" --media "https://example.com/photo.png"

# Send file
weclaw send --to "user_id@im.wechat" --media "https://example.com/report.pdf"
```

**HTTP API** (runs on `127.0.0.1:18011` when `weclaw start` is running):

```bash
# Send text
curl -X POST http://127.0.0.1:18011/api/send \
  -H "Content-Type: application/json" \
  -d '{"to": "user_id@im.wechat", "text": "Hello from weclaw"}'

# Send image
curl -X POST http://127.0.0.1:18011/api/send \
  -H "Content-Type: application/json" \
  -d '{"to": "user_id@im.wechat", "media_url": "https://example.com/photo.png"}'

# Send text + media
curl -X POST http://127.0.0.1:18011/api/send \
  -H "Content-Type: application/json" \
  -d '{"to": "user_id@im.wechat", "text": "See this", "media_url": "https://example.com/photo.png"}'
```

Supported media types: images (png, jpg, gif, webp), videos (mp4, mov), files (pdf, doc, zip, etc.).

Set `WECLAW_API_ADDR` to change the listen address (e.g. `0.0.0.0:18011`).

## Configuration

Config file: `~/.weclaw/config.json`

```json
{
  "default_agent": "claude",
  "agents": {
    "claude": {
      "type": "acp",
      "command": "/usr/local/bin/claude-agent-acp",
      "env": {
        "ANTHROPIC_API_KEY": "sk-ant-xxx"
      },
      "model": "sonnet"
    },
    "codex": {
      "type": "acp",
      "command": "/usr/local/bin/codex-acp",
      "env": {
        "OPENAI_API_KEY": "sk-xxx"
      }
    },
    "openclaw": {
      "type": "http",
      "endpoint": "https://api.example.com/v1/chat/completions",
      "api_key": "sk-xxx",
      "model": "openclaw:main"
    }
  }
}
```

Environment variables:
- `WECLAW_DEFAULT_AGENT` — override default agent
- `OPENCLAW_GATEWAY_URL` — OpenClaw HTTP fallback endpoint
- `OPENCLAW_GATEWAY_TOKEN` — OpenClaw API token

Custom agent CLI environment variables:

```json
{
  "default_agent": "...",
  "agents": {
    "...": {
      ...
      "env": {
        "ENV_NAME": "ENV_VALUE"
      }
    },
  }
}
```

### Permission bypass

By default, some agents require interactive permission approval which doesn't work in WeChat. Add `args` to your agent config to bypass:

| Agent | Flag | What it does |
|-------|------|-------------|
| Claude (CLI) | `--dangerously-skip-permissions` | Skip all tool permission prompts |
| Codex (CLI) | `--skip-git-repo-check` | Allow running outside git repos |

Example:

```json
{
  "claude": {
    "type": "cli",
    "command": "/usr/local/bin/claude",
    "cwd": "/home/user/my-project",
    "args": ["--dangerously-skip-permissions"]
  },
  "codex": {
    "type": "cli",
    "command": "/usr/local/bin/codex",
    "cwd": "/home/user/my-project",
    "args": ["--skip-git-repo-check"]
  }
}
```

Set `cwd` to specify the agent's working directory (workspace). If omitted, defaults to `~/.weclaw/workspace`.

> **Warning:** These flags disable safety checks. Only enable them if you understand the risks. ACP agents handle permissions automatically and don't need these flags.

## Background Mode

```bash
# Start (runs in background by default)
weclaw start

# Check if running
weclaw status

# Stop
weclaw stop

# Run in foreground (for debugging)
weclaw start -f
```

Logs are written to `~/.weclaw/weclaw.log`.

### System service (auto-start on boot)

**macOS (launchd):**

```bash
cp service/com.fastclaw.weclaw.plist ~/Library/LaunchAgents/
launchctl load ~/Library/LaunchAgents/com.fastclaw.weclaw.plist
```

**Linux (systemd):**

```bash
sudo cp service/weclaw.service /etc/systemd/system/
sudo systemctl enable --now weclaw
```

## Docker

```bash
# Build
docker build -t weclaw .

# Login (interactive — scan QR code)
docker run -it -v ~/.weclaw:/root/.weclaw weclaw login

# Start with HTTP agent
docker run -d --name weclaw \
  -v ~/.weclaw:/root/.weclaw \
  -e OPENCLAW_GATEWAY_URL=https://api.example.com \
  -e OPENCLAW_GATEWAY_TOKEN=sk-xxx \
  weclaw

# View logs
docker logs -f weclaw
```

> Note: ACP and CLI agents require the agent binary inside the container.
> The Docker image ships only WeClaw itself. For ACP/CLI agents, mount
> the binary or build a custom image. HTTP agents work out of the box.

## Release

```bash
# Tag a new version to trigger GitHub Actions build & release
git tag v0.1.0
git push origin v0.1.0
```

The workflow builds binaries for `darwin/linux/windows` x `amd64/arm64`, creates a GitHub Release, and uploads all artifacts with checksums.

## Update

```bash
# Update to the latest version (auto-restarts if running)
weclaw update

# Check current version
weclaw version
```

## Development

```bash
# Hot reload
make dev

# Build
go build -o weclaw .

# Run
./weclaw start
```

## Contributors

<a href="https://github.com/fastclaw-ai/weclaw/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=fastclaw-ai/weclaw" />
</a>

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=fastclaw-ai/weclaw&type=Timeline)](https://star-history.com/#fastclaw-ai/weclaw&Timeline)

## License

[MIT](LICENSE)

````

[⬆ 回到目录](#toc)

<a name="file-agent-acp_agent.go"></a>
## 📄 agent/acp_agent.go

````go
package agent

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// ACPAgent communicates with ACP-compatible agents (claude-agent-acp, codex-acp, cursor agent, etc.) via stdio JSON-RPC 2.0.
type ACPAgent struct {
	command      string
	args         []string
	model        string
	systemPrompt string
	cwd          string
	env          map[string]string
	protocol     string // "legacy_acp" or "codex_app_server"

	mu       sync.Mutex
	cmd      *exec.Cmd
	stdin    io.WriteCloser
	scanner  *bufio.Scanner
	started  bool
	nextID   atomic.Int64
	sessions map[string]string // conversationID -> sessionID (legacy ACP)
	threads  map[string]string // conversationID -> threadID (codex app-server)

	// pending tracks in-flight JSON-RPC requests
	pendingMu sync.Mutex
	pending   map[int64]chan *rpcResponse

	// notifications channel for session/update events
	notifyMu sync.Mutex
	notifyCh map[string]chan *sessionUpdate // sessionID -> channel
	turnCh   map[string]chan *codexTurnEvent

	stderr *acpStderrWriter // captures stderr for error reporting
}

// ACPAgentConfig holds configuration for the ACP agent.
type ACPAgentConfig struct {
	Command      string   // path to ACP agent binary (claude-agent-acp, codex-acp, cursor agent, etc.)
	Args         []string // extra args for command (e.g. ["acp"] for cursor)
	Model        string
	SystemPrompt string
	Cwd          string            // working directory
	Env          map[string]string // extra environment variables
}

// --- JSON-RPC types ---

type rpcRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      int64       `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type rpcResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      *int64          `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *rpcError       `json:"error,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// --- ACP protocol types ---

type initParams struct {
	ProtocolVersion    int                `json:"protocolVersion"`
	ClientCapabilities clientCapabilities `json:"clientCapabilities"`
}

type clientCapabilities struct {
	FS *fsCapabilities `json:"fs,omitempty"`
}

type fsCapabilities struct {
	ReadTextFile  bool `json:"readTextFile"`
	WriteTextFile bool `json:"writeTextFile"`
}

type newSessionParams struct {
	Cwd        string        `json:"cwd"`
	McpServers []interface{} `json:"mcpServers"`
}

type newSessionResult struct {
	SessionID string `json:"sessionId"`
}

type promptParams struct {
	SessionID string        `json:"sessionId"`
	Prompt    []promptEntry `json:"prompt"`
}

type promptEntry struct {
	Type     string `json:"type"`
	Text     string `json:"text,omitempty"`
	URL      string `json:"url,omitempty"`
	Path     string `json:"path,omitempty"`
	Data     string `json:"data,omitempty"`
	MimeType string `json:"mimeType,omitempty"`
}

type promptResult struct {
	StopReason string `json:"stopReason"`
}

type sessionUpdateParams struct {
	SessionID string        `json:"sessionId"`
	Update    sessionUpdate `json:"update"`
}

type sessionUpdate struct {
	SessionUpdate string          `json:"sessionUpdate"`
	Content       json.RawMessage `json:"content,omitempty"`
	// For agent_message_chunk
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

type permissionRequestParams struct {
	ToolCall json.RawMessage    `json:"toolCall"`
	Options  []permissionOption `json:"options"`
}

type permissionOption struct {
	OptionID string `json:"optionId"`
	Name     string `json:"name"`
	Kind     string `json:"kind"`
}

// Codex app-server protocol constants and types.
const (
	protocolLegacyACP      = "legacy_acp"
	protocolCodexAppServer = "codex_app_server"
)

type codexTurnStartParams struct {
	ThreadID       string           `json:"threadId"`
	ApprovalPolicy string           `json:"approvalPolicy,omitempty"`
	Input          []codexUserInput `json:"input"`
	SandboxPolicy  interface{}      `json:"sandboxPolicy,omitempty"`
	Model          string           `json:"model,omitempty"`
	Cwd            string           `json:"cwd,omitempty"`
}

type codexUserInput struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

type codexTurnEvent struct {
	Kind  string
	Delta string
	Text  string
}

func detectACPProtocol(command string, args []string) string {
	base := strings.ToLower(filepath.Base(command))
	// codex-acp is a standard ACP wrapper, NOT codex app-server
	// Only `codex app-server` uses the codex-native protocol
	if base == "codex" || base == "codex.exe" {
		for _, arg := range args {
			if arg == "app-server" {
				return protocolCodexAppServer
			}
		}
	}
	return protocolLegacyACP
}

// NewACPAgent creates a new ACP agent.
func NewACPAgent(cfg ACPAgentConfig) *ACPAgent {
	if cfg.Command == "" {
		cfg.Command = "claude-agent-acp"
	}
	if cfg.Cwd == "" {
		cfg.Cwd = defaultWorkspace()
	}
	protocol := detectACPProtocol(cfg.Command, cfg.Args)
	return &ACPAgent{
		command:      cfg.Command,
		args:         cfg.Args,
		model:        cfg.Model,
		systemPrompt: cfg.SystemPrompt,
		cwd:          cfg.Cwd,
		env:          cfg.Env,
		protocol:     protocol,
		sessions:     make(map[string]string),
		threads:      make(map[string]string),
		pending:      make(map[int64]chan *rpcResponse),
		notifyCh:     make(map[string]chan *sessionUpdate),
		turnCh:       make(map[string]chan *codexTurnEvent),
	}
}

// Start launches the claude-agent-acp subprocess and initializes the connection.
func (a *ACPAgent) Start(ctx context.Context) error {
	a.mu.Lock()
	if a.started {
		a.mu.Unlock()
		return nil
	}

	a.cmd = exec.CommandContext(ctx, a.command, a.args...)
	a.cmd.Dir = a.cwd
	if len(a.env) > 0 {
		cmdEnv, err := mergeEnv(os.Environ(), a.env)
		if err != nil {
			a.mu.Unlock()
			return fmt.Errorf("build acp env: %w", err)
		}
		a.cmd.Env = cmdEnv
	}
	// Capture stderr for debugging and error reporting
	a.stderr = &acpStderrWriter{prefix: "[acp-stderr]"}
	a.cmd.Stderr = a.stderr

	var err error
	a.stdin, err = a.cmd.StdinPipe()
	if err != nil {
		a.mu.Unlock()
		return fmt.Errorf("create stdin pipe: %w", err)
	}

	stdout, err := a.cmd.StdoutPipe()
	if err != nil {
		a.mu.Unlock()
		return fmt.Errorf("create stdout pipe: %w", err)
	}

	if err := a.cmd.Start(); err != nil {
		a.mu.Unlock()
		return fmt.Errorf("start acp agent %s: %w", a.command, err)
	}

	pid := a.cmd.Process.Pid
	log.Printf("[acp] started subprocess (command=%s, pid=%d)", a.command, pid)

	a.scanner = bufio.NewScanner(stdout)
	a.scanner.Buffer(make([]byte, 0, 4*1024*1024), 4*1024*1024) // 4MB
	a.started = true

	// Start reading loop
	go a.readLoop()

	// Release lock before calling initialize — call() needs a.mu to write to stdin
	a.mu.Unlock()

	// Initialize handshake with timeout
	initCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	log.Printf("[acp] sending initialize handshake (pid=%d, protocol=%s)...", pid, a.protocol)
	var result json.RawMessage
	if a.protocol == protocolCodexAppServer {
		result, err = a.call(initCtx, "initialize", map[string]interface{}{
			"clientInfo": map[string]string{"name": "weclaw", "version": "0.3.0"},
		})
		if err == nil {
			// codex app-server expects an "initialized" notification after initialize response
			err = a.notify("initialized", nil)
		}
	} else {
		result, err = a.call(initCtx, "initialize", initParams{
			ProtocolVersion: 1,
			ClientCapabilities: clientCapabilities{
				FS: &fsCapabilities{ReadTextFile: true, WriteTextFile: true},
			},
		})
	}
	if err != nil {
		a.mu.Lock()
		a.started = false
		a.mu.Unlock()
		a.stdin.Close()
		a.cmd.Process.Kill()
		a.cmd.Wait()
		// Use stderr detail if available (e.g. "connect ECONNREFUSED")
		if detail := a.stderr.LastError(); detail != "" {
			return fmt.Errorf("agent startup failed: %s", detail)
		}
		// Provide a helpful hint when the binary looks like a Claude CLI that doesn't support ACP
		base := strings.ToLower(filepath.Base(a.command))
		if base == "claude" || base == "claude.exe" {
			return fmt.Errorf("agent startup failed (pid=%d): %w\n\nHint: the 'claude' CLI does not support ACP protocol directly.\nSet type to \"cli\" in your config, or install claude-agent-acp and set command to \"claude-agent-acp\".", pid, err)
		}
		return fmt.Errorf("agent startup failed (pid=%d): %w", pid, err)
	}

	log.Printf("[acp] initialized (pid=%d): %s", pid, string(result))
	return nil
}

// Stop terminates the subprocess.
func (a *ACPAgent) Stop() {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.started {
		return
	}
	a.stdin.Close()
	a.cmd.Process.Kill()
	a.cmd.Wait()
	a.started = false
}

// SetCwd changes the working directory for subsequent sessions.
func (a *ACPAgent) SetCwd(cwd string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.cwd = cwd
}

// ResetSession clears the existing session for the given conversationID and
// immediately creates a new one, returning the new session ID.
func (a *ACPAgent) ResetSession(ctx context.Context, conversationID string) (string, error) {
	a.mu.Lock()
	delete(a.sessions, conversationID)
	a.mu.Unlock()
	log.Printf("[acp] session reset (conversation=%s), creating new session", conversationID)

	sessionID, _, err := a.getOrCreateSession(ctx, conversationID)
	if err != nil {
		return "", fmt.Errorf("create new session: %w", err)
	}
	return sessionID, nil
}

// Chat sends a message and returns the full response.
func (a *ACPAgent) Chat(ctx context.Context, conversationID string, message string) (string, error) {
	if !a.started {
		if err := a.Start(ctx); err != nil {
			return "", err
		}
	}

	// Route to codex app-server protocol if applicable
	if a.protocol == protocolCodexAppServer {
		return a.chatCodexAppServer(ctx, conversationID, message)
	}

	// Get or create session
	sessionID, isNew, err := a.getOrCreateSession(ctx, conversationID)
	if err != nil {
		return "", fmt.Errorf("session error: %w", err)
	}

	pid := a.cmd.Process.Pid
	if isNew {
		log.Printf("[acp] new session created (pid=%d, session=%s, conversation=%s)", pid, sessionID, conversationID)
	} else {
		log.Printf("[acp] reusing session (pid=%d, session=%s, conversation=%s)", pid, sessionID, conversationID)
	}

	// Register notification channel for this session
	notifyCh := make(chan *sessionUpdate, 256)
	a.notifyMu.Lock()
	a.notifyCh[sessionID] = notifyCh
	a.notifyMu.Unlock()

	defer func() {
		a.notifyMu.Lock()
		delete(a.notifyCh, sessionID)
		a.notifyMu.Unlock()
	}()

	// Send prompt (this blocks until the prompt completes)
	type promptDoneMsg struct {
		result json.RawMessage
		err    error
	}
	promptDone := make(chan promptDoneMsg, 1)
	go func() {
		result, err := a.call(ctx, "session/prompt", promptParams{
			SessionID: sessionID,
			Prompt:    []promptEntry{{Type: "text", Text: message}},
		})
		if result != nil {
			log.Printf("[acp] prompt result (session=%s): %s", sessionID, string(result))
		}
		promptDone <- promptDoneMsg{result: result, err: err}
	}()

	// Collect text chunks from notifications
	var textParts []string

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case update := <-notifyCh:
			if update.SessionUpdate == "agent_message_chunk" {
				text := extractChunkText(update)
				if text != "" {
					textParts = append(textParts, text)
				}
			}
		case done := <-promptDone:
			// Drain remaining notifications
			for {
				select {
				case update := <-notifyCh:
					if update.SessionUpdate == "agent_message_chunk" {
						text := extractChunkText(update)
						if text != "" {
							textParts = append(textParts, text)
						}
					}
				default:
					goto drained
				}
			}
		drained:
			if done.err != nil {
				return "", fmt.Errorf("prompt error: %w", done.err)
			}
			result := strings.TrimSpace(strings.Join(textParts, ""))
			if result == "" {
				// Try extracting from prompt result (some agents return content here)
				result = extractPromptResultText(done.result)
			}
			if result == "" {
				return "", fmt.Errorf("agent returned empty response")
			}
			return result, nil
		}
	}
}

// ChatWithMedia sends a message with media attachments and returns the full response.
func (a *ACPAgent) ChatWithMedia(ctx context.Context, conversationID string, message string, media []MediaEntry) (string, error) {
	if !a.started {
		if err := a.Start(ctx); err != nil {
			return "", err
		}
	}

	// Route to codex app-server protocol if applicable
	if a.protocol == protocolCodexAppServer {
		return a.chatCodexAppServerWithMedia(ctx, conversationID, message, media)
	}

	// Get or create session
	sessionID, isNew, err := a.getOrCreateSession(ctx, conversationID)
	if err != nil {
		return "", fmt.Errorf("session error: %w", err)
	}

	pid := a.cmd.Process.Pid
	if isNew {
		log.Printf("[acp] new session created (pid=%d, session=%s, conversation=%s)", pid, sessionID, conversationID)
	} else {
		log.Printf("[acp] reusing session (pid=%d, session=%s, conversation=%s)", pid, sessionID, conversationID)
	}

	// Register notification channel for this session
	notifyCh := make(chan *sessionUpdate, 256)
	a.notifyMu.Lock()
	a.notifyCh[sessionID] = notifyCh
	a.notifyMu.Unlock()

	defer func() {
		a.notifyMu.Lock()
		delete(a.notifyCh, sessionID)
		a.notifyMu.Unlock()
	}()

	// Build prompt entries with media
	prompt := buildPromptEntries(message, media)

	// Send prompt (this blocks until the prompt completes)
	type promptDoneMsg struct {
		result json.RawMessage
		err    error
	}
	promptDone := make(chan promptDoneMsg, 1)
	go func() {
		result, err := a.call(ctx, "session/prompt", promptParams{
			SessionID: sessionID,
			Prompt:    prompt,
		})
		if result != nil {
			log.Printf("[acp] prompt result (session=%s): %s", sessionID, string(result))
		}
		promptDone <- promptDoneMsg{result: result, err: err}
	}()

	// Collect text chunks from notifications
	var textParts []string

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case update := <-notifyCh:
			if update.SessionUpdate == "agent_message_chunk" {
				text := extractChunkText(update)
				if text != "" {
					textParts = append(textParts, text)
				}
			}
		case done := <-promptDone:
			// Drain remaining notifications
			for {
				select {
				case update := <-notifyCh:
					if update.SessionUpdate == "agent_message_chunk" {
						text := extractChunkText(update)
						if text != "" {
							textParts = append(textParts, text)
						}
					}
				default:
					goto drainedMedia
				}
			}
		drainedMedia:
			if done.err != nil {
				return "", fmt.Errorf("prompt error: %w", done.err)
			}
			result := strings.TrimSpace(strings.Join(textParts, ""))
			if result == "" {
				// Try extracting from prompt result (some agents return content here)
				result = extractPromptResultText(done.result)
			}
			if result == "" {
				return "", fmt.Errorf("agent returned empty response")
			}
			return result, nil
		}
	}
}

// buildPromptEntries builds prompt entries from message and media.
func buildPromptEntries(message string, media []MediaEntry) []promptEntry {
	var entries []promptEntry

	// Add media entries first
	for _, m := range media {
		entry := promptEntry{Type: m.Type}
		switch m.Type {
		case "image":
			if m.URL != "" {
				entry.URL = m.URL
			} else if m.Path != "" {
				entry.Path = m.Path
			}
		case "file":
			entry.Type = "file"
			if m.URL != "" {
				entry.URL = m.URL
			} else if m.Path != "" {
				entry.Path = m.Path
			}
			entry.MimeType = m.MIMEType
		case "video":
			entry.Type = "video"
			if m.URL != "" {
				entry.URL = m.URL
			} else if m.Path != "" {
				entry.Path = m.Path
			}
		}
		entries = append(entries, entry)
	}

	// Add text entry
	if message != "" {
		entries = append(entries, promptEntry{Type: "text", Text: message})
	}

	return entries
}

// chatCodexAppServerWithMedia handles media for codex app-server protocol.
func (a *ACPAgent) chatCodexAppServerWithMedia(ctx context.Context, conversationID string, message string, media []MediaEntry) (string, error) {
	threadID, isNew, err := a.getOrCreateThread(ctx, conversationID)
	if err != nil {
		return "", fmt.Errorf("thread error: %w", err)
	}

	pid := 0
	a.mu.Lock()
	if a.cmd != nil && a.cmd.Process != nil {
		pid = a.cmd.Process.Pid
	}
	a.mu.Unlock()

	if isNew {
		log.Printf("[acp] new thread created (pid=%d, thread=%s, conversation=%s)", pid, threadID, conversationID)
	} else {
		log.Printf("[acp] reusing thread (pid=%d, thread=%s, conversation=%s)", pid, threadID, conversationID)
	}

	// Register turn event channel
	turnCh := make(chan *codexTurnEvent, 256)
	a.notifyMu.Lock()
	a.turnCh[threadID] = turnCh
	a.notifyMu.Unlock()

	defer func() {
		a.notifyMu.Lock()
		delete(a.turnCh, threadID)
		a.notifyMu.Unlock()
	}()

	// Build input entries
	var input []codexUserInput
	for _, m := range media {
		input = append(input, codexUserInput{Type: m.Type, Text: m.URL})
	}
	if message != "" {
		input = append(input, codexUserInput{Type: "text", Text: message})
	}

	// Start turn (call returns quickly with turn info, actual content comes via events)
	go func() {
		_, err := a.call(ctx, "turn/start", codexTurnStartParams{
			ThreadID:       threadID,
			ApprovalPolicy: "never",
			Input:          input,
			SandboxPolicy:  map[string]interface{}{"type": "dangerFullAccess"},
			Model:          a.model,
			Cwd:            a.cwd,
		})
		if err != nil {
			// If call itself fails, signal via turn channel
			turnCh <- &codexTurnEvent{Kind: "error", Text: err.Error()}
		}
	}()

	// Collect text from events until turn/completed
	var textParts []string
	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case evt := <-turnCh:
			if evt.Kind == "error" {
				return "", fmt.Errorf("turn error: %s", evt.Text)
			}
			if evt.Delta != "" {
				textParts = append(textParts, evt.Delta)
			}
			if evt.Text != "" {
				textParts = append(textParts, evt.Text)
			}
			if evt.Kind == "completed" {
				result := strings.TrimSpace(strings.Join(textParts, ""))
				if result == "" {
					return "", fmt.Errorf("agent returned empty response")
				}
				return result, nil
			}
		}
	}
}

func (a *ACPAgent) getOrCreateSession(ctx context.Context, conversationID string) (string, bool, error) {
	a.mu.Lock()
	sid, exists := a.sessions[conversationID]
	a.mu.Unlock()

	if exists {
		return sid, false, nil
	}

	result, err := a.call(ctx, "session/new", newSessionParams{
		Cwd:        a.cwd,
		McpServers: []interface{}{},
	})
	if err != nil {
		return "", false, err
	}

	var sessionResult newSessionResult
	if err := json.Unmarshal(result, &sessionResult); err != nil {
		return "", false, fmt.Errorf("parse session result: %w", err)
	}

	a.mu.Lock()
	a.sessions[conversationID] = sessionResult.SessionID
	a.mu.Unlock()

	return sessionResult.SessionID, true, nil
}

// --- Codex app-server protocol ---

func (a *ACPAgent) getOrCreateThread(ctx context.Context, conversationID string) (string, bool, error) {
	a.mu.Lock()
	tid, exists := a.threads[conversationID]
	a.mu.Unlock()

	if exists {
		return tid, false, nil
	}

	params := map[string]interface{}{
		"approvalPolicy": "never",
		"cwd":            a.cwd,
		"sandbox":        "danger-full-access",
	}
	if a.model != "" {
		params["model"] = a.model
	}
	result, err := a.call(ctx, "thread/start", params)
	if err != nil {
		return "", false, err
	}

	var threadResult struct {
		Thread struct {
			ID string `json:"id"`
		} `json:"thread"`
	}
	if err := json.Unmarshal(result, &threadResult); err != nil {
		return "", false, fmt.Errorf("parse thread/start result: %w", err)
	}
	if threadResult.Thread.ID == "" {
		return "", false, fmt.Errorf("thread/start returned empty thread id")
	}

	a.mu.Lock()
	a.threads[conversationID] = threadResult.Thread.ID
	a.mu.Unlock()

	return threadResult.Thread.ID, true, nil
}

func (a *ACPAgent) chatCodexAppServer(ctx context.Context, conversationID string, message string) (string, error) {
	threadID, isNew, err := a.getOrCreateThread(ctx, conversationID)
	if err != nil {
		return "", fmt.Errorf("thread error: %w", err)
	}

	pid := 0
	a.mu.Lock()
	if a.cmd != nil && a.cmd.Process != nil {
		pid = a.cmd.Process.Pid
	}
	a.mu.Unlock()

	if isNew {
		log.Printf("[acp] new thread created (pid=%d, thread=%s, conversation=%s)", pid, threadID, conversationID)
	} else {
		log.Printf("[acp] reusing thread (pid=%d, thread=%s, conversation=%s)", pid, threadID, conversationID)
	}

	// Register turn event channel
	turnCh := make(chan *codexTurnEvent, 256)
	a.notifyMu.Lock()
	a.turnCh[threadID] = turnCh
	a.notifyMu.Unlock()

	defer func() {
		a.notifyMu.Lock()
		delete(a.turnCh, threadID)
		a.notifyMu.Unlock()
	}()

	// Start turn (call returns quickly with turn info, actual content comes via events)
	go func() {
		_, err := a.call(ctx, "turn/start", codexTurnStartParams{
			ThreadID:       threadID,
			ApprovalPolicy: "never",
			Input:          []codexUserInput{{Type: "text", Text: message}},
			SandboxPolicy:  map[string]interface{}{"type": "dangerFullAccess"},
			Model:          a.model,
			Cwd:            a.cwd,
		})
		if err != nil {
			// If call itself fails, signal via turn channel
			turnCh <- &codexTurnEvent{Kind: "error", Text: err.Error()}
		}
	}()

	// Collect text from events until turn/completed
	var textParts []string
	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case evt := <-turnCh:
			if evt.Kind == "error" {
				return "", fmt.Errorf("turn error: %s", evt.Text)
			}
			if evt.Delta != "" {
				textParts = append(textParts, evt.Delta)
			}
			if evt.Text != "" {
				textParts = append(textParts, evt.Text)
			}
			if evt.Kind == "completed" {
				result := strings.TrimSpace(strings.Join(textParts, ""))
				if result == "" {
					return "", fmt.Errorf("agent returned empty response")
				}
				return result, nil
			}
		}
	}
}

// notify sends a JSON-RPC notification (no id, no response expected).
func (a *ACPAgent) notify(method string, params interface{}) error {
	msg := struct {
		JSONRPC string      `json:"jsonrpc"`
		Method  string      `json:"method"`
		Params  interface{} `json:"params,omitempty"`
	}{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal notification: %w", err)
	}

	a.mu.Lock()
	_, err = fmt.Fprintf(a.stdin, "%s\n", data)
	a.mu.Unlock()
	return err
}

// call sends a JSON-RPC request and waits for the response.
func (a *ACPAgent) call(ctx context.Context, method string, params interface{}) (json.RawMessage, error) {
	id := a.nextID.Add(1)

	ch := make(chan *rpcResponse, 1)
	a.pendingMu.Lock()
	a.pending[id] = ch
	a.pendingMu.Unlock()

	defer func() {
		a.pendingMu.Lock()
		delete(a.pending, id)
		a.pendingMu.Unlock()
	}()

	req := rpcRequest{
		JSONRPC: "2.0",
		ID:      id,
		Method:  method,
		Params:  params,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	a.mu.Lock()
	_, err = fmt.Fprintf(a.stdin, "%s\n", data)
	a.mu.Unlock()
	if err != nil {
		return nil, fmt.Errorf("write to stdin: %w", err)
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case resp := <-ch:
		if resp.Error != nil {
			msg := resp.Error.Message
			// Enrich with stderr context if available
			if a.stderr != nil {
				if detail := a.stderr.LastError(); detail != "" {
					msg = detail
				}
			}
			return nil, fmt.Errorf("agent error: %s", msg)
		}
		return resp.Result, nil
	}
}

// readLoop reads NDJSON lines from stdout and dispatches to pending requests or notification channels.
func (a *ACPAgent) readLoop() {
	for a.scanner.Scan() {
		line := a.scanner.Text()
		if line == "" {
			continue
		}

		var msg rpcResponse
		if err := json.Unmarshal([]byte(line), &msg); err != nil {
			log.Printf("[acp] failed to parse message: %v", err)
			continue
		}

		// Response to a request we made (has id, no method)
		if msg.ID != nil && msg.Method == "" {
			a.pendingMu.Lock()
			ch, ok := a.pending[*msg.ID]
			a.pendingMu.Unlock()
			if ok {
				ch <- &msg
			}
			continue
		}

		// Request from agent or notification
		switch msg.Method {
		case "session/update":
			a.handleSessionUpdate(msg.Params)

		case "session/request_permission":
			// Auto-allow all permissions
			a.handlePermissionRequest(line)

		// Codex app-server events (multiple protocol versions)
		case "codex/event/agent_message_delta":
			a.handleCodexDelta(msg.Params)
		case "item/agentMessage/delta":
			a.handleCodexItemDelta(msg.Params)
		case "item/started":
			a.handleCodexItemStarted(msg.Params)
		case "turn/started", "turn/completed":
			a.handleCodexTurnEvent(msg.Method, msg.Params)
		case "codex/event/agent_message", "codex/event/task_complete",
			"codex/event/item_completed", "codex/event/token_count",
			"item/completed", "thread/tokenUsage/updated",
			"account/rateLimits/updated", "thread/status/changed":
			// Known events we don't need to act on
		case "turn/approval/request":
			a.handlePermissionRequest(line)

		default:
			if msg.Method != "" {
				log.Printf("[acp] unhandled method: %s (raw: %.200s)", msg.Method, line)
			}
		}
	}

	if err := a.scanner.Err(); err != nil {
		log.Printf("[acp] read loop error: %v", err)
	}
	log.Println("[acp] read loop ended")
}

func (a *ACPAgent) handleSessionUpdate(params json.RawMessage) {
	var p sessionUpdateParams
	if err := json.Unmarshal(params, &p); err != nil {
		log.Printf("[acp] failed to parse session/update: %v (raw: %s)", err, string(params))
		return
	}

	// Only log non-streaming events (skip chunks to reduce noise)
	switch p.Update.SessionUpdate {
	case "agent_message_chunk", "agent_thought_chunk":
		// skip — too noisy
	default:
		log.Printf("[acp] session/update (session=%s, type=%s)", p.SessionID, p.Update.SessionUpdate)
	}

	a.notifyMu.Lock()
	ch, ok := a.notifyCh[p.SessionID]
	a.notifyMu.Unlock()

	if ok {
		select {
		case ch <- &p.Update:
		default:
			log.Printf("[acp] notification channel full, dropping update (session=%s)", p.SessionID)
		}
	}
}

func (a *ACPAgent) handleCodexDelta(params json.RawMessage) {
	var p struct {
		Msg struct {
			Type  string `json:"type"`
			Delta string `json:"delta"`
		} `json:"msg"`
		ConversationID string `json:"conversationId"`
		ThreadID       string `json:"threadId"` // some versions use threadId
	}
	if err := json.Unmarshal(params, &p); err != nil {
		return
	}

	// Try conversationId first (codex uses this), fallback to threadId
	key := p.ConversationID
	if key == "" {
		key = p.ThreadID
	}

	delta := p.Msg.Delta
	if delta == "" {
		return
	}

	// Find the turn channel by thread ID — we need to match against stored threads
	a.notifyMu.Lock()
	ch, ok := a.turnCh[key]
	if !ok {
		// Try matching by iterating all turn channels (codex uses conversationId, not threadId)
		for _, c := range a.turnCh {
			ch = c
			ok = true
			break
		}
	}
	a.notifyMu.Unlock()

	if ok {
		select {
		case ch <- &codexTurnEvent{Delta: delta}:
		default:
		}
	}
}

// handleCodexItemDelta handles "item/agentMessage/delta" events.
// These contain incremental text deltas for the agent's response.
func (a *ACPAgent) handleCodexItemDelta(params json.RawMessage) {
	var p struct {
		ThreadID string `json:"threadId"`
		ItemID   string `json:"itemId"`
		Delta    string `json:"delta"`
	}
	if err := json.Unmarshal(params, &p); err != nil {
		return
	}

	if p.Delta == "" {
		return
	}

	a.dispatchToTurnCh(p.ThreadID, &codexTurnEvent{Delta: p.Delta})
}

// handleCodexItemStarted handles "item/started" events.
// When type=agentMessage, extracts text from content array.
func (a *ACPAgent) handleCodexItemStarted(params json.RawMessage) {
	var p struct {
		ThreadID string `json:"threadId"`
		Item     struct {
			Type    string `json:"type"`
			Content []struct {
				Type string `json:"type"`
				Text string `json:"text"`
			} `json:"content"`
		} `json:"item"`
	}
	if err := json.Unmarshal(params, &p); err != nil {
		return
	}

	if p.Item.Type != "agentMessage" {
		return
	}

	for _, c := range p.Item.Content {
		if c.Type == "text" && c.Text != "" {
			a.dispatchToTurnCh(p.ThreadID, &codexTurnEvent{Text: c.Text})
		}
	}
}

// handleCodexTurnEvent handles "turn/started" and "turn/completed" notifications.
func (a *ACPAgent) handleCodexTurnEvent(method string, params json.RawMessage) {
	var p struct {
		ThreadID string `json:"threadId"`
		Status   string `json:"status"`
	}
	if err := json.Unmarshal(params, &p); err != nil {
		return
	}

	if method == "turn/completed" {
		a.dispatchToTurnCh(p.ThreadID, &codexTurnEvent{Kind: "completed"})
	}
}

// dispatchToTurnCh sends an event to the turn channel for a thread.
func (a *ACPAgent) dispatchToTurnCh(threadID string, evt *codexTurnEvent) {
	a.notifyMu.Lock()
	ch, ok := a.turnCh[threadID]
	if !ok {
		// Fallback: try any active turn channel
		for _, c := range a.turnCh {
			ch = c
			ok = true
			break
		}
	}
	a.notifyMu.Unlock()

	if ok {
		select {
		case ch <- evt:
		default:
		}
	}
}

func (a *ACPAgent) handlePermissionRequest(raw string) {
	// Parse the request to get the ID and auto-allow
	var req struct {
		ID     json.RawMessage         `json:"id"`
		Params permissionRequestParams `json:"params"`
	}
	if err := json.Unmarshal([]byte(raw), &req); err != nil {
		log.Printf("[acp] failed to parse permission request: %v", err)
		return
	}

	// Find the "allow" option
	optionID := "allow"
	for _, opt := range req.Params.Options {
		if opt.Kind == "allow" {
			optionID = opt.OptionID
			break
		}
	}

	// Send response
	resp := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      req.ID,
		"result": map[string]interface{}{
			"outcome": map[string]interface{}{
				"outcome":  "selected",
				"optionId": optionID,
			},
		},
	}

	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("[acp] failed to marshal permission response: %v", err)
		return
	}

	a.mu.Lock()
	fmt.Fprintf(a.stdin, "%s\n", data)
	a.mu.Unlock()

	log.Printf("[acp] auto-allowed permission request")
}

// Info returns metadata about this agent.
func (a *ACPAgent) Info() AgentInfo {
	info := AgentInfo{
		Name:    a.command,
		Type:    "acp",
		Model:   a.model,
		Command: a.command,
	}
	a.mu.Lock()
	if a.cmd != nil && a.cmd.Process != nil {
		info.PID = a.cmd.Process.Pid
	}
	a.mu.Unlock()
	return info
}

func extractChunkText(update *sessionUpdate) string {
	// The content field in agent_message_chunk can be a text content block
	if update.Text != "" {
		return update.Text
	}

	// Try to extract from content JSON
	if update.Content != nil {
		var content struct {
			Type string `json:"type"`
			Text string `json:"text"`
		}
		if err := json.Unmarshal(update.Content, &content); err == nil && content.Text != "" {
			return content.Text
		}
	}

	return ""
}

// extractPromptResultText tries to extract text from the session/prompt response.
// Some ACP agents include response content in the result alongside stopReason.
func extractPromptResultText(result json.RawMessage) string {
	if result == nil {
		return ""
	}

	// Try to extract content array from result
	var r struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
		// Some agents use a flat text field
		Text string `json:"text"`
	}
	if err := json.Unmarshal(result, &r); err != nil {
		return ""
	}

	if r.Text != "" {
		return r.Text
	}

	var parts []string
	for _, c := range r.Content {
		if c.Type == "text" && c.Text != "" {
			parts = append(parts, c.Text)
		}
	}
	return strings.Join(parts, "")
}

// acpStderrWriter forwards the ACP subprocess stderr to the application log
// and captures the last meaningful error line.
type acpStderrWriter struct {
	prefix string
	mu     sync.Mutex
	last   string // last non-empty, non-traceback line
}

func (w *acpStderrWriter) Write(p []byte) (int, error) {
	lines := strings.Split(strings.TrimRight(string(p), "\n"), "\n")
	w.mu.Lock()
	for _, line := range lines {
		if line != "" {
			log.Printf("%s %s", w.prefix, line)
			// Capture lines that look like actual error messages (not traceback frames)
			if !strings.HasPrefix(line, "  ") && !strings.HasPrefix(line, "Traceback") && !strings.HasPrefix(line, "...") {
				w.last = line
			}
		}
	}
	w.mu.Unlock()
	return len(p), nil
}

// LastError returns the last captured error line and resets it.
func (w *acpStderrWriter) LastError() string {
	w.mu.Lock()
	defer w.mu.Unlock()
	s := w.last
	w.last = ""
	return s
}

````

[⬆ 回到目录](#toc)

<a name="file-agent-agent.go"></a>
## 📄 agent/agent.go

````go
package agent

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// MediaEntry represents a media item (image, file, video) in a message.
type MediaEntry struct {
	Type     string // "image", "file", "video"
	URL      string // download URL (if available)
	Path     string // local file path (after download)
	MIMEType string // MIME type (if known)
	FileName string // original filename (for files)
}

// AgentInfo holds metadata about an agent for logging/debugging.
type AgentInfo struct {
	Name    string // e.g. "claude-acp", "claude", "gpt-4o"
	Type    string // e.g. "acp", "cli", "http"
	Model   string // e.g. "sonnet", "gpt-4o-mini"
	Command string // binary path, e.g. "/usr/local/bin/claude-agent-acp"
	PID     int    // subprocess PID (0 if not applicable, e.g. http agent)
}

// String returns a human-readable summary for logging.
func (i AgentInfo) String() string {
	s := fmt.Sprintf("name=%s, type=%s, model=%s, command=%s", i.Name, i.Type, i.Model, i.Command)
	if i.PID > 0 {
		s += fmt.Sprintf(", pid=%d", i.PID)
	}
	return s
}

// defaultWorkspace returns ~/.weclaw/workspace as the default working directory.
func defaultWorkspace() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return os.TempDir()
	}
	dir := filepath.Join(home, ".weclaw", "workspace")
	os.MkdirAll(dir, 0o755)
	return dir
}

// mergeEnv merges extra environment variables into the base environment.
func mergeEnv(base []string, extra map[string]string) ([]string, error) {
	if len(extra) == 0 {
		return base, nil
	}

	merged := append([]string(nil), base...)
	indexByKey := make(map[string]int, len(base))
	for i, entry := range merged {
		key, _, found := strings.Cut(entry, "=")
		if !found || key == "" {
			continue
		}
		indexByKey[key] = i
	}

	newKeys := make([]string, 0, len(extra))
	for key, value := range extra {
		if key == "" || strings.Contains(key, "=") {
			return nil, fmt.Errorf("invalid env key %q", key)
		}
		entry := key + "=" + value
		if idx, ok := indexByKey[key]; ok {
			merged[idx] = entry
			continue
		}
		newKeys = append(newKeys, key)
	}

	sort.Strings(newKeys)
	for _, key := range newKeys {
		merged = append(merged, key+"="+extra[key])
	}

	return merged, nil
}

// Agent is the interface for AI chat agents.
type Agent interface {
	// Chat sends a message to the agent and returns the response.
	// conversationID is used to maintain conversation history per user.
	Chat(ctx context.Context, conversationID string, message string) (string, error)

	// ChatWithMedia sends a message with media attachments to the agent.
	// media can contain images, files, videos, etc.
	ChatWithMedia(ctx context.Context, conversationID string, message string, media []MediaEntry) (string, error)

	// ResetSession clears the existing session for the given conversationID and
	// starts a new one. Returns the new session ID if immediately available
	// (ACP mode), or an empty string if the ID will be assigned on next Chat
	// (CLI mode) or is not applicable (HTTP mode).
	ResetSession(ctx context.Context, conversationID string) (string, error)

	// Info returns metadata about this agent.
	Info() AgentInfo

	// SetCwd changes the working directory for subsequent operations.
	SetCwd(cwd string)
}

````

[⬆ 回到目录](#toc)

<a name="file-agent-cli_agent.go"></a>
## 📄 agent/cli_agent.go

````go
package agent

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

// CLIAgent invokes a local CLI agent (claude, codex, etc.) via streaming JSON.
type CLIAgent struct {
	name         string
	command      string
	args         []string          // extra args from config
	cwd          string            // working directory
	env          map[string]string // extra environment variables
	model        string
	systemPrompt string
	mu           sync.Mutex
	sessions     map[string]string // conversationID -> session ID for multi-turn
}

// CLIAgentConfig holds configuration for a CLI agent.
type CLIAgentConfig struct {
	Name         string            // agent name for logging, e.g. "claude", "codex"
	Command      string            // path to binary
	Args         []string          // extra args (e.g. ["--dangerously-skip-permissions"])
	Cwd          string            // working directory (workspace)
	Env          map[string]string // extra environment variables
	Model        string
	SystemPrompt string
}

// NewCLIAgent creates a new CLI agent.
func NewCLIAgent(cfg CLIAgentConfig) *CLIAgent {
	cwd := cfg.Cwd
	if cwd == "" {
		cwd = defaultWorkspace()
	}
	return &CLIAgent{
		name:         cfg.Name,
		command:      cfg.Command,
		args:         cfg.Args,
		cwd:          cwd,
		env:          cfg.Env,
		model:        cfg.Model,
		systemPrompt: cfg.SystemPrompt,
		sessions:     make(map[string]string),
	}
}

// streamEvent represents a single event from claude's stream-json output.
type streamEvent struct {
	Type      string         `json:"type"`
	SessionID string         `json:"session_id"`
	Result    string         `json:"result"`
	IsError   bool           `json:"is_error"`
	Message   *streamMessage `json:"message,omitempty"`
}

// streamMessage represents the message field in an assistant event.
type streamMessage struct {
	Content []streamContent `json:"content"`
}

// streamContent represents a content block in an assistant message.
type streamContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// Info returns metadata about this agent.
func (a *CLIAgent) Info() AgentInfo {
	return AgentInfo{
		Name:    a.name,
		Type:    "cli",
		Model:   a.model,
		Command: a.command,
	}
}

// ResetSession clears the existing session for the given conversationID.
// Returns an empty string because the new session ID is only known after the
// next Chat call (claude assigns it during the conversation).
func (a *CLIAgent) ResetSession(_ context.Context, conversationID string) (string, error) {
	a.mu.Lock()
	delete(a.sessions, conversationID)
	a.mu.Unlock()
	log.Printf("[cli] session reset (command=%s, conversation=%s)", a.command, conversationID)
	return "", nil
}

// SetCwd changes the working directory for subsequent CLI invocations.
func (a *CLIAgent) SetCwd(cwd string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.cwd = cwd
}

// Chat sends a message to the CLI agent and returns the response.
func (a *CLIAgent) Chat(ctx context.Context, conversationID string, message string) (string, error) {
	switch a.name {
	case "codex":
		return a.chatCodex(ctx, message)
	default:
		return a.chatClaude(ctx, conversationID, message)
	}
}

// ChatWithMedia sends a message with media attachments.
// CLI agents currently don't support media natively, so we add media info to the message.
func (a *CLIAgent) ChatWithMedia(ctx context.Context, conversationID string, message string, media []MediaEntry) (string, error) {
	// Build enhanced message with media descriptions
	enhancedMessage := message
	for _, m := range media {
		switch m.Type {
		case "image":
			enhancedMessage += fmt.Sprintf("\n[图片: %s]", m.URL)
		case "file":
			enhancedMessage += fmt.Sprintf("\n[文件: %s]", m.FileName)
		case "video":
			enhancedMessage += fmt.Sprintf("\n[视频: %s]", m.URL)
		}
	}
	return a.Chat(ctx, conversationID, enhancedMessage)
}

// chatClaude uses claude -p with stream-json to get structured output and session persistence.
func (a *CLIAgent) chatClaude(ctx context.Context, conversationID string, message string) (string, error) {
	args := []string{"-p", message, "--output-format", "stream-json", "--verbose"}

	if a.model != "" {
		args = append(args, "--model", a.model)
	}
	if a.systemPrompt != "" {
		args = append(args, "--append-system-prompt", a.systemPrompt)
	}
	// Append extra args from config (e.g. --dangerously-skip-permissions)
	args = append(args, a.args...)

	// Resume existing session for multi-turn conversation
	a.mu.Lock()
	sessionID, hasSession := a.sessions[conversationID]
	a.mu.Unlock()

	if hasSession {
		args = append(args, "--resume", sessionID)
		log.Printf("[cli] resuming session (command=%s, session=%s, conversation=%s)", a.command, sessionID, conversationID)
	} else {
		log.Printf("[cli] starting new conversation (command=%s, conversation=%s)", a.command, conversationID)
	}

	cmd := exec.CommandContext(ctx, a.command, args...)
	if a.cwd != "" {
		cmd.Dir = a.cwd
	}
	if len(a.env) > 0 {
		cmdEnv, err := mergeEnv(os.Environ(), a.env)
		if err != nil {
			return "", fmt.Errorf("build %s env: %w", a.name, err)
		}
		cmd.Env = cmdEnv
	}
	var stderr strings.Builder
	cmd.Stderr = &stderr

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("create stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("start %s: %w", a.name, err)
	}

	log.Printf("[cli] spawned process (command=%s, pid=%d, conversation=%s)", a.command, cmd.Process.Pid, conversationID)

	// Parse streaming JSON events
	var result string
	var newSessionID string
	var assistantTexts []string

	scanner := bufio.NewScanner(stdout)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024) // 1MB buffer for large responses

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var event streamEvent
		if err := json.Unmarshal([]byte(line), &event); err != nil {
			continue
		}

		// Capture session ID from any event
		if event.SessionID != "" {
			newSessionID = event.SessionID
		}

		switch event.Type {
		case "result":
			if event.IsError {
				return "", fmt.Errorf("%s returned error: %s", a.name, event.Result)
			}
			result = event.Result
		case "assistant":
			// Newer claude CLI versions send text in assistant events
			// instead of the result event's result field.
			if event.Message != nil {
				for _, c := range event.Message.Content {
					if c.Type == "text" && c.Text != "" {
						assistantTexts = append(assistantTexts, c.Text)
					}
				}
			}
		}
	}

	// If the result event had an empty result, fall back to accumulated assistant texts.
	if result == "" && len(assistantTexts) > 0 {
		result = strings.Join(assistantTexts, "")
	}

	if err := cmd.Wait(); err != nil {
		if result == "" {
			errMsg := strings.TrimSpace(stderr.String())
			if errMsg != "" {
				return "", fmt.Errorf("%s exited with error: %w, stderr: %s", a.name, err, errMsg)
			}
			return "", fmt.Errorf("%s exited with error: %w", a.name, err)
		}
		// If we got a result but exit code is non-zero (e.g. hook failures), still return the result
	}

	log.Printf("[cli] process exited (command=%s, pid=%d)", a.command, cmd.Process.Pid)

	// Save session ID for multi-turn conversation
	if newSessionID != "" {
		a.mu.Lock()
		a.sessions[conversationID] = newSessionID
		a.mu.Unlock()
		log.Printf("[cli] saved session (session=%s, conversation=%s)", newSessionID, conversationID)
	}

	result = strings.TrimSpace(result)
	if result == "" {
		return "", fmt.Errorf("%s returned empty response", a.name)
	}

	return result, nil
}

// chatCodex handles codex CLI invocation using "codex exec".
func (a *CLIAgent) chatCodex(ctx context.Context, message string) (string, error) {
	args := []string{"exec", message}
	if a.model != "" {
		args = append(args, "--model", a.model)
	}
	// Append extra args from config (e.g. --skip-git-repo-check)
	args = append(args, a.args...)

	log.Printf("[cli] running codex exec (command=%s)", a.command)
	cmd := exec.CommandContext(ctx, a.command, args...)
	if a.cwd != "" {
		cmd.Dir = a.cwd
	}
	if len(a.env) > 0 {
		cmdEnv, err := mergeEnv(os.Environ(), a.env)
		if err != nil {
			return "", fmt.Errorf("build %s env: %w", a.name, err)
		}
		cmd.Env = cmdEnv
	}
	var stderr strings.Builder
	cmd.Stderr = &stderr

	out, err := cmd.Output()
	if err != nil {
		errMsg := strings.TrimSpace(stderr.String())
		if errMsg != "" {
			return "", fmt.Errorf("codex error: %w, stderr: %s", err, errMsg)
		}
		return "", fmt.Errorf("codex error: %w", err)
	}

	result := strings.TrimSpace(string(out))
	if result == "" {
		return "", fmt.Errorf("codex returned empty response")
	}
	return result, nil
}

````

[⬆ 回到目录](#toc)

<a name="file-agent-env_test.go"></a>
## 📄 agent/env_test.go

````go
package agent

import (
	"reflect"
	"testing"
)

func TestMergeEnvOverridesAndAppends(t *testing.T) {
	base := []string{"PATH=/usr/bin", "KEEP=1", "DUP=old"}
	extra := map[string]string{
		"NEW":   "2",
		"DUP":   "new",
		"EMPTY": "",
	}

	got, err := mergeEnv(base, extra)
	if err != nil {
		t.Fatalf("mergeEnv returned error: %v", err)
	}

	want := []string{"PATH=/usr/bin", "KEEP=1", "DUP=new", "EMPTY=", "NEW=2"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("mergeEnv() = %#v, want %#v", got, want)
	}
}

func TestMergeEnvRejectsInvalidKey(t *testing.T) {
	_, err := mergeEnv(nil, map[string]string{"BAD=KEY": "value"})
	if err == nil {
		t.Fatal("mergeEnv() error = nil, want invalid env key error")
	}
}

func TestMergeEnvPreservesBaseWhenNoExtra(t *testing.T) {
	base := []string{"PATH=/usr/bin", "KEEP=1"}

	got, err := mergeEnv(base, nil)
	if err != nil {
		t.Fatalf("mergeEnv returned error: %v", err)
	}
	if !reflect.DeepEqual(got, base) {
		t.Fatalf("mergeEnv() = %#v, want %#v", got, base)
	}
}

func TestMergeEnvRejectsEmptyKey(t *testing.T) {
	_, err := mergeEnv(nil, map[string]string{"": "value"})
	if err == nil {
		t.Fatal("mergeEnv() error = nil, want empty env key error")
	}
}

func TestMergeEnvOverridesExistingKeyWithEmptyValue(t *testing.T) {
	got, err := mergeEnv([]string{"EMPTY=old"}, map[string]string{"EMPTY": ""})
	if err != nil {
		t.Fatalf("mergeEnv returned error: %v", err)
	}
	want := []string{"EMPTY="}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("mergeEnv() = %#v, want %#v", got, want)
	}
}

````

[⬆ 回到目录](#toc)

<a name="file-agent-http_agent.go"></a>
## 📄 agent/http_agent.go

````go
package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// ChatMessage represents a single message in a conversation.
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// HTTPAgent is an OpenAI-compatible chat completions API client.
type HTTPAgent struct {
	endpoint     string
	apiKey       string
	headers      map[string]string
	model        string
	systemPrompt string
	httpClient   *http.Client
	mu           sync.Mutex
	history      map[string][]ChatMessage // conversationID -> messages
	maxHistory   int
}

// HTTPAgentConfig holds configuration for the HTTP agent.
type HTTPAgentConfig struct {
	Endpoint     string
	APIKey       string
	Headers      map[string]string
	Model        string
	SystemPrompt string
	MaxHistory   int
}

// NewHTTPAgent creates a new OpenAI-compatible HTTP agent.
func NewHTTPAgent(cfg HTTPAgentConfig) *HTTPAgent {
	if cfg.MaxHistory == 0 {
		cfg.MaxHistory = 20
	}
	if cfg.Model == "" {
		cfg.Model = "gpt-4o-mini"
	}
	return &HTTPAgent{
		endpoint:     cfg.Endpoint,
		apiKey:       cfg.APIKey,
		headers:      cfg.Headers,
		model:        cfg.Model,
		systemPrompt: cfg.SystemPrompt,
		httpClient:   &http.Client{Timeout: 120 * time.Second},
		history:      make(map[string][]ChatMessage),
		maxHistory:   cfg.MaxHistory,
	}
}

// Info returns metadata about this agent.
func (a *HTTPAgent) Info() AgentInfo {
	return AgentInfo{
		Name:    "http",
		Type:    "http",
		Model:   a.model,
		Command: a.endpoint,
	}
}

// SetCwd is a no-op for HTTP agents (they have no working directory).
func (a *HTTPAgent) SetCwd(_ string) {}

// ResetSession clears the conversation history for the given conversationID.
// HTTP agents have no server-side session ID, so an empty string is returned.
func (a *HTTPAgent) ResetSession(_ context.Context, conversationID string) (string, error) {
	a.mu.Lock()
	delete(a.history, conversationID)
	a.mu.Unlock()
	return "", nil
}

// Chat sends a message to the OpenAI-compatible API and returns the response.
func (a *HTTPAgent) Chat(ctx context.Context, conversationID string, message string) (string, error) {
	a.mu.Lock()
	messages := a.buildMessages(conversationID, message)
	a.mu.Unlock()

	reqBody := map[string]interface{}{
		"model":    a.model,
		"messages": messages,
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, a.endpoint, bytes.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if a.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+a.apiKey)
	}
	for k, v := range a.headers {
		req.Header.Set(k, v)
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error HTTP %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("parse response: %w", err)
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}

	reply := result.Choices[0].Message.Content

	// Save to history
	a.mu.Lock()
	a.history[conversationID] = append(a.history[conversationID],
		ChatMessage{Role: "user", Content: message},
		ChatMessage{Role: "assistant", Content: reply},
	)
	// Trim history
	if len(a.history[conversationID]) > a.maxHistory*2 {
		a.history[conversationID] = a.history[conversationID][len(a.history[conversationID])-a.maxHistory*2:]
	}
	a.mu.Unlock()

	return reply, nil
}

// ChatWithMedia sends a message with media attachments.
// For HTTP agents, media is converted to text description (limited support).
func (a *HTTPAgent) ChatWithMedia(ctx context.Context, conversationID string, message string, media []MediaEntry) (string, error) {
	// Build enhanced message with media descriptions
	enhancedMessage := message
	for _, m := range media {
		switch m.Type {
		case "image":
			enhancedMessage += fmt.Sprintf("\n[图片: %s]", m.URL)
		case "file":
			enhancedMessage += fmt.Sprintf("\n[文件: %s (%s)]", m.FileName, m.URL)
		case "video":
			enhancedMessage += fmt.Sprintf("\n[视频: %s]", m.URL)
		}
	}
	return a.Chat(ctx, conversationID, enhancedMessage)
}

func (a *HTTPAgent) buildMessages(conversationID string, message string) []ChatMessage {
	var messages []ChatMessage
	if a.systemPrompt != "" {
		messages = append(messages, ChatMessage{Role: "system", Content: a.systemPrompt})
	}
	if hist, ok := a.history[conversationID]; ok {
		messages = append(messages, hist...)
	}
	messages = append(messages, ChatMessage{Role: "user", Content: message})
	return messages
}

````

[⬆ 回到目录](#toc)

<a name="file-api-server.go"></a>
## 📄 api/server.go

````go
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fastclaw-ai/weclaw/ilink"
	"github.com/fastclaw-ai/weclaw/messaging"
)

// Server provides an HTTP API for sending messages.
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
	mux.HandleFunc("/api/send", s.handleSend)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

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

````

[⬆ 回到目录](#toc)

<a name="file-cmd-login.go"></a>
## 📄 cmd/login.go

````go
package cmd

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Add a WeChat account via QR code scan",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()

		creds, err := doLogin(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("Account %s added. Run 'weclaw start' to begin.\n", creds.ILinkBotID)
		return nil
	},
}

````

[⬆ 回到目录](#toc)

<a name="file-cmd-proc_unix.go"></a>
## 📄 cmd/proc_unix.go

````go
//go:build !windows

package cmd

import (
	"os/exec"
	"syscall"
)

func setSysProcAttr(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
}

````

[⬆ 回到目录](#toc)

<a name="file-cmd-proc_windows.go"></a>
## 📄 cmd/proc_windows.go

````go
//go:build windows

package cmd

import "os/exec"

func setSysProcAttr(_ *exec.Cmd) {
	// No Setsid on Windows — process is already detached via Start()
}

````

[⬆ 回到目录](#toc)

<a name="file-cmd-restart.go"></a>
## 📄 cmd/restart.go

````go
package cmd

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(restartCmd)
}

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the background weclaw process",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Stop if running
		pid, err := readPid()
		if err == nil && processExists(pid) {
			fmt.Printf("Stopping weclaw (pid=%d)...\n", pid)
			if p, err := os.FindProcess(pid); err == nil {
				p.Signal(syscall.SIGTERM)
			}
			for i := 0; i < 20; i++ {
				if !processExists(pid) {
					break
				}
				time.Sleep(500 * time.Millisecond)
			}
			os.Remove(pidFile())
		}

		// Start
		fmt.Println("Starting weclaw...")
		return runDaemon()
	},
}

````

[⬆ 回到目录](#toc)

<a name="file-cmd-root.go"></a>
## 📄 cmd/root.go

````go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Version is set at build time via -ldflags.
var Version = "dev"

var rootCmd = &cobra.Command{
	Use:     "weclaw",
	Short:   "WeChat AI agent bridge",
	Long:    "weclaw bridges WeChat messages to AI agents via the iLink API.",
	Version: Version,
	RunE:    runStart, // default command is start
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

````

[⬆ 回到目录](#toc)

<a name="file-cmd-send.go"></a>
## 📄 cmd/send.go

````go
package cmd

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/fastclaw-ai/weclaw/ilink"
	"github.com/fastclaw-ai/weclaw/messaging"
	"github.com/spf13/cobra"
)

var (
	sendTo       string
	sendText     string
	sendMediaURL string
)

func init() {
	sendCmd.Flags().StringVar(&sendTo, "to", "", "Target user ID (ilink user ID)")
	sendCmd.Flags().StringVar(&sendText, "text", "", "Message text to send")
	sendCmd.Flags().StringVar(&sendMediaURL, "media", "", "Media URL to send (image/video/file)")
	sendCmd.MarkFlagRequired("to")
	rootCmd.AddCommand(sendCmd)
}

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message to a WeChat user",
	Example: `  weclaw send --to "user_id@im.wechat" --text "Hello"
  weclaw send --to "user_id@im.wechat" --media "https://example.com/image.png"
  weclaw send --to "user_id@im.wechat" --text "See this" --media "https://example.com/image.png"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if sendText == "" && sendMediaURL == "" {
			return fmt.Errorf("at least one of --text or --media is required")
		}

		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()

		accounts, err := ilink.LoadAllCredentials()
		if err != nil {
			return fmt.Errorf("load credentials: %w", err)
		}
		if len(accounts) == 0 {
			return fmt.Errorf("no accounts found, run 'weclaw start' first")
		}

		client := ilink.NewClient(accounts[0])

		if sendText != "" {
			if err := messaging.SendTextReply(ctx, client, sendTo, sendText, "", ""); err != nil {
				return fmt.Errorf("send text failed: %w", err)
			}
			fmt.Println("Text sent")
		}

		if sendMediaURL != "" {
			if err := messaging.SendMediaFromURL(ctx, client, sendTo, sendMediaURL, ""); err != nil {
				return fmt.Errorf("send media failed: %w", err)
			}
			fmt.Println("Media sent")
		}

		return nil
	},
}

````

[⬆ 回到目录](#toc)

<a name="file-cmd-start.go"></a>
## 📄 cmd/start.go

````go
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/fastclaw-ai/weclaw/agent"
	"github.com/fastclaw-ai/weclaw/api"
	"github.com/fastclaw-ai/weclaw/config"
	"github.com/fastclaw-ai/weclaw/ilink"
	"github.com/fastclaw-ai/weclaw/messaging"
	"github.com/mdp/qrterminal/v3"
	"github.com/spf13/cobra"
)

var (
	foregroundFlag bool
	apiAddrFlag    string
)

func init() {
	startCmd.Flags().BoolVarP(&foregroundFlag, "foreground", "f", false, "Run in foreground (default is background)")
	startCmd.Flags().StringVar(&apiAddrFlag, "api-addr", "", "API server listen address (default 127.0.0.1:18011)")
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the WeChat message bridge (auto-login if needed)",
	RunE:  runStart,
}

func runStart(cmd *cobra.Command, args []string) error {
	if !foregroundFlag {
		// Check if login is needed — if so, do it in foreground first, then daemon
		accounts, _ := ilink.LoadAllCredentials()
		if len(accounts) == 0 {
			fmt.Println("No WeChat accounts found, starting login...")
			ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
			_, err := doLogin(ctx)
			cancel()
			if err != nil {
				return fmt.Errorf("login failed: %w", err)
			}
		}
		return runDaemon()
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// Load all accounts
	accounts, err := ilink.LoadAllCredentials()
	if err != nil {
		return fmt.Errorf("failed to load credentials: %w", err)
	}

	// No accounts — trigger login
	if len(accounts) == 0 {
		log.Println("No WeChat accounts found, starting login...")
		creds, err := doLogin(ctx)
		if err != nil {
			return fmt.Errorf("login failed: %w", err)
		}
		accounts = append(accounts, creds)
	}

	// Load config and auto-detect agents
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if config.DetectAndConfigure(cfg) {
		if err := config.Save(cfg); err != nil {
			log.Printf("Warning: failed to save auto-detected config: %v", err)
		} else {
			path, _ := config.ConfigPath()
			log.Printf("Auto-detected agents saved to %s", path)
		}
	}

	// Log all available agents
	if len(cfg.Agents) > 0 {
		names := make([]string, 0, len(cfg.Agents))
		for name := range cfg.Agents {
			names = append(names, name)
		}
		log.Printf("Available agents: %v (default: %s)", names, cfg.DefaultAgent)
	}

	// Create handler with an agent factory for on-demand agent creation
	handler := messaging.NewHandler(
		func(ctx context.Context, name string) agent.Agent {
			return createAgentByName(ctx, cfg, name)
		},
		func(name string) error {
			cfg.DefaultAgent = name
			return config.Save(cfg)
		},
	)

	// Populate agent metas for /status
	var metas []messaging.AgentMeta
	workDirs := make(map[string]string, len(cfg.Agents))
	for name, agCfg := range cfg.Agents {
		command := agCfg.Command
		if agCfg.Type == "http" {
			command = agCfg.Endpoint
		}
		metas = append(metas, messaging.AgentMeta{
			Name:    name,
			Type:    agCfg.Type,
			Command: command,
			Model:   agCfg.Model,
		})
		if agCfg.Cwd != "" {
			workDirs[name] = agCfg.Cwd
		}
	}
	handler.SetAgentMetas(metas)
	handler.SetAgentWorkDirs(workDirs)

	// Load custom aliases from agent configs
	handler.SetCustomAliases(config.BuildAliasMap(cfg.Agents))

	// Set save directory for images/files if configured
	if cfg.SaveDir != "" {
		handler.SetSaveDir(cfg.SaveDir)
		log.Printf("Image save directory: %s", cfg.SaveDir)
	}

	// Start default agent initialization in background so monitors can start immediately
	go func() {
		if cfg.DefaultAgent == "" {
			log.Println("No default agent configured, staying in echo mode")
			return
		}
		log.Printf("Initializing default agent %q in background...", cfg.DefaultAgent)
		ag := createAgentByName(ctx, cfg, cfg.DefaultAgent)
		if ag == nil {
			log.Printf("Failed to initialize default agent %q, staying in echo mode", cfg.DefaultAgent)
		} else {
			handler.SetDefaultAgent(cfg.DefaultAgent, ag)
		}
	}()

	// Start HTTP API server for sending messages
	var clients []*ilink.Client
	for _, c := range accounts {
		clients = append(clients, ilink.NewClient(c))
	}
	// Resolve API addr: flag > env/config > default
	apiAddr := cfg.APIAddr // already includes env override from loadEnv
	if apiAddrFlag != "" {
		apiAddr = apiAddrFlag
	}
	apiServer := api.NewServer(clients, apiAddr)
	go func() {
		if err := apiServer.Run(ctx); err != nil {
			log.Printf("API server error: %v", err)
		}
	}()

	// Start monitors immediately — they will use echo mode until agent is ready
	log.Printf("Starting message bridge for %d account(s)...", len(accounts))

	var wg sync.WaitGroup
	for _, creds := range accounts {
		wg.Add(1)
		go func(c *ilink.Credentials) {
			defer wg.Done()
			runMonitorWithRestart(ctx, c, handler)
		}(creds)
	}

	wg.Wait()
	log.Println("All monitors stopped")
	return nil
}

// runMonitorWithRestart runs a monitor with automatic restart on failure.
func runMonitorWithRestart(ctx context.Context, creds *ilink.Credentials, handler *messaging.Handler) {
	const maxRestartDelay = 30 * time.Second
	restartDelay := 3 * time.Second

	for {
		log.Printf("[%s] Starting monitor...", creds.ILinkBotID)

		client := ilink.NewClient(creds)
		monitor, err := ilink.NewMonitor(client, handler.HandleMessage)
		if err != nil {
			log.Printf("[%s] Failed to create monitor: %v", creds.ILinkBotID, err)
		} else {
			err = monitor.Run(ctx)
		}

		// If context is cancelled, exit
		if ctx.Err() != nil {
			return
		}

		log.Printf("[%s] Monitor stopped: %v, restarting in %s", creds.ILinkBotID, err, restartDelay)
		select {
		case <-time.After(restartDelay):
		case <-ctx.Done():
			return
		}

		// Exponential backoff for restarts, capped
		restartDelay *= 2
		if restartDelay > maxRestartDelay {
			restartDelay = maxRestartDelay
		}
	}
}

// createAgentByName creates and starts an agent by its config name.
// Returns nil if the agent is not configured or fails to start.
func createAgentByName(ctx context.Context, cfg *config.Config, name string) agent.Agent {
	agCfg, ok := cfg.Agents[name]
	if !ok {
		log.Printf("[agent] %q not found in config", name)
		return nil
	}

	switch agCfg.Type {
	case "acp":
		ag := agent.NewACPAgent(agent.ACPAgentConfig{
			Command:      agCfg.Command,
			Args:         agCfg.Args,
			Cwd:          agCfg.Cwd,
			Env:          agCfg.Env,
			Model:        agCfg.Model,
			SystemPrompt: agCfg.SystemPrompt,
		})
		if err := ag.Start(ctx); err != nil {
			log.Printf("[agent] failed to start ACP agent %q: %v", name, err)
			return nil
		}
		log.Printf("[agent] started ACP agent: %s (command=%s, type=%s, model=%s)", name, agCfg.Command, agCfg.Type, agCfg.Model)
		return ag
	case "cli":
		ag := agent.NewCLIAgent(agent.CLIAgentConfig{
			Name:         name,
			Command:      agCfg.Command,
			Args:         agCfg.Args,
			Cwd:          agCfg.Cwd,
			Env:          agCfg.Env,
			Model:        agCfg.Model,
			SystemPrompt: agCfg.SystemPrompt,
		})
		log.Printf("[agent] created CLI agent: %s (command=%s, type=%s, model=%s)", name, agCfg.Command, agCfg.Type, agCfg.Model)
		return ag
	case "http":
		if agCfg.Endpoint == "" {
			log.Printf("[agent] HTTP agent %q has no endpoint", name)
			return nil
		}
		ag := agent.NewHTTPAgent(agent.HTTPAgentConfig{
			Endpoint:     agCfg.Endpoint,
			APIKey:       agCfg.APIKey,
			Headers:      agCfg.Headers,
			Model:        agCfg.Model,
			SystemPrompt: agCfg.SystemPrompt,
			MaxHistory:   agCfg.MaxHistory,
		})
		log.Printf("[agent] created HTTP agent: %s (endpoint=%s, model=%s)", name, agCfg.Endpoint, agCfg.Model)
		return ag
	default:
		log.Printf("[agent] unknown type %q for %q", agCfg.Type, name)
		return nil
	}
}

// doLogin runs the interactive QR login flow and returns credentials.
func doLogin(ctx context.Context) (*ilink.Credentials, error) {
	fmt.Println("Fetching QR code...")
	qr, err := ilink.FetchQRCode(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch QR code: %w", err)
	}

	fmt.Println("\nScan this QR code with WeChat:")
	fmt.Println()
	qrterminal.GenerateWithConfig(qr.QRCodeImgContent, qrterminal.Config{
		Level:          qrterminal.L,
		Writer:         os.Stdout,
		HalfBlocks:     true,
		BlackChar:      qrterminal.BLACK_BLACK,
		WhiteBlackChar: qrterminal.WHITE_BLACK,
		WhiteChar:      qrterminal.WHITE_WHITE,
		BlackWhiteChar: qrterminal.BLACK_WHITE,
		QuietZone:      1,
	})
	fmt.Printf("\nQR URL: %s\n", qr.QRCodeImgContent)
	fmt.Println("\nWaiting for scan...")

	lastStatus := ""
	creds, err := ilink.PollQRStatus(ctx, qr.QRCode, func(status string) {
		if status != lastStatus {
			lastStatus = status
			switch status {
			case "scaned":
				fmt.Println("QR code scanned! Please confirm on your phone.")
			case "confirmed":
				fmt.Println("Login confirmed!")
			case "expired":
				fmt.Println("QR code expired.")
			}
		}
	})
	if err != nil {
		return nil, err
	}

	if err := ilink.SaveCredentials(creds); err != nil {
		return nil, fmt.Errorf("failed to save credentials: %w", err)
	}

	dir, _ := ilink.CredentialsPath()
	fmt.Printf("\nLogin successful! Credentials saved to %s\n", dir)
	fmt.Printf("Bot ID: %s\n\n", creds.ILinkBotID)
	return creds, nil
}

// --- Daemon mode ---

func weclawDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".weclaw")
}

func pidFile() string {
	return filepath.Join(weclawDir(), "weclaw.pid")
}

func logFile() string {
	return filepath.Join(weclawDir(), "weclaw.log")
}

// runDaemon spawns weclaw start (without --daemon) as a background process.
func runDaemon() error {
	// Kill any existing weclaw processes before starting a new one
	stopAllWeclaw()

	// Ensure log directory exists
	if err := os.MkdirAll(weclawDir(), 0o700); err != nil {
		return fmt.Errorf("create weclaw dir: %w", err)
	}

	// Open log file
	lf, err := os.OpenFile(logFile(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return fmt.Errorf("open log file: %w", err)
	}

	// Re-exec ourselves without --daemon
	exe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("find executable: %w", err)
	}

	cmd := exec.Command(exe, "start", "-f")
	cmd.Stdout = lf
	cmd.Stderr = lf
	setSysProcAttr(cmd)

	if err := cmd.Start(); err != nil {
		lf.Close()
		return fmt.Errorf("start daemon: %w", err)
	}

	// Save PID
	pid := cmd.Process.Pid
	os.WriteFile(pidFile(), []byte(fmt.Sprintf("%d", pid)), 0o644)

	// Detach — don't wait
	cmd.Process.Release()
	lf.Close()

	fmt.Printf("weclaw started in background (pid=%d)\n", pid)
	fmt.Printf("Log: %s\n", logFile())
	fmt.Printf("Stop: weclaw stop\n")
	return nil
}

func readPid() (int, error) {
	data, err := os.ReadFile(pidFile())
	if err != nil {
		return 0, err
	}
	var pid int
	if _, err := fmt.Sscanf(string(data), "%d", &pid); err != nil {
		return 0, err
	}
	return pid, nil
}

func processExists(pid int) bool {
	p, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	// Signal 0 checks if process exists without killing it
	return p.Signal(syscall.Signal(0)) == nil
}

// stopAllWeclaw kills all running weclaw processes (by PID file and by process scan).
func stopAllWeclaw() {
	// 1. Kill by PID file
	if pid, err := readPid(); err == nil && processExists(pid) {
		if p, err := os.FindProcess(pid); err == nil {
			_ = p.Signal(syscall.SIGTERM)
		}
	}
	os.Remove(pidFile())

	// 2. Kill any remaining weclaw processes by scanning
	exe, err := os.Executable()
	if err != nil {
		return
	}
	// Use pkill to kill all processes matching the executable path
	_ = exec.Command("pkill", "-f", exe+" start").Run()
	time.Sleep(500 * time.Millisecond)
}

````

[⬆ 回到目录](#toc)

<a name="file-cmd-status.go"></a>
## 📄 cmd/status.go

````go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check if weclaw is running in background",
	RunE: func(cmd *cobra.Command, args []string) error {
		pid, err := readPid()
		if err != nil {
			fmt.Println("weclaw is not running")
			return nil
		}

		if processExists(pid) {
			fmt.Printf("weclaw is running (pid=%d)\n", pid)
			fmt.Printf("Log: %s\n", logFile())
		} else {
			fmt.Println("weclaw is not running (stale pid file)")
		}
		return nil
	},
}

````

[⬆ 回到目录](#toc)

<a name="file-cmd-stop.go"></a>
## 📄 cmd/stop.go

````go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the background weclaw process",
	RunE: func(cmd *cobra.Command, args []string) error {
		stopAllWeclaw()
		fmt.Println("weclaw stopped")
		return nil
	},
}

````

[⬆ 回到目录](#toc)

<a name="file-cmd-update.go"></a>
## 📄 cmd/update.go

````go
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const githubRepo = "fastclaw-ai/weclaw"

func init() {
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(upgradeCmd)
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("weclaw %s (%s/%s)\n", Version, runtime.GOOS, runtime.GOARCH)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update weclaw to the latest version and restart",
	RunE:  runUpdate,
}

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Update weclaw to the latest version and restart (alias for update)",
	RunE:  runUpdate,
}

func runUpdate(cmd *cobra.Command, args []string) error {
	// 1. Get latest version
	fmt.Println("Checking for updates...")
	latest, err := getLatestVersion()
	if err != nil {
		return fmt.Errorf("failed to check latest version: %w", err)
	}

	if latest == Version {
		fmt.Printf("Already up to date (%s)\n", Version)
		return nil
	}

	fmt.Printf("Current: %s -> Latest: %s\n", Version, latest)

	// 2. Download new binary
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	filename := fmt.Sprintf("weclaw_%s_%s", goos, goarch)
	url := fmt.Sprintf("https://github.com/%s/releases/download/%s/%s", githubRepo, latest, filename)

	fmt.Printf("Downloading %s...\n", url)
	tmpFile, err := downloadFile(url)
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}
	defer os.Remove(tmpFile)

	// 3. Replace current binary
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("find executable: %w", err)
	}
	// Resolve symlinks
	if resolved, err := resolveSymlink(exePath); err == nil {
		exePath = resolved
	}

	if err := replaceBinary(tmpFile, exePath); err != nil {
		return fmt.Errorf("replace binary: %w", err)
	}

	// Clear macOS quarantine/provenance attributes to avoid Gatekeeper killing the binary
	if runtime.GOOS == "darwin" {
		exec.Command("xattr", "-d", "com.apple.quarantine", exePath).Run()
		exec.Command("xattr", "-d", "com.apple.provenance", exePath).Run()
	}

	fmt.Printf("Updated to %s\n", latest)

	// 4. Restart if running in background
	pid, pidErr := readPid()
	if pidErr == nil && processExists(pid) {
		fmt.Println("Stopping old process...")
		if p, err := os.FindProcess(pid); err == nil {
			p.Signal(os.Interrupt)
		}
		// Wait for old process to exit
		for i := 0; i < 20; i++ {
			if !processExists(pid) {
				break
			}
			time.Sleep(500 * time.Millisecond)
		}
		os.Remove(pidFile())

		fmt.Println("Starting new version...")
		if err := runDaemon(); err != nil {
			log.Printf("Failed to restart: %v", err)
			fmt.Println("Update complete. Please run 'weclaw start' manually.")
		}
	} else {
		fmt.Println("Update complete. Run 'weclaw start' to start.")
	}

	return nil
}

func getLatestVersion() (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", githubRepo))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API returned %d", resp.StatusCode)
	}

	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err
	}
	return release.TagName, nil
}

func downloadFile(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	tmp, err := os.CreateTemp("", "weclaw-update-*")
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(tmp, resp.Body); err != nil {
		tmp.Close()
		os.Remove(tmp.Name())
		return "", err
	}
	tmp.Close()

	if err := os.Chmod(tmp.Name(), 0o755); err != nil {
		os.Remove(tmp.Name())
		return "", err
	}

	return tmp.Name(), nil
}

func replaceBinary(src, dst string) error {
	// Check if we can write directly
	if err := os.Rename(src, dst); err == nil {
		return nil
	}

	// Try with sudo on Unix
	if runtime.GOOS != "windows" {
		fmt.Printf("Installing to %s (requires sudo)...\n", dst)
		cmd := exec.Command("sudo", "cp", src, dst)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	return fmt.Errorf("cannot write to %s", dst)
}

func resolveSymlink(path string) (string, error) {
	for {
		target, err := os.Readlink(path)
		if err != nil {
			return path, nil
		}
		if !strings.HasPrefix(target, "/") {
			// Relative symlink
			dir := path[:strings.LastIndex(path, "/")+1]
			target = dir + target
		}
		path = target
	}
}

````

[⬆ 回到目录](#toc)

<a name="file-config-config.go"></a>
## 📄 config/config.go

````go
package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Config holds the application configuration.
type Config struct {
	DefaultAgent string                 `json:"default_agent"`
	APIAddr      string                 `json:"api_addr,omitempty"`
	SaveDir      string                 `json:"save_dir,omitempty"`
	Agents       map[string]AgentConfig `json:"agents"`
}

// AgentConfig holds configuration for a single agent.
type AgentConfig struct {
	Type         string            `json:"type"`                    // "acp", "cli", or "http"
	Command      string            `json:"command,omitempty"`       // binary path (cli/acp type)
	Args         []string          `json:"args,omitempty"`          // extra args for command (e.g. ["acp"] for cursor)
	Aliases      []string          `json:"aliases,omitempty"`       // custom trigger commands (e.g. ["gpt", "4o"])
	Cwd          string            `json:"cwd,omitempty"`           // working directory (workspace)
	Env          map[string]string `json:"env,omitempty"`           // extra environment variables (cli/acp type)
	Model        string            `json:"model,omitempty"`         // model name
	SystemPrompt string            `json:"system_prompt,omitempty"` // system prompt
	Endpoint     string            `json:"endpoint,omitempty"`      // API endpoint (http type)
	APIKey       string            `json:"api_key,omitempty"`       // API key (http type)
	Headers      map[string]string `json:"headers,omitempty"`       // extra HTTP headers (http type)
	MaxHistory   int               `json:"max_history,omitempty"`   // max history (http type)
}

// BuildAliasMap builds a map from custom alias to agent name from all agent configs.
// It logs warnings for conflicts: duplicate aliases and aliases shadowing agent keys.
func BuildAliasMap(agents map[string]AgentConfig) map[string]string {
	// Built-in commands that cannot be overridden
	reserved := map[string]bool{
		"info": true, "help": true, "new": true, "clear": true, "cwd": true,
	}

	m := make(map[string]string)
	for name, cfg := range agents {
		for _, alias := range cfg.Aliases {
			if reserved[alias] {
				log.Printf("[config] WARNING: alias %q for agent %q conflicts with built-in command, ignored", alias, name)
				continue
			}
			if existing, ok := m[alias]; ok {
				log.Printf("[config] WARNING: alias %q is defined by both %q and %q, using %q", alias, existing, name, name)
			}
			m[alias] = name
		}
	}

	// Warn if a custom alias shadows an agent key
	for alias, target := range m {
		if _, isAgent := agents[alias]; isAgent && alias != target {
			log.Printf("[config] WARNING: alias %q (-> %q) shadows agent key %q", alias, target, alias)
		}
	}

	return m
}

// DefaultConfig returns an empty configuration.
func DefaultConfig() *Config {
	return &Config{
		Agents: make(map[string]AgentConfig),
	}
}

// ConfigPath returns the path to the config file.
func ConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".weclaw", "config.json"), nil
}

// Load loads configuration from disk and environment variables.
func Load() (*Config, error) {
	cfg := DefaultConfig()

	path, err := ConfigPath()
	if err != nil {
		return cfg, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			loadEnv(cfg)
			return cfg, nil
		}
		return nil, fmt.Errorf("read config: %w", err)
	}

	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	if cfg.Agents == nil {
		cfg.Agents = make(map[string]AgentConfig)
	}

	loadEnv(cfg)
	return cfg, nil
}

func loadEnv(cfg *Config) {
	if v := os.Getenv("WECLAW_DEFAULT_AGENT"); v != "" {
		cfg.DefaultAgent = v
	}
	if v := os.Getenv("WECLAW_API_ADDR"); v != "" {
		cfg.APIAddr = v
	}
	if v := os.Getenv("WECLAW_SAVE_DIR"); v != "" {
		cfg.SaveDir = v
	}
}

// Save saves the configuration to disk.
func Save(cfg *Config) error {
	path, err := ConfigPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return fmt.Errorf("create config dir: %w", err)
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}

	return os.WriteFile(path, data, 0o600)
}

````

[⬆ 回到目录](#toc)

<a name="file-config-config_test.go"></a>
## 📄 config/config_test.go

````go
package config

import (
	"encoding/json"
	"testing"
)

func TestAgentConfigUnmarshalEnv(t *testing.T) {
	var cfg Config
	data := []byte(`{
		"agents": {
			"claude": {
				"type": "cli",
				"command": "claude",
				"env": {
					"ANTHROPIC_API_KEY": "test-key",
					"EMPTY": ""
				}
			}
		}
	}`)

	if err := json.Unmarshal(data, &cfg); err != nil {
		t.Fatalf("unmarshal config: %v", err)
	}

	ag, ok := cfg.Agents["claude"]
	if !ok {
		t.Fatalf("expected claude agent config")
	}
	if got := ag.Env["ANTHROPIC_API_KEY"]; got != "test-key" {
		t.Fatalf("ANTHROPIC_API_KEY = %q, want %q", got, "test-key")
	}
	if got, ok := ag.Env["EMPTY"]; !ok || got != "" {
		t.Fatalf("EMPTY = %q, present=%v; want empty string present", got, ok)
	}
}

func TestAgentConfigMarshalEnvRoundTrip(t *testing.T) {
	cfg := Config{
		Agents: map[string]AgentConfig{
			"claude": {
				Type:    "cli",
				Command: "claude",
				Env: map[string]string{
					"ANTHROPIC_API_KEY": "test-key",
					"EMPTY":             "",
				},
			},
		},
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("marshal config: %v", err)
	}

	var decoded Config
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("round-trip unmarshal: %v", err)
	}

	got := decoded.Agents["claude"].Env
	if got["ANTHROPIC_API_KEY"] != "test-key" || got["EMPTY"] != "" {
		t.Fatalf("round-trip env = %#v", got)
	}
}

func TestAgentConfigWithoutEnvStillLoads(t *testing.T) {
	var cfg Config
	data := []byte(`{
		"agents": {
			"claude": {
				"type": "cli",
				"command": "claude"
			}
		}
	}`)

	if err := json.Unmarshal(data, &cfg); err != nil {
		t.Fatalf("unmarshal config without env: %v", err)
	}

	if cfg.Agents["claude"].Env != nil {
		t.Fatalf("Env = %#v, want nil", cfg.Agents["claude"].Env)
	}
}

func TestDefaultConfigInitializesAgentsMap(t *testing.T) {
	cfg := DefaultConfig()
	if cfg.Agents == nil {
		t.Fatal("DefaultConfig() Agents = nil, want initialized map")
	}
}

func TestLoadEnvOverridesTopLevelOnly(t *testing.T) {
	t.Setenv("WECLAW_DEFAULT_AGENT", "codex")
	t.Setenv("WECLAW_API_ADDR", "127.0.0.1:18011")

	cfg := DefaultConfig()
	cfg.Agents["claude"] = AgentConfig{
		Type: "cli",
		Env: map[string]string{
			"KEEP": "value",
		},
	}

	loadEnv(cfg)

	if cfg.DefaultAgent != "codex" {
		t.Fatalf("DefaultAgent = %q, want %q", cfg.DefaultAgent, "codex")
	}
	if cfg.APIAddr != "127.0.0.1:18011" {
		t.Fatalf("APIAddr = %q, want %q", cfg.APIAddr, "127.0.0.1:18011")
	}
	if got := cfg.Agents["claude"].Env["KEEP"]; got != "value" {
		t.Fatalf("agent env = %q, want preserved value", got)
	}
}

````

[⬆ 回到目录](#toc)

<a name="file-config-detect.go"></a>
## 📄 config/detect.go

````go
package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// agentCandidate defines one way to run an agent.
// Multiple candidates can map to the same agent name; the first detected wins.
type agentCandidate struct {
	Name      string   // agent name (e.g. "claude", "codex")
	Binary    string   // binary to look up in PATH
	Args      []string // extra args (e.g. ["acp"] for cursor)
	CheckArgs []string // optional capability probe args (must exit 0)
	Type      string   // "acp", "cli"
	Model     string   // default model
}

// agentCandidates is ordered by priority: for each agent name, earlier entries
// are preferred. E.g. claude ACP is tried before claude CLI.
var agentCandidates = []agentCandidate{
	// claude: prefer ACP, fallback to CLI
	{Name: "claude", Binary: "claude-agent-acp", Type: "acp", Model: "sonnet"},
	{Name: "claude", Binary: "claude", Type: "cli", Model: "sonnet"},
	// codex: prefer ACP, fallback to CLI
	{Name: "codex", Binary: "codex-acp", Type: "acp", Model: ""},
	{Name: "codex", Binary: "codex", Args: []string{"app-server", "--listen", "stdio://"}, CheckArgs: []string{"app-server", "--help"}, Type: "acp", Model: ""},
	{Name: "codex", Binary: "codex", Type: "cli", Model: ""},
	// ACP-only agents
	{Name: "cursor", Binary: "agent", Args: []string{"acp"}, Type: "acp", Model: ""},
	{Name: "kimi", Binary: "kimi", Args: []string{"acp"}, Type: "acp", Model: ""},
	{Name: "gemini", Binary: "gemini", Args: []string{"--acp"}, Type: "acp", Model: ""},
	{Name: "opencode", Binary: "opencode", Args: []string{"acp"}, Type: "acp", Model: ""},
	{Name: "openclaw", Binary: "openclaw", Type: "acp", Model: "openclaw:main"}, // args built dynamically
	{Name: "pi", Binary: "pi-acp", Type: "acp", Model: ""},
	{Name: "copilot", Binary: "copilot", Args: []string{"--acp", "--stdio"}, Type: "acp", Model: ""},
	{Name: "droid", Binary: "droid", Args: []string{"exec", "--output-format", "acp"}, Type: "acp", Model: ""},
	{Name: "iflow", Binary: "iflow", Args: []string{"--experimental-acp"}, Type: "acp", Model: ""},
	{Name: "kiro", Binary: "kiro-cli", Args: []string{"acp"}, Type: "acp", Model: ""},
	{Name: "qwen", Binary: "qwen", Args: []string{"--acp"}, Type: "acp", Model: ""},
}

// defaultOrder defines the priority for choosing the default agent.
// Lower index = higher priority.
var defaultOrder = []string{
	"claude", "codex", "cursor", "kimi", "gemini", "opencode", "openclaw",
	"pi", "copilot", "droid", "iflow", "kiro", "qwen",
}

// DetectAndConfigure auto-detects local agents and populates the config.
// For each agent name, it picks the highest-priority candidate (acp > cli).
// Returns true if the config was modified.
func DetectAndConfigure(cfg *Config) bool {
	modified := false

	for _, candidate := range agentCandidates {
		// Skip if this agent name is already configured
		if _, exists := cfg.Agents[candidate.Name]; exists {
			continue
		}

		path, err := lookPath(candidate.Binary)
		if err != nil {
			continue
		}

		// Run capability probe if specified
		if len(candidate.CheckArgs) > 0 && !commandProbe(path, candidate.CheckArgs) {
			log.Printf("[config] skipping %s at %s (type=%s): probe failed (%v)", candidate.Name, path, candidate.Type, candidate.CheckArgs)
			continue
		}

		log.Printf("[config] auto-detected %s at %s (type=%s)", candidate.Name, path, candidate.Type)
		cfg.Agents[candidate.Name] = AgentConfig{
			Type:    candidate.Type,
			Command: path,
			Args:    candidate.Args,
			Model:   candidate.Model,
		}
		modified = true
	}

	// Special handling for openclaw: prefer HTTP mode over ACP to avoid
	// session routing conflicts with openclaw-weixin plugin (see #9).
	// Priority: HTTP (gateway) > ACP (with user-configured --session) > skip.
	if agCfg, exists := cfg.Agents["openclaw"]; exists && agCfg.Type == "acp" && len(agCfg.Args) == 0 {
		gwURL, gwToken, gwPassword := loadOpenclawGateway()
		if gwURL != "" {
			// Prefer HTTP mode — no session routing issues
			httpURL := gwURL
			httpURL = strings.Replace(httpURL, "wss://", "https://", 1)
			httpURL = strings.Replace(httpURL, "ws://", "http://", 1)
			endpoint := strings.TrimRight(httpURL, "/") + "/v1/chat/completions"
			log.Printf("[config] openclaw using HTTP mode: %s", endpoint)
			cfg.Agents["openclaw"] = AgentConfig{
				Type:     "http",
				Endpoint: endpoint,
				APIKey:   gwToken,
				Headers:  map[string]string{"x-openclaw-scopes": "operator.write"},
				Model:    "openclaw:main",
			}
			modified = true

			// Also register openclaw-acp as a separate agent for users who want ACP
			if _, apcExists := cfg.Agents["openclaw-acp"]; !apcExists {
				args := []string{"acp", "--url", gwURL}
				if gwToken != "" {
					args = append(args, "--token", gwToken)
				} else if gwPassword != "" {
					args = append(args, "--password", gwPassword)
				}
				cfg.Agents["openclaw-acp"] = AgentConfig{
					Type:    "acp",
					Command: agCfg.Command,
					Args:    args,
					Model:   "openclaw:main",
				}
				log.Printf("[config] openclaw ACP also available as 'openclaw-acp' (use /openclaw-acp to switch)")
			}
		} else {
			log.Printf("[config] openclaw binary found but no gateway config, skipping")
			delete(cfg.Agents, "openclaw")
			modified = true
		}
	}

	// Fallback: if openclaw still not configured, try HTTP via gateway config.
	if _, exists := cfg.Agents["openclaw"]; !exists {
		gwURL, gwToken, _ := loadOpenclawGateway()
		if gwURL != "" {
			httpURL := gwURL
			httpURL = strings.Replace(httpURL, "wss://", "https://", 1)
			httpURL = strings.Replace(httpURL, "ws://", "http://", 1)
			endpoint := strings.TrimRight(httpURL, "/") + "/v1/chat/completions"
			log.Printf("[config] using openclaw HTTP: %s", endpoint)
			cfg.Agents["openclaw"] = AgentConfig{
				Type:     "http",
				Endpoint: endpoint,
				APIKey:   gwToken,
				Headers:  map[string]string{"x-openclaw-scopes": "operator.write"},
				Model:    "openclaw:main",
			}
			modified = true
		}
	}

	// Pick the highest-priority default agent.
	if cfg.DefaultAgent == "" || !agentExists(cfg, cfg.DefaultAgent) {
		for _, name := range defaultOrder {
			if _, ok := cfg.Agents[name]; ok {
				if cfg.DefaultAgent != name {
					log.Printf("[config] setting default agent: %s", name)
					cfg.DefaultAgent = name
					modified = true
				}
				break
			}
		}
	}

	return modified
}

// loadOpenclawGateway resolves openclaw gateway connection info.
// Priority: env vars > ~/.openclaw/openclaw.json.
// Returns (url, token, password). url="" means not configured.
func loadOpenclawGateway() (gwURL, gwToken, gwPassword string) {
	// 1. Environment variables take priority
	gwURL = os.Getenv("OPENCLAW_GATEWAY_URL")
	gwToken = os.Getenv("OPENCLAW_GATEWAY_TOKEN")
	gwPassword = os.Getenv("OPENCLAW_GATEWAY_PASSWORD")
	if gwURL != "" {
		return
	}

	// 2. Read from ~/.openclaw/openclaw.json
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}

	data, err := os.ReadFile(filepath.Join(home, ".openclaw", "openclaw.json"))
	if err != nil {
		return
	}

	var ocCfg struct {
		Gateway struct {
			Port int    `json:"port"`
			Mode string `json:"mode"`
			Auth struct {
				Mode     string `json:"mode"`
				Token    string `json:"token"`
				Password string `json:"password"`
			} `json:"auth"`
			Remote struct {
				URL   string `json:"url"`
				Token string `json:"token"`
			} `json:"remote"`
		} `json:"gateway"`
	}
	if err := json.Unmarshal(data, &ocCfg); err != nil {
		log.Printf("[config] failed to parse openclaw config: %v", err)
		return
	}

	gw := ocCfg.Gateway

	// Remote gateway (gateway.remote.url)
	if gw.Remote.URL != "" {
		gwURL = gw.Remote.URL
		gwToken = gw.Remote.Token
		return
	}

	// Local gateway (gateway.port + gateway.auth)
	if gw.Port > 0 {
		gwURL = fmt.Sprintf("ws://127.0.0.1:%d", gw.Port)
		switch gw.Auth.Mode {
		case "token":
			gwToken = gw.Auth.Token
		case "password":
			gwPassword = gw.Auth.Password
		}
		return
	}

	return
}

// commandProbe runs a binary with args and returns true if it exits 0.
func commandProbe(binary string, args []string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, binary, args...)
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	return cmd.Run() == nil
}

func agentExists(cfg *Config, name string) bool {
	_, ok := cfg.Agents[name]
	return ok
}

// lookPath finds a binary by name. It first tries exec.LookPath (fast, uses
// current PATH). If that fails, it falls back to resolving via a login shell
// which sources the user's profile (~/.zshrc, ~/.bashrc) — this picks up
// binaries installed through version managers like nvm, mise, etc. that only
// add their paths in interactive shells.
func lookPath(binary string) (string, error) {
	// Fast path: binary is in current PATH
	if p, err := exec.LookPath(binary); err == nil {
		return p, nil
	}

	// Fallback: resolve via login interactive shell (sources .zshrc/.bashrc)
	shell := "zsh"
	if runtime.GOOS != "darwin" {
		shell = "bash"
	}
	out, err := exec.Command(shell, "-lic", "which "+binary).Output()
	if err != nil {
		return "", fmt.Errorf("not found: %s", binary)
	}
	p := strings.TrimSpace(string(out))
	if p == "" || strings.Contains(p, "not found") {
		return "", fmt.Errorf("not found: %s", binary)
	}
	log.Printf("[config] resolved %s via login shell: %s", binary, p)
	return p, nil
}

````

[⬆ 回到目录](#toc)

<a name="file-config-detect_test.go"></a>
## 📄 config/detect_test.go

````go
package config

import (
	"os"
	"os/exec"
	"testing"
)

// TestLookPath_InPath verifies that lookPath finds binaries already in PATH.
func TestLookPath_InPath(t *testing.T) {
	p, err := lookPath("ls")
	if err != nil {
		t.Fatalf("expected to find ls, got error: %v", err)
	}
	if p == "" {
		t.Fatal("expected non-empty path for ls")
	}
}

// TestLookPath_NotExist verifies that lookPath returns an error for missing binaries.
func TestLookPath_NotExist(t *testing.T) {
	_, err := lookPath("nonexistent-binary-xyz-12345")
	if err == nil {
		t.Fatal("expected error for nonexistent binary")
	}
}

// TestLookPath_LoginShellFallback reproduces the daemon scenario:
// PATH is stripped to system-only dirs (no nvm), so exec.LookPath fails,
// but lookPath resolves claude via login shell fallback.
func TestLookPath_LoginShellFallback(t *testing.T) {
	// Precondition: claude must be discoverable via login shell (i.e. nvm in .zshrc)
	fullPath, err := exec.LookPath("claude")
	if err != nil {
		t.Skip("claude not installed, skipping login shell fallback test")
	}

	// Simulate daemon environment: strip PATH to system-only dirs
	origPath := os.Getenv("PATH")
	os.Setenv("PATH", "/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin")
	defer os.Setenv("PATH", origPath)

	// Reproduce the bug: exec.LookPath must fail under stripped PATH
	_, err = exec.LookPath("claude")
	if err == nil {
		t.Skip("claude found in minimal PATH, cannot reproduce nvm issue")
	}

	// Verify fix: lookPath should find claude via login shell
	p, err := lookPath("claude")
	if err != nil {
		t.Fatalf("lookPath should find claude via login shell, got: %v", err)
	}
	if p != fullPath {
		t.Logf("resolved path differs: direct=%s, login-shell=%s (acceptable)", fullPath, p)
	}
	t.Logf("lookPath resolved claude via login shell: %s", p)
}

// TestDetectAndConfigure_StrippedPath is an end-to-end test:
// empty config + stripped PATH → DetectAndConfigure should still find claude.
func TestDetectAndConfigure_StrippedPath(t *testing.T) {
	if _, err := exec.LookPath("claude"); err != nil {
		t.Skip("claude not installed, skipping")
	}

	origPath := os.Getenv("PATH")
	os.Setenv("PATH", "/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin")
	defer os.Setenv("PATH", origPath)

	cfg := DefaultConfig()
	DetectAndConfigure(cfg)

	agent, ok := cfg.Agents["claude"]
	if !ok {
		t.Fatal("expected claude to be detected via login shell fallback")
	}
	if agent.Type != "cli" {
		t.Fatalf("expected type=cli, got %s", agent.Type)
	}
	t.Logf("detected claude: type=%s, command=%s", agent.Type, agent.Command)
}

````

[⬆ 回到目录](#toc)

<a name="file-docs-readme_cn.md"></a>
## 📄 docs/README_CN.md

````markdown
# WeClaw

[English](README.md)

微信 AI Agent 桥接器 — 将微信消息接入 AI Agent（Claude、Codex、Gemini、Kimi 等）。

> 本项目参考 [@tencent-weixin/openclaw-weixin](https://npmx.dev/package/@tencent-weixin/openclaw-weixin) 实现，仅限个人学习，勿做他用。

|                                                 |                                                 |                                                 |
| :---------------------------------------------: | :---------------------------------------------: | :---------------------------------------------: |
| <img src="previews/preview1.png" width="280" /> | <img src="previews/preview2.png" width="280" /> | <img src="previews/preview3.png" width="280" /> |

## 快速开始

```bash
# 一键安装
curl -sSL https://raw.githubusercontent.com/fastclaw-ai/weclaw/main/install.sh | sh

# 启动（首次运行会弹出微信扫码登录）
weclaw start
```

就这么简单。首次启动时，WeClaw 会：

1. 显示二维码 — 用微信扫码登录
2. 自动检测已安装的 AI Agent（Claude、Codex、Gemini 等）
3. 保存配置到 `~/.weclaw/config.json`
4. 开始接收和回复微信消息

使用 `weclaw login` 可以添加更多微信账号。

### 其他安装方式

```bash
# 通过 Go 安装
go install github.com/fastclaw-ai/weclaw@latest

# 通过 Docker
docker run -it -v ~/.weclaw:/root/.weclaw ghcr.io/fastclaw-ai/weclaw start
```

## 架构

<p align="center">
  <img src="previews/architecture.png" width="600" />
</p>

**Agent 接入模式：**

| 模式 | 工作方式                                                         | 支持的 Agent                                            |
| ---- | ---------------------------------------------------------------- | ------------------------------------------------------- |
| ACP  | 长驻子进程，通过 stdio JSON-RPC 通信。速度最快，复用进程和会话。 | Claude, Codex, Kimi, Gemini, Cursor, OpenCode, OpenClaw |
| CLI  | 每条消息启动一个新进程，支持通过 `--resume` 恢复会话。           | Claude (`claude -p`)、Codex (`codex exec`)              |
| HTTP | OpenAI 兼容的 Chat Completions API。                             | OpenClaw（HTTP 回退）                                   |

同时存在 ACP 和 CLI 时，自动优先选择 ACP。

## 聊天命令

在微信中发送以下命令：

| 命令                    | 说明                     |
| ----------------------- | ------------------------ |
| `你好`                  | 发送给默认 Agent         |
| `/codex 写一个排序函数` | 发送给指定 Agent         |
| `/cc 解释一下这段代码`  | 通过别名发送             |
| `/claude`               | 切换默认 Agent 为 Claude |
| `/cwd /path/to/project` | 切换工作目录             |
| `/new`                  | 开始新对话（清除会话）   |
| `/info`                 | 查看当前 Agent 信息      |
| `/help`                 | 查看帮助信息             |

### 快捷别名

| 别名   | Agent    |
| ------ | -------- |
| `/cc`  | Claude   |
| `/cx`  | Codex    |
| `/cs`  | Cursor   |
| `/km`  | Kimi     |
| `/gm`  | Gemini   |
| `/ocd` | OpenCode |
| `/oc`  | OpenClaw |

也可以在配置文件中为每个 Agent 自定义触发命令：

```json
{
  "agents": {
    "claude": {
      "type": "acp",
      "aliases": ["ai", "c"]
    }
  }
}
```

然后 `/ai 你好` 或 `/c 你好` 就会路由到 claude。

切换默认 Agent 会写入配置文件，重启后仍然生效。

## 富媒体消息

WeClaw 支持收发图片、视频、文件和语音消息。

**语音消息：** 在微信中发送语音消息时，WeClaw 会自动使用微信的语音转文字功能，将转写后的文本发送给 AI Agent。重复的语音消息事件会自动去重。

**Agent 回复自动处理：** 当 AI Agent 返回包含图片的 markdown（`![](url)`）时，WeClaw 会自动提取图片 URL，下载文件，上传到微信 CDN（AES-128-ECB 加密），然后作为图片消息发送。

**Markdown 转换：** Agent 的回复会自动从 markdown 转为纯文本再发送 — 代码块去掉围栏、链接只保留文字、加粗斜体标记去除等。

## 主动推送消息

无需等待用户发消息，主动向微信用户推送消息。

**命令行：**

```bash
# 发送文本
weclaw send --to "user_id@im.wechat" --text "你好，来自 weclaw"

# 发送图片
weclaw send --to "user_id@im.wechat" --media "https://example.com/photo.png"

# 发送文本 + 图片
weclaw send --to "user_id@im.wechat" --text "看看这个" --media "https://example.com/photo.png"

# 发送文件
weclaw send --to "user_id@im.wechat" --media "https://example.com/report.pdf"
```

**HTTP API**（`weclaw start` 运行时，默认监听 `127.0.0.1:18011`）：

```bash
# 发送文本
curl -X POST http://127.0.0.1:18011/api/send \
  -H "Content-Type: application/json" \
  -d '{"to": "user_id@im.wechat", "text": "你好，来自 weclaw"}'

# 发送图片
curl -X POST http://127.0.0.1:18011/api/send \
  -H "Content-Type: application/json" \
  -d '{"to": "user_id@im.wechat", "media_url": "https://example.com/photo.png"}'

# 发送文本 + 媒体
curl -X POST http://127.0.0.1:18011/api/send \
  -H "Content-Type: application/json" \
  -d '{"to": "user_id@im.wechat", "text": "看看这个", "media_url": "https://example.com/photo.png"}'
```

支持的媒体类型：图片（png、jpg、gif、webp）、视频（mp4、mov）、文件（pdf、doc、zip 等）。

设置 `WECLAW_API_ADDR` 环境变量可更改监听地址（如 `0.0.0.0:18011`）。

## 配置

配置文件路径：`~/.weclaw/config.json`

```json
{
  "default_agent": "claude",
  "agents": {
    "claude": {
      "type": "acp",
      "command": "/usr/local/bin/claude-agent-acp",
      "env": {
        "ANTHROPIC_API_KEY": "sk-ant-xxx"
      },
      "model": "sonnet"
    },
    "codex": {
      "type": "acp",
      "command": "/usr/local/bin/codex-acp",
      "env": {
        "OPENAI_API_KEY": "sk-xxx"
      }
    },
    "openclaw": {
      "type": "http",
      "endpoint": "https://api.example.com/v1/chat/completions",
      "api_key": "sk-xxx",
      "model": "openclaw:main"
    }
  }
}
```

环境变量：

- `WECLAW_DEFAULT_AGENT` — 覆盖默认 Agent
- `OPENCLAW_GATEWAY_URL` — OpenClaw HTTP 回退地址
- `OPENCLAW_GATEWAY_TOKEN` — OpenClaw API Token

自定义 agent cli 环境变量

```json
{
  "default_agent": "...",
  "agents": {
    "...": {
      ...
      "env": {
        "ENV_NAME": "ENV_VALUE"
      }
    },
  }
}
```

### 权限配置

部分 Agent 默认需要交互式权限确认，在微信场景下无法操作会导致卡住。可通过 `args` 配置跳过：

| Agent | 参数 | 说明 |
|-------|------|------|
| Claude (CLI) | `--dangerously-skip-permissions` | 跳过所有工具权限确认 |
| Codex (CLI) | `--skip-git-repo-check` | 允许在非 git 仓库目录运行 |

配置示例：

```json
{
  "claude": {
    "type": "cli",
    "command": "/usr/local/bin/claude",
    "cwd": "/home/user/my-project",
    "args": ["--dangerously-skip-permissions"]
  },
  "codex": {
    "type": "cli",
    "command": "/usr/local/bin/codex",
    "cwd": "/home/user/my-project",
    "args": ["--skip-git-repo-check"]
  }
}
```

通过 `cwd` 指定 Agent 的工作目录（workspace）。不设置则默认为 `~/.weclaw/workspace`。

> **注意：** 这些参数会跳过安全检查，请了解风险后再启用。ACP 模式的 Agent 会自动处理权限，无需配置。

## 后台运行

```bash
# 启动（默认后台运行）
weclaw start

# 查看状态
weclaw status

# 停止
weclaw stop

# 前台运行（调试用）
weclaw start -f
```

日志输出到 `~/.weclaw/weclaw.log`。

### 系统服务（开机自启）

**macOS (launchd)：**

```bash
cp service/com.fastclaw.weclaw.plist ~/Library/LaunchAgents/
launchctl load ~/Library/LaunchAgents/com.fastclaw.weclaw.plist
```

**Linux (systemd)：**

```bash
sudo cp service/weclaw.service /etc/systemd/system/
sudo systemctl enable --now weclaw
```

## Docker

```bash
# 构建
docker build -t weclaw .

# 登录（交互式，扫描二维码）
docker run -it -v ~/.weclaw:/root/.weclaw weclaw login

# 使用 HTTP Agent 启动
docker run -d --name weclaw \
  -v ~/.weclaw:/root/.weclaw \
  -e OPENCLAW_GATEWAY_URL=https://api.example.com \
  -e OPENCLAW_GATEWAY_TOKEN=sk-xxx \
  weclaw

# 查看日志
docker logs -f weclaw
```

> 注意：ACP 和 CLI 模式需要容器内有对应的 Agent 二进制文件。
> 默认镜像只包含 WeClaw 本体。如需使用 ACP/CLI Agent，请挂载二进制文件或构建自定义镜像。
> HTTP 模式开箱即用。

## 发版

```bash
# 打 tag 触发 GitHub Actions 自动构建发版
git tag v0.1.0
git push origin v0.1.0
```

自动构建 `darwin/linux/windows` x `amd64/arm64` 的二进制，创建 GitHub Release 并上传所有产物和校验文件。

## 更新

```bash
# 更新到最新版本（运行中会自动重启）
weclaw update

# 查看当前版本
weclaw version
```

## 开发

```bash
# 热重载
make dev

# 编译
go build -o weclaw .

# 运行
./weclaw start
```

## 贡献者

<a href="https://github.com/fastclaw-ai/weclaw/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=fastclaw-ai/weclaw" />
</a>

## Star 趋势

[![Star History Chart](https://api.star-history.com/svg?repos=fastclaw-ai/weclaw&type=Timeline)](https://star-history.com/#fastclaw-ai/weclaw&Timeline)

## 许可证

[MIT](LICENSE)

````

[⬆ 回到目录](#toc)

<a name="file-docs-项目学习.md"></a>
## 📄 docs/项目学习.md

````markdown
# WeClaw 项目学习笔记

> 对话时间：2026-03-30

---

## 1. 项目概述

**WeClaw** 是一个用 **Go 语言** 开发的微信 AI Agent 桥接器，将微信消息连接到各种 AI 代理（Claude、Codex、Gemini、Kimi 等）。

### 项目定位
- **核心功能**: 作为微信与 AI Agent 之间的桥梁
- **技术栈**: 纯 Go 语言实现，使用 Cobra CLI 框架
- **许可协议**: MIT 开源许可

### 项目灵感
```go
// 项目灵感来自腾讯官方 @tencent-weixin/openclaw-weixin
// 但 WeClaw 是独立实现的 Go 版本
```

---

## 2. 项目结构

```
weclaw/
├── main.go              # 程序入口点
├── cmd/                 # CLI 命令实现
│   ├── root.go          # 根命令 (Cobra)
│   ├── start.go         # 启动服务
│   ├── login.go         # 微信登录
│   ├── send.go          # 主动发送消息
│   ├── stop.go          # 停止服务
│   ├── status.go        # 查看状态
│   ├── update.go        # 更新版本
│   └── proc_*.go        # 进程管理 (跨平台)
├── agent/               # AI Agent 适配层
│   ├── agent.go         # Agent 接口定义
│   ├── acp_agent.go     # ACP 协议 Agent (1267行)
│   ├── cli_agent.go     # CLI 模式 Agent
│   └── http_agent.go    # HTTP API Agent
├── ilink/               # 微信 iLink 协议实现
│   ├── client.go        # iLink API 客户端
│   ├── auth.go          # 二维码登录认证
│   ├── monitor.go       # 消息长轮询监听
│   └── types.go         # 协议数据类型定义
├── messaging/           # 消息处理
│   ├── handler.go       # 消息路由与处理
│   ├── sender.go        # 消息发送
│   ├── media.go         # 媒体文件处理
│   ├── cdn.go           # 微信 CDN 上传/下载
│   └── markdown.go      # Markdown 转纯文本
├── api/                 # HTTP API 服务
│   └── server.go        # 主动消息推送 API
├── config/              # 配置管理
│   ├── config.go        # 配置加载/保存
│   └── detect.go        # Agent 自动检测
└── service/             # 系统服务配置
```

---

## 3. 多模式 Agent 支持

### 统一接口定义 (agent/agent.go)

```go
type Agent interface {
    Chat(ctx context.Context, conversationID string, message string) (string, error)
    ChatWithMedia(ctx context.Context, conversationID string, message string, media []MediaEntry) (string, error)
    ResetSession(ctx context.Context, conversationID string) (string, error)
    Info() AgentInfo
    SetCwd(cwd string)
}
```

### 三种模式对比

| 模式 | 说明 | 优势 | 实现文件 |
|------|------|------|----------|
| **ACP** | 长驻子进程，JSON-RPC 2.0 通信 | 速度最快，会话复用 | acp_agent.go (1267行) |
| **CLI** | 每条消息启动新进程 | 兼容性好，支持 `--resume` | cli_agent.go |
| **HTTP** | OpenAI 兼容 API | 易于集成，零代码接入 | http_agent.go |

### 支持的 Agent
`claude`、`codex`、`cursor`、`kimi`、`gemini`、`openclaw`、`opencode`、`pi`、`copilot`、`droid`、`iflow`、`kiro`、`qwen` 等

---

## 4. HTTP Agent 接入方式

### 配置示例 (~/.weclaw/config.json)

```json
{
  "default_agent": "gpt",
  "agents": {
    "gpt": {
      "type": "http",
      "endpoint": "https://api.openai.com/v1/chat/completions",
      "api_key": "sk-xxx",
      "model": "gpt-4o-mini",
      "system_prompt": "你是一个有用的助手",
      "aliases": ["4o", "chatgpt"]
    },
    "deepseek": {
      "type": "http",
      "endpoint": "https://api.deepseek.com/v1/chat/completions",
      "api_key": "sk-xxx",
      "model": "deepseek-chat",
      "aliases": ["ds"]
    },
    "本地模型": {
      "type": "http",
      "endpoint": "http://localhost:11434/v1/chat/completions",
      "model": "llama3",
      "aliases": ["llama"]
    }
  }
}
```

### 可接入的 API

| 服务商 | Endpoint |
|--------|----------|
| OpenAI | `https://api.openai.com/v1/chat/completions` |
| Azure OpenAI | `https://YOUR_RESOURCE.openai.azure.com/...` |
| DeepSeek | `https://api.deepseek.com/v1/chat/completions` |
| Moonshot | `https://api.moonshot.cn/v1/chat/completions` |
| 智谱 AI | `https://open.bigmodel.cn/api/paas/v4/chat/completions` |
| Ollama 本地 | `http://localhost:11434/v1/chat/completions` |
| LM Studio | `http://localhost:1234/v1/chat/completions` |
| vLLM | `http://localhost:8000/v1/chat/completions` |

### HTTP Agent 历史管理原理

**关键：客户端维护历史**

```go
type HTTPAgent struct {
    history    map[string][]ChatMessage  // conversationID -> messages
    maxHistory int                        // 默认 20 轮
}
```

**工作流程**：
1. 构建请求时带上历史 (`buildMessages`)
2. 收到回复后保存用户消息 + AI 回复到历史
3. 超过 `maxHistory*2` 时裁剪历史

```go
func (a *HTTPAgent) buildMessages(conversationID string, message string) []ChatMessage {
    var messages []ChatMessage
    // 1. 先加 system prompt
    if a.systemPrompt != "" {
        messages = append(messages, ChatMessage{Role: "system", Content: a.systemPrompt})
    }
    // 2. 加历史对话
    if hist, ok := a.history[conversationID]; ok {
        messages = append(messages, hist...)
    }
    // 3. 加当前消息
    messages = append(messages, ChatMessage{Role: "user", Content: message})
    return messages
}
```

**特点**：
- 多会话隔离 (`map[conversationID][]ChatMessage`)
- 重启后历史清空（内存存储）
- 每次请求带上完整历史（消耗更多 token）

---

## 5. ACP Agent 实现原理

### 架构图

```
┌─────────────────┐                    ┌─────────────────┐
│    WeClaw       │  ──── stdin ────▶  │  claude-agent   │
│   (父进程)      │                    │   (子进程)      │
│                 │  ◀──── stdout ──── │                 │
└─────────────────┘                    └─────────────────┘
        │                                      │
        │         JSON-RPC 2.0 over NDJSON     │
        └──────────────────────────────────────┘
```

### 核心架构

#### 1. 长驻子进程 + 懒加载

```go
func (a *ACPAgent) Start(ctx context.Context) error {
    a.cmd = exec.CommandContext(ctx, a.command, a.args...)
    a.cmd.Dir = a.cwd

    // 创建 stdin/stdout 管道
    a.stdin, _ = a.cmd.StdinPipe()
    stdout, _ := a.cmd.StdoutPipe()

    // 启动子进程
    a.cmd.Start()

    // 启动读取协程
    go a.readLoop()

    // 初始化握手
    a.call(ctx, "initialize", initParams{...})
}
```

- 子进程**只启动一次**，后续请求复用
- 懒加载：首次 `Chat()` 时才启动

#### 2. 双协议支持

| 协议 | 适用 Agent | 会话模型 |
|------|-----------|---------|
| `legacy_acp` | claude-agent-acp, cursor agent | Session 模型 |
| `codex_app_server` | codex app-server | Thread/Turn 模型 |

#### 3. 请求-响应关联 (pending map)

```go
type ACPAgent struct {
    pending   map[int64]chan *rpcResponse  // 请求ID -> 响应channel
    nextID    atomic.Int64                  // 自增ID生成器
}

// 发送请求
id := a.nextID.Add(1)
a.pending[id] = responseCh
a.stdin.Write(request)

// readLoop 收到响应
if msg.ID != nil {
    a.pending[*msg.ID] <- response  // 唤醒等待的调用
}
```

#### 4. 流式响应处理

Agent 的回复是**分块推送**的，通过 `session/update` 通知：

```go
// 注册通知 channel
notifyCh := make(chan *sessionUpdate, 256)
a.notifyCh[sessionID] = notifyCh

// 异步发送 prompt
go func() {
    a.call(ctx, "session/prompt", params)
    promptDone <- struct{}{}
}()

// 收集流式文本块
var textParts []string
for {
    select {
    case update := <-notifyCh:
        if update.SessionUpdate == "agent_message_chunk" {
            textParts = append(textParts, extractChunkText(update))
        }
    case <-promptDone:
        // 排空剩余通知后返回
        return strings.Join(textParts, "")
    }
}
```

**消息流**：
```
WeClaw                          Agent
  │                               │
  │──── session/prompt ──────────▶│
  │                               │
  │◀─── session/update (chunk) ───│  "你"
  │◀─── session/update (chunk) ───│  "好"
  │◀─── session/update (chunk) ───│  "！"
  │      ...                      │
  │◀─── prompt response ──────────│  {stopReason: "end"}
  │                               │
  └── 返回完整文本 ────────────────┘
```

#### 5. 会话管理与隔离

```go
type ACPAgent struct {
    sessions map[string]string  // conversationID -> sessionID (legacy ACP)
    threads  map[string]string  // conversationID -> threadID (codex)
}
```

- 每个微信对话独立 session/thread
- 自动创建，按需复用

#### 6. 自动权限处理

```go
func (a *ACPAgent) handlePermissionRequest(raw string) {
    // 找到 "allow" 选项
    optionID := "allow"
    for _, opt := range req.Params.Options {
        if opt.Kind == "allow" {
            optionID = opt.OptionID
            break
        }
    }

    // 自动发送允许响应
    resp := map[string]interface{}{
        "jsonrpc": "2.0",
        "id":      req.ID,
        "result":  map[string]interface{}{
            "outcome": map[string]interface{}{
                "outcome":  "selected",
                "optionId": optionID,
            },
        },
    }
    a.stdin.Write(resp)
}
```

---

## 6. JSON-RPC 协议详解

### 核心概念

**JSON-RPC** 是轻量级远程过程调用协议，使用 JSON 作为数据格式。

### 消息格式

#### 请求 (Request)
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "session/prompt",
  "params": {"sessionId": "xxx", "prompt": [...]}
}
```

#### 响应 (Response)
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {"stopReason": "end"}
}
```

#### 错误响应
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "error": {
    "code": -32600,
    "message": "Invalid request"
  }
}
```

#### 通知 (Notification)
```json
{
  "jsonrpc": "2.0",
  "method": "session/update",
  "params": {...}
}
```
**没有 id，不需要响应**，用于单向推送（如流式文本块）。

### JSON-RPC vs REST

| 特性 | JSON-RPC | REST |
|------|----------|------|
| **URL** | 单一端点 | 多个资源路径 |
| **HTTP 方法** | 通常 POST | GET/POST/PUT/DELETE |
| **语义** | `method: "createUser"` | `POST /users` |
| **批量请求** | ✅ 原生支持 | ❌ 需自定义 |
| **通知** | ✅ 支持 | ❌ 需 WebSocket |
| **传输层** | 任意 | 通常 HTTP |

### 为什么 ACP 选择 JSON-RPC？

1. **简单** - 只有请求、响应、通知三种消息
2. **灵活** - 不依赖 HTTP，可以用 stdio、socket 等
3. **双向** - 服务端可以主动推送通知
4. **标准化** - 规范明确，易于实现

---

## 7. 微信端透明性

### 完全解耦架构

```
┌─────────────┐         ┌─────────────┐         ┌─────────────┐
│   微信用户   │  ◀────▶ │   WeClaw    │  ◀────▶ │  AI Agent   │
│             │         │   (中间层)   │         │ (任意后端)   │
└─────────────┘         └─────────────┘         └─────────────┘
      │                        │                       │
      │    只看到"机器人"回复    │                       │
      │    不知道背后是什么 AI   │                       │
      └────────────────────────┘                       │
```

### 微信协议层只包含纯文本

```go
SendMessage {
    To: "user_xxx@im.wechat",
    Items: [{
        Type: 1,  // 文本
        Text: "你好！有什么可以帮助你的？"  // 只有纯文本
    }]
}
```

### 设计优势

| 优势 | 说明 |
|------|------|
| **后端无关** | 微信用户无感知，随时切换后端 |
| **多 Agent** | 命令路由 (`/gpt`, `/claude`) |
| **安全性** | 不暴露技术架构 |
| **灵活性** | 可随时添加/移除 Agent |

### 与其他平台对比

| 平台 | 是否显示后端 |
|------|-------------|
| ChatGPT | ✅ 显示 "GPT-4" |
| Claude | ✅ 显示 "Claude 3.5" |
| Poe | ✅ 显示模型名称 |
| **WeClaw** | ❌ **完全隐藏** |

---

## 8. 微信 iLink 协议实现

### 核心文件

| 文件 | 功能 |
|------|------|
| ilink/types.go | 协议数据类型定义 |
| ilink/client.go | API 客户端实现 |
| ilink/auth.go | 登录认证流程 |
| ilink/monitor.go | 消息监听 |
| messaging/handler.go | 消息处理逻辑 |

### API 端点

```go
const defaultBaseURL = "https://ilinkai.weixin.qq.com"

/ilink/bot/get_bot_qrcode    // 获取登录二维码
/ilink/bot/get_qrcode_status // 查询扫码状态
/ilink/bot/getupdates        // 长轮询获取消息 (35秒超时)
/ilink/bot/sendmessage       // 发送消息
/ilink/bot/sendtyping        // 发送输入状态
/ilink/bot/getconfig         // 获取配置（含 typing_ticket）
/ilink/bot/getuploadurl      // 获取 CDN 上传地址
```

### 消息类型处理

```go
// 消息类型
MessageTypeUser = 1   // 用户消息
MessageTypeBot  = 2   // 机器人消息

// 消息状态
MessageStateFinish = 2  // 已完成

// 内容类型
ItemTypeText  = 1   // 文本
ItemTypeImage = 2   // 图片
ItemTypeVoice = 3   // 语音
ItemTypeFile  = 4   // 文件
ItemTypeVideo = 5   // 视频
```

### 各类型处理流程

#### 文本 (ItemTypeText = 1)
- 直接提取 `TextItem.Text`
- 解析命令 (`/gpt`, `@claude` 等)
- 路由到对应 Agent

#### 语音 (ItemTypeVoice = 3)
- **微信服务端已做 ASR 转文字**
- 直接使用 `VoiceItem.Text`
- 无需本地语音识别

#### 图片/文件/视频 (ItemType 2/4/5)
- 优先使用 HTTP URL
- 否则从 CDN 下载 + AES-128-ECB 解密
- 保存到本地后传给 Agent

### CDN 加密通信

```go
// 加密方案
- AES-128-ECB 模式
- PKCS7 填充
- 随机 16 字节 AES 密钥

// 解密流程
cdnURL := "https://novac2c.cdn.weixin.qq.com/c2c/download?encrypted_query_param=xxx"
encryptedData := download(cdnURL)
aesKey := base64Decode(media.AESKey) -> hexDecode()
decrypted := aes128ECBDecrypt(encryptedData, aesKey)

// 解密代码
func decryptAES128ECB(encrypted, key []byte) ([]byte, error) {
    block, _ := aes.NewCipher(key)
    decrypted := make([]byte, len(encrypted))
    for i := 0; i < len(encrypted); i += aes.BlockSize {
        block.Decrypt(decrypted[i:i+aes.BlockSize], encrypted[i:i+aes.BlockSize])
    }
    // PKCS7 去填充
    padding := int(decrypted[len(decrypted)-1])
    return decrypted[:len(decrypted)-padding], nil
}
```

### 与官方协议对齐情况

| 方面 | 状态 | 说明 |
|------|------|------|
| API 端点 | ✅ 完全对齐 | 使用腾讯官方域名 |
| 认证流程 | ✅ 完全对齐 | QRCode → 扫码 → BotToken |
| 消息结构 | ✅ 完全对齐 | `WeixinMessage` 结构完整 |
| 消息类型 | ✅ 5 种全支持 | Text/Image/Voice/File/Video |
| CDN 加密 | ✅ 完全对齐 | AES-128-ECB + PKCS7 |
| 输入状态 | ✅ 支持 | `sendtyping` API |
| 会话管理 | ✅ 支持 | `context_token` 传递 |

### 对比官方 SDK

| 特性 | 官方 OpenClaw | WeClaw |
|------|--------------|--------|
| 语言 | TypeScript | Go |
| Agent 支持 | 仅 Claude | 多 Agent (ACP/CLI/HTTP) |
| 部署 | 需要 Node.js | 单二进制 |
| 消息类型 | 基础 | 完整 (含媒体) |
| CDN 加密 | ❓ | ✅ 完整实现 |

---

## 9. ACP vs HTTP vs CLI 全面对比

| 特性 | ACP | HTTP | CLI |
|------|-----|------|-----|
| **启动方式** | 长驻子进程 | HTTP 请求 | 每次新进程 |
| **通信协议** | stdio + JSON-RPC | REST API | 命令行参数 |
| **会话管理** | Agent 内部 | WeClaw 本地 | --resume 参数 |
| **流式响应** | ✅ 实时推送 | ❌ 批量返回 | ✅ 实时输出 |
| **性能** | ⚡ 最快 | 🚀 快 | 🐢 慢（启动开销） |
| **适用场景** | 本地 Agent | 云端 API | 简单集成 |
| **历史管理** | Agent 内部维护 | WeClaw 内存维护 | 外部文件 |

---

## 10. 命令系统

### 内置命令

```
/info           → 显示当前 Agent 状态
/help           → 显示帮助信息
/new 或 /clear  → 重置会话
/cwd /path      → 切换工作目录
```

### Agent 路由

```
"hello"              → 发送给默认 Agent
"/gpt 你好"          → 发送给 gpt Agent
"@claude @codex 你好" → 广播给多个 Agent
"/claude"            → 切换默认 Agent 为 claude
```

### 内置别名

```go
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
```

---

## 11. 依赖分析

```go
require (
    github.com/google/uuid v1.6.0      // UUID 生成
    github.com/mdp/qrterminal/v3 v3.2.1 // 终端二维码显示
    github.com/spf13/cobra v1.10.2     // CLI 框架
    golang.org/x/net v0.52.0           // HTTP 客户端
    rsc.io/qr v0.2.0                   // QR 码生成
)
```

项目依赖精简，全部使用标准库和少量必要第三方库。

---

## 12. 设计亮点总结

| 亮点 | 说明 |
|------|------|
| **统一接口抽象** | Agent 接口支持插件式扩展 |
| **长驻进程** | ACP 模式避免重复启动，响应最快 |
| **异步 readLoop** | 单协程处理所有响应，避免并发复杂性 |
| **pending map** | 优雅的请求-响应关联，支持并发请求 |
| **notifyCh 分发** | 按会话 ID 路由通知，支持多会话并行 |
| **双协议兼容** | 同时支持 legacy ACP 和 codex app-server |
| **自动权限** | 无感处理工具调用权限，用户体验好 |
| **流式聚合** | 实时收集文本块，最终返回完整响应 |
| **后端无关** | 微信用户无感知，随时切换后端 |
| **零代码接入** | HTTP Agent 纯配置即可接入任意 OpenAI 兼容 API |
| **协议完整** | 完整实现微信 iLink 协议和多种 AI Agent 协议 |
| **安全可靠** | 完整的 CDN 加密通信实现 |
| **运维成熟** | 完善的部署和更新机制 |

---

## 13. 学习价值

WeClaw 是一个优秀的学习案例，涵盖：

1. **Go 并发编程** - goroutine、channel、sync.Map、atomic
2. **进程间通信** - stdio 管道、JSON-RPC 协议
3. **协议实现** - 微信 iLink、OpenAI API 兼容
4. **加密算法** - AES-128-ECB、PKCS7 填充
5. **架构设计** - 接口抽象、插件模式、中间层设计
6. **CLI 开发** - Cobra 框架、系统服务集成

该项目适合作为学习微信机器人开发和 AI Agent 集成的优秀参考案例。

````

[⬆ 回到目录](#toc)

<a name="file-go.mod"></a>
## 📄 go.mod

````text
module github.com/fastclaw-ai/weclaw

go 1.25.0

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mdp/qrterminal/v3 v3.2.1 // indirect
	github.com/spf13/cobra v1.10.2 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	golang.org/x/net v0.52.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/term v0.41.0 // indirect
	rsc.io/qr v0.2.0 // indirect
)

````

[⬆ 回到目录](#toc)

<a name="file-ilink-auth.go"></a>
## 📄 ilink/auth.go

````go
package ilink

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	qrCodeURL     = "https://ilinkai.weixin.qq.com/ilink/bot/get_bot_qrcode?bot_type=3"
	qrStatusURL   = "https://ilinkai.weixin.qq.com/ilink/bot/get_qrcode_status?qrcode="
	statusWait     = "wait"
	statusScanned  = "scaned"
	statusConfirmed = "confirmed"
	statusExpired  = "expired"
)

// FetchQRCode retrieves a new QR code for login.
func FetchQRCode(ctx context.Context) (*QRCodeResponse, error) {
	c := NewUnauthenticatedClient()
	var resp QRCodeResponse
	if err := c.doGet(ctx, qrCodeURL, &resp); err != nil {
		return nil, fmt.Errorf("fetch QR code: %w", err)
	}
	return &resp, nil
}

// PollQRStatus polls for QR code scan status until confirmed or expired.
// It calls onStatus for each status change so the caller can display progress.
func PollQRStatus(ctx context.Context, qrcode string, onStatus func(status string)) (*Credentials, error) {
	c := NewUnauthenticatedClient()
	url := qrStatusURL + qrcode

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		pollCtx, cancel := context.WithTimeout(ctx, 40*time.Second)
		var resp QRStatusResponse
		err := c.doGet(pollCtx, url, &resp)
		cancel()

		if err != nil {
			// Timeout is normal for long-poll, retry
			if ctx.Err() != nil {
				return nil, ctx.Err()
			}
			continue
		}

		if onStatus != nil {
			onStatus(resp.Status)
		}

		switch resp.Status {
		case statusConfirmed:
			creds := &Credentials{
				BotToken:    resp.BotToken,
				ILinkBotID:  resp.ILinkBotID,
				BaseURL:     resp.BaseURL,
				ILinkUserID: resp.ILinkUserID,
			}
			return creds, nil
		case statusExpired:
			return nil, fmt.Errorf("QR code expired")
		case statusWait, statusScanned:
			// Continue polling
		default:
			// Unknown status, continue
		}
	}
}

// AccountsDir returns the directory where account credentials are stored.
func AccountsDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".weclaw", "accounts"), nil
}

// NormalizeAccountID converts raw bot ID to filesystem-safe format.
func NormalizeAccountID(raw string) string {
	s := raw
	for _, ch := range []string{"@", ".", ":"} {
		s = filepath.Clean(s)
		s = replaceAll(s, ch, "-")
	}
	return s
}

func replaceAll(s, old, new string) string {
	for {
		i := indexOf(s, old)
		if i < 0 {
			return s
		}
		s = s[:i] + new + s[i+len(old):]
	}
}

func indexOf(s, sub string) int {
	for i := range s {
		if i+len(sub) <= len(s) && s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}

// SaveCredentials saves credentials to disk under ~/.weclaw/accounts/{id}.json.
func SaveCredentials(creds *Credentials) error {
	dir, err := AccountsDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return fmt.Errorf("create accounts dir: %w", err)
	}

	id := NormalizeAccountID(creds.ILinkBotID)
	path := filepath.Join(dir, id+".json")

	data, err := json.MarshalIndent(creds, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal credentials: %w", err)
	}

	if err := os.WriteFile(path, data, 0o600); err != nil {
		return fmt.Errorf("write credentials: %w", err)
	}
	return nil
}

// LoadAllCredentials loads all saved account credentials.
func LoadAllCredentials() ([]*Credentials, error) {
	dir, err := AccountsDir()
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("read accounts dir: %w", err)
	}

	var result []*Credentials
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".json" {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			continue
		}
		var creds Credentials
		if json.Unmarshal(data, &creds) == nil && creds.BotToken != "" {
			result = append(result, &creds)
		}
	}
	return result, nil
}

// CredentialsPath returns the path for display purposes.
func CredentialsPath() (string, error) {
	return AccountsDir()
}

````

[⬆ 回到目录](#toc)

<a name="file-ilink-client.go"></a>
## 📄 ilink/client.go

````go
package ilink

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	defaultBaseURL     = "https://ilinkai.weixin.qq.com"
	longPollTimeout    = 35 * time.Second
	sendTimeout        = 15 * time.Second
)

// Client is an iLink HTTP API client.
type Client struct {
	baseURL    string
	botToken   string
	botID      string
	httpClient *http.Client
	wechatUIN  string
}

// NewClient creates a new iLink API client.
func NewClient(creds *Credentials) *Client {
	baseURL := creds.BaseURL
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	return &Client{
		baseURL:    baseURL,
		botToken:   creds.BotToken,
		botID:      creds.ILinkBotID,
		httpClient: &http.Client{},
		wechatUIN:  generateWechatUIN(),
	}
}

// NewUnauthenticatedClient creates a client without credentials for login flow.
func NewUnauthenticatedClient() *Client {
	return &Client{
		baseURL:    defaultBaseURL,
		httpClient: &http.Client{Timeout: 40 * time.Second},
		wechatUIN:  generateWechatUIN(),
	}
}

// BotID returns the bot's user ID.
func (c *Client) BotID() string {
	return c.botID
}

// GetUpdates performs a long-poll for new messages.
func (c *Client) GetUpdates(ctx context.Context, buf string) (*GetUpdatesResponse, error) {
	reqBody := GetUpdatesRequest{
		GetUpdatesBuf: buf,
		BaseInfo:      BaseInfo{ChannelVersion: "1.0.0"},
	}

	ctx, cancel := context.WithTimeout(ctx, longPollTimeout+5*time.Second)
	defer cancel()

	var resp GetUpdatesResponse
	if err := c.doPost(ctx, "/ilink/bot/getupdates", reqBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendMessage sends a message through iLink.
func (c *Client) SendMessage(ctx context.Context, msg *SendMessageRequest) (*SendMessageResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, sendTimeout)
	defer cancel()

	var resp SendMessageResponse
	if err := c.doPost(ctx, "/ilink/bot/sendmessage", msg, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetConfig fetches bot config for a user (includes typing_ticket).
func (c *Client) GetConfig(ctx context.Context, userID, contextToken string) (*GetConfigResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	req := GetConfigRequest{
		ILinkUserID:  userID,
		ContextToken: contextToken,
		BaseInfo:     BaseInfo{},
	}

	var resp GetConfigResponse
	if err := c.doPost(ctx, "/ilink/bot/getconfig", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendTyping sends a typing indicator to a user.
func (c *Client) SendTyping(ctx context.Context, userID, typingTicket string, status int) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	req := SendTypingRequest{
		ILinkUserID:  userID,
		TypingTicket: typingTicket,
		Status:       status,
		BaseInfo:     BaseInfo{},
	}

	var resp SendTypingResponse
	if err := c.doPost(ctx, "/ilink/bot/sendtyping", req, &resp); err != nil {
		return err
	}
	if resp.Ret != 0 {
		return fmt.Errorf("sendtyping failed: ret=%d errmsg=%s", resp.Ret, resp.ErrMsg)
	}
	return nil
}

// GetUploadURL gets a pre-signed CDN upload URL for media files.
func (c *Client) GetUploadURL(ctx context.Context, req *GetUploadURLRequest) (*GetUploadURLResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, sendTimeout)
	defer cancel()

	var resp GetUploadURLResponse
	if err := c.doPost(ctx, "/ilink/bot/getuploadurl", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BaseURL returns the base URL for CDN operations.
func (c *Client) BaseURL() string {
	return c.baseURL
}

func (c *Client) doPost(ctx context.Context, path string, body interface{}, result interface{}) error {
	data, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+path, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(respBody))
	}

	if err := json.Unmarshal(respBody, result); err != nil {
		return fmt.Errorf("unmarshal response: %w", err)
	}
	return nil
}

func (c *Client) doGet(ctx context.Context, url string, result interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(respBody))
	}

	if err := json.Unmarshal(respBody, result); err != nil {
		return fmt.Errorf("unmarshal response: %w", err)
	}
	return nil
}

func (c *Client) setHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("AuthorizationType", "ilink_bot_token")
	req.Header.Set("Authorization", "Bearer "+c.botToken)
	req.Header.Set("X-WECHAT-UIN", c.wechatUIN)
}

// SetRequestHeaders sets authentication headers on an HTTP request.
// This can be used for CDN downloads that require authentication.
func (c *Client) SetRequestHeaders(req *http.Request) {
	c.setHeaders(req)
}

func generateWechatUIN() string {
	var n uint32
	_ = binary.Read(rand.Reader, binary.LittleEndian, &n)
	s := fmt.Sprintf("%d", n)
	return base64.StdEncoding.EncodeToString([]byte(s))
}

````

[⬆ 回到目录](#toc)

<a name="file-ilink-monitor.go"></a>
## 📄 ilink/monitor.go

````go
package ilink

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	maxConsecutiveFailures = 5
	initialBackoff         = 3 * time.Second
	maxBackoff             = 60 * time.Second
	sessionExpiredBackoff  = 5 * time.Second
	errCodeSessionExpired  = -14
)

// MessageHandler is called for each received message.
type MessageHandler func(ctx context.Context, client *Client, msg WeixinMessage)

// Monitor manages the long-poll loop for receiving messages.
type Monitor struct {
	client        *Client
	handler       MessageHandler
	getUpdatesBuf string
	bufPath       string
	failures      int
	lastActivity  time.Time
}

// NewMonitor creates a new long-poll monitor.
func NewMonitor(client *Client, handler MessageHandler) (*Monitor, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	accountID := NormalizeAccountID(client.BotID())
	bufPath := filepath.Join(home, ".weclaw", "accounts", accountID+".sync.json")

	m := &Monitor{
		client:       client,
		handler:      handler,
		bufPath:      bufPath,
		lastActivity: time.Now(),
	}
	m.loadBuf()
	return m, nil
}

// Run starts the long-poll loop. It blocks until ctx is cancelled.
// Automatically recovers from errors with exponential backoff.
func (m *Monitor) Run(ctx context.Context) error {
	log.Println("[monitor] starting long-poll loop")

	for {
		select {
		case <-ctx.Done():
			log.Println("[monitor] shutting down")
			return ctx.Err()
		default:
		}

		resp, err := m.client.GetUpdates(ctx, m.getUpdatesBuf)
		if err != nil {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			m.failures++
			backoff := m.calcBackoff()
			log.Printf("[monitor] GetUpdates error (%d/%d, backoff=%s): %v",
				m.failures, maxConsecutiveFailures, backoff, err)
			if m.failures == maxConsecutiveFailures {
				log.Printf("[monitor] WARNING: %d consecutive failures. If this persists, run `weclaw login` to re-authenticate.", maxConsecutiveFailures)
			}
			select {
			case <-time.After(backoff):
			case <-ctx.Done():
				return ctx.Err()
			}
			continue
		}

		// Reset failure counter on any successful response
		m.failures = 0
		m.lastActivity = time.Now()

		// Session expired — reset sync buf and reconnect silently
		if resp.ErrCode == errCodeSessionExpired {
			if m.getUpdatesBuf != "" {
				log.Printf("[monitor] session expired, resetting sync buf")
				m.getUpdatesBuf = ""
				m.saveBuf()
			} else {
				// Sync buf already empty but still getting session expired:
				// the bot token itself has expired. The user needs to re-login.
				log.Printf("[monitor] WARNING: WeChat session expired and cannot be auto-recovered. Run `weclaw login` to re-authenticate.")
			}
			select {
			case <-time.After(sessionExpiredBackoff):
			case <-ctx.Done():
				return ctx.Err()
			}
			continue
		}

		// Other server errors
		if resp.Ret != 0 && resp.ErrCode != 0 {
			log.Printf("[monitor] server error: ret=%d errcode=%d errmsg=%s", resp.Ret, resp.ErrCode, resp.ErrMsg)
			continue
		}

		// Update buf for next poll
		if resp.GetUpdatesBuf != "" {
			m.getUpdatesBuf = resp.GetUpdatesBuf
			m.saveBuf()
		}

		// Process messages concurrently — don't block the poll loop
		for _, msg := range resp.Msgs {
			go m.handler(ctx, m.client, msg)
		}
	}
}

// calcBackoff returns an exponential backoff duration capped at maxBackoff.
func (m *Monitor) calcBackoff() time.Duration {
	d := initialBackoff
	for i := 1; i < m.failures; i++ {
		d *= 2
		if d > maxBackoff {
			return maxBackoff
		}
	}
	return d
}

type syncData struct {
	GetUpdatesBuf string `json:"get_updates_buf"`
}

func (m *Monitor) loadBuf() {
	data, err := os.ReadFile(m.bufPath)
	if err != nil {
		return
	}
	var s syncData
	if json.Unmarshal(data, &s) == nil && s.GetUpdatesBuf != "" {
		m.getUpdatesBuf = s.GetUpdatesBuf
		log.Printf("[monitor] loaded sync buf from %s", m.bufPath)
	}
}

func (m *Monitor) saveBuf() {
	dir := filepath.Dir(m.bufPath)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		log.Printf("[monitor] failed to create buf dir: %v", err)
		return
	}
	data, _ := json.Marshal(syncData{GetUpdatesBuf: m.getUpdatesBuf})
	if err := os.WriteFile(m.bufPath, data, 0o600); err != nil {
		log.Printf("[monitor] failed to save buf: %v", err)
	}
}

// FormatMessageSummary returns a short description of a message for logging.
func FormatMessageSummary(msg WeixinMessage) string {
	text := ""
	for _, item := range msg.ItemList {
		if item.Type == ItemTypeText && item.TextItem != nil {
			text = item.TextItem.Text
			break
		}
	}
	if len(text) > 50 {
		text = text[:50] + "..."
	}
	return fmt.Sprintf("from=%s type=%d state=%d text=%q", msg.FromUserID, msg.MessageType, msg.MessageState, text)
}

````

[⬆ 回到目录](#toc)

<a name="file-ilink-types.go"></a>
## 📄 ilink/types.go

````go
package ilink

// Message types
const (
	MessageTypeNone = 0
	MessageTypeUser = 1
	MessageTypeBot  = 2
)

// Message states
const (
	MessageStateNew        = 0
	MessageStateGenerating = 1
	MessageStateFinish     = 2
)

// Item types
const (
	ItemTypeNone  = 0
	ItemTypeText  = 1
	ItemTypeImage = 2
	ItemTypeVoice = 3
	ItemTypeFile  = 4
	ItemTypeVideo = 5
)

// QRCodeResponse is the response from get_bot_qrcode.
type QRCodeResponse struct {
	QRCode        string `json:"qrcode"`
	QRCodeImgContent string `json:"qrcode_img_content"`
}

// QRStatusResponse is the response from get_qrcode_status.
type QRStatusResponse struct {
	Status     string `json:"status"`
	BotToken   string `json:"bot_token"`
	ILinkBotID string `json:"ilink_bot_id"`
	BaseURL    string `json:"baseurl"`
	ILinkUserID string `json:"ilink_user_id"`
}

// Credentials stores login session data.
type Credentials struct {
	BotToken   string `json:"bot_token"`
	ILinkBotID string `json:"ilink_bot_id"`
	BaseURL    string `json:"baseurl"`
	ILinkUserID string `json:"ilink_user_id"`
}

// BaseInfo is included in request bodies.
type BaseInfo struct {
	ChannelVersion string `json:"channel_version,omitempty"`
}

// GetUpdatesRequest is the body for getupdates.
type GetUpdatesRequest struct {
	GetUpdatesBuf string   `json:"get_updates_buf"`
	BaseInfo      BaseInfo `json:"base_info"`
}

// GetUpdatesResponse is the response from getupdates.
type GetUpdatesResponse struct {
	Ret                 int              `json:"ret"`
	ErrCode             int              `json:"errcode,omitempty"`
	ErrMsg              string           `json:"errmsg,omitempty"`
	Msgs                []WeixinMessage  `json:"msgs"`
	GetUpdatesBuf       string           `json:"get_updates_buf"`
	LongPollingTimeoutMs int             `json:"longpolling_timeout_ms,omitempty"`
}

// WeixinMessage represents a message from WeChat.
type WeixinMessage struct {
	Seq          int           `json:"seq,omitempty"`
	MessageID    int64         `json:"message_id,omitempty"`
	FromUserID   string        `json:"from_user_id"`
	ToUserID     string        `json:"to_user_id"`
	MessageType  int           `json:"message_type"`
	MessageState int           `json:"message_state"`
	ItemList     []MessageItem `json:"item_list"`
	ContextToken string        `json:"context_token"`
}

// MessageItem is a single item in a message.
type MessageItem struct {
	Type      int        `json:"type"`
	TextItem  *TextItem  `json:"text_item,omitempty"`
	ImageItem *ImageItem `json:"image_item,omitempty"`
	VoiceItem *VoiceItem `json:"voice_item,omitempty"`
	VideoItem *VideoItem `json:"video_item,omitempty"`
	FileItem  *FileItem  `json:"file_item,omitempty"`
}

// CDN media type constants.
const (
	CDNMediaTypeImage = 1
	CDNMediaTypeVideo = 2
	CDNMediaTypeFile  = 3
)

// GetUploadURLRequest is the body for getuploadurl.
type GetUploadURLRequest struct {
	FileKey     string   `json:"filekey"`
	MediaType   int      `json:"media_type"`
	ToUserID    string   `json:"to_user_id"`
	RawSize     int      `json:"rawsize"`
	RawFileMD5  string   `json:"rawfilemd5"`
	FileSize    int      `json:"filesize"`
	NoNeedThumb bool     `json:"no_need_thumb"`
	AESKey      string   `json:"aeskey"`
	BaseInfo    BaseInfo `json:"base_info"`
}

// GetUploadURLResponse is the response from getuploadurl.
type GetUploadURLResponse struct {
	Ret           int    `json:"ret"`
	ErrMsg        string `json:"errmsg,omitempty"`
	UploadParam   string `json:"upload_param"`
	UploadFullURL string `json:"upload_full_url,omitempty"`
}

// TextItem holds text content.
type TextItem struct {
	Text string `json:"text"`
}

// MediaInfo holds CDN media reference for uploaded files.
type MediaInfo struct {
	EncryptQueryParam string `json:"encrypt_query_param"`
	AESKey            string `json:"aes_key"`    // base64-encoded
	EncryptType       int    `json:"encrypt_type"` // 1 = AES-128-ECB
}

// VoiceItem holds voice content.
type VoiceItem struct {
	Media         *MediaInfo `json:"media,omitempty"`
	VoiceSize     int        `json:"voice_size,omitempty"`
	EncodeType    int        `json:"encode_type,omitempty"`    // 1=pcm 2=adpcm 3=feature 4=speex 5=amr 6=silk 7=mp3
	BitsPerSample int       `json:"bits_per_sample,omitempty"`
	SampleRate    int        `json:"sample_rate,omitempty"`    // Hz
	Playtime      int        `json:"playtime,omitempty"`       // duration in milliseconds
	Text          string     `json:"text,omitempty"`           // speech-to-text transcription from WeChat
}

// ImageItem holds image content.
type ImageItem struct {
	URL     string     `json:"url,omitempty"`
	Media   *MediaInfo `json:"media,omitempty"`
	MidSize int        `json:"mid_size,omitempty"` // ciphertext size
}

// VideoItem holds video content.
type VideoItem struct {
	Media     *MediaInfo `json:"media,omitempty"`
	VideoSize int        `json:"video_size,omitempty"`
}

// FileItem holds file content.
type FileItem struct {
	Media    *MediaInfo `json:"media,omitempty"`
	FileName string     `json:"file_name,omitempty"`
	Len      string     `json:"len,omitempty"` // plaintext size as string
}

// SendMessageRequest is the body for sendmessage.
type SendMessageRequest struct {
	Msg      SendMsg  `json:"msg"`
	BaseInfo BaseInfo `json:"base_info"`
}

// SendMsg is the message payload for sending.
type SendMsg struct {
	FromUserID   string        `json:"from_user_id"`
	ToUserID     string        `json:"to_user_id"`
	ClientID     string        `json:"client_id"`
	MessageType  int           `json:"message_type"`
	MessageState int           `json:"message_state"`
	ItemList     []MessageItem `json:"item_list"`
	ContextToken string        `json:"context_token"`
}

// SendMessageResponse is the response from sendmessage.
type SendMessageResponse struct {
	Ret    int    `json:"ret"`
	ErrMsg string `json:"errmsg,omitempty"`
}

// Typing status constants.
const (
	TypingStatusTyping = 1
	TypingStatusCancel = 2
)

// GetConfigRequest is the body for getconfig.
type GetConfigRequest struct {
	ILinkUserID  string   `json:"ilink_user_id"`
	ContextToken string   `json:"context_token,omitempty"`
	BaseInfo     BaseInfo `json:"base_info"`
}

// GetConfigResponse is the response from getconfig.
type GetConfigResponse struct {
	Ret           int    `json:"ret"`
	ErrMsg        string `json:"errmsg,omitempty"`
	TypingTicket  string `json:"typing_ticket,omitempty"`
}

// SendTypingRequest is the body for sendtyping.
type SendTypingRequest struct {
	ILinkUserID  string   `json:"ilink_user_id"`
	TypingTicket string   `json:"typing_ticket"`
	Status       int      `json:"status"`
	BaseInfo     BaseInfo `json:"base_info"`
}

// SendTypingResponse is the response from sendtyping.
type SendTypingResponse struct {
	Ret    int    `json:"ret"`
	ErrMsg string `json:"errmsg,omitempty"`
}

````

[⬆ 回到目录](#toc)

<a name="file-install.sh"></a>
## 📄 install.sh

````bash
#!/bin/sh
set -e

REPO="fastclaw-ai/weclaw"
BINARY="weclaw"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

# Detect OS
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
case "$OS" in
  darwin|linux) ;;
  *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

# Detect architecture
ARCH=$(uname -m)
case "$ARCH" in
  x86_64|amd64) ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

echo "Detected: ${OS}/${ARCH}"

# Get latest version
echo "Fetching latest release..."
VERSION=$(curl -fsSL -H "User-Agent: weclaw-installer" "https://api.github.com/repos/${REPO}/releases/latest" | sed -n 's/.*"tag_name" *: *"\([^"]*\)".*/\1/p')

if [ -z "$VERSION" ]; then
  echo "Error: could not determine latest version. Is there a release on GitHub?"
  exit 1
fi

echo "Latest version: ${VERSION}"

# Download
FILENAME="${BINARY}_${OS}_${ARCH}"
URL="https://github.com/${REPO}/releases/download/${VERSION}/${FILENAME}"

echo "Downloading ${URL}..."
TMP=$(mktemp)
curl -fsSL -o "$TMP" "$URL"

# Install
chmod +x "$TMP"
if [ -d "$INSTALL_DIR" ] && [ -w "$INSTALL_DIR" ]; then
  mv "$TMP" "${INSTALL_DIR}/${BINARY}"
else
  echo "Installing to ${INSTALL_DIR} (requires sudo)..."
  sudo mkdir -p "$INSTALL_DIR"
  sudo mv "$TMP" "${INSTALL_DIR}/${BINARY}"
fi

# Clear macOS quarantine attributes
if [ "$OS" = "darwin" ]; then
  xattr -d com.apple.quarantine "${INSTALL_DIR}/${BINARY}" 2>/dev/null || true
  xattr -d com.apple.provenance "${INSTALL_DIR}/${BINARY}" 2>/dev/null || true
fi

echo ""
echo "weclaw ${VERSION} installed to ${INSTALL_DIR}/${BINARY}"
echo ""
echo "Get started:"
echo "  weclaw start"

````

[⬆ 回到目录](#toc)

<a name="file-main.go"></a>
## 📄 main.go

````go
package main

import "github.com/fastclaw-ai/weclaw/cmd"

func main() {
	cmd.Execute()
}

````

[⬆ 回到目录](#toc)

<a name="file-messaging-attachment.go"></a>
## 📄 messaging/attachment.go

````go
package messaging

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
)

var supportedAttachmentExts = []string{
	".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx",
	".zip", ".txt", ".csv",
	".png", ".jpg", ".jpeg", ".gif", ".webp",
	".mp4", ".mov",
}

func defaultAttachmentWorkspace() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return filepath.Clean(os.TempDir())
	}
	return filepath.Join(home, ".weclaw", "workspace")
}

func extractLocalAttachmentPaths(text string) []string {
	var paths []string
	seen := make(map[string]struct{})

	for _, line := range strings.Split(text, "\n") {
		candidate := strings.TrimSpace(line)
		if candidate == "" || !filepath.IsAbs(candidate) {
			continue
		}
		if !isSupportedAttachmentPath(candidate) {
			continue
		}
		info, err := os.Stat(candidate)
		if err != nil || info.IsDir() {
			continue
		}
		if _, ok := seen[candidate]; ok {
			continue
		}
		seen[candidate] = struct{}{}
		paths = append(paths, candidate)
	}

	return paths
}

func isAllowedAttachmentPath(path string, allowedRoots []string) bool {
	cleanPath, err := canonicalizePath(path, true)
	if err != nil {
		return false
	}

	for _, root := range allowedRoots {
		if root == "" {
			continue
		}
		cleanRoot, err := canonicalizePath(root, false)
		if err != nil {
			continue
		}
		rel, err := filepath.Rel(cleanRoot, cleanPath)
		if err != nil {
			continue
		}
		if rel == "." || (rel != ".." && !strings.HasPrefix(rel, ".."+string(os.PathSeparator))) {
			return true
		}
	}

	return false
}

func rewriteReplyWithAttachmentResults(reply string, sentPaths, failedPaths []string) string {
	sentMap := make(map[string]string, len(sentPaths))
	for _, path := range sentPaths {
		sentMap[path] = "已发送附件：" + filepath.Base(path)
	}

	lines := strings.Split(reply, "\n")
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if replacement, ok := sentMap[trimmed]; ok {
			lines[i] = replacement
		}
	}

	rewritten := strings.Join(lines, "\n")

	var failureLines []string
	seenFailures := make(map[string]struct{})
	for _, path := range failedPaths {
		if _, ok := seenFailures[path]; ok {
			continue
		}
		seenFailures[path] = struct{}{}
		failureLines = append(failureLines, "附件发送失败："+filepath.Base(path))
	}
	if len(failureLines) == 0 {
		return rewritten
	}
	if strings.TrimSpace(rewritten) == "" {
		return strings.Join(failureLines, "\n")
	}
	return rewritten + "\n" + strings.Join(failureLines, "\n")
}

func isSupportedAttachmentPath(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return slices.Contains(supportedAttachmentExts, ext)
}

func canonicalizePath(path string, mustExist bool) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	if realPath, err := filepath.EvalSymlinks(absPath); err == nil {
		return filepath.Clean(realPath), nil
	} else if mustExist {
		return "", err
	}
	return filepath.Clean(absPath), nil
}

````

[⬆ 回到目录](#toc)

<a name="file-messaging-attachment_test.go"></a>
## 📄 messaging/attachment_test.go

````go
package messaging

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExtractLocalAttachmentPaths(t *testing.T) {
	dir := t.TempDir()
	pdfPath := filepath.Join(dir, "report.pdf")
	txtPath := filepath.Join(dir, "notes.txt")
	if err := os.WriteFile(pdfPath, []byte("pdf"), 0o644); err != nil {
		t.Fatalf("write pdf: %v", err)
	}
	if err := os.WriteFile(txtPath, []byte("txt"), 0o644); err != nil {
		t.Fatalf("write txt: %v", err)
	}

	reply := strings.Join([]string{
		"这里是内联路径，不应该命中 " + pdfPath,
		pdfPath,
		"1. " + txtPath,
		txtPath,
		"file://" + pdfPath,
		filepath.Join(dir, "missing.pdf"),
		filepath.Join(dir, "folder"),
	}, "\n")

	got := extractLocalAttachmentPaths(reply)
	if len(got) != 2 {
		t.Fatalf("expected 2 paths, got %d (%v)", len(got), got)
	}
	if got[0] != pdfPath {
		t.Fatalf("got[0] = %q, want %q", got[0], pdfPath)
	}
	if got[1] != txtPath {
		t.Fatalf("got[1] = %q, want %q", got[1], txtPath)
	}
}

func TestIsAllowedAttachmentPath(t *testing.T) {
	workspaceRoot := filepath.Join(t.TempDir(), "workspace")
	otherRoot := filepath.Join(t.TempDir(), "other")
	if err := os.MkdirAll(workspaceRoot, 0o755); err != nil {
		t.Fatalf("mkdir workspace: %v", err)
	}
	if err := os.MkdirAll(otherRoot, 0o755); err != nil {
		t.Fatalf("mkdir other: %v", err)
	}

	allowedPath := filepath.Join(workspaceRoot, "artifacts", "report.pdf")
	deniedPath := filepath.Join(otherRoot, "report.pdf")
	if err := os.MkdirAll(filepath.Dir(allowedPath), 0o755); err != nil {
		t.Fatalf("mkdir allowed dir: %v", err)
	}
	if err := os.WriteFile(allowedPath, []byte("ok"), 0o644); err != nil {
		t.Fatalf("write allowed file: %v", err)
	}
	if err := os.WriteFile(deniedPath, []byte("no"), 0o644); err != nil {
		t.Fatalf("write denied file: %v", err)
	}

	if !isAllowedAttachmentPath(allowedPath, []string{workspaceRoot}) {
		t.Fatalf("expected %q to be allowed", allowedPath)
	}
	if isAllowedAttachmentPath(deniedPath, []string{workspaceRoot}) {
		t.Fatalf("expected %q to be denied", deniedPath)
	}
}

func TestRewriteReplyWithAttachmentResults(t *testing.T) {
	sentPath := "/tmp/report.pdf"
	failedPath := "/tmp/archive.zip"
	reply := strings.Join([]string{
		"已生成文件：",
		sentPath,
		"这里再次内联提到 " + sentPath + "，不应该被替换。",
		failedPath,
	}, "\n")

	got := rewriteReplyWithAttachmentResults(reply, []string{sentPath}, []string{failedPath})

	if strings.Contains(got, "\n"+sentPath+"\n") {
		t.Fatalf("expected sent path line to be replaced, got %q", got)
	}
	if !strings.Contains(got, "已发送附件：report.pdf") {
		t.Fatalf("expected sent replacement, got %q", got)
	}
	if !strings.Contains(got, "这里再次内联提到 "+sentPath+"，不应该被替换。") {
		t.Fatalf("expected inline path to remain, got %q", got)
	}
	if !strings.Contains(got, failedPath) {
		t.Fatalf("expected failed path to remain, got %q", got)
	}
	if !strings.Contains(got, "附件发送失败：archive.zip") {
		t.Fatalf("expected failure note, got %q", got)
	}
}

````

[⬆ 回到目录](#toc)

<a name="file-messaging-cdn.go"></a>
## 📄 messaging/cdn.go

````go
package messaging

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/fastclaw-ai/weclaw/ilink"
)

const cdnBaseURL = "https://novac2c.cdn.weixin.qq.com/c2c"

// UploadedFile holds the result of a CDN upload.
type UploadedFile struct {
	DownloadParam string // encrypted query param for download
	AESKeyHex     string // hex-encoded AES key
	FileSize      int    // plaintext size
	CipherSize    int    // ciphertext size
}

// UploadFileToCDN encrypts and uploads a file to the WeChat CDN.
func UploadFileToCDN(ctx context.Context, client *ilink.Client, data []byte, toUserID string, mediaType int) (*UploadedFile, error) {
	// Generate random filekey and AES key
	filekey := make([]byte, 16)
	aeskey := make([]byte, 16)
	if _, err := rand.Read(filekey); err != nil {
		return nil, fmt.Errorf("generate filekey: %w", err)
	}
	if _, err := rand.Read(aeskey); err != nil {
		return nil, fmt.Errorf("generate aeskey: %w", err)
	}

	filekeyHex := hex.EncodeToString(filekey)
	aeskeyHex := hex.EncodeToString(aeskey)

	// Calculate MD5 of plaintext
	hash := md5.Sum(data)
	rawMD5 := hex.EncodeToString(hash[:])

	// Calculate ciphertext size (PKCS7 padding)
	cipherSize := aesECBPaddedSize(len(data))

	// Get upload URL from iLink API
	uploadReq := &ilink.GetUploadURLRequest{
		FileKey:     filekeyHex,
		MediaType:   mediaType,
		ToUserID:    toUserID,
		RawSize:     len(data),
		RawFileMD5:  rawMD5,
		FileSize:    cipherSize,
		NoNeedThumb: true,
		AESKey:      aeskeyHex,
		BaseInfo:    ilink.BaseInfo{},
	}

	uploadResp, err := client.GetUploadURL(ctx, uploadReq)
	if err != nil {
		return nil, fmt.Errorf("get upload URL: %w", err)
	}
	if uploadResp.Ret != 0 {
		return nil, fmt.Errorf("get upload URL failed: ret=%d errmsg=%s", uploadResp.Ret, uploadResp.ErrMsg)
	}

	// Encrypt data with AES-128-ECB
	encrypted, err := encryptAESECB(data, aeskey)
	if err != nil {
		return nil, fmt.Errorf("encrypt: %w", err)
	}

	// Upload to CDN: prefer server-provided full URL, fall back to param-based construction
	cdnURL := strings.TrimSpace(uploadResp.UploadFullURL)
	if cdnURL == "" {
		if uploadResp.UploadParam == "" {
			return nil, fmt.Errorf("getuploadurl returned no upload URL (need upload_full_url or upload_param)")
		}
		cdnURL = fmt.Sprintf("%s/upload?encrypted_query_param=%s&filekey=%s",
			cdnBaseURL, url.QueryEscape(uploadResp.UploadParam), url.QueryEscape(filekeyHex))
	}

	downloadParam, err := uploadToCDN(ctx, encrypted, cdnURL)
	if err != nil {
		return nil, fmt.Errorf("CDN upload: %w", err)
	}

	return &UploadedFile{
		DownloadParam: downloadParam,
		AESKeyHex:     aeskeyHex,
		FileSize:      len(data),
		CipherSize:    cipherSize,
	}, nil
}

// AESKeyToBase64 converts a hex AES key to base64 format for message items.
func AESKeyToBase64(hexKey string) string {
	return base64.StdEncoding.EncodeToString([]byte(hexKey))
}

// DownloadFileFromCDN downloads and decrypts a file from the WeChat CDN.
func DownloadFileFromCDN(ctx context.Context, encryptQueryParam, aesKeyBase64 string) ([]byte, error) {
	// Decode AES key: base64 -> hex string -> raw bytes
	aesKeyHexBytes, err := base64.StdEncoding.DecodeString(aesKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("decode AES key base64: %w", err)
	}
	aesKey, err := hex.DecodeString(string(aesKeyHexBytes))
	if err != nil {
		return nil, fmt.Errorf("decode AES key hex: %w", err)
	}

	// Download encrypted data from CDN
	downloadURL := fmt.Sprintf("%s/download?encrypted_query_param=%s",
		cdnBaseURL, url.QueryEscape(encryptQueryParam))

	reqCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, downloadURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create download request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("download from CDN: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("CDN download HTTP %d: %s", resp.StatusCode, string(body))
	}

	encrypted, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read CDN response: %w", err)
	}

	// Decrypt AES-128-ECB
	return decryptAESECB(encrypted, aesKey)
}

// decryptAESECB decrypts data encrypted with AES-128-ECB and removes PKCS7 padding.
func decryptAESECB(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of block size")
	}

	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += aes.BlockSize {
		block.Decrypt(plaintext[i:i+aes.BlockSize], ciphertext[i:i+aes.BlockSize])
	}

	// Remove PKCS7 padding
	if len(plaintext) == 0 {
		return plaintext, nil
	}
	padLen := int(plaintext[len(plaintext)-1])
	if padLen > aes.BlockSize || padLen == 0 {
		return nil, fmt.Errorf("invalid PKCS7 padding")
	}
	return plaintext[:len(plaintext)-padLen], nil
}

func uploadToCDN(ctx context.Context, encrypted []byte, cdnURL string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, cdnURL, bytes.NewReader(encrypted))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("CDN upload HTTP %d: %s", resp.StatusCode, string(body))
	}

	downloadParam := resp.Header.Get("X-Encrypted-Param")
	if downloadParam == "" {
		return "", fmt.Errorf("CDN upload: missing X-Encrypted-Param header")
	}

	return downloadParam, nil
}

// encryptAESECB encrypts data using AES-128-ECB with PKCS7 padding.
func encryptAESECB(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// PKCS7 padding
	padLen := aes.BlockSize - (len(plaintext) % aes.BlockSize)
	padded := make([]byte, len(plaintext)+padLen)
	copy(padded, plaintext)
	for i := len(plaintext); i < len(padded); i++ {
		padded[i] = byte(padLen)
	}

	// ECB mode: encrypt each block independently
	encrypted := make([]byte, len(padded))
	for i := 0; i < len(padded); i += aes.BlockSize {
		block.Encrypt(encrypted[i:i+aes.BlockSize], padded[i:i+aes.BlockSize])
	}

	return encrypted, nil
}

func aesECBPaddedSize(plaintextSize int) int {
	return (plaintextSize/aes.BlockSize + 1) * aes.BlockSize
}

````

[⬆ 回到目录](#toc)

<a name="file-messaging-handler.go"></a>
## 📄 messaging/handler.go

````go
package messaging

import (
	"context"
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fastclaw-ai/weclaw/agent"
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
	mu            sync.RWMutex
	defaultName   string
	agents        map[string]agent.Agent // name -> running agent
	agentMetas    []AgentMeta            // all configured agents (for /status)
	agentWorkDirs map[string]string      // agent name -> configured/runtime cwd
	customAliases map[string]string      // custom alias -> agent name (from config)
	factory       AgentFactory
	saveDefault   SaveDefaultFunc
	contextTokens sync.Map   // map[userID]contextToken
	saveDir       string     // directory to save images/files to
	seenMsgs      sync.Map   // map[int64]time.Time — dedup by message_id
}

// NewHandler creates a new message handler.
func NewHandler(factory AgentFactory, saveDefault SaveDefaultFunc) *Handler {
	return &Handler{
		agents:        make(map[string]agent.Agent),
		agentWorkDirs: make(map[string]string),
		factory:       factory,
		saveDefault:   saveDefault,
	}
}

// SetSaveDir sets the directory for saving images and files.
func (h *Handler) SetSaveDir(dir string) {
	h.saveDir = dir
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
func (h *Handler) resolveAlias(name string) string {
	h.mu.RLock()
	custom := h.customAliases
	h.mu.RUnlock()
	if custom != nil {
		if full, ok := custom[name]; ok {
			return full
		}
	}
	if full, ok := agentAliases[name]; ok {
		return full
	}
	return name
}

// parseCommand checks if text starts with "/" or "@" followed by agent name(s).
// Supports multiple agents: "@cc @cx hello" returns (["claude","codex"], "hello").
// Returns (agentNames, actualMessage). Aliases are resolved automatically.
// If no command prefix, returns (nil, originalText).
func (h *Handler) parseCommand(text string) ([]string, string) {
	if !strings.HasPrefix(text, "/") && !strings.HasPrefix(text, "@") {
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

	// Intercept URLs: save to Linkhoard directly without AI agent
	trimmed := strings.TrimSpace(text)
	if h.saveDir != "" && IsURL(trimmed) {
		rawURL := ExtractURL(trimmed)
		if rawURL != "" {
			log.Printf("[handler] saving URL to linkhoard: %s", rawURL)
			title, err := SaveLinkToLinkhoard(ctx, h.saveDir, rawURL)
			var reply string
			if err != nil {
				log.Printf("[handler] link save failed: %v", err)
				reply = fmt.Sprintf("保存失败: %v", err)
			} else {
				reply = fmt.Sprintf("已保存: %s", title)
			}
			if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
				log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
			}
			return
		}
	}

	// Built-in commands (no typing needed)
	if trimmed == "/info" {
		reply := h.buildStatus()
		if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
			log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
		}
		return
	} else if trimmed == "/help" {
		reply := buildHelpText()
		if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
			log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
		}
		return
	} else if trimmed == "/new" || trimmed == "/clear" {
		reply := h.resetDefaultSession(ctx, msg.FromUserID)
		if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
			log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
		}
		return
	} else if strings.HasPrefix(trimmed, "/cwd") {
		reply := h.handleCwd(trimmed)
		if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
			log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
		}
		return
	}

	// Route: "/agentname message" or "@agent1 @agent2 message" -> specific agent(s)
	agentNames, message := h.parseCommand(text)

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
		// No known agents -> forward entire text to default agent
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
		reply, err = h.chatWithAgent(ctx, ag, msg.FromUserID, text)
		if err != nil {
			reply = fmt.Sprintf("Error: %v", err)
		}
	} else {
		log.Printf("[handler] agent not ready, using echo mode for %s", msg.FromUserID)
		reply = "[echo] " + text
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

	reply, err := h.chatWithAgent(ctx, ag, msg.FromUserID, message)
	if err != nil {
		reply = fmt.Sprintf("Error: %v", err)
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
			reply, err := h.chatWithAgent(ctx, ag, msg.FromUserID, message)
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
func (h *Handler) chatWithAgent(ctx context.Context, ag agent.Agent, userID, message string) (string, error) {
	info := ag.Info()
	log.Printf("[handler] dispatching to agent (%s) for %s", info, userID)

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

func buildHelpText() string {
	return `Available commands:
@agent or /agent - Switch default agent
@agent msg or /agent msg - Send to a specific agent
@a @b msg - Broadcast to multiple agents
/new or /clear - Start a new session
/cwd /path - Switch workspace directory
/info - Show current agent info
/help - Show this help message

Aliases: /cc(claude) /cx(codex) /cs(cursor) /km(kimi) /gm(gemini) /oc(openclaw) /ocd(opencode) /pi(pi) /cp(copilot) /dr(droid) /if(iflow) /kr(kiro) /qw(qwen)`
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
		// No encryption key — data is plaintext
		fileData = encryptedData
		log.Printf("[handler] no AES key, using raw data (no decryption)")
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

````

[⬆ 回到目录](#toc)

<a name="file-messaging-handler_test.go"></a>
## 📄 messaging/handler_test.go

````go
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

````

[⬆ 回到目录](#toc)

<a name="file-messaging-linkhoard.go"></a>
## 📄 messaging/linkhoard.go

````go
package messaging

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"bufio"

	"github.com/google/uuid"
	"golang.org/x/net/html"
)

var reURL = regexp.MustCompile(`https?://\S+`)

// IsURL checks if the text is (or starts with) a URL.
func IsURL(text string) bool {
	trimmed := strings.TrimSpace(text)
	return strings.HasPrefix(trimmed, "http://") || strings.HasPrefix(trimmed, "https://")
}

// ExtractURL extracts the first URL from text.
func ExtractURL(text string) string {
	match := reURL.FindString(text)
	return match
}

// LinkMetadata holds extracted metadata from a web page.
type LinkMetadata struct {
	Title       string
	Description string
	Author      string
	OGImage     string
	Published   string
	Body        string
}

// FetchLinkMetadata fetches a URL and extracts metadata from the HTML.
func FetchLinkMetadata(ctx context.Context, rawURL string) (*LinkMetadata, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")
	req.Header.Set("Referer", "https://mp.weixin.qq.com/")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parse HTML: %w", err)
	}

	meta := &LinkMetadata{}
	extractMeta(doc, meta)

	// Fallback title from URL if empty
	if meta.Title == "" {
		meta.Title = rawURL
	}

	return meta, nil
}

// extractMeta walks the HTML tree and extracts metadata.
func extractMeta(n *html.Node, meta *LinkMetadata) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "meta":
			handleMeta(n, meta)
		case "title":
			if meta.Title == "" && n.FirstChild != nil {
				meta.Title = strings.TrimSpace(n.FirstChild.Data)
			}
		case "div":
			// WeChat article body
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == "js_content" {
					meta.Body = extractNodeText(n)
					return
				}
			}
		case "article":
			if meta.Body == "" {
				meta.Body = extractNodeText(n)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractMeta(c, meta)
	}
}

// handleMeta extracts og: and other meta tag values.
func handleMeta(n *html.Node, meta *LinkMetadata) {
	var property, name, content string
	for _, a := range n.Attr {
		switch a.Key {
		case "property":
			property = a.Val
		case "name":
			name = a.Val
		case "content":
			content = a.Val
		}
	}
	if content == "" {
		return
	}
	switch {
	case property == "og:title" && meta.Title == "":
		meta.Title = content
	case property == "og:description" && meta.Description == "":
		meta.Description = content
	case property == "og:image" && meta.OGImage == "":
		meta.OGImage = content
	case property == "article:published_time" && meta.Published == "":
		meta.Published = content
	case name == "author" && meta.Author == "":
		meta.Author = content
	case name == "description" && meta.Description == "":
		meta.Description = content
	}
}

// extractText recursively extracts visible text from an HTML node.
func extractNodeText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var sb strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && (c.Data == "script" || c.Data == "style") {
			continue
		}
		text := extractNodeText(c)
		if text != "" {
			// Add paragraph breaks for block elements
			if c.Type == html.ElementNode {
				switch c.Data {
				case "p", "div", "br", "h1", "h2", "h3", "h4", "h5", "h6", "li", "section":
					sb.WriteString("\n\n")
				}
			}
			sb.WriteString(text)
		}
	}
	return sb.String()
}

// sanitizeFileName removes characters unsafe for filenames.
func sanitizeFileName(name string) string {
	replacer := strings.NewReplacer(
		"/", "", "\\", "", ":", "", "*", "",
		"?", "", "\"", "", "<", "", ">", "", "|", "",
	)
	result := replacer.Replace(name)
	// Trim and limit length
	result = strings.TrimSpace(result)
	if len(result) > 200 {
		result = result[:200]
	}
	if result == "" {
		result = "untitled"
	}
	return result
}

// isWeChatURL checks if a URL is a WeChat article.
func isWeChatURL(rawURL string) bool {
	return strings.Contains(rawURL, "mp.weixin.qq.com") || strings.Contains(rawURL, "weixin.qq.com/s/")
}

// FetchViaJina fetches a URL via Jina Reader API and returns metadata + markdown body.
func FetchViaJina(ctx context.Context, rawURL string) (*LinkMetadata, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	jinaURL := "https://r.jina.ai/" + rawURL
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, jinaURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "text/plain")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Jina HTTP %d", resp.StatusCode)
	}

	meta := &LinkMetadata{}
	scanner := bufio.NewScanner(resp.Body)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	// Parse Jina header lines: "Title:", "URL Source:", "Published Time:", then "Markdown Content:"
	inBody := false
	var body strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		if inBody {
			body.WriteString(line)
			body.WriteString("\n")
			continue
		}
		if strings.HasPrefix(line, "Title: ") {
			meta.Title = strings.TrimPrefix(line, "Title: ")
		} else if strings.HasPrefix(line, "Published Time: ") {
			meta.Published = strings.TrimPrefix(line, "Published Time: ")
		} else if line == "Markdown Content:" {
			inBody = true
		}
	}

	if meta.Title == "" {
		meta.Title = rawURL
	}
	meta.Body = strings.TrimSpace(body.String())

	// Check for Jina failure (CAPTCHA, empty content)
	if meta.Body == "" || strings.Contains(meta.Body, "环境异常") || strings.Contains(meta.Body, "CAPTCHA") {
		return nil, fmt.Errorf("Jina returned empty or blocked content")
	}

	return meta, nil
}

// SaveLinkToLinkhoard fetches a URL and saves it as a Linkhoard-compatible markdown file.
// WeChat articles use direct fetch with browser headers; other sites use Jina Reader.
func SaveLinkToLinkhoard(ctx context.Context, saveDir, rawURL string) (string, error) {
	var meta *LinkMetadata
	var err error

	if isWeChatURL(rawURL) {
		meta, err = FetchLinkMetadata(ctx, rawURL)
	} else {
		meta, err = FetchViaJina(ctx, rawURL)
		if err != nil {
			// Fallback to direct fetch
			log.Printf("[linkhoard] Jina failed (%v), falling back to direct fetch", err)
			meta, err = FetchLinkMetadata(ctx, rawURL)
		}
	}
	if err != nil {
		return "", fmt.Errorf("fetch failed: %w", err)
	}

	// Ensure save directory exists
	if err := os.MkdirAll(saveDir, 0o755); err != nil {
		return "", fmt.Errorf("create dir: %w", err)
	}

	// Build frontmatter
	title := sanitizeFileName(meta.Title)
	created := time.Now().UTC().Format(time.RFC3339)
	itemID := uuid.New().String()

	// Normalize body text
	body := strings.TrimSpace(meta.Body)
	// Collapse excessive newlines
	for strings.Contains(body, "\n\n\n") {
		body = strings.ReplaceAll(body, "\n\n\n", "\n\n")
	}

	// Build author field
	authorField := "author: []\n"
	if meta.Author != "" {
		authorField = fmt.Sprintf("author:\n  - '[[%s]]'\n", meta.Author)
	}

	// Build markdown content
	var sb strings.Builder
	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("title: '%s'\n", strings.ReplaceAll(meta.Title, "'", "''")))
	sb.WriteString(fmt.Sprintf("source: '%s'\n", rawURL))
	sb.WriteString(fmt.Sprintf("published: '%s'\n", meta.Published))
	sb.WriteString(fmt.Sprintf("created: '%s'\n", created))
	sb.WriteString(fmt.Sprintf("description: '%s'\n", strings.ReplaceAll(meta.Description, "'", "''")))
	if meta.OGImage != "" {
		sb.WriteString(fmt.Sprintf("openGraphImage: '%s'\n", meta.OGImage))
	}
	sb.WriteString(authorField)
	sb.WriteString("---\n\n")
	if body != "" {
		sb.WriteString(body)
		sb.WriteString("\n")
	}

	// Write markdown file
	filePath := filepath.Join(saveDir, title+".md")
	if err := os.WriteFile(filePath, []byte(sb.String()), 0o644); err != nil {
		return "", fmt.Errorf("write file: %w", err)
	}

	// Write sidecar
	sidecarPath := filePath + ".sidecar.md"
	sidecarContent := fmt.Sprintf("---\nid: %s\n---\n", itemID)
	if err := os.WriteFile(sidecarPath, []byte(sidecarContent), 0o644); err != nil {
		log.Printf("[linkhoard] failed to write sidecar: %v", err)
	}

	log.Printf("[linkhoard] saved %q to %s", meta.Title, filePath)
	return meta.Title, nil
}

````

[⬆ 回到目录](#toc)

<a name="file-messaging-markdown.go"></a>
## 📄 messaging/markdown.go

````go
package messaging

import (
	"regexp"
	"strings"
)

var (
	// Code blocks: strip fences, keep code content
	reCodeBlock = regexp.MustCompile("(?s)```[^\n]*\n?(.*?)```")
	// Inline code: strip backticks, keep content
	reInlineCode = regexp.MustCompile("`([^`]+)`")
	// Images: remove entirely
	reImage = regexp.MustCompile(`!\[[^\]]*\]\([^)]*\)`)
	// Links: keep display text only
	reLink = regexp.MustCompile(`\[([^\]]+)\]\([^)]*\)`)
	// Table separator rows: remove
	reTableSep = regexp.MustCompile(`(?m)^\|[\s:|\-]+\|$`)
	// Table rows: convert pipe-delimited to space-delimited
	reTableRow = regexp.MustCompile(`(?m)^\|(.+)\|$`)
	// Headers: remove # prefix
	reHeader = regexp.MustCompile(`(?m)^#{1,6}\s+`)
	// Bold: **text** or __text__
	reBold = regexp.MustCompile(`\*\*(.+?)\*\*|__(.+?)__`)
	// Italic: *text* or _text_
	reItalic = regexp.MustCompile(`(?:^|[^*])\*([^*]+)\*(?:[^*]|$)|(?:^|[^_])_([^_]+)_(?:[^_]|$)`)
	// Strikethrough: ~~text~~
	reStrike = regexp.MustCompile(`~~(.+?)~~`)
	// Blockquote: > prefix
	reBlockquote = regexp.MustCompile(`(?m)^>\s?`)
	// Horizontal rule
	reHR = regexp.MustCompile(`(?m)^[-*_]{3,}\s*$`)
	// Unordered list markers: -, *, +
	reUL = regexp.MustCompile(`(?m)^(\s*)[-*+]\s+`)
)

// MarkdownToPlainText converts markdown to readable plain text for WeChat.
func MarkdownToPlainText(text string) string {
	result := text

	// Code blocks: strip fences, keep code content
	result = reCodeBlock.ReplaceAllStringFunc(result, func(match string) string {
		parts := reCodeBlock.FindStringSubmatch(match)
		if len(parts) > 1 {
			return strings.TrimSpace(parts[1])
		}
		return match
	})

	// Images: remove entirely
	result = reImage.ReplaceAllString(result, "")

	// Links: keep display text only
	result = reLink.ReplaceAllString(result, "$1")

	// Table separator rows: remove
	result = reTableSep.ReplaceAllString(result, "")

	// Table rows: pipe-delimited to space-delimited
	result = reTableRow.ReplaceAllStringFunc(result, func(match string) string {
		parts := reTableRow.FindStringSubmatch(match)
		if len(parts) > 1 {
			cells := strings.Split(parts[1], "|")
			for i := range cells {
				cells[i] = strings.TrimSpace(cells[i])
			}
			return strings.Join(cells, "  ")
		}
		return match
	})

	// Headers: remove # prefix
	result = reHeader.ReplaceAllString(result, "")

	// Bold
	result = reBold.ReplaceAllStringFunc(result, func(match string) string {
		parts := reBold.FindStringSubmatch(match)
		if parts[1] != "" {
			return parts[1]
		}
		return parts[2]
	})

	// Strikethrough
	result = reStrike.ReplaceAllString(result, "$1")

	// Blockquote
	result = reBlockquote.ReplaceAllString(result, "")

	// Horizontal rule -> empty line
	result = reHR.ReplaceAllString(result, "")

	// Unordered list: replace markers with "• "
	result = reUL.ReplaceAllString(result, "${1}• ")

	// Inline code: strip backticks (do after code blocks)
	result = reInlineCode.ReplaceAllString(result, "$1")

	// Clean up excessive blank lines
	result = regexp.MustCompile(`\n{3,}`).ReplaceAllString(result, "\n\n")

	return strings.TrimSpace(result)
}

````

[⬆ 回到目录](#toc)

<a name="file-messaging-media.go"></a>
## 📄 messaging/media.go

````go
package messaging

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/fastclaw-ai/weclaw/ilink"
)

// reMarkdownImage matches markdown image syntax: ![alt](url)
var reMarkdownImage = regexp.MustCompile(`!\[[^\]]*\]\(([^)]+)\)`)

// ExtractImageURLs extracts image URLs from markdown text.
func ExtractImageURLs(text string) []string {
	matches := reMarkdownImage.FindAllStringSubmatch(text, -1)
	var urls []string
	for _, m := range matches {
		url := strings.TrimSpace(m[1])
		if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
			urls = append(urls, url)
		}
	}
	return urls
}

// SendMediaFromURL downloads a file from a URL and sends it as a media message.
func SendMediaFromURL(ctx context.Context, client *ilink.Client, toUserID, mediaURL, contextToken string) error {
	data, contentType, err := downloadFile(ctx, mediaURL)
	if err != nil {
		return fmt.Errorf("download %s: %w", mediaURL, err)
	}

	return sendMediaData(ctx, client, toUserID, filenameFromURL(mediaURL), mediaURL, data, contentType, contextToken)
}

// SendMediaFromPath reads a local file and sends it as a media message.
func SendMediaFromPath(ctx context.Context, client *ilink.Client, toUserID, path, contextToken string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}

	return sendMediaData(ctx, client, toUserID, filepath.Base(path), path, data, inferContentType(path), contextToken)
}

func sendMediaData(ctx context.Context, client *ilink.Client, toUserID, fileName, source string, data []byte, contentType, contextToken string) error {
	if fileName == "" {
		fileName = "file"
	}

	cdnMediaType, itemType := classifyMedia(contentType, source)

	log.Printf("[media] uploading %s (%s, %d bytes) for %s", source, contentType, len(data), toUserID)

	uploaded, err := UploadFileToCDN(ctx, client, data, toUserID, cdnMediaType)
	if err != nil {
		return fmt.Errorf("upload to CDN: %w", err)
	}

	media := &ilink.MediaInfo{
		EncryptQueryParam: uploaded.DownloadParam,
		AESKey:            AESKeyToBase64(uploaded.AESKeyHex),
		EncryptType:       1,
	}

	var item ilink.MessageItem
	switch itemType {
	case ilink.ItemTypeImage:
		item = ilink.MessageItem{
			Type: ilink.ItemTypeImage,
			ImageItem: &ilink.ImageItem{
				Media:   media,
				MidSize: uploaded.CipherSize,
			},
		}
	case ilink.ItemTypeVideo:
		item = ilink.MessageItem{
			Type: ilink.ItemTypeVideo,
			VideoItem: &ilink.VideoItem{
				Media:     media,
				VideoSize: uploaded.CipherSize,
			},
		}
	default:
		item = ilink.MessageItem{
			Type: ilink.ItemTypeFile,
			FileItem: &ilink.FileItem{
				Media:    media,
				FileName: fileName,
				Len:      fmt.Sprintf("%d", uploaded.FileSize),
			},
		}
	}

	req := &ilink.SendMessageRequest{
		Msg: ilink.SendMsg{
			FromUserID:   client.BotID(),
			ToUserID:     toUserID,
			ClientID:     NewClientID(),
			MessageType:  ilink.MessageTypeBot,
			MessageState: ilink.MessageStateFinish,
			ItemList:     []ilink.MessageItem{item},
			ContextToken: contextToken,
		},
		BaseInfo: ilink.BaseInfo{},
	}

	resp, err := client.SendMessage(ctx, req)
	if err != nil {
		return fmt.Errorf("send media message: %w", err)
	}
	if resp.Ret != 0 {
		return fmt.Errorf("send media failed: ret=%d errmsg=%s", resp.Ret, resp.ErrMsg)
	}

	log.Printf("[media] sent %s to %s from %s", contentType, toUserID, source)
	return nil
}

func downloadFile(ctx context.Context, url string) ([]byte, string, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		contentType = inferContentType(url)
	}

	return data, contentType, nil
}

func classifyMedia(contentType, url string) (cdnMediaType int, itemType int) {
	ct := strings.ToLower(contentType)

	if strings.HasPrefix(ct, "image/") || isImageExt(url) {
		return ilink.CDNMediaTypeImage, ilink.ItemTypeImage
	}
	if strings.HasPrefix(ct, "video/") || isVideoExt(url) {
		return ilink.CDNMediaTypeVideo, ilink.ItemTypeVideo
	}
	return ilink.CDNMediaTypeFile, ilink.ItemTypeFile
}

func isImageExt(url string) bool {
	ext := strings.ToLower(filepath.Ext(stripQuery(url)))
	switch ext {
	case ".png", ".jpg", ".jpeg", ".gif", ".webp", ".bmp":
		return true
	}
	return false
}

func isVideoExt(url string) bool {
	ext := strings.ToLower(filepath.Ext(stripQuery(url)))
	switch ext {
	case ".mp4", ".mov", ".webm", ".mkv", ".avi":
		return true
	}
	return false
}

func inferContentType(url string) string {
	ext := filepath.Ext(stripQuery(url))
	if ct := mime.TypeByExtension(ext); ct != "" {
		return ct
	}
	return "application/octet-stream"
}

func filenameFromURL(rawURL string) string {
	u := stripQuery(rawURL)
	name := filepath.Base(u)
	if name == "" || name == "." || name == "/" {
		return "file"
	}
	return name
}

func stripQuery(rawURL string) string {
	if i := strings.IndexByte(rawURL, '?'); i >= 0 {
		return rawURL[:i]
	}
	return rawURL
}

````

[⬆ 回到目录](#toc)

<a name="file-messaging-media_test.go"></a>
## 📄 messaging/media_test.go

````go
package messaging

import "testing"

func TestExtractImageURLs(t *testing.T) {
	text := "check ![img](https://example.com/a.png) and ![](https://example.com/b.jpg)"
	urls := ExtractImageURLs(text)
	if len(urls) != 2 {
		t.Fatalf("expected 2 urls, got %d", len(urls))
	}
	if urls[0] != "https://example.com/a.png" {
		t.Errorf("urls[0] = %q", urls[0])
	}
	if urls[1] != "https://example.com/b.jpg" {
		t.Errorf("urls[1] = %q", urls[1])
	}
}

func TestExtractImageURLs_NoImages(t *testing.T) {
	urls := ExtractImageURLs("just plain text")
	if len(urls) != 0 {
		t.Errorf("expected 0 urls, got %d", len(urls))
	}
}

func TestExtractImageURLs_RelativeURL(t *testing.T) {
	text := "![img](./local.png)"
	urls := ExtractImageURLs(text)
	if len(urls) != 0 {
		t.Errorf("expected 0 urls for relative path, got %d", len(urls))
	}
}

func TestFilenameFromURL(t *testing.T) {
	tests := []struct {
		url  string
		want string
	}{
		{"https://example.com/photo.png", "photo.png"},
		{"https://example.com/path/to/report.pdf", "report.pdf"},
		{"https://example.com/file", "file"},
	}
	for _, tt := range tests {
		got := filenameFromURL(tt.url)
		if got != tt.want {
			t.Errorf("filenameFromURL(%q) = %q, want %q", tt.url, got, tt.want)
		}
	}
}

func TestFilenameFromURL_WithQuery(t *testing.T) {
	got := filenameFromURL("https://example.com/photo.png?token=abc")
	if got != "photo.png" {
		t.Errorf("got %q, want %q", got, "photo.png")
	}
}

func TestStripQuery(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"https://example.com/a?b=c", "https://example.com/a"},
		{"https://example.com/a", "https://example.com/a"},
		{"https://example.com/?x=1&y=2", "https://example.com/"},
	}
	for _, tt := range tests {
		got := stripQuery(tt.input)
		if got != tt.want {
			t.Errorf("stripQuery(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

````

[⬆ 回到目录](#toc)

<a name="file-messaging-sender.go"></a>
## 📄 messaging/sender.go

````go
package messaging

import (
	"context"
	"fmt"
	"log"

	"github.com/fastclaw-ai/weclaw/ilink"
	"github.com/google/uuid"
)

// NewClientID generates a new unique client ID for message correlation.
func NewClientID() string {
	return uuid.New().String()
}

// SendTypingState sends a typing indicator to a user via the iLink sendtyping API.
// It first fetches a typing_ticket via getconfig, then sends the typing status.
func SendTypingState(ctx context.Context, client *ilink.Client, userID, contextToken string) error {
	// Get typing ticket
	configResp, err := client.GetConfig(ctx, userID, contextToken)
	if err != nil {
		return fmt.Errorf("get config for typing: %w", err)
	}
	if configResp.TypingTicket == "" {
		return fmt.Errorf("no typing_ticket returned from getconfig")
	}

	// Send typing
	if err := client.SendTyping(ctx, userID, configResp.TypingTicket, ilink.TypingStatusTyping); err != nil {
		return fmt.Errorf("send typing: %w", err)
	}

	log.Printf("[sender] sent typing indicator to %s", userID)
	return nil
}

// SendTextReply sends a text reply to a user through the iLink API.
// If clientID is empty, a new one is generated.
func SendTextReply(ctx context.Context, client *ilink.Client, toUserID, text, contextToken, clientID string) error {
	if clientID == "" {
		clientID = NewClientID()
	}

	// Convert markdown to plain text for WeChat display
	plainText := MarkdownToPlainText(text)

	req := &ilink.SendMessageRequest{
		Msg: ilink.SendMsg{
			FromUserID:   client.BotID(),
			ToUserID:     toUserID,
			ClientID:     clientID,
			MessageType:  ilink.MessageTypeBot,
			MessageState: ilink.MessageStateFinish,
			ItemList: []ilink.MessageItem{
				{
					Type: ilink.ItemTypeText,
					TextItem: &ilink.TextItem{
						Text: plainText,
					},
				},
			},
			ContextToken: contextToken,
		},
		BaseInfo: ilink.BaseInfo{},
	}

	resp, err := client.SendMessage(ctx, req)
	if err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	if resp.Ret != 0 {
		return fmt.Errorf("send message failed: ret=%d errmsg=%s", resp.Ret, resp.ErrMsg)
	}

	log.Printf("[sender] sent reply to %s: %q", toUserID, truncate(text, 50))
	return nil
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

````

[⬆ 回到目录](#toc)

<a name="file-service-com.fastclaw.weclaw.plist"></a>
## 📄 service/com.fastclaw.weclaw.plist

````text
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.fastclaw.weclaw</string>
    <key>ProgramArguments</key>
    <array>
        <string>/usr/local/bin/weclaw</string>
        <string>start</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
    <key>KeepAlive</key>
    <true/>
    <key>StandardOutPath</key>
    <string>/tmp/weclaw.log</string>
    <key>StandardErrorPath</key>
    <string>/tmp/weclaw.log</string>
</dict>
</plist>

````

[⬆ 回到目录](#toc)

<a name="file-service-weclaw.service"></a>
## 📄 service/weclaw.service

````text
[Unit]
Description=WeClaw - WeChat AI Agent Bridge
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/bin/weclaw start
Restart=always
RestartSec=5
StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target

````

[⬆ 回到目录](#toc)

---
### 📊 最终统计汇总
- **文件总数:** 48
- **代码总行数:** 8409
- **物理总大小:** 225.10 KB

```

[⬆ 回到目录](#toc)

## docs/项目学习.md

```markdown
# WeClaw 项目学习笔记

> 对话时间：2026-03-30

---

## 1. 项目概述

**WeClaw** 是一个用 **Go 语言** 开发的微信 AI Agent 桥接器，将微信消息连接到各种 AI 代理（Claude、Codex、Gemini、Kimi 等）。

### 项目定位
- **核心功能**: 作为微信与 AI Agent 之间的桥梁
- **技术栈**: 纯 Go 语言实现，使用 Cobra CLI 框架
- **许可协议**: MIT 开源许可

### 项目灵感
```go
// 项目灵感来自腾讯官方 @tencent-weixin/openclaw-weixin
// 但 WeClaw 是独立实现的 Go 版本
```

---

## 2. 项目结构

```
weclaw/
├── main.go              # 程序入口点
├── cmd/                 # CLI 命令实现
│   ├── root.go          # 根命令 (Cobra)
│   ├── start.go         # 启动服务
│   ├── login.go         # 微信登录
│   ├── send.go          # 主动发送消息
│   ├── stop.go          # 停止服务
│   ├── status.go        # 查看状态
│   ├── update.go        # 更新版本
│   └── proc_*.go        # 进程管理 (跨平台)
├── agent/               # AI Agent 适配层
│   ├── agent.go         # Agent 接口定义
│   ├── acp_agent.go     # ACP 协议 Agent (1267行)
│   ├── cli_agent.go     # CLI 模式 Agent
│   └── http_agent.go    # HTTP API Agent
├── ilink/               # 微信 iLink 协议实现
│   ├── client.go        # iLink API 客户端
│   ├── auth.go          # 二维码登录认证
│   ├── monitor.go       # 消息长轮询监听
│   └── types.go         # 协议数据类型定义
├── messaging/           # 消息处理
│   ├── handler.go       # 消息路由与处理
│   ├── sender.go        # 消息发送
│   ├── media.go         # 媒体文件处理
│   ├── cdn.go           # 微信 CDN 上传/下载
│   └── markdown.go      # Markdown 转纯文本
├── api/                 # HTTP API 服务
│   └── server.go        # 主动消息推送 API
├── config/              # 配置管理
│   ├── config.go        # 配置加载/保存
│   └── detect.go        # Agent 自动检测
└── service/             # 系统服务配置
```

---

## 3. 多模式 Agent 支持

### 统一接口定义 (agent/agent.go)

```go
type Agent interface {
    Chat(ctx context.Context, conversationID string, message string) (string, error)
    ChatWithMedia(ctx context.Context, conversationID string, message string, media []MediaEntry) (string, error)
    ResetSession(ctx context.Context, conversationID string) (string, error)
    Info() AgentInfo
    SetCwd(cwd string)
}
```

### 三种模式对比

| 模式 | 说明 | 优势 | 实现文件 |
|------|------|------|----------|
| **ACP** | 长驻子进程，JSON-RPC 2.0 通信 | 速度最快，会话复用 | acp_agent.go (1267行) |
| **CLI** | 每条消息启动新进程 | 兼容性好，支持 `--resume` | cli_agent.go |
| **HTTP** | OpenAI 兼容 API | 易于集成，零代码接入 | http_agent.go |

### 支持的 Agent
`claude`、`codex`、`cursor`、`kimi`、`gemini`、`openclaw`、`opencode`、`pi`、`copilot`、`droid`、`iflow`、`kiro`、`qwen` 等

---

## 4. HTTP Agent 接入方式

### 配置示例 (~/.weclaw/config.json)

```json
{
  "default_agent": "gpt",
  "agents": {
    "gpt": {
      "type": "http",
      "endpoint": "https://api.openai.com/v1/chat/completions",
      "api_key": "sk-xxx",
      "model": "gpt-4o-mini",
      "system_prompt": "你是一个有用的助手",
      "aliases": ["4o", "chatgpt"]
    },
    "deepseek": {
      "type": "http",
      "endpoint": "https://api.deepseek.com/v1/chat/completions",
      "api_key": "sk-xxx",
      "model": "deepseek-chat",
      "aliases": ["ds"]
    },
    "本地模型": {
      "type": "http",
      "endpoint": "http://localhost:11434/v1/chat/completions",
      "model": "llama3",
      "aliases": ["llama"]
    }
  }
}
```

### 可接入的 API

| 服务商 | Endpoint |
|--------|----------|
| OpenAI | `https://api.openai.com/v1/chat/completions` |
| Azure OpenAI | `https://YOUR_RESOURCE.openai.azure.com/...` |
| DeepSeek | `https://api.deepseek.com/v1/chat/completions` |
| Moonshot | `https://api.moonshot.cn/v1/chat/completions` |
| 智谱 AI | `https://open.bigmodel.cn/api/paas/v4/chat/completions` |
| Ollama 本地 | `http://localhost:11434/v1/chat/completions` |
| LM Studio | `http://localhost:1234/v1/chat/completions` |
| vLLM | `http://localhost:8000/v1/chat/completions` |

### HTTP Agent 历史管理原理

**关键：客户端维护历史**

```go
type HTTPAgent struct {
    history    map[string][]ChatMessage  // conversationID -> messages
    maxHistory int                        // 默认 20 轮
}
```

**工作流程**：
1. 构建请求时带上历史 (`buildMessages`)
2. 收到回复后保存用户消息 + AI 回复到历史
3. 超过 `maxHistory*2` 时裁剪历史

```go
func (a *HTTPAgent) buildMessages(conversationID string, message string) []ChatMessage {
    var messages []ChatMessage
    // 1. 先加 system prompt
    if a.systemPrompt != "" {
        messages = append(messages, ChatMessage{Role: "system", Content: a.systemPrompt})
    }
    // 2. 加历史对话
    if hist, ok := a.history[conversationID]; ok {
        messages = append(messages, hist...)
    }
    // 3. 加当前消息
    messages = append(messages, ChatMessage{Role: "user", Content: message})
    return messages
}
```

**特点**：
- 多会话隔离 (`map[conversationID][]ChatMessage`)
- 重启后历史清空（内存存储）
- 每次请求带上完整历史（消耗更多 token）

---

## 5. ACP Agent 实现原理

### 架构图

```
┌─────────────────┐                    ┌─────────────────┐
│    WeClaw       │  ──── stdin ────▶  │  claude-agent   │
│   (父进程)      │                    │   (子进程)      │
│                 │  ◀──── stdout ──── │                 │
└─────────────────┘                    └─────────────────┘
        │                                      │
        │         JSON-RPC 2.0 over NDJSON     │
        └──────────────────────────────────────┘
```

### 核心架构

#### 1. 长驻子进程 + 懒加载

```go
func (a *ACPAgent) Start(ctx context.Context) error {
    a.cmd = exec.CommandContext(ctx, a.command, a.args...)
    a.cmd.Dir = a.cwd

    // 创建 stdin/stdout 管道
    a.stdin, _ = a.cmd.StdinPipe()
    stdout, _ := a.cmd.StdoutPipe()

    // 启动子进程
    a.cmd.Start()

    // 启动读取协程
    go a.readLoop()

    // 初始化握手
    a.call(ctx, "initialize", initParams{...})
}
```

- 子进程**只启动一次**，后续请求复用
- 懒加载：首次 `Chat()` 时才启动

#### 2. 双协议支持

| 协议 | 适用 Agent | 会话模型 |
|------|-----------|---------|
| `legacy_acp` | claude-agent-acp, cursor agent | Session 模型 |
| `codex_app_server` | codex app-server | Thread/Turn 模型 |

#### 3. 请求-响应关联 (pending map)

```go
type ACPAgent struct {
    pending   map[int64]chan *rpcResponse  // 请求ID -> 响应channel
    nextID    atomic.Int64                  // 自增ID生成器
}

// 发送请求
id := a.nextID.Add(1)
a.pending[id] = responseCh
a.stdin.Write(request)

// readLoop 收到响应
if msg.ID != nil {
    a.pending[*msg.ID] <- response  // 唤醒等待的调用
}
```

#### 4. 流式响应处理

Agent 的回复是**分块推送**的，通过 `session/update` 通知：

```go
// 注册通知 channel
notifyCh := make(chan *sessionUpdate, 256)
a.notifyCh[sessionID] = notifyCh

// 异步发送 prompt
go func() {
    a.call(ctx, "session/prompt", params)
    promptDone <- struct{}{}
}()

// 收集流式文本块
var textParts []string
for {
    select {
    case update := <-notifyCh:
        if update.SessionUpdate == "agent_message_chunk" {
            textParts = append(textParts, extractChunkText(update))
        }
    case <-promptDone:
        // 排空剩余通知后返回
        return strings.Join(textParts, "")
    }
}
```

**消息流**：
```
WeClaw                          Agent
  │                               │
  │──── session/prompt ──────────▶│
  │                               │
  │◀─── session/update (chunk) ───│  "你"
  │◀─── session/update (chunk) ───│  "好"
  │◀─── session/update (chunk) ───│  "！"
  │      ...                      │
  │◀─── prompt response ──────────│  {stopReason: "end"}
  │                               │
  └── 返回完整文本 ────────────────┘
```

#### 5. 会话管理与隔离

```go
type ACPAgent struct {
    sessions map[string]string  // conversationID -> sessionID (legacy ACP)
    threads  map[string]string  // conversationID -> threadID (codex)
}
```

- 每个微信对话独立 session/thread
- 自动创建，按需复用

#### 6. 自动权限处理

```go
func (a *ACPAgent) handlePermissionRequest(raw string) {
    // 找到 "allow" 选项
    optionID := "allow"
    for _, opt := range req.Params.Options {
        if opt.Kind == "allow" {
            optionID = opt.OptionID
            break
        }
    }

    // 自动发送允许响应
    resp := map[string]interface{}{
        "jsonrpc": "2.0",
        "id":      req.ID,
        "result":  map[string]interface{}{
            "outcome": map[string]interface{}{
                "outcome":  "selected",
                "optionId": optionID,
            },
        },
    }
    a.stdin.Write(resp)
}
```

---

## 6. JSON-RPC 协议详解

### 核心概念

**JSON-RPC** 是轻量级远程过程调用协议，使用 JSON 作为数据格式。

### 消息格式

#### 请求 (Request)
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "session/prompt",
  "params": {"sessionId": "xxx", "prompt": [...]}
}
```

#### 响应 (Response)
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {"stopReason": "end"}
}
```

#### 错误响应
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "error": {
    "code": -32600,
    "message": "Invalid request"
  }
}
```

#### 通知 (Notification)
```json
{
  "jsonrpc": "2.0",
  "method": "session/update",
  "params": {...}
}
```
**没有 id，不需要响应**，用于单向推送（如流式文本块）。

### JSON-RPC vs REST

| 特性 | JSON-RPC | REST |
|------|----------|------|
| **URL** | 单一端点 | 多个资源路径 |
| **HTTP 方法** | 通常 POST | GET/POST/PUT/DELETE |
| **语义** | `method: "createUser"` | `POST /users` |
| **批量请求** | ✅ 原生支持 | ❌ 需自定义 |
| **通知** | ✅ 支持 | ❌ 需 WebSocket |
| **传输层** | 任意 | 通常 HTTP |

### 为什么 ACP 选择 JSON-RPC？

1. **简单** - 只有请求、响应、通知三种消息
2. **灵活** - 不依赖 HTTP，可以用 stdio、socket 等
3. **双向** - 服务端可以主动推送通知
4. **标准化** - 规范明确，易于实现

---

## 7. 微信端透明性

### 完全解耦架构

```
┌─────────────┐         ┌─────────────┐         ┌─────────────┐
│   微信用户   │  ◀────▶ │   WeClaw    │  ◀────▶ │  AI Agent   │
│             │         │   (中间层)   │         │ (任意后端)   │
└─────────────┘         └─────────────┘         └─────────────┘
      │                        │                       │
      │    只看到"机器人"回复    │                       │
      │    不知道背后是什么 AI   │                       │
      └────────────────────────┘                       │
```

### 微信协议层只包含纯文本

```go
SendMessage {
    To: "user_xxx@im.wechat",
    Items: [{
        Type: 1,  // 文本
        Text: "你好！有什么可以帮助你的？"  // 只有纯文本
    }]
}
```

### 设计优势

| 优势 | 说明 |
|------|------|
| **后端无关** | 微信用户无感知，随时切换后端 |
| **多 Agent** | 命令路由 (`/gpt`, `/claude`) |
| **安全性** | 不暴露技术架构 |
| **灵活性** | 可随时添加/移除 Agent |

### 与其他平台对比

| 平台 | 是否显示后端 |
|------|-------------|
| ChatGPT | ✅ 显示 "GPT-4" |
| Claude | ✅ 显示 "Claude 3.5" |
| Poe | ✅ 显示模型名称 |
| **WeClaw** | ❌ **完全隐藏** |

---

## 8. 微信 iLink 协议实现

### 核心文件

| 文件 | 功能 |
|------|------|
| ilink/types.go | 协议数据类型定义 |
| ilink/client.go | API 客户端实现 |
| ilink/auth.go | 登录认证流程 |
| ilink/monitor.go | 消息监听 |
| messaging/handler.go | 消息处理逻辑 |

### API 端点

```go
const defaultBaseURL = "https://ilinkai.weixin.qq.com"

/ilink/bot/get_bot_qrcode    // 获取登录二维码
/ilink/bot/get_qrcode_status // 查询扫码状态
/ilink/bot/getupdates        // 长轮询获取消息 (35秒超时)
/ilink/bot/sendmessage       // 发送消息
/ilink/bot/sendtyping        // 发送输入状态
/ilink/bot/getconfig         // 获取配置（含 typing_ticket）
/ilink/bot/getuploadurl      // 获取 CDN 上传地址
```

### 消息类型处理

```go
// 消息类型
MessageTypeUser = 1   // 用户消息
MessageTypeBot  = 2   // 机器人消息

// 消息状态
MessageStateFinish = 2  // 已完成

// 内容类型
ItemTypeText  = 1   // 文本
ItemTypeImage = 2   // 图片
ItemTypeVoice = 3   // 语音
ItemTypeFile  = 4   // 文件
ItemTypeVideo = 5   // 视频
```

### 各类型处理流程

#### 文本 (ItemTypeText = 1)
- 直接提取 `TextItem.Text`
- 解析命令 (`/gpt`, `@claude` 等)
- 路由到对应 Agent

#### 语音 (ItemTypeVoice = 3)
- **微信服务端已做 ASR 转文字**
- 直接使用 `VoiceItem.Text`
- 无需本地语音识别

#### 图片/文件/视频 (ItemType 2/4/5)
- 优先使用 HTTP URL
- 否则从 CDN 下载 + AES-128-ECB 解密
- 保存到本地后传给 Agent

### CDN 加密通信

```go
// 加密方案
- AES-128-ECB 模式
- PKCS7 填充
- 随机 16 字节 AES 密钥

// 解密流程
cdnURL := "https://novac2c.cdn.weixin.qq.com/c2c/download?encrypted_query_param=xxx"
encryptedData := download(cdnURL)
aesKey := base64Decode(media.AESKey) -> hexDecode()
decrypted := aes128ECBDecrypt(encryptedData, aesKey)

// 解密代码
func decryptAES128ECB(encrypted, key []byte) ([]byte, error) {
    block, _ := aes.NewCipher(key)
    decrypted := make([]byte, len(encrypted))
    for i := 0; i < len(encrypted); i += aes.BlockSize {
        block.Decrypt(decrypted[i:i+aes.BlockSize], encrypted[i:i+aes.BlockSize])
    }
    // PKCS7 去填充
    padding := int(decrypted[len(decrypted)-1])
    return decrypted[:len(decrypted)-padding], nil
}
```

### 与官方协议对齐情况

| 方面 | 状态 | 说明 |
|------|------|------|
| API 端点 | ✅ 完全对齐 | 使用腾讯官方域名 |
| 认证流程 | ✅ 完全对齐 | QRCode → 扫码 → BotToken |
| 消息结构 | ✅ 完全对齐 | `WeixinMessage` 结构完整 |
| 消息类型 | ✅ 5 种全支持 | Text/Image/Voice/File/Video |
| CDN 加密 | ✅ 完全对齐 | AES-128-ECB + PKCS7 |
| 输入状态 | ✅ 支持 | `sendtyping` API |
| 会话管理 | ✅ 支持 | `context_token` 传递 |

### 对比官方 SDK

| 特性 | 官方 OpenClaw | WeClaw |
|------|--------------|--------|
| 语言 | TypeScript | Go |
| Agent 支持 | 仅 Claude | 多 Agent (ACP/CLI/HTTP) |
| 部署 | 需要 Node.js | 单二进制 |
| 消息类型 | 基础 | 完整 (含媒体) |
| CDN 加密 | ❓ | ✅ 完整实现 |

---

## 9. ACP vs HTTP vs CLI 全面对比

| 特性 | ACP | HTTP | CLI |
|------|-----|------|-----|
| **启动方式** | 长驻子进程 | HTTP 请求 | 每次新进程 |
| **通信协议** | stdio + JSON-RPC | REST API | 命令行参数 |
| **会话管理** | Agent 内部 | WeClaw 本地 | --resume 参数 |
| **流式响应** | ✅ 实时推送 | ❌ 批量返回 | ✅ 实时输出 |
| **性能** | ⚡ 最快 | 🚀 快 | 🐢 慢（启动开销） |
| **适用场景** | 本地 Agent | 云端 API | 简单集成 |
| **历史管理** | Agent 内部维护 | WeClaw 内存维护 | 外部文件 |

---

## 10. 命令系统

### 内置命令

```
/info           → 显示当前 Agent 状态
/help           → 显示帮助信息
/new 或 /clear  → 重置会话
/cwd /path      → 切换工作目录
```

### Agent 路由

```
"hello"              → 发送给默认 Agent
"/gpt 你好"          → 发送给 gpt Agent
"@claude @codex 你好" → 广播给多个 Agent
"/claude"            → 切换默认 Agent 为 claude
```

### 内置别名

```go
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
```

---

## 11. 依赖分析

```go
require (
    github.com/google/uuid v1.6.0      // UUID 生成
    github.com/mdp/qrterminal/v3 v3.2.1 // 终端二维码显示
    github.com/spf13/cobra v1.10.2     // CLI 框架
    golang.org/x/net v0.52.0           // HTTP 客户端
    rsc.io/qr v0.2.0                   // QR 码生成
)
```

项目依赖精简，全部使用标准库和少量必要第三方库。

---

## 12. 设计亮点总结

| 亮点 | 说明 |
|------|------|
| **统一接口抽象** | Agent 接口支持插件式扩展 |
| **长驻进程** | ACP 模式避免重复启动，响应最快 |
| **异步 readLoop** | 单协程处理所有响应，避免并发复杂性 |
| **pending map** | 优雅的请求-响应关联，支持并发请求 |
| **notifyCh 分发** | 按会话 ID 路由通知，支持多会话并行 |
| **双协议兼容** | 同时支持 legacy ACP 和 codex app-server |
| **自动权限** | 无感处理工具调用权限，用户体验好 |
| **流式聚合** | 实时收集文本块，最终返回完整响应 |
| **后端无关** | 微信用户无感知，随时切换后端 |
| **零代码接入** | HTTP Agent 纯配置即可接入任意 OpenAI 兼容 API |
| **协议完整** | 完整实现微信 iLink 协议和多种 AI Agent 协议 |
| **安全可靠** | 完整的 CDN 加密通信实现 |
| **运维成熟** | 完善的部署和更新机制 |

---

## 13. 学习价值

WeClaw 是一个优秀的学习案例，涵盖：

1. **Go 并发编程** - goroutine、channel、sync.Map、atomic
2. **进程间通信** - stdio 管道、JSON-RPC 协议
3. **协议实现** - 微信 iLink、OpenAI API 兼容
4. **加密算法** - AES-128-ECB、PKCS7 填充
5. **架构设计** - 接口抽象、插件模式、中间层设计
6. **CLI 开发** - Cobra 框架、系统服务集成

该项目适合作为学习微信机器人开发和 AI Agent 集成的优秀参考案例。

---

## 14. Linkhoard 网页剪藏功能

### 功能概述

当用户在微信中发送 URL 时，WeClaw 会自动拦截并将网页内容保存为本地 Markdown 文件。

### 核心文件

| 文件 | 功能 |
|------|------|
| messaging/linkhoard.go | 网页抓取与 Markdown 生成 |
| messaging/markdown.go | Markdown 转纯文本 |

### 双抓取策略

```go
func SaveLinkToLinkhoard(ctx context.Context, saveDir, rawURL string) (*LinkMetadata, error) {
    var meta *LinkMetadata
    var err error

    if isWeChatURL(rawURL) {
        // 微信文章：直接抓取（带浏览器 Header）
        meta, err = FetchLinkMetadata(ctx, rawURL)
    } else {
        // 外部链接：使用 Jina Reader API
        meta, err = FetchViaJina(ctx, rawURL)
        if err != nil {
            // 降级到直接抓取
            meta, err = FetchLinkMetadata(ctx, rawURL)
        }
    }
    // ...
}
```

#### 微信文章抓取

```go
// 伪造浏览器 Header，绕过反爬
req.Header.Set("User-Agent", "Mozilla/5.0 ...")
req.Header.Set("Referer", "https://mp.weixin.qq.com/")

// 解析 HTML，提取元数据
extractMeta(doc, meta)  // og:title, og:description, author, etc.
```

#### 外部链接 - Jina Reader

```go
// Jina Reader API: https://r.jina.ai/{url}
jinaURL := "https://r.jina.ai/" + rawURL
// 返回格式：
// Title: xxx
// URL Source: xxx
// Published Time: xxx
// Markdown Content:
// [正文 Markdown]
```

### 保存格式

生成的 Markdown 文件包含 Frontmatter：

```markdown
---
title: '文章标题'
source: 'https://original.url'
published: '2024-01-01T00:00:00Z'
created: '2024-01-01T12:00:00Z'
description: '文章描述'
openGraphImage: 'https://cover.image.jpg'
author:
  - '[[作者名]]'
---

[正文内容]
```

### 侧边文件 (Sidecar)

每个保存的文件都附带 `.sidecar.md` 文件：

```markdown
---
id: uuid-v4-generated-id
---
```

用于与 Linkhoard / Obsidian 等工具集成。

---

## 15. 微信文章自动分析

### 功能概述

在保存微信文章后，自动触发 `nanobot` Agent 进行内容分析，实现「保存 → 分析 → 推送」的闭环。

### 实现流程

```
用户发送微信文章链接
        ↓
Linkhoard 保存为 Markdown
        ↓
判断：是否为微信文章？
        ↓ 是
异步调用 nanobot 分析
        ↓
发送分析结果给用户
```

### 核心代码 (handler.go:696-728)

```go
// URL 拦截处理
if h.saveDir != "" && IsURL(trimmed) {
    meta, err := SaveLinkToLinkhoard(ctx, h.saveDir, rawURL)
    if err == nil {
        reply = fmt.Sprintf("已保存: %s", meta.Title)
        // 如果是微信文章，触发自动分析
        if isWeChatURL(rawURL) {
            go h.analyzeWithNanobot(ctx, client, msg, meta)
        }
    }
}

// 自动分析函数
func (h *Handler) analyzeWithNanobot(ctx context.Context, client *ilink.Client,
    msg ilink.WeixinMessage, meta *LinkMetadata) {
    // 获取 nanobot agent
    ag, err := h.getAgent(ctx, "nanobot")

    // 构建分析提示词（发送全文）
    prompt := fmt.Sprintf("请分析这篇微信文章，给出摘要和关键观点：\n\n标题：%s\n\n文章内容：\n%s",
        meta.Title, meta.Body)

    // 获取分析结果
    reply, err := h.chatWithAgent(ctx, ag, msg.FromUserID, prompt)

    // 发送分析结果
    SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID)
}
```

### 数据流向

```
微信文章 URL
    ↓
SaveLinkToLinkhoard()
    ├─→ FetchLinkMetadata()  → meta.Title + meta.Body
    └─→ 保存 .md 文件
    ↓
analyzeWithNanobot()
    ├─→ getAgent("nanobot")
    ├─→ chatWithAgent(prompt={标题+正文})
    └─→ SendTextReply(分析结果)
```

### 设计亮点

| 特性 | 说明 |
|------|------|
| **异步处理** | `go h.analyzeWithNanobot()` 不阻塞保存确认 |
| **全文发送** | 发送 `meta.Body` 而非 URL，nanobot 无需访问网络 |
| **失败隔离** | 保存失败时不触发分析，分析失败不影响保存 |
| **仅限微信** | 只对微信文章触发，避免分析无关链接 |

---

## 16. Attachment 自动抓取与推送

### 功能概述

当 AI Agent 的回复包含本地文件路径时，WeClaw 会自动：
1. 提取文件路径（正则匹配）
2. 校验路径安全性
3. 上传到微信 CDN
4. 发送给用户

### 核心代码 (handler.go:501-534)

```go
func (h *Handler) sendReplyWithMedia(ctx context.Context, client *ilink.Client,
    msg ilink.WeixinMessage, agentName, reply, clientID string) {
    // 1. 提取本地文件路径
    attachmentPaths := extractLocalAttachmentPaths(reply)
    // 正则: `/path/to/file.(pdf|png|jpg|xlsx|...)`

    // 2. 获取允许的根目录
    allowedRoots := h.allowedAttachmentRoots(agentName)
    // 默认: 当前工作目录 + agent 的 cwd

    // 3. 校验并上传
    for _, attachmentPath := range attachmentPaths {
        if !isAllowedAttachmentPath(attachmentPath, allowedRoots) {
            log.Printf("[handler] rejected attachment outside allowed roots")
            continue  // 跳过非安全路径
        }
        SendMediaFromPath(ctx, client, msg.FromUserID, attachmentPath, ...)
    }
}
```

### 安全性设计

```go
func isAllowedAttachmentPath(path string, roots []string) bool {
    for _, root := range roots {
        if strings.HasPrefix(filepath.Clean(path), root) {
            return true  // 只允许白名单目录
        }
    }
    return false
}
```

### 支持的文件类型

| 类别 | 扩展名 |
|------|--------|
| 图片 | png, jpg, jpeg, gif, webp |
| 文档 | pdf, txt, md, csv |
| 表格 | xlsx, xls |
| 代码 | py, js, ts, go, java, etc. |

### 典型应用场景

```
用户: "帮我分析 data.csv 并生成图表"
    ↓
Agent: (运行 Python 生成 /tmp/output.png)
Agent: "图表已生成：/tmp/output.png"
    ↓
WeClaw: 检测到文件路径 → 上传到微信 → 发送图片给用户
```

---

## 17. 消息去重机制

### 问题背景

微信服务器在网络不稳定时会重试推送相同的消息：
- 语音消息转文字可能触发多次状态变更
- 同一条消息可能收到多次 `MessageStateFinish`

### 解决方案 (handler.go:276-284)

```go
type Handler struct {
    seenMsgs sync.Map  // map[int64]time.Time — dedup by message_id
}

func (h *Handler) HandleMessage(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage) {
    // 去重检查
    if msg.MessageID != 0 {
        if _, loaded := h.seenMsgs.LoadOrStore(msg.MessageID, time.Now()); loaded {
            return  // 已处理过，直接丢弃
        }
        // 异步清理旧条目（5分钟过期）
        go h.cleanSeenMsgs()
    }
    // ...
}

func (h *Handler) cleanSeenMsgs() {
    cutoff := time.Now().Add(-5 * time.Minute)
    h.seenMsgs.Range(func(key, value any) bool {
        if t, ok := value.(time.Time); ok && t.Before(cutoff) {
            h.seenMsgs.Delete(key)
        }
        return true
    })
}
```

### 设计特点

| 特性 | 说明 |
|------|------|
| **并发安全** | `sync.Map` 无需加锁 |
| **自动清理** | 5分钟后自动删除旧记录 |
| **异步清理** | 不阻塞消息处理 |
| **MessageID** | 使用微信服务器的唯一 ID |

---

## 18. Markdown 转纯文本

### 功能概述

AI Agent 返回的 Markdown 需要转换为微信可显示的纯文本。

### 转换规则 (markdown.go)

| Markdown | 转换后 | 说明 |
|----------|--------|------|
| `**bold**` | `bold` | 移除加粗标记 |
| `[text](url)` | `text` | 保留显示文本，移除链接 |
| `![alt](img)` | (删除) | 移除图片 |
| `` `code` `` | `code` | 保留代码内容 |
| `# Header` | `Header` | 移除 # 前缀 |
| `> quote` | `quote` | 移除引用标记 |
| `- item` | `• item` | 转换为圆点 |
| `\n\n\n` | `\n\n` | 折叠多余空行 |

### 核心正则

```go
var (
    reCodeBlock   = regexp.MustCompile("(?s)```[^\n]*\n?(.*?)```")
    reInlineCode  = regexp.MustCompile("`([^`]+)`")
    reImage       = regexp.MustCompile(`!\[[^\]]*\]\([^)]*\)`)
    reLink        = regexp.MustCompile(`\[([^\]]+)\]\([^)]*\)`)
    reTableSep    = regexp.MustCompile(`(?m)^\|[\s:|\-]+\|$`)
    reTableRow    = regexp.MustCompile(`(?m)^\|(.+)\|$`)
    reHeader      = regexp.MustCompile(`(?m)^#{1,6}\s+`)
)
```

### 转换示例

```
输入:
# 标题
**加粗** 和 [链接](https://example.com)

```go
code block
```

输出:
标题
加粗 和 链接

code block
```

---

## 19. 媒体文件处理

### CDN 加密下载

```go
func downloadCDNMedia(ctx context.Context, client *ilink.Client,
    media *ilink.MediaInfo, saveDir string, ext string) (string, error) {

    // 1. 构建 CDN URL
    cdnURL := fmt.Sprintf("https://novac2c.cdn.weixin.qq.com/c2c/download?encrypted_query_param=%s",
        url.QueryEscape(media.EncryptQueryParam))

    // 2. 下载加密数据
    encryptedData := download(cdnURL)

    // 3. 解密 (AES-128-ECB + PKCS7)
    if media.AESKey != "" {
        aesKey := base64Decode(media.AESKey) -> hexDecode()
        fileData = decryptAES128ECB(encryptedData, aesKey)
    }

    // 4. 保存到本地
    filePath := filepath.Join(saveDir, uuid.New().String() + ext)
    os.WriteFile(filePath, fileData, 0644)
    return filePath, nil
}
```

### 支持的媒体类型

| 类型 | ItemType | 处理方式 |
|------|----------|----------|
| 图片 | 2 | 保存 .jpg，传给 Agent |
| 语音 | 3 | 使用服务端转文字结果 |
| 文件 | 4 | 保留扩展名，传路径给 Agent |
| 视频 | 5 | 保存 .mp4，传路径给 Agent |

---

## 20. 后台守护进程 (Daemon)

### 实现原理 (cmd/start.go)

WeClaw 通过**自己重启自己**的方式实现跨平台后台运行：

```go
if !foreground {
    // 构建后台命令
    cmd := exec.Command(os.Args[0], "start", "-f")
    cmd.Dir = cwd

    // Unix: 脱离终端
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Setsid: true,  // 创建新会话
    }

    // Windows: 不使用控制台
    if runtime.GOOS == "windows" {
        cmd.SysProcAttr = &syscall.SysProcAttr{
            HideWindow:    true,
            CreationFlags: 0x08000000, // CREATE_NO_WINDOW
        }
    }

    // 重定向输出到日志文件
    logFile, _ := os.OpenFile("weclaw.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
    cmd.Stdout = logFile
    cmd.Stderr = logFile

    // 启动并释放进程
    cmd.Start()
    cmd.Process.Release()  // 让子进程独立运行
}
```

### 跨平台支持

| 平台 | 实现方式 |
|------|----------|
| Linux/macOS | `SysProcAttr.Setsid = true` |
| Windows | `CREATE_NO_WINDOW` 标志 |
| 通用 | 输出重定向到 `weclaw.log` |

---

## 21. 总结与补充

### 额外的设计亮点

| 亮点 | 说明 |
|------|------|
| **自己重启自己** | Daemon 实现，无需第三方库 |
| **消息去重** | `sync.Map` + MessageID，防止重复处理 |
| **Attachment 自动推送** | 正则提取路径，安全校验，CDN 上传 |
| **Markdown 转文本** | 完整的正则转换，适配微信显示 |
| **双抓取策略** | 微信直接抓，外部用 Jina Reader |
| **自动分析** | 微信文章保存后触发 nanobot 分析 |
| **流式聚合** | 实时收集文本块，最终返回完整响应 |

### 项目依赖精简

```go
require (
    github.com/google/uuid v1.6.0      // UUID 生成
    github.com/mdp/qrterminal/v3 v3.2.1 // 终端二维码
    github.com/spf13/cobra v1.10.2     // CLI 框架
    golang.org/x/net v0.52.0           // HTTP 客户端
    rsc.io/qr v0.2.0                   // QR 码生成
)
```

**零外部依赖**：除了 CLI 框架和二维码库，全盘使用 Go 标准库。

---

## 22. 进度通知机制 - 解决"等待焦虑"

### 问题背景

当 Agent 执行复杂任务（如遍历项目目录、读写多个文件）时，可能耗时 30~60 秒。虽然底层 ACP 协议支持流式推送，但 WeClaw 的处理方式是**收集所有文本块直到完成才一次性发送**，导致微信端长期处于"对方正在输入"状态，用户体验很差。

### 解决方案：进度通知回调

实现了一套进度通知机制，让 Agent 在执行耗时操作时主动向微信推送进度提示。

### 架构设计

```
┌─────────────┐      progress events      ┌─────────────┐
│  ACPAgent   │ ─────────────────────────▶ │   Handler   │
│             │                            │             │
│ readLoop()  │  ProgressCallback(event)   │ SendText()  │
└─────────────┘                            └─────────────┘
       │                                          │
       │  检测到:                                  │
       │  - 工具调用开始 (permission request)      │
       │  - 非消息 item 开始 (item/started)        │
       └──────────────────────────────────────────┘
                                               │
                                               ▼
                                    ┌─────────────┐
                                    │   微信用户   │
                                    │ ⏳ 正在调用  │
                                    │   工具: xxx  │
                                    └─────────────┘
```

### 核心代码实现

#### 1. 进度事件定义 (agent/agent.go:30-52)

```go
// ProgressType represents the type of progress event.
type ProgressType string

const (
    ProgressTypeToolStart   ProgressType = "tool_start"   // Tool execution started
    ProgressTypeToolEnd     ProgressType = "tool_end"     // Tool execution ended
    ProgressTypeThought     ProgressType = "thought"      // Agent thinking
    ProgressTypeFileRead    ProgressType = "file_read"    // Reading file
    ProgressTypeFileWrite   ProgressType = "file_write"   // Writing file
    ProgressTypeProcessing  ProgressType = "processing"   // General processing
    ProgressTypeSearching   ProgressType = "searching"    // Searching
)

// ProgressEvent represents a progress notification from an agent.
type ProgressEvent struct {
    Type     ProgressType // Type of progress event
    Message  string       // Human-readable progress message
    ToolName string       // Name of the tool being used (optional)
}

// ProgressCallback is called when an agent reports progress.
type ProgressCallback func(ctx context.Context, event ProgressEvent)
```

#### 2. Agent 接口扩展 (agent/agent.go:109-111)

```go
type Agent interface {
    // ... existing methods ...

    // SetProgressCallback sets a callback for progress notifications.
    SetProgressCallback(callback ProgressCallback)
}
```

#### 3. ACPAgent 进度跟踪 (agent/acp_agent.go:334-351)

```go
// ACPAgent 结构体中添加
type ACPAgent struct {
    // ... existing fields ...
    progressCallback ProgressCallback // progress notification callback
}

// SetProgressCallback sets a callback for progress notifications.
func (a *ACPAgent) SetProgressCallback(callback ProgressCallback) {
    a.mu.Lock()
    defer a.mu.Unlock()
    a.progressCallback = callback
}

// sendProgress sends a progress event if a callback is registered.
func (a *ACPAgent) sendProgress(ctx context.Context, event ProgressEvent) {
    a.mu.Lock()
    callback := a.progressCallback
    a.mu.Unlock()

    if callback != nil {
        // Call callback in goroutine to avoid blocking
        go callback(ctx, event)
    }
}
```

#### 4. 工具调用检测 (agent/acp_agent.go:1132-1147)

```go
func (a *ACPAgent) handlePermissionRequest(raw string) {
    var req struct {
        ID     json.RawMessage         `json:"id"`
        Params permissionRequestParams `json:"params"`
    }
    json.Unmarshal([]byte(raw), &req)

    // Extract tool name for progress notification
    var toolName string
    if req.Params.ToolCall != nil {
        var toolCall struct {
            Name string `json:"name"`
        }
        json.Unmarshal(req.Params.ToolCall, &toolCall)
        if toolCall.Name != "" {
            toolName = toolCall.Name
            // Send progress notification
            a.sendProgress(context.Background(), ProgressEvent{
                Type:     ProgressTypeToolStart,
                Message:  fmt.Sprintf("正在调用工具: %s", toolName),
                ToolName: toolName,
            })
        }
    }
    // ... auto-allow permission ...
}
```

#### 5. Codex 进度检测 (agent/acp_agent.go:1073-1090)

```go
func (a *ACPAgent) handleCodexItemStarted(params json.RawMessage) {
    // ... parse params ...

    // Send progress notification for non-agentMessage items
    if p.Item.Type != "agentMessage" {
        var message string
        switch p.Item.Type {
        case "tool_use":
            message = "正在执行工具..."
        case "thinking":
            message = "正在思考..."
        default:
            message = fmt.Sprintf("正在处理: %s", p.Item.Type)
        }
        a.sendProgress(context.Background(), ProgressEvent{
            Type:    ProgressTypeProcessing,
            Message: message,
        })
        return
    }
    // ... handle agentMessage ...
}
```

#### 6. Handler 进度处理 (messaging/handler.go:54-62, 606-626)

```go
// progressContext holds context for sending progress notifications.
type progressContext struct {
    client   *ilink.Client
    userID   string
    token    string
    lastTime time.Time // last progress notification time
    mu       sync.Mutex
}

// handleProgressEvent handles a progress event from an agent.
func (h *Handler) handleProgressEvent(ctx context.Context, pCtx *progressContext, event agent.ProgressEvent) {
    // Check rate limit: at most 1 notification per 3 seconds
    pCtx.mu.Lock()
    now := time.Now()
    if !pCtx.lastTime.IsZero() && now.Sub(pCtx.lastTime) < 3*time.Second {
        pCtx.mu.Unlock()
        return
    }
    pCtx.lastTime = now
    pCtx.mu.Unlock()

    // Send progress notification to WeChat
    message := fmt.Sprintf("⏳ %s", event.Message)
    SendTextReply(ctx, pCtx.client, pCtx.userID, message, pCtx.token, NewClientID())
}
```

#### 7. 设置回调 (messaging/handler.go:556-584)

```go
func (h *Handler) chatWithAgent(ctx context.Context, ag agent.Agent, userID, message string,
    clientAndToken ...interface{}) (string, error) {

    // Set up progress callback if client and token are provided
    if len(clientAndToken) >= 2 {
        if client, ok := clientAndToken[0].(*ilink.Client); ok {
            if token, ok := clientAndToken[1].(string); ok {
                if contextTokenVal, ok := h.contextTokens.Load(userID); ok {
                    if contextToken, ok := contextTokenVal.(string); ok {
                        // Create progress context
                        pCtx := &progressContext{
                            client:   client,
                            userID:   userID,
                            token:    contextToken,
                            lastTime: time.Time{},
                        }

                        // Set progress callback on the agent
                        ag.SetProgressCallback(func(ctx context.Context, event agent.ProgressEvent) {
                            h.handleProgressEvent(ctx, pCtx, event)
                        })

                        // Clean up after chat completes
                        defer func() {
                            h.setProgressContext(nil)
                        }()
                    }
                }
            }
        }
    }
    // ... call agent.Chat ...
}
```

### 功能特点

| 特性 | 说明 |
|------|------|
| **异步回调** | `go callback(ctx, event)` 不阻塞 Agent 处理 |
| **限流保护** | 最多每 3 秒发送一次，避免刷屏 |
| **仅限 ACP** | 只有 ACP 协议的 agent 支持进度通知 |
| **友好提示** | 消息格式 `⏳ 正在调用工具: xxx` |
| **自动清理** | 请求完成后自动清除回调 |

### 使用效果

```
用户: 帮我分析一下项目的所有 Go 文件

[Agent 开始工作...]

微信接收:
⏳ 正在调用工具: list_directory
⏳ 正在调用工具: read_file
⏳ 正在调用工具: read_file

[最终回复]
我已经分析了项目的所有 Go 文件...
```

### 设计亮点

| 亮点 | 说明 |
|------|------|
| **接口扩展性** | `SetProgressCallback` 方法可选实现，CLI/HTTP 返回空操作 |
| **非阻塞设计** | 回调在独立 goroutine 中执行，不影响 Agent 性能 |
| **上下文隔离** | 每个请求独立的 `progressContext`，避免并发冲突 |
| **用户体验** | 从"黑盒等待"变为"可见进度"，大幅提升体验 |

```

[⬆ 回到目录](#toc)

## go.mod

```text
module github.com/fastclaw-ai/weclaw

go 1.25.0

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mdp/qrterminal/v3 v3.2.1 // indirect
	github.com/spf13/cobra v1.10.2 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	golang.org/x/net v0.52.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/term v0.41.0 // indirect
	rsc.io/qr v0.2.0 // indirect
)

```

[⬆ 回到目录](#toc)

## go.sum

```text
github.com/cpuguy83/go-md2man/v2 v2.0.6/go.mod h1:oOW0eioCTA6cOiMLiUPZOpcVxMig6NIQQ7OS05n1F4g=
github.com/google/uuid v1.6.0 h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=
github.com/google/uuid v1.6.0/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
github.com/inconshreveable/mousetrap v1.1.0 h1:wN+x4NVGpMsO7ErUn/mUI3vEoE6Jt13X2s0bqwp9tc8=
github.com/inconshreveable/mousetrap v1.1.0/go.mod h1:vpF70FUmC8bwa3OWnCshd2FqLfsEA9PFc4w1p2J65bw=
github.com/mdp/qrterminal/v3 v3.2.1 h1:6+yQjiiOsSuXT5n9/m60E54vdgFsw0zhADHhHLrFet4=
github.com/mdp/qrterminal/v3 v3.2.1/go.mod h1:jOTmXvnBsMy5xqLniO0R++Jmjs2sTm9dFSuQ5kpz/SU=
github.com/russross/blackfriday/v2 v2.1.0/go.mod h1:+Rmxgy9KzJVeS9/2gXHxylqXiyQDYRxCVz55jmeOWTM=
github.com/spf13/cobra v1.10.2 h1:DMTTonx5m65Ic0GOoRY2c16WCbHxOOw6xxezuLaBpcU=
github.com/spf13/cobra v1.10.2/go.mod h1:7C1pvHqHw5A4vrJfjNwvOdzYu0Gml16OCs2GRiTUUS4=
github.com/spf13/pflag v1.0.9 h1:9exaQaMOCwffKiiiYk6/BndUBv+iRViNW+4lEMi0PvY=
github.com/spf13/pflag v1.0.9/go.mod h1:McXfInJRrz4CZXVZOBLb0bTZqETkiAhM9Iw0y3An2Bg=
go.yaml.in/yaml/v3 v3.0.4/go.mod h1:DhzuOOF2ATzADvBadXxruRBLzYTpT36CKvDb3+aBEFg=
golang.org/x/net v0.52.0 h1:He/TN1l0e4mmR3QqHMT2Xab3Aj3L9qjbhRm78/6jrW0=
golang.org/x/net v0.52.0/go.mod h1:R1MAz7uMZxVMualyPXb+VaqGSa3LIaUqk0eEt3w36Sw=
golang.org/x/sys v0.29.0 h1:TPYlXGxvx1MGTn2GiZDhnjPA9wZzZeGKHHmKhHYvgaU=
golang.org/x/sys v0.29.0/go.mod h1:/VUhepiaJMQUp4+oa/7Zr1D23ma6VTLIYjOOTFZPUcA=
golang.org/x/sys v0.42.0 h1:omrd2nAlyT5ESRdCLYdm3+fMfNFE/+Rf4bDIQImRJeo=
golang.org/x/sys v0.42.0/go.mod h1:4GL1E5IUh+htKOUEOaiffhrAeqysfVGipDYzABqnCmw=
golang.org/x/term v0.13.0 h1:bb+I9cTfFazGW51MZqBVmZy7+JEJMouUHTUSKVQLBek=
golang.org/x/term v0.13.0/go.mod h1:LTmsnFJwVN6bCy1rVCoS+qHT1HhALEFxKncY3WNNh4U=
golang.org/x/term v0.41.0 h1:QCgPso/Q3RTJx2Th4bDLqML4W6iJiaXFq2/ftQF13YU=
golang.org/x/term v0.41.0/go.mod h1:3pfBgksrReYfZ5lvYM0kSO0LIkAl4Yl2bXOkKP7Ec2A=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
rsc.io/qr v0.2.0 h1:6vBLea5/NRMVTz8V66gipeLycZMl/+UlFmk8DvqQ6WY=
rsc.io/qr v0.2.0/go.mod h1:IF+uZjkb9fqyeF/4tlBoynqmQxUoPfWEKh921coOuXs=

```

[⬆ 回到目录](#toc)

## ilink/auth.go

```go
package ilink

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	qrCodeURL     = "https://ilinkai.weixin.qq.com/ilink/bot/get_bot_qrcode?bot_type=3"
	qrStatusURL   = "https://ilinkai.weixin.qq.com/ilink/bot/get_qrcode_status?qrcode="
	statusWait     = "wait"
	statusScanned  = "scaned"
	statusConfirmed = "confirmed"
	statusExpired  = "expired"
)

// FetchQRCode retrieves a new QR code for login.
func FetchQRCode(ctx context.Context) (*QRCodeResponse, error) {
	c := NewUnauthenticatedClient()
	var resp QRCodeResponse
	if err := c.doGet(ctx, qrCodeURL, &resp); err != nil {
		return nil, fmt.Errorf("fetch QR code: %w", err)
	}
	return &resp, nil
}

// PollQRStatus polls for QR code scan status until confirmed or expired.
// It calls onStatus for each status change so the caller can display progress.
func PollQRStatus(ctx context.Context, qrcode string, onStatus func(status string)) (*Credentials, error) {
	c := NewUnauthenticatedClient()
	url := qrStatusURL + qrcode

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		pollCtx, cancel := context.WithTimeout(ctx, 40*time.Second)
		var resp QRStatusResponse
		err := c.doGet(pollCtx, url, &resp)
		cancel()

		if err != nil {
			// Timeout is normal for long-poll, retry
			if ctx.Err() != nil {
				return nil, ctx.Err()
			}
			continue
		}

		if onStatus != nil {
			onStatus(resp.Status)
		}

		switch resp.Status {
		case statusConfirmed:
			creds := &Credentials{
				BotToken:    resp.BotToken,
				ILinkBotID:  resp.ILinkBotID,
				BaseURL:     resp.BaseURL,
				ILinkUserID: resp.ILinkUserID,
			}
			return creds, nil
		case statusExpired:
			return nil, fmt.Errorf("QR code expired")
		case statusWait, statusScanned:
			// Continue polling
		default:
			// Unknown status, continue
		}
	}
}

// AccountsDir returns the directory where account credentials are stored.
func AccountsDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".weclaw", "accounts"), nil
}

// NormalizeAccountID converts raw bot ID to filesystem-safe format.
func NormalizeAccountID(raw string) string {
	s := raw
	for _, ch := range []string{"@", ".", ":"} {
		s = filepath.Clean(s)
		s = replaceAll(s, ch, "-")
	}
	return s
}

func replaceAll(s, old, new string) string {
	for {
		i := indexOf(s, old)
		if i < 0 {
			return s
		}
		s = s[:i] + new + s[i+len(old):]
	}
}

func indexOf(s, sub string) int {
	for i := range s {
		if i+len(sub) <= len(s) && s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}

// SaveCredentials saves credentials to disk under ~/.weclaw/accounts/{id}.json.
func SaveCredentials(creds *Credentials) error {
	dir, err := AccountsDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return fmt.Errorf("create accounts dir: %w", err)
	}

	id := NormalizeAccountID(creds.ILinkBotID)
	path := filepath.Join(dir, id+".json")

	data, err := json.MarshalIndent(creds, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal credentials: %w", err)
	}

	if err := os.WriteFile(path, data, 0o600); err != nil {
		return fmt.Errorf("write credentials: %w", err)
	}
	return nil
}

// LoadAllCredentials loads all saved account credentials.
func LoadAllCredentials() ([]*Credentials, error) {
	dir, err := AccountsDir()
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("read accounts dir: %w", err)
	}

	var result []*Credentials
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".json" {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			continue
		}
		var creds Credentials
		if json.Unmarshal(data, &creds) == nil && creds.BotToken != "" {
			result = append(result, &creds)
		}
	}
	return result, nil
}

// CredentialsPath returns the path for display purposes.
func CredentialsPath() (string, error) {
	return AccountsDir()
}

```

[⬆ 回到目录](#toc)

## ilink/client.go

```go
package ilink

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	defaultBaseURL     = "https://ilinkai.weixin.qq.com"
	longPollTimeout    = 35 * time.Second
	sendTimeout        = 15 * time.Second
)

// Client is an iLink HTTP API client.
type Client struct {
	baseURL    string
	botToken   string
	botID      string
	httpClient *http.Client
	wechatUIN  string
}

// NewClient creates a new iLink API client.
func NewClient(creds *Credentials) *Client {
	baseURL := creds.BaseURL
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	return &Client{
		baseURL:    baseURL,
		botToken:   creds.BotToken,
		botID:      creds.ILinkBotID,
		httpClient: &http.Client{},
		wechatUIN:  generateWechatUIN(),
	}
}

// NewUnauthenticatedClient creates a client without credentials for login flow.
func NewUnauthenticatedClient() *Client {
	return &Client{
		baseURL:    defaultBaseURL,
		httpClient: &http.Client{Timeout: 40 * time.Second},
		wechatUIN:  generateWechatUIN(),
	}
}

// BotID returns the bot's user ID.
func (c *Client) BotID() string {
	return c.botID
}

// GetUpdates performs a long-poll for new messages.
func (c *Client) GetUpdates(ctx context.Context, buf string) (*GetUpdatesResponse, error) {
	reqBody := GetUpdatesRequest{
		GetUpdatesBuf: buf,
		BaseInfo:      BaseInfo{ChannelVersion: "1.0.0"},
	}

	ctx, cancel := context.WithTimeout(ctx, longPollTimeout+5*time.Second)
	defer cancel()

	var resp GetUpdatesResponse
	if err := c.doPost(ctx, "/ilink/bot/getupdates", reqBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendMessage sends a message through iLink.
func (c *Client) SendMessage(ctx context.Context, msg *SendMessageRequest) (*SendMessageResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, sendTimeout)
	defer cancel()

	var resp SendMessageResponse
	if err := c.doPost(ctx, "/ilink/bot/sendmessage", msg, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetConfig fetches bot config for a user (includes typing_ticket).
func (c *Client) GetConfig(ctx context.Context, userID, contextToken string) (*GetConfigResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	req := GetConfigRequest{
		ILinkUserID:  userID,
		ContextToken: contextToken,
		BaseInfo:     BaseInfo{},
	}

	var resp GetConfigResponse
	if err := c.doPost(ctx, "/ilink/bot/getconfig", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendTyping sends a typing indicator to a user.
func (c *Client) SendTyping(ctx context.Context, userID, typingTicket string, status int) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	req := SendTypingRequest{
		ILinkUserID:  userID,
		TypingTicket: typingTicket,
		Status:       status,
		BaseInfo:     BaseInfo{},
	}

	var resp SendTypingResponse
	if err := c.doPost(ctx, "/ilink/bot/sendtyping", req, &resp); err != nil {
		return err
	}
	if resp.Ret != 0 {
		return fmt.Errorf("sendtyping failed: ret=%d errmsg=%s", resp.Ret, resp.ErrMsg)
	}
	return nil
}

// GetUploadURL gets a pre-signed CDN upload URL for media files.
func (c *Client) GetUploadURL(ctx context.Context, req *GetUploadURLRequest) (*GetUploadURLResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, sendTimeout)
	defer cancel()

	var resp GetUploadURLResponse
	if err := c.doPost(ctx, "/ilink/bot/getuploadurl", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BaseURL returns the base URL for CDN operations.
func (c *Client) BaseURL() string {
	return c.baseURL
}

func (c *Client) doPost(ctx context.Context, path string, body interface{}, result interface{}) error {
	data, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+path, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(respBody))
	}

	if err := json.Unmarshal(respBody, result); err != nil {
		return fmt.Errorf("unmarshal response: %w", err)
	}
	return nil
}

func (c *Client) doGet(ctx context.Context, url string, result interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(respBody))
	}

	if err := json.Unmarshal(respBody, result); err != nil {
		return fmt.Errorf("unmarshal response: %w", err)
	}
	return nil
}

func (c *Client) setHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("AuthorizationType", "ilink_bot_token")
	req.Header.Set("Authorization", "Bearer "+c.botToken)
	req.Header.Set("X-WECHAT-UIN", c.wechatUIN)
}

// SetRequestHeaders sets authentication headers on an HTTP request.
// This can be used for CDN downloads that require authentication.
func (c *Client) SetRequestHeaders(req *http.Request) {
	c.setHeaders(req)
}

func generateWechatUIN() string {
	var n uint32
	_ = binary.Read(rand.Reader, binary.LittleEndian, &n)
	s := fmt.Sprintf("%d", n)
	return base64.StdEncoding.EncodeToString([]byte(s))
}

```

[⬆ 回到目录](#toc)

## ilink/monitor.go

```go
package ilink

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	maxConsecutiveFailures = 5
	initialBackoff         = 3 * time.Second
	maxBackoff             = 60 * time.Second
	sessionExpiredBackoff  = 5 * time.Second
	errCodeSessionExpired  = -14
)

// MessageHandler is called for each received message.
type MessageHandler func(ctx context.Context, client *Client, msg WeixinMessage)

// Monitor manages the long-poll loop for receiving messages.
type Monitor struct {
	client        *Client
	handler       MessageHandler
	getUpdatesBuf string
	bufPath       string
	failures      int
	lastActivity  time.Time
}

// NewMonitor creates a new long-poll monitor.
func NewMonitor(client *Client, handler MessageHandler) (*Monitor, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	accountID := NormalizeAccountID(client.BotID())
	bufPath := filepath.Join(home, ".weclaw", "accounts", accountID+".sync.json")

	m := &Monitor{
		client:       client,
		handler:      handler,
		bufPath:      bufPath,
		lastActivity: time.Now(),
	}
	m.loadBuf()
	return m, nil
}

// Run starts the long-poll loop. It blocks until ctx is cancelled.
// Automatically recovers from errors with exponential backoff.
func (m *Monitor) Run(ctx context.Context) error {
	log.Println("[monitor] starting long-poll loop")

	for {
		select {
		case <-ctx.Done():
			log.Println("[monitor] shutting down")
			return ctx.Err()
		default:
		}

		resp, err := m.client.GetUpdates(ctx, m.getUpdatesBuf)
		if err != nil {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			m.failures++
			backoff := m.calcBackoff()
			log.Printf("[monitor] GetUpdates error (%d/%d, backoff=%s): %v",
				m.failures, maxConsecutiveFailures, backoff, err)
			if m.failures == maxConsecutiveFailures {
				log.Printf("[monitor] WARNING: %d consecutive failures. If this persists, run `weclaw login` to re-authenticate.", maxConsecutiveFailures)
			}
			select {
			case <-time.After(backoff):
			case <-ctx.Done():
				return ctx.Err()
			}
			continue
		}

		// Reset failure counter on any successful response
		m.failures = 0
		m.lastActivity = time.Now()

		// Session expired — reset sync buf and reconnect silently
		if resp.ErrCode == errCodeSessionExpired {
			if m.getUpdatesBuf != "" {
				log.Printf("[monitor] session expired, resetting sync buf")
				m.getUpdatesBuf = ""
				m.saveBuf()
			} else {
				// Sync buf already empty but still getting session expired:
				// the bot token itself has expired. The user needs to re-login.
				log.Printf("[monitor] WARNING: WeChat session expired and cannot be auto-recovered. Run `weclaw login` to re-authenticate.")
			}
			select {
			case <-time.After(sessionExpiredBackoff):
			case <-ctx.Done():
				return ctx.Err()
			}
			continue
		}

		// Other server errors
		if resp.Ret != 0 && resp.ErrCode != 0 {
			log.Printf("[monitor] server error: ret=%d errcode=%d errmsg=%s", resp.Ret, resp.ErrCode, resp.ErrMsg)
			continue
		}

		// Update buf for next poll
		if resp.GetUpdatesBuf != "" {
			m.getUpdatesBuf = resp.GetUpdatesBuf
			m.saveBuf()
		}

		// Process messages concurrently — don't block the poll loop
		for _, msg := range resp.Msgs {
			go m.handler(ctx, m.client, msg)
		}
	}
}

// calcBackoff returns an exponential backoff duration capped at maxBackoff.
func (m *Monitor) calcBackoff() time.Duration {
	d := initialBackoff
	for i := 1; i < m.failures; i++ {
		d *= 2
		if d > maxBackoff {
			return maxBackoff
		}
	}
	return d
}

type syncData struct {
	GetUpdatesBuf string `json:"get_updates_buf"`
}

func (m *Monitor) loadBuf() {
	data, err := os.ReadFile(m.bufPath)
	if err != nil {
		return
	}
	var s syncData
	if json.Unmarshal(data, &s) == nil && s.GetUpdatesBuf != "" {
		m.getUpdatesBuf = s.GetUpdatesBuf
		log.Printf("[monitor] loaded sync buf from %s", m.bufPath)
	}
}

func (m *Monitor) saveBuf() {
	dir := filepath.Dir(m.bufPath)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		log.Printf("[monitor] failed to create buf dir: %v", err)
		return
	}
	data, _ := json.Marshal(syncData{GetUpdatesBuf: m.getUpdatesBuf})
	if err := os.WriteFile(m.bufPath, data, 0o600); err != nil {
		log.Printf("[monitor] failed to save buf: %v", err)
	}
}

// FormatMessageSummary returns a short description of a message for logging.
func FormatMessageSummary(msg WeixinMessage) string {
	text := ""
	for _, item := range msg.ItemList {
		if item.Type == ItemTypeText && item.TextItem != nil {
			text = item.TextItem.Text
			break
		}
	}
	if len(text) > 50 {
		text = text[:50] + "..."
	}
	return fmt.Sprintf("from=%s type=%d state=%d text=%q", msg.FromUserID, msg.MessageType, msg.MessageState, text)
}

```

[⬆ 回到目录](#toc)

## ilink/types.go

```go
package ilink

// Message types
const (
	MessageTypeNone = 0
	MessageTypeUser = 1
	MessageTypeBot  = 2
)

// Message states
const (
	MessageStateNew        = 0
	MessageStateGenerating = 1
	MessageStateFinish     = 2
)

// Item types
const (
	ItemTypeNone  = 0
	ItemTypeText  = 1
	ItemTypeImage = 2
	ItemTypeVoice = 3
	ItemTypeFile  = 4
	ItemTypeVideo = 5
)

// QRCodeResponse is the response from get_bot_qrcode.
type QRCodeResponse struct {
	QRCode        string `json:"qrcode"`
	QRCodeImgContent string `json:"qrcode_img_content"`
}

// QRStatusResponse is the response from get_qrcode_status.
type QRStatusResponse struct {
	Status     string `json:"status"`
	BotToken   string `json:"bot_token"`
	ILinkBotID string `json:"ilink_bot_id"`
	BaseURL    string `json:"baseurl"`
	ILinkUserID string `json:"ilink_user_id"`
}

// Credentials stores login session data.
type Credentials struct {
	BotToken   string `json:"bot_token"`
	ILinkBotID string `json:"ilink_bot_id"`
	BaseURL    string `json:"baseurl"`
	ILinkUserID string `json:"ilink_user_id"`
}

// BaseInfo is included in request bodies.
type BaseInfo struct {
	ChannelVersion string `json:"channel_version,omitempty"`
}

// GetUpdatesRequest is the body for getupdates.
type GetUpdatesRequest struct {
	GetUpdatesBuf string   `json:"get_updates_buf"`
	BaseInfo      BaseInfo `json:"base_info"`
}

// GetUpdatesResponse is the response from getupdates.
type GetUpdatesResponse struct {
	Ret                 int              `json:"ret"`
	ErrCode             int              `json:"errcode,omitempty"`
	ErrMsg              string           `json:"errmsg,omitempty"`
	Msgs                []WeixinMessage  `json:"msgs"`
	GetUpdatesBuf       string           `json:"get_updates_buf"`
	LongPollingTimeoutMs int             `json:"longpolling_timeout_ms,omitempty"`
}

// WeixinMessage represents a message from WeChat.
type WeixinMessage struct {
	Seq          int           `json:"seq,omitempty"`
	MessageID    int64         `json:"message_id,omitempty"`
	FromUserID   string        `json:"from_user_id"`
	ToUserID     string        `json:"to_user_id"`
	MessageType  int           `json:"message_type"`
	MessageState int           `json:"message_state"`
	ItemList     []MessageItem `json:"item_list"`
	ContextToken string        `json:"context_token"`
}

// MessageItem is a single item in a message.
type MessageItem struct {
	Type      int        `json:"type"`
	TextItem  *TextItem  `json:"text_item,omitempty"`
	ImageItem *ImageItem `json:"image_item,omitempty"`
	VoiceItem *VoiceItem `json:"voice_item,omitempty"`
	VideoItem *VideoItem `json:"video_item,omitempty"`
	FileItem  *FileItem  `json:"file_item,omitempty"`
}

// CDN media type constants.
const (
	CDNMediaTypeImage = 1
	CDNMediaTypeVideo = 2
	CDNMediaTypeFile  = 3
)

// GetUploadURLRequest is the body for getuploadurl.
type GetUploadURLRequest struct {
	FileKey     string   `json:"filekey"`
	MediaType   int      `json:"media_type"`
	ToUserID    string   `json:"to_user_id"`
	RawSize     int      `json:"rawsize"`
	RawFileMD5  string   `json:"rawfilemd5"`
	FileSize    int      `json:"filesize"`
	NoNeedThumb bool     `json:"no_need_thumb"`
	AESKey      string   `json:"aeskey"`
	BaseInfo    BaseInfo `json:"base_info"`
}

// GetUploadURLResponse is the response from getuploadurl.
type GetUploadURLResponse struct {
	Ret           int    `json:"ret"`
	ErrMsg        string `json:"errmsg,omitempty"`
	UploadParam   string `json:"upload_param"`
	UploadFullURL string `json:"upload_full_url,omitempty"`
}

// TextItem holds text content.
type TextItem struct {
	Text string `json:"text"`
}

// MediaInfo holds CDN media reference for uploaded files.
type MediaInfo struct {
	EncryptQueryParam string `json:"encrypt_query_param"`
	AESKey            string `json:"aes_key"`    // base64-encoded
	EncryptType       int    `json:"encrypt_type"` // 1 = AES-128-ECB
}

// VoiceItem holds voice content.
type VoiceItem struct {
	Media         *MediaInfo `json:"media,omitempty"`
	VoiceSize     int        `json:"voice_size,omitempty"`
	EncodeType    int        `json:"encode_type,omitempty"`    // 1=pcm 2=adpcm 3=feature 4=speex 5=amr 6=silk 7=mp3
	BitsPerSample int       `json:"bits_per_sample,omitempty"`
	SampleRate    int        `json:"sample_rate,omitempty"`    // Hz
	Playtime      int        `json:"playtime,omitempty"`       // duration in milliseconds
	Text          string     `json:"text,omitempty"`           // speech-to-text transcription from WeChat
}

// ImageItem holds image content.
type ImageItem struct {
	URL     string     `json:"url,omitempty"`
	Media   *MediaInfo `json:"media,omitempty"`
	MidSize int        `json:"mid_size,omitempty"` // ciphertext size
}

// VideoItem holds video content.
type VideoItem struct {
	Media     *MediaInfo `json:"media,omitempty"`
	VideoSize int        `json:"video_size,omitempty"`
}

// FileItem holds file content.
type FileItem struct {
	Media    *MediaInfo `json:"media,omitempty"`
	FileName string     `json:"file_name,omitempty"`
	Len      string     `json:"len,omitempty"` // plaintext size as string
}

// SendMessageRequest is the body for sendmessage.
type SendMessageRequest struct {
	Msg      SendMsg  `json:"msg"`
	BaseInfo BaseInfo `json:"base_info"`
}

// SendMsg is the message payload for sending.
type SendMsg struct {
	FromUserID   string        `json:"from_user_id"`
	ToUserID     string        `json:"to_user_id"`
	ClientID     string        `json:"client_id"`
	MessageType  int           `json:"message_type"`
	MessageState int           `json:"message_state"`
	ItemList     []MessageItem `json:"item_list"`
	ContextToken string        `json:"context_token"`
}

// SendMessageResponse is the response from sendmessage.
type SendMessageResponse struct {
	Ret    int    `json:"ret"`
	ErrMsg string `json:"errmsg,omitempty"`
}

// Typing status constants.
const (
	TypingStatusTyping = 1
	TypingStatusCancel = 2
)

// GetConfigRequest is the body for getconfig.
type GetConfigRequest struct {
	ILinkUserID  string   `json:"ilink_user_id"`
	ContextToken string   `json:"context_token,omitempty"`
	BaseInfo     BaseInfo `json:"base_info"`
}

// GetConfigResponse is the response from getconfig.
type GetConfigResponse struct {
	Ret           int    `json:"ret"`
	ErrMsg        string `json:"errmsg,omitempty"`
	TypingTicket  string `json:"typing_ticket,omitempty"`
}

// SendTypingRequest is the body for sendtyping.
type SendTypingRequest struct {
	ILinkUserID  string   `json:"ilink_user_id"`
	TypingTicket string   `json:"typing_ticket"`
	Status       int      `json:"status"`
	BaseInfo     BaseInfo `json:"base_info"`
}

// SendTypingResponse is the response from sendtyping.
type SendTypingResponse struct {
	Ret    int    `json:"ret"`
	ErrMsg string `json:"errmsg,omitempty"`
}

```

[⬆ 回到目录](#toc)

## install.sh

```bash
#!/bin/sh
set -e

REPO="fastclaw-ai/weclaw"
BINARY="weclaw"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

# Detect OS
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
case "$OS" in
  darwin|linux) ;;
  *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

# Detect architecture
ARCH=$(uname -m)
case "$ARCH" in
  x86_64|amd64) ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

echo "Detected: ${OS}/${ARCH}"

# Get latest version
echo "Fetching latest release..."
VERSION=$(curl -fsSL -H "User-Agent: weclaw-installer" "https://api.github.com/repos/${REPO}/releases/latest" | sed -n 's/.*"tag_name" *: *"\([^"]*\)".*/\1/p')

if [ -z "$VERSION" ]; then
  echo "Error: could not determine latest version. Is there a release on GitHub?"
  exit 1
fi

echo "Latest version: ${VERSION}"

# Download
FILENAME="${BINARY}_${OS}_${ARCH}"
URL="https://github.com/${REPO}/releases/download/${VERSION}/${FILENAME}"

echo "Downloading ${URL}..."
TMP=$(mktemp)
curl -fsSL -o "$TMP" "$URL"

# Install
chmod +x "$TMP"
if [ -d "$INSTALL_DIR" ] && [ -w "$INSTALL_DIR" ]; then
  mv "$TMP" "${INSTALL_DIR}/${BINARY}"
else
  echo "Installing to ${INSTALL_DIR} (requires sudo)..."
  sudo mkdir -p "$INSTALL_DIR"
  sudo mv "$TMP" "${INSTALL_DIR}/${BINARY}"
fi

# Clear macOS quarantine attributes
if [ "$OS" = "darwin" ]; then
  xattr -d com.apple.quarantine "${INSTALL_DIR}/${BINARY}" 2>/dev/null || true
  xattr -d com.apple.provenance "${INSTALL_DIR}/${BINARY}" 2>/dev/null || true
fi

echo ""
echo "weclaw ${VERSION} installed to ${INSTALL_DIR}/${BINARY}"
echo ""
echo "Get started:"
echo "  weclaw start"

```

[⬆ 回到目录](#toc)

## main.go

```go
package main

import "github.com/fastclaw-ai/weclaw/cmd"

func main() {
	cmd.Execute()
}

```

[⬆ 回到目录](#toc)

## messaging/attachment.go

```go
package messaging

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
)

var supportedAttachmentExts = []string{
	".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx",
	".zip", ".txt", ".csv",
	".png", ".jpg", ".jpeg", ".gif", ".webp",
	".mp4", ".mov",
}

func defaultAttachmentWorkspace() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return filepath.Clean(os.TempDir())
	}
	return filepath.Join(home, ".weclaw", "workspace")
}

func extractLocalAttachmentPaths(text string) []string {
	var paths []string
	seen := make(map[string]struct{})

	for _, line := range strings.Split(text, "\n") {
		candidate := strings.TrimSpace(line)
		if candidate == "" || !filepath.IsAbs(candidate) {
			continue
		}
		if !isSupportedAttachmentPath(candidate) {
			continue
		}
		info, err := os.Stat(candidate)
		if err != nil || info.IsDir() {
			continue
		}
		if _, ok := seen[candidate]; ok {
			continue
		}
		seen[candidate] = struct{}{}
		paths = append(paths, candidate)
	}

	return paths
}

func isAllowedAttachmentPath(path string, allowedRoots []string) bool {
	cleanPath, err := canonicalizePath(path, true)
	if err != nil {
		return false
	}

	for _, root := range allowedRoots {
		if root == "" {
			continue
		}
		cleanRoot, err := canonicalizePath(root, false)
		if err != nil {
			continue
		}
		rel, err := filepath.Rel(cleanRoot, cleanPath)
		if err != nil {
			continue
		}
		if rel == "." || (rel != ".." && !strings.HasPrefix(rel, ".."+string(os.PathSeparator))) {
			return true
		}
	}

	return false
}

func rewriteReplyWithAttachmentResults(reply string, sentPaths, failedPaths []string) string {
	sentMap := make(map[string]string, len(sentPaths))
	for _, path := range sentPaths {
		sentMap[path] = "已发送附件：" + filepath.Base(path)
	}

	lines := strings.Split(reply, "\n")
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if replacement, ok := sentMap[trimmed]; ok {
			lines[i] = replacement
		}
	}

	rewritten := strings.Join(lines, "\n")

	var failureLines []string
	seenFailures := make(map[string]struct{})
	for _, path := range failedPaths {
		if _, ok := seenFailures[path]; ok {
			continue
		}
		seenFailures[path] = struct{}{}
		failureLines = append(failureLines, "附件发送失败："+filepath.Base(path))
	}
	if len(failureLines) == 0 {
		return rewritten
	}
	if strings.TrimSpace(rewritten) == "" {
		return strings.Join(failureLines, "\n")
	}
	return rewritten + "\n" + strings.Join(failureLines, "\n")
}

func isSupportedAttachmentPath(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return slices.Contains(supportedAttachmentExts, ext)
}

func canonicalizePath(path string, mustExist bool) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	if realPath, err := filepath.EvalSymlinks(absPath); err == nil {
		return filepath.Clean(realPath), nil
	} else if mustExist {
		return "", err
	}
	return filepath.Clean(absPath), nil
}

```

[⬆ 回到目录](#toc)

## messaging/attachment_test.go

```go
package messaging

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExtractLocalAttachmentPaths(t *testing.T) {
	dir := t.TempDir()
	pdfPath := filepath.Join(dir, "report.pdf")
	txtPath := filepath.Join(dir, "notes.txt")
	if err := os.WriteFile(pdfPath, []byte("pdf"), 0o644); err != nil {
		t.Fatalf("write pdf: %v", err)
	}
	if err := os.WriteFile(txtPath, []byte("txt"), 0o644); err != nil {
		t.Fatalf("write txt: %v", err)
	}

	reply := strings.Join([]string{
		"这里是内联路径，不应该命中 " + pdfPath,
		pdfPath,
		"1. " + txtPath,
		txtPath,
		"file://" + pdfPath,
		filepath.Join(dir, "missing.pdf"),
		filepath.Join(dir, "folder"),
	}, "\n")

	got := extractLocalAttachmentPaths(reply)
	if len(got) != 2 {
		t.Fatalf("expected 2 paths, got %d (%v)", len(got), got)
	}
	if got[0] != pdfPath {
		t.Fatalf("got[0] = %q, want %q", got[0], pdfPath)
	}
	if got[1] != txtPath {
		t.Fatalf("got[1] = %q, want %q", got[1], txtPath)
	}
}

func TestIsAllowedAttachmentPath(t *testing.T) {
	workspaceRoot := filepath.Join(t.TempDir(), "workspace")
	otherRoot := filepath.Join(t.TempDir(), "other")
	if err := os.MkdirAll(workspaceRoot, 0o755); err != nil {
		t.Fatalf("mkdir workspace: %v", err)
	}
	if err := os.MkdirAll(otherRoot, 0o755); err != nil {
		t.Fatalf("mkdir other: %v", err)
	}

	allowedPath := filepath.Join(workspaceRoot, "artifacts", "report.pdf")
	deniedPath := filepath.Join(otherRoot, "report.pdf")
	if err := os.MkdirAll(filepath.Dir(allowedPath), 0o755); err != nil {
		t.Fatalf("mkdir allowed dir: %v", err)
	}
	if err := os.WriteFile(allowedPath, []byte("ok"), 0o644); err != nil {
		t.Fatalf("write allowed file: %v", err)
	}
	if err := os.WriteFile(deniedPath, []byte("no"), 0o644); err != nil {
		t.Fatalf("write denied file: %v", err)
	}

	if !isAllowedAttachmentPath(allowedPath, []string{workspaceRoot}) {
		t.Fatalf("expected %q to be allowed", allowedPath)
	}
	if isAllowedAttachmentPath(deniedPath, []string{workspaceRoot}) {
		t.Fatalf("expected %q to be denied", deniedPath)
	}
}

func TestRewriteReplyWithAttachmentResults(t *testing.T) {
	sentPath := "/tmp/report.pdf"
	failedPath := "/tmp/archive.zip"
	reply := strings.Join([]string{
		"已生成文件：",
		sentPath,
		"这里再次内联提到 " + sentPath + "，不应该被替换。",
		failedPath,
	}, "\n")

	got := rewriteReplyWithAttachmentResults(reply, []string{sentPath}, []string{failedPath})

	if strings.Contains(got, "\n"+sentPath+"\n") {
		t.Fatalf("expected sent path line to be replaced, got %q", got)
	}
	if !strings.Contains(got, "已发送附件：report.pdf") {
		t.Fatalf("expected sent replacement, got %q", got)
	}
	if !strings.Contains(got, "这里再次内联提到 "+sentPath+"，不应该被替换。") {
		t.Fatalf("expected inline path to remain, got %q", got)
	}
	if !strings.Contains(got, failedPath) {
		t.Fatalf("expected failed path to remain, got %q", got)
	}
	if !strings.Contains(got, "附件发送失败：archive.zip") {
		t.Fatalf("expected failure note, got %q", got)
	}
}

```

[⬆ 回到目录](#toc)

## messaging/cdn.go

```go
package messaging

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/fastclaw-ai/weclaw/ilink"
)

const cdnBaseURL = "https://novac2c.cdn.weixin.qq.com/c2c"

// UploadedFile holds the result of a CDN upload.
type UploadedFile struct {
	DownloadParam string // encrypted query param for download
	AESKeyHex     string // hex-encoded AES key
	FileSize      int    // plaintext size
	CipherSize    int    // ciphertext size
}

// UploadFileToCDN encrypts and uploads a file to the WeChat CDN.
func UploadFileToCDN(ctx context.Context, client *ilink.Client, data []byte, toUserID string, mediaType int) (*UploadedFile, error) {
	// Generate random filekey and AES key
	filekey := make([]byte, 16)
	aeskey := make([]byte, 16)
	if _, err := rand.Read(filekey); err != nil {
		return nil, fmt.Errorf("generate filekey: %w", err)
	}
	if _, err := rand.Read(aeskey); err != nil {
		return nil, fmt.Errorf("generate aeskey: %w", err)
	}

	filekeyHex := hex.EncodeToString(filekey)
	aeskeyHex := hex.EncodeToString(aeskey)

	// Calculate MD5 of plaintext
	hash := md5.Sum(data)
	rawMD5 := hex.EncodeToString(hash[:])

	// Calculate ciphertext size (PKCS7 padding)
	cipherSize := aesECBPaddedSize(len(data))

	// Get upload URL from iLink API
	uploadReq := &ilink.GetUploadURLRequest{
		FileKey:     filekeyHex,
		MediaType:   mediaType,
		ToUserID:    toUserID,
		RawSize:     len(data),
		RawFileMD5:  rawMD5,
		FileSize:    cipherSize,
		NoNeedThumb: true,
		AESKey:      aeskeyHex,
		BaseInfo:    ilink.BaseInfo{},
	}

	uploadResp, err := client.GetUploadURL(ctx, uploadReq)
	if err != nil {
		return nil, fmt.Errorf("get upload URL: %w", err)
	}
	if uploadResp.Ret != 0 {
		return nil, fmt.Errorf("get upload URL failed: ret=%d errmsg=%s", uploadResp.Ret, uploadResp.ErrMsg)
	}

	// Encrypt data with AES-128-ECB
	encrypted, err := encryptAESECB(data, aeskey)
	if err != nil {
		return nil, fmt.Errorf("encrypt: %w", err)
	}

	// Upload to CDN: prefer server-provided full URL, fall back to param-based construction
	cdnURL := strings.TrimSpace(uploadResp.UploadFullURL)
	if cdnURL == "" {
		if uploadResp.UploadParam == "" {
			return nil, fmt.Errorf("getuploadurl returned no upload URL (need upload_full_url or upload_param)")
		}
		cdnURL = fmt.Sprintf("%s/upload?encrypted_query_param=%s&filekey=%s",
			cdnBaseURL, url.QueryEscape(uploadResp.UploadParam), url.QueryEscape(filekeyHex))
	}

	downloadParam, err := uploadToCDN(ctx, encrypted, cdnURL)
	if err != nil {
		return nil, fmt.Errorf("CDN upload: %w", err)
	}

	return &UploadedFile{
		DownloadParam: downloadParam,
		AESKeyHex:     aeskeyHex,
		FileSize:      len(data),
		CipherSize:    cipherSize,
	}, nil
}

// AESKeyToBase64 converts a hex AES key to base64 format for message items.
func AESKeyToBase64(hexKey string) string {
	return base64.StdEncoding.EncodeToString([]byte(hexKey))
}

// DownloadFileFromCDN downloads and decrypts a file from the WeChat CDN.
func DownloadFileFromCDN(ctx context.Context, encryptQueryParam, aesKeyBase64 string) ([]byte, error) {
	// Decode AES key: base64 -> hex string -> raw bytes
	aesKeyHexBytes, err := base64.StdEncoding.DecodeString(aesKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("decode AES key base64: %w", err)
	}
	aesKey, err := hex.DecodeString(string(aesKeyHexBytes))
	if err != nil {
		return nil, fmt.Errorf("decode AES key hex: %w", err)
	}

	// Download encrypted data from CDN
	downloadURL := fmt.Sprintf("%s/download?encrypted_query_param=%s",
		cdnBaseURL, url.QueryEscape(encryptQueryParam))

	reqCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, downloadURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create download request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("download from CDN: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("CDN download HTTP %d: %s", resp.StatusCode, string(body))
	}

	encrypted, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read CDN response: %w", err)
	}

	// Decrypt AES-128-ECB
	return decryptAESECB(encrypted, aesKey)
}

// decryptAESECB decrypts data encrypted with AES-128-ECB and removes PKCS7 padding.
func decryptAESECB(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of block size")
	}

	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += aes.BlockSize {
		block.Decrypt(plaintext[i:i+aes.BlockSize], ciphertext[i:i+aes.BlockSize])
	}

	// Remove PKCS7 padding
	if len(plaintext) == 0 {
		return plaintext, nil
	}
	padLen := int(plaintext[len(plaintext)-1])
	if padLen > aes.BlockSize || padLen == 0 {
		return nil, fmt.Errorf("invalid PKCS7 padding")
	}
	return plaintext[:len(plaintext)-padLen], nil
}

func uploadToCDN(ctx context.Context, encrypted []byte, cdnURL string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, cdnURL, bytes.NewReader(encrypted))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("CDN upload HTTP %d: %s", resp.StatusCode, string(body))
	}

	downloadParam := resp.Header.Get("X-Encrypted-Param")
	if downloadParam == "" {
		return "", fmt.Errorf("CDN upload: missing X-Encrypted-Param header")
	}

	return downloadParam, nil
}

// encryptAESECB encrypts data using AES-128-ECB with PKCS7 padding.
func encryptAESECB(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// PKCS7 padding
	padLen := aes.BlockSize - (len(plaintext) % aes.BlockSize)
	padded := make([]byte, len(plaintext)+padLen)
	copy(padded, plaintext)
	for i := len(plaintext); i < len(padded); i++ {
		padded[i] = byte(padLen)
	}

	// ECB mode: encrypt each block independently
	encrypted := make([]byte, len(padded))
	for i := 0; i < len(padded); i += aes.BlockSize {
		block.Encrypt(encrypted[i:i+aes.BlockSize], padded[i:i+aes.BlockSize])
	}

	return encrypted, nil
}

func aesECBPaddedSize(plaintextSize int) int {
	return (plaintextSize/aes.BlockSize + 1) * aes.BlockSize
}

```

[⬆ 回到目录](#toc)

## messaging/handler.go

```go
package messaging

import (
	"context"
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fastclaw-ai/weclaw/agent"
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
	mu            sync.RWMutex
	defaultName   string
	agents        map[string]agent.Agent // name -> running agent
	agentMetas    []AgentMeta            // all configured agents (for /status)
	agentWorkDirs map[string]string      // agent name -> configured/runtime cwd
	customAliases map[string]string      // custom alias -> agent name (from config)
	factory       AgentFactory
	saveDefault   SaveDefaultFunc
	contextTokens sync.Map   // map[userID]contextToken
	saveDir       string     // directory to save images/files to
	seenMsgs      sync.Map   // map[int64]time.Time — dedup by message_id
	progressCtx   *progressContext // current request context for progress notifications
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

// NewHandler creates a new message handler.
func NewHandler(factory AgentFactory, saveDefault SaveDefaultFunc) *Handler {
	return &Handler{
		agents:        make(map[string]agent.Agent),
		agentWorkDirs: make(map[string]string),
		factory:       factory,
		saveDefault:   saveDefault,
	}
}

// SetSaveDir sets the directory for saving images and files.
func (h *Handler) SetSaveDir(dir string) {
	h.saveDir = dir
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
func (h *Handler) resolveAlias(name string) string {
	h.mu.RLock()
	custom := h.customAliases
	h.mu.RUnlock()
	if custom != nil {
		if full, ok := custom[name]; ok {
			return full
		}
	}
	if full, ok := agentAliases[name]; ok {
		return full
	}
	return name
}

// parseCommand checks if text starts with "/" or "@" followed by agent name(s).
// Supports multiple agents: "@cc @cx hello" returns (["claude","codex"], "hello").
// Returns (agentNames, actualMessage). Aliases are resolved automatically.
// If no command prefix, returns (nil, originalText).
func (h *Handler) parseCommand(text string) ([]string, string) {
	if !strings.HasPrefix(text, "/") && !strings.HasPrefix(text, "@") {
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

	// Built-in commands (no typing needed)
	if trimmed == "/info" {
		reply := h.buildStatus()
		if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
			log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
		}
		return
	} else if trimmed == "/help" {
		reply := buildHelpText()
		if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
			log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
		}
		return
	} else if trimmed == "/new" || trimmed == "/clear" {
		reply := h.resetDefaultSession(ctx, msg.FromUserID)
		if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
			log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
		}
		return
	} else if strings.HasPrefix(trimmed, "/cwd") {
		reply := h.handleCwd(trimmed)
		if err := SendTextReply(ctx, client, msg.FromUserID, reply, msg.ContextToken, clientID); err != nil {
			log.Printf("[handler] failed to send reply to %s: %v", msg.FromUserID, err)
		}
		return
	}

	// Route: "/agentname message" or "@agent1 @agent2 message" -> specific agent(s)
	agentNames, message := h.parseCommand(text)

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
		// No known agents -> forward entire text to default agent
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
}

func buildHelpText() string {
	return `Available commands:
@agent or /agent - Switch default agent
@agent msg or /agent msg - Send to a specific agent
@a @b msg - Broadcast to multiple agents
/new or /clear - Start a new session
/cwd /path - Switch workspace directory
/info - Show current agent info
/help - Show this help message

Aliases: /cc(claude) /cx(codex) /cs(cursor) /km(kimi) /gm(gemini) /oc(openclaw) /ocd(opencode) /pi(pi) /cp(copilot) /dr(droid) /if(iflow) /kr(kiro) /qw(qwen)`
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
		// No encryption key — data is plaintext
		fileData = encryptedData
		log.Printf("[handler] no AES key, using raw data (no decryption)")
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

```

[⬆ 回到目录](#toc)

## messaging/handler_test.go

```go
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

```

[⬆ 回到目录](#toc)

## messaging/linkhoard.go

```go
package messaging

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"bufio"

	"github.com/google/uuid"
	"golang.org/x/net/html"
)

var reURL = regexp.MustCompile(`https?://\S+`)

// IsURL checks if the text is (or starts with) a URL.
func IsURL(text string) bool {
	trimmed := strings.TrimSpace(text)
	return strings.HasPrefix(trimmed, "http://") || strings.HasPrefix(trimmed, "https://")
}

// ExtractURL extracts the first URL from text.
func ExtractURL(text string) string {
	match := reURL.FindString(text)
	return match
}

// LinkMetadata holds extracted metadata from a web page.
type LinkMetadata struct {
	Title       string
	Description string
	Author      string
	OGImage     string
	Published   string
	Body        string
}

// FetchLinkMetadata fetches a URL and extracts metadata from the HTML.
func FetchLinkMetadata(ctx context.Context, rawURL string) (*LinkMetadata, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")
	req.Header.Set("Referer", "https://mp.weixin.qq.com/")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parse HTML: %w", err)
	}

	meta := &LinkMetadata{}
	extractMeta(doc, meta)

	// Fallback title from URL if empty
	if meta.Title == "" {
		meta.Title = rawURL
	}

	return meta, nil
}

// extractMeta walks the HTML tree and extracts metadata.
func extractMeta(n *html.Node, meta *LinkMetadata) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "meta":
			handleMeta(n, meta)
		case "title":
			if meta.Title == "" && n.FirstChild != nil {
				meta.Title = strings.TrimSpace(n.FirstChild.Data)
			}
		case "div":
			// WeChat article body
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == "js_content" {
					meta.Body = extractNodeText(n)
					return
				}
			}
		case "article":
			if meta.Body == "" {
				meta.Body = extractNodeText(n)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractMeta(c, meta)
	}
}

// handleMeta extracts og: and other meta tag values.
func handleMeta(n *html.Node, meta *LinkMetadata) {
	var property, name, content string
	for _, a := range n.Attr {
		switch a.Key {
		case "property":
			property = a.Val
		case "name":
			name = a.Val
		case "content":
			content = a.Val
		}
	}
	if content == "" {
		return
	}
	switch {
	case property == "og:title" && meta.Title == "":
		meta.Title = content
	case property == "og:description" && meta.Description == "":
		meta.Description = content
	case property == "og:image" && meta.OGImage == "":
		meta.OGImage = content
	case property == "article:published_time" && meta.Published == "":
		meta.Published = content
	case name == "author" && meta.Author == "":
		meta.Author = content
	case name == "description" && meta.Description == "":
		meta.Description = content
	}
}

// extractText recursively extracts visible text from an HTML node.
func extractNodeText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var sb strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && (c.Data == "script" || c.Data == "style") {
			continue
		}
		text := extractNodeText(c)
		if text != "" {
			// Add paragraph breaks for block elements
			if c.Type == html.ElementNode {
				switch c.Data {
				case "p", "div", "br", "h1", "h2", "h3", "h4", "h5", "h6", "li", "section":
					sb.WriteString("\n\n")
				}
			}
			sb.WriteString(text)
		}
	}
	return sb.String()
}

// sanitizeFileName removes characters unsafe for filenames.
func sanitizeFileName(name string) string {
	replacer := strings.NewReplacer(
		"/", "", "\\", "", ":", "", "*", "",
		"?", "", "\"", "", "<", "", ">", "", "|", "",
	)
	result := replacer.Replace(name)
	// Trim and limit length
	result = strings.TrimSpace(result)
	if len(result) > 200 {
		result = result[:200]
	}
	if result == "" {
		result = "untitled"
	}
	return result
}

// isWeChatURL checks if a URL is a WeChat article.
func isWeChatURL(rawURL string) bool {
	return strings.Contains(rawURL, "mp.weixin.qq.com") || strings.Contains(rawURL, "weixin.qq.com/s/")
}

// FetchViaJina fetches a URL via Jina Reader API and returns metadata + markdown body.
func FetchViaJina(ctx context.Context, rawURL string) (*LinkMetadata, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	jinaURL := "https://r.jina.ai/" + rawURL
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, jinaURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "text/plain")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Jina HTTP %d", resp.StatusCode)
	}

	meta := &LinkMetadata{}
	scanner := bufio.NewScanner(resp.Body)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	// Parse Jina header lines: "Title:", "URL Source:", "Published Time:", then "Markdown Content:"
	inBody := false
	var body strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		if inBody {
			body.WriteString(line)
			body.WriteString("\n")
			continue
		}
		if strings.HasPrefix(line, "Title: ") {
			meta.Title = strings.TrimPrefix(line, "Title: ")
		} else if strings.HasPrefix(line, "Published Time: ") {
			meta.Published = strings.TrimPrefix(line, "Published Time: ")
		} else if line == "Markdown Content:" {
			inBody = true
		}
	}

	if meta.Title == "" {
		meta.Title = rawURL
	}
	meta.Body = strings.TrimSpace(body.String())

	// Check for Jina failure (CAPTCHA, empty content)
	if meta.Body == "" || strings.Contains(meta.Body, "环境异常") || strings.Contains(meta.Body, "CAPTCHA") {
		return nil, fmt.Errorf("Jina returned empty or blocked content")
	}

	return meta, nil
}

// SaveLinkToLinkhoard fetches a URL and saves it as a Linkhoard-compatible markdown file.
// WeChat articles use direct fetch with browser headers; other sites use Jina Reader.
// Returns the link metadata for further processing (e.g., AI analysis).
func SaveLinkToLinkhoard(ctx context.Context, saveDir, rawURL string) (*LinkMetadata, error) {
	var meta *LinkMetadata
	var err error

	if isWeChatURL(rawURL) {
		meta, err = FetchLinkMetadata(ctx, rawURL)
	} else {
		meta, err = FetchViaJina(ctx, rawURL)
		if err != nil {
			// Fallback to direct fetch
			log.Printf("[linkhoard] Jina failed (%v), falling back to direct fetch", err)
			meta, err = FetchLinkMetadata(ctx, rawURL)
		}
	}
	if err != nil {
		return nil, fmt.Errorf("fetch failed: %w", err)
	}

	// Ensure save directory exists
	if err := os.MkdirAll(saveDir, 0o755); err != nil {
		return nil, fmt.Errorf("create dir: %w", err)
	}

	// Build frontmatter
	title := sanitizeFileName(meta.Title)
	created := time.Now().UTC().Format(time.RFC3339)
	itemID := uuid.New().String()

	// Normalize body text
	body := strings.TrimSpace(meta.Body)
	// Collapse excessive newlines
	for strings.Contains(body, "\n\n\n") {
		body = strings.ReplaceAll(body, "\n\n\n", "\n\n")
	}

	// Build author field
	authorField := "author: []\n"
	if meta.Author != "" {
		authorField = fmt.Sprintf("author:\n  - '[[%s]]'\n", meta.Author)
	}

	// Build markdown content
	var sb strings.Builder
	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("title: '%s'\n", strings.ReplaceAll(meta.Title, "'", "''")))
	sb.WriteString(fmt.Sprintf("source: '%s'\n", rawURL))
	sb.WriteString(fmt.Sprintf("published: '%s'\n", meta.Published))
	sb.WriteString(fmt.Sprintf("created: '%s'\n", created))
	sb.WriteString(fmt.Sprintf("description: '%s'\n", strings.ReplaceAll(meta.Description, "'", "''")))
	if meta.OGImage != "" {
		sb.WriteString(fmt.Sprintf("openGraphImage: '%s'\n", meta.OGImage))
	}
	sb.WriteString(authorField)
	sb.WriteString("---\n\n")
	if body != "" {
		sb.WriteString(body)
		sb.WriteString("\n")
	}

	// Write markdown file
	filePath := filepath.Join(saveDir, title+".md")
	if err := os.WriteFile(filePath, []byte(sb.String()), 0o644); err != nil {
		return nil, fmt.Errorf("write file: %w", err)
	}

	// Write sidecar
	sidecarPath := filePath + ".sidecar.md"
	sidecarContent := fmt.Sprintf("---\nid: %s\n---\n", itemID)
	if err := os.WriteFile(sidecarPath, []byte(sidecarContent), 0o644); err != nil {
		log.Printf("[linkhoard] failed to write sidecar: %v", err)
	}

	log.Printf("[linkhoard] saved %q to %s", meta.Title, filePath)
	return meta, nil
}

```

[⬆ 回到目录](#toc)

## messaging/markdown.go

```go
package messaging

import (
	"regexp"
	"strings"
)

var (
	// Code blocks: strip fences, keep code content
	reCodeBlock = regexp.MustCompile("(?s)```[^\n]*\n?(.*?)```")
	// Inline code: strip backticks, keep content
	reInlineCode = regexp.MustCompile("`([^`]+)`")
	// Images: remove entirely
	reImage = regexp.MustCompile(`!\[[^\]]*\]\([^)]*\)`)
	// Links: keep display text only
	reLink = regexp.MustCompile(`\[([^\]]+)\]\([^)]*\)`)
	// Table separator rows: remove
	reTableSep = regexp.MustCompile(`(?m)^\|[\s:|\-]+\|$`)
	// Table rows: convert pipe-delimited to space-delimited
	reTableRow = regexp.MustCompile(`(?m)^\|(.+)\|$`)
	// Headers: remove # prefix
	reHeader = regexp.MustCompile(`(?m)^#{1,6}\s+`)
	// Bold: **text** or __text__
	reBold = regexp.MustCompile(`\*\*(.+?)\*\*|__(.+?)__`)
	// Italic: *text* or _text_
	reItalic = regexp.MustCompile(`(?:^|[^*])\*([^*]+)\*(?:[^*]|$)|(?:^|[^_])_([^_]+)_(?:[^_]|$)`)
	// Strikethrough: ~~text~~
	reStrike = regexp.MustCompile(`~~(.+?)~~`)
	// Blockquote: > prefix
	reBlockquote = regexp.MustCompile(`(?m)^>\s?`)
	// Horizontal rule
	reHR = regexp.MustCompile(`(?m)^[-*_]{3,}\s*$`)
	// Unordered list markers: -, *, +
	reUL = regexp.MustCompile(`(?m)^(\s*)[-*+]\s+`)
)

// MarkdownToPlainText converts markdown to readable plain text for WeChat.
func MarkdownToPlainText(text string) string {
	result := text

	// Code blocks: strip fences, keep code content
	result = reCodeBlock.ReplaceAllStringFunc(result, func(match string) string {
		parts := reCodeBlock.FindStringSubmatch(match)
		if len(parts) > 1 {
			return strings.TrimSpace(parts[1])
		}
		return match
	})

	// Images: remove entirely
	result = reImage.ReplaceAllString(result, "")

	// Links: keep display text only
	result = reLink.ReplaceAllString(result, "$1")

	// Table separator rows: remove
	result = reTableSep.ReplaceAllString(result, "")

	// Table rows: pipe-delimited to space-delimited
	result = reTableRow.ReplaceAllStringFunc(result, func(match string) string {
		parts := reTableRow.FindStringSubmatch(match)
		if len(parts) > 1 {
			cells := strings.Split(parts[1], "|")
			for i := range cells {
				cells[i] = strings.TrimSpace(cells[i])
			}
			return strings.Join(cells, "  ")
		}
		return match
	})

	// Headers: remove # prefix
	result = reHeader.ReplaceAllString(result, "")

	// Bold
	result = reBold.ReplaceAllStringFunc(result, func(match string) string {
		parts := reBold.FindStringSubmatch(match)
		if parts[1] != "" {
			return parts[1]
		}
		return parts[2]
	})

	// Strikethrough
	result = reStrike.ReplaceAllString(result, "$1")

	// Blockquote
	result = reBlockquote.ReplaceAllString(result, "")

	// Horizontal rule -> empty line
	result = reHR.ReplaceAllString(result, "")

	// Unordered list: replace markers with "• "
	result = reUL.ReplaceAllString(result, "${1}• ")

	// Inline code: strip backticks (do after code blocks)
	result = reInlineCode.ReplaceAllString(result, "$1")

	// Clean up excessive blank lines
	result = regexp.MustCompile(`\n{3,}`).ReplaceAllString(result, "\n\n")

	return strings.TrimSpace(result)
}

```

[⬆ 回到目录](#toc)

## messaging/media.go

```go
package messaging

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/fastclaw-ai/weclaw/ilink"
)

// reMarkdownImage matches markdown image syntax: ![alt](url)
var reMarkdownImage = regexp.MustCompile(`!\[[^\]]*\]\(([^)]+)\)`)

// ExtractImageURLs extracts image URLs from markdown text.
func ExtractImageURLs(text string) []string {
	matches := reMarkdownImage.FindAllStringSubmatch(text, -1)
	var urls []string
	for _, m := range matches {
		url := strings.TrimSpace(m[1])
		if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
			urls = append(urls, url)
		}
	}
	return urls
}

// SendMediaFromURL downloads a file from a URL and sends it as a media message.
func SendMediaFromURL(ctx context.Context, client *ilink.Client, toUserID, mediaURL, contextToken string) error {
	data, contentType, err := downloadFile(ctx, mediaURL)
	if err != nil {
		return fmt.Errorf("download %s: %w", mediaURL, err)
	}

	return sendMediaData(ctx, client, toUserID, filenameFromURL(mediaURL), mediaURL, data, contentType, contextToken)
}

// SendMediaFromPath reads a local file and sends it as a media message.
func SendMediaFromPath(ctx context.Context, client *ilink.Client, toUserID, path, contextToken string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}

	return sendMediaData(ctx, client, toUserID, filepath.Base(path), path, data, inferContentType(path), contextToken)
}

func sendMediaData(ctx context.Context, client *ilink.Client, toUserID, fileName, source string, data []byte, contentType, contextToken string) error {
	if fileName == "" {
		fileName = "file"
	}

	cdnMediaType, itemType := classifyMedia(contentType, source)

	log.Printf("[media] uploading %s (%s, %d bytes) for %s", source, contentType, len(data), toUserID)

	uploaded, err := UploadFileToCDN(ctx, client, data, toUserID, cdnMediaType)
	if err != nil {
		return fmt.Errorf("upload to CDN: %w", err)
	}

	media := &ilink.MediaInfo{
		EncryptQueryParam: uploaded.DownloadParam,
		AESKey:            AESKeyToBase64(uploaded.AESKeyHex),
		EncryptType:       1,
	}

	var item ilink.MessageItem
	switch itemType {
	case ilink.ItemTypeImage:
		item = ilink.MessageItem{
			Type: ilink.ItemTypeImage,
			ImageItem: &ilink.ImageItem{
				Media:   media,
				MidSize: uploaded.CipherSize,
			},
		}
	case ilink.ItemTypeVideo:
		item = ilink.MessageItem{
			Type: ilink.ItemTypeVideo,
			VideoItem: &ilink.VideoItem{
				Media:     media,
				VideoSize: uploaded.CipherSize,
			},
		}
	default:
		item = ilink.MessageItem{
			Type: ilink.ItemTypeFile,
			FileItem: &ilink.FileItem{
				Media:    media,
				FileName: fileName,
				Len:      fmt.Sprintf("%d", uploaded.FileSize),
			},
		}
	}

	req := &ilink.SendMessageRequest{
		Msg: ilink.SendMsg{
			FromUserID:   client.BotID(),
			ToUserID:     toUserID,
			ClientID:     NewClientID(),
			MessageType:  ilink.MessageTypeBot,
			MessageState: ilink.MessageStateFinish,
			ItemList:     []ilink.MessageItem{item},
			ContextToken: contextToken,
		},
		BaseInfo: ilink.BaseInfo{},
	}

	resp, err := client.SendMessage(ctx, req)
	if err != nil {
		return fmt.Errorf("send media message: %w", err)
	}
	if resp.Ret != 0 {
		return fmt.Errorf("send media failed: ret=%d errmsg=%s", resp.Ret, resp.ErrMsg)
	}

	log.Printf("[media] sent %s to %s from %s", contentType, toUserID, source)
	return nil
}

func downloadFile(ctx context.Context, url string) ([]byte, string, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		contentType = inferContentType(url)
	}

	return data, contentType, nil
}

func classifyMedia(contentType, url string) (cdnMediaType int, itemType int) {
	ct := strings.ToLower(contentType)

	if strings.HasPrefix(ct, "image/") || isImageExt(url) {
		return ilink.CDNMediaTypeImage, ilink.ItemTypeImage
	}
	if strings.HasPrefix(ct, "video/") || isVideoExt(url) {
		return ilink.CDNMediaTypeVideo, ilink.ItemTypeVideo
	}
	return ilink.CDNMediaTypeFile, ilink.ItemTypeFile
}

func isImageExt(url string) bool {
	ext := strings.ToLower(filepath.Ext(stripQuery(url)))
	switch ext {
	case ".png", ".jpg", ".jpeg", ".gif", ".webp", ".bmp":
		return true
	}
	return false
}

func isVideoExt(url string) bool {
	ext := strings.ToLower(filepath.Ext(stripQuery(url)))
	switch ext {
	case ".mp4", ".mov", ".webm", ".mkv", ".avi":
		return true
	}
	return false
}

func inferContentType(url string) string {
	ext := filepath.Ext(stripQuery(url))
	if ct := mime.TypeByExtension(ext); ct != "" {
		return ct
	}
	return "application/octet-stream"
}

func filenameFromURL(rawURL string) string {
	u := stripQuery(rawURL)
	name := filepath.Base(u)
	if name == "" || name == "." || name == "/" {
		return "file"
	}
	return name
}

func stripQuery(rawURL string) string {
	if i := strings.IndexByte(rawURL, '?'); i >= 0 {
		return rawURL[:i]
	}
	return rawURL
}

```

[⬆ 回到目录](#toc)

## messaging/media_test.go

```go
package messaging

import "testing"

func TestExtractImageURLs(t *testing.T) {
	text := "check ![img](https://example.com/a.png) and ![](https://example.com/b.jpg)"
	urls := ExtractImageURLs(text)
	if len(urls) != 2 {
		t.Fatalf("expected 2 urls, got %d", len(urls))
	}
	if urls[0] != "https://example.com/a.png" {
		t.Errorf("urls[0] = %q", urls[0])
	}
	if urls[1] != "https://example.com/b.jpg" {
		t.Errorf("urls[1] = %q", urls[1])
	}
}

func TestExtractImageURLs_NoImages(t *testing.T) {
	urls := ExtractImageURLs("just plain text")
	if len(urls) != 0 {
		t.Errorf("expected 0 urls, got %d", len(urls))
	}
}

func TestExtractImageURLs_RelativeURL(t *testing.T) {
	text := "![img](./local.png)"
	urls := ExtractImageURLs(text)
	if len(urls) != 0 {
		t.Errorf("expected 0 urls for relative path, got %d", len(urls))
	}
}

func TestFilenameFromURL(t *testing.T) {
	tests := []struct {
		url  string
		want string
	}{
		{"https://example.com/photo.png", "photo.png"},
		{"https://example.com/path/to/report.pdf", "report.pdf"},
		{"https://example.com/file", "file"},
	}
	for _, tt := range tests {
		got := filenameFromURL(tt.url)
		if got != tt.want {
			t.Errorf("filenameFromURL(%q) = %q, want %q", tt.url, got, tt.want)
		}
	}
}

func TestFilenameFromURL_WithQuery(t *testing.T) {
	got := filenameFromURL("https://example.com/photo.png?token=abc")
	if got != "photo.png" {
		t.Errorf("got %q, want %q", got, "photo.png")
	}
}

func TestStripQuery(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"https://example.com/a?b=c", "https://example.com/a"},
		{"https://example.com/a", "https://example.com/a"},
		{"https://example.com/?x=1&y=2", "https://example.com/"},
	}
	for _, tt := range tests {
		got := stripQuery(tt.input)
		if got != tt.want {
			t.Errorf("stripQuery(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

```

[⬆ 回到目录](#toc)

## messaging/sender.go

```go
package messaging

import (
	"context"
	"fmt"
	"log"

	"github.com/fastclaw-ai/weclaw/ilink"
	"github.com/google/uuid"
)

// NewClientID generates a new unique client ID for message correlation.
func NewClientID() string {
	return uuid.New().String()
}

// SendTypingState sends a typing indicator to a user via the iLink sendtyping API.
// It first fetches a typing_ticket via getconfig, then sends the typing status.
func SendTypingState(ctx context.Context, client *ilink.Client, userID, contextToken string) error {
	// Get typing ticket
	configResp, err := client.GetConfig(ctx, userID, contextToken)
	if err != nil {
		return fmt.Errorf("get config for typing: %w", err)
	}
	if configResp.TypingTicket == "" {
		return fmt.Errorf("no typing_ticket returned from getconfig")
	}

	// Send typing
	if err := client.SendTyping(ctx, userID, configResp.TypingTicket, ilink.TypingStatusTyping); err != nil {
		return fmt.Errorf("send typing: %w", err)
	}

	log.Printf("[sender] sent typing indicator to %s", userID)
	return nil
}

// SendTextReply sends a text reply to a user through the iLink API.
// If clientID is empty, a new one is generated.
func SendTextReply(ctx context.Context, client *ilink.Client, toUserID, text, contextToken, clientID string) error {
	if clientID == "" {
		clientID = NewClientID()
	}

	// Convert markdown to plain text for WeChat display
	plainText := MarkdownToPlainText(text)

	req := &ilink.SendMessageRequest{
		Msg: ilink.SendMsg{
			FromUserID:   client.BotID(),
			ToUserID:     toUserID,
			ClientID:     clientID,
			MessageType:  ilink.MessageTypeBot,
			MessageState: ilink.MessageStateFinish,
			ItemList: []ilink.MessageItem{
				{
					Type: ilink.ItemTypeText,
					TextItem: &ilink.TextItem{
						Text: plainText,
					},
				},
			},
			ContextToken: contextToken,
		},
		BaseInfo: ilink.BaseInfo{},
	}

	resp, err := client.SendMessage(ctx, req)
	if err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	if resp.Ret != 0 {
		return fmt.Errorf("send message failed: ret=%d errmsg=%s", resp.Ret, resp.ErrMsg)
	}

	log.Printf("[sender] sent reply to %s: %q", toUserID, truncate(text, 50))
	return nil
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

```

[⬆ 回到目录](#toc)

## service/com.fastclaw.weclaw.plist

```text
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.fastclaw.weclaw</string>
    <key>ProgramArguments</key>
    <array>
        <string>/usr/local/bin/weclaw</string>
        <string>start</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
    <key>KeepAlive</key>
    <true/>
    <key>StandardOutPath</key>
    <string>/tmp/weclaw.log</string>
    <key>StandardErrorPath</key>
    <string>/tmp/weclaw.log</string>
</dict>
</plist>

```

[⬆ 回到目录](#toc)

## service/weclaw.service

```text
[Unit]
Description=WeClaw - WeChat AI Agent Bridge
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/bin/weclaw start
Restart=always
RestartSec=5
StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target

```

[⬆ 回到目录](#toc)

---
### 📊 最终统计汇总
- **文件总数:** 50
- **代码总行数:** 18264
- **物理总大小:** 488.15 KB
