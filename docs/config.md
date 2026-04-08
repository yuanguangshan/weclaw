# `config` 文件夹详细分析

这个文件夹实现了**配置管理 + Agent 自动检测**机制。

---

## 📁 文件清单

| 文件 | 行数 | 作用 |
|------|------|------|
| `config.go` | ~130 | **配置加载/保存** |
| `detect.go` | ~250 | **Agent 自动检测** |
| `config_test.go` | ~100 | **配置测试** |
| `detect_test.go` | ~80 | **检测测试** |

---

## 1️⃣ `config.go` - 配置加载/保存

**作用：** 管理 `~/.weclaw/config.json` 的读写

### **数据结构**

```go
type Config struct {
    DefaultAgent string                 `json:"default_agent"`  // 默认 Agent 名称
    APIAddr      string                 `json:"api_addr"`       // API 监听地址
    SaveDir      string                 `json:"save_dir"`       // 文件保存目录
    Agents       map[string]AgentConfig `json:"agents"`         // Agent 配置映射
}

type AgentConfig struct {
    Type         string            `json:"type"`           // "acp" | "cli" | "http"
    Command      string            `json:"command"`        // 二进制路径
    Args         []string          `json:"args"`           // 额外参数
    Aliases      []string          `json:"aliases"`        // 自定义触发名
    Cwd          string            `json:"cwd"`            // 工作目录
    Env          map[string]string `json:"env"`            // 环境变量
    Model        string            `json:"model"`          // 模型名称
    SystemPrompt string            `json:"system_prompt"`  // 系统提示词
    Endpoint     string            `json:"endpoint"`       // HTTP 端点 (http 类型)
    APIKey       string            `json:"api_key"`        // API 密钥 (http 类型)
    Headers      map[string]string `json:"headers"`        // 额外请求头 (http 类型)
    MaxHistory   int               `json:"max_history"`    // 历史长度 (http 类型)
}
```

### **配置加载优先级**

```
1. 磁盘文件 (~/.weclaw/config.json)
2. 环境变量覆盖
   - WECLAW_DEFAULT_AGENT → DefaultAgent
   - WECLAW_API_ADDR      → APIAddr
   - WECLAW_SAVE_DIR      → SaveDir
3. 默认值（空配置）
```

**加载流程：**
```go
func Load() (*Config, error) {
    cfg := DefaultConfig()
    
    // 1. 读取 JSON 文件
    data := os.ReadFile(path)
    json.Unmarshal(data, cfg)
    
    // 2. 环境变量覆盖（只覆盖顶层字段，不覆盖 Agent 配置）
    loadEnv(cfg)
    
    return cfg, nil
}
```

### **别名映射**

```go
func BuildAliasMap(agents map[string]AgentConfig) map[string]string {
    // 从所有 Agent 配置中提取自定义别名
    // 例: {"gpt": "claude", "4o": "claude"}
    
    // 保护机制：
    // 1. 保留内置命令不能被覆盖 (info, help, new, clear, cwd)
    // 2. 检测别名冲突并警告
    // 3. 检测别名遮蔽 Agent 名称并警告
}
```

---

## 2️⃣ `detect.go` - Agent 自动检测

**作用：** 启动时自动扫描系统中已安装的 AI Agent，并写入配置

### **检测候选列表**

```go
var agentCandidates = []agentCandidate{
    // claude: 优先 ACP，降级 CLI
    {Name: "claude", Binary: "claude-agent-acp", Type: "acp", Model: "sonnet"},
    {Name: "claude", Binary: "claude",             Type: "cli", Model: "sonnet"},
    
    // codex: 优先 ACP，其次 app-server，最后 CLI
    {Name: "codex", Binary: "codex-acp",           Type: "acp"},
    {Name: "codex", Binary: "codex", Args: ["app-server", "--listen", "stdio://"], 
                          CheckArgs: ["app-server", "--help"], Type: "acp"},
    {Name: "codex", Binary: "codex",               Type: "cli"},
    
    // 纯 ACP Agent
    {Name: "cursor",   Binary: "agent",  Args: ["acp"],              Type: "acp"},
    {Name: "kimi",     Binary: "kimi",   Args: ["acp"],              Type: "acp"},
    {Name: "gemini",   Binary: "gemini", Args: ["--acp"],            Type: "acp"},
    {Name: "opencode", Binary: "opencode", Args: ["acp"],            Type: "acp"},
    {Name: "openclaw", Binary: "openclaw",                            Type: "acp"},
    {Name: "pi",       Binary: "pi-acp",                              Type: "acp"},
    {Name: "copilot",  Binary: "copilot", Args: ["--acp", "--stdio"], Type: "acp"},
    {Name: "droid",    Binary: "droid",   Args: ["exec", "--output-format", "acp"], Type: "acp"},
    {Name: "iflow",    Binary: "iflow",   Args: ["--experimental-acp"], Type: "acp"},
    {Name: "kiro",     Binary: "kiro-cli", Args: ["acp"],             Type: "acp"},
    {Name: "qwen",     Binary: "qwen",    Args: ["--acp"],            Type: "acp"},
}
```

