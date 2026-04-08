# `messaging` 文件夹详细分析

这是项目的**核心业务逻辑层**，负责消息路由、Agent 调度、文件处理和任务管理。

---

## 📁 文件清单

| 文件 | 行数 | 作用 |
|------|------|------|
| `handler.go` | ~2716 | **消息处理器**（核心路由） |
| `sender.go` | ~85 | **消息发送器** |
| `timer.go` | ~436 | **定时器管理** |
| `todo.go` | ~380 | **待办事项管理** |
| `media.go` | ~195 | **媒体文件处理** |
| `cdn.go` | ~200 | **CDN 上传/下载** |
| `attachment.go` | ~120 | **附件路径处理** |
| `linkhoard.go` | ~300 | **网页内容抓取与保存** |
| `markdown.go` | ~90 | **Markdown 转纯文本** |
| `*_test.go` | ~400 | **单元测试** |

---

## 1️⃣ `handler.go` - 消息处理器（核心文件）

**作用：** 接收微信消息 → 解析命令 → 路由到 Agent → 返回回复

**代码占比：** 2716 行，占整个 messaging 目录的 **53.4%**，是**项目的绝对核心**

---

### 📦 核心结构

```go
type Handler struct {
    mu              sync.RWMutex
    defaultName     string              // 默认 Agent 名称
    agents          map[string]agent.Agent  // 运行中的 Agent
    agentMetas      []AgentMeta         // 所有配置的 Agent
    agentWorkDirs   map[string]string   // Agent 工作目录
    customAliases   map[string]string   // 自定义别名
    factory         AgentFactory        // Agent 工厂函数
    hub             *hub.Hub            // 共享文件中心
    contextTokens   sync.Map            // 用户上下文令牌
    saveDir         string              // 文件保存目录
    seenMsgs        sync.Map            // 消息去重缓存
    todoStore       *TodoStore          // 待办存储
    timerStore      *TimerStore         // 定时器存储
    clients         []*ilink.Client     // iLink 客户端
}
```

---

### 🔄 消息处理流程 `HandleMessage()`

```
┌─────────────────────────────────────────────────────────┐
│  1. 过滤消息                                             │
│     - 只处理用户消息 (MessageTypeUser)                   │
│     - 只处理完成状态 (MessageStateFinish)                │
│     - 消息去重 (seenMsgs)                                │
├─────────────────────────────────────────────────────────┤
│  2. 提取内容                                             │
│     - extractText() → 文本                               │
│     - extractVoiceText() → 语音转文字                     │
│     - extractAllMedia() → 图片/文件/视频                  │
├─────────────────────────────────────────────────────────┤
│  3. 媒体消息处理                                        │
│     - 有媒体 → sendMediaToAgent() → 直接转发              │
├─────────────────────────────────────────────────────────┤
│  4. URL 拦截                                            │
│     - 纯 URL → SaveLinkToLinkhoard() → 保存网页           │
│     - 微信文章 → analyzeWithNanobot() → AI 分析           │
├─────────────────────────────────────────────────────────┤
│  5. 命令解析                                             │
│     - parseCommand() → 提取 @agent /agent                │
│     - 支持多 Agent: @cc @cx hello                        │
│     - 别名解析: cc→claude, cx→codex                      │
├─────────────────────────────────────────────────────────┤
│  6. 内置命令路由                                         │
│     /info, /help, /new, /clear, /cwd                    │
│     /save, /hub, /sh, /podcast, /debate                 │
│     /todo, /timer                                       │
├─────────────────────────────────────────────────────────┤
│  7. Agent 路由                                           │
│     - 无命令 → sendToDefaultAgent()                      │
│     - 单 Agent → sendToNamedAgent()                      │
│     - 多 Agent → broadcastToAgents() (并发)               │
│     - 仅 Agent 名 → switchDefault()                      │
└─────────────────────────────────────────────────────────┘
```

---

### 🎯 内置命令详解

