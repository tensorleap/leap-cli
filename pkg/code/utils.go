package code

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/local"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

var ErrEmptyCodeIntegrationVersion = fmt.Errorf("CodeIntegration is empty")

const BindingFilePath = "leap_mapping.yaml"

func CreateCodeIntegration(ctx context.Context, codeIntegrations []CodeIntegration) (*CodeIntegration, error) {

	name, err := AskForCodeIntegrationName(codeIntegrations)
	if err != nil {
		return nil, err
	}
	return AddCodeIntegration(ctx, name)
}

func GetCodeIntegrationFromFlag(ctx context.Context, codeIntegrationIdFlag string, askForNewProjectFirst bool) (*CodeIntegration, error) {
	codeIntegrations, err := GetCodeIntegrations(ctx)
	if err != nil {
		return nil, err
	}
	var selected *CodeIntegration
	if len(codeIntegrationIdFlag) > 0 {
		selected, err = entity.GetEntityById(codeIntegrationIdFlag, codeIntegrations, CodeIntegrationEntityDesc)
	} else {
		selected, err = SelectOrCreateCodeIntegration(ctx, codeIntegrations, askForNewProjectFirst)
	}
	if err != nil {
		return nil, err
	}
	return selected, nil
}

func SelectOrCreateCodeIntegration(ctx context.Context, codeIntegrations []CodeIntegration, askIsCreateNewFirst bool) (*tensorleapapi.Dataset, error) {
	createCodeIntegration := func() (*CodeIntegration, error) {
		return CreateCodeIntegration(ctx, codeIntegrations)
	}
	return entity.SelectEntityOrCreateOne(codeIntegrations, createCodeIntegration, askIsCreateNewFirst, CodeIntegrationEntityDesc)
}

func AskForCodeIntegrationName(codeIntegrations []CodeIntegration) (name string, err error) {

	existingNames := entity.GetNames(codeIntegrations, CodeIntegrationEntityDesc)

	name, err = entity.AskForName(existingNames, CodeIntegrationEntityDesc)
	return
}

func isDatasetVersionEmpty(datasetVersion *tensorleapapi.DatasetVersion) bool {
	return len(datasetVersion.GetBlobPath()) == 0
}

func CloneCodeIntegrationVersion(ctx context.Context, codeIntegrationVersion *tensorleapapi.DatasetVersion, outputDir string) ([]string, error) {
	if isDatasetVersionEmpty(codeIntegrationVersion) {
		return []string{}, ErrEmptyCodeIntegrationVersion
	}
	blobPath := codeIntegrationVersion.GetBlobPath()
	res, _, err := api.ApiClient.GetDownloadSignedUrl(ctx).
		GetDownloadSignedUrlParams(*tensorleapapi.NewGetDownloadSignedUrlParams(blobPath)).
		Execute()
	if err != nil {
		return nil, err
	}

	files, err := local.DownloadAndExtractTarFile(res.GetUrl(), outputDir)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func FetchFileFromTarGz(blobURL string, filename string) (string, error) {
	resp, err := http.Get(blobURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return "", err
	}
	defer gz.Close()

	tarReader := tar.NewReader(gz)
	for {
		header, err := tarReader.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return "", err
		}

		if header.Name == filename {
			var buf bytes.Buffer
			if _, err := io.Copy(&buf, tarReader); err != nil {
				return "", err
			}
			return buf.String(), nil
		}
	}

	return "", errors.New("file not found in tar.gz")
}

func BundleCodeIntoTempFile(filesDir string, workspaceConfig *workspace.WorkspaceConfig) (close func(), tarGzFile *os.File, err error) {
	filePaths, err := getDatasetFiles(filesDir, workspaceConfig)
	if err != nil {
		return
	}

	tarGzFile, err = os.CreateTemp("", "tensorleap-*.tar.gz")
	if err != nil {
		return
	}
	close = func() { local.CleanupTempFile(tarGzFile) }

	if err = local.CreateTarGzFile(filesDir, filePaths, tarGzFile); err != nil {
		close()
		return
	}
	_, err = tarGzFile.Seek(0, 0)

	return
}

func getDatasetFiles(filesDir string, workspaceConfig *workspace.WorkspaceConfig) ([]string, error) {
	currentDirFs := os.DirFS(filesDir)
	var allMatchedFiles []string
	allFilePaths := append(workspaceConfig.IncludePatterns, BindingFilePath)
	for _, pattern := range allFilePaths {
		matches, err := fs.Glob(currentDirFs, pattern)
		if err != nil {
			return nil, err
		}
		allMatchedFiles = append(allMatchedFiles, matches...)
	}
	return allMatchedFiles, nil
}

func GetAndUpdateCodeIntegrationIfNotExists(ctx context.Context, workspaceConfig *workspace.WorkspaceConfig) (*CodeIntegration, error) {
	codeIntegrations, err := GetCodeIntegrations(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get code integration: %v", err)
	}
	for _, codeIntegration := range codeIntegrations {
		if codeIntegration.Cid == workspaceConfig.CodeIntegrationId {
			return &codeIntegration, nil
		}
	}

	log.Infof("Not found code integration id: %s. Select or create new code integration", workspaceConfig.CodeIntegrationId)

	codeIntegration, err := SelectOrCreateCodeIntegration(ctx, codeIntegrations, true)
	if err != nil {
		return nil, err
	}

	workspaceConfig.CodeIntegrationId = codeIntegration.GetCid()
	log.Info("Updating codeIntegrationId")
	err = workspace.SetWorkspaceConfig(workspaceConfig, "")
	if err != nil {
		return nil, err
	}
	return codeIntegration, nil
}

