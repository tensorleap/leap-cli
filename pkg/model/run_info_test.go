package model

import (
	"testing"

	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

// Graph validation describes what a push checked about the code integration.
// An evaluate runs against an already-validated version, so the section says
// nothing about why that run failed and must not be collected for it.
func TestOnlyAPushCollectsValidationErrors(t *testing.T) {
	subType := tensorleapapi.JOBSUBTYPE_EVALUATE
	evaluate := &tensorleapapi.Job{
		Cid:       "job-1",
		Type:      tensorleapapi.JOBTYPE_TRAINING,
		SubType:   &subType,
		Status:    tensorleapapi.JOBSTATUS_TERMINATED,
		ProjectId: strPtr("project-1"),
		VersionId: strPtr("version-1"),
	}

	if wantsFullPushReport(evaluate) {
		t.Fatal("an evaluate must not collect the push validation report")
	}

	pushSubType := tensorleapapi.JOBSUBTYPE_PUSH
	push := &tensorleapapi.Job{
		Cid:       "job-2",
		Type:      tensorleapapi.JOBTYPE_PUSH,
		SubType:   &pushSubType,
		Status:    tensorleapapi.JOBSTATUS_FAILED,
		ProjectId: strPtr("project-1"),
		VersionId: strPtr("version-1"),
	}

	if !wantsFullPushReport(push) {
		t.Fatal("a push must collect the full report, validation included")
	}
}

// A push whose project/version can't be resolved has no validation data to
// fetch either, so it falls back the same way an evaluate does.
func TestPushWithoutAVersionFallsBack(t *testing.T) {
	subType := tensorleapapi.JOBSUBTYPE_PUSH
	push := &tensorleapapi.Job{
		Cid:     "job-3",
		Type:    tensorleapapi.JOBTYPE_PUSH,
		SubType: &subType,
		Status:  tensorleapapi.JOBSTATUS_FAILED,
	}

	if wantsFullPushReport(push) {
		t.Fatal("a push with no version cannot collect validation errors")
	}
}

func strPtr(s string) *string { return &s }
