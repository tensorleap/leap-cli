package run

import (
	"strings"
	"testing"
)

func TestCleanBuildLogLinesStripsFraming(t *testing.T) {
	tests := []struct {
		name string
		line string
		want string
	}{
		{
			name: "python logging prefix",
			line: "2026-01-15 10:32:01,441 - INFO - Building image with BuildKit",
			want: "Building image with BuildKit",
		},
		{
			name: "buildkit prefix on top of logging prefix",
			line: "2026-01-15 10:32:01,441 - INFO - [BuildKit] #8 12.4 Collecting torch",
			want: "#8 12.4 Collecting torch",
		},
		{
			name: "unframed line passes through",
			line: "#8 DONE 1204.7s",
			want: "#8 DONE 1204.7s",
		},
		{
			name: "indentation preserved after stripping",
			line: "2026-01-15 10:32:01,441 - INFO - [BuildKit] #8   Downloading torch",
			want: "#8   Downloading torch",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CleanBuildLogLines([]string{tt.line})
			if len(got) != 1 {
				t.Fatalf("got %d lines, want 1: %q", len(got), got)
			}
			if got[0] != tt.want {
				t.Errorf("got %q, want %q", got[0], tt.want)
			}
		})
	}
}

func TestCleanBuildLogLinesDropsNoise(t *testing.T) {
	dropped := []struct {
		name string
		line string
	}{
		{
			name: "rate limiter bookkeeping",
			line: "2026-01-15 10:32:01,441 - INFO - [BuildKit] ... 4213 lines suppressed (log rate limit) ...",
		},
		{
			name: "blank line",
			line: "   ",
		},
	}

	for _, tt := range dropped {
		t.Run(tt.name, func(t *testing.T) {
			if got := CleanBuildLogLines([]string{tt.line}); len(got) != 0 {
				t.Errorf("expected line to be dropped, got %q", got)
			}
		})
	}
}

// The dependency archive URL is presigned — it must never reach the user's
// terminal or scrollback.
func TestCleanBuildLogLinesRedactsSignedUrls(t *testing.T) {
	lines := []string{
		"2026-01-15 10:32:01,441 - INFO - DEPENDENCY_URL: https://bucket.s3.amazonaws.com/deps.tar.gz?X-Amz-Signature=deadbeef",
		"2026-01-15 10:32:01,441 - INFO - Downloading dependency archive from https://bucket.s3.amazonaws.com/deps.tar.gz?X-Amz-Signature=deadbeef",
		"2026-01-15 10:32:01,441 - INFO - Collecting torch",
	}

	got := CleanBuildLogLines(lines)

	for _, line := range got {
		if strings.Contains(line, "X-Amz-Signature") || strings.Contains(strings.ToUpper(line), "DEPENDENCY_URL") {
			t.Errorf("signed URL leaked through: %q", line)
		}
	}
	if len(got) != 1 || got[0] != "Collecting torch" {
		t.Errorf("expected only the safe line to survive, got %q", got)
	}
}

func TestCleanBuildLogLinesPreservesOrder(t *testing.T) {
	got := CleanBuildLogLines([]string{"first", "", "second", "third"})
	want := []string{"first", "second", "third"}

	if len(got) != len(want) {
		t.Fatalf("got %d lines, want %d: %q", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("line %d: got %q, want %q", i, got[i], want[i])
		}
	}
}
