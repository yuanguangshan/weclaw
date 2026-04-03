package hub

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// MaxHubFileSize is the maximum allowed file size (1MB)
const MaxHubFileSize = 1 * 1024 * 1024

// Hub manages shared context files for cross-agent collaboration.
type Hub struct {
	mu        sync.RWMutex // protects all file operations
	sharedDir string       // directory for shared context files
}

// New creates a new Hub with the given shared directory.
func New(sharedDir string) *Hub {
	os.MkdirAll(sharedDir, 0o755)
	return &Hub{sharedDir: sharedDir}
}

// DefaultDir returns the default hub shared directory path.
func DefaultDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return filepath.Join(os.TempDir(), "weclaw-hub", "shared")
	}
	return filepath.Join(home, ".weclaw", "hub", "shared")
}

// SharedDir returns the hub's shared directory path.
func (h *Hub) SharedDir() string {
	return h.sharedDir
}

// Save writes content to a file in the shared directory with YAML frontmatter.
// agentName identifies which agent produced the content.
// If file already exists, auto-renames with timestamp suffix to avoid overwrite.
func (h *Hub) Save(filename, content, agentName string) (string, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	// Check size limit
	if len(content) > MaxHubFileSize {
		return "", fmt.Errorf("file too large (%.1f MB), limit is %d MB",
			float64(len(content))/(1024*1024), MaxHubFileSize/(1024*1024))
	}

	// Sanitize filename
	filename = sanitizeFilename(filename)
	if !strings.HasSuffix(filename, ".md") {
		filename += ".md"
	}

	filePath := filepath.Join(h.sharedDir, filename)

	// Check for conflict and auto-rename
	if _, err := os.Stat(filePath); err == nil {
		// File exists, add timestamp suffix
		base := strings.TrimSuffix(filename, ".md")
		ts := time.Now().Format("20060102-150405")
		newFilename := fmt.Sprintf("%s_%s.md", base, ts)
		filePath = filepath.Join(h.sharedDir, newFilename)
		filename = newFilename
	}

	// Build frontmatter with UTC timestamp
	timestamp := time.Now().UTC().Format(time.RFC3339)
	frontmatter := fmt.Sprintf("---\nagent: %s\ntimestamp: %s\n---\n\n", agentName, timestamp)

	fullContent := frontmatter + content

	if err := os.WriteFile(filePath, []byte(fullContent), 0o644); err != nil {
		return "", fmt.Errorf("save hub file: %w", err)
	}

	return filePath, nil
}

// SaveRaw writes raw content to a file (no frontmatter) in the shared directory.
func (h *Hub) SaveRaw(filename, content string) (string, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	filename = sanitizeFilename(filename)
	filePath := filepath.Join(h.sharedDir, filename)

	if err := os.WriteFile(filePath, []byte(content), 0o644); err != nil {
		return "", fmt.Errorf("save hub file: %w", err)
	}

	return filePath, nil
}

// ReadFile reads a specific file from the shared directory.
func (h *Hub) ReadFile(filename string) (string, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	filename = sanitizeFilename(filename)
	filePath := filepath.Join(h.sharedDir, filename)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("read hub file: %w", err)
	}

	return string(data), nil
}

// ReadAll reads all files from the shared directory and returns their combined content.
// Returns a formatted context string ready for injection into agent prompts.
func (h *Hub) ReadAll() (string, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	entries, err := os.ReadDir(h.sharedDir)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil // empty hub is fine
		}
		return "", fmt.Errorf("read hub directory: %w", err)
	}

	if len(entries) == 0 {
		return "", nil
	}

	// Sort by modification time (oldest first)
	type fileEntry struct {
		name string
		info os.FileInfo
	}
	var files []fileEntry
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		info, err := e.Info()
		if err != nil {
			continue
		}
		files = append(files, fileEntry{name: e.Name(), info: info})
	}
	sort.Slice(files, func(i, j int) bool {
		return files[i].info.ModTime().Before(files[j].info.ModTime())
	})

	var sb strings.Builder
	sb.WriteString("=== Agent Hub Shared Context ===\n\n")

	for _, f := range files {
		data, err := os.ReadFile(filepath.Join(h.sharedDir, f.name))
		if err != nil {
			continue
		}

		sb.WriteString(fmt.Sprintf("--- %s ---\n", f.name))
		sb.Write(data)
		sb.WriteString("\n\n")
	}

	sb.WriteString("=== End Hub Context ===\n")
	return sb.String(), nil
}

// List returns all filenames in the shared directory.
func (h *Hub) List() ([]string, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	entries, err := os.ReadDir(h.sharedDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("list hub directory: %w", err)
	}

	var names []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		names = append(names, e.Name())
	}

	sort.Strings(names)
	return names, nil
}

// FileInfo holds filename and modification time.
type FileInfo struct {
	Name    string
	ModTime time.Time
}

