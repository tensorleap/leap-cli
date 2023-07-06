package datasets

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
	. "github.com/tensorleap/cli-go/pkg/api"
	"github.com/tensorleap/cli-go/pkg/config"
	. "github.com/tensorleap/cli-go/pkg/local"
	"github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func init() {
	RootCommand.AddCommand(&cobra.Command{
		Use:   "push",
		Short: "Push dataset script",
		Long:  `Push dataset script`,
		RunE: func(cmd *cobra.Command, args []string) error {
			datasetConfig, err := config.GetDatasetConfig()
			if err != nil {
				return err
			}
			if err := config.CheckLoggedIn(); err != nil {
				return err
			}

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

			data, _, err := ApiClient.GetDatasetVersionUploadUrl(cmd.Context()).Execute()
			if err != nil {
				return err
			}

			tarGzFile.Seek(0, 0)
			uploadUrl := data.GetUrl()
			if err := UploadFile(uploadUrl, tarGzFile); err != nil {
				return err
			}

			fmt.Println("Creating new dataset version...")
			if _, _, err = ApiClient.SaveDatasetVersion(cmd.Context()).
				SaveDatasetVersionParams(*tensorleapapi.NewSaveDatasetVersionParams(
					datasetConfig.DatasetId,
					uploadUrl,
					datasetConfig.EntryFile,
				)).
				Execute(); err != nil {
				return err
			}

			fmt.Println("Done!")
			return nil
		},
	})
}

func getDatasetFiles(datasetConfig *config.DatasetConfig) ([]string, error) {
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
