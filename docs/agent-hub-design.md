# Agent Hub Design — 跨 Agent 上下文共享

## Problem

WeClaw supports multiple AI agents (Claude, Codex, Gemini, etc.), each with isolated conversation sessions. Users cannot easily share context between agents for collaborative workflows like:
- Multi-agent debate
- Chain-of-thought analysis (Agent A outputs → Agent B reviews)
- Collaborative content creation

## Solution: Agent Hub

A shared context layer built directly into WeClaw's Go codebase, using the filesystem as a persistent message board.

### Architecture

```
WeChat ←→ WeClaw (handler.go)
              │
              ├── Agent A (isolated session)
              ├── Agent B (isolated session)
              └── Agent Hub (shared context)
                    │
                    ├── ~/.weclaw/hub/shared/    # shared context files
                    ├── ~/.weclaw/hub/templates/  # prompt templates
                    └── /hub, /save, /hub pipe commands
```

### New Commands

| Command | Description | Example |
|---------|-------------|---------|
| `/hub` | Read all shared files and inject as context | `/hub 基于以上分析，给出你的反驳` |
| `/hub {filename}` | Read specific file from shared | `/hub round1_claude.md 基于此反驳` |
| `/save {filename} {message}` | Send message and save reply to shared | `/save round1.md 分析AI未来` |
| `/hub ls` | List files with numbers (newest first) | `/hub ls` |
| `/hub cat {编号}` | View file content by number | `/hub cat 1` |
| `/hub clear` | Clear all shared files | `/hub clear` |
| `/hub pipe {agent} {msg}` | Chain: default agent → target agent | `/hub pipe gemini 量子计算` |
| `/hub pipe {agent} @1 {msg}` | Chain using Hub file reference | `/hub pipe claude @1 继续分析` |
| `/hub pipe {agent} @-1 {msg}` | Chain using latest file | `/hub pipe claude @-1 补充说明` |

### Workflow Examples

#### Multi-Agent Debate (with Pipe)
```
1. /hub pipe claude 从哲学角度分析AI是否会替代人类
   → nanobot replies, saved as [@1] pipe_xxx_nanobot.md
   → claude reads it, replies with philosophical analysis
   → saved as [@2] pipe_xxx_claude_final.md

2. /hub pipe gemini @2 从技术角度反驳
   → gemini reads claude's analysis (@2)
   → replies with technical rebuttal
   → saved as [@3] pipe_xxx_gemini_final.md

3. /hub pipe deepseek @3 总结双方观点
   → deepseek synthesizes both perspectives
```

#### Chain Collaboration (with reference syntax)
```
# 方法一：使用相对编号 @-1（最新文件）
/hub pipe gemini 写一个量子计算博客大纲
/hub pipe claude @-1 基于大纲扩写完整文章
/hub pipe deepseek @-1 审查文章质量并优化

# 方法二：使用绝对编号
/hub pipe gemini 写一个量子计算博客大纲    # 结果保存为 @1
/hub pipe claude @1 基于大纲扩写完整文章    # 结果保存为 @2
/hub pipe deepseek @2 审查文章质量并优化    # 结果保存为 @3

# 方法三：使用文件名引用
/hub pipe gemini 写一个量子计算博客大纲
/hub pipe claude @pipe_xxx_gemini.md 继续扩写
```

### Implementation Status

#### ✅ Phase 1: File-based shared context (COMPLETE)
- `~/.weclaw/hub/shared/` — shared context files
- `~/.weclaw/hub/templates/` — prompt templates
- Commands in `handler.go`: `/hub`, `/save`, `/hub ls`, `/hub cat`, `/hub clear`

#### ✅ Phase 2: Auto-save with context injection (COMPLETE)
- Agent replies auto-saved when `/save` is used
- `/hub` auto-injects shared files as system prompt prefix
- File naming with timestamp and agent name

#### ✅ Phase 3: Chain mode with reference syntax (COMPLETE)
- `/hub pipe {agent}` — automatic chain: send → save → next
- `@1`, `@2` — absolute file number references
- `@-1`, `@-2` — relative references (latest, second latest)
- `@filename.md` — direct filename reference
- Auto-display file numbers in results for easy continuation
- **Thread-safe operations** with `sync.RWMutex` protection

### File Format

Shared files use Markdown with YAML frontmatter:

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

### Integration Points in weclaw

1. **`messaging/handler.go`** — Command parsing for `/hub`, `/save`, `/hub pipe`
2. **`hub/hub.go`** — Hub logic: read/write shared files, inject context (with concurrency protection)
3. **`cmd/start.go`** — Initialize Hub with default directory

### Key Design Decisions

1. **Filesystem over database** — Simple, inspectable, no extra dependencies
2. **Markdown with frontmatter** — Human-readable, agent-friendly, extensible
3. **Opt-in via commands** — No automatic cross-contamination of agent sessions
4. **Go-native** — No Python dependencies, fits weclaw's architecture
5. **Reference syntax** — `@1`, `@-1`, `@filename.md` for flexible file referencing
6. **Thread-safe** — `sync.RWMutex` protects all file operations in concurrent scenarios
