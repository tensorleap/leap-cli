package mcp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
	"github.com/tensorleap/helm-charts/pkg/local"
)

// The install skill keeps its record of this machine's setup in <data dir>/install-notes.md, outside the
// workspace, so Cursor asks before every write to it, even with auto-review. One allow rule for exactly
// that file in Cursor's CLI config removes the question; nothing else is allowed.
const (
	installNotesFile = "install-notes.md"
	// the key leap server install saves its data dir under in the CLI config (cmd/server DATA_DIR_CONFIG_PATH)
	dataDirConfigKey = "data_dir"
)

func cursorCLIConfigPath() string {
	return home(".cursor", "cli-config.json")
}

// installDataDir is where leap server install puts (or will put) its data: $TL_DATA_DIR, else the data dir
// the last install saved, else the default location
func installDataDir() string {
	if d := os.Getenv(local.DATA_DIR_ENV_NAME); d != "" {
		return d
	}
	if d := viper.GetString(dataDirConfigKey); d != "" {
		return d
	}
	return local.DEFAULT_DATA_DIR
}

func cursorInstallNotesRule() string {
	return fmt.Sprintf("Write(%s)", filepath.Join(installDataDir(), installNotesFile))
}

type cursorPermissions struct {
	Permissions *struct {
		Allow *[]string `json:"allow"`
	} `json:"permissions"`
}

// cursorRuleInPlace says whether Cursor's CLI config already allows rule
func cursorRuleInPlace(rule string) bool {
	b, err := os.ReadFile(cursorCLIConfigPath())
	if err != nil {
		return false
	}
	var c cursorPermissions
	if json.Unmarshal(b, &c) != nil || c.Permissions == nil || c.Permissions.Allow == nil {
		return false
	}
	for _, r := range *c.Permissions.Allow {
		if r == rule {
			return true
		}
	}
	return false
}

// addCursorAllowRule adds rule to permissions.allow in Cursor's CLI config, editing only that list's text,
// so every other setting keeps its value, formatting and order. A missing file is created with the fields
// Cursor requires; a file that isn't plain JSON, or whose permissions aren't the expected shape, is left alone.
func addCursorAllowRule(rule string) error {
	path := cursorCLIConfigPath()
	old, err := os.ReadFile(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if len(bytes.TrimSpace(old)) == 0 {
		fresh := map[string]any{
			"version":     1,
			"editor":      map[string]any{"vimMode": false},
			"permissions": map[string]any{"allow": []string{rule}, "deny": []string{}},
		}
		out, _ := json.MarshalIndent(fresh, "", "  ")
		_, err := writeWithBackup(path, old, append(out, '\n'))
		return err
	}
	if !json.Valid(old) {
		return fmt.Errorf("%s is not plain JSON, so it was left as is", tilde(path))
	}
	var c cursorPermissions
	if err := json.Unmarshal(old, &c); err != nil {
		return fmt.Errorf("%s: permissions.allow is not a list of rules, so it was left as is", tilde(path))
	}
	allow := []string{}
	if c.Permissions != nil && c.Permissions.Allow != nil {
		allow = *c.Permissions.Allow
	}
	for _, r := range allow {
		if r == rule {
			return nil
		}
	}
	out, err := patchJSON(old, "permissions", "allow", append(allow, rule))
	if err == nil && c.Permissions == nil {
		// Cursor requires both lists once permissions exists
		out, err = patchJSON(out, "permissions", "deny", []string{})
	}
	if err != nil {
		return fmt.Errorf("%s: %v, so it was left as is", tilde(path), err)
	}
	_, err = writeWithBackup(path, old, out)
	return err
}

// allowInstallNotesInCursor adds the rule when allow is set (--cursor-allow-install-notes), asks first in a
// terminal, and otherwise only says how to turn it on. Declining keeps Cursor's question on each write.
func allowInstallNotesInCursor(out ui, allow bool) {
	if runtime.GOOS == "windows" {
		return // leap server install doesn't run on Windows, so there are no install notes to write
	}
	rule := cursorInstallNotesRule()
	path := tilde(cursorCLIConfigPath())
	if cursorRuleInPlace(rule) {
		out.ok("Cursor writes the install notes without asking (%s in %s)", rule, path)
		return
	}
	if !allow && interactive() {
		i, err := choose(out, "Let Cursor write Tensorleap's install notes without asking? (adds "+rule+" to "+path+")", "",
			[]choice{
				{label: "Yes, allow writing that one file", detail: "the install skill records this machine's setup there, outside your workspace"},
				{label: "No, ask me each time"},
			}, 0)
		allow = err == nil && i == 0
	}
	if !allow {
		out.skip("Cursor asks before each write to the install notes; to stop that, run: leap mcp config cursor --cursor-allow-install-notes")
		return
	}
	if err := addCursorAllowRule(rule); err != nil {
		out.warn("Could not update %s: %v", path, err)
		return
	}
	out.ok("Cursor writes the install notes without asking (%s in %s)", rule, path)
}