| 命令 | 作用 | 实现函数 |
|------|------|---------|
| `/info` | 显示当前 Agent 状态 | `buildStatus()` |
| `/help` | 显示帮助文本 | `buildHelpText()` |
| `/new` `/clear` | 重置会话 | `resetDefaultSession()` |
| `/cwd` | 切换工作目录 | `handleCwd()` |
| `/save` | 发送消息并保存回复到 Hub | `handleSave()` |
| `/hub` | 读取共享文件并注入上下文 | `handleHub()` |
| `/sh` `/$` | 执行 Shell 命令 | `handleShell()` |
| `/podcast` | 生成播客 | `handlePodcast()` |
| `/debate` | 多 Agent 辩论 | `handleDebate()` |
| `/todo` | 待办事项管理 | `handleTodo()` |
| `/timer` | 定时器管理 | `handleTimer()` |

---

### 🔗 Hub 命令高级用法

```bash
# 基础用法
/hub                    # 列出共享文件
/hub ls                 # 详细列表（带编号）
/hub cat 1              # 查看编号 1 的文件
/hub clear              # 清空所有文件

# 注入上下文
/hub 分析量子计算         # 注入所有共享文件
/hub file.md 继续分析     # 注入指定文件

# Pipe 链式协作
/hub pipe gemini 分析量子计算           # 默认 Agent → Gemini
/hub pipe claude @1 商业应用            # 引用编号 1 的文件
/hub pipe deepseek @-1 投资建议          # 引用最新文件
/hub pipe claude @gemini 继续分析        # 引用部分匹配的文件
```

**Pipe 工作流程：**
```
1. 默认 Agent 处理消息 → 回复保存到 Hub
2. 目标 Agent 读取 Hub 文件 → 继续分析
3. 最终回复也保存到 Hub
4. 返回结果（含文件编号引用信息）
```

---

### 🎭 辩论功能 `/debate`

```go
func (h *Handler) runDebate(ctx, client, msg, proAgent, conAgent, topic) {
    for round := 1; round <= 3; round++ {
        // 正方发言
        proReply = proAgent.Chat(proPrompt)
        
        // 反方基于正方观点反驳
        conReply = conAgent.Chat(conPrompt)
        
        // 发送本轮结果
        sendMsg(roundText)
    }
    
    // 生成完整 Markdown 文档
    // 包含所有轮次的观点
    sendMsg(fullDebateDoc)
}
```

**辩论提示词设计：**
- **第 1 轮**：立论（3-5 个要点，500 字）
- **第 2-3 轮**：反驳 + 强化观点（400 字）
- **最终**：生成结构化 Markdown 文档

---

### 🖥️ Shell 模式

```go
// 安全限制
- 白名单命令：ls, cat, pwd, find, grep, head, tail, cd
- 禁止特殊字符：>, <, |, &&, ||, ;, `, $()
- 路径沙箱：基于 allowedRoots（Agent 工作目录）

// 快捷别名
ll     → ls -lh
..     → cd ..
...    → cd ../..

// 状态持久化
type shellModeState struct {
    enabled bool   // 是否启用
    cwd     string // 当前目录
    baseDir string // 基础目录（沙箱）
}
```

---

## 2️⃣ `sender.go` - 消息发送器

**作用：** 封装 iLink 消息发送 API

| 函数 | 作用 |
|------|------|
| `NewClientID()` | 生成唯一消息 ID（UUID） |
| `SendTypingState()` | 发送"正在输入"状态 |
| `SendTextReply()` | 发送文本回复（自动转纯文本） |

**关键流程：**
```go
func SendTextReply(ctx, client, toUserID, text, contextToken, clientID) {
    plainText := MarkdownToPlainText(text)  // Markdown → 纯文本
    
    req := &ilink.SendMessageRequest{
        Msg: ilink.SendMsg{
            FromUserID:   client.BotID(),
            ToUserID:     toUserID,
            ClientID:     clientID,
            MessageType:  ilink.MessageTypeBot,
            MessageState: ilink.MessageStateFinish,
            ItemList: []ilink.MessageItem{
                {Type: ilink.ItemTypeText, TextItem: &ilink.TextItem{Text: plainText}},
            },
            ContextToken: contextToken,
        },
    }
    
    client.SendMessage(ctx, req)
}
```

---

## 3️⃣ `timer.go` - 定时器管理

**作用：** 实现倒计时定时器功能

### **核心结构**

```go
type TimerItem struct {
    ID        int    `json:"id"`
    UserID    string `json:"user_id"`
    Label     string `json:"label"`      // 标签（如"写报告"）
    Duration  int64  `json:"duration"`   // 时长（秒）
    EndTime   int64  `json:"end_time"`   // 到期时间戳
    Status    int    `json:"status"`     // 0=运行, 1=完成, 2=取消
    Reminded  bool   `json:"reminded"`   // 是否已提醒
}
```

### **时间解析**

```go
// 直接解析（快速路径）
"25"        → 25 分钟 = 1500 秒
"2h"        → 2 小时 = 7200 秒
"30m 休息"   → 30 分钟 + 标签"休息"
"1.5h 写报告" → 1.5 小时 + 标签"写报告"

