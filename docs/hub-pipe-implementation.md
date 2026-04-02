# Hub Pipe 功能实现文档

## 概述

`/hub pipe` 是 weclaw 中的一个核心功能，用于实现 **Agent 之间的链式协作**。它允许用户将消息依次发送给多个 Agent，每个 Agent 可以基于前一个 Agent 的输出进行工作，实现智能流的传递。

### 核心价值

1. **上下文传递**：自动保存中间结果到 Hub，实现 Agent 间共享上下文
2. **智能链式**：支持任意 Agent 之间的组合与协作
3. **透明可追溯**：保存每步结果，便于审查和调试

---

## 命令格式

```
/hub pipe <目标agent> <消息>
```

**示例：**
```
/hub pipe gemini 量子计算的基本原理
```

**执行流程：**
1. 将消息发送给**默认 Agent**（如 nanobot）
2. 自动保存默认 Agent 的回复到 Hub
3. 将保存的内容注入给**目标 Agent**（如 gemini）
4. 返回目标 Agent 的分析结果

---

## 实现原理

### 架构图

```
用户消息
    │
    ▼
┌─────────────────────────────────────┐
│  Step 1: 发送给默认 Agent           │
│  (如 nanobot)                       │
└─────────────────────────────────────┘
    │
    ▼ 回复内容
┌─────────────────────────────────────┐
│  Step 2: 自动保存到 Hub              │
│  文件名: pipe_时间戳_源agent.md      │
└─────────────────────────────────────┘
    │
    ▼ Hub 上下文
┌─────────────────────────────────────┐
│  Step 3: 读取 Hub 内容               │
│  构造带上下文的 Prompt               │
└─────────────────────────────────────┘
    │
    ▼
┌─────────────────────────────────────┐
│  Step 4: 发送给目标 Agent            │
│  (如 gemini)                        │
└─────────────────────────────────────┘
    │
    ▼ 最终结果
┌─────────────────────────────────────┐
│  Step 5: 保存最终结果                │
│  返回给用户                          │
└─────────────────────────────────────┘
```

---

## 代码实现

### 核心函数：`handlePipe`

**文件位置：** `messaging/handler.go`

**函数签名：**
```go
func (h *Handler) handlePipe(ctx context.Context, client *ilink.Client,
    msg ilink.WeixinMessage, targetAgent, message, clientID string) string
```

### 实现步骤

#### 1. 获取源 Agent（默认 Agent）

```go
sourceAgent := h.getDefaultAgent()
if sourceAgent == nil {
    return "❌ 没有可用的默认 agent，请先设置默认 agent（如 /claude）"
}

// 使用配置名称而不是 Info().Name（后者可能返回进程路径）
h.mu.RLock()
sourceAgentName := h.defaultName
h.mu.RUnlock()
```

**关键点：**
- 使用 `h.defaultName` 而非 `sourceAgent.Info().Name`
- 避免显示进程路径（如 `/home/nanobot/.nanobot/.venv/bin/python3`）
- 使用读写锁保证线程安全

#### 2. 发送消息给源 Agent

```go
log.Printf("[hub/pipe] step1: sending to default agent (%s)", sourceAgentName)
reply1, err := h.chatWithAgent(ctx, sourceAgent, msg.FromUserID,
    message, client, msg.ContextToken)
if err != nil {
    return fmt.Sprintf("❌ 第一步（默认 agent %s）失败: %v", sourceAgentName, err)
}
```

**关键点：**
- 使用 `chatWithAgent` 发送消息
- 传入 `clientID` 确保会话隔离
- 记录日志便于调试

#### 3. 自动保存到 Hub

```go
timestamp := time.Now().Format("20060102-150405")

// 使用简洁的文件名：pipe_<timestamp>_<agent>.md
shortAgentName := sourceAgentName
if idx := strings.LastIndex(sourceAgentName, "/"); idx >= 0 {
    shortAgentName = sourceAgentName[idx+1:]
}
filename := fmt.Sprintf("pipe_%s_%s.md", timestamp, shortAgentName)

savePath, err := h.hub.Save(filename, reply1, sourceAgentName)
if err != nil {
    log.Printf("[hub/pipe] save failed: %v", err)
    // 即使保存失败，仍继续执行第二步（降级处理）
    filename = ""
}
```

