package auth

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/log"
)

func NewLogoutCmd() *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Remove api key from the machine",
		Long:  `Remove api key from the machine`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Track logout started
			properties := map[string]interface{}{
				"environment_name": name,
				"has_environment_name": len(name) > 0,
			}
			
			err := auth.Logout(name)
			if err != nil {
				// Track logout failed
				properties["error"] = err.Error()
				properties["stage"] = "logout_execution"
				if err := analytics.SendEvent(analytics.EventAuthLogoutFailed, properties); err != nil {
					log.Warnf("Failed to track logout failure event: %v", err)
				}
				return err
			}
			
			// Track logout success
			properties["stage"] = "logout_execution"
			if err := analytics.SendEvent(analytics.EventAuthLogoutSuccess, properties); err != nil {
				log.Warnf("Failed to track logout success event: %v", err)
			}
			
			return nil
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Name of the environment to logout from, by default logout from the current environment")
	return cmd
}

func init() {
	RootCommand.AddCommand(NewLogoutCmd())
}
