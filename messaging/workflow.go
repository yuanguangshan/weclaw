package messaging

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fastclaw-ai/weclaw/ilink"
)

// WorkflowStep represents a single step in a workflow.
type WorkflowStep struct {
	Index    int               // 1-based step number
	Agent    string            // resolved agent name (for sequential steps)
	Message  string            // prompt template (may contain @N refs)
	SaveName string            // filename for hub.Save (empty = no save, no .md suffix)
	Parallel []*WorkflowBranch // if non-nil, this is a parallel step
}

// WorkflowBranch represents one branch of a parallel step.
type WorkflowBranch struct {
	Agent   string // resolved agent name
	Message string // prompt template (may contain @N refs)
}

const (
	maxWorkflowSteps    = 10
	maxParallelBranches = 5
	workflowStepTimeout = 3 * time.Minute
	// Save filename whitelist: alphanumeric, underscore, hyphen, dot, CJK Unified Ideographs (U+4E00–U+9FFF)
	saveNamePattern = `^[a-zA-Z0-9_\-.一-龟]{1,64}$`
)

// saveNameRe validates save filenames.
var saveNameRe = regexp.MustCompile(saveNamePattern)

// stepRe matches step lines: "step", "step1", "step:", "step1:", etc.
var stepRe = regexp.MustCompile(`(?i)^(step\d*\s|step\d*:)`)

// stepRefRe matches @N and @N.B references.
var stepRefRe = regexp.MustCompile(`@(\d+)\.(\d+)|@(\d+)`)

// workflowHelpText is returned when /workflow is called without arguments.
const workflowHelpText = `🔄 工作流 · 多步骤 Agent 编排

语法:
step1 @agent 消息内容
save 文件名

step2 parallel
branch @agent1 @1 消息
branch @agent2 @1 消息

step3 @agent 合并 @2 @3
save 最终结果

说明:
  step<N>   定义步骤（按顺序执行）
  parallel  步骤内并行执行
  branch    并行分支（每个分支发给不同 agent）
  save      保存步骤输出到 Hub
  @N        引用步骤 N 的输出
  @N.B      引用步骤 N 的第 B 个并行分支

示例:
/workflow
step1 @claude 分析这段代码
save code_analysis

step2 parallel
branch @gemini @1 找出安全漏洞
branch @qwen @1 写单元测试

step3 @claude 合并 @2.1 和 @2.2 的结果，输出最终报告
save security_report`

// handleWorkflow is the synchronous entry point for /workflow command.
func (h *Handler) handleWorkflow(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage, trimmed, clientID string) string {
	rest := strings.TrimSpace(strings.TrimPrefix(trimmed, "/workflow"))
	if rest == "" {
		return workflowHelpText
	}

	steps, parseErr := parseWorkflow(h, rest)
	if parseErr != "" {
		return "❌ DSL 解析错误:\n" + parseErr
	}
	if len(steps) == 0 {
		return "❌ 至少需要一个步骤"
	}

	// Validate all agents are available
	for _, s := range steps {
		for _, name := range stepAgents(s) {
			if _, err := h.getAgent(ctx, name); err != nil {
				return fmt.Sprintf("❌ Agent %q 不可用: %v", name, err)
			}
		}
	}

	// Use independent context with cancellation support.
	// The cancel func could be stored in a registry for future /workflow cancel support.
	wfCtx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		h.runWorkflow(wfCtx, client, msg, steps)
	}()

	return fmt.Sprintf("🔄 工作流已启动！共 %d 个步骤，结果将陆续发送...", len(steps))
}

