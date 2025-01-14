package projects

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/project"
	"github.com/tensorleap/leap-cli/pkg/secret"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func NewInitCmd() *cobra.Command {
	var projectId string
	var secretId string
	var codeIntegrationId string
	var codeIntegrationBranch string

	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Create initial project environment files in the current directory",
		Long:  `Create initial project environment files in the current directory`,
		PreRunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			selectedProject, _, err := project.GetProjectFromFlag(ctx, projectId, true)
			if err != nil {
				return err
			}

			codeIntegration, wasCreatedCodeIntegration, err := code.GetCodeIntegrationFromFlag(ctx, codeIntegrationId, true)
			if err != nil {
				return err
			}

			isCreatingEmptyTemplate := true
			if !wasCreatedCodeIntegration {
				latestVersion, err := code.GetLatestVersion(ctx, codeIntegration.GetCid(), codeIntegrationBranch)
				if err != nil && code.ErrEmptyCodeIntegrationVersion != err {
					return err
				} else if err == nil {
					if secretId == "" {
						secretId = latestVersion.Metadata.GetSecretManagerId()
					}
					isCreatingEmptyTemplate = false
					files, err := code.CloneCodeIntegrationVersion(ctx, latestVersion, ".", "")
					if err != nil {
						return err
					}

					err = workspace.OverrideWorkspaceConfig(
						codeIntegration.GetCid(),
						selectedProject.GetCid(),
						latestVersion.GetCodeEntryFile(),
						secretId,
						latestVersion.Branch,
						files,
						".",
					)

					if err != nil {
						return err
					}
				}
			}

			if wasCreatedCodeIntegration && secretId == "" {
				selectedSecret, _, err := secret.CreateOrSelectSecretIfInUse(ctx)
				if err != nil {
					return err
				}
				if selectedSecret != nil {
					secretId = selectedSecret.GetCid()
				}
			}

			if isCreatingEmptyTemplate {
				err = workspace.CreateCodeTemplate(
					codeIntegration.GetCid(),
					selectedProject.GetCid(),
					secretId,
					codeIntegrationBranch,
					".",
				)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&projectId, "projectId", "", "ProjectId is the id of the project")
	cmd.Flags().StringVar(&projectId, "codeId", "", "CodeIntegrationId is the id of the code integration to bind")
	cmd.Flags().StringVar(&codeIntegrationBranch, "branch", "", "Branch of the code integration to bind")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewInitCmd())
}
