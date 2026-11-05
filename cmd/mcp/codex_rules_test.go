package mcp

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/config"
)

// runCodexConfig runs leap mcp config codex in a fresh HOME (no terminal, as under go test) and returns its output and event
func runCodexConfig(t *testing.T, dir string, args ...string) (string, map[string]interface{}) {
	t.Helper()
	t.Setenv("HOME", dir)
	viper.Reset()
	t.Cleanup(viper.Reset)
	if err := config.InitConfig(filepath.Join(dir, "config.yaml")); err != nil {
		t.Fatal(err)
	}
	var props map[string]interface{}
	t.Cleanup(func() { sendEvent = func(analytics.EventType, map[string]interface{}) {} })
	sendEvent = func(_ analytics.EventType, p map[string]interface{}) { props = p }
	var out bytes.Buffer
	cmd := newConfigCmd()
	cmd.SetOut(&out)
	cmd.SetArgs(append([]string{"codex", "--no-skills"}, args...))
	if err := cmd.Execute(); err != nil {
		t.Fatalf("%v: %v\n%s", args, err, out.String())
	}
	return out.String(), props
}

func TestCodexRuleIsNotWrittenWithoutATerminalOrTheFlag(t *testing.T) {
	dir := t.TempDir()
	out, props := runCodexConfig(t, dir)
	if _, err := os.Stat(filepath.Join(dir, ".codex", "rules", "tensorleap.rules")); !os.IsNotExist(err) {
		t.Fatalf("rules file written without asking: %v", err)
	}
	if !strings.Contains(out, "leap mcp config codex --codex-allow-leap") {
		t.Errorf("no line on how to enable it:\n%s", out)
	}
	if props["codex_allow_leap"] != false {
		t.Errorf("props %v", props)
	}
}

func TestCodexAllowLeapWritesOnlyItsOwnRulesFile(t *testing.T) {
	dir := t.TempDir()
	rulesDir := filepath.Join(dir, ".codex", "rules")
	if err := os.MkdirAll(rulesDir, 0o755); err != nil {
		t.Fatal(err)
	}
	mine := `prefix_rule(pattern=["git"], decision="allow")` + "\n"
	if err := os.WriteFile(filepath.Join(rulesDir, "default.rules"), []byte(mine), 0o644); err != nil {
		t.Fatal(err)
	}
	// an older or edited tensorleap.rules is replaced
	if err := os.WriteFile(filepath.Join(rulesDir, "tensorleap.rules"), []byte("old"), 0o644); err != nil {
		t.Fatal(err)
	}

	out, props := runCodexConfig(t, dir, "--codex-allow-leap")
	got, err := os.ReadFile(filepath.Join(rulesDir, "tensorleap.rules"))
	if err != nil || string(got) != codexRules || !strings.Contains(string(got), `prefix_rule(pattern=["leap"], decision="allow")`) {
		t.Fatalf("tensorleap.rules: %q %v", got, err)
	}
	if other, _ := os.ReadFile(filepath.Join(rulesDir, "default.rules")); string(other) != mine {
		t.Fatalf("the user's own rules file changed: %q", other)
	}
	if !strings.Contains(out, "Codex runs leap commands without asking") || props["codex_allow_leap"] != true {
		t.Errorf("props %v\n%s", props, out)
	}

	// a second run, even without the flag, sees it in place and changes nothing
	out, props = runCodexConfig(t, dir)
	if !strings.Contains(out, "Codex already runs leap commands without asking") || props["codex_allow_leap"] != true {
		t.Errorf("second run: props %v\n%s", props, out)
	}
}

func TestCodexPrintShowsTheRule(t *testing.T) {
	var buf bytes.Buffer
	printConfig(&buf, "codex", "/usr/local/bin/leap", []string{"mcp"})
	if !strings.Contains(buf.String(), "tensorleap.rules") || !strings.Contains(buf.String(), `prefix_rule(pattern=["leap"], decision="allow")`) {
		t.Fatalf("codex --print:\n%s", buf.String())
	}
}
