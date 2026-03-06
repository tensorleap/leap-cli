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

func WaitForConditionWithSteps(ctx context.Context, condition func() (bool, []log.Step, error), sleepDuration time.Duration, timeoutDuration time.Duration) error {
	lastProgressTime := time.Now()
	lastFingerprint := ""

	renderer := log.NewRenderer()
	renderer.Start()
	defer renderer.Stop()

	var doneTime *time.Time

	for time.Since(lastProgressTime) < timeoutDuration {
		select {
		case <-ctx.Done():
			return ErrorTimeout
		default:
			done, steps, err := condition()
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

func markLastStepAsFailed(steps []log.Step) {
	for _, step := range steps {
		if step.Status != log.StepStatusDone {
			step.Status = log.StepStatusFailed
			break
		}
	}
}

func isAllStepsEnded(steps []log.Step) bool {
	for _, step := range steps {
		if step.Status != log.StepStatusDone && step.Status != log.StepStatusFailed {
			return false
		}
	}
	return true
}
