package hub

import (
	"fmt"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestHub_SaveAndRead(t *testing.T) {
	dir := t.TempDir()
	h := New(dir)

	// Test Save
	path, err := h.Save("test.md", "hello world", "test-agent")
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	// Verify path
	expectedPath := filepath.Join(dir, "test.md")
	if path != expectedPath {
		t.Errorf("expected path %s, got %s", expectedPath, path)
	}

	// Verify content with frontmatter
	content, err := h.ReadFile("test.md")
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}

	if !strings.Contains(content, "agent: test-agent") {
		t.Error("frontmatter missing agent")
	}
	if !strings.Contains(content, "hello world") {
		t.Error("content missing")
	}
	if !strings.Contains(content, "timestamp:") {
		t.Error("frontmatter missing timestamp")
	}
}

func TestHub_SaveAutoExtension(t *testing.T) {
	dir := t.TempDir()
	h := New(dir)

	// Save without .md extension
	path, err := h.Save("myfile", "content", "agent")
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	if !strings.HasSuffix(path, ".md") {
		t.Errorf("expected .md extension, got %s", path)
	}

	// Verify we can read it with or without extension
	content, err := h.ReadFile("myfile.md")
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}
	if !strings.Contains(content, "content") {
		t.Error("content mismatch")
	}
}

func TestHub_SaveConflictAutoRename(t *testing.T) {
	dir := t.TempDir()
	h := New(dir)

	// Save first file
	path1, err := h.Save("test.md", "v1", "agent1")
	if err != nil {
		t.Fatalf("First save failed: %v", err)
	}

	// Save same filename again - should auto-rename
	path2, err := h.Save("test.md", "v2", "agent2")
	if err != nil {
		t.Fatalf("Second save failed: %v", err)
	}

	// Paths should be different (auto-renamed)
	if path1 == path2 {
		t.Error("expected auto-rename on conflict, but paths are same")
	}

	// Original file should still have v1
	content1, err := h.ReadFile("test.md")
	if err != nil {
		t.Fatalf("ReadFile original failed: %v", err)
	}
	if !strings.Contains(content1, "v1") {
		t.Error("original file should contain v1")
	}
}

func TestHub_ReadAll(t *testing.T) {
	dir := t.TempDir()
	h := New(dir)

	// Save multiple files
	h.Save("file1.md", "content1", "agent1")
	h.Save("file2.md", "content2", "agent2")

	// Read all
	all, err := h.ReadAll()
	if err != nil {
		t.Fatalf("ReadAll failed: %v", err)
	}

	if !strings.Contains(all, "content1") {
		t.Error("missing content1")
	}
	if !strings.Contains(all, "content2") {
		t.Error("missing content2")
	}
	if !strings.Contains(all, "=== Agent Hub Shared Context ===") {
		t.Error("missing header")
	}
}

func TestHub_List(t *testing.T) {
	dir := t.TempDir()
	h := New(dir)

	// Empty hub
	names, err := h.List()
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}
	if len(names) != 0 {
		t.Errorf("expected empty list, got %d", len(names))
	}

	// Add files
	h.Save("alpha.md", "a", "agent")
	h.Save("beta.md", "b", "agent")

	names, err = h.List()
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}
	if len(names) != 2 {
		t.Errorf("expected 2 files, got %d", len(names))
	}
}

func TestHub_ListWithInfo(t *testing.T) {
	dir := t.TempDir()
	h := New(dir)

	// Add files with slight delay to ensure different timestamps
	h.Save("old.md", "old content", "agent")
	time.Sleep(10 * time.Millisecond)
	h.Save("new.md", "new content", "agent")

	files, err := h.ListWithInfo()
	if err != nil {
		t.Fatalf("ListWithInfo failed: %v", err)
	}

	if len(files) != 2 {
		t.Fatalf("expected 2 files, got %d", len(files))
	}

	// Should be sorted newest first
	if files[0].Name != "new.md" {
		t.Errorf("expected newest first, got %s", files[0].Name)
	}
	if files[1].Name != "old.md" {
		t.Errorf("expected oldest second, got %s", files[1].Name)
	}
}

func TestHub_FindByPartialName(t *testing.T) {
	dir := t.TempDir()
	h := New(dir)

	h.Save("pipe_gemini_analysis.md", "content1", "gemini")
	h.Save("pipe_claude_review.md", "content2", "claude")

	tests := []struct {
		partial string
		expect  string
		wantErr bool
	}{
		{"gemini", "pipe_gemini_analysis.md", false},
		{"gem", "pipe_gemini_analysis.md", false},
		{"claude", "pipe_claude_review.md", false},
		{"nonexistent", "", true},
		{"", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.partial, func(t *testing.T) {
			name, err := h.FindByPartialName(tt.partial)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if name != tt.expect {
					t.Errorf("expected %s, got %s", tt.expect, name)
				}
			}
		})
	}
}

