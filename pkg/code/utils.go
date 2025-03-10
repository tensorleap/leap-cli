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
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/local"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
	"github.com/tensorleap/leap-cli/pkg/workspace"
	"k8s.io/kubectl/pkg/util/slice"
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

func GetCodeIntegrationFromFlag(ctx context.Context, codeIntegrationIdFlag string, askForNewProjectFirst bool) (code *CodeIntegration, wasCreated bool, err error) {
	codeIntegrations, err := GetCodeIntegrations(ctx)
	if err != nil {
		return nil, false, err
	}
	var selected *CodeIntegration
	if len(codeIntegrationIdFlag) > 0 {
		selected, err = entity.GetEntityById(codeIntegrationIdFlag, codeIntegrations, CodeIntegrationEntityDesc)
	} else {
		selected, wasCreated, err = SelectOrCreateCodeIntegration(ctx, codeIntegrations, askForNewProjectFirst)
	}
	if err != nil {
		return nil, false, err
	}
	return selected, wasCreated, nil
}

func SelectOrCreateCodeIntegration(ctx context.Context, codeIntegrations []CodeIntegration, askIsCreateNewFirst bool) (*tensorleapapi.Dataset, bool, error) {
	createCodeIntegration := func() (*CodeIntegration, error) {
		return CreateCodeIntegration(ctx, codeIntegrations)
	}
	return entity.SelectEntityOrCreateOne(codeIntegrations, createCodeIntegration, askIsCreateNewFirst, CodeIntegrationEntityDesc)
}

func AskForCodeIntegrationName(codeIntegrations []CodeIntegration) (name string, err error) {

	existingNames := entity.GetNames(codeIntegrations, CodeIntegrationEntityDesc)

	name, err = entity.AskForName(existingNames, "", CodeIntegrationEntityDesc)
	return
}

func AskForCodeIntegrationNameIfExisted(name string, codeIntegrations []CodeIntegration) (string, error) {

	existingNames := entity.GetNames(codeIntegrations, CodeIntegrationEntityDesc)
	if len(name) > 0 {
		if !slice.ContainsString(existingNames, name, nil) {
			return name, nil
		}
		log.Warnf("Code integration name '%s' already exists, Please select new one", name)
	}

	return entity.AskForName(existingNames, name, CodeIntegrationEntityDesc)
}

func AskForPythonVersion(defaultVersionId string, baseImageTypes []tensorleapapi.GenericBaseImage) (string, error) {
	displayNames := []string{}
	versionIds := []string{}
	defaultVersionDisplay := ""
	for _, version := range baseImageTypes {
		displayNames = append(displayNames, version.DisplayName)
		versionIds = append(versionIds, version.Id)
		if version.Id == defaultVersionId {
			defaultVersionDisplay = version.DisplayName
		}
	}
	prompt := &survey.Select{
		Message: "Select python version",
		Options: displayNames,
		Default: defaultVersionDisplay,
	}

	var pythonVersion string
	err := survey.AskOne(prompt, &pythonVersion)
	if err != nil {
		return "", err
	}
	for i, displayName := range displayNames {
		if displayName == pythonVersion {
			return versionIds[i], nil
		}
	}
	return "", fmt.Errorf("python version %s is not available", pythonVersion)
}

func GetPythonVersions(ctx context.Context) ([]tensorleapapi.GenericBaseImage, string, error) {
	currentVersions, res, err := api.ApiClient.GetGenericBaseImageTypes(ctx).Execute()
	if err := api.CheckRes(res, err); err != nil {
		return nil, "", fmt.Errorf("failed to get python versions: %v", err)
	}

	return currentVersions.BaseImageTypes, currentVersions.DefaultBaseImageType, nil
}

func selectPythonVersionFromFlag(flag string, baseImages []tensorleapapi.GenericBaseImage) (string, error) {
	if len(flag) > 0 {
		id := ""
		for _, baseImage := range baseImages {
			if baseImage.Id == flag || strings.HasSuffix(baseImage.DisplayName, flag) {
				id = baseImage.Id
				break
			}
		}

		if len(id) == 0 {
			return "", fmt.Errorf("python version %s is not available", flag)
		}
		return id, nil
	}
	return "", nil
}

func GetPythonVersionFromFlag(ctx context.Context, flag string) (string, error) {
	return SyncPythonVersionFromFlagAndConfig(ctx, flag, nil)
}

