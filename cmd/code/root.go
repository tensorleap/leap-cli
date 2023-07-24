package code

import (
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:     "code",
	Aliases: []string{"code-integration"},
	Short:   "Manage tensorleap code integration",
	Long:    `Manage tensorleap code integration`,
}
