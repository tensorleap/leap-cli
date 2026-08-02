package api

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/tensorleap/leap-cli/pkg/log"
)

var ErrorTimeout = errors.New("timeout")

func WaitForCondition(ctx context.Context, message string, condition func() (bool, error), sleepDuration time.Duration, timeoutDuration time.Duration) error {
	startTime := time.Now()

	s := log.NewSpinner(message)
	s.Start()
	defer s.Stop()

	for time.Since(startTime) < timeoutDuration {
		select {
		case <-ctx.Done():
			return ErrorTimeout
		default:
			done, err := condition()
			if err != nil {
				return err
			}
			if done {
				return nil
			}

			time.Sleep(sleepDuration)
		}
	}

	return ErrorTimeout
}

var DELAY_NOT_ENDED_STEP_DURATION = 4 * time.Second

func stepsFingerprint(steps []log.Step) string {
	var result string
	for _, s := range steps {
		result += fmt.Sprintf("%s:%s:%.0f:%.0f|", s.Name, s.Status, s.Current, s.Total)
	}
	return result
}

// LogTail streams the output of one long-running step beneath the step list,
// for steps whose progress isn't visible from status alone. Fetch is called
// only while the step identified by StepID is RUNNING.
type LogTail struct {
	StepID string
	Fetch  func() ([]string, error)
}

// logTailFetchEvery throttles the tail relative to the step poll. At the usual
// 3s poll that's a refresh roughly every 6s, matching the web UI's cadence.
const logTailFetchEvery = 2

// BuildDependenciesStepID is the server-side event id for pippin's dependency
// build — the only step whose logs we stream today. Mirrors
// BUILD_DEPENDENCIES_EVENT_ID in node-server and the id the engine scheduler
// publishes.
//
// ponytail: a one-element set lives here rather than as a job-event field. If a
// second step ever wants a tail, move the decision server-side so changing it
// doesn't need a CLI release.
const BuildDependenciesStepID = "build_dependencies"

func WaitForConditionWithSteps(ctx context.Context, condition func() (bool, []log.Step, error), sleepDuration time.Duration, timeoutDuration time.Duration, tail *LogTail) error {
	lastProgressTime := time.Now()
	lastFingerprint := ""

	renderer := log.NewRenderer()
	renderer.Start()
	defer renderer.Stop()

	var doneTime *time.Time
	poll := 0

	for time.Since(lastProgressTime) < timeoutDuration {
		select {
		case <-ctx.Done():
			return ErrorTimeout
		default:
			done, steps, err := condition()
			poll++
			if len(steps) == 0 && renderer.IsTTY {
				status := log.StepStatusRunning
				if done {
					status = log.StepStatusDone
				} else if err != nil {
					status = log.StepStatusFailed
				}
				steps = []log.Step{
					{Name: "Pending...", Status: status},
				}
			}

			fp := stepsFingerprint(steps)
			if fp != lastFingerprint {
				lastProgressTime = time.Now()
				lastFingerprint = fp
			}

			if err != nil {
				markLastStepAsFailed(steps)
				renderer.Update(steps)
				time.Sleep(log.FrameDuration)
				return err
			}
			renderer.Update(steps)
			updateLogTail(renderer, tail, steps, poll)

			if done {
				isAllStepsDone := isAllStepsEnded(steps)
				if isAllStepsDone {
					time.Sleep(log.FrameDuration)
					return nil
				}
				if doneTime == nil {
					now := time.Now()
					doneTime = &now
				}

				if time.Since(*doneTime) > DELAY_NOT_ENDED_STEP_DURATION {
					time.Sleep(log.FrameDuration)
					return nil
				}
			}
			time.Sleep(sleepDuration)
		}
	}

	return ErrorTimeout
}

// updateLogTail refreshes the tail while its step is running, and takes it down
// once the step ends — except on failure, where the last lines are the most
// useful thing on screen and are left frozen for the error report that follows.
func updateLogTail(renderer *log.Renderer, tail *LogTail, steps []log.Step, poll int) {
	if tail == nil || tail.Fetch == nil {
		return
	}

	// No step, no tail — the engine replaces the whole events array when it
	// takes over, so a vanished step must not leave stale output on screen.
	step := findStep(steps, tail.StepID)
	if step == nil {
		renderer.UpdateLogs(nil)
		return
	}

	if step.Status != log.StepStatusRunning {
		if step.Status != log.StepStatusFailed {
			renderer.UpdateLogs(nil)
		}
		return
	}

	if poll%logTailFetchEvery != 0 {
		return
	}

	lines, err := tail.Fetch()
	if err != nil {
		// The tail is decoration — a failed fetch must never fail the wait.
		return
	}
	renderer.UpdateLogs(lines)
}

func findStep(steps []log.Step, id string) *log.Step {
	for i := range steps {
		if steps[i].ID == id {
			return &steps[i]
		}
	}
	return nil
}

// Terminal-success step statuses — work is done or was intentionally
// skipped. Treated equivalently when deciding whether the run is over
// or where to mark the failure when one occurs.
func isStepStatusSucceeded(status log.StepStatus) bool {
	return status == log.StepStatusDone || status == log.StepStatusSkipped
}

func markLastStepAsFailed(steps []log.Step) {
	// Index, not range-copy: log.Step is a value type, so assigning to the loop
	// variable updated a copy and left the slice untouched.
	for i := range steps {
		if !isStepStatusSucceeded(steps[i].Status) {
			steps[i].Status = log.StepStatusFailed
			break
		}
	}
}

func isAllStepsEnded(steps []log.Step) bool {
	for _, step := range steps {
		if !isStepStatusSucceeded(step.Status) && step.Status != log.StepStatusFailed {
			return false
		}
	}
	return true
}