func TestHub_FindByPartialNameNewest(t *testing.T) {
	dir := t.TempDir()
	h := New(dir)

	// Save multiple files matching same partial
	h.Save("gemini_v1.md", "old", "agent")
	time.Sleep(10 * time.Millisecond)
	h.Save("gemini_v2.md", "new", "agent")

	// Should return newest
	name, err := h.FindByPartialName("gemini")
	if err != nil {
		t.Fatalf("FindByPartialName failed: %v", err)
	}

	if name != "gemini_v2.md" {
		t.Errorf("expected newest file gemini_v2.md, got %s", name)
	}
}

func TestHub_Exists(t *testing.T) {
	dir := t.TempDir()
	h := New(dir)

	if h.Exists("nonexistent.md") {
		t.Error("expected false for nonexistent file")
	}

	h.Save("exists.md", "content", "agent")
	if !h.Exists("exists.md") {
		t.Error("expected true for existing file")
	}
}

func TestHub_Clear(t *testing.T) {
	dir := t.TempDir()
	h := New(dir)

	h.Save("file1.md", "a", "agent")
	h.Save("file2.md", "b", "agent")

	count, err := h.Clear()
	if err != nil {
		t.Fatalf("Clear failed: %v", err)
	}
	if count != 2 {
		t.Errorf("expected 2 files cleared, got %d", count)
	}

	files, _ := h.List()
	if len(files) != 0 {
		t.Errorf("expected empty hub after clear, got %d files", len(files))
	}
}

func TestHub_ReadSpecific(t *testing.T) {
	dir := t.TempDir()
	h := New(dir)

	h.Save("file1.md", "content1", "agent1")
	h.Save("file2.md", "content2", "agent2")

	result, err := h.ReadSpecific([]string{"file1.md", "file2.md", "nonexistent.md"})
	if err != nil {
		t.Fatalf("ReadSpecific failed: %v", err)
	}

	if !strings.Contains(result, "content1") {
		t.Error("missing content1")
	}
	if !strings.Contains(result, "content2") {
		t.Error("missing content2")
	}
	if !strings.Contains(result, "nonexistent.md (not found)") {
		t.Error("missing not found indicator")
	}
}

func TestHub_ConcurrentAccess(t *testing.T) {
	dir := t.TempDir()
	h := New(dir)

	var wg sync.WaitGroup
	errCh := make(chan error, 20)

	// Concurrent writes
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			_, err := h.Save(fmt.Sprintf("file%d.md", id), fmt.Sprintf("content%d", id), "agent")
			if err != nil {
				errCh <- err
			}
		}(i)
	}

	// Concurrent reads
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := h.ReadAll()
			if err != nil {
				errCh <- err
			}
		}()
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		t.Errorf("concurrent access error: %v", err)
	}
}

func TestHub_SaveRaw(t *testing.T) {
	dir := t.TempDir()
	h := New(dir)

	_, err := h.SaveRaw("raw.txt", "raw content")
	if err != nil {
		t.Fatalf("SaveRaw failed: %v", err)
	}

	content, err := h.ReadFile("raw.txt")
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}

	// Should NOT have frontmatter
	if strings.Contains(content, "agent:") {
		t.Error("SaveRaw should not add frontmatter")
	}
	if content != "raw content" {
		t.Errorf("expected 'raw content', got %q", content)
	}
}

func TestSanitizeFilename(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"normal.md", "normal.md"},
		{"../../etc/passwd", "passwd"},
		{"test\x00file.md", "testfile.md"},
		{"", "untitled.md"},
		{".", "untitled.md"},
		{"..", "untitled.md"},
		{"   ", "untitled.md"},
		{"/path/to/file.md", "file.md"},
		// Long filename should be truncated to 255 chars
		{strings.Repeat("a", 300), strings.Repeat("a", 255)},
		// Path with illegal chars - filepath.Base extracts last component after /
		{"test<>:\"/\\|?*.md", "____.md"},
		// Windows reserved names without extension get .md added by caller
		{"CON", "_CON"},
		{"con.md", "_con.md"},
		{"PRN.md", "_PRN.md"},
		{"AUX", "_AUX"},
		{"NUL.md", "_NUL.md"},
		{"COM1.md", "_COM1.md"},
		{"LPT1.md", "_LPT1.md"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%q", tt.input), func(t *testing.T) {
			got := sanitizeFilename(tt.input)
			if got != tt.expect {
				t.Errorf("sanitizeFilename(%q) = %q, want %q", tt.input, got, tt.expect)
			}
		})
	}
}

func TestHub_SizeLimit(t *testing.T) {
	dir := t.TempDir()
	h := New(dir)

	// Try to save a file larger than limit
	largeContent := strings.Repeat("x", 2*1024*1024) // 2MB
	_, err := h.Save("large.md", largeContent, "agent")
	if err == nil {
		t.Error("expected error for oversized file")
	}
	if !strings.Contains(err.Error(), "too large") {
		t.Errorf("expected size limit error, got: %v", err)
	}
}
