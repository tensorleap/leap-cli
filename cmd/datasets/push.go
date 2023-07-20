package datasets

import (
	"context"
	"fmt"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
	"github.com/tensorleap/cli-go/pkg/auth"
	. "github.com/tensorleap/cli-go/pkg/api"
	"github.com/tensorleap/cli-go/pkg/datasets"
	. "github.com/tensorleap/cli-go/pkg/local"
	"github.com/tensorleap/cli-go/pkg/log"
	"github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func init() {
	RootCommand.AddCommand(&cobra.Command{
		Use:   "push",
		Short: "Push dataset script",
		Long:  `Push dataset script`,
		RunE: func(cmd *cobra.Command, args []string) error {
			datasetConfig, err := datasets.GetDatasetConfig()
			if err != nil {
				return err
			}
			if err := auth.CheckLoggedIn(); err != nil {
				return err
			}
			ctx := cmd.Context()

			filePaths, err := getDatasetFiles(datasetConfig)
			if err != nil {
				return err
			}

			tarGzFile, err := os.CreateTemp("", "tensorleap-*.tar.gz")
			if err != nil {
				return err
			}
			defer CleanupTempFile(tarGzFile)

			if err := CreateTarGzFile(filePaths, tarGzFile); err != nil {
				return err
			}

			data, _, err := ApiClient.GetDatasetVersionUploadUrl(ctx).Execute()
			if err != nil {
				return err
			}

			tarGzFile.Seek(0, 0)
			uploadUrl := data.GetUrl()
			if err := UploadFile(uploadUrl, tarGzFile); err != nil {
				return err
			}

			if err := addDatasetIfNotExisted(ctx, datasetConfig); err != nil {
				return err
			}

			log.Info("Creating new dataset version...")
			if _, _, err = ApiClient.SaveDatasetVersion(ctx).
				SaveDatasetVersionParams(*tensorleapapi.NewSaveDatasetVersionParams(
					datasetConfig.DatasetId,
					uploadUrl,
					datasetConfig.EntryFile,
				)).
				Execute(); err != nil {
				return err
			}

			log.Info("Done!")
			return nil
		},
	})
}

func getDatasetFiles(datasetConfig *datasets.DatasetConfig) ([]string, error) {
	currentDirFs := os.DirFS(".")
	var allMatchedFiles []string
	for _, pattern := range datasetConfig.IncludePatterns {
		matches, err := fs.Glob(currentDirFs, pattern)
		if err != nil {
			return nil, err
		}
		allMatchedFiles = append(allMatchedFiles, matches...)
	}
	return allMatchedFiles, nil
}

func addDatasetIfNotExisted(ctx context.Context, datasetConfig *datasets.DatasetConfig) error {
	datasetsRes, _, err := ApiClient.GetDatasets(ctx).Execute()
	if err != nil {
		return fmt.Errorf("Failed to get datasets: %v", err)
	}
	isDatasetExisted := false
	for _, dataset := range datasetsRes.Datasets {
		if dataset.Cid == datasetConfig.DatasetId {
			isDatasetExisted = true
			break
		}
	}

	if !isDatasetExisted {
		log.Infof("Not found dataset id: %s. Creating new dataset", datasetConfig.DatasetId)

		name := datasets.AskForDatasetName()
		dataset, err := datasets.CreateNewDataset(ctx, name)
		if err != nil {
			return err
		}

		datasetConfig.DatasetId = dataset.GetCid()

		err = datasets.SetDatasetConfig(datasetConfig, "")
		if err != nil {
			return fmt.Errorf("Failed to update tensorleap config: %v", err)
		}
	}
	return nil
}
