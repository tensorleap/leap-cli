package auth

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
)

func NewLicenseCmd() *cobra.Command {
	licenseFlag := auth.NewLicenseFlag()

	cmd := &cobra.Command{
		Use:   "license",
		Short: "Set license",
		Long:  `Set license`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var licenseToken, err = licenseFlag.PrepareLicenseToken()
			if err != nil {
				return err
			}

			ctx, err := auth.InitMaybeUnauthedContext(cmd.Context(), "")
			if err != nil {
				return err
			}

			return auth.SetLicense(ctx, licenseToken)
		},
	}

	licenseFlag.AddShortFlags(cmd)

	return cmd
}

func init() {
	RootCommand.AddCommand(NewLicenseCmd())
}
