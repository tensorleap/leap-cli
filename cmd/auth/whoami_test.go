package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/config"
)

func TestWhoamiWhenSignedOutFailsWithoutPrompting(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)
	if err := config.InitConfig(filepath.Join(t.TempDir(), "config.yaml")); err != nil {
		t.Fatal(err)
	}
	cmd := NewWhoamiCmd()
	var out strings.Builder
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetIn(strings.NewReader("")) // a login prompt would hit EOF here, as it does in an AI assistant's shell
	cmd.SetArgs(nil)
	err := cmd.Execute()
	if err == nil || !strings.HasPrefix(err.Error(), "Not signed in. Run: leap auth login ") {
		t.Fatalf("got %v", err)
	}
}

// a saved sign-in whose server can't be reached (e.g. from Codex's sandbox) is not reported as signed out
func TestWhoamiWhenTheServerIsUnreachableSaysSo(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)
	if err := config.InitConfig(filepath.Join(t.TempDir(), "config.yaml")); err != nil {
		t.Fatal(err)
	}
	down := httptest.NewServer(http.NotFoundHandler())
	down.Close()
	viper.Set(auth.API_URL_CONFIG_PATH, down.URL+"/api/v2")
	viper.Set(auth.API_KEY_CONFIG_PATH, "k")

	cmd := NewWhoamiCmd()
	cmd.SetOut(&strings.Builder{})
	cmd.SetArgs(nil)
	err := cmd.ExecuteContext(context.Background())
	if err == nil || !strings.HasPrefix(err.Error(), "Can't reach the Tensorleap server at "+down.URL+": ") {
		t.Fatalf("got %v", err)
	}
}
