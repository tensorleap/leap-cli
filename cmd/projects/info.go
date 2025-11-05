package projects

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/project"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func NewInfoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Show project and code integration information",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			workspaceConfig, err := workspace.GetWorkspaceConfig()
			if err != nil {
				return fmt.Errorf("failed to load workspace config: %w", err)
			}

			if workspaceConfig.ProjectId == "" {
				log.Info("No project configured in workspace")
			} else {
				projects, err := project.GetProjects(ctx)
				if err != nil {
					return fmt.Errorf("failed to fetch projects: %w", err)
				}
				proj, _ := entity.GetEntityById(workspaceConfig.ProjectId, projects, project.ProjectEntityDesc)
				if proj != nil {
					fmt.Printf("Project Name: %s\n", project.ProjectEntityDesc.GetDisplayName(proj))
				} else {
					fmt.Printf("Project ID: %s (not found)\n", workspaceConfig.ProjectId)
				}
			}

			return nil
		},
	}
}

func init() {
	RootCommand.AddCommand(NewInfoCmd())
}
