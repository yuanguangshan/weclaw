# `hub`、`ilink`、`web` 文件夹详细分析

---

## 📁 1. `hub` - Agent 共享文件中心

**作用：** 实现多 Agent 之间的**文件共享与上下文传递**机制

### 文件清单

| 文件 | 行数 | 作用 |
|------|------|------|
| `hub.go` | ~350 | **核心实现** |
| `hub_test.go` | ~300 | **单元测试** |

### 核心功能

#### **A. 文件管理**

| 方法 | 作用 |
|------|------|
| `Save()` | 保存文件（自动添加 YAML frontmatter） |
| `SaveRaw()` | 保存原始内容（无 frontmatter） |
| `ReadFile()` | 读取单个文件 |
| `ReadAll()` | 读取所有文件（按时间排序） |
| `ReadSpecific()` | 读取指定文件列表 |
| `List()` / `ListWithInfo()` | 列出文件 |
| `Clear()` | 清空所有文件 |
| `Exists()` | 检查文件是否存在 |
| `FindByPartialName()` | 模糊匹配文件名 |

#### **B. 设计特点**

```go
// 1. 自动添加 frontmatter（元数据）
---
agent: claude-acp
timestamp: 2026-04-08T10:30:00Z
---

实际内容...

// 2. 文件名冲突自动重命名
test.md → test_20260408-103005.md

// 3. 大小限制（1MB）
const MaxHubFileSize = 1 * 1024 * 1024

// 4. 文件名安全处理
sanitizeFilename() → 防止路径穿越、Windows 保留名
```

#### **C. 并发安全**

```go
mu sync.RWMutex  // 读写锁
- Save()     → Lock()   (写)
- ReadAll()  → RLock()  (读)
```

#### **D. 用途**

```
Agent A (Claude)  ──Save()──→  hub/shared/analysis.md
Agent B (GPT-4)   ──ReadAll()──→  获取上下文继续处理
```

**典型场景：**
- 代码审查：Claude 生成报告 → GPT-4 基于报告优化
- 多阶段处理：每个 Agent 输出中间结果
- 上下文持久化：服务重启后恢复状态

---

## 📁 2. `ilink` - 微信 iLink API 客户端

**作用：** 封装微信 iLink 协议的 HTTP API，实现消息收发

### 文件清单

| 文件 | 行数 | 作用 |
|------|------|------|
| `types.go` | ~200 | **数据类型定义** |
| `client.go` | ~180 | **HTTP 客户端实现** |
| `auth.go` | ~130 | **登录认证** |
| `monitor.go` | ~150 | **消息监听循环** |

### 2️⃣ `types.go` - 数据类型定义

**作用：** 定义 iLink 协议的所有请求/响应结构

#### **核心类型**

```go
// 登录凭证
type Credentials struct {
    BotToken    string  // Bot 认证 Token
    ILinkBotID  string  // Bot ID
    BaseURL     string  // API 地址
    ILinkUserID string  // 用户 ID
}

// 微信消息
type WeixinMessage struct {
    FromUserID   string        // 发送者
    MessageType  int           // 1=用户消息, 2=Bot 消息
    MessageState int           // 0=新消息, 1=生成中, 2=完成
    ItemList     []MessageItem // 内容项（文本/图片/语音/视频/文件）
    ContextToken string        // 上下文令牌
}

// 消息项（多模态）
type MessageItem struct {
    Type      int        // 1=文本, 2=图片, 3=语音, 4=文件, 5=视频
    TextItem  *TextItem
    ImageItem *ImageItem
    VoiceItem *VoiceItem
    FileItem  *FileItem
    VideoItem *VideoItem
}
```

#### **常量定义**

```go
// 消息类型
MessageTypeUser = 1  // 用户发送
MessageTypeBot  = 2  // Bot 回复

// 消息状态
MessageStateNew        = 0  // 新消息
MessageStateGenerating = 1  // 生成中
MessageStateFinish     = 2  // 完成

// 内容类型
ItemTypeText  = 1
ItemTypeImage = 2
ItemTypeVoice = 3
ItemTypeFile  = 4
ItemTypeVideo = 5
```

---