func SyncPythonVersionFromFlagAndConfig(ctx context.Context, flag string, workspaceConfig *workspace.WorkspaceConfig) (string, error) {
	baseImages, defaultVersionId, err := GetPythonVersions(ctx)
	if err != nil {
		return "", err
	}

	flagVersionId, err := selectPythonVersionFromFlag(flag, baseImages)
	if err != nil {
		return "", err
	}

	updateConfig := func(versionId string) error {
		if workspaceConfig == nil {
			return nil
		}
		workspaceConfig.PythonVersion = versionId
		err = workspace.SetWorkspaceConfig(workspaceConfig, ".")
		return err
	}
	if len(flagVersionId) > 0 {
		err = updateConfig(flagVersionId)
		if err != nil {
			return "", err
		}
		return flagVersionId, nil
	}

	if workspaceConfig != nil && len(workspaceConfig.PythonVersion) > 0 {
		found, _ := selectPythonVersionFromFlag(workspaceConfig.PythonVersion, baseImages)
		if len(found) >= 0 {
			return found, nil
		}
		log.Warnf("Python version %s is not available", workspaceConfig.PythonVersion)
	}

	selectedVersionId, err := AskForPythonVersion(defaultVersionId, baseImages)
	if err != nil {
		return "", err
	}

	err = updateConfig(selectedVersionId)
	if err != nil {
		return "", err
	}
	return selectedVersionId, nil
}

func isDatasetVersionEmpty(datasetVersion *tensorleapapi.DatasetVersion) bool {
	return len(datasetVersion.GetBlobPath()) == 0
}

