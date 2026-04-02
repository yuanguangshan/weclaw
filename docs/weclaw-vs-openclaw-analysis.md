# WeClaw 项目架构与功能分析报告

## 一、项目概述

**WeClaw** 是一个用 **Go 语言** 编写的 **WeChat AI Agent Bridge**，将微信消息连接到多个 AI 代理（Claude、Codex、Gemini、Kimi 等）。

**灵感来源**：项目受到 `@tencent-weixin/openclaw-weixin` 的启发。

**技术栈**：
- 语言：Go 1.25.0
- 许可：MIT
- 仓库：`github.com/fastclaw-ai/weclaw`

---

## 二、整体架构对比

### 原版 openclaw-weixin 架构

```
WeChat ←→ openclaw-weixin (Node.js Plugin) ←→ OpenClaw AI Gateway
```

- **语言**：Node.js/TypeScript
- **定位**：OpenClaw 的微信通道插件
- **依赖**：需要安装 OpenClaw >= 2026.3.22
- **消息处理**：长轮询架构，分层处理（监控、处理、发送）

### WeClaw 架构

```
                                        HTTP API (18011)
                                              ↑
                                              │ 主动消息
                                              │
WeChat ←→ iLink API ←→ WeClaw (Go) ←→ 多个 AI 代理
                      │                    (ACP/CLI/HTTP)
                      │
                  ├── messaging/ (消息管道)
                  ├── agent/ (代理抽象层)
                  ├── hub/ (跨 Agent 共享上下文) ⭐
                  ├── config/ (配置与自动检测)
                  └── ilink/ (微信 API 客户端)
```

---

## 三、目录结构

```
/Users/ygs/ygs/weclaw/
|-- main.go                     # 入口点
|-- go.mod / go.sum             # Go 模块定义
|-- Makefile                    # 开发热重载
|-- Dockerfile                  # 多阶段 Docker 构建
|-- install.sh                  # 一键安装脚本
|
|-- cmd/                        # CLI 命令 (Cobra)
|   |-- root.go                 # 根命令
|   |-- start.go                # 启动/登录/守护进程
|   |-- login.go                # 微信 QR 登录
|   |-- send.go                 # 主动消息 CLI
|   |-- stop.go                 # 停止守护进程
|   |-- status.go               # 检查状态
|   |-- restart.go              # 重启
|   |-- update.go               # 自更新 + 版本
|
|-- agent/                      # AI Agent 抽象层
|   |-- agent.go                # Agent 接口 + 类型
|   |-- acp_agent.go            # ACP 协议 (JSON-RPC stdio)
|   |-- cli_agent.go            # CLI 代理 (每消息子进程)
|   |-- http_agent.go           # HTTP/OpenAI 兼容
|
|-- ilink/                      # 微信 iLink API 客户端
|   |-- types.go                # API 类型/常量
|   |-- client.go               # HTTP 客户端
|   |-- auth.go                 # QR 登录 + 凭据存储
|   |-- monitor.go              # 长轮询消息监控
|
|-- messaging/                  # 消息处理管道
|   |-- handler.go              # 核心路由/命令处理
|   |-- sender.go               # 文本回复 + 输入指示器
|   |-- media.go                # 媒体发送 (图片/视频/文件)
|   |-- cdn.go                  # 微信 CDN 上传/下载 (AES-128-ECB)
|   |-- markdown.go             # Markdown 转纯文本
|   |-- attachment.go           # 本地附件检测
|   |-- linkhoard.go            # URL 保存 (Jina Reader)
|
|-- hub/                        # 跨代理共享上下文 ⭐
|   |-- hub.go                  # Hub 文件存储 (保存/读取/列表/清除)
|
|-- config/                     # 配置管理
|   |-- config.go               # 配置加载/保存 + 类型
|   |-- detect.go               # 自动检测已安装 AI 代理
|
|-- api/                        # HTTP API 服务器
|   |-- server.go               # REST API (主动消息)
|
|-- service/                    # 系统服务定义
|   |-- com.fastclaw.weclaw.plist  # macOS launchd
|   |-- weclaw.service             # Linux systemd
|
|-- docs/                       # 文档
```