### 3️⃣ `client.go` - HTTP 客户端

**作用：** 封装 iLink HTTP API 调用

#### **核心方法**

| 方法 | 作用 |
|------|------|
| `GetUpdates()` | **长轮询获取新消息**（35 秒超时） |
| `SendMessage()` | 发送消息 |
| `GetConfig()` | 获取 Bot 配置（含 typing_ticket） |
| `SendTyping()` | 发送"正在输入"状态 |
| `GetUploadURL()` | 获取 CDN 上传预签名 URL |

#### **认证机制**

```go
func (c *Client) setHeaders(req *http.Request) {
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("AuthorizationType", "ilink_bot_token")
    req.Header.Set("Authorization", "Bearer "+c.botToken)
    req.Header.Set("X-WECHAT-UIN", c.wechatUIN)  // 随机生成
}
```

#### **长轮询机制**

```go
func (c *Client) GetUpdates(ctx, buf) {
    // 长轮询：服务器挂起 35 秒直到有新消息
    // buf 用于增量同步，避免重复接收
}
```

---

### 4️⃣ `auth.go` - 登录认证

**作用：** 实现二维码登录流程

#### **登录流程**

```
┌─────────────────────────────────────────┐
│  1. FetchQRCode()                       │
│     POST /ilink/bot/get_bot_qrcode      │
│     → 返回二维码 URL 和内容             │
├─────────────────────────────────────────┤
│  2. 用户扫码（终端显示二维码）            │
├─────────────────────────────────────────┤
│  3. PollQRStatus()                      │
│     GET /ilink/bot/get_qrcode_status    │
│     轮询直到状态变为 confirmed/expired   │
├─────────────────────────────────────────┤
│  4. SaveCredentials()                   │
│     保存到 ~/.weclaw/accounts/{id}.json │
└─────────────────────────────────────────┘
```

#### **状态机**

```go
statusWait      = "wait"      // 等待扫码
statusScanned   = "scaned"    // 已扫码，等待确认
statusConfirmed = "confirmed" // 登录成功
statusExpired   = "expired"   // 二维码过期
```

#### **凭证存储**

```go
// 文件路径：~/.weclaw/accounts/wx123456-im-wechat-com.json
{
  "bot_token": "xxx",
  "ilink_bot_id": "wx123456@im.wechat.com",
  "baseurl": "https://ilinkai.weixin.qq.com",
  "ilink_user_id": "user123"
}
```

---

### 5️⃣ `monitor.go` - 消息监听

**作用：** 实现**长轮询消息循环**，自动重连和错误恢复

#### **核心循环**

```go
func (m *Monitor) Run(ctx) {
    for {
        // 1. 长轮询获取消息
        resp, err := m.client.GetUpdates(ctx, m.getUpdatesBuf)
        
        // 2. 错误处理（指数退避）
        if err != nil {
            m.failures++
            backoff := m.calcBackoff()  // 3s → 6s → 12s → ... → 60s
            time.Sleep(backoff)
            continue
        }
        
        // 3. 会话过期处理
        if resp.ErrCode == -14 {
            m.getUpdatesBuf = ""  // 重置同步缓冲区
            m.saveBuf()
            continue
        }
        
        // 4. 处理消息（并发）
        for _, msg := range resp.Msgs {
            go m.handler(ctx, m.client, msg)  // 每个消息独立处理
        }
    }
}
```

#### **容错机制**

| 场景 | 处理策略 |
|------|---------|
| **网络错误** | 指数退避重试（3s → 60s） |
| **连续失败 5 次** | 日志警告，提示用户重新登录 |
| **会话过期** | 重置同步缓冲区，自动重连 |
| **Token 过期** | 日志警告，需手动 `weclaw login` |

#### **同步缓冲区持久化**

```go
// 文件：~/.weclaw/accounts/{id}.sync.json
{
  "get_updates_buf": "base64_encoded_sync_state"
}
```

**作用：** 服务重启后从断点继续，避免消息丢失或重复

---

## 📁 3. `web` - 管理后台前端

**作用：** 嵌入管理后台 HTML 页面

### 文件清单