// AI 解析（慢速路径，用于自然语言）
"半小时后提醒我开会" → Agent.Chat() → {"seconds": 1800, "label": "开会"}
```

### **后台调度器**

```go
func (h *Handler) StartTimerScheduler(ctx) {
    ticker := time.NewTicker(10 * time.Second)
    go func() {
        for {
            <-ticker.C
            h.checkTimerExpirations(ctx)  // 检查到期定时器
        }
    }()
}
```

---

## 4️⃣ `todo.go` - 待办事项管理

**作用：** 实现待办事项 CRUD + 到期提醒

### **核心结构**

```go
type TodoItem struct {
    ID        int    `json:"id"`
    UserID    string `json:"user_id"`
    Title     string `json:"title"`
    DueTime   int64  `json:"due_time"` // 0 = 无截止时间
    Status    int    `json:"status"`   // 0=待办, 1=已完成
    Reminded  bool   `json:"reminded"` // 是否已提醒
}
```

### **时间提取**

```go
// 使用 AI 解析自然语言
func (h *Handler) createTodo(ctx, userID, text) {
    prompt := `从这句话中提取时间和待办事项。只返回 JSON：
{"time": "YYYY-MM-DD HH:MM:SS", "title": "事项"}
句子：明天下午 3 点开会`
    
    reply, _ := ag.Chat(ctx, userID+"_todo", prompt)
    // 解析 JSON → {"time": "2026-04-09 15:00:00", "title": "开会"}
}
```

### **后台调度器**

```go
func (h *Handler) StartTodoScheduler(ctx) {
    ticker := time.NewTicker(1 * time.Minute)  // 每分钟检查
    go func() {
        for {
            <-ticker.C
            h.checkReminders(ctx)  // 发送到期提醒
        }
    }()
}
```

---

## 5️⃣ `media.go` - 媒体文件处理

**作用：** 处理图片/文件/视频的下载和发送

| 函数 | 作用 |
|------|------|
| `ExtractImageURLs()` | 从 Markdown 提取图片 URL |
| `SendMediaFromURL()` | 下载 URL 文件并发送 |
| `SendMediaFromPath()` | 读取本地文件并发送 |
| `downloadFile()` | HTTP 下载文件 |
| `classifyMedia()` | 分类媒体类型（图片/视频/文件） |

**发送流程：**
```
1. 下载文件（HTTP 或本地）
2. 分类媒体类型（image/video/file）
3. 上传到微信 CDN（encrypt + upload）
4. 构造消息（含 MediaInfo）
5. 发送消息
```

---

## 6️⃣ `cdn.go` - CDN 上传/下载

**作用：** 封装微信 CDN 的加密上传和解密下载

### **上传流程**

```go
func UploadFileToCDN(ctx, client, data, toUserID, mediaType) {
    // 1. 生成随机密钥
    filekey = rand(16 bytes)
    aeskey  = rand(16 bytes)
    
    // 2. 计算 MD5
    rawMD5 = md5(data)
    
    // 3. 获取上传 URL
    uploadResp = client.GetUploadURL(ctx, &GetUploadURLRequest{
        FileKey:    filekeyHex,
        MediaType:  mediaType,
        RawSize:    len(data),
        RawFileMD5: rawMD5,
        FileSize:   cipherSize,  // PKCS7 padding 后大小
        AESKey:     aeskeyHex,
    })
    
    // 4. AES-128-ECB 加密
    encrypted = encryptAESECB(data, aeskey)
    
    // 5. 上传到 CDN
    downloadParam = uploadToCDN(ctx, encrypted, cdnURL)
    
    return &UploadedFile{
        DownloadParam: downloadParam,
        AESKeyHex:     aeskeyHex,
        FileSize:      len(data),
        CipherSize:    cipherSize,
    }
}
```

### **下载流程**

```go
func DownloadFileFromCDN(ctx, encryptQueryParam, aesKeyBase64) {
    // 1. 构造下载 URL
    downloadURL = "https://novac2c.cdn.weixin.qq.com/c2c/download?encrypted_query_param=..."
    
    // 2. 下载加密数据
    encrypted = http.Get(downloadURL)
    
    // 3. AES-128-ECB 解密
    plaintext = decryptAESECB(encrypted, aesKey)
    
    return plaintext
}
```

---

## 7️⃣ `attachment.go` - 附件路径处理

**作用：** 从 Agent 回复中提取本地文件路径并发送

### **提取逻辑**

```go
func extractLocalAttachmentPaths(text string) []string {
    // 1. 逐行扫描
    for _, line := range strings.Split(text, "\n") {
        // 2. 检查绝对路径
        if filepath.IsAbs(candidate) {
            // 3. 检查扩展名
            if isSupportedAttachmentPath(candidate) {
                // 4. 检查文件存在
                if os.Stat(candidate) != nil {
                    paths = append(paths, candidate)
                }
            }
        }
    }
}
```

### **安全限制**

```go
func isAllowedAttachmentPath(path string, allowedRoots []string) bool {
    // 1. 规范化路径（解析符号链接）
    cleanPath = canonicalizePath(path)
    
    // 2. 检查是否在允许根目录下
    for _, root := range allowedRoots {
        rel, _ = filepath.Rel(root, cleanPath)
        if !strings.HasPrefix(rel, "..") {
            return true  // 在沙箱内
        }
    }
    return false
}
```

**支持的扩展名：**
```
文档: .pdf, .doc, .docx, .xls, .xlsx, .ppt, .pptx, .zip, .txt, .csv
图片: .png, .jpg, .jpeg, .gif, .webp
视频: .mp4, .mov
```

---

## 8️⃣ `linkhoard.go` - 网页内容抓取

**作用：** 抓取网页内容并保存为 Markdown 文件

### **抓取策略**

```go
func SaveLinkToLinkhoard(ctx, saveDir, rawURL) {
    if isWeChatURL(rawURL) {
        // 微信文章：直接抓取（带浏览器 Header）
        meta = FetchLinkMetadata(ctx, rawURL)
    } else {
        // 其他网站：使用 Jina Reader API
        meta = FetchViaJina(ctx, rawURL)
        if err != nil {
            // 降级：直接抓取
            meta = FetchLinkMetadata(ctx, rawURL)
        }
    }
    
    // 保存为 Markdown
    saveMarkdownFile(meta, rawURL)
}
```

### **元数据提取**

```go
type LinkMetadata struct {
    Title       string  // 标题
    Description string  // 描述
    Author      string  // 作者
    OGImage     string  // OpenGraph 图片
    Published   string  // 发布时间
    Body        string  // 正文内容
}
```

**HTML 解析：**
```go
func extractMeta(n *html.Node, meta *LinkMetadata) {
    // <meta property="og:title" content="...">
    // <meta name="author" content="...">
    // <title>...</title>
    // <div id="js_content">...</div>  // 微信文章正文
}
```

---

## 9️⃣ `markdown.go` - Markdown 转纯文本

**作用：** 将 Markdown 格式转换为微信可读的纯文本

### **转换规则**

| Markdown | 纯文本 |
|----------|--------|
| `# 标题` | `标题` |
| `**粗体**` | `粗体` |
| `*斜体*` | `斜体` |
| `~~删除~~` | `删除` |
| `> 引用` | `引用` |
| `- 列表` | `• 列表` |
| `` `代码` `` | `代码` |
| `![图片](url)` | *(删除)* |
| `[链接](url)` | `链接` |
| `\| 表格 \|` | `表格` (空格分隔) |

