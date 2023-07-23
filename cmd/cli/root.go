package cli

import (
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "cli",
	Short: "Manage tensorleap cli",
	Long:  `Manage tensorleap cli`,
}
