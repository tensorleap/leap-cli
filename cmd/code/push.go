package code

import (
	"fmt"

	"github.com/spf13/cobra"
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

	cmd := &cobra.Command{
		Use:   "push",
		Short: "Push code integration",
		Long:  `Push code integration`,
		RunE: func(cmd *cobra.Command, args []string) error {
			workspaceConfig, err := workspace.GetWorkspaceConfig()
			if err != nil {
				return err
			}
			if err := auth.CheckLoggedIn(); err != nil {
				return err
			}
			close, tarGzFile, err := code.BundleCodeIntoTempFile(".", workspaceConfig)
			if err != nil {
				return err
			}
			defer close()

			ctx := cmd.Context()
			codeIntegration, _, err := code.GetAndUpdateCodeIntegrationIfNotExists(ctx, workspaceConfig)
			if err != nil {
				return err
			}

			secretId, err := secret.SyncSecretIdFromFlagAndConfig(ctx, secretId, workspaceConfig)
			if err != nil {
				return err
			}

			pythonVersion, err = code.SyncPythonVersionFromFlagAndConfig(ctx, pythonVersion, workspaceConfig)
			if err != nil {
				return err
			}

			branches := code.BranchesFromCodeIntegration(codeIntegration)
			branch, err := code.SyncBranchFromFlagAndConfig(branch, workspaceConfig, branches, codeIntegration.DefaultBranch)
			if err != nil {
				return err
			}

			_, currentVersion, err := code.PushCode(ctx, force, codeIntegration.Cid, tarGzFile, workspaceConfig.EntryFile, secretId, branch, message, pythonVersion)
			if err != nil {
				return err
			}
			supposedToWait := !noWait
			waitNeeded := supposedToWait && code.IsCodeParsing(currentVersion)

			if waitNeeded {
				ok, codeIntegrationVersion, err := code.WaitForCodeIntegrationStatus(ctx, currentVersion.Cid)
				if err != nil {
					return err
				}
				if ok {
					log.Info("Code parsed successfully")
				} else {
					code.PrintCodeIntegrationVersionParserErr(codeIntegrationVersion)
					return fmt.Errorf("code parsing failed")
				}
			} else if supposedToWait && code.IsCodeParseFailed(currentVersion) {
				code.PrintCodeIntegrationVersionParserErr(currentVersion)
				return fmt.Errorf("latest code parsing failed, add --force to push anyway")
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

	return cmd
}

func init() {
	RootCommand.AddCommand(NewPushCmd())
}