func PrintCodeIntegrationVersionParserErr(civ *CodeIntegrationVersion) {
	log.Error("Code parsing failed, see error below:")
	fmt.Println(*civ.Metadata.SetupStatus.GeneralError)
	fmt.Println(*civ.Metadata.SetupStatus.PrintLog)
}

func PushCode(ctx context.Context, force bool, codeIntegrationId string, tarGzFile *os.File, entryFile, secretId string) (pushed bool, current *CodeIntegrationVersion, err error) {
	if !force {
		log.Info("Checking if code has changed")

		latestVersion, err := GetLatestVersion(ctx, codeIntegrationId)

		if err != nil {
			log.Warnf("Failed to get latest code integration version: %v", err)
		} else {
			change, err := CompareCodeVersion(ctx, latestVersion, tarGzFile, entryFile, secretId)
			if err != nil {
				log.Warnf("Failed to check if code changed: %v", err)
			}
			if !change {
				log.Info("No change in code, skipping push")
				return false, latestVersion, nil
			} else {
				log.Info("Code changed, pushing new version")
			}
		}
	}
	codeIntegrationVersion, err := AddCodeIntegrationVersion(ctx, tarGzFile, codeIntegrationId, entryFile, secretId)
	if err != nil {
		return false, nil, err
	}
	return true, codeIntegrationVersion, nil
}

func CompareCodeVersion(ctx context.Context, compareVersion *CodeIntegrationVersion, tarGzFile *os.File, entryFile, secretId string) (bool, error) {

	if isDatasetVersionEmpty(compareVersion) {
		return true, nil
	}

	if compareVersion.Metadata.GetSecretManagerId() != secretId || compareVersion.CodeEntryFile != entryFile {
		return true, nil
	}

	latestVersionBlobPath := compareVersion.GetBlobPath()
	tempLatestVersionFile, err := os.CreateTemp("", "tensorleap-*.tar.gz")
	if err != nil {
		return true, fmt.Errorf("failed to create temp file: %v", err)

	}
	defer local.CleanupTempFile(tempLatestVersionFile)

	res, _, err := api.ApiClient.GetDownloadSignedUrl(ctx).
		GetDownloadSignedUrlParams(*tensorleapapi.NewGetDownloadSignedUrlParams(latestVersionBlobPath)).
		Execute()
	if err != nil {
		return true, fmt.Errorf("failed to get download signed url: %v", err)
	}
	err = api.DownloadFile(res.GetUrl(), tempLatestVersionFile)
	if err != nil {
		return true, fmt.Errorf("failed to download latest code integration version: %v", err)
	}
	_, err = tempLatestVersionFile.Seek(0, 0)
	if err != nil {
		return true, fmt.Errorf("failed to seek tempLatestVersionFile: %v", err)
	}
	latestChecksum, err := local.GetFileChecksum(tempLatestVersionFile)
	if err != nil {
		return true, fmt.Errorf("failed to get checksum of latest code integration version: %v", err)
	}
	newChecksum, err := local.GetFileChecksum(tarGzFile)

	if err != nil {
		return true, fmt.Errorf("failed to get checksum of new code integration version: %v", err)
	}
	_, err = tarGzFile.Seek(0, 0)
	if err != nil {
		return true, fmt.Errorf("failed to seek tarGzFile: %v", err)
	}

	return latestChecksum != newChecksum, nil
}

func IsCodeParseFailed(codeIntegrationVersion *CodeIntegrationVersion) bool {
	return codeIntegrationVersion.TestStatus == tensorleapapi.TESTSTATUS_TEST_FAIL
}

func IsCodeParsing(codeIntegrationVersion *CodeIntegrationVersion) bool {
	return codeIntegrationVersion.TestStatus == tensorleapapi.TESTSTATUS_DURING_TEST || codeIntegrationVersion.TestStatus == tensorleapapi.TESTSTATUS_BEFORE_TEST
}

func GetDatasetMappingYaml(ctx context.Context, codeIntegrationId string) string {
	codeIntegrationVersion, err := GetLatestVersion(ctx, codeIntegrationId)
	if err != nil {
		return ""
	}

	if isDatasetVersionEmpty(codeIntegrationVersion) {
		return ""
	}

	blobPath := codeIntegrationVersion.GetBlobPath()
	res, _, err := api.ApiClient.GetDownloadSignedUrl(ctx).
		GetDownloadSignedUrlParams(*tensorleapapi.NewGetDownloadSignedUrlParams(blobPath)).
		Execute()
	if err != nil {
		return ""
	}

	content, err := FetchFileFromTarGz(res.Url, BindingFilePath)
	if err != nil {
		return ""
	}

	return string(content)
}