func CloneCodeIntegrationVersion(ctx context.Context, codeIntegrationVersion *tensorleapapi.DatasetVersion, outputDir string, specificFileName string) ([]string, error) {
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

	files, err := local.DownloadAndExtractTarFile(res.GetUrl(), outputDir, specificFileName)
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

func BundleCodeIntoTempFile(filesDir string, workspaceConfig *workspace.WorkspaceConfig, externalLeapMappingPath string) (close func(), tarGzFile *os.File, err error) {
	filePaths, err := getDatasetFiles(filesDir, workspaceConfig)
	if err != nil {
		return
	}

	if externalLeapMappingPath != "" {
		filePaths = filterOutBySuffix(filePaths, BindingFilePath)
	}

	tarGzFile, err = os.CreateTemp("", "tensorleap-*.tar.gz")
	if err != nil {
		return
	}
	close = func() { local.CleanupTempFile(tarGzFile) }

	if err = local.CreateTarGzFile(filesDir, filePaths, tarGzFile, externalLeapMappingPath, BindingFilePath); err != nil {
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
		pattern = local.ConvertPathPatternToUnix(pattern)
		matches, err := fs.Glob(currentDirFs, pattern)
		if err != nil {
			return nil, err
		}

		allMatchedFiles = append(allMatchedFiles, matches...)
	}

	uniqueFilePaths := distinctStrings(allMatchedFiles)
	return uniqueFilePaths, nil
}

func filterOutBySuffix(paths []string, suffix string) []string {
	var filtered []string
	for _, path := range paths {
		if !strings.HasSuffix(path, suffix) {
			filtered = append(filtered, path)
		}
	}
	return filtered
}

func distinctStrings(arr []string) []string {
	uniqueMap := make(map[string]bool)
	var result []string

	for _, str := range arr {
		if !uniqueMap[str] {
			uniqueMap[str] = true
			result = append(result, str)
		}
	}

	return result
}

func GetAndUpdateCodeIntegrationIfNotExists(ctx context.Context, workspaceConfig *workspace.WorkspaceConfig) (code *CodeIntegration, wasCreated bool, err error) {
	codeIntegrations, err := GetCodeIntegrations(ctx)

	if err != nil {
		return nil, false, fmt.Errorf("failed to get code integration: %v", err)
	}
	for _, codeIntegration := range codeIntegrations {
		if codeIntegration.Cid == workspaceConfig.CodeIntegrationId {
			return &codeIntegration, false, nil
		}
	}

	log.Infof("Not found code integration id: %s. Select or create new code integration", workspaceConfig.CodeIntegrationId)

	codeIntegration, wasCreated, err := SelectOrCreateCodeIntegration(ctx, codeIntegrations, true)
	if err != nil {
		return nil, false, err
	}

	workspaceConfig.CodeIntegrationId = codeIntegration.GetCid()
	log.Info("Updating codeIntegrationId")
	err = workspace.SetWorkspaceConfig(workspaceConfig, "")
	if err != nil {
		return nil, false, err
	}
	return codeIntegration, wasCreated, nil
}

func PrintCodeIntegrationVersionParserErr(civ *CodeIntegrationVersion) {

	log.Error("Code parsing failed, see error below:")
	if civ.Metadata.SetupStatus.GeneralError != nil {
		log.Error("General error:")
		fmt.Println(*civ.Metadata.SetupStatus.GeneralError)
	}

	for _, binderStatus := range civ.Metadata.SetupStatus.BindersStatus {
		if !binderStatus.IsPassed && len(binderStatus.Display) > 0 {
			log.Errorf("binder error: %s", binderStatus.Name)
			for key, display := range binderStatus.Display {
				fmt.Printf("%s: %s\n", key, display)
			}
		}
	}

	if civ.Metadata.SetupStatus.PrintLog != nil {
		log.Info("Log:")
		fmt.Println(*civ.Metadata.SetupStatus.PrintLog)
		return
	}
}

func PushCode(ctx context.Context, force bool, codeIntegrationId string, tarGzFile *os.File, entryFile, secretId, branch, message, pythonVersion string) (pushed bool, current *CodeIntegrationVersion, err error) {
	if !force {
		log.Info("Checking if code has changed")

		latestVersion, err := GetLatestVersion(ctx, codeIntegrationId, branch)

		if err != nil && !errors.Is(err, ErrEmptyCodeIntegrationVersion) {
			log.Warnf("Failed to get latest code integration version: %v", err)
		}
		if err == nil {
			change, err := CompareCodeVersion(ctx, latestVersion, tarGzFile, entryFile, secretId, pythonVersion)
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
	fileStat, err := tarGzFile.Stat()
	if err != nil {
		return false, nil, fmt.Errorf("failed to get file stat: %v", err)
	}

	codeIntegrationVersion, err := AddCodeIntegrationVersion(ctx, tarGzFile, fileStat.Size(), codeIntegrationId, entryFile, secretId, branch, message, pythonVersion)
	if err != nil {
		return false, nil, err
	}
	return true, codeIntegrationVersion, nil
}

func CompareCodeVersion(ctx context.Context, compareVersion *CodeIntegrationVersion, tarGzFile *os.File, entryFile, secretId, pythonVersion string) (bool, error) {

	if isDatasetVersionEmpty(compareVersion) {
		return true, nil
	}

	if compareVersion.Metadata.GetSecretManagerId() != secretId || compareVersion.CodeEntryFile != entryFile || compareVersion.GetGenericBaseImageType() != pythonVersion {
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

func GetDatasetMappingYaml(ctx context.Context, codeIntegrationId, branch string) string {
	codeIntegrationVersion, err := GetLatestVersion(ctx, codeIntegrationId, branch)
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

func SyncBranchFromFlagAndConfig(flagBranch string, workspaceConfig *workspace.WorkspaceConfig, branches []string, defaultBranch string) (string, error) {
	var branch string
	if len(flagBranch) == 0 {
		branch = workspaceConfig.Branch
	} else {
		branch = flagBranch
	}

	if len(branch) > 0 {
		branch, err := CreateOrSelectBranch(branch, branches, defaultBranch)
		if err != nil {
			return "", err
		}
		if workspaceConfig.Branch != branch {
			log.Infof("Updating leap.yaml branch to %s", branch)
			workspaceConfig.Branch = branch
			err = workspace.SetWorkspaceConfig(workspaceConfig, ".")
			if err != nil {
				return "", err
			}
		}
	}
	return branch, nil
}

func CreateOrSelectBranch(branch string, branches []string, defaultBranch string) (selectedBranch string, err error) {
	if branch != "" {
		selectedBranch = branch
		return
	}
	// autocomplete
	prompt := &survey.Input{
		Message: "Enter branch name",
		Default: defaultBranch,
		Help:    fmt.Sprintf("Available branches: %s", branches),
	}
	err = survey.AskOne(prompt, &selectedBranch)
	if err != nil {
		return "", err
	}
	return selectedBranch, nil
}

func SelectBranch(branches []string, defaultBranch string) (selectedBranch string, err error) {
	if len(branches) == 0 {
		return "", fmt.Errorf("no version available")
	}
	if len(branches) == 1 {
		selectedBranch = branches[0]
		log.Infof("Selected branch: %s", selectedBranch)
		return selectedBranch, nil
	}

	prompt := &survey.Select{
		Message: "Select branch",
		Options: branches,
		Default: defaultBranch,
	}
	err = survey.AskOne(prompt, &selectedBranch)
	if err != nil {
		return "", err
	}
	return selectedBranch, nil
}

func BranchesFromCodeIntegration(codeIntegration *CodeIntegration) []string {
	var branches []string
	for _, latest := range codeIntegration.LatestVersions {
		branches = append(branches, latest.Branch)
	}
	return branches
}
