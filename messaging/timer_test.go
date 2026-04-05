package messaging

import (
	"testing"
)

func TestParseTimerDuration(t *testing.T) {
	tests := []struct {
		input       string
		wantSeconds int64
		wantLabel   string
	}{
		{"25", 1500, ""},
		{"2h", 7200, ""},
		{"30m 休息", 1800, "休息"},
		{"1.5h 写报告", 5400, "写报告"},
		{"45s", 45, ""},
		{"3", 180, ""},
		{"10min", 600, ""},
		{"2hours", 7200, ""},
		{"1h30m", 3600, "30m"}, // only parses first number+unit
		{"", 0, ""},
		{"写报告", 0, "写报告"}, // no number, returns as label for AI
	}

	for _, tt := range tests {
		seconds, label := parseTimerDuration(tt.input)
		if seconds != tt.wantSeconds {
			t.Errorf("parseTimerDuration(%q) seconds = %d, want %d", tt.input, seconds, tt.wantSeconds)
		}
		if label != tt.wantLabel {
			t.Errorf("parseTimerDuration(%q) label = %q, want %q", tt.input, label, tt.wantLabel)
		}
	}
}
