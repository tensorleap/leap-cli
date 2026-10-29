package auth

import (
	"encoding/base64"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/tensorleap/leap-cli/pkg/config"
)

func testToken(payload string) string {
	enc := base64.RawURLEncoding.EncodeToString
	return enc([]byte(`{"alg":"RS256"}`)) + "." + enc([]byte(payload)) + ".sig"
}

func initTestConfig(t *testing.T) string {
	viper.Reset()
	t.Cleanup(viper.Reset)
	path := filepath.Join(t.TempDir(), "config.yaml")
	if err := config.InitConfig(path); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestLicenseEmailRejectsTokensWithoutAnEmail(t *testing.T) {
	for name, token := range map[string]string{
		"garbage":  "not-a-token",
		"no email": testToken(`{"activated":false,"user":{}}`),
		"not json": testToken(`nope`),
	} {
		if _, err := LicenseEmail(token); err != ErrInvalidLicense {
			t.Errorf("%s: got %v", name, err)
		}
	}
	if email, err := LicenseEmail(" " + testToken(`{"activated":false,"user":{"email":"a@b.c"}}`) + "\n"); err != nil || email != "a@b.c" {
		t.Fatalf("got %q %v", email, err)
	}
}

func TestSavedLicenseIsPrivateAndUsedOnlyWithoutAFlag(t *testing.T) {
	path := initTestConfig(t)
	saved := testToken(`{"activated":false,"user":{"email":"a@b.c"}}`)
	if err := SaveLicense(saved); err != nil {
		t.Fatal(err)
	}
	if st, _ := os.Stat(path); st.Mode().Perm() != 0o600 {
		t.Fatalf("config mode %v, want 0600", st.Mode().Perm())
	}

	// a later leap run reads it back from the file
	viper.Reset()
	if err := config.InitConfig(path); err != nil {
		t.Fatal(err)
	}
	f := NewLicenseFlag()
	f.ResolveWithSaved()
	if got, err := f.PrepareLicenseToken(); err != nil || got != saved {
		t.Fatalf("saved token not resolved: %q %v", got, err)
	}

	explicit := &LicenseFlag{LicenseToken: "explicit"}
	explicit.ResolveWithSaved()
	fromFile := &LicenseFlag{FilePath: "/some/license"}
	fromFile.ResolveWithSaved()
	if explicit.LicenseToken != "explicit" || fromFile.LicenseToken != "" {
		t.Fatalf("a flag must win over the saved license: %+v %+v", explicit, fromFile)
	}

	viper.Reset()
	none := NewLicenseFlag()
	none.ResolveWithSaved()
	if none.HasLicense() {
		t.Fatal("nothing saved, so no license")
	}
}
