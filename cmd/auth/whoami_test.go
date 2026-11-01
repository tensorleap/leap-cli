package auth

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"
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
