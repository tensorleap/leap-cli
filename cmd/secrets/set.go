package secrets

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/secret"
)

func NewConfigSecret() *cobra.Command {
	var secretId string
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set secret on local leap config",
		Long:  "Set secret on local leap config",
		RunE: func(cmd *cobra.Command, args []string) error {
			return secret.SetSecret(cmd.Context(), secretId)
		},
	}

	cmd.Flags().StringVar(&secretId, "secret-id", "", "Secret ID")
	return cmd
}

func init() {
	RootCommand.AddCommand(NewConfigSecret())
}
