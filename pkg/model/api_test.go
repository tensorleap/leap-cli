package model

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/log"
)

func stubWaiter(calls *int, results ...error) func(context.Context, func() (bool, []log.Step, error), time.Duration, time.Duration) error {
	return func(context.Context, func() (bool, []log.Step, error), time.Duration, time.Duration) error {
		*calls++
		if *calls <= len(results) {
			return results[*calls-1]
		}
		return results[len(results)-1]
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
