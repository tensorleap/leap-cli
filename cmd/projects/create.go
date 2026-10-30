package projects

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/project"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func NewCreateCmd() *cobra.Command {
	projectDetails := &project.AddProjectDetails{}

	var cmd = &cobra.Command{
		Use:   "create [projectName]",
		Short: "Create new project",
		Long:  `Create new project`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			projects, err := project.GetProjects(ctx)
			if err != nil {
				return err
			}
			if len(args) > 0 {
				projectDetails.Name = args[0]
			}
			var newProject *project.ProjectEntity
			if len(projectDetails.Name) == 0 {
				newProject, err = project.AskAndAddProject(ctx, projectDetails, projects)
			} else {
				newProject, err = project.AddProject(ctx, projectDetails)
			}
			if err != nil {
				return err
			}

			return useInWorkspace(newProject.GetCid())
		},
	}

	cmd.Flags().StringVarP(&projectDetails.Description, "description", "d", "", "Project description")

	return cmd
}

// useInWorkspace points this folder's leap.yaml at the new project, so leap push doesn't ask which project to use
func useInWorkspace(projectId string) error {
	written, existingProjectId, err := workspace.SetProjectIdIfMissing(projectId)
	if err != nil {
		// the project exists by now, so this isn't a failed create
		log.Warnf("Could not set projectId %s in %s: %v", projectId, workspace.CONFIG_FILE_NAME, err)
		return nil
	}
	if written {
		log.Infof("Set projectId %s in %s", projectId, workspace.CONFIG_FILE_NAME)
	} else if existingProjectId != "" && existingProjectId != projectId {
		log.Infof("%s already uses project %s; to push to the new project, set projectId: %s", workspace.CONFIG_FILE_NAME, existingProjectId, projectId)
	}
	return nil
}

func init() {
	RootCommand.AddCommand(NewCreateCmd())
}
