package project

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/hub"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

type ProjectEntity = tensorleapapi.Project

var ProjectEntityDesc = entity.NewEntityDescriptor[ProjectEntity](
	"project",
	"projects",
	func(p *ProjectEntity) string { return p.GetName() },
	func(p *ProjectEntity) string { return p.GetCid() },
)

func GetProjects(ctx context.Context) ([]ProjectEntity, error) {
	res, _, err := api.ApiClient.GetProjects(ctx).Execute()
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

type AddProjectDetails struct {
	Name        string
	Description string
}

func AddProject(ctx context.Context, projectDetails *AddProjectDetails) (*ProjectEntity, error) {
	res, _, err := api.ApiClient.AddProject(ctx).ProjectMeta(*tensorleapapi.NewProjectMeta(
		projectDetails.Name,
		projectDetails.Description,
		[]string{},
		tensorleapapi.HUBPUBLISHPOLICY_NO_PUBLISH,
	)).Execute()

	if err != nil {
		return nil, err
	}
	newProject := &res.Project
	entity.InfoCreation(newProject, ProjectEntityDesc)
	return newProject, nil
}

func DeleteProject(ctx context.Context, project *ProjectEntity) error {
	_, err := api.ApiClient.DeleteProject(ctx).
		DeleteProjectParams(*tensorleapapi.NewDeleteProjectParams(project.GetCid())).
		Execute()
	if err != nil {
		return err
	}
	entity.InfoDeletion(project, ProjectEntityDesc)
	return nil
}

func ImportProject(ctx context.Context, projectName, importUrl string, meta *hub.ProjectMeta) error {
	reqParams := tensorleapapi.NewImportProjectRequest(
		projectName,
		meta.Description,
		meta.Tags,
		tensorleapapi.HUBPUBLISHPOLICY_NO_PUBLISH,
		importUrl,
	)
	reqParams.SetCategories(meta.Categories)
	reqParams.SetBgImageUrl(meta.BgImagePath)

	res, err := api.ApiClient.ImportProject(ctx).ImportProjectRequest(*reqParams).Execute()
	if err = api.CheckRes(res, err); err != nil {
		return fmt.Errorf("failed to import project %s/n %v", projectName, err)
	}
	log.Infof("Import process for '%s' is in progress on the server.", projectName)
	return nil
}

func ImportProjectFromFile(ctx context.Context, filePath, projectName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	projectCtx, err := hub.ExtractProjectContextFromTar(file)
	if err != nil {
		return fmt.Errorf("failed to extract project context: %v", err)
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}
	projects, err := GetProjects(ctx)
	if err != nil {
		return err
	}
	projectName, err = ValidateProjectName(projectName, projectCtx.Meta.Name, projects)
	if err != nil {
		return err
	}

	getUrl, err := uploadProjectToTempFile(ctx, projectName, file)
	if err != nil {
		return err
	}
	return ImportProject(ctx, projectName, getUrl, &projectCtx.Meta)
}

func ImportProjectFromStream(ctx context.Context, projectName string, meta *hub.ProjectMeta, stream io.Reader) error {
	getUrl, err := uploadProjectToTempFile(ctx, projectName, stream)
	if err != nil {
		return err
	}
	return ImportProject(ctx, projectName, getUrl, meta)
}

func uploadProjectToTempFile(ctx context.Context, projectName string, projectReader io.Reader) (string, error) {
	tempSignedUploadUrlParams := *tensorleapapi.NewGetUploadSignedUrlParams(projectName)
	uploadUrl, res, err := api.ApiClient.GetUploadSignedUrl(ctx).GetUploadSignedUrlParams(tempSignedUploadUrlParams).Execute()
	if err = api.CheckRes(res, err); err != nil {
		return "", fmt.Errorf("failed to get signed url: %v", err)
	}
	err = api.UploadFile(uploadUrl.Url, projectReader)
	if err != nil {
		return "", fmt.Errorf("failed to upload project: %v", err)
	}
	log.Infof("Uploaded successfully project: %s", projectName)

	getUrlParams := *tensorleapapi.NewGetDownloadSignedUrlParams(uploadUrl.GetFileName())
	getUrlParams.SetNotUseRequestOrigin(true)
	getUrl, res, err := api.ApiClient.GetDownloadSignedUrl(ctx).
		GetDownloadSignedUrlParams(getUrlParams).Execute()
	if err = api.CheckRes(res, err); err != nil {
		return "", fmt.Errorf("failed to get signed url: %v", err)
	}
	return getUrl.Url, nil
}

func ExportProjectIntoFile(ctx context.Context, project *ProjectEntity, outputDir string) error {
	res, err := ExportProject(ctx, project.GetCid())
	if err != nil {
		return err
	}

	fileName := getFileNameFromResponse(res)

	defer res.Body.Close()

	filePath := path.Join(outputDir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	start, stop, _ := log.NewSpinner(fmt.Sprintf("Exporting project '%s', into: '%s'", project.Name, filePath))
	start()
	defer stop()

	_, err = io.Copy(file, res.Body)
	if err != nil && err != io.EOF {
		defer os.Remove(filePath)
		return err
	}
	return nil
}

func PublishProject(ctx context.Context, projectId, tarSignedUrl string) error {
	start, stop, _ := log.NewSpinner("Publishing project content by signed url")
	start()
	defer stop()
	res, err := api.ApiClient.PublishProject(ctx).PublishProjectRequest(*tensorleapapi.NewPublishProjectRequest(projectId, tarSignedUrl)).Execute()
	if err = api.CheckRes(res, err); err != nil {
		return fmt.Errorf("failed to publish project: %v", err)
	}
	return nil
}

// Not using apiClient because of the response type (file)
func ExportProject(ctx context.Context, projectId string) (*http.Response, error) {
	baseUrl, apiKey := api.GetAuthFromContext(ctx)
	url := fmt.Sprintf("%s/projects/downloadProject/%s", baseUrl, projectId)
	req, _ := http.NewRequest("GET", url, nil)

	// steam response
	req.Header.Set("Accept", "application/octet-stream")
	api.AddAuthToRequestHeader(&req.Header, apiKey)
	client := &http.Client{}
	res, err := client.Do(req)
	if err = api.CheckRes(res, err); err != nil {
		return nil, fmt.Errorf("failed to export project %s /n%s", projectId, err)
	}
	return res, nil
}

func getFileNameFromResponse(res *http.Response) string {
	contentDisposition := res.Header.Get("Content-Disposition")
	// e.g attachment; filename="mnist-v95.tar.zip"
	fileName := contentDisposition[len("attachment; filename="):]
	fileName = strings.Trim(fileName, "\"")
	return fileName
}
