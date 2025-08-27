package projects

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/model"
	"github.com/tensorleap/leap-cli/pkg/project"
	"github.com/tensorleap/leap-cli/pkg/secret"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func NewPushCmd() *cobra.Command {

	var secretId string
	var modelVersionName string
	var codeVersionMessage string
	var modelType string
	var modelBranch string
	var codeBranch string
	var transformInput bool
	var force bool
	var noWait bool
	var pythonVersion string
	var leapMappingPath string

	var cmd = &cobra.Command{
		Use:   "push <modelPath>",
		Short: "Push new version into a project with its model and code integration",
		Long:  `Push new version into a project with its model and code integration`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("missing model path argument")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// Define base properties for all analytics events
			properties := map[string]interface{}{
				"secret_id":            secretId,
				"model_version_name":   modelVersionName,
				"code_version_message": codeVersionMessage,
				"model_type":           modelType,
				"model_branch":         modelBranch,
				"code_branch":          codeBranch,
				"transform_input":      transformInput,
				"force":                force,
				"no_wait":              noWait,
				"python_version":       pythonVersion,
				"leap_mapping_path":    leapMappingPath,
			}

			// Track projects push started
			if err := analytics.SendEvent(analytics.EventCliProjectsPushStarted, properties); err != nil {
				log.Warnf("Failed to track projects push start event: %v", err)
			}

			ctx := cmd.Context()
			modelPath := args[0]
			properties["model_path"] = modelPath

			err := model.SelectModelType(&modelType, modelPath)
			if err != nil {
				// Track projects push failed
				properties["error"] = err.Error()
				properties["stage"] = "select_model_type"
				if err := analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties); err != nil {
					log.Warnf("Failed to track projects push failure event: %v", err)
				}
				return err
			}
			err = model.InitMessage(&modelVersionName)
			if err != nil {
				// Track projects push failed
				properties["error"] = err.Error()
				properties["stage"] = "init_message"
				if err := analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties); err != nil {
					log.Warnf("Failed to track projects push failure event: %v", err)
				}
				return err
			}

			workspaceConfig, err := workspace.GetWorkspaceConfig()
			if err != nil {
				// Track projects push failed
				properties["error"] = err.Error()
				properties["stage"] = "get_workspace_config"
				if err := analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties); err != nil {
					log.Warnf("Failed to track projects push failure event: %v", err)
				}
				return err
			}

			currentProject, err := project.SyncProjectIdToWorkspaceConfig(ctx, workspaceConfig)
			if err != nil {
				// Track projects push failed
				properties["error"] = err.Error()
				properties["stage"] = "sync_project"
				if err := analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties); err != nil {
					log.Warnf("Failed to track projects push failure event: %v", err)
				}
				return err
			}

			codeIntegration, _, err := code.GetAndUpdateCodeIntegrationIfNotExists(ctx, workspaceConfig)
			if err != nil {
				// Track projects push failed
				properties["error"] = err.Error()
				properties["stage"] = "get_code_integration"
				if err := analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties); err != nil {
					log.Warnf("Failed to track projects push failure event: %v", err)
				}
				return err
			}

			codeIntegrationBranches := code.BranchesFromCodeIntegration(codeIntegration)
			codeBranch, err = code.SyncBranchFromFlagAndConfig(codeBranch, workspaceConfig, codeIntegrationBranches, codeIntegration.GetDefaultBranch())
			if err != nil {
				// Track projects push failed
				properties["error"] = err.Error()
				properties["stage"] = "sync_branch"
				if err := analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties); err != nil {
					log.Warnf("Failed to track projects push failure event: %v", err)
				}
				return err
			}

			secretId, err := secret.SyncSecretIdFromFlagAndConfig(ctx, secretId, workspaceConfig)
			if err != nil {
				// Track projects push failed
				properties["error"] = err.Error()
				properties["stage"] = "sync_secret"
				if err := analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties); err != nil {
					log.Warnf("Failed to track projects push failure event: %v", err)
				}
				return err
			}

			pythonVersion, err = code.SyncPythonVersionFromFlagAndConfig(ctx, pythonVersion, workspaceConfig)
			if err != nil {
				// Track projects push failed
				properties["error"] = err.Error()
				properties["stage"] = "sync_python_version"
				if err := analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties); err != nil {
					log.Warnf("Failed to track projects push failure event: %v", err)
				}
				return err
			}

			close, tarGzFile, err := code.BundleCodeIntoTempFile(".", workspaceConfig, leapMappingPath)
			if err != nil {
				// Track projects push failed
				properties["error"] = err.Error()
				properties["stage"] = "bundle_code"
				if err := analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties); err != nil {
					log.Warnf("Failed to track projects push failure event: %v", err)
				}
				return err
			}
			defer close()

			pushed, currentVersion, err := code.PushCode(ctx, force, codeIntegration.Cid, tarGzFile, workspaceConfig.EntryFile, secretId, codeBranch, codeVersionMessage, pythonVersion)
			if err != nil {
				// Track projects push failed
				properties["error"] = err.Error()
				properties["stage"] = "push_code"
				properties["code_integration_id"] = codeIntegration.Cid
				if err := analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties); err != nil {
					log.Warnf("Failed to track projects push failure event: %v", err)
				}
				return err
			}
			if pushed || code.IsCodeParsing(currentVersion) {

				if err != nil {
					return err
				}

				ok, codeIntegrationVersion, err := code.WaitForCodeIntegrationStatus(ctx, currentVersion.Cid)
				if err != nil {
					// Track projects push failed
					properties["error"] = err.Error()
					properties["stage"] = "wait_for_code_parsing"
					properties["code_integration_id"] = codeIntegration.Cid
					properties["version_id"] = currentVersion.Cid
					if err := analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties); err != nil {
						log.Warnf("Failed to track projects push failure event: %v", err)
					}
					return err
				}
				if ok {
					log.Info("Code parsed successfully")
				} else {
					code.PrintCodeIntegrationVersionParserErr(codeIntegrationVersion)
					// Track projects push failed due to code parsing failure
					properties["error"] = "code parsing failed"
					properties["stage"] = "code_parsing"
					properties["code_integration_id"] = codeIntegration.Cid
					properties["version_id"] = currentVersion.Cid
					if err := analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties); err != nil {
						log.Warnf("Failed to track projects push failure event: %v", err)
					}
					return fmt.Errorf("code parsing failed")
				}
			} else if code.IsCodeParseFailed(currentVersion) {
				code.PrintCodeIntegrationVersionParserErr(currentVersion)
				// Track projects push failed due to previous code parsing failure
				properties["error"] = "latest code parsing failed"
				properties["stage"] = "previous_code_parsing_failed"
				properties["code_integration_id"] = codeIntegration.Cid
				properties["version_id"] = currentVersion.Cid
				if err := analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties); err != nil {
					log.Warnf("Failed to track projects push failure event: %v", err)
				}
				return fmt.Errorf("latest code parsing failed, add --force to push anyway")
			}

			err = model.ImportModel(ctx, modelPath, currentProject.GetCid(), modelVersionName, modelType, modelBranch, codeIntegration.GetCid(), codeBranch, transformInput, !noWait)
			if err != nil {
				// Track projects push failed
				properties["error"] = err.Error()
				properties["stage"] = "import_model"
				properties["code_integration_id"] = codeIntegration.Cid
				properties["version_id"] = currentVersion.Cid
				properties["project_id"] = currentProject.GetCid()
				if err := analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties); err != nil {
					log.Warnf("Failed to track projects push failure event: %v", err)
				}
				return err
			}
			// Track projects push success
			properties["code_integration_id"] = codeIntegration.Cid
			properties["version_id"] = currentVersion.Cid
			properties["project_id"] = currentProject.GetCid()
			properties["final_secret_id"] = secretId
			properties["final_code_branch"] = codeBranch
			properties["final_python_version"] = pythonVersion
			properties["final_model_type"] = modelType
			properties["final_model_branch"] = modelBranch
			properties["final_transform_input"] = transformInput
			properties["final_wait"] = !noWait
			properties["code_pushed"] = pushed
			if err := analytics.SendEvent(analytics.EventCliProjectsPushSuccess, properties); err != nil {
				log.Warnf("Failed to track projects push success event: %v", err)
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&modelVersionName, "model-name", "m", "", "Model version name")
	cmd.Flags().StringVar(&codeVersionMessage, "code-message", "", "Code version message")
	cmd.Flags().StringVar(&modelType, "type", "", "Type is the type of the model file [JSON_TF2 / ONNX / PB_TF2 / H5_TF2]")
	cmd.Flags().StringVar(&modelBranch, "model-branch", "", "Name of the model branch [OPTIONAL]")
	cmd.Flags().StringVar(&codeBranch, "code-branch", "", "Name of the code branch [OPTIONAL]")
	cmd.Flags().StringVar(&secretId, "secretId", "", "Secret id")
	cmd.Flags().BoolVarP(&force, "force", "f", false, "Force push code integration")
	cmd.Flags().BoolVar(&transformInput, "transform-input", false, "Transpose the input data to channel-last format")
	cmd.Flags().BoolVar(&noWait, "no-wait", false, "Do not wait for push to complete")
	cmd.Flags().StringVar(&leapMappingPath, "leap-mapping", "", "Path to external leap mapping file")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewPushCmd())
}
