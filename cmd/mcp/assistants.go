package mcp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type assistant struct {
	id, title string
	// how the user reloads the assistant so it starts leap mcp
	restart string
	// install.sh --tool value for the Tensorleap skills; "" when the assistant has no skills support
	skillsTool string
	// where the skills land, for the summary
	skillsDir string
	detect    func() bool
}

var assistants = []assistant{
	{id: "claude-code", title: "Claude Code", restart: "Start a new Claude Code session",
		detect: func() bool { return onPath("claude") }},
	{id: "cursor", title: "Cursor", restart: "Restart Cursor (or run Developer: Reload Window)", skillsTool: "cursor", skillsDir: "~/.cursor/skills",
		detect: func() bool { return onPath("cursor") || exists(home(".cursor")) }},
	{id: "vscode", title: "VS Code (GitHub Copilot)", restart: "Reload VS Code (Developer: Reload Window)", skillsTool: "copilot", skillsDir: "~/.copilot/skills",
		detect: func() bool { return onPath("code") || exists(filepath.Dir(vscodeConfigPath())) }},
	{id: "codex", title: "OpenAI Codex", restart: "Restart Codex",
		detect: func() bool { return onPath("codex") || exists(home(".codex")) }},
	{id: "windsurf", title: "Windsurf", restart: "Restart Windsurf",
		detect: func() bool { return onPath("windsurf") || exists(home(".codeium", "windsurf")) }},
	{id: "claude-desktop", title: "Claude Desktop", restart: "Quit and reopen Claude Desktop",
		detect: func() bool { return exists(filepath.Dir(desktopConfigPath())) }},
}

var clientNames = func() []string {
	ids := make([]string, len(assistants))
	for i, a := range assistants {
		ids[i] = a.id
	}
	return ids
}()

func findAssistant(id string) *assistant {
	for i := range assistants {
		if assistants[i].id == id {
			return &assistants[i]
		}
	}
	return nil
}

func onPath(bin string) bool {
	_, err := exec.LookPath(bin)
	return err == nil
}

func exists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

func home(parts ...string) string {
	h, _ := os.UserHomeDir()
	return filepath.Join(append([]string{h}, parts...)...)
}

func appData(parts ...string) string {
	if d := os.Getenv("APPDATA"); d != "" {
		return filepath.Join(append([]string{d}, parts...)...)
	}
	return home(append([]string{"AppData", "Roaming"}, parts...)...)
}

func desktopConfigPath() string {
	switch runtime.GOOS {
	case "windows":
		return appData("Claude", "claude_desktop_config.json")
	case "darwin":
		return home("Library", "Application Support", "Claude", "claude_desktop_config.json")
	}
	return home(".config", "Claude", "claude_desktop_config.json")
}

// VS Code's user-level MCP file applies to every workspace
func vscodeConfigPath() string {
	switch runtime.GOOS {
	case "windows":
		return appData("Code", "User", "mcp.json")
	case "darwin":
		return home("Library", "Application Support", "Code", "User", "mcp.json")
	}
	return home(".config", "Code", "User", "mcp.json")
}

// tilde shortens a path under the home directory for display
func tilde(p string) string {
	h, err := os.UserHomeDir()
	if err == nil && strings.HasPrefix(p, h+string(filepath.Separator)) {
		return "~" + p[len(h):]
	}
	return p
}

// configFile is where an assistant keeps its MCP servers and the key that holds them
func configFile(id string) (path, key string) {
	switch id {
	case "cursor":
		return home(".cursor", "mcp.json"), "mcpServers"
	case "windsurf":
		return home(".codeium", "windsurf", "mcp_config.json"), "mcpServers"
	case "claude-desktop":
		return desktopConfigPath(), "mcpServers"
	case "vscode":
		return vscodeConfigPath(), "servers"
	case "codex":
		return home(".codex", "config.toml"), "mcp_servers"
	}
	return "", ""
}

var errUnchanged = errors.New("already set up")

