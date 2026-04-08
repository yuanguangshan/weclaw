# `api` 文件夹详细分析

这个文件夹实现了 **HTTP API 服务层**，提供 RESTful 接口用于消息发送和管理后台。

---

## 📁 文件清单

| 文件 | 行数 | 作用 |
|------|------|------|
| `server.go` | ~140 | **HTTP 服务器 + 路由** |
| `admin.go` | ~908 | **管理后台 API 实现** |

---

## 1️⃣ `server.go` - HTTP 服务器核心

**作用：** 启动 HTTP 服务，注册所有路由

**关键结构：**
```go
type Server struct {
    clients  []*ilink.Client  // iLink 客户端列表
    addr     string           // 监听地址 (默认 127.0.0.1:18011)
    todosMu  sync.Mutex       // 保护 todos.json
    timersMu sync.Mutex       // 保护 timers.json
}
```

**路由分类：**

| 路由 | 方法 | 作用 |
|------|------|------|
| `/api/send` | POST | **发送消息**（核心功能） |
| `/health` | GET | 健康检查 |
| `/api/config` | GET/PUT | 配置读写 |
| `/api/agents` | GET/POST/PUT/DELETE | Agent 管理 |
| `/api/status` | GET | 服务状态 |
| `/api/accounts` | GET/DELETE | 账号管理 |
| `/api/login/*` | GET | 登录二维码/状态 |
| `/api/service/*` | POST | 重启/更新服务 |
| `/api/logs` | GET | 查看日志 |
| `/api/hub/*` | GET/DELETE/POST | 文件共享中心 |
| `/api/todos` | GET/POST/PUT/DELETE | 待办事项 |
| `/api/timers` | GET/POST/PUT | 定时器 |
| `/admin` | GET | 管理后台 UI |

**核心流程：**
```go
func (s *Server) Run(ctx) {
    mux := http.NewServeMux()
    // 注册路由...
    srv := &http.Server{Addr: s.addr, Handler: mux}
    srv.ListenAndServe()
}
```

---

## 2️⃣ `admin.go` - 管理后台 API 实现

**作用：** 实现所有管理接口的业务逻辑

---

### 📦 模块分解

#### **A. 配置与 Agent 管理** (行 40-170)

| 函数 | 作用 |
|------|------|
| `handleGetConfig` | 加载 `config.json` |
| `handleUpdateConfig` | 更新默认 Agent、API 地址、保存目录 |
| `handleListAgents` | 列出所有已配置的 Agent |
| `handleAddAgent` | 添加新 Agent（ACP/CLI/HTTP） |
| `handleUpdateAgent` | 修改 Agent 配置 |
| `handleDeleteAgent` | 删除 Agent |
| `handleDetectAgents` | 自动检测系统中可用的 Agent |

**数据流：**
```
请求 → 解析 JSON → 修改 config.json → 返回状态
```

---

#### **B. 状态与账号管理** (行 172-280)

| 函数 | 作用 |
|------|------|
| `handleStatus` | 返回服务运行状态（PID、运行时间、Agent 数量等） |
| `handleListAccounts` | 列出所有 iLink 账号（隐藏敏感 token） |
| `handleDeleteAccount` | 删除账号配置文件 |

**状态信息示例：**
```json
{
  "running": true,
  "pid": 12345,
  "uptime": "2h30m",
  "agent_count": 3,
  "account_count": 2,
  "hub_files": 5,
  "default_agent": "claude-acp",
  "version": "dev"
}
```

---

#### **C. 日志查看** (行 282-310)

| 函数 | 作用 |
|------|------|
| `handleLogs` | 读取 `~/.weclaw/weclaw.log` 最后 N 行 |

**实现特点：**
- 默认 200 行，最多 2000 行
- 使用滑动窗口读取，避免大文件内存溢出

---

#### **D. Hub 文件管理** (行 312-365)

| 函数 | 作用 |
|------|------|
| `handleListHub` | 列出共享目录中的所有文件 |
| `handleReadHubFile` | 读取指定文件内容 |
| `handleDeleteHubFile` | 删除指定文件 |
| `handleClearHub` | 清空整个共享目录 |

**用途：** Hub 是 Agent 之间共享文件的目录（如生成的图片、文档）

---

#### **E. 待办事项 (Todos)** (行 367-490)

| 函数 | 作用 |
|------|------|
| `handleListTodos` | 列出所有待办 |
| `handleAddTodo` | 添加新待办 |
| `handleDoneTodo` | 标记完成（status=1） |
| `handleDeleteTodo` | 删除待办 |

**数据结构：**
```go
type Todo struct {
    ID        int    `json:"id"`
    UserID    string `json:"user_id"`
    Title     string `json:"title"`
    DueTime   int64  `json:"due_time"`
    Status    int    `json:"status"` // 0=待办, 1=已完成
    CreatedAt int64  `json:"created_at"`
    Reminded  bool   `json:"reminded"`
}
```

