package hub

import "github.com/spf13/cobra"

var RootCommand = &cobra.Command{
	Use:     "hub",
	Aliases: []string{"hub"},
	Short:   "Manage tensorleap hub",
	Long:    `Manage tensorleap hub`,
}
