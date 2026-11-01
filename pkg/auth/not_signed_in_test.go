package auth

import (
	"testing"

	"github.com/spf13/viper"
)

func TestNotSignedInErrorNamesTheKnownServer(t *testing.T) {
	cases := []struct {
		name, apiUrl, installed, want string
	}{
		{name: "configured api url", apiUrl: "https://api.acme.tensorleap.ai/api/v2", installed: "http://localhost:4589", want: "Not signed in. Run: leap auth login https://acme.tensorleap.ai"},
		{name: "local install", installed: "http://localhost:4589", want: "Not signed in. Run: leap auth login http://localhost:4589"},
		{name: "nothing known", want: "Not signed in. Run: leap auth login <url>"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			initTestConfig(t)
			viper.Set(API_URL_CONFIG_PATH, c.apiUrl)
			orig := installedServerUrl
			t.Cleanup(func() { installedServerUrl = orig })
			installedServerUrl = func() string { return c.installed }

			if got := NotSignedInError().Error(); got != c.want {
				t.Fatalf("got %q, want %q", got, c.want)
			}
		})
	}
}
