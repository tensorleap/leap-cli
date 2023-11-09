package hub

import (
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"time"

	"github.com/tensorleap/leap-cli/pkg/hub/storage"
	"github.com/tensorleap/leap-cli/pkg/log"
	"gopkg.in/yaml.v3"
)

type ProjectContext struct {
	Meta    ProjectMeta
	BgImage Image
}

type HubStorageApi struct {
	Namespace   string
	FilesClient storage.StorageClient
}

type ProjectFilePaths struct {
	BgImagePath string
	TarPath     string
	MetaPath    string
}

type UploadCtx struct {
	Meta  ProjectMeta
	Files ProjectFilePaths
}

func NewHubStorageApi(namespace string, filesClient storage.StorageClient) *HubStorageApi {
	return &HubStorageApi{
		Namespace:   namespace,
		FilesClient: filesClient,
	}
}

func (hs *HubStorageApi) UploadProjectContentBySignedUrl(meta *ProjectMeta) (string, *ProjectFilePaths, error) {
	filePaths := GenerateProjectFilePaths(hs.Namespace, meta)

	log.Infof("HubStorage - Create signed url for content of project: '%s' version: %d", meta.Name, meta.SchemaVersion)

	signedTarUrl, err := hs.FilesClient.CreateWriteableSignedUrl(filePaths.TarPath, time.Now().Add(4*time.Hour))
	if err != nil {
		return "", nil, err
	}

	return signedTarUrl, filePaths, nil
}

func (hs *HubStorageApi) UploadProjectContent(fileStream io.Reader, meta *ProjectMeta) (*ProjectFilePaths, error) {
	name := meta.Name
	schemaVersion := meta.SchemaVersion

	filePaths := GenerateProjectFilePaths(hs.Namespace, meta)

	start, stop, _ := log.NewSpinner(fmt.Sprintf("HubStorage - Uploading content of project: '%s' version: %d", name, schemaVersion))
	start()
	defer stop()
	err := hs.FilesClient.UploadFile(filePaths.TarPath, fileStream, nil)
	if err != nil {
		return nil, err
	}

	return filePaths, nil
}

func (hs *HubStorageApi) UpdateProjectMeta(filePaths *ProjectFilePaths, ctx *ProjectContext) (*HubProjectVersion, error) {
	metaBytes, _ := json.Marshal(ctx.Meta)

	log.Info("HubStorage - Uploading background image")
	err := hs.FilesClient.UploadFileBuffer(filePaths.BgImagePath, ctx.BgImage.Buffer, nil)
	if err != nil {
		return nil, err
	}

	log.Info("HubStorage - Uploading meta file")
	err = hs.FilesClient.UploadFileBuffer(filePaths.MetaPath, metaBytes, nil)
	if err != nil {
		return nil, err
	}

	tarPathWithoutNamespace := strings.TrimPrefix(filePaths.TarPath, fmt.Sprintf("%s/", hs.Namespace))
	bgImagePathWithoutNamespace := strings.TrimPrefix(filePaths.BgImagePath, fmt.Sprintf("%s/", hs.Namespace))

	fileMeta, err := hs.FilesClient.GetFileMeta(filePaths.TarPath)
	if err != nil {
		return nil, err
	}

	projectMetaVersion := &HubProjectVersion{
		ProjectMeta: ctx.Meta,
		Size:        int(fileMeta.Size),
		Path:        tarPathWithoutNamespace,
		CreatedAt:   time.Now().Unix(),
	}
	projectMetaVersion.BgImagePath = bgImagePathWithoutNamespace

	return projectMetaVersion, nil
}

func (hs *HubStorageApi) DownloadProject(projectName string, projectVersion int, destFileName string) error {
	projectVersionDir := fmt.Sprintf("%s/%s", hs.Namespace, CalcProjectVersionPath(projectName, projectVersion))
	files, _ := hs.FilesClient.ListDirectoryObjects(projectVersionDir, true)
	var fileName string
	for _, file := range files {
		if strings.HasSuffix(file, PROJECT_TAR_FILE_SUFFIX_NAME) {
			fileName = file
			break
		}
	}
	if fileName == "" {
		return fmt.Errorf("not found tar file for project: %s version: %d", projectName, projectVersion)
	}

	start, stop, _ := log.NewSpinner(fmt.Sprintf("Downloading project: %s, into: %s", projectName, destFileName))
	start()
	defer stop()
	return hs.FilesClient.DownloadFile(fileName, destFileName)
}

