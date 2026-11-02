package mcp

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/tensorleap/helm-charts/pkg/local"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/config"
)

// runCursorConfig runs leap mcp config cursor in a fresh HOME (no terminal, as under go test)
func runCursorConfig(t *testing.T, dir, dataDir string, args ...string) (string, map[string]interface{}) {
	t.Helper()
	t.Setenv("HOME", dir)
	t.Setenv(local.DATA_DIR_ENV_NAME, "")
	viper.Reset()
	t.Cleanup(viper.Reset)
	if err := config.InitConfig(filepath.Join(dir, "config.yaml")); err != nil {
		t.Fatal(err)
	}
	if dataDir != "" {
		viper.Set(dataDirConfigKey, dataDir)
	}
	var props map[string]interface{}
	t.Cleanup(func() { sendEvent = func(analytics.EventType, map[string]interface{}) {} })
	sendEvent = func(_ analytics.EventType, p map[string]interface{}) { props = p }
	var out bytes.Buffer
	cmd := newConfigCmd()
	cmd.SetOut(&out)
	cmd.SetArgs(append([]string{"cursor", "--no-skills"}, args...))
	if err := cmd.Execute(); err != nil {
		t.Fatalf("%v: %v\n%s", args, err, out.String())
	}
	return out.String(), props
}

func readCursorConfig(t *testing.T, dir string) map[string]any {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(dir, ".cursor", "cli-config.json"))
	if err != nil {
		t.Fatal(err)
	}
	var c map[string]any
	if err := json.Unmarshal(b, &c); err != nil {
		t.Fatalf("%v\n%s", err, b)
	}
	return c
}

func TestCursorRuleIsNotWrittenWithoutATerminalOrTheFlag(t *testing.T) {
	dir := t.TempDir()
	out, props := runCursorConfig(t, dir, "")
	if _, err := os.Stat(filepath.Join(dir, ".cursor", "cli-config.json")); !os.IsNotExist(err) {
		t.Fatalf("cli-config.json written without asking: %v", err)
	}
	if !strings.Contains(out, "leap mcp config cursor --cursor-allow-install-notes") {
		t.Errorf("no line on how to enable it:\n%s", out)
	}
	if props["cursor_allow_install_notes"] != false {
		t.Errorf("props %v", props)
	}
}

func TestCursorAllowRuleCreatesAValidConfigForTheDefaultDataDir(t *testing.T) {
	dir := t.TempDir()
	_, props := runCursorConfig(t, dir, "", "--cursor-allow-install-notes")
	c := readCursorConfig(t, dir)
	perms := c["permissions"].(map[string]any)
	allow := perms["allow"].([]any)
	if len(allow) != 1 || allow[0] != "Write(/var/lib/tensorleap/standalone/install-notes.md)" {
		t.Fatalf("allow = %v", allow)
	}
	// the fields Cursor requires in a new file
	if c["version"] != float64(1) || perms["deny"] == nil || c["editor"].(map[string]any)["vimMode"] != false {
		t.Fatalf("config %v", c)
	}
	if props["cursor_allow_install_notes"] != true {
		t.Errorf("props %v", props)
	}
}

func TestCursorAllowRuleMergesIntoTheUsersConfig(t *testing.T) {
	dir := t.TempDir()
	mine := `{
  "version": 1,
  "editor": { "vimMode": true },
  "model": "gpt-5",
  "permissions": {
    "allow": ["Shell(ls)", "Read(**)"],
    "deny": ["Shell(rm)"]
  }
}
`
	path := filepath.Join(dir, ".cursor", "cli-config.json")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(mine), 0o644); err != nil {
		t.Fatal(err)
	}
	// the install's data dir, as leap server install saved it
	runCursorConfig(t, dir, "/data/tl", "--cursor-allow-install-notes")
	runCursorConfig(t, dir, "/data/tl", "--cursor-allow-install-notes") // a second run adds nothing

	got, _ := os.ReadFile(path)
	want := strings.Replace(mine, `"allow": ["Shell(ls)", "Read(**)"]`,
		"\"allow\": [\n      \"Shell(ls)\",\n      \"Read(**)\",\n      \"Write(/data/tl/install-notes.md)\"\n    ]", 1)
	if string(got) != want {
		t.Fatalf("only the allow list changes:\n%s\nwant:\n%s", got, want)
	}
	if backup, _ := os.ReadFile(path + ".bak"); string(backup) != mine {
		t.Errorf("backup %q", backup)
	}
}

func TestCursorAllowRuleAddsPermissionsWhenMissing(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, ".cursor", "cli-config.json")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(`{"version": 1, "editor": {"vimMode": false}}`), 0o644); err != nil {
		t.Fatal(err)
	}
	runCursorConfig(t, dir, "", "--cursor-allow-install-notes")
	perms := readCursorConfig(t, dir)["permissions"].(map[string]any)
	if len(perms["allow"].([]any)) != 1 || perms["deny"] == nil {
		t.Fatalf("permissions %v", perms)
	}
}

func TestCursorAllowRuleLeavesAFileItCannotReadAlone(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, ".cursor", "cli-config.json")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	for _, mine := range []string{"{ // comment\n}", `{"permissions": {"allow": "Shell(ls)"}}`} {
		if err := os.WriteFile(path, []byte(mine), 0o644); err != nil {
			t.Fatal(err)
		}
		out, _ := runCursorConfig(t, dir, "", "--cursor-allow-install-notes")
		if got, _ := os.ReadFile(path); string(got) != mine || !strings.Contains(out, "Could not update") {
			t.Fatalf("%q was changed to %q\n%s", mine, got, out)
		}
	}
}

// leap server install prefers $TL_DATA_DIR over the data dir it saved, so the rule does too
func TestCursorRuleNamesTheDataDirInstallUses(t *testing.T) {
	t.Setenv(local.DATA_DIR_ENV_NAME, "/from/env")
	viper.Reset()
	t.Cleanup(viper.Reset)
	viper.Set(dataDirConfigKey, "/saved")
	if got := cursorInstallNotesRule(); got != "Write(/from/env/install-notes.md)" {
		t.Fatalf("with TL_DATA_DIR: %s", got)
	}
	t.Setenv(local.DATA_DIR_ENV_NAME, "")
	if got := cursorInstallNotesRule(); got != "Write(/saved/install-notes.md)" {
		t.Fatalf("with a saved data_dir: %s", got)
	}
}
