package projects

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/project"
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
			if len(projectDetails.Name) == 0 {
				_, err = project.AskAndAddProject(ctx, projectDetails, projects)
			} else {
				_, err = project.AddProject(ctx, projectDetails)
			}
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&projectDetails.Description, "description", "d", "", "Project description")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewCreateCmd())
}
