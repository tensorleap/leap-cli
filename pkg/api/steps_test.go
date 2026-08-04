package api

import (
	"testing"

	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

// The log tail is scoped by event id, so the id has to survive the conversion —
// matching on the display name would break the moment anyone rewords it.
func TestStepsFromJobCarriesEventID(t *testing.T) {
	job := &tensorleapapi.Job{
		EventsSnapshot: &tensorleapapi.EventsSnapshot{
			Events: []tensorleapapi.JobEvent{
				{
					Id:     "build_dependencies",
					Name:   "Build Dependencies Preprocess",
					Status: tensorleapapi.STATUSENUM_STARTED,
				},
			},
		},
	}

	steps := StepsFromJob(job)

	if len(steps) != 1 {
		t.Fatalf("got %d steps, want 1", len(steps))
	}
	if steps[0].ID != "build_dependencies" {
		t.Errorf("got ID %q, want %q", steps[0].ID, "build_dependencies")
	}
	if steps[0].Status != log.StepStatusRunning {
		t.Errorf("got status %q, want %q", steps[0].Status, log.StepStatusRunning)
	}
}

// Regression: this iterated by value over a slice of a value type, so the
// status assignment landed on a copy and the slice was never modified.
func TestMarkLastStepAsFailedMutatesSlice(t *testing.T) {
	steps := []log.Step{
		{ID: "a", Status: log.StepStatusDone},
		{ID: "b", Status: log.StepStatusSkipped},
		{ID: "c", Status: log.StepStatusRunning},
		{ID: "d", Status: log.StepStatusPending},
	}

	markLastStepAsFailed(steps)

	if steps[2].Status != log.StepStatusFailed {
		t.Errorf("expected the first unfinished step to be FAILED, got %q", steps[2].Status)
	}
	if steps[3].Status != log.StepStatusPending {
		t.Errorf("expected later steps untouched, got %q", steps[3].Status)
	}
	if steps[0].Status != log.StepStatusDone || steps[1].Status != log.StepStatusSkipped {
		t.Error("expected terminal-success steps untouched")
	}
}
