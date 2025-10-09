package project

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func GetProjectFromProjectId(ctx context.Context, projectId string, askForNewProjectFirst bool) (project *ProjectEntity, wasCreated bool, err error) {
	projects, err := GetProjects(ctx)
	if err != nil {
		return nil, false, err
	}
	if len(projectId) > 0 {
		project, _ = entity.GetEntityById(projectId, projects, ProjectEntityDesc)
		if project == nil {
			log.Warnf("Project with ID '%s' not found, please select a valid project", projectId)
		}
	}

	if project == nil {
		project, wasCreated, err = SelectOrCreateProject(ctx, projects, askForNewProjectFirst)
		if err != nil {
			return nil, false, err
		}
	}
	return project, wasCreated, nil
}

func SelectOrCreateProject(ctx context.Context, projects []ProjectEntity, askForNewProjectFirst bool) (*ProjectEntity, bool, error) {
	createProject := func() (*ProjectEntity, error) {
		projectDetails := AddProjectDetails{}
		return AskAndAddProject(ctx, &projectDetails, projects)
	}
	return entity.SelectEntityOrCreateOne(projects, createProject, askForNewProjectFirst, true, ProjectEntityDesc)
}

func SyncProjectIdToWorkspaceConfig(ctx context.Context, workspaceConfig *workspace.WorkspaceConfig) (*ProjectEntity, error) {

	currentProject, _, err := GetProjectFromProjectId(ctx, workspaceConfig.ProjectId, true)

	if err != nil {
		return nil, err
	}
	if currentProject.Cid != workspaceConfig.ProjectId {

		log.Info("Updating projectId")
		workspaceConfig.ProjectId = currentProject.GetCid()
		err = workspace.SetWorkspaceConfig(workspaceConfig, "")
		if err != nil {
			return nil, err
		}
	}
	return currentProject, nil
}

func AskAndAddProject(ctx context.Context, projectDetails *AddProjectDetails, projects []ProjectEntity) (*ProjectEntity, error) {

	existingNames := entity.GetNames(projects, ProjectEntityDesc)
	err := AskForNewProject(projectDetails, existingNames)
	if err != nil {
		return nil, err
	}

	return AddProject(ctx, projectDetails)
}

func AskForNewProject(projectDetails *AddProjectDetails, existingProjectNames []string) error {

	// Create a survey with questions
	questions := []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "Project name:",
				Default: projectDetails.Name,
			},
			Validate: entity.CreateUniqueNameValidator(existingProjectNames),
		},
		{
			Name: "description",
			Prompt: &survey.Input{
				Message: "Project description (optional):",
				Default: projectDetails.Description,
			},
		},
	}

	// Perform the survey
	err := survey.Ask(questions, projectDetails)
	if err != nil {
		return err
	}

	return nil
}

func SelectProjectsToPublish(ctx context.Context, projectNamesFromArgs []string, allProjectFlag bool) ([]ProjectEntity, error) {
	projects, err := GetProjects(ctx)
	if err != nil {
		return nil, err
	}

	if allProjectFlag {
		selectedProjects := []ProjectEntity{}
		for _, project := range projects {
			isPublicProject := project.HubPublishPolicy == tensorleapapi.HUBPUBLISHPOLICY_PUBLIC
			if isPublicProject {
				selectedProjects = append(selectedProjects, project)
			}
		}
		return selectedProjects, nil
	}

	projectEntities, err := entity.SelectEntitiesByNamesOrAsk(projectNamesFromArgs, projects, ProjectEntityDesc)
	if err != nil {
		return nil, err
	}
	return projectEntities, nil
}