---

## 四、核心功能对比表

| 功能特性 | 原版 openclaw-weixin | WeClaw |
|---------|---------------------|------------------|
| **实现语言** | Node.js/TypeScript | Go 1.25 |
| **依赖要求** | 需要 OpenClaw 网关 | 独立运行，无外部依赖 |
| **支持的 AI 代理** | OpenClaw 网关路由 | Claude, Codex, Gemini, Kimi, Cursor, OpenCode, OpenClaw 等 |
| **代理模式** | 单一插件模式 | ACP、CLI、HTTP 三种模式 ⭐ |
| **账户支持** | 单账户 | 多账户并行支持 ⭐ |
| **消息类型** | 文本（流式） | 文本、图片、视频、文件、语音 ⭐ |
| **主动消息** | ❌ | ✅ CLI + HTTP API ⭐ |
| **代理路由** | ❌ | ✅ `/agentname`、`@agent1 @agent2`、别名 ⭐ |
| **跨 Agent 协作** | ❌ | ✅ Agent Hub 系统 ⭐⭐⭐ |
| **会话管理** | 基础 | ✅ `/new`、`/cwd` 工作目录切换 ⭐ |
| **URL 自动保存** | ❌ | ✅ Linkhoard + Jina Reader ⭐⭐ |
| **进度通知** | ❌ | ✅ 实时工具执行状态 ⭐ |
| **部署方式** | npm 安装 | 二进制/Docker/系统服务 ⭐ |
| **自动更新** | ❌ | ✅ `weclaw update` ⭐ |

---

## 五、新增功能详解

### 1. 多代理系统与路由 ⭐⭐

```
原版: WeChat → openclaw-weixin → OpenClaw Gateway → 单一路由

WeClaw: WeChat → WeClaw → 多个 AI 代理（同时支持）
         ├── /claude → Claude Agent
         ├── /codex → Codex Agent
         ├── @cc @cx → 广播到多个代理
         └── 默认代理
```

**核心文件**：
- [agent/acp_agent.go](../agent/acp_agent.go) - ACP 协议（1343 行，支持 JSON-RPC 2.0）
- [agent/cli_agent.go](../agent/cli_agent.go) - CLI 模式（支持会话恢复）
- [agent/http_agent.go](../agent/http_agent.go) - HTTP/OpenAI 兼容模式

**Agent 接口定义** ([agent/agent.go](../agent/agent.go))：
```go
type Agent interface {
    Chat(ctx, conversationID, message) (string, error)
    ChatWithMedia(ctx, conversationID, message, media) (string, error)
    ResetSession(ctx, conversationID) (string, error)
    Info() AgentInfo
    SetCwd(cwd string)
    SetProgressCallback(callback ProgressCallback)
}
```

**支持的路由方式**：
- `/claude message` - 发送到指定代理
- `/cc message` - 使用别名发送
- `@agent1 @agent2 message` - 广播到多个代理
- `/claude` - 切换默认代理

### 2. Agent Hub - 跨 Agent 上下文共享 ⭐⭐⭐

这是**最核心的创新功能**。

```
原版: 各 Agent 会话完全隔离

WeClaw: Agent Hub 共享上下文层
         ~/.weclaw/hub/shared/    # 共享上下文文件
         ~/.weclaw/hub/templates/  # 提示词模板
```

**新增命令** ([messaging/handler.go](../messaging/handler.go))：

| 命令 | 描述 | 示例 |
|------|------|------|
| `/hub` | 读取所有共享文件并注入上下文 | `/hub 基于以上分析，给出反驳` |
| `/hub {filename}` | 读取特定文件 | `/hub round1_claude.md 基于此反驳` |
| `/save {filename} {msg}` | 发送消息并保存回复 | `/save round1.md 分析AI未来` |
| `/hub ls` | 列出共享文件 | `/hub ls` |
| `/hub clear` | 清空共享文件 | `/hub clear` |
| `/hub pipe {target}` | 自动链式调用 | `/hub pipe gemini` |