**持久化：** 存储在 `~/.weclaw/hub/todos.json`

---

#### **F. 定时器 (Timers)** (行 492-580)

| 函数 | 作用 |
|------|------|
| `handleListTimers` | 列出所有定时器 |
| `handleAddTimer` | 创建定时器（最大 24 小时） |
| `handleCancelTimer` | 取消定时器（status=2） |

**数据结构：**
```go
type Timer struct {
    ID        int    `json:"id"`
    UserID    string `json:"user_id"`
    Label     string `json:"label"`
    Duration  int64  `json:"duration"`  // 秒
    EndTime   int64  `json:"end_time"`  // 到期时间戳
    Status    int    `json:"status"`    // 0=运行, 2=已取消
}
```

**用途：** Agent 可以设置定时提醒（如"30 分钟后提醒我开会"）

---

#### **G. 登录管理** (行 582-630)

| 函数 | 作用 |
|------|------|
| `handleLoginQRCode` | 获取 iLink 登录二维码 |
| `handleLoginStatus` | 轮询登录状态，成功后保存凭证 |

**流程：**
```
1. 前端请求 /api/login/qrcode → 获取二维码 URL
2. 用户扫码后，前端轮询 /api/login/status?qrcode=xxx
3. 登录成功 → 保存 Credentials 到 ~/.weclaw/accounts/
```

---

#### **H. 服务控制** (行 632-700)

| 函数 | 作用 |
|------|------|
| `handleServiceRestart` | 调用 `systemctl restart weclaw` |
| `handleServiceUpdate` | **在线更新**：编译 → 停服 → 替换二进制 → 启动 |

**更新流程（关键）：**
```go
1. go build -o /tmp/weclaw-new  // 编译到临时文件
2. 先响应客户端（避免断开连接）
3. 后台执行：
   - systemctl stop weclaw
   - cp /tmp/weclaw-new /usr/local/bin/weclaw
   - systemctl start weclaw
```

**设计亮点：**
- 先响应再更新，避免客户端超时
- 使用临时文件，避免覆盖正在运行的二进制
- 失败回退：如果替换失败，仍会启动旧版本

---

#### **I. 辅助函数** (行 702-750)

| 函数 | 作用 |
|------|------|
| `writeJSON` | 统一 JSON 响应格式 |
| `writeError` | 统一错误响应 |
| `CopyFile` | 文件复制（用于更新二进制） |
| `normalizeAccountID` | 将 Bot ID 转为合法文件名 |

---

## 🏗️ 架构设计

```
┌─────────────────────────────────────────┐
│         前端 (web/admin.html)           │
│         或外部调用方 (curl/API)          │
└──────────────┬──────────────────────────┘
               │ HTTP REST API
               ▼
┌─────────────────────────────────────────┐
│          api.Server                     │
│  ┌───────────────────────────────────┐  │
│  │  server.go: 路由 + HTTP 服务器     │  │
│  └───────────────────────────────────┘  │
│  ┌───────────────────────────────────┐  │
│  │  admin.go: 业务逻辑实现           │  │
│  │  - 配置管理                        │  │
│  │  - Agent 管理                      │  │
│  │  - Todos/Timers                    │  │
│  │  - Hub 文件                        │  │
│  │  - 服务控制                        │  │
│  └───────────────────────────────────┘  │
└───────┬──────────┬──────────┬───────────┘
        │          │          │
        ▼          ▼          ▼
   config.json  hub/      ilink.Client
   accounts/    todos.json  messaging
```

---

## 📊 代码统计验证

| 文件 | 实际行数 | 统计显示 |
|------|---------|---------|
| `server.go` | ~140 | 包含在 908 行中 |
| `admin.go` | ~908 | ✓ |
| **总计** | **~1048** | **1084 行**（含空行/注释） |

---

## 🎯 总结

`api` 文件夹实现了**完整的管理后台**：

| 功能模块 | 复杂度 | 说明 |
|---------|--------|------|
| **消息发送** | ⭐ | 核心功能，调用 `messaging` 层 |
| **配置管理** | ⭐⭐ | CRUD 操作，持久化到 JSON |
| **Agent 管理** | ⭐⭐ | 动态增删改查 AI 代理 |
| **Todos/Timers** | ⭐⭐ | 简单的 CRUD + 定时任务 |
| **Hub 文件** | ⭐⭐ | 文件浏览/读取/删除 |
| **服务控制** | ⭐⭐⭐ | 在线更新，需处理进程生命周期 |
| **登录管理** | ⭐⭐⭐ | 二维码轮询，凭证保存 |

**设计特点：**
- **RESTful 风格**：资源导向，HTTP 动词操作
- **无状态**：每次请求读取最新配置
- **简单持久化**：JSON 文件，无数据库
- **单文件实现**：`admin.go` 908 行包含所有业务逻辑，适合小型项目
