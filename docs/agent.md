# `agent` 文件夹详细分析

这个文件夹实现了**统一的 AI Agent 抽象层**，支持多种 AI 后端。

---

## 📁 文件清单

| 文件 | 行数 | 作用 |
|------|------|------|
| `agent.go` | ~130 | **核心接口定义** |
| `acp_agent.go` | ~1343 | **ACP 协议代理**（主力） |
| `http_agent.go` | ~195 | **HTTP API 代理** |
| `cli_agent.go` | ~280 | **CLI 工具代理** |
| `env_test.go` | ~60 | **环境变量测试** |

---

## 1️⃣ `agent.go` - 核心接口定义

**作用：** 定义所有 Agent 必须实现的统一接口

**关键类型：**
```go
type Agent interface {
    Chat(ctx, conversationID, message)           // 基础对话
    ChatWithMedia(ctx, conversationID, message, media) // 带媒体
    ResetSession(ctx, conversationID)            // 重置会话
    Info() AgentInfo                             // 获取元信息
    SetCwd(cwd)                                  // 设置工作目录
    SetProgressCallback(callback)                // 进度回调
}
```

**辅助类型：**
- `MediaEntry`: 媒体文件（图片/文件/视频）
- `AgentInfo`: Agent 元数据（名称、类型、模型、PID）
- `ProgressEvent`: 进度通知事件

**设计模式：** 策略模式（Strategy Pattern），允许运行时切换不同 Agent 实现

---

## 2️⃣ `acp_agent.go` - ACP 协议代理（核心实现）

**作用：** 通过 **JSON-RPC 2.0** 与本地 ACP 代理进程通信（如 `claude-agent-acp`、`codex-acp`）

**关键特性：**

| 特性 | 说明 |
|------|------|
| **双协议支持** | `legacy_acp`（标准 ACP）和 `codex_app_server`（Codex 原生协议） |
| **会话管理** | 维护 `conversationID -> sessionID/threadID` 映射 |
| **异步通信** | `pending` map 跟踪未完成的 RPC 请求 |
| **流式响应** | `readLoop` 解析 NDJSON 输出，实时收集文本块 |
| **媒体支持** | `buildPromptEntries` 处理图片/文件/视频 |

**通信流程：**
```
1. Start() → 启动子进程，建立 stdin/stdout 管道
2. initialize → 握手（ACP 协议）
3. session/new → 创建会话
4. session/prompt → 发送消息
5. readLoop → 持续读取 stdout，分发 session/update 通知
```

**代码占比：** 1343 行，占整个 agent 目录的 **65.8%**，是**最复杂的实现**

---

## 3️⃣ `http_agent.go` - HTTP API 代理

**作用：** 对接 **OpenAI 兼容的 HTTP API**（如 GPT-4、本地 LLM）

**关键特性：**
- 标准 OpenAI Chat Completions API 格式
- 本地维护对话历史（`history map`）
- 支持自定义 Headers 和 API Key
- 媒体支持有限（转为文本描述）

**适用场景：**
- 调用云端 API（GPT-4、Claude API）
- 本地 LLM 服务（Ollama、vLLM）

**代码量：** 195 行，**最简洁的实现**

---

## 4️⃣ `cli_agent.go` - CLI 工具代理

**作用：** 通过**命令行调用**本地 AI 工具（`claude` CLI、`codex`）

**关键特性：**
- 支持 `claude -p --output-format stream-json`
- 支持 `codex exec` 命令
- 会话恢复（`--resume sessionID`）
- 流式解析 JSON 事件

**限制：**
- 每次调用启动新进程（性能较差）
- 不支持进度回调
- 媒体支持有限

**代码量：** 280 行

---

## 5️⃣ `env_test.go` - 环境变量测试

**作用：** 测试 `mergeEnv` 函数的正确性

**测试用例：**
- 覆盖已有变量
- 添加新变量
- 拒绝非法 key（包含 `=`）
- 处理空值

**代码量：** 60 行

---

## 🏗️ 架构设计

```
┌─────────────────────────────────────┐
│         业务层 (messaging)          │
│   调用 Agent.Chat(message)          │
└──────────────┬──────────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│      Agent 接口 (agent.go)          │
│  - Chat                             │
│  - ChatWithMedia                    │
│  - ResetSession                     │
└───────┬──────────┬──────────┬───────┘
        │          │          │
        ▼          ▼          ▼
   ┌────────┐ ┌────────┐ ┌────────┐
   │  ACP   │ │  HTTP  │ │  CLI   │
   │ Agent  │ │ Agent  │ │ Agent  │
   └───┬────┘ └───┬────┘ └───┬────┘
       │          │          │
       ▼          ▼          ▼
   本地进程   HTTP API    CLI 命令
   (stdio)   (REST)     (exec)
```

---

## 📊 统计差异解释

之前统计显示 `agent` 占 54 KB / 2038 行，但实际只有 ~2000 行。差异原因：
- 统计工具可能**包含了依赖文件**或**重复计算**
- 或者 `acp_agent.go` 的 1343 行被统计工具**错误放大**

**实际占比：**
```
agent.go:       ~130 行 (6.4%)
acp_agent.go:   ~1343 行 (65.8%)  ← 绝对主力
http_agent.go:  ~195 行 (9.6%)
cli_agent.go:   ~280 行 (13.7%)
env_test.go:    ~60 行 (2.9%)
─────────────────────────────────
总计:           ~2008 行
```

---

## 🎯 总结

`agent` 文件夹实现了一个**灵活的 AI Agent 抽象层**：
- **统一接口**：业务代码无需关心底层 AI 实现
- **多后端支持**：ACP（本地）、HTTP（云端）、CLI（命令行）
- **核心复杂度**：集中在 `acp_agent.go`（处理 JSON-RPC 协议）
- **代码精简**：总约 2000 行，属于**轻量级适配层**
