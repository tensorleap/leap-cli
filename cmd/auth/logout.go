package auth

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/cli-go/pkg/auth"
)

func init() {
	RootCommand.AddCommand(&cobra.Command{
		Use:   "logout",
		Short: "Remove api key from the machine",
		Long:  `Remove api key from the machine`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return auth.Logout()
		},
	})
}
