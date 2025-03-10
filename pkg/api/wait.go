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