| 文件 | 行数 | 作用 |
|------|------|------|
| `embed.go` | ~5 | **嵌入静态资源** |
| `admin.html` | ~1326 | **管理界面**（未读取） |

### 核心实现

```go
package web

import _ "embed"

//go:embed admin.html
var AdminHTML []byte
```

**技术：** Go 1.16+ 的 `embed` 指令，将 HTML 编译进二进制

**优点：**
- 单二进制部署，无需额外文件
- 启动速度快
- 版本一致性保证

**使用方式：**
```go
// api/server.go
func (s *Server) handleAdminUI(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.Write(web.AdminHTML)
}
```

---

## 🏗️ 整体架构

```
┌─────────────────────────────────────────────────────────┐
│                    用户交互层                             │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  │
│  │  CLI (cmd/)  │  │  Web UI      │  │  HTTP API    │  │
│  │  weclaw send │  │  admin.html  │  │  REST API    │  │
│  └──────┬───────┘  └──────┬───────┘  └──────┬───────┘  │
└─────────┼─────────────────┼─────────────────┼──────────┘
          │                 │                 │
          ▼                 ▼                 ▼
┌─────────────────────────────────────────────────────────┐
│                    业务逻辑层                             │
│  ┌──────────────────────────────────────────────────┐  │
│  │  messaging/handler.go (消息路由 + Agent 调度)     │  │
│  └──────────────────────────────────────────────────┘  │
└───────┬──────────────────────────┬──────────────────────┘
        │                          │
        ▼                          ▼
┌───────────────┐        ┌──────────────────┐
│  Agent 层     │        │  共享存储层       │
│  ┌─────────┐  │        │  ┌────────────┐  │
│  │ ACP     │  │        │  │  Hub       │  │
│  │ CLI     │  │        │  │  (文件共享) │  │
│  │ HTTP    │  │        │  └────────────┘  │
│  └─────────┘  │        └──────────────────┘
└───────┬───────┘
        │
        ▼
┌─────────────────────────────────────────────────────────┐
│                    外部服务层                             │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  │
│  │  iLink API   │  │  AI Agents   │  │  CDN/OSS     │  │
│  │  (微信协议)   │  │  (本地/云端)  │  │  (媒体文件)   │  │
│  └──────────────┘  └──────────────┘  └──────────────┘  │
└─────────────────────────────────────────────────────────┘
```

---

## 📊 代码统计验证

| 文件夹 | 文件数 | 实际行数 | 说明 |
|--------|--------|---------|------|
| **hub** | 2 | ~650 | hub.go(350) + hub_test.go(300) |
| **ilink** | 4 | ~660 | types(200) + client(180) + auth(130) + monitor(150) |
| **web** | 2 | ~1330 | embed.go(5) + admin.html(1326) |

---

## 🎯 总结

### `hub` - Agent 协作中心

| 特性 | 说明 |
|------|------|
| **用途** | 多 Agent 文件共享与上下文传递 |
| **存储** | `~/.weclaw/hub/shared/` |
| **格式** | Markdown + YAML frontmatter |
| **安全** | 读写锁 + 文件名消毒 + 大小限制 |
| **复杂度** | ⭐⭐ |

### `ilink` - 微信协议客户端

| 特性 | 说明 |
|------|------|
| **用途** | 封装微信 iLink HTTP API |
| **核心** | 长轮询消息循环 + 自动重连 |
| **认证** | 二维码登录 + 凭证持久化 |
| **容错** | 指数退避 + 会话恢复 + 同步缓冲 |
| **复杂度** | ⭐⭐⭐⭐ |

### `web` - 管理后台前端

| 特性 | 说明 |
|------|------|
| **用途** | 提供 Web 管理界面 |
| **技术** | Go embed 嵌入 HTML |
| **部署** | 单二进制，无需静态文件服务器 |
| **复杂度** | ⭐ |

---

**三个文件夹的共同特点：**
- **职责单一**：每个文件夹专注于一个领域
- **测试覆盖**：`hub` 有完整单元测试，`ilink` 依赖外部服务
- **错误处理**：统一的错误传播和日志记录
- **并发安全**：`hub` 用读写锁，`ilink` 用 goroutine 并发处理消息
