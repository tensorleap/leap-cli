package mcp

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/tensorleap/leap-cli/pkg/config"
)

func TestSigninReadsALoginMadeAfterStart(t *testing.T) {
	t.Setenv("TL_API_URL", "")
	viper.Reset()
	t.Cleanup(viper.Reset)
	file := filepath.Join(t.TempDir(), "config.yaml")
	if err := config.InitConfig(file); err != nil {
		t.Fatal(err)
	}
	get := signin("")
	if _, err := get(); err == nil || !strings.Contains(err.Error(), "leap auth login") {
		t.Fatalf("before login: %v", err)
	}
	// what leap auth login writes, from another process
	login := "auth:\n  api_url: http://localhost:4589/api/v2\n  api_key: key-1\ncurrent_env: local\n"
	if err := os.WriteFile(file, []byte(login), 0o600); err != nil {
		t.Fatal(err)
	}
	client, err := get()
	if err != nil || client.BaseURL != "http://localhost:4589/api/v2" {
		t.Fatalf("after login: %v %+v", err, client)
	}
}
