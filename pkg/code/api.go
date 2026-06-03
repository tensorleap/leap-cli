package code

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/interactive_pages"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

type CodeSnapshot = tensorleapapi.CodeSnapshot

func GetCodeSnapshot(ctx context.Context, projectId, id string) (*CodeSnapshot, error) {
	res, response, err := api.ApiClient.GetCodeSnapshot(ctx).GetCodeSnapshotParams(*tensorleapapi.NewGetCodeSnapshotParams(projectId, id)).Execute()
	if err = api.CheckRes(response, err); err != nil {
		return nil, err
	}
	return &res.CodeSnapshot, err
}

// PushCodeAndModel uploads the code bundle and triggers the combined push job
// (code parse + model import in one job) via the new push endpoint, replacing
// the separate pushCodeSnapshot + importModel calls.
func PushCodeAndModel(
	ctx context.Context, tarGzFile io.Reader, fileSize int64,
	entryFile, secretId, pythonVersion,
	versionName, projectId, branch string,
	overwriteVersionId string,
	modelInfo tensorleapapi.ImportModelInfo,
) (*tensorleapapi.PushResponse, error) {

	uploadUrl, err := GetCodeSnapshotUploadUrl(ctx, projectId)
	if err != nil {
		return nil, err
	}

	if err := api.UploadFile(uploadUrl, tarGzFile, fileSize); err != nil {
		return nil, err
	}

	pushParams := *tensorleapapi.NewPushParams(
		projectId,
		uploadUrl,
		entryFile,
		versionName,
		modelInfo,
	)

	if len(overwriteVersionId) > 0 {
		pushParams.SetOverwriteVersionId(overwriteVersionId)
	}

	if len(pythonVersion) > 0 {
		pushParams.GenericBaseImageType = &pythonVersion
	}

	if len(branch) > 0 {
		pushParams.SetBranchName(branch)
	}

	if len(secretId) > 0 {
		pushParams.SecretManagerId = &secretId
	}

	log.Info("Pushing code and model...")
	result, response, err := api.ApiClient.Push(ctx).
		PushParams(pushParams).
		Execute()
	if err = api.CheckRes(response, err); err != nil {
		return nil, err
	}

	return result, nil
}

// PushOverride re-pushes code to an existing version and re-validates its
// existing model via the override-push endpoint — no new model upload, no name
// (both come from the overwritten version).
func PushOverride(
	ctx context.Context, tarGzFile io.Reader, fileSize int64,
	entryFile, secretId, pythonVersion, projectId, branch string,
	overwriteVersionId string,
) (*tensorleapapi.PushResponse, error) {

	uploadUrl, err := GetCodeSnapshotUploadUrl(ctx, projectId)
	if err != nil {
		return nil, err
	}

	if err := api.UploadFile(uploadUrl, tarGzFile, fileSize); err != nil {
		return nil, err
	}

	params := *tensorleapapi.NewPushOverrideParams(
		projectId,
		uploadUrl,
		entryFile,
		overwriteVersionId,
	)

	if len(pythonVersion) > 0 {
		params.GenericBaseImageType = &pythonVersion
	}

	if len(branch) > 0 {
		params.SetBranchName(branch)
	}

	if len(secretId) > 0 {
		params.SecretManagerId = &secretId
	}

	log.Info("Pushing code (override)...")
	result, response, err := api.ApiClient.PushOverride(ctx).
		PushOverrideParams(params).
		Execute()
	if err = api.CheckRes(response, err); err != nil {
		return nil, err
	}

	return result, nil
}

const TIMEOUT_FOR_CODE_INTEGRATION_STATUS = 90 * time.Minute

var ErrCodeIntegrationTimeout = fmt.Errorf("timeout occurred while waiting for the integration code status")

func WaitForCodeIntegrationStatus(ctx context.Context, projectId, codeSnapshotId string) (bool, error) {
	log.Info("Waiting for code parser result...")
	sleepDuration := 3 * time.Second
	commandStartTime := time.Now()

	var jobId string

	condition := func() (bool, []log.Step, error) {
		getJobParams := *tensorleapapi.NewGetJobsFilterParams()
		getJobParams.SetCodeSnapshotId(codeSnapshotId)
		codeIntegrationJobs, _, err := api.ApiClient.GetSlimJobs(ctx).GetJobsFilterParams(
			getJobParams,
		).Execute()
		if err != nil {
			return false, nil, fmt.Errorf("failed to get the code integration job: %v", err)
		}
		if len(codeIntegrationJobs.Jobs) == 0 {
			return false, nil, fmt.Errorf("failed to get the code integration job")
		}
		codeIntegrationJob := codeIntegrationJobs.Jobs[0]
		jobId = codeIntegrationJob.Cid
		steps := api.StepsFromJob(&codeIntegrationJob)
		switch true {
		case api.IsJobFinished(codeIntegrationJob.Status):
			return true, steps, nil
		case api.IsJobFailed(codeIntegrationJob.Status):
			return false, steps, fmt.Errorf("code integration job failed")
		}
		return false, steps, nil
	}

	err := api.WaitForConditionWithSteps(ctx, condition, sleepDuration, TIMEOUT_FOR_CODE_INTEGRATION_STATUS)

	if err == api.ErrorTimeout {
		return false, ErrCodeIntegrationTimeout
	}
	if err != nil {
		log.Errorf("failed to wait for the integration code status: %v", err)

		codeSnapshot, getErr := GetCodeSnapshot(ctx, projectId, codeSnapshotId)
		if getErr != nil {
			return false, fmt.Errorf("failed to get the code integration version: %v", getErr)
		}
		report, collectErr := CollectErrorsOnCodeParseFailed(ctx, jobId, codeSnapshot, commandStartTime)
		if collectErr != nil {
			return false, fmt.Errorf("failed to collect errors on code parse failed: %v", collectErr)
		}
		asPages := report.ToReportPages()
		runErr := interactive_pages.RunInteractivePages(asPages)
		if runErr != nil {
			return false, fmt.Errorf("failed to run interactive pages: %v", runErr)
		}
		return false, fmt.Errorf("code parsing failed see errors above, for more logs run: leap run logs %s", jobId)
	}
	return true, nil
}

func GetCodeSnapshotUploadUrl(ctx context.Context, projectId string) (string, error) {
	baseURL := api.GetBaseUrlFromContext(ctx)
	params := *tensorleapapi.NewGetCodeSnapshotUploadUrlParams(projectId)
	params.SetOrigin(baseURL)
	data, _, err := api.ApiClient.GetCodeSnapshotUploadUrl(ctx).
		GetCodeSnapshotUploadUrlParams(params).
		Execute()
	if err != nil {
		return "", err
	}
	return data.GetUrl(), nil
}
