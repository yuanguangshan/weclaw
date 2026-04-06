# WeClaw 项目深度分析

## 一、项目定位

**WeClaw** 是一个**微信 AI 代理桥接网关**，通过腾讯 iLink API 将微信消息连接到多种 AI 代理（Claude、Codex、Gemini、Kimi 等），实现双向通信、媒体文件传输、多代理协作等功能。

---

## 二、核心架构特色

### 2.1 分层架构设计

```
┌─────────────────────────────────────────────────────────────────┐
│                        CLI / Web UI 层                          │
│  (Cobra 命令 + 嵌入式 HTML 管理面板)                              │
└─────────────────────────────────────────────────────────────────┘
                                 ↓
┌─────────────────────────────────────────────────────────────────┐
│                      API 服务层 (api/)                          │
│  HTTP API (127.0.0.1:18011) + 管理后台端点                       │
└─────────────────────────────────────────────────────────────────┘
                                 ↓
┌─────────────────────────────────────────────────────────────────┐
│                    消息处理引擎 (messaging/)                     │
│  Handler · Media · Sender · CDN · Todo · Timer                 │
└─────────────────────────────────────────────────────────────────┘
                                 ↓
┌─────────────────┬───────────────────────┬───────────────────────┐
│  微信 iLink 层   │    AI 代理抽象层        │     共享上下文层        │
│  (ilink/)       │    (agent/)            │     (hub/)            │
│  · 长轮询监听    │    · ACP 协议          │  · 跨代理协作           │
│  · QR 登录      │    · CLI 协议          │  · YAML frontmatter    │
│  · CDN 上传     │    · HTTP 协议         │  · Pipe 机制           │
└─────────────────┴───────────────────────┴───────────────────────┘
```

### 2.2 设计模式运用

