package project

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tensorleap/leap-cli/pkg/entity"
)

func GetProjectFromFlag(ctx context.Context, projectIdFlag string, askForNewProjectFirst bool) (project *ProjectEntity, wasCreated bool, err error) {
	projects, err := GetProjects(ctx)
	if err != nil {
		return nil, false, err
	}
	if len(projectIdFlag) > 0 {
		project, err = entity.GetEntityById(projectIdFlag, projects, ProjectEntityDesc)
	} else {
		project, wasCreated, err = SelectOrCreateProject(ctx, projects, askForNewProjectFirst)
	}
	if err != nil {
		return nil, false, err
	}
	return project, wasCreated, nil
}

func SelectOrCreateProject(ctx context.Context, projects []ProjectEntity, askForNewProjectFirst bool) (*ProjectEntity, bool, error) {
	createProject := func() (*ProjectEntity, error) {
		projectDetails := AddProjectDetails{}
		return AskAndAddProject(ctx, &projectDetails, projects)
	}
	return entity.SelectEntityOrCreateOne(projects, createProject, askForNewProjectFirst, ProjectEntityDesc)
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
			Name:     "name",
			Prompt:   &survey.Input{Message: "Project name:"},
			Validate: entity.CreateUniqueNameValidator(existingProjectNames),
		},
		{
			Name:   "description",
			Prompt: &survey.Input{Message: "Project description (optional):"},
		},
	}

	// Perform the survey
	err := survey.Ask(questions, projectDetails)
	if err != nil {
		return err
	}

	return nil
}
