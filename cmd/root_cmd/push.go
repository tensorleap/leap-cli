package root_cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/model"
	"github.com/tensorleap/leap-cli/pkg/project"
	"github.com/tensorleap/leap-cli/pkg/secret"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func NewPushCmd() *cobra.Command {

	var secretId string
	var modelVersionName string
	var codeVersionMessage string
	var modelType string
	var modelPath string
	var branch string
	var transformInput bool
	var noWait bool
	var pythonVersion string
	var runEval bool
	var batchSize int
	var overwriteVersionRef string
	var updateParts []string

	var cmd = &cobra.Command{
		Use:   "push",
		Short: "Push new or overwrite model version",
		Long: `Push new or overwrite model version into a project with its code integration.

Examples:
  # Push a model
  leap push

  # Push and run evaluation after
  leap push -e

  # Overwrite a version by id (non-interactive) and refresh metadata after eval
  leap push --overwrite-version <versionId> -e -u metadata
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if batchSize > 0 && !runEval {
				return fmt.Errorf("--batch requires --eval")
			}
			if len(updateParts) > 0 && !runEval {
				return fmt.Errorf("--update (-u) requires --eval (-e)")
			}

			ctx, _, err := auth.RequireAuth(cmd.Context())
			if err != nil {
				return err
			}

			properties := map[string]interface{}{
				"secret_id":               secretId,
				"model_version_name":      modelVersionName,
				"code_version_message":    codeVersionMessage,
				"model_type":              modelType,
				"branch":                  branch,
				"transform_input":         transformInput,
				"no_wait":                 noWait,
				"python_version":          pythonVersion,
				"model_path":              modelPath,
				"overwrite_version_ref":   overwriteVersionRef,
				"update_parts":            updateParts,
			}

			analytics.SendEvent(analytics.EventCliProjectsPushStarted, properties)

			workspaceConfig, err := workspace.GetWorkspaceConfig()
			if err != nil {
				properties["error"] = err.Error()
				properties["stage"] = "get_workspace_config"
				analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
				return err
			}
			currentProject, err := project.SyncProjectIdToWorkspaceConfig(ctx, workspaceConfig)
			if err != nil {
				properties["error"] = err.Error()
				properties["stage"] = "sync_project"
				analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
				return err
			}

			var isOverwrite bool
			var overwriteVersionInfo *model.VersionInfo
			if overwriteVersionRef != "" {
				overwriteVersionInfo, err = model.ResolveVersionInfoFromRef(ctx, currentProject.GetCid(), overwriteVersionRef)
				if err != nil {
					return err
				}
				isOverwrite = overwriteVersionInfo.HasModel || overwriteVersionInfo.HasUploadedModel
				if !isOverwrite && len(modelPath) == 0 {
					return fmt.Errorf("version %q has no model; set --model-path (-m)", overwriteVersionRef)
				}
			} else if len(modelPath) == 0 {
				isOverwrite, overwriteVersionInfo, modelPath, err = model.AskUserForModelPathOrOverwrite(ctx, currentProject.GetCid())
				if err != nil {
					return err
				}
			} else {
				isOverwrite = false
			}

			properties["is_overwrite"] = isOverwrite
			var overwriteVersionId string
			var overwriteVersionHasModel bool
			if overwriteVersionInfo != nil {
				properties["version_id"] = overwriteVersionInfo.VersionId
				properties["version_name"] = overwriteVersionInfo.VersionName
				overwriteVersionId = overwriteVersionInfo.VersionId
				overwriteVersionHasModel = overwriteVersionInfo.HasModel
				properties["version_has_model"] = overwriteVersionHasModel

			}
			properties["model_path"] = modelPath

			overwriteVersionHasUploadedModel := false
			if overwriteVersionInfo != nil {
				overwriteVersionHasUploadedModel = overwriteVersionInfo.HasUploadedModel
			}

			needsNewModel := !isOverwrite || (!overwriteVersionHasModel && !overwriteVersionHasUploadedModel)
			if needsNewModel {
				err := model.SelectModelType(&modelType, modelPath)
				if err != nil {
					properties["error"] = err.Error()
					properties["stage"] = "select_model_type"
					analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
					return err
				}
			}
			if overwriteVersionInfo == nil {
				defaultMessage := model.GetDefaultMessageFromModelPath(modelPath)
				err = model.InitMessage(&modelVersionName, defaultMessage)
				if err != nil {
					properties["error"] = err.Error()
					properties["stage"] = "init_message"
					analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
					return err
				}
			} else {
				modelVersionName = overwriteVersionInfo.VersionName
			}

			branch, err = code.SyncBranchFromFlagAndConfig(branch, workspaceConfig)
			if err != nil {
				properties["error"] = err.Error()
				properties["stage"] = "sync_branch"
				analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
				return err
			}

			secretId, err := secret.SyncSecretIdFromFlagAndConfig(ctx, secretId, workspaceConfig)
			if err != nil {
				properties["error"] = err.Error()
				properties["stage"] = "sync_secret"
				analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
				return err
			}

			pythonVersion, err = code.SyncPythonVersionFromFlagAndConfig(ctx, pythonVersion, workspaceConfig)
			if err != nil {
				properties["error"] = err.Error()
				properties["stage"] = "sync_python_version"
				analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
				return err
			}

			var evalBatchSize int
			var updateActions []tensorleapapi.UpdateAction
			runUpdateEvaluate := false
			if len(updateParts) > 0 && !isOverwrite {
				return fmt.Errorf("--update (-u) only applies when overwriting an existing version (use --overwrite-version or choose overwrite in the prompt)")
			}
			if runEval {
				askBatchSize := func() error {
					if batchSize > 0 {
						evalBatchSize = batchSize
						return nil
					}
					defaultBatchSize, err := model.GetLatestEvaluateBatchSize(ctx, currentProject.GetCid())
					if err != nil {
						log.Warnf("failed to get latest evaluate batch size: %v", err)
					}
					evalBatchSize, err = model.AskForBatchSize(defaultBatchSize)
					if err != nil {
						return fmt.Errorf("failed to get batch size: %w", err)
					}
					return nil
				}

				if isOverwrite {
					updateActions, err = model.ParseUpdateActionsFromFlags(updateParts)
					if err != nil {
						return err
					}
					if len(updateActions) > 0 {
						runUpdateEvaluate = true
					} else {
						hasEvalData, evalErr := model.HasEvaluatedAncestorOrSelf(ctx, currentProject.GetCid(), overwriteVersionId)
						if evalErr != nil {
							log.Warnf("failed to check evaluation data for override chain: %v", evalErr)
							hasEvalData = true
						}
						if !hasEvalData {
							log.Info("No evaluation data found in the override chain — running a fresh evaluate.")
							if err := askBatchSize(); err != nil {
								return err
							}
						} else {
							plan, planErr := model.AskForEvaluatePlan()
							if planErr != nil {
								return fmt.Errorf("failed to get update plan: %w", planErr)
							}
							if plan.Kind == model.EvaluatePlanReset {
								if err := askBatchSize(); err != nil {
									return err
								}
							} else {
								updateActions = plan.UpdateActions
								runUpdateEvaluate = true
							}
						}
					}
				} else {
					if err := askBatchSize(); err != nil {
						return err
					}
				}
			}

			close, tarGzFile, err := code.BundleCodeIntoTempFile(".", workspaceConfig)
			if err != nil {
				properties["error"] = err.Error()
				properties["stage"] = "bundle_code"
				analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
				return err
			}
			defer close()

			pushed, codeSnapshotResponse, err := code.PushCode(ctx, tarGzFile, workspaceConfig.EntryFile, secretId, pythonVersion, modelVersionName, currentProject.GetCid(), branch, overwriteVersionId)
			if err != nil {
				properties["error"] = err.Error()
				properties["stage"] = "push_code"
				if codeSnapshotResponse != nil {
					properties["code_snapshot_id"] = codeSnapshotResponse.CodeSnapshot.Cid
					properties["version_id"] = codeSnapshotResponse.VersionId
				}
				analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
				return err
			}
			if pushed || !code.IsCodeEnded(&codeSnapshotResponse.CodeSnapshot) {

				ok, err := code.WaitForCodeIntegrationStatus(ctx, currentProject.GetCid(), codeSnapshotResponse.CodeSnapshot.Cid)
				if err != nil {
					properties["error"] = err.Error()
					properties["stage"] = "wait_for_code_parsing"
					properties["code_snapshot_id"] = codeSnapshotResponse.CodeSnapshot.Cid
					properties["version_id"] = codeSnapshotResponse.VersionId
					analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
					return err
				}
				if ok {
					log.Info("Code parsed successfully")
				} else {
					properties["error"] = "code parsing failed"
					properties["stage"] = "code_parsing"
					properties["code_snapshot_id"] = codeSnapshotResponse.CodeSnapshot.Cid
					properties["version_id"] = codeSnapshotResponse.VersionId
					analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
					return fmt.Errorf("code parsing failed")
				}
			} else if code.IsCodeParseFailed(&codeSnapshotResponse.CodeSnapshot) {
				properties["error"] = "latest code parsing failed"
				properties["stage"] = "previous_code_parsing_failed"
				properties["code_snapshot_id"] = codeSnapshotResponse.CodeSnapshot.Cid
				properties["version_id"] = codeSnapshotResponse.VersionId
				analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
				return fmt.Errorf("latest code parsing failed, add --force to push anyway")
			}
			if !isOverwrite {
				importModelInfo, err := model.PrepareImportModelFromFilePath(ctx, currentProject.GetCid(), modelPath, transformInput, modelType)
				if err != nil {
					return err
				}
				_, err = model.ImportModel(ctx, currentProject.GetCid(), codeSnapshotResponse.VersionId, importModelInfo, !noWait)
				if err != nil {
					properties["error"] = err.Error()
					properties["stage"] = "import_model"
					properties["code_snapshot_id"] = codeSnapshotResponse.CodeSnapshot.Cid
					properties["version_id"] = codeSnapshotResponse.VersionId
					analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
					return err
				}
			} else {
				var importModelInfo *tensorleapapi.ImportModelInfo
				if !overwriteVersionHasModel && !overwriteVersionHasUploadedModel {
					importModelInfo, err = model.PrepareImportModelFromFilePath(ctx, currentProject.GetCid(), modelPath, transformInput, modelType)
					if err != nil {
						return err
					}
				}
				_, err := model.OverrideModel(ctx, currentProject.GetCid(), codeSnapshotResponse.VersionId, !noWait, importModelInfo)
				if err != nil {
					properties["error"] = err.Error()
					properties["stage"] = "override_model"
					properties["code_snapshot_id"] = codeSnapshotResponse.CodeSnapshot.Cid
					properties["version_id"] = codeSnapshotResponse.VersionId
					analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)

					return err
				}
			}

			properties["code_snapshot_id"] = codeSnapshotResponse.CodeSnapshot.Cid
			properties["version_id"] = codeSnapshotResponse.VersionId
			properties["project_id"] = currentProject.GetCid()
			properties["final_secret_id"] = secretId
			properties["final_python_version"] = pythonVersion
			properties["final_model_type"] = modelType
			properties["final_transform_input"] = transformInput
			properties["final_wait"] = !noWait
			properties["code_pushed"] = pushed
			analytics.SendEvent(analytics.EventCliProjectsPushSuccess, properties)

			if runEval {
				if runUpdateEvaluate {
					err = model.RunUpdateEvaluateArtifact(ctx, currentProject.GetCid(), codeSnapshotResponse.VersionId, updateActions)
					if err != nil {
						return fmt.Errorf("failed to run update evaluate: %w", err)
					}
				} else {
					err = model.RunEvaluate(ctx, currentProject.GetCid(), codeSnapshotResponse.VersionId, evalBatchSize)
					if err != nil {
						return fmt.Errorf("failed to run evaluation: %w", err)
					}
				}
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&modelVersionName, "name", "n", "", "Model version name")
	cmd.Flags().StringVar(&modelType, "type", "", "Type is the type of the model file [JSON_TF2 / ONNX / PB_TF2 / H5_TF2]")
	cmd.Flags().StringVar(&branch, "branch", "", "Name of the branch [OPTIONAL]")
	cmd.Flags().StringVar(&secretId, "secretId", "", "Secret id")
	cmd.Flags().BoolVar(&transformInput, "transform-input", false, "Transpose the input data to channel-last format")
	cmd.Flags().BoolVar(&noWait, "no-wait", false, "Do not wait for push to complete")
	cmd.Flags().StringVarP(&modelPath, "model-path", "m", "", "Path to the model file")
	cmd.Flags().BoolVarP(&runEval, "eval", "e", false, "Run evaluation on the model after push completes")
	cmd.Flags().IntVar(&batchSize, "batch", 0, "Batch size for evaluation (only valid with --eval)")
	cmd.Flags().StringVar(&overwriteVersionRef, "overwrite-version", "", "Overwrite this existing version (version id from UI/API; a name is accepted only if exactly one version in the project has that name)")
	cmd.Flags().StringSliceVarP(&updateParts, "update", "u", nil, "With --eval on overwrite: artifact(s) to refresh (repeatable). Values: metadata, metric_config, insights, visualizations")
	return cmd
}