// runWorkflow executes all steps sequentially/parallel and sends progress updates.
func (h *Handler) runWorkflow(ctx context.Context, client *ilink.Client, msg ilink.WeixinMessage, steps []*WorkflowStep) {
	// Protect against panics in agent calls
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[workflow] panic recovered: %v\n%s", r, debug.Stack())
			cid := NewClientID()
			plainText := MarkdownToPlainText(fmt.Sprintf("❌ 工作流异常终止: %v", r))
			req := &ilink.SendMessageRequest{
				Msg: ilink.SendMsg{
					FromUserID:   client.BotID(),
					ToUserID:     msg.FromUserID,
					ClientID:     cid,
					MessageType:  ilink.MessageTypeBot,
					MessageState: ilink.MessageStateFinish,
					ItemList: []ilink.MessageItem{
						{Type: ilink.ItemTypeText, TextItem: &ilink.TextItem{Text: plainText}},
					},
				},
			}
			client.SendMessage(context.Background(), req)
		}
	}()

	// sendMsg sends a message to the user. Uses a short-lived background context
	// so that cancellation/final-status messages can still be delivered after ctx is cancelled.
	sendMsg := func(text string) {
		cid := NewClientID()
		plainText := MarkdownToPlainText(text)
		req := &ilink.SendMessageRequest{
			Msg: ilink.SendMsg{
				FromUserID:   client.BotID(),
				ToUserID:     msg.FromUserID,
				ClientID:     cid,
				MessageType:  ilink.MessageTypeBot,
				MessageState: ilink.MessageStateFinish,
				ItemList: []ilink.MessageItem{
					{Type: ilink.ItemTypeText, TextItem: &ilink.TextItem{Text: plainText}},
				},
			},
		}
		sendCtx, sendCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer sendCancel()
		resp, err := client.SendMessage(sendCtx, req)
		if err != nil || resp.Ret != 0 {
			log.Printf("[workflow] send failed: err=%v ret=%d", err, resp.Ret)
		}
	}

	// Thread-safety invariant: stepResults and branchResults are only written by the main
	// goroutine (sequential loop in this function). Parallel branches receive read-only
	// snapshots via refSnap/branchSnap in runParallelStep. This invariant must be preserved
	// if the execution model changes (e.g. to DAG or nested parallelism).
	stepResults := make(map[int]string)     // step index -> combined content for @N
	branchResults := make(map[string]string) // "stepIdx.branchIdx" -> content for @N.B

	for _, step := range steps {
		sendMsg(fmt.Sprintf("⏳ 步骤 %d/%d 开始...", step.Index, len(steps)))

		if len(step.Parallel) > 0 {
			h.runParallelStep(ctx, msg.FromUserID, step, stepResults, branchResults, sendMsg)
		} else {
			h.runSequentialStep(ctx, msg.FromUserID, step, stepResults, branchResults, sendMsg)
		}

		// Brief pause between steps, responsive to cancellation
		if step.Index < len(steps) {
			select {
			case <-time.After(2 * time.Second):
			case <-ctx.Done():
				sendMsg("⚠️ 工作流已取消")
				return
			}
		}
	}

	sendMsg(buildWorkflowSummary(steps, stepResults))
}

// runSequentialStep executes a single sequential step.
func (h *Handler) runSequentialStep(ctx context.Context, userID string, step *WorkflowStep, stepResults map[int]string, branchResults map[string]string, sendMsg func(string)) {
	// Check cancellation before starting
	if ctx.Err() != nil {
		sendMsg(fmt.Sprintf("⚠️ 步骤 %d: 工作流已取消", step.Index))
		return
	}

	resolvedMsg := resolveStepRefs(step.Message, stepResults, branchResults)
	if resolvedMsg == "" {
		sendMsg(fmt.Sprintf("⚠️ 步骤 %d: 消息为空，跳过", step.Index))
		return
	}

	ag, err := h.getAgent(ctx, step.Agent)
	if err != nil {
		sendMsg(fmt.Sprintf("❌ 步骤 %d: Agent %q 不可用: %v", step.Index, step.Agent, err))
		return
	}

	convID := fmt.Sprintf("wf:%s:%s:step%d", userID, step.Agent, step.Index)
	stepCtx, cancel := context.WithTimeout(ctx, workflowStepTimeout)
	defer cancel()

	reply, chatErr := ag.Chat(stepCtx, convID, hubReplyHint+resolvedMsg)
	if chatErr != nil {
		sendMsg(fmt.Sprintf("❌ 步骤 %d (%s) 失败: %v", step.Index, step.Agent, chatErr))
		return
	}

	stepResults[step.Index] = reply

	stepSaved := false
	if step.SaveName != "" {
		savedName, saveErr := h.hub.Save(step.SaveName, reply, step.Agent)
		if saveErr != nil {
			sendMsg(fmt.Sprintf("⚠️ 步骤 %d 保存失败: %v", step.Index, saveErr))
		} else {
			sendMsg(fmt.Sprintf("✅ 步骤 %d (%s) 完成，已保存: %s", step.Index, step.Agent, savedName))
			stepSaved = true
		}
	}

	if !stepSaved {
		sendMsg(fmt.Sprintf("✅ 步骤 %d (%s) 完成\n%s", step.Index, step.Agent, truncate(reply, 300)))
	}
}