// upsertJSON sets <key>.tensorleap in a JSON config, keeping every other server and setting as written.
// A file it cannot parse (VS Code allows comments) is left alone.
func upsertJSON(path, key string, entry serverEntry) (backup string, err error) {
	doc := map[string]json.RawMessage{}
	old, err := os.ReadFile(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return "", err
	}
	if len(bytes.TrimSpace(old)) > 0 {
		if err := json.Unmarshal(old, &doc); err != nil {
			return "", fmt.Errorf("%s is not plain JSON (%v), so it was left as is", tilde(path), err)
		}
	}
	servers := map[string]json.RawMessage{}
	if raw, ok := doc[key]; ok {
		if err := json.Unmarshal(raw, &servers); err != nil {
			return "", fmt.Errorf("%s: %q is not an object, so it was left as is", tilde(path), key)
		}
	}
	want, _ := json.Marshal(entry)
	if have, ok := servers["tensorleap"]; ok {
		var a, b any
		_ = json.Unmarshal(have, &a)
		_ = json.Unmarshal(want, &b)
		if fmt.Sprint(a) == fmt.Sprint(b) {
			return "", errUnchanged
		}
	}
	servers["tensorleap"] = want
	doc[key], _ = json.Marshal(servers)
	out, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return "", err
	}
	return writeWithBackup(path, old, append(out, '\n'))
}

// upsertCodexTOML replaces or appends the [mcp_servers.tensorleap] table in Codex's config.toml
func upsertCodexTOML(path string, entry serverEntry) (backup string, err error) {
	old, err := os.ReadFile(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return "", err
	}
	quoted := make([]string, len(entry.Args))
	for i, a := range entry.Args {
		quoted[i] = fmt.Sprintf("%q", a)
	}
	table := fmt.Sprintf("[mcp_servers.tensorleap]\ncommand = %q\nargs = [%s]\n", entry.Command, strings.Join(quoted, ", "))
	lines := strings.SplitAfter(string(old), "\n")
	start, end := -1, len(lines)
	for i, l := range lines {
		t := strings.TrimSpace(l)
		if start < 0 && t == "[mcp_servers.tensorleap]" {
			start = i
		} else if start >= 0 && strings.HasPrefix(t, "[") {
			end = i
			break
		}
	}
	var out string
	if start < 0 {
		out = strings.TrimRight(string(old), "\n")
		if out != "" {
			out += "\n\n"
		}
		out += table
	} else {
		if strings.TrimSpace(strings.Join(lines[start:end], "")) == strings.TrimSpace(table) {
			return "", errUnchanged
		}
		rest := strings.Join(lines[end:], "")
		if rest != "" {
			rest = "\n" + rest
		}
		out = strings.Join(lines[:start], "") + table + rest
	}
	return writeWithBackup(path, old, []byte(out))
}

func writeWithBackup(path string, old, data []byte) (backup string, err error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return "", err
	}
	if len(old) > 0 {
		backup = path + ".bak"
		if err := os.WriteFile(backup, old, 0o600); err != nil {
			return "", err
		}
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o600); err != nil {
		return "", err
	}
	return backup, os.Rename(tmp, path)
}

// closest suggests the option a typo was probably meant to be (edit distance up to 2)
func closest(s string, options []string) string {
	best, bestD := "", 3
	for _, o := range options {
		if d := editDistance(strings.ToLower(s), strings.ToLower(o)); d < bestD {
			best, bestD = o, d
		}
	}
	return best
}

func editDistance(a, b string) int {
	prev := make([]int, len(b)+1)
	for j := range prev {
		prev[j] = j
	}
	for i := 1; i <= len(a); i++ {
		cur := make([]int, len(b)+1)
		cur[0] = i
		for j := 1; j <= len(b); j++ {
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}
			cur[j] = min(prev[j]+1, cur[j-1]+1, prev[j-1]+cost)
		}
		prev = cur
	}
	return prev[len(b)]
}