**关键点：**
- 文件命名：`pipe_时间戳_agent名.md`
- 提取 basename 处理路径情况
- **降级处理**：保存失败不中断流程

#### 4. 获取目标 Agent

```go
targetAg, err := h.getAgent(ctx, targetAgent)
if err != nil {
    return fmt.Sprintf("❌ 目标 agent %q 不可用: %v", targetAgent, err)
}
```

#### 5. 构造第二步 Prompt（注入 Hub 上下文）

```go
var hubContext string
if filename != "" {
    hubContext, err = h.hub.ReadSpecific([]string{filename})
    if err != nil {
        log.Printf("[hub/pipe] read saved file failed: %v", err)
        hubContext = ""
    }
}

// 构造第二步的 prompt
prompt2 := message
if hubContext != "" {
    prompt2 = fmt.Sprintf(
        "【第一步 %s 的回复】\n\n%s\n\n---\n\n"+
        "请基于以上内容，用你的视角进行分析或补充：\n\n%s",
        sourceAgentName, hubContext, message)
}
```

**关键点：**
- 使用 `hub.ReadSpecific()` 读取特定文件
- Prompt 结构：上下文 + 分隔符 + 任务指令
- 确保 Agent 理解这是基于前一步的延续

#### 6. 发送给目标 Agent

```go
log.Printf("[hub/pipe] step2: sending to target agent (%s)", targetAgent)
reply2, err := h.chatWithAgent(ctx, targetAg, msg.FromUserID,
    prompt2, client, msg.ContextToken)
if err != nil {
    return fmt.Sprintf("❌ 第二步（目标 agent %s）失败: %v", targetAgent, err)
}
```

#### 7. 自动保存最终结果

```go
finalFilename := fmt.Sprintf("pipe_%s_%s_final.md", timestamp, targetAgent)
_, err = h.hub.Save(finalFilename, reply2, targetAgent)
if err != nil {
    log.Printf("[hub/pipe] save final result failed: %v", err)
    finalFilename = ""
}
```

#### 8. 返回格式化结果

```go
var result strings.Builder
result.WriteString(fmt.Sprintf("【第一步 %s 的回复】\n%s\n\n", sourceAgentName, reply1))
result.WriteString(fmt.Sprintf("【第二步 %s 的分析】\n%s", targetAgent, reply2))

result.WriteString(fmt.Sprintf("\n\n📁 Pipe 流程: %s → %s\n💾 已保存: %s, %s",
    sourceAgentName, targetAgent, filename, finalFilename))

return result.String()
```

---

## Hub API 依赖

Pipe 功能依赖 Hub 包提供以下 API：

### 1. `Save(filename, content, agentName string)`

保存内容到 Hub，自动添加 YAML frontmatter：
```yaml
---
agent: agent名称
timestamp: 2026-04-02T20:03:12+08:00
---

内容
```

### 2. `ReadSpecific(filenames []string) string`

读取指定文件并返回格式化的上下文字符串：
```
=== Agent Hub Shared Context ===

--- filename.md ---
文件内容

=== End Hub Context ===
```

### 3. `Exists(filename string) bool`

检查文件是否存在。

### 4. `ListWithInfo() ([]FileInfo, error)`

获取文件列表及修改时间，按最新优先排序（新增功能，支持 `/hub cat`）。

---

## 路由集成

**文件：** `messaging/handler.go:1003-1011`

```go
case strings.HasPrefix(rest, "pipe "):
    // /hub pipe <target_agent> <message>
    parts := strings.Fields(rest)
    if len(parts) < 3 {
        return "用法: /hub pipe <目标agent> <消息>\n" +
               "示例: /hub pipe gemini 分析量子计算对密码学的影响"
    }
    targetAgent := parts[1]
    message := strings.Join(parts[2:], " ")
    return h.handlePipe(ctx, client, msg, targetAgent, message, clientID)
```

---

## 关键设计决策