// runParallelStep executes a parallel step with multiple branches.
func (h *Handler) runParallelStep(ctx context.Context, userID string, step *WorkflowStep, stepResults map[int]string, branchResults map[string]string, sendMsg func(string)) {
	type branchOutput struct {
		idx     int
		agent   string
		content string
		err     error
	}

	ch := make(chan branchOutput, len(step.Parallel))

	// Snapshot reference data before spawning goroutines to guarantee read-only access
	refSnap := make(map[int]string, len(stepResults))
	for k, v := range stepResults {
		refSnap[k] = v
	}
	branchSnap := make(map[string]string, len(branchResults))
	for k, v := range branchResults {
		branchSnap[k] = v
	}

	for bi, branch := range step.Parallel {
		go func(bIdx int, br *WorkflowBranch) {
			// Check cancellation before starting
			if ctx.Err() != nil {
				ch <- branchOutput{idx: bIdx, agent: br.Agent, err: fmt.Errorf("工作流已取消")}
				return
			}

			resolvedMsg := resolveStepRefs(br.Message, refSnap, branchSnap)
			if resolvedMsg == "" {
				ch <- branchOutput{idx: bIdx, err: fmt.Errorf("消息为空")}
				return
			}

			ag, agErr := h.getAgent(ctx, br.Agent)
			if agErr != nil {
				ch <- branchOutput{idx: bIdx, agent: br.Agent, err: agErr}
				return
			}

			convID := fmt.Sprintf("wf:%s:%s:step%d.branch%d", userID, br.Agent, step.Index, bIdx)
			stepCtx, cancel := context.WithTimeout(ctx, workflowStepTimeout)
			defer cancel()

			reply, chatErr := ag.Chat(stepCtx, convID, hubReplyHint+resolvedMsg)
			if chatErr != nil {
				ch <- branchOutput{idx: bIdx, agent: br.Agent, err: chatErr}
				return
			}
			ch <- branchOutput{idx: bIdx, agent: br.Agent, content: reply}
		}(bi, branch)
	}

	// Collect all branch results and sort by branch index for stable output order
	outputs := make([]branchOutput, 0, len(step.Parallel))
	for range step.Parallel {
		outputs = append(outputs, <-ch)
	}
	sort.Slice(outputs, func(i, j int) bool {
		return outputs[i].idx < outputs[j].idx
	})

	var allContent []string
	var branchErrors []string
	for _, out := range outputs {
		if out.err != nil {
			branchErrors = append(branchErrors, fmt.Sprintf("分支 %d (%s): %v", out.idx+1, out.agent, out.err))
		} else {
			branchKey := fmt.Sprintf("%d.%d", step.Index, out.idx+1)
			branchResults[branchKey] = out.content
			allContent = append(allContent, fmt.Sprintf("[%s]\n%s", out.agent, out.content))
		}
	}

	combined := strings.Join(allContent, "\n\n---\n\n")
	stepResults[step.Index] = combined

	// Save if requested
	stepSaved := false
	if step.SaveName != "" && combined != "" {
		savedName, saveErr := h.hub.Save(step.SaveName, combined, "workflow-step"+fmt.Sprint(step.Index))
		if saveErr != nil {
			sendMsg(fmt.Sprintf("⚠️ 步骤 %d 保存失败: %v", step.Index, saveErr))
		} else {
			sendMsg(fmt.Sprintf("✅ 步骤 %d (parallel, %d/%d 成功)，已保存: %s", step.Index, len(allContent), len(step.Parallel), savedName))
			stepSaved = true
		}
	}

	// Report branch errors if any
	if len(branchErrors) > 0 {
		sendMsg("⚠️ 部分分支失败:\n" + strings.Join(branchErrors, "\n"))
	}

	// Only send success message if not already sent during save
	if !stepSaved {
		if len(branchErrors) > 0 {
			sendMsg(fmt.Sprintf("⚠️ 步骤 %d 完成 (%d/%d 成功)", step.Index, len(allContent), len(step.Parallel)))
		} else {
			sendMsg(fmt.Sprintf("✅ 步骤 %d (parallel, %d 分支全部成功)\n%s", step.Index, len(step.Parallel), truncate(combined, 300)))
		}
	}
}