func (hs *HubStorageApi) SetMeta(meta HubMeta) error {
	metaAsYaml, err := yaml.Marshal(meta)
	if err != nil {
		return err
	}

	metaPath := fmt.Sprintf("%s/%s", hs.Namespace, NAMESPACE_META_FILE_NAME)
	err = hs.FilesClient.UploadFileBuffer(metaPath, metaAsYaml, &storage.UploadOptions{
		NoCache: true,
	})
	if err != nil {
		return err
	}

	log.Infof("HubStorage - Meta updated successfully")
	return nil
}

func (hs *HubStorageApi) DeleteProject(projectName string) (bool, error) {
	projectDir := fmt.Sprintf("%s/%s", hs.Namespace, CalcProjectPath(projectName))
	isDeleted, err := hs.FilesClient.DeleteFilesInDirectory(projectDir)
	if err != nil {
		return false, err
	}

	if isDeleted {
		log.Infof("HubStorage - project: '%s' deleted successfully", projectName)
	}

	return isDeleted, nil
}

func (hs *HubStorageApi) DeleteProjectVersion(projectName string, projectVersion int) (bool, error) {
	projectVersionDir := fmt.Sprintf("%s/%s", hs.Namespace, CalcProjectVersionPath(projectName, projectVersion))
	isDeleted, err := hs.FilesClient.DeleteFilesInDirectory(projectVersionDir)
	if err != nil {
		return false, err
	}

	if isDeleted {
		log.Infof("HubStorage - project: '%s' version: %d deleted successfully", projectName, projectVersion)
	}

	return isDeleted, nil
}

func (hs *HubStorageApi) GetMetaWrapper() (*HubMetaWrapper, error) {
	meta, err := hs.GetMeta()
	if err != nil {
		return nil, fmt.Errorf("failed to get meta: %v", err)
	}

	return NewHubMetaWrapper(*meta), nil
}

func (hs *HubStorageApi) GetMeta() (*HubMeta, error) {
	filePath := fmt.Sprintf("%s/%s", hs.Namespace, NAMESPACE_META_FILE_NAME)
	exists, err := hs.FilesClient.CheckIfFileExists(filePath)
	if err != nil {
		return nil, err
	}

	if exists {
		data, err := hs.FilesClient.GetFileBuffer(filePath)
		if err != nil {
			return nil, err
		}

		var meta HubMeta
		err = yaml.Unmarshal(data, &meta)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal meta: %v", err)
		}

		return &meta, nil
	}

	return &HubMeta{
		Namespace: hs.Namespace,
		Projects:  map[string]HubProject{},
	}, nil
}

func CalcProjectPath(projectName string) string {
	return fmt.Sprintf("projects/%s", projectName)
}

func CalcProjectVersionPath(projectName string, projectVersion int) string {
	return fmt.Sprintf("%s/versions/%d", CalcProjectPath(projectName), projectVersion)
}

func GenerateProjectFilePaths(namespace string, meta *ProjectMeta) *ProjectFilePaths {
	name := meta.Name
	schemaVersion := meta.SchemaVersion
	bgImageName := meta.BgImagePath

	imgType := filepath.Ext(filepath.Base(bgImageName))
	projectRelativeDir := CalcProjectVersionPath(name, schemaVersion)
	now := time.Now().Unix()
	tarPath := fmt.Sprintf("%s/%d-%s", projectRelativeDir, now, PROJECT_TAR_FILE_SUFFIX_NAME)
	bgImagePath := fmt.Sprintf("%s/%d-%s%s", projectRelativeDir, now, PROJECT_BG_IMAGE_FILE_NAME, imgType)
	metaPath := fmt.Sprintf("%s/%d-%s", projectRelativeDir, now, PROJECT_META_FILE_NAME)

	return &ProjectFilePaths{
		BgImagePath: fmt.Sprintf("%s/%s", namespace, bgImagePath),
		TarPath:     fmt.Sprintf("%s/%s", namespace, tarPath),
		MetaPath:    fmt.Sprintf("%s/%s", namespace, metaPath),
	}
}
