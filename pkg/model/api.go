package model

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/log"
	tlApi "github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func ImportModel(ctx context.Context, filePath, projectId, message, modelType, branchName, datasetId, codeBranch string, transformInput bool, waitForResults bool) error {
	fileName := filepath.Base(filePath)
	versionName := message
	modelName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	tempSignedUploadUrlParams := *tlApi.NewGetUploadSignedUrlParams(fileName)
	signedUrlData, _, err := api.ApiClient.GetUploadSignedUrl(ctx).
		GetUploadSignedUrlParams(tempSignedUploadUrlParams).Execute()
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
	importModelParams := *tlApi.NewImportNewModelParams(projectId, uploadFileName, modelName, versionName, tlApi.ImportModelType(modelType))
	if len(branchName) > 0 {
		importModelParams.BranchName = &branchName
	}

	var mappingYaml string = ""
	if len(datasetId) > 0 {
		importModelParams.DatasetId = &datasetId
		mappingYaml = code.GetDatasetMappingYaml(ctx, datasetId, codeBranch)
		importModelParams.MappingYaml = &mappingYaml
	}
	if transformInput && modelType == string(tlApi.IMPORTMODELTYPE_ONNX) {
		importModelParams.TransformInputs = &transformInput
	}
	if codeBranch != "" {
		importModelParams.SetCodeIntegrationBranch(codeBranch)
	}
	importModelData, _, err := api.ApiClient.ImportModel(ctx).ImportNewModelParams(importModelParams).Execute()
	if err != nil {
		return fmt.Errorf("failed to import model: %v", err)
	}

	importModelJobId := importModelData.GetImportModelJobId()
	if waitForResults {
		okStatus, job, err := waitForImportModelJob(ctx, projectId, importModelJobId)
		if err != nil {
			return err
		}
		if okStatus {
			log.Println("Successfully imported model")

			displayMappingErrorsIfMappingProvided := mappingYaml != ""
			if displayMappingErrorsIfMappingProvided {
				printMappingValidationErrors(ctx, job.GetVersion(), projectId, mappingYaml)
			}
		} else {
			return fmt.Errorf("failed to import model")
		}
		return nil
	}

	log.Println("Starting import model job. JobId: ", importModelJobId)
	return nil
}

func waitForImportModelJob(ctx context.Context, projectId, importModelJobId string) (ok bool, job *tlApi.Job, err error) {
	message := "Waiting for import model result..."
	sleepDuration := 3 * time.Second
	timeoutDuration := 10 * time.Minute

	getJobParams := *tlApi.NewGetJobsFilterParams()
	getJobParams.SetProjectId(projectId)

	condition := func() (bool, error) {
		data, _, err := api.ApiClient.GetSlimJobs(ctx).GetJobsFilterParams(
			getJobParams,
		).Execute()
		if err != nil {
			return false, fmt.Errorf("failed to wait for the import model job status: %v", err)
		}
		if len(data.Jobs) == 0 {
			return false, fmt.Errorf("failed to wait for the import model job status")
		}

		job = findRunProcessByJobId(data.Jobs, importModelJobId)
		switch job.Status {
		case tlApi.JOBSTATUS_FAILED:
			ok = true
			return true, nil
		case tlApi.JOBSTATUS_FINISHED:
			ok = true
			return true, nil
		}

		return false, nil
	}

	err = api.WaitForCondition(ctx, message, condition, sleepDuration, timeoutDuration)

	if err == api.ErrorTimeout {
		return false, nil, fmt.Errorf("timeout occurred while waiting for import job status")
	}
	if err != nil {
		return false, nil, fmt.Errorf("failed to wait for import job result status: %v", err)
	}

	return
}

func findRunProcessByJobId(runProcesses []tlApi.Job, jobId string) *tlApi.Job {
	for _, rp := range runProcesses {
		if rp.GetCid() == jobId {
			return &rp
		}
	}
	return nil
}

func printMappingValidationErrors(ctx context.Context, versionId string, projectId string, mapping string) {
	getValidationErrorsParams := *tlApi.NewCodeIntegrationMappingErrorsParams(projectId, versionId, mapping)

	data, _, err := api.ApiClient.GetCodeIntegrationMappingErrorsByVersionId(ctx).CodeIntegrationMappingErrorsParams(getValidationErrorsParams).Execute()
	if err != nil {
		log.Println("failed getting mapping validation errors:", err)
		return
	}
	if data != nil && data.MappingErrors != nil && len(data.MappingErrors) > 0 {
		log.Println("encountered the following validation errors while applying mapping to graph:")

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Title", "Message"})
		table.SetBorder(true)
		table.SetRowLine(true)

		for _, entry := range data.MappingErrors {
			table.Append([]string{entry.Title, entry.Message})
		}
		table.Render()

	} else {
		log.Println("mapping was applied successfully applied with no validation errors")
	}
}