---

## 🏗️ 架构设计

```
┌─────────────────────────────────────────────────────────┐
│                    微信用户                               │
└──────────────────────┬──────────────────────────────────┘
                       │ 微信消息
                       ▼
┌─────────────────────────────────────────────────────────┐
│              Handler.HandleMessage()                     │
│  ┌───────────────────────────────────────────────────┐  │
│  │  1. 消息过滤 + 去重                                │  │
│  │  2. 提取文本/媒体                                   │  │
│  │  3. 命令解析 (@agent /agent)                       │  │
│  │  4. 内置命令路由                                   │  │
│  │  5. Agent 路由 (单播/广播)                          │  │
│  └───────────────────────────────────────────────────┘  │
└───────┬──────────┬──────────┬──────────┬────────────────┘
        │          │          │          │
        ▼          ▼          ▼          ▼
   ┌────────┐ ┌────────┐ ┌────────┐ ┌────────┐
   │ Agent  │ │  Hub   │ │ Todo/  │ │ Media/ │
   │ 调度   │ │ 上下文  │ │ Timer  │ │  CDN   │
   └───┬────┘ └───┬────┘ └───┬────┘ └───┬────┘
       │          │          │          │
       ▼          ▼          ▼          ▼
   AI 回复    文件共享    定时提醒    文件发送
```

