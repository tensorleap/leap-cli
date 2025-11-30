package push

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/analytics"
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

	var cmd = &cobra.Command{
		Use:   "push",
		Short: "Push new or overwrite model version",
		Long:  `Push new or overwrite model version into a project with its code integration`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Define base properties for all analytics events
			properties := map[string]interface{}{
				"secret_id":            secretId,
				"model_version_name":   modelVersionName,
				"code_version_message": codeVersionMessage,
				"model_type":           modelType,
				"branch":               branch,
				"transform_input":      transformInput,
				"no_wait":              noWait,
				"python_version":       pythonVersion,
				"model_path":           modelPath,
			}

			// Track projects push started
			analytics.SendEvent(analytics.EventCliProjectsPushStarted, properties)

			workspaceConfig, err := workspace.GetWorkspaceConfig()
			if err != nil {
				// Track projects push failed
				properties["error"] = err.Error()
				properties["stage"] = "get_workspace_config"
				analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
				return err
			}
			ctx := cmd.Context()
			currentProject, err := project.SyncProjectIdToWorkspaceConfig(ctx, workspaceConfig)
			if err != nil {
				// Track projects push failed
				properties["error"] = err.Error()
				properties["stage"] = "sync_project"
				analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
				return err
			}

			var isOverwrite bool
			var overwriteVersionInfo *model.VersionInfo
			if len(modelPath) == 0 {
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

			if !isOverwrite || !overwriteVersionHasModel {
				err := model.SelectModelType(&modelType, modelPath)
				if err != nil {
					// Track projects push failed
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
					// Track projects push failed
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
				// Track projects push failed
				properties["error"] = err.Error()
				properties["stage"] = "sync_branch"
				analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
				return err
			}

			secretId, err := secret.SyncSecretIdFromFlagAndConfig(ctx, secretId, workspaceConfig)
			if err != nil {
				// Track projects push failed
				properties["error"] = err.Error()
				properties["stage"] = "sync_secret"
				analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
				return err
			}

			pythonVersion, err = code.SyncPythonVersionFromFlagAndConfig(ctx, pythonVersion, workspaceConfig)
			if err != nil {
				// Track projects push failed
				properties["error"] = err.Error()
				properties["stage"] = "sync_python_version"
				analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
				return err
			}

			close, tarGzFile, err := code.BundleCodeIntoTempFile(".", workspaceConfig)
			if err != nil {
				// Track projects push failed
				properties["error"] = err.Error()
				properties["stage"] = "bundle_code"
				analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
				return err
			}
			defer close()

			pushed, codeSnapshotResponse, err := code.PushCode(ctx, tarGzFile, workspaceConfig.EntryFile, secretId, pythonVersion, modelVersionName, currentProject.GetCid(), branch, overwriteVersionId)
			if err != nil {
				// Track projects push failed
				properties["error"] = err.Error()
				properties["stage"] = "push_code"
				properties["code_snapshot_id"] = codeSnapshotResponse.CodeSnapshot.Cid
				properties["version_id"] = codeSnapshotResponse.VersionId
				analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
				return err
			}
			if pushed || !code.IsCodeEnded(&codeSnapshotResponse.CodeSnapshot) {

				ok, codeSnapshot, err := code.WaitForCodeIntegrationStatus(ctx, currentProject.GetCid(), codeSnapshotResponse.CodeSnapshot.Cid)
				if err != nil {
					// Track projects push failed
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
					code.PrintCodeSnapshotParserErr(codeSnapshot)
					// Track projects push failed due to code parsing failure
					properties["error"] = "code parsing failed"
					properties["stage"] = "code_parsing"
					properties["code_snapshot_id"] = codeSnapshotResponse.CodeSnapshot.Cid
					properties["version_id"] = codeSnapshotResponse.VersionId
					analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
					return fmt.Errorf("code parsing failed")
				}
			} else if code.IsCodeParseFailed(&codeSnapshotResponse.CodeSnapshot) {
				code.PrintCodeSnapshotParserErr(&codeSnapshotResponse.CodeSnapshot)
				// Track projects push failed due to previous code parsing failure
				properties["error"] = "latest code parsing failed"
				properties["stage"] = "previous_code_parsing_failed"
				properties["code_snapshot_id"] = codeSnapshotResponse.CodeSnapshot.Cid
				properties["version_id"] = codeSnapshotResponse.VersionId
				analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
				return fmt.Errorf("latest code parsing failed, add --force to push anyway")
			}
			if !isOverwrite {
				importModelInfo, err := model.PrepareImportModelFromFilePath(ctx, modelPath, transformInput, modelType)
				if err != nil {
					return err
				}
				err = model.ImportModel(ctx, currentProject.GetCid(), codeSnapshotResponse.VersionId, importModelInfo, !noWait)
				if err != nil {
					// Track projects push failed
					properties["error"] = err.Error()
					properties["stage"] = "import_model"
					properties["code_snapshot_id"] = codeSnapshotResponse.CodeSnapshot.Cid
					properties["version_id"] = codeSnapshotResponse.VersionId
					analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
					return err
				}
			} else {
				var importModelInfo *tensorleapapi.ImportModelInfo
				if !overwriteVersionHasModel {
					importModelInfo, err = model.PrepareImportModelFromFilePath(ctx, modelPath, transformInput, modelType)
					if err != nil {
						return err
					}
				}
				err = model.OverrideModel(ctx, currentProject.GetCid(), codeSnapshotResponse.VersionId, !noWait, importModelInfo)
				if err != nil {
					// Track projects push failed
					properties["error"] = err.Error()
					properties["stage"] = "override_model"
					properties["code_snapshot_id"] = codeSnapshotResponse.CodeSnapshot.Cid
					properties["version_id"] = codeSnapshotResponse.VersionId
					analytics.SendEvent(analytics.EventCliProjectsPushFailed, properties)
					return err
				}
			}

			// Track projects push success
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
	return cmd
}
