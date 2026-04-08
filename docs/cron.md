# Cron 定时任务功能文档

WeClaw 的 Cron 定时任务功能允许用户通过自然语言或标准 cron 表达式创建、管理和执行定时任务。

---

## 📋 功能概述

- ✅ **自然语言解析** - 使用 AI 将自然语言转换为 cron 表达式
- ✅ **规则解析** - 内置规则解析器支持常见时间表达
- ✅ **多种任务类型** - 支持文本消息、Workflow DSL、Agent 调用
- ✅ **任务管理** - 通过管理后台或命令行管理任务
- ✅ **持久化存储** - 任务保存在 `~/.weclaw/cron_jobs.json`
- ✅ **启用/禁用** - 可随时启用或禁用任务
- ✅ **API 接口** - RESTful API 支持集成开发

---

## 📁 文件结构

| 文件 | 行数 | 作用 |
|------|------|------|
| `messaging/cron.go` | ~382 | Cron 核心实现（调度、存储、执行） |
| `api/server.go` | ~268 | Cron API 端点（增删改查） |
| `web/admin.html` | ~114 | Cron 管理界面 |
| `messaging/handler.go` | ~200 | 自然语言命令解析 |

---

## 🎯 核心数据结构

```go
// CronJob 定时任务
type CronJob struct {
    ID        string      // 任务 ID (格式: cron_<timestamp>)
    UserID    string      // 用户 ID
    CronExpr  string      // Cron 表达式
    Command   CronCommand // 执行命令
    Enabled   bool        // 是否启用
    CreatedAt int64       // 创建时间
    NextRun   int64       // 下次运行时间
}

// CronCommand 执行命令
type CronCommand struct {
    Type    string // "text" | "workflow" | "agent"
    Content string // 消息文本或 DSL
    Agent   string // 指定 Agent (仅 type=agent 时)
}
```

---

## 🚀 使用方式

### 1️⃣ 自然语言创建任务（推荐）

通过 AI 智能解析自然语言：

```
每天早上9点提醒我开会
每周一到周五早上9点开例会
每周四凌晨2:30提醒睡觉
每30分钟检查状态
每天早上9点生成周报
```

**示例响应：**
```
✅ 定时任务已添加（AI 解析）

ID: cron_1775672695739713860
表达式: 0 30 2 * * 4
类型: text
内容: 提醒睡觉
```

### 2️⃣ 标准 Cron 表达式

使用标准的 6 段式 cron 表达式（包含秒）：

```
/cron add "0 9 * * *" 每天早上9点的提醒
/cron add "0 0 9 * * 1-5" 工作日早上9点
/cron add "*/30 * * * *" 每30分钟
```

**Cron 表达式格式：**
```
┌───────────── 秒 (0-59)
│ ┌─────────── 分 (0-59)
│ │ ┌───────── 时 (0-23)
│ │ │ ┌─────── 日 (1-31)
│ │ │ │ ┌───── 月 (1-12)
│ │ │ │ │ ┌─── 星期 (0-6, 0=周日)
│ │ │ │ │ │
* * * * * *
```

### 3️⃣ Workflow 任务

创建执行 Workflow DSL 的定时任务：

```
每天早上9点
step: @claude 生成今日工作计划
step: @gemini 根据计划生成待办清单
save: daily_plan
```

### 4️⃣ 指定 Agent 任务

```
/cron add "0 8 * * *" @deepseek 早报摘要
```

---

## 🎨 管理后台

访问 `http://localhost:18011/admin` 进入管理后台，点击 **Cron** 菜单：

**功能按钮：**
- **Add Job** - 添加新任务
- **Refresh** - 刷新列表
- **Enable/Disable** - 启用/禁用任务
- **Delete** - 删除任务

**任务列表显示：**
| 列 | 说明 |
|------|------|
| ID | 任务 ID（前 20 字符） |
| Cron Expr | Cron 表达式 |
| Type | 任务类型（Text/Workflow/Agent） |
| Content | 任务内容（截断显示） |
| User | 创建用户（前 15 字符） |
| Status | 状态（Enabled/Disabled） |
| Actions | 操作按钮 |

---

## 📡 API 接口

### 列出任务
```bash
GET /api/cron
```

**响应：**
```json
{
  "jobs": [
    {
      "id": "cron_1775672695739713860",
      "user_id": "o9cq80wpGQpRIUxH2LGdGFrksGak@im.wechat",
      "cron_expr": "0 30 2 * * 4",
      "command": {
        "type": "text",
        "content": "提醒睡觉"
      },
      "enabled": true,
      "created_at": 1775672695,
      "next_run": 1775672755
    }
  ]
}
```

