package model

import (
	"context"
	"os"
	"path/filepath"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func ImportModel(ctx context.Context, filePath, projectId, modelName, versionName, modelType, branchName, datasetId string) error {
	fileName := filepath.Base(filePath)
	tempSignedUploadUrlParams := *tensorleapapi.NewGetUploadSignedUrlParams(fileName)
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
	if err := api.UploadFile(uploadUrl, file); err != nil {
		return err
	}

	uploadFileName := signedUrlData.GetFileName()
	importModelParams := *tensorleapapi.NewImportNewModelParams(projectId, uploadFileName, modelName, versionName, tensorleapapi.ImportModelType(modelType))
	if len(branchName) > 0 {
		importModelParams.BranchName = &branchName
	}
	if len(datasetId) > 0 {
		importModelParams.DatasetId = &datasetId
	}
	importModelData, _, err := api.ApiClient.ImportModel(ctx).ImportNewModelParams(importModelParams).Execute()
	if err != nil {
		return err
	}

	log.Println("Starting import model job. JobId: ", importModelData.GetImportModelJobId())
	return nil
}
