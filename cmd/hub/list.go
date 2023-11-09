package hub

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/hub"
	"github.com/tensorleap/leap-cli/pkg/log"
)

func NewListCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List projects from the hub",
		Long:    `List projects from the hub`,
		RunE: func(cmd *cobra.Command, args []string) error {
			meta, err := hub.CreateWrapperMetaFromUrl(hub.GetPublicUrl())
			if err != nil {
				return err
			}
			projectNames := meta.GetProjectNames()
			if len(projectNames) == 0 {
				log.Warn(
					"Projects not found use 'upload [path to tar] to create a new one",
				)
				return nil
			}
			fmt.Println(strings.Join(projectNames, "\n"))
			return nil
		},
	}
}

func init() {
	RootCommand.AddCommand(NewListCmd())
}
