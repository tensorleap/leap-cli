package auth

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
  RootCommand.AddCommand(&cobra.Command{
    Use:   "login",
    Short: "Loging with Tensorleap API key",
    Long:  `Loging with Tensorleap API key`,
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Login command")
    },
  })
}
