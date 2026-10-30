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
	if msg, fallback := notActivatedMessage(DecodeLicenseToken(licenseToken)); msg != "" {
		log.Warn(msg)
		if fallback != "" {
			log.Info(fallback)
		}
	}
	return nil
}

// notActivatedMessage tells how to activate a colored (not yet activated) license; "" for any other.
// A Free Tier signup activates with one press in its email, so the key is only the fallback there.
func notActivatedMessage(p LicenseTokenPayload) (msg, fallback string) {
	if p.Activated == nil || *p.Activated {
		return "", ""
	}
	if p.LicenseType != freeLicenseType {
		return "License installed but not activated. Apply the key from your Tensorleap email with `leap auth license -t <key>` or in the UI (Manage License)", ""
	}
	email := "the email we sent you"
	if p.User.Email != "" {
		email = "the email we sent to " + p.User.Email
	}
	return fmt.Sprintf("Press Activate my license in %s. Tensorleap unlocks by itself.", email),
		"If the button doesn't work, apply the key from that email with `leap auth license -t <key>` or in the UI (Manage License)."
}

func registeredUserEmail(r *tensorleapapi.ExtendTrialResponse) string {
	if u, ok := r.AdditionalProperties["registeredUser"].(map[string]interface{}); ok {
		email, _ := u["email"].(string)
		return email
	}
	return ""
}
