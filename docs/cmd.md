# `cmd` 文件夹详细分析

这个文件夹实现了 **CLI 命令行工具**，是用户与 weclaw 交互的主要入口。

---

## 📁 文件清单

| 文件 | 行数 | 作用 |
|------|------|------|
| `root.go` | ~25 | **根命令定义** |
| `start.go` | ~453 | **启动服务**（核心） |
| `send.go` | ~70 | **发送消息** |
| `status.go` | ~30 | **查看状态** |
| `stop.go` | ~20 | **停止服务** |
| `restart.go` | ~45 | **重启服务** |
| `update.go` | ~200 | **在线更新** |
| `login.go` | ~35 | **扫码登录** |
| `proc_unix.go` | ~12 | **Unix 进程配置** |
| `proc_windows.go` | ~10 | **Windows 进程配置** |

---

## 1️⃣ `root.go` - 根命令定义

**作用：** 定义 `weclaw` 命令的入口

```go
var rootCmd = &cobra.Command{
    Use:     "weclaw",
    Short:   "WeChat AI agent bridge",
    RunE:    runStart,  // 默认执行 start
}
```

**特点：**
- 使用 [Cobra](https://github.com/spf13/cobra) 框架
- 默认命令是 `start`（直接运行 `weclaw` 等同于 `weclaw start`）
- 版本号通过 `-ldflags` 编译时注入

---

## 2️⃣ `start.go` - 启动服务（核心文件）

**作用：** 启动整个 weclaw 服务（前台/后台模式）

**代码占比：** 453 行，占 cmd 目录的 **50.5%**，是**最复杂的文件**

---

### 📦 功能模块

#### **A. 命令定义** (行 15-25)

```go
startCmd.Flags().BoolVarP(&foregroundFlag, "foreground", "f", false, "前台运行")
startCmd.Flags().StringVar(&apiAddrFlag, "api-addr", "", "API 监听地址")
```

**标志：**
- `-f, --foreground`: 前台运行（默认后台守护）
- `--api-addr`: 自定义 API 端口

---

#### **B. 启动流程** `runStart()` (行 27-180)

```
┌─────────────────────────────────────────┐
│  1. 检查是否前台模式                     │
│     - 否 → 检查登录 → 启动守护进程      │
│     - 是 ↓                              │
├─────────────────────────────────────────┤
│  2. 加载账号凭证                         │
│     - 无账号 → 触发 doLogin()            │
├─────────────────────────────────────────┤
│  3. 加载配置 + 自动检测 Agent            │
│     - DetectAndConfigure()               │
├─────────────────────────────────────────┤
│  4. 创建 Handler（消息处理器）            │
│     - Agent 工厂函数                     │
│     - 设置默认 Agent                     │
│     - 启动 Todo/Timer 调度器             │
├─────────────────────────────────────────┤
│  5. 启动 HTTP API 服务器                 │
│     - api.NewServer(clients, addr)       │
├─────────────────────────────────────────┤
│  6. 启动 Monitor（消息监听）              │
│     - 每个账号一个 goroutine              │
│     - runMonitorWithRestart()            │
└─────────────────────────────────────────┘
```

**关键代码：**
```go
// Agent 工厂：按需创建 Agent
handler := messaging.NewHandler(
    func(ctx, name) agent.Agent {
        return createAgentByName(ctx, cfg, name)
    },
    func(name) error {
        cfg.DefaultAgent = name
        return config.Save(cfg)
    },
)
```

---

#### **C. Agent 工厂** `createAgentByName()` (行 210-280)

**作用：** 根据配置动态创建 Agent 实例

```go
switch agCfg.Type {
case "acp":
    ag := agent.NewACPAgent(...)
    ag.Start(ctx)  // ACP 需要启动子进程
case "cli":
    ag := agent.NewCLIAgent(...)  // CLI 无需启动，每次调用 exec
case "http":
    ag := agent.NewHTTPAgent(...) // HTTP 无需启动，每次 HTTP 请求
}
```

**特点：**
- 支持**大小写不敏感**的 Agent 名称匹配
- ACP Agent 会立即启动子进程
- CLI/HTTP Agent 延迟创建（按需调用）

---

#### **D. 登录流程** `doLogin()` (行 282-320)

```go
1. FetchQRCode() → 获取二维码
2. qrterminal.Generate() → 终端打印二维码
3. PollQRStatus() → 轮询等待用户扫码确认
4. SaveCredentials() → 保存凭证到 ~/.weclaw/accounts/
```

**用户体验：**
```
Scan this QR code with WeChat:
█████████████████████
█ ▄▄▄▄▄ █▀█ ▄▄▄▄▄ █
█ █   █ █▀▀ █   █ █
█ █▄▄▄█ █▀▄ █▄▄▄█ █
█▄▄▄▄▄▄▄█▄▀▄█▄▄▄▄▄▄█
QR URL: https://...
Waiting for scan...
QR code scanned! Please confirm on your phone.
Login confirmed!
```

---

#### **E. 守护进程模式** `runDaemon()` (行 340-390)

```go
1. stopAllWeclaw() → 杀掉旧进程
2. 创建 ~/.weclaw/ 目录
3. 打开 weclaw.log 文件
4. 重新执行自己：exec.Command(exe, "start", "-f")
5. 输出重定向到日志文件
6. 保存 PID 到 weclaw.pid
7. Process.Release() → 脱离父进程
```

**关键：**
```go
cmd := exec.Command(exe, "start", "-f")  // 前台模式
cmd.Stdout = lf
cmd.Stderr = lf
setSysProcAttr(cmd)  // Unix: Setsid=true 脱离会话
cmd.Start()
cmd.Process.Release()  // 不等待子进程
```

---

#### **F. Monitor 重启逻辑** `runMonitorWithRestart()` (行 182-208)

```go
for {
    monitor.Run(ctx)  // 运行直到断开连接
    
    if ctx.Err() != nil { return }  // 用户取消
    
    // 指数退避重启：3s → 6s → 12s → 30s（上限）
    restartDelay *= 2
    if restartDelay > maxRestartDelay {
        restartDelay = maxRestartDelay
    }
    
    time.Sleep(restartDelay)
}
```

**设计目的：** 网络断开时自动重连，避免频繁重启消耗资源

---

## 3️⃣ `send.go` - 发送消息

**作用：** 通过 CLI 快速发送消息（无需启动服务）

```bash
weclaw send --to "user_id@im.wechat" --text "Hello"
weclaw send --to "user_id" --media "https://example.com/image.png"
```

**实现：**
```go
1. 加载账号凭证
2. 创建 ilink.Client
3. messaging.SendTextReply() / SendMediaFromURL()
```

**用途：** 脚本自动化、定时任务、CI/CD 通知

---

## 4️⃣ `status.go` - 查看状态

**作用：** 检查后台服务是否运行

```go
pid, err := readPid()
if processExists(pid) {
    fmt.Printf("weclaw is running (pid=%d)\n", pid)
} else {
    fmt.Println("weclaw is not running (stale pid file)")
}
```

**进程检测：**
```go
func processExists(pid int) bool {
    p, _ := os.FindProcess(pid)
    return p.Signal(syscall.Signal(0)) == nil  // 信号 0 检测存活
}
```

---

## 5️⃣ `stop.go` - 停止服务

**作用：** 终止后台守护进程

```go
func stopAllWeclaw() {
    // 1. 通过 PID 文件杀进程
    p.Signal(syscall.SIGTERM)
    os.Remove(pidFile())
    
    // 2. 扫描残留进程（防止 PID 文件丢失）
    exec.Command("pkill", "-f", exe+" start").Run()
}
```

**双重保险：**
- 先尝试 PID 文件
- 再用 `pkill` 清理残留

---

## 6️⃣ `restart.go` - 重启服务

**作用：** 优雅重启（stop → start）

```go
1. 停止旧进程（等待最多 10 秒）
2. 清理 PID 文件
3. runDaemon() → 启动新进程
```

---

## 7️⃣ `update.go` - 在线更新

**作用：** 从 GitHub 自动更新到最新版本

**更新流程：**
```
┌─────────────────────────────────────────┐
│  1. GET /repos/fastclaw-ai/weclaw/releases/latest  │
│     → 获取最新版本号                     │
├─────────────────────────────────────────┤
│  2. 对比当前版本                         │
│     - 相同 → "Already up to date"        │
├─────────────────────────────────────────┤
│  3. 下载新二进制                         │
│     URL: github.com/.../weclaw_darwin_amd64 │
├─────────────────────────────────────────┤
│  4. 替换当前二进制                       │
│     - os.Rename() → 直接移动             │
│     - 失败 → sudo cp                     │
│     - macOS → xattr 清除隔离属性         │
├─────────────────────────────────────────┤
│  5. 重启服务（如果正在运行）              │
└─────────────────────────────────────────┘
```

**macOS 特殊处理：**
```go
if runtime.GOOS == "darwin" {
    exec.Command("xattr", "-d", "com.apple.quarantine", exePath).Run()
    exec.Command("xattr", "-d", "com.apple.provenance", exePath).Run()
}
```

**原因：** macOS Gatekeeper 会拦截未签名的下载文件，需清除隔离属性

---

## 8️⃣ `login.go` - 扫码登录

**作用：** 独立登录命令（可在服务未启动时添加账号）

```bash
weclaw login
```

**实现：** 调用 `doLogin()`（与 `start.go` 共享）

---

## 9️⃣ `proc_unix.go` / `proc_windows.go` - 跨平台进程配置

**作用：** 守护进程脱离会话的跨平台实现

**Unix:**
```go
cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
```
- `Setsid=true`：创建新会话，脱离控制终端

**Windows:**
```go
func setSysProcAttr(_ *exec.Cmd) {}  // 空操作
```
- Windows 无需特殊处理

**构建标签：**
```go
//go:build !windows   // proc_unix.go
//go:build windows    // proc_windows.go
```

---

## 🏗️ 命令结构

```
weclaw                          # 默认执行 start
├── start [-f] [--api-addr]     # 启动服务（前台/后台）
├── login                       # 扫码添加账号
├── send --to --text --media    # 发送消息
├── status                      # 查看运行状态
├── stop                        # 停止服务
├── restart                     # 重启服务
├── update / upgrade            # 在线更新
└── version                     # 查看版本
```

---

## 📊 代码统计验证

| 文件 | 实际行数 | 说明 |
|------|---------|------|
| `start.go` | ~453 | ✓ 核心启动逻辑 |
| `update.go` | ~200 | ✓ 在线更新 |
| `send.go` | ~70 | ✓ 消息发送 |
| `restart.go` | ~45 | ✓ 重启 |
| `login.go` | ~35 | ✓ 登录 |
| `status.go` | ~30 | ✓ 状态 |
| `root.go` | ~25 | ✓ 根命令 |
| `stop.go` | ~20 | ✓ 停止 |
| `proc_unix.go` | ~12 | ✓ Unix 配置 |
| `proc_windows.go` | ~10 | ✓ Windows 配置 |
| **总计** | **~890** | **898 行**（含空行/注释） |

---

## 🎯 总结

`cmd` 文件夹实现了**完整的 CLI 工具**：

| 功能模块 | 复杂度 | 说明 |
|---------|--------|------|
| **启动服务** | ⭐⭐⭐⭐⭐ | 核心逻辑，Agent 工厂、守护进程、Monitor |
| **在线更新** | ⭐⭐⭐⭐ | GitHub API、下载、替换、重启 |
| **登录管理** | ⭐⭐⭐ | 二维码生成、轮询确认 |
| **消息发送** | ⭐⭐ | 简单调用 messaging 层 |
| **进程管理** | ⭐⭐⭐ | start/stop/restart/status |
| **跨平台** | ⭐⭐ | Unix/Windows 进程脱离 |

**设计特点：**
- **Cobra 框架**：标准 Go CLI 结构
- **守护进程模式**：自实现，无外部依赖
- **指数退避重启**：网络不稳定时自动恢复
- **跨平台兼容**：构建标签区分 Unix/Windows
- **单二进制部署**：`update` 命令直接替换二进制