### 1. 会话隔离

每个用户使用独立的 `clientID`（基于 `FromUserID`），确保不同用户的 pipe 流程互不干扰：

```go
reply1, err := h.chatWithAgent(ctx, sourceAgent, msg.FromUserID,
    message, client, msg.ContextToken)
```

### 2. 降级处理

当 Hub 保存/读取失败时，不中断整个流程：

```go
if err != nil {
    log.Printf("[hub/pipe] save failed: %v", err)
    filename = ""  // 标记为空，后续跳过 Hub 读取
}
```

### 3. 文件命名规范

- 源文件：`pipe_时间戳_agent.md`
- 目标文件：`pipe_时间戳_target_final.md`
- 使用时间戳避免冲突，便于排序

### 4. 日志标记

所有 pipe 相关日志使用 `[hub/pipe]` 前缀，便于过滤调试：

```bash
journalctl -u weclaw | grep "\[hub/pipe\]"
```

---

## 使用示例

### 示例 1：翻译流程

```
/hub pipe gemini 你好世界
```

**流程：**
1. nanobot（默认）回复中文问候
2. gemini 接收中文内容，用英文重新表达

### 示例 2：代码审查流程

```
/hub pipe claude 解释一下 Go 的 interface 用法
```

**流程：**
1. nanobot 提供 Go interface 的中文解释
2. claude 基于解释提供英文补充和最佳实践

### 示例 3：多视角分析

```
/hub pipe gemini 量子计算对密码学的影响
```

**流程：**
1. nanobot 提供技术背景
2. gemini 从安全角度分析影响

---

## 多级 Pipe 传递

### 当前实现支持情况

**单条命令：** 支持两级传递（默认 Agent → 目标 Agent）

```
/hub pipe <目标agent> <消息>
```

**多级传递：** 可通过命令链或 Hub 文件实现

### 方法一：连续 Pipe 命令

实现 `nanobot → gemini → claude` 的三级传递：

```
# 第一步：nanobot → gemini
/hub pipe gemini 量子计算的基本原理

# 第二步：gemini 的结果已保存到 Hub，查看文件编号
/hub list

# 第三步：读取 gemini 的结果，发送给 claude
/hub pipe claude 请基于刚才 gemini 的分析，补充商业应用前景
```

**注意：** 第二次 pipe 时，claude 会接收到新的 nanobot 回复 + 你的提示。如需传递 gemini 的内容，需要先保存 gemini 的回复或使用方法二。

### 方法二：使用 Hub 文件编号（推荐）

更精确地控制传递链：

```
# 步骤 1: nanobot → gemini
/hub pipe gemini 分析量子计算的安全性

# 步骤 2: 查看保存的文件编号
/hub list
# 输出:
# 📁 Hub 文件列表 (最新优先):
#   [1] pipe_20260402-200312_gemini.md (04-02 20:03)
#   [2] pipe_20260402-200312_nanobot.md (04-02 20:03)

# 步骤 3: 读取 gemini 的结果并发送给 claude
# 方法 A: 直接读取文件内容（如果需要查看）
/hub cat 1

# 方法 B: 使用文件名发送给指定 agent（需要实现 /hub @agent <文件> 功能）
# 当前可用方式：复制内容后手动发送
```

### 方法三：手动构造传递链

对于复杂的 Agent 链，可以手动控制每一步：

```
# 步骤 1: 先获取第一个 agent 的回复（不使用 pipe）
@nanobot 解释什么是零知识证明

# 步骤 2: 将回复保存到 hub
/hub save zkp-explanation （然后粘贴内容）

# 步骤 3: 发送给第二个 agent
/hub zkp-explanation @gemini 分析其局限性

# 步骤 4: 保存第二个结果
/hub save zkp-limitation （粘贴内容）

# 步骤 5: 发送给第三个 agent
/hub zkp-limitation @claude 给出改进建议
```

### 多级 Pipe 实战示例

**场景：** 分析一项技术从原理到商业落地的完整路径

**Agent 链：** nanobot（技术） → gemini（安全） → claude（商业）

