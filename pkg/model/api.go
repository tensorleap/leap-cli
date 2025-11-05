package model

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/run"
	tlApi "github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func ImportModel(ctx context.Context, filePath, projectId, modelType, versionId string, transformInput bool, waitForResults bool) error {
	fileName := filepath.Base(filePath)
	signedUrlData, err := api.GetUploadSignedUrl(ctx, fileName)
	if err != nil {
		return err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	uploadUrl := signedUrlData.GetUrl()
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	if err := api.UploadFile(uploadUrl, file, fileInfo.Size()); err != nil {
		return err
	}

	uploadFileName := signedUrlData.GetFileName()
	modelInfo := *tlApi.NewImportModelInfo(uploadFileName, tlApi.ImportModelType(modelType))

	if transformInput && modelType == string(tlApi.IMPORTMODELTYPE_ONNX) {
		modelInfo.TransformInputs = &transformInput
	}
	importModelParams := *tlApi.NewImportNewModelParams(projectId, versionId, modelInfo)

	importModelData, _, err := api.ApiClient.ImportModel(ctx).ImportNewModelParams(importModelParams).Execute()
	if err != nil {
		return fmt.Errorf("failed to import model: %v", err)
	}

	importModelJobId := importModelData.GetImportModelJobId()
	if waitForResults {
		okStatus, _, err := waitForImportModelJob(ctx, projectId, importModelJobId)
		if err != nil {
			return err
		}
		if okStatus {
			log.Println("Successfully imported model")

		} else {
			topLogs, err := run.GetRunLogs(ctx, importModelJobId)
			if err != nil {
				return err
			}
			logs := run.GetTopLogs(topLogs, "import-model", 40)
			if logs == "" {
				logs = "-- logs not found --"
			}
			fmt.Printf("%s\n", logs)
			return fmt.Errorf("failed to import model show the logs above for more details run `leap runs logs %s` to view the full logs", importModelJobId)
		}
		return nil
	}

	log.Println("Starting import model job. JobId: ", importModelJobId)
	return nil
}

const TIMEOUT_FOR_IMPORT_MODEL_JOB = 30 * time.Minute

func waitForImportModelJob(ctx context.Context, projectId, importModelJobId string) (ok bool, job *tlApi.Job, err error) {
	fmt.Println("Waiting for import model result...")
	sleepDuration := 3 * time.Second

	getJobParams := *tlApi.NewGetJobsFilterParams()
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
			return false, steps, fmt.Errorf("import model job failed")
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

func findRunProcessByJobId(runProcesses []tlApi.Job, jobId string) *tlApi.Job {
	for _, rp := range runProcesses {
		if rp.GetCid() == jobId {
			return &rp
		}
	}
	return nil
}
