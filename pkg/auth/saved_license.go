package auth

import (
	"errors"
	"os"
	"strings"

	"github.com/spf13/viper"
	"github.com/tensorleap/leap-cli/pkg/config"
)

// the license leap mcp config saved, so leap server install applies it without the token in an assistant's chat
const savedLicenseKey = "license.token"

var ErrInvalidLicense = errors.New("this is not a Tensorleap license token; copy it again from the page you signed up on")

// LicenseEmail checks that a token can be saved and returns the email it was issued to
func LicenseEmail(token string) (string, error) {
	email := DecodeLicenseToken(strings.TrimSpace(token)).User.Email
	if email == "" {
		return "", ErrInvalidLicense
	}
	return email, nil
}

// SaveLicense keeps the token in the leap config, which becomes readable only by the user
func SaveLicense(token string) error {
	viper.Set(savedLicenseKey, strings.TrimSpace(token))
	if err := config.Save(); err != nil {
		return err
	}
	// viper writes 0644 and keeps an existing file's mode
	return os.Chmod(viper.ConfigFileUsed(), 0o600)
}

// ResolveWithSaved falls back to the saved license when neither --license-token nor --license-file is given
func (f *LicenseFlag) ResolveWithSaved() {
	if !f.HasLicense() {
		f.LicenseToken = viper.GetString(savedLicenseKey)
	}
}
