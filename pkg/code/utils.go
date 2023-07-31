package code

import (
	"context"
	"fmt"
	"io/fs"
	"os"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/local"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

var ErrEmptyCodeIntegrationVersion = fmt.Errorf("CodeIntegration is empty")

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
	for _, pattern := range workspaceConfig.IncludePatterns {
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
