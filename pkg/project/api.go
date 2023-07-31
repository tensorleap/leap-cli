package project

import (
	"context"

	. "github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/entity"
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
	res, _, err := ApiClient.GetProjects(ctx).Execute()
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
	res, _, err := ApiClient.AddProject(ctx).ProjectMeta(*tensorleapapi.NewProjectMeta(
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
	_, err := ApiClient.DeleteProject(ctx).
		DeleteProjectParams(*tensorleapapi.NewDeleteProjectParams(project.GetCid())).
		Execute()
	if err != nil {
		return err
	}
	entity.InfoDeletion(project, ProjectEntityDesc)
	return nil
}
