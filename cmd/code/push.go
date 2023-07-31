package code

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func NewPushCmd() *cobra.Command {
	var secretId string
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

			_, err = code.AddCodeIntegrationVersion(ctx, tarGzFile, codeIntegration, workspaceConfig.EntryFile, secretId)
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&secretId, "secretManagerId", "", "Secret manager id")
	return cmd
}

func init() {
	RootCommand.AddCommand(NewPushCmd())
}
