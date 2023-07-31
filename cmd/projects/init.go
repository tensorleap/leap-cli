package projects

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/project"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func NewInitCmd() *cobra.Command {
	var projectId string
	var codeIntegrationId string

	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Create initial project environment files in the current directory",
		Long:  `Create initial project environment files in the current directory`,
		PreRunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			selectedProject, err := project.GetProjectFromFlag(ctx, projectId, true)
			if err != nil {
				return err
			}

			codeIntegration, err := code.GetCodeIntegrationFromFlag(ctx, codeIntegrationId, true)
			if err != nil {
				return err
			}

			err = workspace.CreateCodeTemplate(
				codeIntegration.GetCid(),
				selectedProject.GetCid(),
				".",
			)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&projectId, "projectId", "", "ProjectId is the id of the project")
	cmd.Flags().StringVar(&projectId, "codeId", "", "CodeIntegrationId is the id of the code integration to bind")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewInitCmd())
}
