package code

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/tensorleap/leap-cli/pkg/api"
	. "github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

type CodeSnapshot = tensorleapapi.CodeSnapshot

func GetCodeSnapshot(ctx context.Context, projectId, id string) (*CodeSnapshot, error) {
	res, response, err := ApiClient.GetCodeSnapshot(ctx).GetCodeSnapshotParams(*tensorleapapi.NewGetCodeSnapshotParams(projectId, id)).Execute()
	if err = api.CheckRes(response, err); err != nil {
		return nil, err
	}
	return &res.CodeSnapshot, err
}

func PushCodeSnapshot(
	ctx context.Context, tarGzFile io.Reader, fileSize int64,
	entryFile, secretId, pythonVersion,
	versionName, projectId, branch string,
	overwriteVersionId string,
) (*tensorleapapi.PushCodeSnapshotResponse, error) {

	uploadUrl, err := GetCodeSnapshotUploadUrl(ctx, projectId)
	if err != nil {
		return nil, err
	}

	if err := UploadFile(uploadUrl, tarGzFile, fileSize); err != nil {
		return nil, err
	}

	saveCodeSnapshotParams := *tensorleapapi.NewPushCodeSnapshotParams(
		projectId,
		uploadUrl,
		entryFile,
		versionName,
	)

	if len(overwriteVersionId) > 0 {
		saveCodeSnapshotParams.SetOverwriteVersionId(overwriteVersionId)
	}

	if len(pythonVersion) > 0 {
		saveCodeSnapshotParams.GenericBaseImageType = &pythonVersion
	}

	if len(branch) > 0 {
		saveCodeSnapshotParams.SetBranchName(branch)
	}

	if len(secretId) > 0 {
		saveCodeSnapshotParams.SecretManagerId = &secretId
	}

	log.Info("Pushing code snapshot...")
	result, response, err := ApiClient.PushCodeSnapshot(ctx).
		PushCodeSnapshotParams(saveCodeSnapshotParams).
		Execute()
	if err = api.CheckRes(response, err); err != nil {
		return nil, err
	}

	return result, nil
}

const TIMEOUT_FOR_CODE_INTEGRATION_STATUS = 30 * time.Minute

func WaitForCodeIntegrationStatus(ctx context.Context, projectId, codeSnapshotId string) (bool, *CodeSnapshot, error) {
	log.Info("Waiting for code parser result...")
	sleepDuration := 3 * time.Second

	condition := func() (bool, []log.Step, error) {
		getJobParams := *tensorleapapi.NewGetJobsFilterParams()
		getJobParams.SetCodeSnapshotId(codeSnapshotId)
		// getJobParams.SetTypes([]tensorleapapi.JobType{tensorleapapi.JOBTYPE_DATASET_PARSE})
		codeIntegrationJobs, _, err := ApiClient.GetSlimJobs(ctx).GetJobsFilterParams(
			getJobParams,
		).Execute()
		if err != nil {
			return false, nil, fmt.Errorf("failed to get the code integration job: %v", err)
		}
		if len(codeIntegrationJobs.Jobs) == 0 {
			return false, nil, fmt.Errorf("failed to get the code integration job")
		}
		codeIntegrationJob := codeIntegrationJobs.Jobs[0]
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

	if err == ErrorTimeout {
		return false, nil, fmt.Errorf("timeout occurred while waiting for the integration code status")
	}
	if err != nil {
		return false, nil, fmt.Errorf("failed to wait for the integration code status: %v", err)
	}
	codeSnapshot, err := GetCodeSnapshot(ctx, projectId, codeSnapshotId)
	if err != nil {
		return false, nil, fmt.Errorf("failed to get the code integration version: %v", err)
	}
	return true, codeSnapshot, nil
}

func GetCodeSnapshotUploadUrl(ctx context.Context, projectId string) (string, error) {
	base_url := api.GetBaseUrlFromContext(ctx)
	params := *tensorleapapi.NewGetCodeSnapshotUploadUrlParams(projectId)
	params.SetOrigin(base_url)
	data, _, err := ApiClient.GetCodeSnapshotUploadUrl(ctx).
		GetCodeSnapshotUploadUrlParams(params).
		Execute()
	if err != nil {
		return "", err
	}
	return data.GetUrl(), nil
}
