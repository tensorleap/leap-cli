package code

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/secret"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func NewCreateCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create [code-integration-name]",
		Short: "Create a new code integration",
		Long:  `Create a new code integration`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var newCodeIntegrationName string
			var err error
			ctx := cmd.Context()
			if len(args) > 0 {
				newCodeIntegrationName = args[0]
			} else {
				codeIntegrations, err := code.GetCodeIntegrations(ctx)
				if err != nil {
					return err
				}
				newCodeIntegrationName, err = code.AskForCodeIntegrationName(codeIntegrations)
				if err != nil {
					return err
				}
			}
			codeIntegration, err := code.AddCodeIntegration(ctx, newCodeIntegrationName)
			if err != nil {
				return err
			}
			log.Infof("Creating code integration: %s", codeIntegration.GetName())

			codeIntegrationName := codeIntegration.GetName()
			_, dirExistsErr := os.Stat(codeIntegrationName)
			if dirExistsErr == nil {
				return fmt.Errorf("can't pull '%s' code integration, directory named '%s' already exists on current directory", codeIntegrationName, codeIntegrationName)
			}

			var secretId string
			selectedSecret, _, err := secret.CreateOrSelectSecretIfInUse(ctx)
			if err != nil {
				return err
			}
			if selectedSecret != nil {
				secretId = selectedSecret.GetCid()
			}

			err = workspace.CreateCodeTemplate(codeIntegration.GetCid(), "", secretId, codeIntegration.GetDefaultBranch(), codeIntegrationName)
			if err != nil {
				return err
			}

			log.Infof("'%s' code integration created into '%s' directory. Run 'cd %s' to navigate into", codeIntegrationName, codeIntegrationName, codeIntegrationName)
			return nil
		},
	}
	return cmd
}

func init() {
	RootCommand.AddCommand(NewCreateCmd())
}
