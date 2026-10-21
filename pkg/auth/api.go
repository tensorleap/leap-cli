package auth

import (
	"context"
	"fmt"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func GetWhoami(ctx context.Context) (*tensorleapapi.UserData, error) {
	userData, _, err := api.ApiClient.WhoAmI(ctx).Execute()
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func SetLicense(ctx context.Context, licenseToken string) error {
	result, response, err := api.ApiClient.ExtendTrial(ctx).ExtendTrialParams(*tensorleapapi.NewExtendTrialParams(licenseToken)).Execute()
	if err = api.CheckRes(response, err); err != nil {
		return err
	}
	if !result.Success {
		return fmt.Errorf("failed to set license")
	}
	log.Info("License set successfully")
	if email := registeredUserEmail(result); email != "" {
		log.Infof("Open %s and register with %s to finish setup", api.GetBaseUrlFromContext(ctx), email)
	}
	if p := DecodeLicenseToken(licenseToken); p.Activated != nil && !*p.Activated {
		log.Warn("License installed but not activated. Apply the key from your Tensorleap email with `leap auth license -t <key>` or in the UI (Manage License)")
	}
	return nil
}

func registeredUserEmail(r *tensorleapapi.ExtendTrialResponse) string {
	if u, ok := r.AdditionalProperties["registeredUser"].(map[string]interface{}); ok {
		email, _ := u["email"].(string)
		return email
	}
	return ""
}
