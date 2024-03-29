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
			var err error
			if len(newCodeIntegrationName) > 0 {
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
				codeIntegrations, err := code.GetCodeIntegrations(ctx)
				if err != nil {
					return err
				}
				codeIntegration, err = entity.GetEntityById(codeIntegrationId, codeIntegrations, code.CodeIntegrationEntityDesc)
				if err != nil {
					return err
				}
				if len(secretId) == 0 {
					latestVersion, err := code.GetLatestVersion(ctx, codeIntegration.GetCid())
					if err != nil && err != code.ErrEmptyCodeIntegrationVersion {
						return err
					}
					secretId = latestVersion.Metadata.GetSecretManagerId()
				}
			}
			return workspace.CreateCodeTemplate(codeIntegration.GetCid(), "", secretId, ".")
		},
	}

	cmd.Flags().StringVar(&codeIntegrationId, "codeId", "", "Code integration id of existing dataset")
	cmd.Flags().StringVar(&newCodeIntegrationName, "new", "", "Name for new database")
	cmd.Flags().StringVar(&secretId, "secretId", "", "Secret manager id for new code integration")
	cmd.MarkFlagsMutuallyExclusive("new", "codeId")
	RootCommand.AddCommand(cmd)
}
