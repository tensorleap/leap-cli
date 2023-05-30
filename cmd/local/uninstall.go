package local

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(&cobra.Command{
		Use:   "uninstall",
		Short: "Remove local Tensorleap installation",
		Long:  `Remove local Tensorleap installation`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Uninstall command")
		},
	})
}