**实现文件**：[hub/hub.go](../hub/hub.go)

**文件格式**：
```markdown
---
agent: claude
timestamp: 2026-04-02T01:30:00+08:00
session: debate-001
round: 1
---

# Claude's Analysis

[content here]
```

**使用场景**：

1. **多代理辩论**：
```
1. /save round1_claude.md 从哲学角度分析AI是否会替代人类
2. @gemini /hub round1_claude.md 从技术可行性角度反驳
3. /save round2_gemini.md @gemini 从技术可行性角度反驳
4. @claude /hub round2_gemini.md 作为哲学派回应
5. /hub 综合两方观点，给出最终结论
```

2. **链式协作**：
```
1. /save draft.md 写一个关于量子计算的技术博客大纲
2. @gemini /hub draft.md 基于大纲扩写完整文章
3. /save article.md @gemini 基于大纲扩写完整文章
4. @claude /hub article.md 审查文章质量并优化
```

### 3. Linkhoard - URL 自动保存与分析 ⭐⭐

在 [messaging/linkhoard.go](../messaging/linkhoard.go) 中实现：

```go
功能：
1. 自动拦截消息中的 URL
2. 抓取 HTML 元数据
3. 对微信文章使用 Jina Reader 提取内容
4. 保存为带 YAML frontmatter 的 Markdown
```

**存储位置**：通过 `WECLAW_SAVE_DIR` 或配置指定

### 4. 主动消息系统 ⭐⭐

**HTTP API** ([api/server.go](../api/server.go)) 运行在 `127.0.0.1:18011`：

```bash
# 主动发送消息到微信用户
POST /api/send
{
  "to": "user_id@im.wechat",
  "text": "Hello from weclaw",
  "media_url": "https://example.com/image.png"
}
```

同时支持 **CLI 命令** ([cmd/send.go](../cmd/send.go))：
```bash
# 发送文本
weclaw send --to "user_id" --text "message"

# 发送图片
weclaw send --to "user_id" --media "url"

# 发送文本 + 媒体
weclaw send --to "user_id" --text "Check this" --media "url"
```

### 5. 代理自动检测系统 ⭐

在 [config/detect.go](../config/detect.go) 中实现：

```go
// 自动扫描已安装的 AI 代理：
- claude, codex, cursor, kimi, gemini
- opencode, openclaw, pi, copilot, droid
- iflow, kiro, qwen

// 优先选择 ACP 模式而非 CLI 模式
// 支持 login shell 回退解析（处理 nvm/mise 管理的二进制文件）
```

### 6. 媒体处理管道 ⭐⭐

**核心文件**：
- [messaging/media.go](../messaging/media.go) - 媒体发送
- [messaging/cdn.go](../messaging/cdn.go) - CDN 上传/下载（AES-128-ECB 加密）
- [messaging/attachment.go](../messaging/attachment.go) - 附件检测

```go
功能：
1. 从 CDN 下载并解密传入的图片/文件/视频
2. 转发给 AI 代理
3. 从代理回复中提取图片 URL (![](url))
4. 上传到微信 CDN（AES-128-ECB 加密）
5. 作为图片消息发送
```

**支持的媒体类型**：
- 图片：png, jpg, gif, webp
- 视频：mp4, mov
- 文件：pdf, doc, zip 等
- 语音：自动使用微信语音转文本

### 7. 进度通知系统 ⭐

AI 代理可以发送实时进度更新给微信用户：

```go
type ProgressEvent struct {
    Type    string  // "thinking", "tool", "processing"
    Message string
}

// 速率限制：每 3 秒一次
```

示例通知：
- "正在执行工具..."
- "正在思考..."
- "正在处理..."

