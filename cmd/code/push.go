package code

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func NewPushCmd() *cobra.Command {
	var secretId string
	var watch bool
	var force bool

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
			codeIntegration, err := code.GetAndUpdateCodeIntegrationIfNotExists(ctx, workspaceConfig)
			if err != nil {
				return err
			}
			if len(secretId) == 0 {
				secretId = workspaceConfig.SecretManagerId
			}

			_, currentVersion, err := code.PushCode(ctx, force, codeIntegration.Cid, tarGzFile, workspaceConfig.EntryFile, secretId)
			if err != nil {
				return err
			}
			if watch && code.IsCodeParsing(currentVersion) {
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
			} else if watch && code.IsCodeParseFailed(currentVersion) {
				code.PrintCodeIntegrationVersionParserErr(currentVersion)
				return fmt.Errorf("latest code parsing failed, add --force to push anyway")
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&secretId, "secretId", "", "Secret manager id")
	cmd.Flags().BoolVarP(&watch, "watch", "w", false, "Watch code integration status")
	cmd.Flags().BoolVarP(&watch, "force", "f", false, "Force push code integration")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewPushCmd())
}
