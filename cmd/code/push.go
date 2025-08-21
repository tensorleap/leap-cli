package code

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/secret"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func NewPushCmd() *cobra.Command {
	var secretId string
	var branch string
	var noWait bool
	var force bool
	var message string
	var pythonVersion string
	var leapMappingPath string

	cmd := &cobra.Command{
		Use:   "push",
		Short: "Push code integration",
		Long:  `Push code integration`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Define base properties for all analytics events
			properties := map[string]interface{}{
				"secret_id": secretId,
				"branch": branch,
				"no_wait": noWait,
				"force": force,
				"message": message,
				"python_version": pythonVersion,
				"leap_mapping_path": leapMappingPath,
			}
			
			// Track code push started
			if err := analytics.SendEvent(analytics.EventCliCodePushStarted, properties); err != nil {
				log.Warnf("Failed to track code push start event: %v", err)
			}
			
			workspaceConfig, err := workspace.GetWorkspaceConfig()
			if err != nil {
				// Track code push failed
				properties["error"] = err.Error()
				properties["stage"] = "get_workspace_config"
				if err := analytics.SendEvent(analytics.EventCliCodePushFailed, properties); err != nil {
					log.Warnf("Failed to track code push failure event: %v", err)
				}
				return err
			}
			if err := auth.CheckLoggedIn(); err != nil {
				// Track code push failed
				properties["error"] = err.Error()
				properties["stage"] = "auth_check"
				if err := analytics.SendEvent(analytics.EventCliCodePushFailed, properties); err != nil {
					log.Warnf("Failed to track code push failure event: %v", err)
				}
				return err
			}
			close, tarGzFile, err := code.BundleCodeIntoTempFile(".", workspaceConfig, leapMappingPath)
			if err != nil {
				// Track code push failed
				properties["error"] = err.Error()
				properties["stage"] = "bundle_code"
				if err := analytics.SendEvent(analytics.EventCliCodePushFailed, properties); err != nil {
					log.Warnf("Failed to track code push failure event: %v", err)
				}
				return err
			}
			defer close()

			ctx := cmd.Context()
			codeIntegration, _, err := code.GetAndUpdateCodeIntegrationIfNotExists(ctx, workspaceConfig)
			if err != nil {
				// Track code push failed
				properties["error"] = err.Error()
				properties["stage"] = "get_code_integration"
				if err := analytics.SendEvent(analytics.EventCliCodePushFailed, properties); err != nil {
					log.Warnf("Failed to track code push failure event: %v", err)
				}
				return err
			}

			secretId, err := secret.SyncSecretIdFromFlagAndConfig(ctx, secretId, workspaceConfig)
			if err != nil {
				// Track code push failed
				properties["error"] = err.Error()
				properties["stage"] = "sync_secret"
				if err := analytics.SendEvent(analytics.EventCliCodePushFailed, properties); err != nil {
					log.Warnf("Failed to track code push failure event: %v", err)
				}
				return err
			}

			pythonVersion, err = code.SyncPythonVersionFromFlagAndConfig(ctx, pythonVersion, workspaceConfig)
			if err != nil {
				// Track code push failed
				properties["error"] = err.Error()
				properties["stage"] = "sync_python_version"
				if err := analytics.SendEvent(analytics.EventCliCodePushFailed, properties); err != nil {
					log.Warnf("Failed to track code push failure event: %v", err)
				}
				return err
			}

			branches := code.BranchesFromCodeIntegration(codeIntegration)
			branch, err := code.SyncBranchFromFlagAndConfig(branch, workspaceConfig, branches, codeIntegration.DefaultBranch)
			if err != nil {
				// Track code push failed
				properties["error"] = err.Error()
				properties["stage"] = "sync_branch"
				if err := analytics.SendEvent(analytics.EventCliCodePushFailed, properties); err != nil {
					log.Warnf("Failed to track code push failure event: %v", err)
				}
				return err
			}

			_, currentVersion, err := code.PushCode(ctx, force, codeIntegration.Cid, tarGzFile, workspaceConfig.EntryFile, secretId, branch, message, pythonVersion)
			if err != nil {
				// Track code push failed
				properties["error"] = err.Error()
				properties["stage"] = "push_code"
				properties["code_integration_id"] = codeIntegration.Cid
				if err := analytics.SendEvent(analytics.EventCliCodePushFailed, properties); err != nil {
					log.Warnf("Failed to track code push failure event: %v", err)
				}
				return err
			}
			supposedToWait := !noWait
			waitNeeded := supposedToWait && code.IsCodeParsing(currentVersion)

			if waitNeeded {
				ok, codeIntegrationVersion, err := code.WaitForCodeIntegrationStatus(ctx, currentVersion.Cid)
				if err != nil {
					// Track code push failed
					properties["error"] = err.Error()
					properties["stage"] = "wait_for_parsing"
					properties["code_integration_id"] = codeIntegration.Cid
					properties["version_id"] = currentVersion.Cid
					if err := analytics.SendEvent(analytics.EventCliCodePushFailed, properties); err != nil {
						log.Warnf("Failed to track code push failure event: %v", err)
					}
					return err
				}
				if ok {
					log.Info("Code parsed successfully")
				} else {
					code.PrintCodeIntegrationVersionParserErr(codeIntegrationVersion)
					// Track code push failed due to parsing failure
					properties["error"] = "code parsing failed"
					properties["stage"] = "code_parsing"
					properties["code_integration_id"] = codeIntegration.Cid
					properties["version_id"] = currentVersion.Cid
					if err := analytics.SendEvent(analytics.EventCliCodePushFailed, properties); err != nil {
						log.Warnf("Failed to track code push failure event: %v", err)
					}
					return fmt.Errorf("code parsing failed")
				}
			} else if supposedToWait && code.IsCodeParseFailed(currentVersion) {
				code.PrintCodeIntegrationVersionParserErr(currentVersion)
				// Track code push failed due to previous parsing failure
				properties["error"] = "latest code parsing failed"
				properties["stage"] = "previous_parsing_failed"
				properties["code_integration_id"] = codeIntegration.Cid
				properties["version_id"] = currentVersion.Cid
				if err := analytics.SendEvent(analytics.EventCliCodePushFailed, properties); err != nil {
					log.Warnf("Failed to track code push failure event: %v", err)
				}
				return fmt.Errorf("latest code parsing failed, add --force to push anyway")
			}

			// Track code push success
			properties["code_integration_id"] = codeIntegration.Cid
			properties["version_id"] = currentVersion.Cid
			properties["final_secret_id"] = secretId
			properties["final_branch"] = branch
			properties["final_python_version"] = pythonVersion
			properties["wait_needed"] = waitNeeded
			properties["parsing_successful"] = !waitNeeded || (!code.IsCodeParsing(currentVersion) && !code.IsCodeParseFailed(currentVersion))
			if err := analytics.SendEvent(analytics.EventCliCodePushSuccess, properties); err != nil {
				log.Warnf("Failed to track code push success event: %v", err)
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&secretId, "secretId", "", "Secret id")
	cmd.Flags().StringVarP(&branch, "branch", "b", "", "Branch")
	cmd.Flags().StringVarP(&message, "message", "m", "", "Commit message")
	cmd.Flags().BoolVar(&noWait, "no-wait", false, "Do not wait for code parsing")
	cmd.Flags().BoolVarP(&force, "force", "f", false, "Force push code integration")
	cmd.Flags().StringVarP(&pythonVersion, "python-version", "p", "", "Python version")
	cmd.Flags().StringVar(&leapMappingPath, "leap-mapping", "", "Path to extrenal leap mapping file")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewPushCmd())
}
