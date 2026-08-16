package model

import (
	"context"
	"fmt"
	"time"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

// RunInfo is the machine-readable description of a single run (job). It is the
// non-interactive counterpart of the push error report: for a failed push it
// carries the same three sources the interactive viewer shows, so the two can
// never disagree about what went wrong.
type RunInfo struct {
	JobId     string `json:"jobId"`
	Type      string `json:"type"`
	SubType   string `json:"subType,omitempty"`
	Status    string `json:"status"`
	ProjectId string `json:"projectId,omitempty"`
	VersionId string `json:"versionId,omitempty"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Steps []RunInfoStep `json:"steps,omitempty"`

	// ChainedEvaluate is present only on a push that asked the server to run an
	// evaluation once it finished (`leap push --eval --no-wait`).
	ChainedEvaluate *RunInfoChainedEvaluate `json:"chainedEvaluate,omitempty"`

	// ErrorReport is populated only for a failed job, matching when the
	// interactive viewer opens.
	ErrorReport *ImportModelErrorReport `json:"errorReport,omitempty"`

	Hints []string `json:"hints,omitempty"`
}

type RunInfoStep struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Status  string  `json:"status"`
	Current float64 `json:"current,omitempty"`
	Total   float64 `json:"total,omitempty"`
}

type RunInfoChainedEvaluate struct {
	Status string `json:"status"`
	JobId  string `json:"jobId,omitempty"`
	Error  string `json:"error,omitempty"`
}

// GetJobById fetches a single job by id. projectId is deliberately not required
// — a caller holding only a job id (the one thing `--no-wait` prints) can still
// resolve it.
func GetJobById(ctx context.Context, jobId string) (*tensorleapapi.Job, error) {
	params := *tensorleapapi.NewGetJobsFilterParams()
	params.SetCid([]string{jobId})
	data, res, err := api.ApiClient.GetSlimJobs(ctx).GetJobsFilterParams(params).Execute()
	if err = api.CheckRes(res, err); err != nil {
		return nil, fmt.Errorf("failed to get run %s: %w", jobId, err)
	}
	if len(data.Jobs) == 0 {
		return nil, fmt.Errorf("run %s not found", jobId)
	}
	return &data.Jobs[0], nil
}

// BuildRunInfo assembles the full report for a run. Errors fetching the
// *optional* parts (the failure detail) are surfaced as warnings rather than
// failing the command: knowing the status is more useful than knowing nothing.
func BuildRunInfo(ctx context.Context, jobId string) (*RunInfo, error) {
	job, err := GetJobById(ctx, jobId)
	if err != nil {
		return nil, err
	}

	info := &RunInfo{
		JobId:     job.Cid,
		Type:      string(job.Type),
		Status:    string(job.Status),
		ProjectId: job.GetProjectId(),
		VersionId: job.GetVersionId(),
		CreatedAt: job.CreatedAt.Format(time.RFC3339),
		UpdatedAt: job.UpdatedAt.Format(time.RFC3339),
		Steps:     runInfoSteps(job),
	}
	if job.SubType != nil {
		info.SubType = string(*job.SubType)
	}
	if chained, ok := job.GetChainedEvaluateOk(); ok && chained != nil {
		info.ChainedEvaluate = &RunInfoChainedEvaluate{
			Status: chained.Status,
			JobId:  chained.GetJobId(),
			Error:  chained.GetError(),
		}
	}

	if api.IsJobFailed(job.Status) {
		info.ErrorReport = collectRunErrorReport(ctx, job)
		info.Hints = append(info.Hints, fmt.Sprintf("for full logs run: leap run logs %s", job.Cid))
	}

	return info, nil
}

// collectRunErrorReport gathers the failure detail for a run.
//
// A push gets the collector the interactive viewer uses, so the two report the
// same thing. Every other job type gets notifications and logs only: graph
// validation records what a push checked about the code integration, and an
// evaluate runs against a version that was validated long before it started —
// the section has nothing to say about why this run failed.
func collectRunErrorReport(ctx context.Context, job *tensorleapapi.Job) *ImportModelErrorReport {
	if !wantsFullPushReport(job) {
		return CollectJobFailureDetail(ctx, job.Cid, job.CreatedAt)
	}

	projectId, versionId := job.GetProjectId(), job.GetVersionId()
	report, err := CollectImportModelJobErrors(ctx, projectId, job.Cid, versionId, job.CreatedAt)
	if err != nil {
		// Most likely the version is gone. Still worth reporting what we can.
		log.Warnf("failed to collect the full error report for run %s: %v", job.Cid, err)
		return CollectJobFailureDetail(ctx, job.Cid, job.CreatedAt)
	}
	return report
}

// wantsFullPushReport reports whether the version-scoped graph-validation
// section applies to this run — only a push, and only one we can resolve a
// version for.
func wantsFullPushReport(job *tensorleapapi.Job) bool {
	return job.Type == tensorleapapi.JOBTYPE_PUSH &&
		job.GetProjectId() != "" &&
		job.GetVersionId() != ""
}

func runInfoSteps(job *tensorleapapi.Job) []RunInfoStep {
	steps := api.StepsFromJob(job)
	out := make([]RunInfoStep, 0, len(steps))
	for _, step := range steps {
		out = append(out, RunInfoStep{
			ID:      step.ID,
			Name:    step.Name,
			Status:  string(step.Status),
			Current: step.Current,
			Total:   step.Total,
		})
	}
	return out
}