// parseWorkflow parses the /workflow DSL text into ordered steps.
// Returns steps and an empty string on success, or an error message on failure.
func parseWorkflow(h *Handler, text string) ([]*WorkflowStep, string) {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil, "workflow 内容不能为空"
	}

	lines := strings.Split(text, "\n")
	var steps []*WorkflowStep
	var currentStep *WorkflowStep
	var inParallel bool

	finishStep := func() {
		if currentStep != nil {
			steps = append(steps, currentStep)
			currentStep = nil
		}
	}

	for _, rawLine := range lines {
		line := strings.TrimSpace(rawLine)
		if line == "" {
			continue
		}

		// Detect step start: "step1", "step2", "step" (case insensitive)
		if isStepLine(line) {
			finishStep()

			if len(steps) >= maxWorkflowSteps {
				return nil, fmt.Sprintf("步骤数超过上限 (%d)", maxWorkflowSteps)
			}

			rest := extractStepRest(line)
			inParallel = false

			// Check for "parallel" keyword after step marker
			rest = strings.TrimSpace(rest)
			if strings.EqualFold(rest, "parallel") {
				currentStep = &WorkflowStep{
					Index:    len(steps) + 1,
					Parallel: []*WorkflowBranch{},
				}
				inParallel = true
				continue
			}

			// Parse @agent from rest
			agentName, message := parseAgentLine(h, rest)
			currentStep = &WorkflowStep{
				Index:   len(steps) + 1,
				Agent:   agentName,
				Message: message,
			}
			continue
		}

		// Detect branch line (only inside parallel step)
		if strings.HasPrefix(strings.ToLower(line), "branch") {
			if currentStep == nil || !inParallel {
				return nil, "branch 只能在 parallel 步骤内使用"
			}
			if len(currentStep.Parallel) >= maxParallelBranches {
				return nil, fmt.Sprintf("并行分支数超过上限 (%d)", maxParallelBranches)
			}
			rest := strings.TrimSpace(line[len("branch"):])
			agentName, message := parseAgentLine(h, rest)
			currentStep.Parallel = append(currentStep.Parallel, &WorkflowBranch{
				Agent:   agentName,
				Message: message,
			})
			continue
		}

		// Detect save directive
		if strings.HasPrefix(strings.ToLower(line), "save ") || strings.EqualFold(line, "save") {
			if currentStep == nil {
				return nil, "save 必须在 step 内使用"
			}
			saveName := strings.TrimSpace(strings.TrimPrefix(strings.ToLower(line), "save "))
			saveName = strings.TrimSpace(strings.TrimPrefix(saveName, "save")) // handle "save" alone
			// Strip .md suffix (hub.Save adds it)
			saveName = strings.TrimSuffix(saveName, ".md")
			if saveName == "" {
				return nil, "save 后必须指定文件名"
			}
			// Validate filename: whitelist alphanumeric, underscore, hyphen, dot, CJK
			if !saveNameRe.MatchString(saveName) {
				return nil, "文件名只能包含字母、数字、下划线、连字符、点和中文，长度 1-64"
			}
			currentStep.SaveName = saveName
			continue
		}

		// Content line: append to current step's message
		if currentStep != nil && !inParallel {
			if currentStep.Message != "" {
				currentStep.Message += "\n" + line
			} else {
				currentStep.Message = line
			}
		}
	}

	finishStep()

	// Validate steps
	for _, s := range steps {
		if len(s.Parallel) == 0 && s.Agent == "" {
			return nil, fmt.Sprintf("步骤 %d 缺少 @agent 指定", s.Index)
		}
		if len(s.Parallel) == 0 && s.Message == "" {
			return nil, fmt.Sprintf("步骤 %d 消息为空", s.Index)
		}
		if len(s.Parallel) > 0 {
			for j, b := range s.Parallel {
				if b.Agent == "" {
					return nil, fmt.Sprintf("步骤 %d 分支 %d 缺少 @agent", s.Index, j+1)
				}
				if b.Message == "" {
					return nil, fmt.Sprintf("步骤 %d 分支 %d 消息为空", s.Index, j+1)
				}
			}
		}
	}

	return steps, ""
}

