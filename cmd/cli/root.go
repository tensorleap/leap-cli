package cli

import (
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "cli",
	Short: "Manage leap cli",
	Long:  `Manage leap cli`,
}
