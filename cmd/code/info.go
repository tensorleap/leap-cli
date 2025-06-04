package code

import (
	"fmt"

	"github.com/tensorleap/leap-cli/pkg/auth"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func NewInfoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Show code integration information",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Info("Fetching code integration information...")
			workspaceConfig, err := workspace.GetWorkspaceConfig()
			if err != nil {
				return err
			}
			log.Infof("Using workspace: %s", workspaceConfig)
			log.Infof("Code Integration ID: %s", workspaceConfig.CodeIntegrationId)
			if err := auth.CheckLoggedIn(); err != nil {
				return err
			}
			if workspaceConfig.CodeIntegrationId == "" {
				log.Info("No code integration configured in workspace")
				return nil
			}

			ctx := cmd.Context()
			log.Infof("Fetching code integration with ID: %s", workspaceConfig.CodeIntegrationId)
			codeIntegration, err := code.GetCodeIntegrationById(ctx, workspaceConfig.CodeIntegrationId)
			if err != nil {
				return fmt.Errorf("failed to get code integration: %w", err)
			}

			code.CodeIntegrationInfo(codeIntegration)

			return nil
		},
	}
}

func init() {
	RootCommand.AddCommand(NewInfoCmd())
}
