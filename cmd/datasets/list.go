package datasets

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
  RootCommand.AddCommand(&cobra.Command{
    Use:   "list",
    Short: "List available datasets",
    Long:  `List available datasets`,
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("List command")
    },
  })
}
