package model

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

const DEFAULT_BATCH_SIZE = 1

// ErrEvaluateCancelled is returned from the interactive concurrent-evaluate
// recovery flow when the user explicitly declines to terminate any running
// job (or aborts the prompt via Ctrl+C). It's a sentinel so RunEvaluate can
// short-circuit to a clean exit instead of treating user intent as a failure.
var ErrEvaluateCancelled = errors.New("evaluation cancelled by user")

func GetLatestEvaluateBatchSize(ctx context.Context, projectId string) (int, error) {
	versions, err := GetVersions(ctx, projectId)
	if err != nil {
		return 0, err
	}

	var latest *tensorleapapi.SlimVersion
	for _, v := range versions {
		if v.HasEvaluateParams() && (latest == nil || latest.GetCreatedAt().Before(v.GetCreatedAt())) {
			latest = &v
		}
	}
	if latest == nil {
		return 0, nil
	}
	return int(latest.EvaluateParams.GetBatchSize()), nil
}

func AskForBatchSize(defaultBatchSize int) (int, error) {

	if defaultBatchSize <= 0 {
		defaultBatchSize = DEFAULT_BATCH_SIZE
	}

	var batchSizeStr string
	prompt := &survey.Input{
		Message: "Enter batch size for evaluation:",
		Default: fmt.Sprintf("%v", defaultBatchSize),
	}
	err := survey.AskOne(prompt, &batchSizeStr, survey.WithValidator(func(val interface{}) error {
		batchSizeStr := val.(string)
		batchSize, err := strconv.Atoi(batchSizeStr)
		if err != nil {
			return fmt.Errorf("invalid batch size: %s", batchSizeStr)
		}
		if batchSize <= 0 {
			return fmt.Errorf("batch size must be greater than 0")
		}
		return nil
	}))
	if err != nil {
		return 0, err
	}

	batchSize, err := strconv.Atoi(batchSizeStr)
	if err != nil {
		return 0, fmt.Errorf("invalid batch size: %s", batchSizeStr)
	}

	return batchSize, nil
}

func RunEvaluate(ctx context.Context, projectId, versionId string, batchSize int) error {
	existingVersionParams := tensorleapapi.NewEvaluateExistingVersionParams(
		versionId,
		projectId,
		float64(batchSize),
		0,
	)

	evaluateParams := tensorleapapi.EvaluateParams{
		EvaluateExistingVersionParams: existingVersionParams,
	}

	// callEvaluate is captured so it can be re-invoked on the retry path after
	// the user terminates one of their running evaluate jobs.
	callEvaluate := func() (*tensorleapapi.Job, error) {
		job, res, err := api.ApiClient.Evaluate(ctx).EvaluateParams(evaluateParams).Execute()
		if err = api.CheckRes(res, err); err != nil {
			return nil, err
		}
		return job, nil
	}

	log.Info("Starting evaluation...")
	job, err := callEvaluate()
	if err != nil {
		// Special-case the structured concurrent-evaluate-limit error so the
		// user gets a guided recovery prompt instead of a JSON dump.
		var limitErr *api.ConcurrentEvaluateLimitError
		if errors.As(err, &limitErr) {
			job, err = handleConcurrentEvaluateLimit(ctx, limitErr, callEvaluate)
			if err != nil {
				// User-initiated cancel (or Ctrl+C) is a deliberate choice,
				// not a failure: log a friendly note and return nil so the
				// parent command exits cleanly without a scary "Error:" line
				// or a usage dump.
				if errors.Is(err, ErrEvaluateCancelled) {
					log.Warn("Evaluation skipped — cancelled by user.")
					return nil
				}
				return err
			}
		} else {
			return fmt.Errorf("failed to start evaluation: %w", err)
		}
	}

	log.Infof("Evaluation started with job ID: %s", job.GetCid())
	return nil
}

// handleConcurrentEvaluateLimit interactively walks the user through
// terminating one of their running Evaluate/Continue-Evaluate jobs and
// retrying the original evaluate call. Returns the newly-started job on
// success, or an error if the user cancels or termination/retry fails.
//
// Mirrors the web-ui's ConcurrentEvaluateLimitDialog terminate-and-retry
// flow so users get the same recovery experience from either entry point.
func handleConcurrentEvaluateLimit(
	ctx context.Context,
	initial *api.ConcurrentEvaluateLimitError,
	callEvaluate func() (*tensorleapapi.Job, error),
) (*tensorleapapi.Job, error) {
	current := initial
	for {
		log.Warn(current.Error())
		if len(current.RunningJobs) == 0 {
			// Shouldn't happen — backend always sends at least one row when
			// the limit is hit — but guard anyway.
			return nil, errors.New(
				"concurrent evaluate jobs limit reached, but the server " +
					"reported no running jobs to terminate",
			)
		}

		options, byLabel := buildEvaluateJobOptions(current.RunningJobs)
		const cancelLabel = "Cancel — do not start this evaluation"
		options = append(options, cancelLabel)

		var selected string
		prompt := &survey.Select{
			Message: "Select a running Evaluate job to terminate, then this evaluation will be retried:",
			Options: options,
		}
		if err := survey.AskOne(prompt, &selected); err != nil {
			// Treat any prompt failure (most commonly Ctrl+C / InterruptErr)
			// as a deliberate cancellation so RunEvaluate exits cleanly.
			return nil, errors.Join(ErrEvaluateCancelled, err)
		}
		if selected == cancelLabel {
			return nil, ErrEvaluateCancelled
		}

		chosen := byLabel[selected]
		log.Infof("Terminating job %s ...", chosen.JobId)
		if err := terminateJob(ctx, chosen.JobId); err != nil {
			return nil, fmt.Errorf("failed to terminate job %s: %w", chosen.JobId, err)
		}

		log.Info("Retrying evaluation...")
		job, err := callEvaluate()
		if err == nil {
			return job, nil
		}

		// If we're still blocked (e.g., a different job started in the
		// meantime), loop with the fresh running-jobs list. Any other error
		// short-circuits with the original error message.
		var nextLimit *api.ConcurrentEvaluateLimitError
		if errors.As(err, &nextLimit) {
			log.Warn("Still at the concurrent-evaluate-jobs limit after termination.")
			current = nextLimit
			continue
		}
		return nil, fmt.Errorf("failed to start evaluation: %w", err)
	}
}

func buildEvaluateJobOptions(
	jobs []api.RunningEvaluateJobInfo,
) ([]string, map[string]api.RunningEvaluateJobInfo) {
	options := make([]string, 0, len(jobs))
	byLabel := make(map[string]api.RunningEvaluateJobInfo, len(jobs))
	for _, j := range jobs {
		startedLabel := j.StartedAt
		if formatted, err := api.ParseAndFormatDateToLocalTime(j.StartedAt); err == nil {
			startedLabel = formatted
		}
		label := fmt.Sprintf(
			"jobId=%s  versionId=%s  started=%s",
			j.JobId, j.VersionId, startedLabel,
		)
		options = append(options, label)
		byLabel[label] = j
	}
	return options, byLabel
}

func terminateJob(ctx context.Context, jobId string) error {
	params := tensorleapapi.NewTerminateJobParams(jobId)
	_, res, err := api.ApiClient.TerminateJob(ctx).TerminateJobParams(*params).Execute()
	return api.CheckRes(res, err)
}