### 8. 多账户支持 ⭐

在 [ilink/](../ilink/) 中实现：

```go
// 多个微信账号可以并行监控
~/.weclaw/accounts/{id}.json
// 每个账号独立的凭据和同步缓冲区
```

### 9. 系统服务集成 ⭐

**macOS (launchd)**：[service/com.fastclaw.weclaw.plist](../service/com.fastclaw.weclaw.plist)

```bash
cp service/com.fastclaw.weclaw.plist ~/Library/LaunchAgents/
launchctl load ~/Library/LaunchAgents/com.fastclaw.weclaw.plist
```

**Linux (systemd)**：[service/weclaw.service](../service/weclaw.service)

```bash
sudo cp service/weclaw.service /etc/systemd/system/
sudo systemctl enable --now weclaw
```

### 10. 自更新功能 ⭐

在 [cmd/update.go](../cmd/update.go) 中实现：

```bash
weclaw update   # 从 GitHub 发布更新（自动重启）
weclaw version  # 检查当前版本
```

### 11. Markdown 转纯文本 ⭐

在 [messaging/markdown.go](../messaging/markdown.go) 中实现：

```go
// 自动转换代理回复的 Markdown 为微信友好的纯文本
// 去除代码围栏、链接标记、表格、标题、粗体/斜体等
```

### 12. 工作目录管理 ⭐

```bash
/cwd /path/to/project   # 切换 Agent 工作目录
/new                    # 重置会话
/info                   # 显示当前 Agent 信息
```

---

## 六、技术架构亮点

### 1. 独立性

| 方面 | 原版 | WeClaw |
|------|------|--------|
| 依赖 | 需要 OpenClaw Gateway | 完全独立 |
| 运行时 | Node.js | 单一二进制 |
| 配置 | Gateway 配置 | `~/.weclaw/config.json` |

### 2. 性能优势

- **Go 协程**：高并发消息处理
- **长轮询优化**：指数退避（3-60 秒），自动恢复
- **ACP 进程复用**：避免每次消息创建新进程
- **自动重连**：最多 5 次连续失败后重启

### 3. 部署灵活性

| 方式 | 原版 | WeClaw |
|------|------|--------|
| 二进制 | ❌ | ✅ 6 平台支持 |
| Docker | ❌ | ✅ 多阶段构建 |
| 系统服务 | ❌ | ✅ launchd/systemd |
| 一键安装 | ✅ npx | ✅ curl 脚本 |

### 4. ACP 协议支持

**完整支持两种 ACP 协议**：
- **传统 ACP**：基于会话（session/new + session/prompt）
- **Codex 应用服务器**：基于线程（turn/start + turn/completed）

**特性**：
- JSON-RPC 2.0 over NDJSON stdio
- 流式响应（session/update 通知）
- 自动权限授予
- 进度回调

---

## 七、消息流处理

```
1. 微信登录
   └──> QR 码显示 → 扫码确认 → 凭据保存 ~/.weclaw/accounts/{id}.json

2. 消息接收
   └──> ilink.Monitor 长轮询 getupdates API
       └──> 消息分派到 messaging.Handler.HandleMessage()

3. 消息路由 (handler.go)
   ├──> 内置命令: /help, /info, /new, /cwd, /save, /hub 等
   ├──> 代理路由: /agentname, @agent1 @agent2, 别名
   └──> 默认代理
       └──> URL 自动保存到 Linkhoard

4. 代理通信
   ├──> ACP 代理: 长期子进程，JSON-RPC stdio
   ├──> CLI 代理: 每消息新进程，支持 --resume
   └──> HTTP 代理: OpenAI 兼容 API

5. 响应发送
   ├──> 图片 URL 自动提取、下载、上传 CDN
   ├──> Markdown 转纯文本
   └──> 通过 CDN 发送（AES-128-ECB 加密）
```

---

## 八、配置示例

**配置文件**：`~/.weclaw/config.json`