// ListWithInfo returns all files with their modification time, sorted by newest first.
func (h *Hub) ListWithInfo() ([]FileInfo, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	entries, err := os.ReadDir(h.sharedDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("list hub directory: %w", err)
	}

	var files []FileInfo
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		info, err := e.Info()
		if err != nil {
			continue
		}
		files = append(files, FileInfo{Name: e.Name(), ModTime: info.ModTime()})
	}

	// Sort by modification time, newest first
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime.After(files[j].ModTime)
	})

	return files, nil
}

// Clear removes all files from the shared directory.
func (h *Hub) Clear() (int, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	entries, err := os.ReadDir(h.sharedDir)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil
		}
		return 0, fmt.Errorf("clear hub directory: %w", err)
	}

	count := 0
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		path := filepath.Join(h.sharedDir, e.Name())
		if err := os.Remove(path); err != nil {
			continue
		}
		count++
	}

	return count, nil
}

// ReadSpecific reads specific files from the shared directory.
// filenames is a list of filenames to read.
func (h *Hub) ReadSpecific(filenames []string) (string, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	var sb strings.Builder
	sb.WriteString("=== Agent Hub Shared Context ===\n\n")

	for _, fname := range filenames {
		fname = sanitizeFilename(fname)
		data, err := os.ReadFile(filepath.Join(h.sharedDir, fname))
		if err != nil {
			sb.WriteString(fmt.Sprintf("--- %s (not found) ---\n\n", fname))
			continue
		}

		sb.WriteString(fmt.Sprintf("--- %s ---\n", fname))
		sb.Write(data)
		sb.WriteString("\n\n")
	}

	sb.WriteString("=== End Hub Context ===\n")
	return sb.String(), nil
}

// Exists checks if a file exists in the shared directory.
func (h *Hub) Exists(filename string) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()

	filename = sanitizeFilename(filename)
	_, err := os.Stat(filepath.Join(h.sharedDir, filename))
	return err == nil
}

// FindByPartialName finds a file by partial name matching.
// Returns the newest matching file, or empty string if not found.
// Supports partial matching: "gemini" matches "pipe_20260402_gemini.md"
func (h *Hub) FindByPartialName(partial string) (string, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if partial == "" {
		return "", fmt.Errorf("partial name is empty")
	}

	partial = strings.ToLower(strings.TrimSpace(partial))
	// Remove .md suffix if user included it
	partial = strings.TrimSuffix(partial, ".md")

	entries, err := os.ReadDir(h.sharedDir)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("no files found")
		}
		return "", fmt.Errorf("read hub directory: %w", err)
	}

	// Find all matching files
	type match struct {
		name    string
		modTime time.Time
	}
	var matches []match

	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		// Remove .md suffix for comparison
		baseName := strings.TrimSuffix(name, ".md")

		// Partial match (case-insensitive)
		if strings.Contains(strings.ToLower(baseName), partial) {
			info, err := e.Info()
			if err != nil {
				continue
			}
			matches = append(matches, match{name: name, modTime: info.ModTime()})
		}
	}

	if len(matches) == 0 {
		return "", fmt.Errorf("no files matching %q", partial)
	}

	// Return newest match
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].modTime.After(matches[j].modTime)
	})

	return matches[0].name, nil
}

// BuildPrompt creates a prompt with hub context injected.
// If context is empty, returns the original message.
func BuildPrompt(context, message string) string {
	if context == "" {
		return message
	}
	return fmt.Sprintf("%s\n\n%s", context, message)
}

// sanitizeFilename removes path traversal attempts and dangerous characters.
// Also handles Windows reserved names and length limits.
func sanitizeFilename(name string) string {
	// Remove path components
	name = filepath.Base(name)
	// Remove null bytes
	name = strings.ReplaceAll(name, "\x00", "")
	name = strings.TrimSpace(name)

	// Replace Windows illegal characters: < > : " / \ | ? *
	illegalChars := []string{"<", ">", ":", "\"", "/", "\\", "|", "?", "*"}
	for _, c := range illegalChars {
		name = strings.ReplaceAll(name, c, "_")
	}

	// Handle Windows reserved names (CON, PRN, AUX, NUL, COM1-9, LPT1-9)
	reserved := []string{"CON", "PRN", "AUX", "NUL",
		"COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9",
		"LPT1", "LPT2", "LPT3", "LPT4", "LPT5", "LPT6", "LPT7", "LPT8", "LPT9"}

	baseName := strings.TrimSuffix(name, ".md")
	for _, r := range reserved {
		if strings.EqualFold(baseName, r) {
			name = "_" + name
			break
		}
	}

	// Length limit (255 chars max on most filesystems)
	if len(name) > 255 {
		ext := filepath.Ext(name)
		base := strings.TrimSuffix(name, ext)
		maxBase := 255 - len(ext)
		if maxBase < 1 {
			maxBase = 250
			ext = ""
		}
		name = base[:maxBase] + ext
	}

	if name == "" || name == "." || name == ".." {
		return "untitled.md"
	}
	return name
}