```
# 第一轮：技术原理
/hub pipe gemini 零知识证明的技术原理
# 保存: pipe_xxx_gemini.md

# 第二轮：安全分析（基于 gemini 的输出）
/hub list  # 记住 gemini 的文件编号，比如 [1]
/hub cat 1  # 查看内容

# 然后手动将内容发送给 claude
@claude 请基于以下 gemini 关于零知识证明的分析，从商业落地角度评估：
[粘贴 cat 1 的内容]

# 或者使用 hub 上下文（如果文件名为 zkp_analysis.md）
/hub zkp_analysis.md @claude 评估商业落地前景
```

### 多级 Pipe 流程图

```
用户消息
    │
    ▼
┌─────────────────────────────────────┐
│  /hub pipe gemini <消息>            │
│  nanobot → gemini                   │
└─────────────────────────────────────┘
    │
    ▼ 保存到 Hub
┌─────────────────────────────────────┐
│  pipe_xxx_gemini.md                 │
└─────────────────────────────────────┘
    │
    ▼ 读取或引用
┌─────────────────────────────────────┐
│  /hub cat 1                         │
│  或 /hub <文件> @claude <指令>      │
└─────────────────────────────────────┘
    │
    ▼
┌─────────────────────────────────────┐
│  claude 基于 gemini 输出继续分析    │
└─────────────────────────────────────┘
```

### 未来扩展：语法糖支持

未来可考虑添加更简洁的多级 pipe 语法：

```
# 提案 1: 链式语法
/hub pipe nanobot → gemini → claude 分析零知识证明

# 提案 2: 分号分隔
/hub pipe gemini 分析原理; /hub pipe claude 评估前景
```

**当前推荐：** 使用方法二（Hub 文件编号）进行精确控制。

---

## 扩展可能性

### 1. 语法糖支持

**当前：** 多级 Pipe 需要手动链式调用

**未来：** 支持更简洁的语法

```text
# 提案 1: 链式语法
/hub pipe nanobot → gemini → claude 分析 AI 安全

# 提案 2: 分号分隔
/hub pipe gemini 分析原理; /hub pipe claude 评估前景

# 提案 3: 引用语法
/hub pipe @1 @claude 评估前景  # @1 表示 Hub 中编号 1 的文件
```

### 2. 条件分支

根据源 Agent 输出内容选择不同目标：
```go
if contains(reply1, "安全相关") {
    targetAgent = "security-expert"
} else {
    targetAgent = "general-assistant"
}
```

### 3. 并行处理

同时发送给多个目标 Agent，汇总结果：
```go
// 并行调用多个 Agent
var wg sync.WaitGroup
results := make(chan string, len(targets))
// ...
```

---

## 故障排查

### 问题 1：Agent 名称显示为进程路径

**症状：**
```
📁 Pipe 流程: /home/nanobot/.nanobot/.venv/bin/python3 → pa
```

**原因：** 使用了 `sourceAgent.Info().Name`

**解决：** 改用 `h.defaultName`（已修复）

### 问题 2：Hub 文件未保存

**检查：**
```bash
# 远程查看 Hub 目录
ssh u "ls -la ~/.weclaw/hub/shared/"

# 查看日志
ssh u "journalctl -u weclaw | grep '\[hub/pipe\]'"
```

### 问题 3：第二步 Agent 无响应

**可能原因：**
- 目标 Agent 不可用
- Prompt 过长被截断
- 网络超时

**解决：** 检查 Agent 配置和日志

---

## 相关文件

| 文件 | 说明 |
|------|------|
| `messaging/handler.go` | Pipe 命令路由和核心实现 |
| `hub/hub.go` | Hub 存储和读取 API |
| `docs/agent-hub-design.md` | Hub 整体设计文档 |
| `docs/agent-hub-deploy.md` | 部署文档 |

---

## 版本历史

| 版本 | 日期 | 变更 |
|------|------|------|
| 1.0 | 2026-04-02 | 初始实现，支持两级 Pipe |
| 1.1 | 2026-04-02 | 修复 Agent 名称显示问题 |
| 1.2 | 2026-04-02 | 添加 `/hub cat` 编号访问功能 |