### 添加任务
```bash
POST /api/cron
Content-Type: application/json

{
  "cron_expr": "0 9 * * *",
  "type": "text",
  "content": "早上好"
}
```

### 删除任务
```bash
DELETE /api/cron/{id}
```

### 启用任务
```bash
PUT /api/cron/{id}/enable
```

### 禁用任务
```bash
PUT /api/cron/{id}/disable
```

---

## 🧠 AI 解析流程

```
用户输入 → AI 解析 → JSON 验证 → 创建任务
                ↓
         解析失败时降级到规则解析
```

**AI Prompt 模板：**
```go
你是时间表达式解析专家。请将自然语言时间描述转换为 cron 表达式。

规则：
1. 返回标准 6 段式 cron 表达式（包含秒）
2. 时间范围：分 0-59，时 0-23
3. 返回纯 JSON 格式

输入格式：
- "每天早上9点" → 每天上午9点
- "每周一早上8点" → 每周一上午8点
- "每30分钟" → 每30分钟一次

输出格式：
{
  "cron_expr": "0 0 9 * * *",
  "message": "提取的消息内容",
  "type": "text"
}
```

**降级策略：**
- AI 解析失败 → 规则解析
- JSON 无效 → 规则解析
- Cron 表达式无效 → 规则解析
- 类型无效 → 规则解析

---

## 🔧 规则解析器

内置规则解析支持以下模式：

| 模式 | 示例 | Cron 表达式 |
|------|------|-------------|
| 每天 X 点 | 每天9点 | `0 0 9 * * *` |
| 每 N 分钟 | 每30分钟 | `0 */30 * * * *` |
| 每周 X | 每周一 | `0 0 0 * * 1` |
| 每周 X Y 点 | 每周一早上8点 | `0 0 8 * * 1` |
| 工作日 | 每周一到周五 | `0 0 9 * * 1-5` |

---

## 📦 存储位置

任务数据存储在：`~/.weclaw/cron_jobs.json`

**文件格式：**
```json
{
  "jobs": [
    {
      "id": "cron_1775672695739713860",
      "user_id": "o9cq80wpGQpRIUxH2LGdGFrksGak@im.wechat",
      "cron_expr": "0 30 2 * * 4",
      "command": {
        "type": "text",
        "content": "提醒睡觉"
      },
      "enabled": true,
      "created_at": 1775672695,
      "next_run": 1775672755
    }
  ]
}
```

---

## 🔄 执行流程

```
Cron 调度器触发
       ↓
根据任务类型分发
       ↓
┌──────┴──────┬─────────────┐
Text          Workflow      Agent
发送消息      执行 DSL      调用 Agent
```

**执行详情：**

1. **Text 类型**
   - 直接发送文本消息到用户微信

2. **Workflow 类型**
   - 解析 Workflow DSL
   - 依次执行各个步骤
   - 返回执行结果

3. **Agent 类型**
   - 指定 Agent 处理任务
   - 返回 Agent 响应

---

## 🛠️ 技术实现

### 调度器
使用 [robfig/cron](https://github.com/robfig/cron) 库：
- 支持 6 段式 cron 表达式（包含秒）
- 使用本地时区
- 秒级精度调度

### 并发安全
```go
type CronManager struct {
    cron    *cron.Cron
    jobs    map[string]cron.EntryID
    store   *CronStore
    mu      sync.RWMutex  // 读写锁保护并发访问
}
```

### 原子写入
文件操作使用临时文件+重命名确保原子性：
```go
tmpPath := s.filePath + ".tmp"
os.WriteFile(tmpPath, jsonData, 0600)
os.Rename(tmpPath, s.filePath)
```

---

## 📝 日志示例

```bash
[cron] started with 3 jobs
[cron] executing job cron_1775672695739713860 for user o9cq80wpGQpRIUxH2LGdGFrksGak@im.wechat
[cron] AI parsing successful, creating job
[cron] sent text message to o9cq80wpGQpRIUxH2LGdGFrksGak@im.wechat
```

---

## ⚠️ 注意事项

1. **时区** - 所有时间使用服务器本地时区
2. **幂等性** - 相同任务会创建不同 ID
3. **持久化** - 服务重启后自动加载已启用任务
4. **执行时长** - 单个任务执行不应超过 1 分钟
5. **并发限制** - 同一任务可能并发执行（如果前一次未完成）

---

## 🔮 未来计划

- [ ] 支持任务执行历史记录
- [ ] 支持任务执行超时控制
- [ ] 支持任务失败重试
- [ ] 支持任务依赖关系
- [ ] 支持更多自然语言时间表达
- [ ] Webhook 回调支持
