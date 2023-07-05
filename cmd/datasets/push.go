package datasets

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	. "github.com/tensorleap/cli-go/pkg/api"
	"github.com/tensorleap/cli-go/pkg/config"
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
			defer cleanupTempFile(tarGzFile)

			if err := createTarGzFile(filePaths, tarGzFile); err != nil {
				return err
			}

			data, _, err := ApiClient.GetDatasetVersionUploadUrl(cmd.Context()).Execute()
			if err != nil {
				return err
			}

			tarGzFile.Seek(0, 0)
			uploadUrl := data.GetUrl()
			if err := uploadFile(uploadUrl, tarGzFile); err != nil {
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

func cleanupTempFile(file *os.File) {
	file.Close()
	os.Remove(file.Name())
}

func uploadFile(url string, file *os.File) error {
	fmt.Println("Uploading...")
	reader, writer := io.Pipe()
	go func() {
		defer writer.Close()
		_, err := io.Copy(writer, file)
		writer.CloseWithError(err)
	}()
	req, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func createTarGzFile(filePaths []string, file *os.File) error {
	fmt.Println("Packing files...")
	gzip := gzip.NewWriter(file)
	defer gzip.Close()
	tarWriter := tar.NewWriter(gzip)
	defer tarWriter.Close()

	for _, filePath := range filePaths {
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			return err
		}
		fmt.Println(filePath)
		header, err := tar.FileInfoHeader(fileInfo, fileInfo.Name())
		if err != nil {
			return err
		}
		header.Name = filePath
		err = tarWriter.WriteHeader(header)
		if err != nil {
			return err
		}
		fileContents, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		tarWriter.Write(fileContents)
	}

	return nil
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
