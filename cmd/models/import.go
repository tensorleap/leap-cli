package models

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	. "github.com/tensorleap/cli-go/pkg/api"
	"github.com/tensorleap/cli-go/pkg/auth"
	"github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func init() {
	var projectId string
	var modelName string
	var versionName string
	var modelType string
	var branchName string
	var datasetId string

	var cmd = &cobra.Command{
		Use:   "import [filePath]",
		Short: "Import a model into tensorleap",
		Long:  `Import a model into tensorleap`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := auth.CheckLoggedIn(); err != nil {
				return err
			}
			filePath := args[0]
			fileName := filepath.Base(filePath)
			tempSignedUploadUrlParams := *tensorleapapi.NewGetUploadSignedUrlParams(fileName)
			signedUrlData, _, err := ApiClient.GetUploadSignedUrl(cmd.Context()).GetUploadSignedUrlParams(tempSignedUploadUrlParams).Execute()
			if err != nil {
				return err
			}
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			uploadUrl := signedUrlData.GetUrl()
			if err := UploadFile(uploadUrl, file); err != nil {
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
			importModelData, _, err := ApiClient.ImportModel(cmd.Context()).ImportNewModelParams(importModelParams).Execute()
			if err != nil {
				return err
			}

			fmt.Println("Starting import model job. JobId: ", importModelData.GetImportModelJobId())
			return nil
		},
	}

	cmd.Flags().StringVar(&projectId, "projectId", "", "ProjectId is the id of the project the model will be imported to")
	cmd.MarkFlagRequired("projectId")

	cmd.Flags().StringVar(&modelName, "name", "", "Name of the model that will be created")
	cmd.MarkFlagRequired("name")

	cmd.Flags().StringVar(&versionName, "version", "", "Version is the name of the version that will be created for the model")
	cmd.MarkFlagRequired("version")

	cmd.Flags().StringVar(&modelType, "type", "", "Type is the type of the model file [JSON_TF2 / ONNX / PB_TF2 / H5_TF2]")
	cmd.MarkFlagRequired("type")

	cmd.Flags().StringVar(&branchName, "branch", "", "Branch is the name of the branch [OPTIONAL]")
	cmd.Flags().StringVar(&datasetId, "datasetId", "", "This is a datasetId (Will use the last valid dataset version)")

	RootCommand.AddCommand(cmd)
}
