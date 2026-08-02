package model

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/log"
)

func stubWaiter(calls *int, results ...error) func(context.Context, func() (bool, []log.Step, error), time.Duration, time.Duration, *api.LogTail) error {
	return func(context.Context, func() (bool, []log.Step, error), time.Duration, time.Duration, *api.LogTail) error {
		*calls++
		if *calls <= len(results) {
			return results[*calls-1]
		}
		return results[len(results)-1]
	}
}

// The push flow waits here, not in code.WaitForCodeIntegrationStatus, so this
// is where the pippin log tail has to be attached. Regression: it was wired
// into the code package, which nothing calls, and this path passed nil.
func TestWaitForImportModelJobAttachesBuildDependenciesLogTail(t *testing.T) {
	orig := waitForSteps
	defer func() { waitForSteps = orig }()

	var got *api.LogTail
	waitForSteps = func(_ context.Context, _ func() (bool, []log.Step, error), _ time.Duration, _ time.Duration, tail *api.LogTail) error {
		got = tail
		return nil
	}

	if _, _, err := waitForImportModelJob(context.Background(), "p", "job-123", "push"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got == nil {
		t.Fatal("expected a log tail to be attached to the push wait, got nil")
	}
	if got.StepID != api.BuildDependenciesStepID {
		t.Errorf("got StepID %q, want %q", got.StepID, api.BuildDependenciesStepID)
	}
	if got.Fetch == nil {
		t.Error("expected a non-nil Fetch on the log tail")
	}
}

func TestWaitForImportModelJobRetriesOnNoProgressTimeout(t *testing.T) {
	orig := waitForSteps
	defer func() { waitForSteps = orig }()

	calls := 0
	waitForSteps = stubWaiter(&calls, api.ErrorTimeout, api.ErrorTimeout, nil)

	ok, _, err := waitForImportModelJob(context.Background(), "p", "j", "push")
	if err != nil || !ok {
		t.Fatalf("expected success after retries, got ok=%v err=%v", ok, err)
	}
	if calls != 3 {
		t.Fatalf("expected 3 waiter calls, got %d", calls)
	}
}

func TestWaitForImportModelJobStopsOnCancelledContext(t *testing.T) {
	orig := waitForSteps
	defer func() { waitForSteps = orig }()

	calls := 0
	waitForSteps = stubWaiter(&calls, api.ErrorTimeout)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, _, err := waitForImportModelJob(ctx, "p", "j", "push")
	if !errors.Is(err, ErrImportModelTimeout) {
		t.Fatalf("expected ErrImportModelTimeout, got %v", err)
	}
	if calls != 1 {
		t.Fatalf("expected no retry on cancelled context, got %d calls", calls)
	}
}

func TestWaitForImportModelJobRespectsAbsoluteCap(t *testing.T) {
	origWait, origMax := waitForSteps, maxWaitForImportModelJob
	defer func() { waitForSteps, maxWaitForImportModelJob = origWait, origMax }()

	maxWaitForImportModelJob = 0
	calls := 0
	waitForSteps = stubWaiter(&calls, api.ErrorTimeout)

	_, _, err := waitForImportModelJob(context.Background(), "p", "j", "push")
	if !errors.Is(err, ErrImportModelTimeout) {
		t.Fatalf("expected ErrImportModelTimeout, got %v", err)
	}
	if calls != 1 {
		t.Fatalf("expected exactly 1 waiter call at zero cap, got %d", calls)
	}
}
