package local

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
  RootCommand.AddCommand(&cobra.Command{
    Use:   "upgrade",
    Short: "Upgrade an existing local tensorleap installation to the latest version",
    Long:  `Upgrade an existing local tensorleap installation to the latest version`,
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Upgrade command")
    },
  })
}