---

## 📊 代码统计验证

| 文件 | 实际行数 | 说明 |
|------|---------|------|
| `handler.go` | ~2716 | ✓ 核心消息路由 |
| `timer.go` | ~436 | ✓ 定时器 |
| `todo.go` | ~380 | ✓ 待办事项 |
| `linkhoard.go` | ~300 | ✓ 网页抓取 |
| `cdn.go` | ~200 | ✓ CDN 加密 |
| `media.go` | ~195 | ✓ 媒体处理 |
| `attachment.go` | ~120 | ✓ 附件处理 |
| `markdown.go` | ~90 | ✓ Markdown 转换 |
| `sender.go` | ~85 | ✓ 消息发送 |
| `*_test.go` | ~400 | ✓ 单元测试 |
| **总计** | **~5022** | **5095 行**（含空行/注释） |

---

## 🎯 总结

`messaging` 文件夹是项目的**业务逻辑核心**：

| 功能模块 | 复杂度 | 代码量 | 说明 |
|---------|--------|--------|------|
| **消息路由** | ⭐⭐⭐⭐⭐ | 2716 行 | 命令解析、Agent 调度、内置命令 |
| **Hub 协作** | ⭐⭐⭐⭐ | 包含在 handler.go | 上下文注入、Pipe 链式调用 |
| **辩论功能** | ⭐⭐⭐⭐ | 包含在 handler.go | 多 Agent 多轮辩论 |
| **定时任务** | ⭐⭐⭐ | 436 + 380 行 | Timer + Todo + 后台调度器 |
| **媒体处理** | ⭐⭐⭐⭐ | 195 + 200 行 | CDN 加密上传/下载 |
| **网页抓取** | ⭐⭐⭐ | 300 行 | Jina API + HTML 解析 |
| **安全限制** | ⭐⭐⭐ | 120 行 | 路径沙箱 + 命令白名单 |

**设计特点：**
- **命令驱动**：所有功能通过 `/command` 触发
- **Agent 工厂**：按需启动 Agent，避免资源浪费
- **并发安全**：所有共享状态用 `sync.RWMutex` 保护
- **降级策略**：AI 解析失败时回退到默认行为
- **链式协作**：`/hub pipe` 实现多 Agent 流水线处理
