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
		Use:   "save",
		Short: "Save dataset script",
		Long:  `Save dataset script`,
		Run: func(cmd *cobra.Command, args []string) {
			datasetConfig, err := config.GetDatasetConfig()
			cobra.CheckErr(err)
			config.VerifyLoggedIn()

			filePaths, err := getDatasetFiles(datasetConfig)
			cobra.CheckErr(err)

			tarGzFile, err := os.CreateTemp("", "tensorleap-*.tar.gz")
			cobra.CheckErr(err)
			defer cleanupTempFile(tarGzFile)

			err = createTarGzFile(filePaths, tarGzFile)
			cobra.CheckErr(err)

			data, _, err := ApiClient.GetDatasetVersionUploadUrl(cmd.Context()).Execute()
			cobra.CheckErr(err)

			tarGzFile.Seek(0, 0)
			uploadUrl := data.GetUrl()
			err = uploadFile(uploadUrl, tarGzFile)
			cobra.CheckErr(err)

			fmt.Println("Creating new dataset version...")
			_, _, err = ApiClient.SaveDatasetVersion(cmd.Context()).
				SaveDatasetVersionParams(*tensorleapapi.NewSaveDatasetVersionParams(
					datasetConfig.DatasetId,
					uploadUrl,
					datasetConfig.EntryFile,
				)).
				Execute()
			cobra.CheckErr(err)

			fmt.Println("Done!")
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
		cobra.CheckErr(err)
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
