package messaging

import (
	"testing"
)

func TestParseWorkflow_SimpleSequential(t *testing.T) {
	h := newTestHandler()
	dsl := `step1 @claude 分析这段代码
save analysis

step2 @gemini 给出优化建议`

	steps, err := parseWorkflow(h, dsl)
	if err != "" {
		t.Fatalf("unexpected error: %s", err)
	}
	if len(steps) != 2 {
		t.Fatalf("expected 2 steps, got %d", len(steps))
	}

	// Step 1
	if steps[0].Agent != "claude" {
		t.Errorf("step1 agent = %q, want claude", steps[0].Agent)
	}
	if steps[0].Message != "分析这段代码" {
		t.Errorf("step1 message = %q, want '分析这段代码'", steps[0].Message)
	}
	if steps[0].SaveName != "analysis" {
		t.Errorf("step1 saveName = %q, want analysis", steps[0].SaveName)
	}

	// Step 2
	if steps[1].Agent != "gemini" {
		t.Errorf("step2 agent = %q, want gemini", steps[1].Agent)
	}
	if steps[1].SaveName != "" {
		t.Errorf("step2 saveName = %q, want empty", steps[1].SaveName)
	}
}

func TestParseWorkflow_Parallel(t *testing.T) {
	h := newTestHandler()
	dsl := `step1 @claude 分析日志
save analysis

step2 parallel
branch @gemini @1 找出漏洞
branch @qwen @1 写测试

step3 @claude 合并 @2.1 和 @2.2`

	steps, err := parseWorkflow(h, dsl)
	if err != "" {
		t.Fatalf("unexpected error: %s", err)
	}
	if len(steps) != 3 {
		t.Fatalf("expected 3 steps, got %d", len(steps))
	}

	// Step 2 should be parallel
	if len(steps[1].Parallel) != 2 {
		t.Fatalf("step2 parallel branches = %d, want 2", len(steps[1].Parallel))
	}
	if steps[1].Parallel[0].Agent != "gemini" {
		t.Errorf("step2 branch1 agent = %q, want gemini", steps[1].Parallel[0].Agent)
	}
	if steps[1].Parallel[1].Agent != "qwen" {
		t.Errorf("step2 branch2 agent = %q, want qwen", steps[1].Parallel[1].Agent)
	}
}

func TestParseWorkflow_Empty(t *testing.T) {
	h := newTestHandler()
	_, err := parseWorkflow(h, "")
	if err == "" {
		t.Fatal("expected error for empty workflow")
	}
}

func TestParseWorkflow_MissingAgent(t *testing.T) {
	h := newTestHandler()
	_, err := parseWorkflow(h, "step1 分析这段代码")
	if err == "" {
		t.Fatal("expected error for step without @agent")
	}
}

func TestParseWorkflow_BranchOutsideParallel(t *testing.T) {
	h := newTestHandler()
	_, err := parseWorkflow(h, `step1 @claude hello
branch @gemini world`)
	if err == "" {
		t.Fatal("expected error for branch outside parallel")
	}
}

func TestParseWorkflow_SaveWithMdSuffix(t *testing.T) {
	h := newTestHandler()
	steps, err := parseWorkflow(h, `step1 @claude hello
save report.md`)
	if err != "" {
		t.Fatalf("unexpected error: %s", err)
	}
	if steps[0].SaveName != "report" {
		t.Errorf("saveName = %q, want 'report' (stripped .md)", steps[0].SaveName)
	}
}

func TestParseWorkflow_SaveInvalidFilename(t *testing.T) {
	h := newTestHandler()
	tests := []struct {
		name     string
		filename string
	}{
		{"path separator", "../etc/passwd"},
		{"special chars", "file@#$%"},
		{"too long", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}, // 65 chars
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parseWorkflow(h, "step1 @claude hello\nsave "+tt.filename)
			if err == "" {
				t.Errorf("expected error for filename %q", tt.filename)
			}
		})
	}
}

func TestParseWorkflow_SaveValidFilenames(t *testing.T) {
	h := newTestHandler()
	tests := []struct {
		name     string
		filename string
	}{
		{"simple", "report"},
		{"with hyphen", "my-report"},
		{"with underscore", "my_report"},
		{"with dot", "report.v2"},
		{"chinese", "分析报告"},
		{"mixed", "report-2026_分析"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			steps, err := parseWorkflow(h, "step1 @claude hello\nsave "+tt.filename)
			if err != "" {
				t.Fatalf("unexpected error for filename %q: %s", tt.filename, err)
			}
			if steps[0].SaveName != tt.filename {
				t.Errorf("saveName = %q, want %q", steps[0].SaveName, tt.filename)
			}
		})
	}
}

