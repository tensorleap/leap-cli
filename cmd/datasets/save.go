package datasets

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(&cobra.Command{
		Use:   "save",
		Short: "Save dataset script",
		Long:  `Save dataset script`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Save command")
		},
	})
}
