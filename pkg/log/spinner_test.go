package log

import (
	"bytes"
	"strings"
	"testing"
)

func TestSpinnerAbortDoesNotSayCompleted(t *testing.T) {
	var out bytes.Buffer
	prev := stdoutLogger.Out
	stdoutLogger.SetOutput(&out)
	t.Cleanup(func() { stdoutLogger.SetOutput(prev) })

	s := NewSpinner("Waiting for authentication...")
	s.Start()
	s.Abort()
	if strings.Contains(out.String(), "Completed") {
		t.Fatalf("Abort printed %q", out.String())
	}

	s = NewSpinner("Waiting for authentication...")
	s.Start()
	s.Stop()
	if !strings.Contains(out.String(), "Waiting for authentication... - Completed") {
		t.Fatalf("Stop printed %q", out.String())
	}
}
