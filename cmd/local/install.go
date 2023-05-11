package local

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
  RootCommand.AddCommand(&cobra.Command{
    Use:   "install",
    Short: "Installs tensorleap on the local machine, running in a docker container",
    Long:  `Installs tensorleap on the local machine, running in a docker container`,
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Install command")
    },
  })
}
