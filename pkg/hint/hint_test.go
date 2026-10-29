package hint

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/tensorleap/leap-cli/pkg/config"
)

func TestEachStepIsSuggestedOnceAndOnlyWithAnAssistant(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)
	path := filepath.Join(t.TempDir(), "config.yaml")
	if err := config.InitConfig(path); err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	if Show(&out, "Next", Install) || out.Len() > 0 {
		t.Fatalf("no assistant yet, so no hint:\n%s", out.String())
	}

	SetAssistant("Cursor")
	if !Show(&out, "Next", Report, "mnist") || !strings.Contains(out.String(), "Next, ask Cursor:\n    Write a Tensorleap analysis report for project mnist") {
		t.Fatalf("got:\n%s", out.String())
	}
	out.Reset()
	if Show(&out, "Next", Report, "mnist") || out.Len() > 0 {
		t.Fatalf("shown twice:\n%s", out.String())
	}
	if !Show(&out, "Next", Install) {
		t.Fatal("another step is still due")
	}

	viper.Reset()
	if err := config.InitConfig(path); err != nil {
		t.Fatal(err)
	}
	out.Reset()
	if Show(&out, "Next", Install) || Show(&out, "Next", Report, "mnist") {
		t.Fatalf("a later leap run repeated a hint:\n%s", out.String())
	}
	if !Show(&out, "Next", Integrate) || !strings.Contains(out.String(), "ask Cursor") {
		t.Fatalf("the assistant should persist too:\n%s", out.String())
	}
}
