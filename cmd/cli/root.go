package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/version"
)

var showVersionInfo bool

var RootCommand = &cobra.Command{
	Use:   "cli",
	Short: "Manage leap cli",
	Long:  `Manage leap cli`,
	Run: func(cmd *cobra.Command, args []string) {
		if !showVersionInfo {
			_ = cmd.Help()
			return
		}

		fmt.Println(version.CliVersion)
	},
}

func init() {
	RootCommand.Flags().BoolVar(&showVersionInfo, "version", false, "Show version information")
}