```json
{
  "default_agent": "claude",
  "api_addr": "127.0.0.1:18011",
  "save_dir": "~/Documents/weclaw-saves",
  "agents": {
    "claude": {
      "type": "acp",
      "command": "/usr/local/bin/claude-agent-acp",
      "aliases": ["cc", "ai", "c"],
      "env": {
        "ANTHROPIC_API_KEY": "sk-ant-xxx"
      },
      "model": "sonnet"
    },
    "codex": {
      "type": "acp",
      "command": "/usr/local/bin/codex-acp",
      "aliases": ["cx"],
      "cwd": "/home/user/my-project",
      "env": {
        "OPENAI_API_KEY": "sk-xxx"
      }
    },
    "openclaw": {
      "type": "http",
      "endpoint": "https://api.example.com/v1/chat/completions",
      "api_key": "sk-xxx",
      "model": "openclaw:main",
      "headers": {
        "X-Custom-Header": "value"
      }
    }
  }
}
```

**环境变量覆盖**：
- `WECLAW_DEFAULT_AGENT` - 覆盖默认代理
- `WECLAW_API_ADDR` - API 监听地址
- `WECLAW_SAVE_DIR` - 保存目录

---

## 九、测试覆盖

| 文件 | 测试内容 |
|------|---------|
| [agent/env_test.go](../agent/env_test.go) | 环境合并 |
| [hub/hub_test.go](../hub/hub_test.go) | Hub 保存/读取/列表/清除 |
| [config/config_test.go](../config/config_test.go) | 配置 JSON 编组 |
| [config/detect_test.go](../config/detect_test.go) | 登录 shell 检测 |
| [messaging/handler_test.go](../messaging/handler_test.go) | 命令解析、别名解析 |
| [messaging/media_test.go](../messaging/media_test.go) | 图片 URL 提取 |
| [messaging/attachment_test.go](../messaging/attachment_test.go) | 附件路径提取 |

---

## 十、别名列表

| 别名 | 代理 |
|------|------|
| `/cc` | claude |
| `/cx` | codex |
| `/cs` | cursor |
| `/km` | kimi |
| `/gm` | gemini |
| `/ocd` | opencode |
| `/oc` | openclaw |

**自定义别名**：在配置文件中为每个代理添加 `aliases` 数组。

---

## 十一、总结

### 核心创新点

1. **Agent Hub 系统**：最大创新，实现跨 Agent 协作能力
2. **多代理路由**：支持同时连接多个 AI 服务
3. **主动消息**：支持服务器推送消息到微信
4. **完整媒体处理**：双向图片/视频/文件传输
5. **独立部署**：不依赖 OpenClaw Gateway

### 与原版的主要区别

```
原版定位: 微信 → OpenClaw 的通道插件
WeClaw定位: 微信 → 多 AI 代理的通用桥梁 + 协作平台
```

### 代码规模

- **总代码量**：约 6000+ 行 Go 代码
- **核心模块**：8 个主要包
- **测试文件**：7 个
- **支持平台**：6 个（darwin/linux × amd64/arm64 + windows）

### 项目特色

1. **零依赖**：仅依赖 6 个外部 Go 包
2. **高性能**：Go 协程 + 长轮询优化
3. **易部署**：单二进制 + Docker + 系统服务
4. **强扩展**：插件化 Agent 接口
5. **生产级**：完整测试 + CI/CD + 自动更新

---

## 附录：依赖项

| 依赖项 | 版本 | 用途 |
|---|---|---|
| `github.com/spf13/cobra` | v1.10.2 | CLI 命令框架 |
| `github.com/mdp/qrterminal/v3` | v3.2.1 | 终端二维码生成 |
| `github.com/google/uuid` | v1.6.0 | 消息关联 ID |
| `golang.org/x/net` | v0.52.0 | HTML 解析 |
| `golang.org/x/term` | v0.41.0 | 终端处理 |
| `rsc.io/qr` | v0.2.0 | 二维码编码 |
