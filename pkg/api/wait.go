package api

import (
	"context"
	"errors"
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

func WaitForConditionWithSteps(ctx context.Context, condition func() (bool, []log.Step, error), sleepDuration time.Duration, timeoutDuration time.Duration) error {
	startTime := time.Now()

	renderer := log.NewRenderer()
	renderer.Start()
	defer renderer.Stop()

	var doneTime *time.Time

	for time.Since(startTime) < timeoutDuration {
		select {
		case <-ctx.Done():
			return ErrorTimeout
		default:
			done, steps, err := condition()
			renderer.Update(steps)
			if err != nil {
				return err
			}
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

func isAllStepsEnded(steps []log.Step) bool {
	for _, step := range steps {
		if step.Status != log.StepStatusDone && step.Status != log.StepStatusFailed {
			return false
		}
	}
	return true
}
