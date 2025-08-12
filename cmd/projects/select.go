package projects

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/project"
	"github.com/tensorleap/leap-cli/pkg/secret"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func NewSelectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "select",
		Short: "Select project and code integration",
		Long:  `Select a project and code integration to work with. This will update the workspace configuration.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			workspaceConfig, err := workspace.GetWorkspaceConfig()
			if err != nil {
				return fmt.Errorf("failed to load workspace config: %w", err)
			}

			projects, err := project.GetProjects(ctx)
			if err != nil {
				return fmt.Errorf("failed to get projects: %w", err)
			}
			selectedProject, projectWasCreated, err := project.SelectOrCreateProject(ctx, projects, true)
			if err != nil {
				return fmt.Errorf("failed to select or create project: %w", err)
			}

			codeIntegrations, err := code.GetCodeIntegrations(ctx)
			if err != nil {
				return fmt.Errorf("failed to get code integrations: %w", err)
			}
			selectedCodeIntegration, codeWasCreated, err := code.SelectOrCreateCodeIntegration(ctx, codeIntegrations, true)
			if err != nil {
				return fmt.Errorf("failed to select or create code integration: %w", err)
			}

			if workspaceConfig.SecretId != "" {
				selectedSecret, wasValid, wasCreated, err := secret.CreateOrSelectIfSecretNotFound(ctx, workspaceConfig.SecretId)
				if err != nil {
					return fmt.Errorf("failed to handle secret: %w", err)
				}
				if selectedSecret != nil {
					workspaceConfig.SecretId = selectedSecret.GetCid()
					if wasCreated {
						log.Infof("Created new secret: %s", selectedSecret.GetName())
					} else if !wasValid {
						log.Infof("Selected secret: %s", selectedSecret.GetName())
					}
				} else {
					workspaceConfig.SecretId = ""
				}
			}

			workspaceConfig.ProjectId = selectedProject.GetCid()
			workspaceConfig.CodeIntegrationId = selectedCodeIntegration.GetCid()

			err = workspace.SetWorkspaceConfig(workspaceConfig, ".")
			if err != nil {
				return fmt.Errorf("failed to update workspace configuration: %w", err)
			}

			projectAction := "selected"
			if projectWasCreated {
				projectAction = "created"
			}
			codeAction := "selected"
			if codeWasCreated {
				codeAction = "created"
			}

			log.Infof("Successfully %s project '%s' and %s code integration '%s'",
				projectAction,
				project.ProjectEntityDesc.GetDisplayName(selectedProject),
				codeAction,
				code.CodeIntegrationEntityDesc.GetDisplayName(selectedCodeIntegration))

			return nil
		},
	}
}

func init() {
	RootCommand.AddCommand(NewSelectCmd())
}
