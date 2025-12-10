package model

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"time"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/run"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func PrepareImportModelFromFilePath(ctx context.Context, filePath string, transformInput bool, modelType string) (*tensorleapapi.ImportModelInfo, error) {
	fileName := filepath.Base(filePath)
	signedUrlData, err := api.GetUploadSignedUrl(ctx, fileName)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()

	uploadUrl := signedUrlData.GetUrl()
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if err := api.UploadFile(uploadUrl, file, fileInfo.Size()); err != nil {
		return nil, err
	}

	uploadFileName := signedUrlData.GetFileName()
	modelInfo := tensorleapapi.NewImportModelInfo(uploadFileName, tensorleapapi.ImportModelType(modelType))
	if transformInput && modelType == string(tensorleapapi.IMPORTMODELTYPE_ONNX) {
		modelInfo.TransformInputs = &transformInput
	}
	return modelInfo, nil
}

func ImportModel(ctx context.Context, projectId, versionId string, modelInfo *tensorleapapi.ImportModelInfo, waitForResults bool) error {
	if modelInfo == nil {
		return fmt.Errorf("model info is required")
	}

	importModelParams := *tensorleapapi.NewImportNewModelParams(projectId, versionId, *modelInfo)

	importModelData, _, err := api.ApiClient.ImportModel(ctx).ImportNewModelParams(importModelParams).Execute()
	if err != nil {
		return fmt.Errorf("failed to import model: %v", err)
	}

	importModelJobId := importModelData.GetJobId()
	if waitForResults {
		okStatus, _, err := waitForImportModelJob(ctx, projectId, importModelJobId)
		if err == ErrJobFailed || !okStatus {
			topLogs, err := GetTopLogs(ctx, importModelJobId)
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", topLogs)
			return fmt.Errorf("failed to push model show the logs above for more details run `leap runs logs %s` to view the full logs", importModelJobId)
		}
		if err != nil {
			return fmt.Errorf("error while waiting for import model job: %v", err)
		}
		log.Println("Model imported successfully")
		return nil

	}

	log.Println("Starting import model job. JobId: ", importModelJobId)
	return nil
}

func GetTopLogs(ctx context.Context, jobId string) (string, error) {
	jobLogs, err := run.GetRunLogs(ctx, jobId)
	if err != nil {
		return "", err
	}
	topLogs := run.GetTopLogs(jobLogs, "import-model", 40)
	if topLogs == "" {
		topLogs = "-- logs not found --"
	}
	return topLogs, nil
}

func OverrideModel(ctx context.Context, projectId, versionId string, waitForResults bool, modelInfo *tensorleapapi.ImportModelInfo) error {
	params := *tensorleapapi.NewOverwriteModelParams(projectId, versionId)
	if modelInfo != nil {
		params.Model = modelInfo
	}
	overrideModelData, _, err := api.ApiClient.OverwriteModel(ctx).OverwriteModelParams(params).Execute()
	if err != nil {
		return err
	}
	if !waitForResults {
		return nil
	}
	overrideModelJobId := overrideModelData.GetJobId()
	okStatus, _, err := waitForImportModelJob(ctx, projectId, overrideModelJobId)
	if err != nil {
		return err
	}
	if okStatus {
		log.Println("Successfully overridden model")
	} else {
		topLogs, err := GetTopLogs(ctx, overrideModelJobId)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", topLogs)
		return fmt.Errorf("failed to override model show the logs above for more details run `leap runs logs %s` to view the full logs", overrideModelJobId)
	}
	return nil
}

const TIMEOUT_FOR_IMPORT_MODEL_JOB = 30 * time.Minute

var ErrJobFailed = fmt.Errorf("import model job failed")

func waitForImportModelJob(ctx context.Context, projectId, importModelJobId string) (ok bool, job *tensorleapapi.Job, err error) {
	fmt.Println("Waiting for import model result...")
	sleepDuration := 3 * time.Second

	getJobParams := *tensorleapapi.NewGetJobsFilterParams()
	getJobParams.SetProjectId(projectId)
	getJobParams.SetCid([]string{importModelJobId})

	condition := func() (bool, []log.Step, error) {
		data, _, err := api.ApiClient.GetSlimJobs(ctx).GetJobsFilterParams(
			getJobParams,
		).Execute()
		if err != nil {
			return false, nil, fmt.Errorf("failed to wait for the import model job status: %v", err)
		}
		if len(data.Jobs) == 0 {
			return false, nil, fmt.Errorf("failed to wait for the import model job status")
		}

		job = findRunProcessByJobId(data.Jobs, importModelJobId)
		steps := api.StepsFromJob(job)
		switch true {
		case api.IsJobFailed(job.Status):
			return false, steps, ErrJobFailed
		case api.IsJobFinished(job.Status):
			return true, steps, nil
		}

		return false, steps, nil
	}

	err = api.WaitForConditionWithSteps(ctx, condition, sleepDuration, TIMEOUT_FOR_IMPORT_MODEL_JOB)

	if err == api.ErrorTimeout {
		return false, nil, fmt.Errorf("timeout occurred while waiting for import job status")
	}
	if err != nil {
		return false, nil, err
	}

	return true, job, nil
}

func findRunProcessByJobId(runProcesses []tensorleapapi.Job, jobId string) *tensorleapapi.Job {
	for _, rp := range runProcesses {
		if rp.GetCid() == jobId {
			return &rp
		}
	}
	return nil
}

func GetVersions(ctx context.Context, projectId string) ([]tensorleapapi.SlimVersion, error) {
	versions, _, err := api.ApiClient.GetProjectSlimVersions(ctx).GetProjectVersionsParams(*tensorleapapi.NewGetProjectVersionsParams(projectId)).Execute()
	if err != nil {
		return nil, err
	}
	return versions.Versions, nil
}

func GetSessionRunsEvaluate(ctx context.Context, projectId string, sessionIds []string) (*tensorleapapi.GetSessionRunsEvaluateResponse, error) {
	params := tensorleapapi.GetSessionRunsEvaluateParams{
		ProjectId: projectId,
	}
	result, res, err := api.ApiClient.GetSessionRunsEvaluate(ctx).GetSessionRunsEvaluateParams(params).Execute()
	if err = api.CheckRes(res, err); err != nil {
		return nil, err
	}
	if len(sessionIds) > 0 {
		result.EvaluateSessionRuns = filterSessionRunsBySessionId(result.EvaluateSessionRuns, sessionIds)
	}
	return result, nil
}

func filterSessionRunsBySessionId(sessionRuns []tensorleapapi.SessionRunData, sessionIds []string) []tensorleapapi.SessionRunData {
	var filteredSessionRuns []tensorleapapi.SessionRunData
	for _, sessionRun := range sessionRuns {
		if slices.Contains(sessionIds, sessionRun.GetSessionId()) {
			filteredSessionRuns = append(filteredSessionRuns, sessionRun)
		}
	}

	return filteredSessionRuns
}
