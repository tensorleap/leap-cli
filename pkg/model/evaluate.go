package model

import (
	"context"
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

const DEFAULT_BATCH_SIZE = 1

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

	log.Info("Starting evaluation...")
	job, res, err := api.ApiClient.Evaluate(ctx).EvaluateParams(evaluateParams).Execute()
	if err = api.CheckRes(res, err); err != nil {
		return fmt.Errorf("failed to start evaluation: %w", err)
	}

	log.Infof("Evaluation started with job ID: %s", job.GetCid())
	return nil
}
