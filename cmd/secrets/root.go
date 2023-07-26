package secrets

import (
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "secrets",
	Short: "Manage secrets",
	Long:  `Manage secrets`,
}