**设计原则：**
- **ACP 优先**：同一 Agent 先尝试 ACP 模式，再降级到 CLI
- **能力探测**：用 `CheckArgs` 探测 Agent 是否支持特定协议
- **顺序优先**：列表中靠前的候选者优先被选用

---

### **检测流程**

```
┌─────────────────────────────────────────────────────────┐
│  DetectAndConfigure(cfg)                                │
├─────────────────────────────────────────────────────────┤
│  1. 遍历 agentCandidates                                │
│     - 已配置？→ 跳过                                    │
│     - lookPath(binary) → 找不到？→ 跳过                 │
│     - CheckArgs 探测失败？→ 跳过                        │
│     - 写入 cfg.Agents[name]                             │
├─────────────────────────────────────────────────────────┤
│  2. openclaw 特殊处理                                    │
│     - 检测到 ACP 但无 Args → 尝试 HTTP 模式              │
│     - 读取 ~/.openclaw/openclaw.json 获取网关配置        │
│     - 优先 HTTP (避免会话路由冲突)                       │
│     - 同时注册 openclaw-acp 供用户手动切换               │
├─────────────────────────────────────────────────────────┤
│  3. 选择默认 Agent                                       │
│     - 按 defaultOrder 优先级选择第一个已检测到的 Agent    │
│     - 优先级: claude > codex > cursor > kimi > ...      │
├─────────────────────────────────────────────────────────┤
│  4. 返回 modified (是否修改了配置)                       │
└─────────────────────────────────────────────────────────┘
```

---

### **openclaw 特殊处理**

openclaw 是一个特殊情况，因为它有**两种运行模式**：

```go
// 模式 1: HTTP 网关模式（优先）
// 避免会话路由冲突（issue #9）
cfg.Agents["openclaw"] = AgentConfig{
    Type:     "http",
    Endpoint: "https://gateway/v1/chat/completions",
    APIKey:   gwToken,
    Headers:  {"x-openclaw-scopes": "operator.write"},
    Model:    "openclaw:main",
}

// 模式 2: ACP 模式（备用）
cfg.Agents["openclaw-acp"] = AgentConfig{
    Type:    "acp",
    Command: "openclaw",
    Args:    ["acp", "--url", gwURL, "--token", gwToken],
    Model:   "openclaw:main",
}
```

**网关配置读取优先级：**
```
1. 环境变量
   - OPENCLAW_GATEWAY_URL
   - OPENCLAW_GATEWAY_TOKEN
   - OPENCLAW_GATEWAY_PASSWORD

2. 配置文件 ~/.openclaw/openclaw.json
   - gateway.remote.url (远程网关)
   - gateway.port + gateway.auth (本地网关)
```

---

### **二进制查找策略 `lookPath()`**

```go
func lookPath(binary string) (string, error) {
    // 快速路径：在当前 PATH 中查找
    if p, err := exec.LookPath(binary); err == nil {
        return p, nil
    }
    
    // 降级路径：通过登录 shell 查找
    // 这会加载 ~/.zshrc / ~/.bashrc，解决 nvm/mise 等版本管理工具的问题
    shell := "zsh"  // macOS 默认
    if runtime.GOOS != "darwin" {
        shell = "bash"  // Linux 默认
    }
    out := exec.Command(shell, "-lic", "which "+binary).Output()
    return strings.TrimSpace(out), nil
}
```

