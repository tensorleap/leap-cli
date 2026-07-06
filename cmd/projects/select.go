package projects

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/project"
	"github.com/tensorleap/leap-cli/pkg/secret"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func NewSelectCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "select [projectId]",
		Short: "Select project and code integration",
		Long: `Select a project and code integration to work with. This will update the workspace configuration.

Pass a project id as an argument to select it non-interactively (no prompts):
  leap projects select 6a3c2eb0b186171463d36eb6`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			workspaceConfig, err := workspace.GetWorkspaceConfig()
			if err != nil {
				return fmt.Errorf("failed to load workspace config: %w", err)
			}

			projectId := ""
			if len(args) > 0 {
				projectId = args[0]
			}

			var selectedProject *project.ProjectEntity
			projectWasCreated := false
			if projectId != "" {
				// Explicit id: resolve it directly and fail if it's unknown, so a
				// headless `select <id>` reports a bad id instead of silently
				// dropping into the interactive picker.
				projects, listErr := project.GetProjects(ctx)
				if listErr != nil {
					return fmt.Errorf("failed to get projects: %w", listErr)
				}
				found, findErr := entity.GetEntityById(projectId, projects, project.ProjectEntityDesc)
				if findErr != nil {
					return fmt.Errorf("project with id %q not found", projectId)
				}
				selectedProject = found
			} else {
				// No arg: interactive create-or-select picker.
				sp, wasCreated, selErr := project.GetProjectFromProjectId(ctx, "", true)
				if selErr != nil {
					return fmt.Errorf("failed to select or create project: %w", selErr)
				}
				selectedProject = sp
				projectWasCreated = wasCreated
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

			// When a project id is passed explicitly, keep the flow non-interactive:
			// preserve the workspace's existing python version instead of prompting.
			if projectId == "" {
				baseImages, defaultVersionId, err := code.GetPythonVersions(ctx)
				if err != nil {
					return fmt.Errorf("failed to get python versions: %w", err)
				}

				if workspaceConfig.PythonVersion != "" {
					defaultVersionId = workspaceConfig.PythonVersion
				}

				workspaceConfig.PythonVersion, err = code.AskForPythonVersion(defaultVersionId, baseImages)
				if err != nil {
					return fmt.Errorf("failed to sync python version: %w", err)
				}
			}

			err = workspace.SetWorkspaceConfig(workspaceConfig, ".")
			if err != nil {
				return fmt.Errorf("failed to update workspace configuration: %w", err)
			}

			projectAction := "selected"
			if projectWasCreated {
				projectAction = "created"
			}

			log.Infof("Successfully %s project '%s'",
				projectAction,
				project.ProjectEntityDesc.GetDisplayName(selectedProject),
			)
			return nil
		},
	}
	return cmd
}

func init() {
	RootCommand.AddCommand(NewSelectCmd())
}
