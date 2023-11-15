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

func ImportProject(ctx context.Context, projectName, importUrl, bgImageUrl string, versionMeta *hub.HubProjectVersion) error {
	reqParams := tensorleapapi.NewImportProjectRequest(
		projectName,
		versionMeta.Description,
		versionMeta.Tags,
		tensorleapapi.HUBPUBLISHPOLICY_NO_PUBLISH,
		importUrl,
	)
	reqParams.SetCategories(versionMeta.Categories)
	reqParams.SetBgImageUrl(bgImageUrl)

	_, err := api.ApiClient.ImportProject(ctx).ImportProjectRequest(*reqParams).Execute()
	if err != nil {
		return err
	}
	log.Infof("Import project %s started", projectName)
	return nil
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
	_, err := api.ApiClient.PublishProject(ctx).PublishProjectRequest(*tensorleapapi.NewPublishProjectRequest(projectId, tarSignedUrl)).Execute()
	return err
}

// Not using apiClient because of the response type (file)
func ExportProject(ctx context.Context, projectId string) (*http.Response, error) {
	baseUrl, apiKey := api.GetAuthFromContext(ctx)
	url := fmt.Sprintf("%s/projects/downloadProject/%s", baseUrl, projectId)
	req, _ := http.NewRequest("GET", url, nil)

	api.AddAuthToRequestHeader(&req.Header, apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
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
