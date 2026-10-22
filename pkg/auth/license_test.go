package auth

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func envInfo(t *testing.T, license string) *tensorleapapi.GetEnvironmentInfoResponse {
	t.Helper()
	raw := `{"clientStoragePrefixUrl":"","schemaVersion":1,"disableDatadogMetrics":false,"isCloud":false,"gatewayUrl":"","proxyUrl":""`
	if license != "" {
		raw += `,"licenseMetadata":` + license
	}
	raw += `}`
	var info tensorleapapi.GetEnvironmentInfoResponse
	if err := json.Unmarshal([]byte(raw), &info); err != nil {
		t.Fatal(err)
	}
	return &info
}

func TestLicenseLine(t *testing.T) {
	cases := []struct{ name, license, want string }{
		{"no license", "", ""},
		{"free default projects", `{"type":"free","maxUsers":1}`,
			"License: Free Tier, 1 project, evaluate up to 5,000 samples, no dataset curation or sample analysis; contact support@tensorleap.ai to unlock the full platform"},
		{"paid with limit and expiry", `{"type":"paid","expirationDate":"2027-01-31T00:00:00.000Z","maxProjects":3}`,
			"License: paid, expires 2027-01-31, 3 projects"},
		{"paid not activated", `{"type":"paid","expirationDate":"2027-01-31T00:00:00.000Z","activated":false}`,
			"License: paid, not activated, expires 2027-01-31"},
		{"free not activated", `{"type":"free","activated":false}`,
			"License: Free Tier, not activated, 1 project, evaluate up to 5,000 samples, no dataset curation or sample analysis; contact support@tensorleap.ai to unlock the full platform"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := LicenseLine(envInfo(t, c.license)); got != c.want {
				t.Fatalf("got %q\nwant %q", got, c.want)
			}
		})
	}
}

func TestDecodeLicenseToken(t *testing.T) {
	payload := base64.RawURLEncoding.EncodeToString([]byte(`{"activated":false,"user":{"email":"a@b.c"}}`))
	p := DecodeLicenseToken("h." + payload + ".s")
	if p.Activated == nil || *p.Activated || p.User.Email != "a@b.c" {
		t.Fatalf("got %+v", p)
	}
	if p := DecodeLicenseToken("not-a-jwt"); p.Activated != nil || p.User.Email != "" {
		t.Fatalf("got %+v", p)
	}
}
