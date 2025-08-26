package code

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/log"
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

			// Define base properties for all analytics events
			properties := map[string]interface{}{
				"code_id":               codeIntegrationId,
				"code_integration_name": newCodeIntegrationName,
				"secret_id":             secretId,
				"branch":                branch,
				"python_version":        pythonVersion,
			}

			// Track code init started
			if err := analytics.SendEvent(analytics.EventCliCodeInitStarted, properties); err != nil {
				log.Warnf("Failed to track code init start event: %v", err)
			}

			var codeIntegration *code.CodeIntegration = nil
			codeIntegrations, err := code.GetCodeIntegrations(ctx)
			if err != nil {
				// Track code init failed
				properties["error"] = err.Error()
				properties["stage"] = "get_code_integrations"
				if err := analytics.SendEvent(analytics.EventCliCodeInitFailed, properties); err != nil {
					log.Warnf("Failed to track code init failure event: %v", err)
				}
				return err
			}

			if len(newCodeIntegrationName) > 0 {
				newCodeIntegrationName, err = code.AskForCodeIntegrationNameIfExisted(newCodeIntegrationName, codeIntegrations)
				if err != nil {
					// Track code init failed
					properties["error"] = err.Error()
					properties["stage"] = "ask_code_integration_name"
					if err := analytics.SendEvent(analytics.EventCliCodeInitFailed, properties); err != nil {
						log.Warnf("Failed to track code init failure event: %v", err)
					}
					return err
				}
				codeIntegration, err = code.AddCodeIntegration(ctx, newCodeIntegrationName)
				if err != nil {
					// Track code init failed
					properties["error"] = err.Error()
					properties["stage"] = "add_code_integration"
					if err := analytics.SendEvent(analytics.EventCliCodeInitFailed, properties); err != nil {
						log.Warnf("Failed to track code init failure event: %v", err)
					}
					return err
				}
				var selectedSecret *secret.SecretEntity
				if len(secretId) == 0 {
					selectedSecret, _, err = secret.CreateOrSelectSecretIfInUse(ctx)
				} else {
					selectedSecret, _, _, err = secret.CreateOrSelectIfSecretNotFound(ctx, secretId)
				}
				if err != nil {
					// Track code init failed
					properties["error"] = err.Error()
					properties["stage"] = "secret_creation_selection"
					if err := analytics.SendEvent(analytics.EventCliCodeInitFailed, properties); err != nil {
						log.Warnf("Failed to track code init failure event: %v", err)
					}
					return err
				}
				if selectedSecret != nil {
					secretId = selectedSecret.GetCid()
				}

			} else if len(codeIntegrationId) > 0 {

				codeIntegration, err = entity.GetEntityById(codeIntegrationId, codeIntegrations, code.CodeIntegrationEntityDesc)
				if err != nil {
					// Track code init failed
					properties["error"] = err.Error()
					properties["stage"] = "get_existing_code_integration"
					properties["requested_code_id"] = codeIntegrationId
					if err := analytics.SendEvent(analytics.EventCliCodeInitFailed, properties); err != nil {
						log.Warnf("Failed to track code init failure event: %v", err)
					}
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
				// Track code init failed
				properties["error"] = err.Error()
				properties["stage"] = "python_version_validation"
				if err := analytics.SendEvent(analytics.EventCliCodeInitFailed, properties); err != nil {
					log.Warnf("Failed to track code init failure event: %v", err)
				}
				return err
			}

			err = workspace.CreateCodeTemplate(codeIntegration.GetCid(), "", secretId, branch, ".", pythonVersion)
			if err != nil {
				// Track code init failed
				properties["error"] = err.Error()
				properties["stage"] = "template_creation"
				properties["code_integration_id"] = codeIntegration.GetCid()
				if err := analytics.SendEvent(analytics.EventCliCodeInitFailed, properties); err != nil {
					log.Warnf("Failed to track code init failure event: %v", err)
				}
				return err
			}

			// Track code init success
			properties["code_integration_id"] = codeIntegration.GetCid()
			properties["final_python_version"] = pythonVersion
			if err := analytics.SendEvent(analytics.EventCliCodeInitSuccess, properties); err != nil {
				log.Warnf("Failed to track code init success event: %v", err)
			}

			return nil
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
