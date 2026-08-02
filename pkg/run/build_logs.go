package run

import (
	"context"
	"regexp"
	"strings"
)

// Pippin (the dependency-builder init container) logs through Python's logging
// module, so every line carries a timestamp and level, and BuildKit output
// carries a second prefix on top of that.
var (
	pippinLogPrefix = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2},\d{3} - \w+ - `)
	buildKitPrefix  = regexp.MustCompile(`^\[BuildKit\] `)
)

// dropLine matches pippin output that must never reach the user's terminal.
var dropLine = []*regexp.Regexp{
	// Pippin's own log rate limiter — an artifact of our plumbing, not the build.
	regexp.MustCompile(`lines suppressed \(log rate limit\)`),
	// Presigned download URL for the dependency archive.
	regexp.MustCompile(`(?i)DEPENDENCY_URL`),
	// Any other signed URL that shows up in the build output.
	regexp.MustCompile(`[?&](X-Amz-Signature|Signature|sig)=`),
}

// CleanBuildLogLines strips log framing and drops lines that are noise or
// unsafe to display, preserving order. Complements GetTopLogs, which keeps
// lines matching a pattern; this one drops them.
func CleanBuildLogLines(raw []string) []string {
	out := make([]string, 0, len(raw))
	for _, line := range raw {
		line = pippinLogPrefix.ReplaceAllString(line, "")
		line = buildKitPrefix.ReplaceAllString(line, "")
		if strings.TrimSpace(line) == "" {
			continue
		}
		if shouldDrop(line) {
			continue
		}
		out = append(out, line)
	}
	return out
}

func shouldDrop(line string) bool {
	for _, pat := range dropLine {
		if pat.MatchString(line) {
			return true
		}
	}
	return false
}

// GetBuildLogTail fetches the job's pod logs and returns the last `count`
// displayable lines. While the dependency-builder init container is running the
// engine container hasn't started, so the pod's log is pippin's output.
func GetBuildLogTail(ctx context.Context, jobId string, count int) ([]string, error) {
	podLogs, err := GetRunLogs(ctx, jobId)
	if err != nil {
		return nil, err
	}

	var lines []string
	for _, podLog := range podLogs {
		if isDescribeLog(podLog.Name) {
			continue
		}
		lines = append(lines, strings.Split(podLog.Logs, "\n")...)
	}

	lines = CleanBuildLogLines(lines)
	if len(lines) > count {
		lines = lines[len(lines)-count:]
	}
	return lines, nil
}

// getPodLogs appends a `describe-<pod>` entry alongside the real logs; it's
// kubectl-describe output, not build progress.
func isDescribeLog(name string) bool {
	return strings.HasPrefix(strings.ToLower(name), "describe")
}