func TestParseWorkflow_MultilineMessage(t *testing.T) {
	h := newTestHandler()
	dsl := `step1 @claude 分析下面这段代码
func main() {
    fmt.Println("hello")
}
save code_analysis`

	steps, err := parseWorkflow(h, dsl)
	if err != "" {
		t.Fatalf("unexpected error: %s", err)
	}
	if steps[0].Agent != "claude" {
		t.Errorf("agent = %q, want claude", steps[0].Agent)
	}
	if steps[0].SaveName != "code_analysis" {
		t.Errorf("saveName = %q, want code_analysis", steps[0].SaveName)
	}
	// Message should contain the code lines
	msg := steps[0].Message
	if len(msg) < 10 {
		t.Errorf("message too short: %q", msg)
	}
}

func TestParseWorkflow_MaxSteps(t *testing.T) {
	h := newTestHandler()
	var dsl string
	for i := 1; i <= 11; i++ {
		dsl += "step @claude hello\n"
	}
	_, err := parseWorkflow(h, dsl)
	if err == "" {
		t.Fatal("expected error for too many steps")
	}
}

func TestParseWorkflow_MaxBranches(t *testing.T) {
	h := newTestHandler()
	dsl := "step1 parallel\n"
	for i := 0; i < 6; i++ {
		dsl += "branch @claude hello\n"
	}
	_, err := parseWorkflow(h, dsl)
	if err == "" {
		t.Fatal("expected error for too many branches")
	}
}

func TestResolveStepRefs_Simple(t *testing.T) {
	stepResults := map[int]string{
		1: "output from step 1",
		2: "output from step 2",
	}
	branchResults := map[string]string{}

	result := resolveStepRefs("分析 @1 然后总结 @2", stepResults, branchResults)
	expected := "分析 output from step 1 然后总结 output from step 2"
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestResolveStepRefs_Branch(t *testing.T) {
	stepResults := map[int]string{}
	branchResults := map[string]string{
		"2.1": "branch 1 output",
		"2.2": "branch 2 output",
	}

	result := resolveStepRefs("合并 @2.1 和 @2.2", stepResults, branchResults)
	expected := "合并 branch 1 output 和 branch 2 output"
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestResolveStepRefs_MissingRef(t *testing.T) {
	stepResults := map[int]string{1: "output1"}
	branchResults := map[string]string{}

	result := resolveStepRefs("引用 @1 和 @99", stepResults, branchResults)
	expected := "引用 output1 和 @99"
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestResolveStepRefs_Empty(t *testing.T) {
	result := resolveStepRefs("", nil, nil)
	if result != "" {
		t.Errorf("expected empty, got %q", result)
	}
}

func TestIsStepLine(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"step1 @claude hello", true},
		{"Step2 @gemini world", true},
		{"step @claude test", true},
		{"step: @claude test", true},
		{"notastep hello", false},
		{"@claude hello", false},
		{"parallel", false},
	}

	for _, tt := range tests {
		got := isStepLine(tt.input)
		if got != tt.want {
			t.Errorf("isStepLine(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestExtractStepRest(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"step1 @claude hello", "@claude hello"},
		{"step2 parallel", "parallel"},
		{"step: @gemini world", "@gemini world"},
	}

	for _, tt := range tests {
		got := extractStepRest(tt.input)
		if got != tt.want {
			t.Errorf("extractStepRest(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestParseWorkflow_NoIndentation(t *testing.T) {
	// WeChat strips indentation, so branches must work flat
	h := newTestHandler()
	dsl := `step1 @claude analyze
step2 parallel
branch @gemini @1 find bugs
branch @qwen @1 write tests
step3 @claude merge @2.1 and @2.2`

	steps, err := parseWorkflow(h, dsl)
	if err != "" {
		t.Fatalf("unexpected error: %s", err)
	}
	if len(steps) != 3 {
		t.Fatalf("expected 3 steps, got %d", len(steps))
	}
	if len(steps[1].Parallel) != 2 {
		t.Errorf("step2 parallel branches = %d, want 2", len(steps[1].Parallel))
	}
}
