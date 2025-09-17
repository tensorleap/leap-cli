package project

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

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

	fileStat, err := file.Stat()
	if err != nil {
		return err
	}

	getUrl, err := uploadProjectToTempFile(ctx, projectName, file, fileStat.Size())
	if err != nil {
		return err
	}
	return ImportProject(ctx, projectName, getUrl, &projectCtx.Meta)
}

func uploadProjectToTempFile(ctx context.Context, projectName string, projectReader io.Reader, projectSize int64) (string, error) {

	res, err := api.GetUploadSignedUrl(ctx, projectName)

	if err != nil {
		return "", fmt.Errorf("failed to get signed url for the uploaded project: %v", err)
	}
	uploadSingedUrl := res.GetUrl()
	uploadUrl := res.GetFileName()
	err = api.UploadFile(uploadSingedUrl, projectReader, projectSize)
	if err != nil {
		return "", fmt.Errorf("failed to upload project: %v", err)
	}
	log.Infof("Uploaded successfully project: %s", projectName)

	expireTime := time.Hour * 24
	getUrl, err := api.GetSignedUrl(ctx, uploadUrl, http.MethodGet, expireTime, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get singed url for the uploaded project: %v", err)
	}
	return getUrl, nil
}

func ExportProjectIntoFile(ctx context.Context, project *ProjectEntity, outputDir string, options ExportProjectParams) error {
	job, err := ExportProject(ctx, project.GetCid(), "", options)
	if err != nil {
		return err
	}

	exportParams := job.Params.ExportProjectParams
	exportUrl := exportParams.ExportUrl

	fileName := fmt.Sprintf("%s-%v.tar.gz", project.Name, exportParams.ProjectVersion)
	filePath := filepath.Join(outputDir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	s := log.NewSpinner(fmt.Sprintf("Exporting project '%s', into: '%s'", project.Name, filePath))
	s.Start()
	defer s.Stop()
	downloadUrl, err := api.GetDownloadSignedUrl(ctx, exportUrl)
	if err != nil {
		return fmt.Errorf("failed to get download signed url: %v", err)
	}
	err = api.DownloadFile(downloadUrl, file)
	if err != nil && err != io.EOF {
		defer os.Remove(filePath)
		return err
	}
	return nil
}

const TIMEOUT_FOR_PROJECT_PUBLISH = 4 * time.Hour

type ExportProjectParams struct {
	NoCache                bool
	ExcludeCalculatedFiles bool
}

func ExportProject(ctx context.Context, projectId string, copyToUrl string, options ExportProjectParams) (*tensorleapapi.Job, error) {

	exportParams := *tensorleapapi.NewExportProjectRequest(projectId)
	exportParams.SetNoCache(options.NoCache)

	exportOptions := tensorleapapi.NewExportOptions()
	exportOptions.SetExcludeCalculatedFiles(options.ExcludeCalculatedFiles)

	exportParams.SetOptions(*exportOptions)

	if len(copyToUrl) > 0 {
		log.Infof("Copy project content by signed url: %v", extractUrl(copyToUrl))
		exportParams.SetCopyToUrl(copyToUrl)
	}
	exportRes, res, err := api.ApiClient.ExportProject(ctx).ExportProjectRequest(exportParams).Execute()
	if err = api.CheckRes(res, err); err != nil {
		return nil, fmt.Errorf("failed to export project: %v", err)
	}

	err = api.WaitForCondition(ctx, "Export project...", func() (bool, error) {
		job, err := getExportProjectJobById(ctx, exportRes.JobId)
		if err != nil {
			return false, err
		}
		switch job.Status {
		case tensorleapapi.JOBSTATUS_FAILED:
			return false, fmt.Errorf("export project failed")
		case tensorleapapi.JOBSTATUS_FINISHED:
			return true, nil
		}

		return false, nil
	}, 20*time.Second, TIMEOUT_FOR_PROJECT_PUBLISH)

	if err != nil {
		return nil, fmt.Errorf("failed to wait for project to be copied: %v", err)
	}
	job, err := getExportProjectJobById(ctx, exportRes.JobId)
	if err != nil {
		return nil, fmt.Errorf("failed to get export job: %v", err)
	}
	return job, nil
}

func getExportProjectJobById(ctx context.Context, cid string) (*tensorleapapi.Job, error) {
	query := *tensorleapapi.NewGetJobsFilterParams()
	query.SetCid([]string{cid})

	res, orgRes, err := api.ApiClient.GetSlimJobs(ctx).GetJobsFilterParams(query).Execute()
	if err != nil {
		return nil, err
	}
	if len(res.Jobs) == 0 {
		return nil, fmt.Errorf("job not found")
	}
	var exportProjectJobs struct {
		Jobs []struct {
			Params tensorleapapi.ExportProjectParams `json:"params"`
		} `json:"jobs"`
	}
	data, err := io.ReadAll(orgRes.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}
	if err := json.Unmarshal(data, &exportProjectJobs); err != nil {
		return nil, err
	}
	job := res.Jobs[0]
	// override job params with export project params - openApi not deserializing correctly
	job.Params = &tensorleapapi.JobParams{
		ExportProjectParams: &exportProjectJobs.Jobs[0].Params,
	}
	return &job, nil
}

// Not using apiClient because of the response type (file)
func DownloadProject(ctx context.Context, projectId string) (*http.Response, error) {
	baseUrl, apiKey := api.GetAuthFromContext(ctx)
	url := fmt.Sprintf("%s/projects/downloadProject/%s", baseUrl, projectId)
	req, _ := http.NewRequest("GET", url, nil)

	// steam response
	req.Header.Set("Accept", "application/octet-stream")
	api.AddAuthToRequestHeader(&req.Header, apiKey)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err = api.CheckRes(res, err); err != nil {
		return nil, fmt.Errorf("failed to export project %s /n%s", projectId, err)
	}
	return res, nil
}
