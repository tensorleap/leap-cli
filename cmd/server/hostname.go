package server

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/pkg/helm"
)

func NewHostnameCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hostname",
		Short: "Print host name",
		Long:  "Print the unique host name generated for the standalone cluster during creation.",
		RunE: func(cmd *cobra.Command, args []string) error {
			hostname, err := helm.ReadHostname()
			if err != nil {
				return err
			}
			fmt.Println(hostname)
			return nil
		},
	}
	return cmd
}
