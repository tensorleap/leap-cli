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
	{id: "codex", title: "OpenAI Codex", restart: "Restart Codex", skillsTool: "codex", skillsDir: "~/.agents/skills",
		detect: func() bool { return onPath("codex") || exists(home(".codex")) }},
	{id: "windsurf", title: "Windsurf", restart: "Restart Windsurf", skillsTool: "windsurf", skillsDir: "~/.codeium/windsurf/skills",
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

// upsertJSON sets <key>.tensorleap in a JSON config by editing only that entry's text, so every
// other server and setting keeps its formatting and order.
// A file it cannot parse (VS Code allows comments) is left alone.
func upsertJSON(path, key string, entry serverEntry) (backup string, err error) {
	old, err := os.ReadFile(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return "", err
	}
	if len(bytes.TrimSpace(old)) == 0 {
		out, _ := json.MarshalIndent(map[string]any{key: map[string]any{"tensorleap": entry}}, "", "  ")
		return writeWithBackup(path, old, append(out, '\n'))
	}
	if !json.Valid(old) {
		return "", fmt.Errorf("%s is not plain JSON, so it was left as is", tilde(path))
	}
	out, err := patchJSON(old, key, "tensorleap", entry)
	if err != nil {
		return "", fmt.Errorf("%s: %v, so it was left as is", tilde(path), err)
	}
	if bytes.Equal(out, old) {
		return "", errUnchanged
	}
	return writeWithBackup(path, old, out)
}

// patchJSON returns src with obj[key][name] = value, touching no other bytes
func patchJSON(src []byte, key, name string, value any) ([]byte, error) {
	dec := json.NewDecoder(bytes.NewReader(src))
	if t, err := dec.Token(); err != nil || t != json.Delim('{') {
		return nil, errors.New("the top level is not an object")
	}
	topLast := -1
	for dec.More() {
		kt, err := dec.Token()
		if err != nil {
			return nil, err
		}
		if kt != key {
			var raw json.RawMessage
			if err := dec.Decode(&raw); err != nil {
				return nil, err
			}
			topLast = int(dec.InputOffset())
			continue
		}
		if t, err := dec.Token(); err != nil || t != json.Delim('{') {
			return nil, fmt.Errorf("%q is not an object", key)
		}
		open := int(dec.InputOffset()) - 1
		last, count := -1, 0
		for dec.More() {
			nt, err := dec.Token()
			if err != nil {
				return nil, err
			}
			keyEnd := int(dec.InputOffset())
			var raw json.RawMessage
			if err := dec.Decode(&raw); err != nil {
				return nil, err
			}
			end := int(dec.InputOffset())
			if nt == name {
				start := keyEnd + bytes.IndexFunc(src[keyEnd:], func(r rune) bool { return r != ' ' && r != '\t' && r != '\n' && r != '\r' && r != ':' })
				var have, want any
				_ = json.Unmarshal(raw, &have)
				b, _ := json.Marshal(value)
				_ = json.Unmarshal(b, &want)
				if fmt.Sprint(have) == fmt.Sprint(want) {
					return src, nil
				}
				return splice(src, start, end, indented(value, lineIndent(src, start))), nil
			}
			last, count = end, count+1
		}
		inner := lineIndent(src, open) + "  "
		if count > 0 {
			inner = lineIndent(src, last)
		}
		entry := fmt.Sprintf("%q: %s", name, indented(value, inner))
		if count == 0 {
			close := open + 1 + bytes.IndexByte(src[open+1:], '}')
			return splice(src, open+1, close, "\n"+inner+entry+"\n"+lineIndent(src, open)), nil
		}
		return splice(src, last, last, ",\n"+inner+entry), nil
	}
	block := map[string]any{name: value}
	if topLast < 0 {
		open := bytes.IndexByte(src, '{')
		close := bytes.LastIndexByte(src, '}')
		return splice(src, open+1, close, fmt.Sprintf("\n  %q: %s\n", key, indented(block, "  "))), nil
	}
	return splice(src, topLast, topLast, fmt.Sprintf(",\n%s%q: %s", lineIndent(src, topLast), key, indented(block, lineIndent(src, topLast)))), nil
}

func splice(src []byte, from, to int, text string) []byte {
	out := append([]byte{}, src[:from]...)
	out = append(out, text...)
	return append(out, src[to:]...)
}

// lineIndent is the leading whitespace of the line containing offset
func lineIndent(src []byte, offset int) string {
	start := bytes.LastIndexByte(src[:offset], '\n') + 1
	end := start
	for end < len(src) && (src[end] == ' ' || src[end] == '\t') {
		end++
	}
	return string(src[start:end])
}

func indented(v any, indent string) string {
	b, _ := json.MarshalIndent(v, indent, "  ")
	return string(b)
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
