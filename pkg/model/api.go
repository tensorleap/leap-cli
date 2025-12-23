package model

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"time"

	"github.com/samber/lo"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/interactive_pages"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/notification"
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

func ImportModel(ctx context.Context, projectId, versionId string, modelInfo *tensorleapapi.ImportModelInfo, waitForResults bool) (jobId string, err error) {
	if modelInfo == nil {
		return "", fmt.Errorf("model info is required")
	}
	commandStartTime := time.Now()

	importModelParams := *tensorleapapi.NewImportNewModelParams(projectId, versionId, *modelInfo)

	importModelData, _, err := api.ApiClient.ImportModel(ctx).ImportNewModelParams(importModelParams).Execute()
	if err != nil {
		return "", fmt.Errorf("failed to import model: %v", err)
	}

	importModelJobId := importModelData.GetJobId()
	if waitForResults {
		okStatus, _, err := waitForImportModelJob(ctx, projectId, importModelJobId)
		if err == ErrJobFailed || !okStatus {
			report, collectErr := CollectImportModelJobErrors(ctx, projectId, importModelJobId, versionId, commandStartTime)
			if collectErr != nil {
				return importModelJobId, collectErr
			}
			runErr := interactive_pages.RunInteractivePages(report.ToReportPages())
			if runErr != nil {
				return importModelJobId, runErr
			}
			return importModelJobId, fmt.Errorf("import model failed see errors above, for more logs run: leap run logs %s", importModelJobId)
		}

		if err != nil {
			return importModelJobId, fmt.Errorf("error while waiting for import model job: %v", err)
		}
		log.Println("Model imported successfully")
		return importModelJobId, nil

	}

	log.Println("Starting import model job. JobId: ", importModelJobId)
	return importModelJobId, nil
}

func CollectImportModelJobErrors(ctx context.Context, projectId, jobId string, versionId string, commandStartTime time.Time) (*ImportModelErrorReport, error) {

	report := &ImportModelErrorReport{}
	var err error
	report.Notifications, err = notification.GetJobFailureNotifications(ctx, jobId, commandStartTime)
	if err != nil {
		log.Warnf("failed to print import model notifications: %v", err)
	}

	versions, err := GetVersions(ctx, projectId)
	if err != nil {
		return nil, fmt.Errorf("failed to get versions: %v", err)
	}
	currentVersion, found := lo.Find(versions, func(version tensorleapapi.SlimVersion) bool {
		return version.GetCid() == versionId
	})
	if !found {
		return nil, fmt.Errorf("failed to find version: %v", err)
	}

	report.ValidateAssetReport, err = GetGraphValidationErrors(currentVersion.GraphValidationData)
	if err != nil {
		return nil, fmt.Errorf("failed to print graph validation errors: %v", err)
	}
	report.TopLogs, err = getImportModelLogs(ctx, jobId)
	if err != nil {
		log.Errorf("failed to fetch import model top error logs: %v", err)
	}

	return report, nil
}

func getImportModelLogs(ctx context.Context, importModelJobId string) ([]string, error) {
	runLogs, err := run.GetRunLogs(ctx, importModelJobId)
	if err != nil {
		return nil, err
	}
	if len(runLogs) == 0 {
		return nil, nil
	}

	errorPattern := regexp.MustCompile(`(?i)"level_name":\s*"Error"|"level":\s*"error"|ERROR|Error:`)
	topLogs := run.GetTopLogs(runLogs, errorPattern, 5)

	return topLogs, nil
}

func GetGraphValidationErrors(validationData *tensorleapapi.GraphValidatorData) (*ValidateAssetReport, error) {
	if validationData == nil {
		return nil, nil
	}

	report := &ValidateAssetReport{}

	if validationData.GeneralError != nil {
		report.GeneralError = *validationData.GeneralError
	}

	addValidationError("losses", validatedLossNodesToInterface(validationData.Losses), report)
	addValidationError("metrics", validatedNodesToInterface(validationData.Metrics), report)
	addValidationError("custom_layers", validatedNodesToInterface(validationData.CustomLayers), report)
	addValidationError("ground_truths", validatedNodesToInterface(validationData.GroundTruths), report)
	addValidationError("prediction_types", validatedNodesToInterface(validationData.PredictionTypes), report)
	addValidationError("visualizers", validatedNodesToInterface(validationData.Visualizers), report)
	addValidationError("metadata", validatedNodesToInterface(validationData.Metadata), report)
	addValidationError("inputs", validatedNodesToInterface(validationData.Inputs), report)

	return report, nil
}