func isStepLine(line string) bool {
	lower := strings.ToLower(line)
	if lower == "step" {
		return true
	}
	return stepRe.MatchString(lower)
}

// extractStepRest returns the text after the step marker (e.g. "step1 @claude msg" -> "@claude msg").
func extractStepRest(line string) string {
	// Find end of "step" + optional digits
	lower := strings.ToLower(line)
	i := len("step")
	for i < len(lower) && lower[i] >= '0' && lower[i] <= '9' {
		i++
	}
	// Skip optional separator (space, colon, dot)
	if i < len(lower) && (lower[i] == ':' || lower[i] == '.' || lower[i] == ' ') {
		i++
	}
	return strings.TrimSpace(line[i:])
}

// parseAgentLine parses "@agent message" from a line, using the handler's parseCommand.
func parseAgentLine(h *Handler, text string) (agentName, message string) {
	if h == nil || text == "" {
		return "", text
	}
	names, rest := h.parseCommand(text)
	if len(names) > 0 {
		return names[0], strings.TrimSpace(rest)
	}
	return "", text
}

// resolveStepRefs replaces @N and @N.B references with actual step output content.
func resolveStepRefs(message string, stepResults map[int]string, branchResults map[string]string) string {
	if message == "" {
		return ""
	}

	return stepRefRe.ReplaceAllStringFunc(message, func(match string) string {
		sub := stepRefRe.FindStringSubmatch(match)
		if sub[2] != "" {
			// @N.B pattern
			branchKey := sub[1] + "." + sub[2]
			if content, ok := branchResults[branchKey]; ok {
				return content
			}
			return match // leave as-is if not found
		}
		if sub[3] != "" {
			// @N pattern
			stepNum, _ := strconv.Atoi(sub[3])
			if content, ok := stepResults[stepNum]; ok {
				return content
			}
			return match
		}
		return match
	})
}

// stepAgents returns all agent names referenced in a step.
func stepAgents(s *WorkflowStep) []string {
	if len(s.Parallel) > 0 {
		names := make([]string, len(s.Parallel))
		for i, b := range s.Parallel {
			names[i] = b.Agent
		}
		return names
	}
	return []string{s.Agent}
}

// buildWorkflowSummary creates a final summary of the workflow execution.
func buildWorkflowSummary(steps []*WorkflowStep, stepResults map[int]string) string {
	var sb strings.Builder
	sb.WriteString("🔄 工作流执行完毕！\n\n")
	for _, step := range steps {
		content, ok := stepResults[step.Index]
		if !ok {
			sb.WriteString(fmt.Sprintf("  %d. ❌ 未完成\n", step.Index))
			continue
		}
		mode := "顺序"
		if len(step.Parallel) > 0 {
			mode = fmt.Sprintf("并行(%d分支)", len(step.Parallel))
		}
		saved := ""
		if step.SaveName != "" {
			saved = " (已保存)"
		}
		sb.WriteString(fmt.Sprintf("  %d. ✅ %s %s%s\n", step.Index, mode, truncate(content, 50), saved))
	}
	sb.WriteString("\n💡 使用 /hub list 查看保存的文件")
	return sb.String()
}
