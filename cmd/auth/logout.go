package auth

import (
	"fmt"
	"github.com/spf13/cobra"
)


func init() {
  RootCommand.AddCommand(&cobra.Command{
    Use:   "logout",
    Short: "Remove api key from the machine",
    Long:  `Remove api key from the machine`,
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Logout command")
    },
  })
}
