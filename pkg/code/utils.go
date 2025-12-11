package code

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/bmatcuk/doublestar/v4"
	"github.com/samber/lo"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/local"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

var ErrEmptyCodeSnapshot = fmt.Errorf("CodeIntegration is empty")

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

func isCodeSnapshotEmpty(codeSnapshot *tensorleapapi.CodeSnapshot) bool {
	return len(codeSnapshot.GetBlobName()) == 0
}

func CloneCodeSnapshot(ctx context.Context, codeSnapshot *tensorleapapi.CodeSnapshot, outputDir string, specificFileName string) ([]string, error) {
	if isCodeSnapshotEmpty(codeSnapshot) {
		return []string{}, ErrEmptyCodeSnapshot
	}
	blobPath := codeSnapshot.GetBlobName()
	blobWithProjectId := fmt.Sprintf("projects/%s/%s", codeSnapshot.GetProjectId(), blobPath)
	log.Infof("Downloading latest code integration version from %s", blobPath)
	downloadUrl, err := api.GetDownloadSignedUrl(ctx, blobWithProjectId)
	if err != nil {
		return nil, err
	}

	files, err := local.DownloadAndExtractTarFile(downloadUrl, outputDir, specificFileName)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func DownloadCodeSnapshotArchive(ctx context.Context, codeSnapshot *tensorleapapi.CodeSnapshot, outputPath string) error {
	if isCodeSnapshotEmpty(codeSnapshot) {
		return ErrEmptyCodeSnapshot
	}
	blobPath := codeSnapshot.GetBlobName()
	blobWithProjectId := fmt.Sprintf("projects/%s/%s", codeSnapshot.GetProjectId(), blobPath)
	log.Infof("Downloading code archive from %s", blobPath)
	downloadUrl, err := api.GetDownloadSignedUrl(ctx, blobWithProjectId)
	if err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer func() { _ = file.Close() }()

	err = api.DownloadFile(downloadUrl, file)
	if err != nil {
		return fmt.Errorf("failed to download archive: %w", err)
	}
	return nil
}

func FetchFileFromTarGz(blobURL string, filename string) (string, error) {
	resp, err := http.Get(blobURL)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return "", err
	}
	defer func() { _ = gz.Close() }()

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

var RequirementsFiles = []string{
	"pyproject.toml",
	"requirements.txt",
	"tensorleap_requirements.txt",
}

func isRequirementsFile(path string) bool {
	return lo.SomeBy(RequirementsFiles, func(file string) bool {
		return path == file
	})
}

// LargeFileExtensions contains extensions for files that are typically large
// and may indicate model weights, media files, datasets, or other binary data
var LargeFileExtensions = []string{
	// Model weights and checkpoints
	".onnx",
	".pth", ".pt", ".ckpt", ".bin",
	".safetensors",
	".pb", ".tflite",
	".h5", ".hdf5", ".keras",
	".pkl", ".pickle", ".joblib",
	".npz", ".npy",
	".mlx", ".gguf", ".ggml",
	".trt", ".engine",

	// Archives
	".tar", ".tar.gz", ".tgz",

	// Images
	".png", ".jpg", ".jpeg", ".gif", ".bmp", ".tiff", ".webp", ".ico",
	".psd", ".ai", ".sketch",

	// Video/Audio
	".mp4", ".mov", ".avi", ".mkv", ".webm",
	".mp3", ".wav", ".flac", ".ogg", ".m4a",

	// 3D models
	".stl", ".obj", ".fbx", ".glb", ".gltf",

	// Data files
	".jsonl", ".parquet", ".arrow", ".orc", ".feather",
	".db", ".sqlite", ".sqlite3",
	".zip", ".7z", ".rar",
	".dump", ".bak",
}

// LargeFileDirPatterns contains directory patterns for typically large/state directories
var LargeFileDirPatterns = []string{
	".git",
}

// LargeFilesResult holds the result of scanning for large files
type LargeFilesResult struct {
	TotalSize       int64               // Total size of all large files in bytes
	SizePerPattern  map[string]int64    // Size sum per glob pattern (e.g., "*.png" -> 1024)
	MatchedPatterns map[string][]string // Map of glob pattern (e.g., "*.png", ".git/**") -> list of matching file paths
	RemainingFiles  []string            // Files that don't match any large file pattern
}

// SumLargeFilesSize calculates the total size of files that match large file extensions
// or are inside state directories like .git
// Returns:
// - LargeFilesResult containing total size, matched patterns with their files, and remaining files
// - error if any
func SumLargeFilesSize(filesDir string, filePaths []string) (*LargeFilesResult, error) {
	result := &LargeFilesResult{
		TotalSize:       0,
		SizePerPattern:  make(map[string]int64),
		MatchedPatterns: make(map[string][]string),
		RemainingFiles:  []string{},
	}

	for _, filePath := range filePaths {
		fullPath := filepath.Join(filesDir, filePath)
		matched := false
		var matchedPattern string

		// Check if file is in a state directory (e.g., .git)
		for _, dirPattern := range LargeFileDirPatterns {
			if strings.HasPrefix(filePath, dirPattern+string(filepath.Separator)) || strings.Contains(filePath, string(filepath.Separator)+dirPattern+string(filepath.Separator)) {
				matched = true
				matchedPattern = dirPattern + "/**"
				break
			}
		}

		// Check if file has a large file extension
		if !matched {
			lowerPath := strings.ToLower(filePath)
			for _, ext := range LargeFileExtensions {
				if strings.HasSuffix(lowerPath, ext) {
					matched = true
					matchedPattern = "*" + ext
					break
				}
			}
		}

		if matched {
			// Add file to matched patterns
			result.MatchedPatterns[matchedPattern] = append(result.MatchedPatterns[matchedPattern], filePath)

			// Add file size to total and per pattern
			info, err := os.Stat(fullPath)
			if err != nil {
				// Skip files that can't be accessed for size calculation
				continue
			}
			fileSize := info.Size()
			result.TotalSize += fileSize
			result.SizePerPattern[matchedPattern] += fileSize
		} else {
			// File doesn't match any large file pattern
			result.RemainingFiles = append(result.RemainingFiles, filePath)
		}
	}

	return result, nil
}

func BundleCodeIntoTempFile(filesDir string, workspaceConfig *workspace.WorkspaceConfig) (close func(), tarGzFile *os.File, err error) {
	filePaths, err := getCodeFiles(filesDir, workspaceConfig)
	if err != nil {
		return
	}
	isCodeIntegrationUsePippinButNoRequirementsTxt := lo.EveryBy(filePaths, func(path string) bool {
		return !isRequirementsFile(path)
	})
	requirementsFilesStr := strings.Join(RequirementsFiles, ", ")

	if isCodeIntegrationUsePippinButNoRequirementsTxt {
		isContinue := false
		prompt := &survey.Confirm{
			Message: fmt.Sprintf("No requirements file (%s) found in the root directory. Continue?", requirementsFilesStr),
			Default: isContinue,
		}
		err = survey.AskOne(prompt, &isContinue)
		if err != nil {
			return nil, nil, err
		}
		if !isContinue {
			return nil, nil, fmt.Errorf("please ensure a requirements file (%s) exists and is specified in leap.yaml", requirementsFilesStr)
		}
	}

	isEntryFileInFilePaths := lo.SomeBy(filePaths, func(path string) bool {
		return path == workspaceConfig.EntryFile
	})
	if !isEntryFileInFilePaths && workspaceConfig.EntryFile != "" {
		log.Infof("Adding entry file %s to file paths", workspaceConfig.EntryFile)
		filePaths = append(filePaths, workspaceConfig.EntryFile)
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

func getCodeFiles(filesDir string, workspaceConfig *workspace.WorkspaceConfig) ([]string, error) {
	includePatterns := workspaceConfig.IncludePatterns
	excludePatterns := workspaceConfig.ExcludePatterns

	if len(includePatterns) == 0 {
		includePatterns = []string{"**"}
	}

	fileSet := make(map[string]struct{})

	for _, pattern := range includePatterns {
		unixPattern := local.ConvertPathPatternToUnix(pattern)
		matches, err := doublestar.FilepathGlob(filepath.Join(filesDir, unixPattern))
		if err != nil {
			return nil, err
		}
		for _, match := range matches {
			info, err := os.Stat(match)
			if err != nil {
				return nil, err
			}
			if !info.IsDir() {
				relPath, err := filepath.Rel(filesDir, match)
				if err != nil {
					return nil, err
				}
				fileSet[relPath] = struct{}{}
			}
		}
	}

	finalFiles := make([]string, 0, len(fileSet))
	for file := range fileSet {
		if !isExcluded(file, excludePatterns) {
			finalFiles = append(finalFiles, file)
		}
	}

	return finalFiles, nil
}

func isExcluded(path string, excludePatterns []string) bool {
	for _, pattern := range excludePatterns {
		unixPattern := local.ConvertPathPatternToUnix(pattern)
		matched, err := doublestar.PathMatch(unixPattern, path)
		if err == nil && matched {
			return true
		}
	}
	return false
}

func PrintCodeSnapshotParserErr(civ *CodeSnapshot) {

	if civ.ParseResult == nil {
		log.Error("Code parsing failed, but no parse result found")
		return
	}
	log.Error("Code parsing failed, see error below:")
	if civ.ParseResult.SetupStatus.GeneralError != nil {
		log.Error("General error:")
		fmt.Println(*civ.ParseResult.SetupStatus.GeneralError)
	}

	for _, binderStatus := range civ.ParseResult.SetupStatus.BindersStatus {
		if !binderStatus.IsPassed && len(binderStatus.Display) > 0 {
			log.Errorf("binder error: %s", binderStatus.Name)
			for key, display := range binderStatus.Display {
				fmt.Printf("%s: %s\n", key, display)
			}
		}
	}

	if civ.ParseResult.SetupStatus.PrintLog != nil {
		log.Info("Log:")
		fmt.Println(*civ.ParseResult.SetupStatus.PrintLog)
		return
	}
}

func PushCode(ctx context.Context, tarGzFile *os.File, entryFile, secretId, pythonVersion, versionName, projectId, branch string, overwriteVersionId string) (pushed bool, current *tensorleapapi.PushCodeSnapshotResponse, err error) {

	fileStat, err := tarGzFile.Stat()
	if err != nil {
		return false, nil, fmt.Errorf("failed to get file stat: %v", err)
	}

	codeSnapshot, err := PushCodeSnapshot(
		ctx, tarGzFile, fileStat.Size(),
		entryFile, secretId, pythonVersion, versionName,
		projectId,
		branch,
		overwriteVersionId,
	)
	if err != nil {
		return false, nil, err
	}
	return true, codeSnapshot, nil
}

func CompareCodeVersion(ctx context.Context, compareVersion *CodeSnapshot, tarGzFile *os.File, entryFile, secretId, pythonVersion string) (bool, error) {

	if isCodeSnapshotEmpty(compareVersion) {
		return true, nil
	}

	if compareVersion.GetSecretManagerId() != secretId || compareVersion.GetCodeEntryFile() != entryFile || compareVersion.GetGenericBaseImageType() != pythonVersion {
		return true, nil
	}

	comperedBlobPath := compareVersion.GetBlobName()
	tempLatestVersionFile, err := os.CreateTemp("", "tensorleap-*.tar.gz")
	if err != nil {
		return true, fmt.Errorf("failed to create temp file: %v", err)

	}
	defer local.CleanupTempFile(tempLatestVersionFile)

	downloadUrl, err := api.GetDownloadSignedUrl(ctx, comperedBlobPath)
	if err != nil {
		return true, fmt.Errorf("failed to get download signed url: %v", err)
	}
	err = api.DownloadFile(downloadUrl, tempLatestVersionFile)
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

func IsCodeParseFailed(codeSnapshot *CodeSnapshot) bool {
	if codeSnapshot.ParseResult == nil {
		return false
	}
	return codeSnapshot.ParseResult.TestStatus == tensorleapapi.TESTSTATUS_TEST_FAIL
}

func IsCodeEnded(codeSnapshot *CodeSnapshot) bool {
	if IsCodeParseFailed(codeSnapshot) || IsCodeFinished(codeSnapshot) {
		return true
	}
	return false
}

func IsCodeFinished(codeSnapshot *CodeSnapshot) bool {
	if codeSnapshot.ParseResult == nil {
		return false
	}
	return codeSnapshot.ParseResult.TestStatus == tensorleapapi.TESTSTATUS_TEST_SUCCESS
}

func SyncBranchFromFlagAndConfig(flagBranch string, workspaceConfig *workspace.WorkspaceConfig) (string, error) {
	var branch string
	if len(flagBranch) == 0 {
		branch = workspaceConfig.Branch
	} else {
		branch = flagBranch
	}

	if len(branch) > 0 {

		if workspaceConfig.Branch != branch {
			log.Infof("Updating leap.yaml branch to %s", branch)
			workspaceConfig.Branch = branch
			err := workspace.SetWorkspaceConfig(workspaceConfig, ".")
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
