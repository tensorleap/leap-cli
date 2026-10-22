package auth

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

type LicenseTokenPayload struct {
	Activated *bool `json:"activated"`
	User      struct {
		Email string `json:"email"`
	} `json:"user"`
}

// Unverified: only feeds hints and analytics, the server validates the signature.
func DecodeLicenseToken(token string) LicenseTokenPayload {
	var p LicenseTokenPayload
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return p
	}
	raw, err := base64.RawURLEncoding.DecodeString(strings.TrimRight(parts[1], "="))
	if err != nil {
		return p
	}
	_ = json.Unmarshal(raw, &p)
	return p
}