| 设计模式 | 应用场景 | 实现位置 |
|---------|---------|---------|
| **策略模式** | Agent 接口多实现 | [agent/agent.go:111](../agent/agent.go#L111) |
| **工厂模式** | 按需创建代理实例 | [messaging/handler.go:158](../messaging/handler.go#L158) |
| **观察者模式** | 进度回调通知 | [agent/agent.go:50](../agent/agent.go#L50) |
| **守护进程模式** | 后台服务管理 | [cmd/start.go](../cmd/start.go) |
| **单例模式** | Hub 共享存储 | [hub/hub.go](../hub/hub.go) |

---

## 三、功能亮点详解

### 3.1 多代理统一抽象 ([agent/](../agent/))

**核心接口设计：**
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

**三种协议实现：**

| 协议 | 特点 | 适用场景 | 实现文件 |
|------|------|---------|---------|
| **ACP** | JSON-RPC 2.0 over stdio，长驻子进程 | Claude ACP, Codex ACP, Cursor, Kimi | [acp_agent.go](../agent/acp_agent.go) |
| **CLI** | 每消息启动新进程，流式 JSON | Claude CLI, Codex exec | [cli_agent.go](../agent/cli_agent.go) |
| **HTTP** | OpenAI 兼容 chat/completions API | 任何兼容 OpenAI 的服务 | [http_agent.go](../agent/http_agent.go) |

**ACP 协议高级特性：**
- 支持两种子协议：`legacy_acp` 和 `codex_app_server`
- 会话/线程复用机制（sessions/threads map）
- 权限自动批准（auto-allow）
- JSON-RPC 请求/响应管理
- stderr 错误捕获

### 3.2 智能消息路由 ([messaging/handler.go](../messaging/handler.go))

**命令解析能力：**
```
普通消息          → 默认代理
/agentname 消息   → 指定代理
@agent1 @agent2   → 多代理广播（并行）
内置命令          → 系统处理
```

**内置命令系统：**

| 命令 | 功能 |
|------|------|
| `/help` | 显示帮助信息 |
| `/info` | 显示系统状态 |
| `/new` / `/clear` | 重置会话 |
| `/cwd [path]` | 设置工作目录 |
| `/save [@agent] file msg` | 保存回复到 Hub |
| `/hub [@agent] file msg` | 从 Hub 读取上下文 |
| `/podcast` | 生成播客脚本 |
| `/debate` | 多代理辩论 |
| `/todo` | 待办事项管理 |
| `/timer` | 倒计时器 |
| `/sh` / `/$` | Shell 模式 |

### 3.3 跨代理协作 Hub ([hub/hub.go](../hub/hub.go))

**核心机制：**
```yaml
---
agent: claude
timestamp: 2026-04-06T12:00:00Z
---

# 用户保存的内容...
```

**Pipe 机制：**
```bash
# Claude 分析后保存到 Hub
/cc 分析这个代码...

# 传给 Codex 继续处理
/hub pipe @cx @latest 用 TypeScript 重写
```

**特性：**
- YAML frontmatter 标记来源和时间
- 文件名冲突自动重命名（时间戳后缀）
- 按编号引用（`@1`、`@-1` 最新）
- 按名称部分匹配（`gemini` 匹配 `pipe_20260402_gemini.md`）

### 3.4 媒体处理系统

**支持流程：**
```
微信媒体消息
    ↓
iLink CDN 下载 (AES-128-ECB 解密)
    ↓
本地保存 (save_dir)
    ↓
发送给 AI 代理 (ChatWithMedia)
    ↓
提取回复中的图片/附件
    ↓
上传到微信 CDN (AES-128-ECB 加密)
    ↓
发送给用户
```

**类型支持：**
- 图片（自动提取 Markdown 中的 URL）
- 视频
- 文件
- 语音（利用微信内置转文字）

### 3.5 待办事项系统 ([messaging/todo.go](../messaging/todo.go))

**数据结构：**
```go
type TodoItem struct {
    ID        int    // 自增 ID
    UserID    string // 用户微信 ID
    Title     string // 任务标题
    DueTime   int64  // 截止时间戳（0=无截止）
    Status    int    // 0=待办, 1=完成
    CreatedAt int64  // 创建时间
    Reminded  bool   // 是否已提醒
}
```

**命令格式：**
```bash
/todo 内容               # 添加待办
/todo 内容 @明天 3pm     # 带时间（AI 解析）
/todo                    # 列出待办
/todo done #1            # 完成
/todo del #1             # 删除
/todo clear              # 清空
```

**定时提醒：**
- 后台调度器每分钟检查
- 到期前 1 分钟发送微信提醒
- 支持自然语言时间解析

### 3.6 倒计时器系统

**使用场景：**
```bash
/timer 5分钟            # 5分钟倒计时
/timer 倒计时结束       # 自定义结束消息
```

**实现机制：**
- 每个用户独立的计时器
- 到期自动发送提醒
- 支持自然语言时间解析

### 3.7 自动检测系统 ([config/detect.go](../config/detect.go))

**支持 15+ 种代理：**
- Claude (ACP/CLI)
- Codex (ACP/CLI/App-Server)
- Cursor Agent
- Kimi
- Gemini
- OpenCode
- OpenClaw
- Pi
- Copilot
- Droid
- iFlow
- Kiro
- Qwen
- ...更多

**检测策略：**
1. 扫描系统 PATH
2. 检查 login shell 配置
3. 优先 ACP > CLI
4. 自动配置别名

### 3.8 Web 管理面板

**嵌入式单页应用：**
- 使用 `go:embed` 嵌入二进制
- 无需额外前端服务
- 完整的 CRUD 操作

**API 端点：**
```
GET/PUT  /api/config          # 配置管理
GET/POST/PUT/DELETE /api/agents  # 代理管理
POST     /api/agents/detect   # 自动检测
GET/DELETE /api/accounts      # 账户管理
GET      /api/login/qrcode    # QR 登录
POST     /api/service/restart # 服务重启
GET      /api/logs            # 日志查看
GET/POST/PUT/DELETE /api/todos  # 待办管理
GET      /admin               # 管理面板
```

---

## 四、数据流转详解

### 4.1 消息接收流程

```
┌────────────────────────────────────────────────────────────┐
│ 1. iLink 长轮询 (ilink/monitor.go)                          │
│    - 35s 超时长轮询                                         │
│    - get_updates_buf 持久化（断点续传）                      │
│    - 消息去重（sync.Map seenMsgs）                          │
└────────────────────────────────────────────────────────────┘
                         ↓
┌────────────────────────────────────────────────────────────┐
│ 2. 消息预过滤                                               │
│    - MessageType == User                                    │
│    - MessageState == Finish                                 │
│    - Dedup by MessageID                                     │
└────────────────────────────────────────────────────────────┘
                         ↓
┌────────────────────────────────────────────────────────────┐
│ 3. 内容提取 (handler.go:367)                                │
│    - 文本消息 extractText()                                 │
│    - 语音转文字 extractVoiceText()                          │
│    - 媒体附件 extractAllMedia()                             │
└────────────────────────────────────────────────────────────┘
                         ↓
┌────────────────────────────────────────────────────────────┐
│ 4. 特殊路径处理                                             │
│    ├─ URL 拦截 → Linkhoard 保存                             │
│    ├─ 媒体消息 → sendMediaToAgent()                         │
│    ├─ Shell 模式 → handleShellWithState()                   │
│    └─ 普通文本 → 命令路由                                   │
└────────────────────────────────────────────────────────────┘
```

### 4.2 命令路由流程

```
┌────────────────────────────────────────────────────────────┐
│ parseCommand(text)                                         │
│    - 解析 /agent 或 @agent 前缀                            │
│    - 支持多代理 @cc @cx msg                                │
│    - 别名解析（cc→claude）                                  │
└────────────────────────────────────────────────────────────┘
                         ↓
         ┌───────────────┼───────────────┐
         ↓               ↓               ↓
   内置命令       单代理指定        多代理广播
 (/hub, /todo)    (/cc msg)      (@cc @cx msg)
         ↓               ↓               ↓
   直接处理      sendToNamedAgent  broadcastToAgents
                                 (并行执行)
```

### 4.3 AI 代理调用流程

```
┌────────────────────────────────────────────────────────────┐
│ 1. 获取代理实例 (getAgent)                                  │
│    - 快速路径：已运行的代理                                 │
│    - 慢速路径：factory 按需创建                             │
└────────────────────────────────────────────────────────────┘
                         ↓
┌────────────────────────────────────────────────────────────┐
│ 2. 发送打字状态                                             │
│    - SendTypingState() 异步发送                             │
│    - 提升用户体验                                           │
└────────────────────────────────────────────────────────────┘
                         ↓
┌────────────────────────────────────────────────────────────┐
│ 3. 调用代理 chatWithAgent()                                 │
│    - Chat() 或 ChatWithMedia()                              │
│    - 传递 conversationID（会话隔离）                         │
│    - 设置进度回调                                           │
└────────────────────────────────────────────────────────────┘
                         ↓
┌────────────────────────────────────────────────────────────┐
│ 4. ACP 协议流程 (acp_agent.go)                              │
│    ┌────────────────────────────────────────────┐          │
│    │ session/new → 建立会话                     │          │
│    │ session/prompt → 发送消息                  │          │
│    │ session/update (流式) → 接收响应           │          │
│    └────────────────────────────────────────────┘          │
└────────────────────────────────────────────────────────────┘
                         ↓
┌────────────────────────────────────────────────────────────┐
│ 5. 进度通知 (ProgressCallback)                              │
│    - tool_start: 工具开始执行                               │
│    - tool_end: 工具执行结束                                 │
│    - thought: AI 思考中                                     │
│    - file_read/write: 文件操作                              │
└────────────────────────────────────────────────────────────┘
                         ↓
┌────────────────────────────────────────────────────────────┐
│ 6. 响应处理 sendReplyWithMedia()                            │
│    - Markdown → 纯文本转换                                  │
│    - 提取本地文件路径                                       │
│    - 下载网络图片                                           │
│    - 上传 CDN 并发送                                        │
└────────────────────────────────────────────────────────────┘
```

### 4.4 Hub 共享上下文流程

**保存流程：**
```
用户消息: /save analysis 这次分析很好
         ↓
handleSave() → 代理调用 → 保存结果
         ↓
hub.Save("analysis.md", content, "claude")
         ↓
写入 ~/.weclaw/hub/shared/analysis_20260406_120000.md
         ↓
YAML frontmatter:
  ---
  agent: claude
  timestamp: 2026-04-06T12:00:00Z
  ---
```

**Pipe 流程：**
```
用户: /hub pipe @cx @analysis 用 TS 重写
         ↓
handleHub() → hub.ReadFile("analysis.md")
         ↓
构建 prompt: [Hub 内容] + [用户消息]
         ↓
调用 codex.Chat(prompt)
         ↓
返回 TS 重写结果
```

### 4.5 进度通知流程

```
┌────────────────────────────────────────────────────────────┐
│ ACP 子进程 stderr                                           │
│    {"type": "tool_start", "toolName": "read_file"}          │
└────────────────────────────────────────────────────────────┘
                         ↓
┌────────────────────────────────────────────────────────────┐
│ readLoop() 解析 JSON-RPC 通知                               │
│    - session/update 事件                                    │
│    - 提取 progress 字段                                     │
└────────────────────────────────────────────────────────────┘
                         ↓
┌────────────────────────────────────────────────────────────┐
│ ProgressCallback 触发                                       │
│    - 构建 "正在使用 xxx..." 消息                            │
│    - 调用 SendTypingState()                                 │
└────────────────────────────────────────────────────────────┘
                         ↓
┌────────────────────────────────────────────────────────────┐
│ 用户微信显示                                                │
│    "Claude 正在使用 read_file..."                          │
└────────────────────────────────────────────────────────────┘
```

---

## 五、技术亮点

### 5.1 并发控制
- **消息并发处理**：每条消息独立 goroutine
- **多代理广播**：并行调用，结果按到达顺序发送
- **sync.Map**：无锁缓存（seenMsgs、contextTokens、lastReplies）

### 5.2 容错机制
- **长轮询自动重连**：指数退避（3s → 60s）
- **会话过期恢复**：自动重置 sync buf
- **代理按需加载**：启动失败不影响其他代理

### 5.3 安全措施
- **文件名清洗**：防止路径遍历（sanitizeFilename）
- **大小限制**：Hub 文件最大 1MB
- **工作目录隔离**：每个代理独立 cwd

### 5.4 跨平台支持
- **系统服务**：macOS launchd + Linux systemd
- **CI/CD**：GitHub Actions 跨平台构建
- **Docker**：Alpine 两阶段构建

---

## 六、项目总结

WeClaw 是一个**工程化水平极高**的 Go 项目，展现了：

1. **清晰的分层架构**：从 CLI 到消息处理再到 AI 代理抽象
2. **灵活的设计模式**：策略、工厂、观察者模式的恰当运用
3. **强大的功能集成**：多代理、媒体处理、Hub 协作、Todo/Timer
4. **完善的工程实践**：并发控制、容错机制、安全措施
5. **优秀的用户体验**：自动检测、进度通知、Web 管理面板

特别值得学习的创新点：
- **Hub 跨代理协作机制**：让不同 AI 共享上下文
- **Pipe 机制**：代理间无缝传递分析结果
- **进度回调**：让用户实时了解 AI 状态
- **自然语言时间解析**：todo/timer 支持 "明天 3pm" 这样的表达

---

## 七、目录结构

```
/Users/ygs/ygs/weclaw/
|-- main.go                  # 入口：调用 cmd.Execute()
|-- go.mod / go.sum          # Go 模块定义
|-- config.json              # 实际配置文件（含多代理配置）
|-- Makefile                 # make dev（热重载）
|-- Dockerfile               # 两阶段构建（golang:1.25-alpine + alpine:3.21）
|-- install.sh               # curl 一键安装脚本
|-- README.md                # 英文文档
|-- LICENSE
|
|-- cmd/                     # CLI 命令层（Cobra 框架）
|   |-- root.go              # 根命令（默认执行 start）
|   |-- start.go             # 核心启动逻辑：登录、配置、消息桥、守护进程
|   |-- login.go             # 微信 QR 码登录
|   |-- send.go              # 主动发送消息
|   |-- stop.go              # 停止后台进程
|   |-- restart.go           # 重启后台进程
|   |-- status.go            # 检查进程状态
|   |-- update.go            # 自更新（从 GitHub Release 下载）
|   |-- proc_unix.go         # Unix 进程属性设置
|   |-- proc_windows.go      # Windows 进程属性设置
|
|-- agent/                   # AI 代理抽象层
|   |-- agent.go             # Agent 接口定义、类型定义
|   |-- acp_agent.go         # ACP 协议代理（JSON-RPC over stdio）
|   |-- cli_agent.go         # CLI 代理（每消息启动新进程）
|   |-- http_agent.go        # HTTP 代理（OpenAI 兼容 API）
|   |-- env_test.go          # 测试
|
|-- config/                  # 配置管理
|   |-- config.go            # 配置数据模型、加载、保存
|   |-- detect.go            # 自动检测已安装的 AI 代理
|   |-- config_test.go
|   |-- detect_test.go
|
|-- ilink/                   # 微信 iLink API 客户端
|   |-- types.go             # API 类型定义（消息、CDN、登录等）
|   |-- client.go            # HTTP API 客户端
|   |-- auth.go              # QR 码登录、凭据管理
|   |-- monitor.go           # 长轮询消息监听器
|
|-- messaging/               # 消息处理引擎
|   |-- handler.go           # 核心消息路由与命令处理
|   |-- sender.go            # 文本回复发送
|   |-- media.go             # 媒体下载/分类/发送
|   |-- cdn.go               # CDN 上传/下载（AES-128-ECB 加密）
|   |-- markdown.go          # Markdown 转纯文本
|   |-- attachment.go        # 附件提取与安全校验
|   |-- linkhoard.go         # URL 抓取与保存（Jina Reader、微信文章）
|   |-- todo.go              # 待办事项管理
|   |-- timer.go             # 计时器管理
|   |-- handler_test.go / media_test.go / attachment_test.go / timer_test.go
|
|-- hub/                     # 代理间共享上下文存储
|   |-- hub.go               # 文件存储、YAML frontmatter、跨代理协作
|   |-- hub_test.go
|
|-- api/                     # HTTP API 服务 + Web 管理面板
|   |-- server.go            # HTTP 服务路由与消息发送 API
|   |-- admin.go             # 管理后台 API（配置、代理、账户、日志、Hub、Todo、Timer、登录、服务控制）
|
|-- web/                     # 嵌入式前端
|   |-- embed.go             # go:embed admin.html
|   |-- admin.html           # Web 管理面板（单文件 HTML）
|
|-- service/                 # 系统服务配置
|   |-- com.fastclaw.weclaw.plist   # macOS launchd
|   |-- weclaw.service             # Linux systemd
|
|-- docs/                    # 项目文档
|-- previews/                # 截图
|-- .github/workflows/       # CI/CD（ci.yml, release.yml）
|-- .ai/                     # AI 工具上下文
|-- .shell/                  # Shell 模式配置
```

---

## 八、技术栈

| 层面 | 技术 |
|------|------|
| 语言 | Go 1.25 |
| CLI 框架 | Cobra (`spf13/cobra`) |
| UUID 生成 | `google/uuid` |
| QR 码终端显示 | `mdp/qrterminal/v3` |
| HTTP 客户端 | 标准库 `net/http` |
| HTML 解析 | `golang.org/x/net/html` |
| 加密 | 标准库 `crypto/aes`（AES-128-ECB） |
| 进程管理 | `os/exec`、`syscall` |
| CI/CD | GitHub Actions（跨平台构建：darwin/linux/windows x amd64/arm64） |
| 容器化 | Docker（Alpine 镜像） |
| 系统服务 | macOS launchd + Linux systemd |
| 前端 | 嵌入式 HTML（单文件管理面板） |
| 热重载开发 | Air (`cosmtrek/air`) |

**外部 API 集成：**
- 微信 iLink API（`ilinkai.weixin.qq.com`）-- 消息收发、QR 登录、CDN 上传
- WeChat CDN（`novac2c.cdn.weixin.qq.com`）-- 媒体文件上传/下载
- Jina Reader API（`r.jina.ai`）-- URL 内容抓取
- GitHub API（`api.github.com`）-- 版本检查与更新
- 各 AI 代理的 API/CLI/ACP 接口
