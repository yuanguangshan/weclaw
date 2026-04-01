# Agent Hub 部署指南

> 将本地已开发的 Agent Hub 功能（`/hub`、`/save` 指令）部署到 nanobot 服务器。

## 前置确认

本地代码已 commit：`f7ea503`（包含 hub.go + handler.go 改动）。
所有 31 个测试通过。

---

## 部署步骤（在 nanobot 服务器上执行）

### Step 1：SSH 登录服务器

```bash
ssh root@nanobot
# 或你的实际服务器地址
```

### Step 2：确认 weclaw 安装方式

先确认 weclaw 当前是怎么运行的：

```bash
# 查看是否通过 Docker 运行
docker ps | grep weclaw

# 查看是否直接安装的二进制
which weclaw && weclaw version

# 查看是否有源码目录
ls -la /root/.weclaw/workspace/weclaw/ 2>/dev/null
```

根据结果选择 **方案 A**（Docker）或 **方案 B**（源码编译）。

---

### 方案 A：Docker 部署（如果 weclaw 跑在 Docker 里）

```bash
# 1. 进入源码目录
cd /root/.weclaw/workspace/weclaw

# 2. 拉取最新代码（从你的 fork）
git remote -v  # 确认 origin 指向 yuanguangshan/weclaw
git pull origin main

# 3. 重新构建 Docker 镜像
docker build -t weclaw:latest .

# 4. 停止旧容器
docker stop weclaw 2>/dev/null || docker stop $(docker ps -q --filter ancestor=weclaw:latest) 2>/dev/null

# 5. 启动新容器（参考你原来的启动方式，示例）
docker run -d \
  --name weclaw \
  --restart unless-stopped \
  -v /root/.weclaw:/root/.weclaw \
  weclaw:latest

# 6. 验证
docker logs -f weclaw  # 应该能看到正常启动日志
```

---

### 方案 B：源码编译（如果 weclaw 是二进制安装的）

```bash
# 1. 进入源码目录
cd /root/.weclaw/workspace/weclaw

# 2. 确认远程仓库
git remote -v
# 如果 origin 不是 yuanguangshan/weclaw，需要添加：
# git remote add myfork git@github.com:yuanguangshan/weclaw.git

# 3. 拉取最新代码
git pull origin main

# 4. 编译
go build -o /usr/local/bin/weclaw .

# 5. 重启 weclaw
weclaw stop
weclaw start

# 6. 验证
weclaw status
tail -f /root/.weclaw/weclaw.log
```

---

### 方案 C：直接二进制更新（如果服务器上没有源码）

```bash
# 1. 确认架构
uname -m   # 应该输出 x86_64 或 aarch64

# 2. 在本地 Mac 上交叉编译 Linux 二进制
# （回到本地 Mac 执行，见下文"本地交叉编译"章节）

# 3. 上传到服务器
scp weclaw_linux_amd64 root@nanobot:/usr/local/bin/weclaw
ssh root@nanobot "chmod +x /usr/local/bin/weclaw"

# 4. 重启
ssh root@nanobot "weclaw stop && weclaw start"
```

---

## 本地交叉编译（从 Mac 编译 Linux 二进制）

在本地 Mac 上执行：

```bash
cd /Users/ygs/ygs/weclaw

# 编译 Linux amd64
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o weclaw_linux_amd64 .

# 如果服务器是 ARM（如树莓派/鲲鹏）
# GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-s -w" -o weclaw_linux_arm64 .

# 验证二进制
file weclaw_linux_amd64
# 应该输出: weclaw_linux_amd64: ELF 64-bit LSB executable, x86-64, ...

# 上传到服务器
scp weclaw_linux_amd64 root@nanobot:/usr/local/bin/weclaw
```

---

## 验证部署成功

部署完成后，在微信中发送：

```
/help
```

应该看到新增的两行：

```
Agent Hub (cross-agent collaboration):
/hub - List shared context files
/hub {msg} - Read all shared files, inject context, send to agent
/save {file} {msg} - Send to agent, save reply to hub
```

然后测试完整流程：

```
# 1. 保存 Claude 的分析
/save round1_claude.md 从哲学角度分析AI代理是否会替代人类决策

# 2. 确认 Hub 中有文件
/hub ls

# 3. 让另一个 agent 读取 Hub 上下文并反驳
@codex /hub round1_claude.md 从技术可行性角度反驳上面的观点

# 4. 综合两方
@claude /hub 综合以上两方观点，给出你的最终判断
```

---

## Hub 文件存储位置

所有共享文件存储在服务器上的：

```
~/.weclaw/hub/shared/
```

每个文件是 Markdown 格式，带 YAML frontmatter：

```yaml
---
agent: claude
timestamp: "2026-04-02T01:30:00Z"
session: wechat_user_id
---

agent 回复的具体内容...
```

手动查看/清理：

```bash
ls -la ~/.weclaw/hub/shared/
cat ~/.weclaw/hub/shared/round1_claude.md
```

---

## 回滚方案

如果部署后出问题：

```bash
# 回退到上一个版本
cd /root/.weclaw/workspace/weclaw
git log --oneline -5
git checkout <上一个commit>

# 重新编译/部署
go build -o /usr/local/bin/weclaw .
weclaw stop && weclaw start
```

---

## 涉及改动的文件清单

| 文件 | 状态 | 说明 |
|------|------|------|
| `hub/hub.go` | **新增** | Hub 核心包（文件系统共享黑板） |
| `hub/hub_test.go` | **新增** | 12 个单元测试 |
| `messaging/handler.go` | **修改** | 集成 `/hub`、`/save` 指令 + help 更新 |
| `docs/agent-hub-design.md` | **新增** | 架构设计文档 |

---

## 注意事项

1. **Hub 存储目录会自动创建**：首次使用 `/save` 时，`~/.weclaw/hub/shared/` 会自动创建
2. **不会影响现有功能**：`/hub` 和 `/save` 是新增命令，不影响原有的 @agent、/new、/cwd 等
3. **跨 agent 协作不会污染会话**：通过 `/hub` 发送的消息使用 `hub:agentName:userID` 作为 conversationID，和正常对话隔离
