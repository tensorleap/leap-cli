package code

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/secret"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func init() {
	var codeIntegrationId string
	var newCodeIntegrationName string
	var secretId string
	var branch string
	var pythonVersion string

	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Create initial code integration files in the current directory",
		Long:  `Create initial code integration files in the current directory`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(codeIntegrationId) == 0 && len(newCodeIntegrationName) == 0 {
				return errors.New("error: flag(s) \"codeId\" or \"new\" must be set")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			var codeIntegration *code.CodeIntegration = nil
			codeIntegrations, err := code.GetCodeIntegrations(ctx)
			if err != nil {
				return err
			}

			if len(newCodeIntegrationName) > 0 {
				newCodeIntegrationName, err = code.AskForCodeIntegrationNameIfExisted(newCodeIntegrationName, codeIntegrations)
				if err != nil {
					return err
				}
				codeIntegration, err = code.AddCodeIntegration(ctx, newCodeIntegrationName)
				if err != nil {
					return err
				}
				var selectedSecret *secret.SecretEntity
				if len(secretId) == 0 {
					selectedSecret, _, err = secret.CreateOrSelectSecretIfInUse(ctx)
				} else {
					selectedSecret, _, _, err = secret.CreateOrSelectIfSecretNotFound(ctx, secretId)
				}
				if err != nil {
					return err
				}
				if selectedSecret != nil {
					secretId = selectedSecret.GetCid()
				}

			} else if len(codeIntegrationId) > 0 {

				codeIntegration, err = entity.GetEntityById(codeIntegrationId, codeIntegrations, code.CodeIntegrationEntityDesc)
				if err != nil {
					return err
				}
				latestVersion, err := code.GetLatestVersion(ctx, codeIntegration.GetCid(), branch)
				if err != nil && err != code.ErrEmptyCodeIntegrationVersion {
					return err
				}
				if err == nil {
					if len(secretId) == 0 {
						secretId = latestVersion.Metadata.GetSecretManagerId()
					}
					if len(pythonVersion) == 0 {
						pythonVersion = latestVersion.GetGenericBaseImageType()
					}
				}
				if len(pythonVersion) == 0 {
					pythonVersion = latestVersion.GetGenericBaseImageType()
				}
			}

			pythonVersion, err = code.GetPythonVersionFromFlag(cmd.Context(), pythonVersion)
			if err != nil {
				return err
			}

			return workspace.CreateCodeTemplate(codeIntegration.GetCid(), "", secretId, branch, ".", pythonVersion)
		},
	}

	cmd.Flags().StringVar(&codeIntegrationId, "codeId", "", "Code integration id of existing dataset")
	cmd.Flags().StringVar(&newCodeIntegrationName, "new", "", "Name for new database")
	cmd.Flags().StringVar(&secretId, "secretId", "", "Secret manager id for new code integration")
	cmd.Flags().StringVarP(&branch, "branch", "b", "", "Branch for code integration")
	cmd.Flags().StringVarP(&pythonVersion, "pythonVersion", "p", "", "Python version for code integration")
	cmd.MarkFlagsMutuallyExclusive("new", "codeId")
	RootCommand.AddCommand(cmd)
}
