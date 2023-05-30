package models

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(&cobra.Command{
		Use:   "import",
		Short: "Import a model into tensorleap",
		Long:  `Import a model into tensorleap`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Import model command")
		},
	})
}
