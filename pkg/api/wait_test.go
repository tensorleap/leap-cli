package api

import (
	"context"
	"errors"
	"testing"
	"time"
)

type recordingSpinner struct{ stopped, aborted bool }

func (r *recordingSpinner) Start() {}
func (r *recordingSpinner) Stop()  { r.stopped = true }
func (r *recordingSpinner) Abort() { r.aborted = true }

// G92: the spinner line said "Completed" even when the wait failed or timed out
func TestWaitForConditionSaysCompletedOnlyWhenDone(t *testing.T) {
	failure := errors.New("boom")
	cases := []struct {
		name      string
		condition func() (bool, error)
		wantErr   error
		completed bool
	}{
		{"done", func() (bool, error) { return true, nil }, nil, true},
		{"condition failed", func() (bool, error) { return false, failure }, failure, false},
		{"timed out", func() (bool, error) { return false, nil }, ErrorTimeout, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			spinner := &recordingSpinner{}
			prev := newWaitSpinner
			newWaitSpinner = func(string) waitSpinner { return spinner }
			t.Cleanup(func() { newWaitSpinner = prev })

			err := WaitForCondition(context.Background(), "Waiting...", c.condition, time.Millisecond, 20*time.Millisecond)
			if !errors.Is(err, c.wantErr) {
				t.Fatalf("err = %v, want %v", err, c.wantErr)
			}
			if spinner.stopped != c.completed || spinner.aborted == c.completed {
				t.Fatalf("Completed printed = %v, aborted = %v; want Completed only when done", spinner.stopped, spinner.aborted)
			}
		})
	}
}