**解决的问题：**
```
守护进程环境: PATH = /usr/local/bin:/usr/bin:/bin
              → exec.LookPath("claude") 失败（claude 在 nvm 管理的目录）
              
登录 shell:   zsh -lic "which claude"
              → 加载 ~/.zshrc → nvm 初始化 → 找到 claude
```

---

### **能力探测 `commandProbe()`**

```go
func commandProbe(binary string, args []string) bool {
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()
    
    cmd := exec.CommandContext(ctx, binary, args...)
    cmd.Stdout = io.Discard
    cmd.Stderr = io.Discard
    
    return cmd.Run() == nil  // 退出码 0 = 支持
}
```

**用途：** 探测 `codex app-server --help` 是否可用，以确定是否支持该协议

---

### **默认 Agent 选择**

```go
var defaultOrder = []string{
    "claude", "codex", "cursor", "kimi", "gemini", "opencode", "openclaw",
    "pi", "copilot", "droid", "iflow", "kiro", "qwen",
}

// 选择第一个已检测到的 Agent 作为默认
for _, name := range defaultOrder {
    if _, ok := cfg.Agents[name]; ok {
        cfg.DefaultAgent = name
        break
    }
}
```

---

## 📊 配置示例

**`~/.weclaw/config.json`：**
```json
{
  "default_agent": "claude",
  "api_addr": "127.0.0.1:18011",
  "save_dir": "/Users/ygs/Documents/weclaw-saves",
  "agents": {
    "claude": {
      "type": "acp",
      "command": "/usr/local/bin/claude-agent-acp",
      "model": "sonnet",
      "cwd": "/Users/ygs/projects/myapp",
      "env": {
        "ANTHROPIC_API_KEY": "sk-ant-xxx"
      }
    },
    "gpt": {
      "type": "http",
      "endpoint": "https://api.openai.com/v1/chat/completions",
      "api_key": "sk-xxx",
      "model": "gpt-4o",
      "aliases": ["4o", "gpt4"]
    }
  }
}
```

---

## 🏗️ 配置生命周期

```
┌─────────────────────────────────────────────────────────┐
│  1. 启动时                                               │
│     config.Load() → 读取 config.json                     │
│     config.DetectAndConfigure() → 自动检测新 Agent        │
│     如果 modified → config.Save()                        │
├─────────────────────────────────────────────────────────┤
│  2. 运行时                                               │
│     用户通过 /api/agents 增删改 Agent                     │
│     每次修改 → config.Save()                             │
├─────────────────────────────────────────────────────────┤
│  3. 环境变量覆盖                                         │
│     WECLAW_DEFAULT_AGENT=codex weclaw start              │
│     → 临时覆盖默认 Agent，不写入磁盘                      │
└─────────────────────────────────────────────────────────┘
```

---

## 📊 代码统计验证

| 文件 | 实际行数 | 说明 |
|------|---------|------|
| `detect.go` | ~250 | ✓ Agent 自动检测 |
| `config.go` | ~130 | ✓ 配置加载/保存 |
| `config_test.go` | ~100 | ✓ 配置测试 |
| `detect_test.go` | ~80 | ✓ 检测测试 |
| **总计** | **~560** | |

---

## 🎯 总结

`config` 文件夹实现了**智能配置管理**：

| 功能模块 | 复杂度 | 说明 |
|---------|--------|------|
| **配置加载** | ⭐⭐ | JSON 读写 + 环境变量覆盖 |
| **别名映射** | ⭐⭐ | 自定义触发名 + 冲突检测 |
| **Agent 检测** | ⭐⭐⭐⭐ | 13 种 Agent，多协议探测 |
| **二进制查找** | ⭐⭐⭐ | PATH + 登录 shell 降级 |
| **openclaw 特殊处理** | ⭐⭐⭐ | HTTP/ACP 双模式 + 网关配置 |
| **默认选择** | ⭐⭐ | 优先级列表 |

**设计特点：**
- **零配置启动**：自动检测已安装的 Agent
- **渐进式配置**：检测到的配置持久化到 JSON
- **环境变量优先**：便于容器化部署
- **多协议支持**：同一 Agent 可注册多种模式（如 openclaw + openclaw-acp）
- **安全警告**：别名冲突和遮蔽检测