func validatedNodesToInterface(nodes []tensorleapapi.ValidatedNode) []ValidatedNode {
	result := make([]ValidatedNode, len(nodes))
	for i := range nodes {
		result[i] = &nodes[i]
	}
	return result
}

func validatedLossNodesToInterface(nodes []tensorleapapi.ValidatedLossNode) []ValidatedNode {
	result := make([]ValidatedNode, len(nodes))
	for i := range nodes {
		result[i] = &nodes[i]
	}
	return result
}

type ValidatedNode interface {
	GetError() string
	GetName() string
	GetNodeId() string
	GetConnectionName() string
}

func addValidationError(name string, nodes []ValidatedNode, report *ValidateAssetReport) bool {

	hasErrors := nodes == nil || lo.SomeBy(nodes, func(node ValidatedNode) bool {
		return node.GetError() != ""
	})
	if !hasErrors {
		return false
	}
	for _, node := range nodes {
		if node.GetError() == "" {
			continue
		}
		validationError := ValidateAssetError{
			Error:          node.GetError(),
			Name:           node.GetName(),
			NodeId:         node.GetNodeId(),
			ConnectionName: node.GetConnectionName(),
			Type:           name,
		}
		report.Nodes = append(report.Nodes, validationError)
	}
	return true
}

func OverrideModel(ctx context.Context, projectId, versionId string, waitForResults bool, modelInfo *tensorleapapi.ImportModelInfo) (jobId string, err error) {
	commandStartTime := time.Now()
	params := *tensorleapapi.NewOverwriteModelParams(projectId, versionId)
	if modelInfo != nil {
		params.Model = modelInfo
	}
	overrideModelData, _, err := api.ApiClient.OverwriteModel(ctx).OverwriteModelParams(params).Execute()
	if err != nil {
		return "", err
	}
	overrideModelJobId := overrideModelData.GetJobId()
	if !waitForResults {
		return overrideModelJobId, nil
	}
	okStatus, _, err := waitForImportModelJob(ctx, projectId, overrideModelJobId)
	if err != nil {
		return overrideModelJobId, err
	}
	if okStatus {
		log.Println("Successfully overridden model")
	} else {
		report, err := CollectImportModelJobErrors(ctx, projectId, overrideModelJobId, versionId, commandStartTime)
		if err != nil {
			return overrideModelJobId, err
		}
		report.Show()
	}
	return overrideModelJobId, nil
}

const TIMEOUT_FOR_IMPORT_MODEL_JOB = 30 * time.Minute

var ErrJobFailed = fmt.Errorf("import model job failed")

var ErrImportModelTimeout = fmt.Errorf("timeout occurred while waiting for import model job status")

func waitForImportModelJob(ctx context.Context, projectId, importModelJobId string) (ok bool, job *tensorleapapi.Job, err error) {
	fmt.Println("Waiting for import model result...")
	sleepDuration := 3 * time.Second

	condition := func() (bool, []log.Step, error) {
		job, err := GetImportModelJob(ctx, importModelJobId, projectId)
		if err != nil {
			return false, nil, fmt.Errorf("failed to wait for the import model job status: %v", err)
		}
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
		return false, nil, ErrImportModelTimeout
	}
	if err != nil {
		return false, nil, err
	}

	return true, job, nil
}

func GetImportModelJob(ctx context.Context, jobId string, projectId string) (*tensorleapapi.Job, error) {
	getJobParams := *tensorleapapi.NewGetJobsFilterParams()
	getJobParams.SetProjectId(projectId)
	getJobParams.SetCid([]string{jobId})
	data, _, err := api.ApiClient.GetSlimJobs(ctx).GetJobsFilterParams(getJobParams).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get job: %v", err)
	}
	if len(data.Jobs) == 0 {
		return nil, fmt.Errorf("failed to get job")
	}
	return &data.Jobs[0], nil
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
