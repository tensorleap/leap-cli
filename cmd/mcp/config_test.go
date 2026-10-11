package mcp

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestEverySnippetIsValidAndCopilotNoteOnlyForVSCode(t *testing.T) {
	for _, c := range clientNames {
		var buf bytes.Buffer
		printConfig(&buf, c, "/usr/local/bin/leap", []string{"mcp"})
		out := buf.String()
		if strings.Contains(out, "Copilot Business") != (c == "vscode") {
			t.Errorf("%s: Copilot note shown=%v", c, !(c == "vscode"))
		}
		if c == "codex" {
			if !strings.Contains(out, "command = \"/usr/local/bin/leap\"\nargs = [\"mcp\"]") {
				t.Errorf("codex snippet:\n%s", out)
			}
			continue
		}
		start, end := strings.Index(out, "{"), strings.LastIndex(out, "}")
		var v map[string]any
		if err := json.Unmarshal([]byte(out[start:end+1]), &v); err != nil {
			t.Errorf("%s: invalid JSON: %v", c, err)
		}
		if strings.Index(out, `"command"`) > strings.Index(out, `"args"`) {
			t.Errorf("%s: command should come before args", c)
		}
	}
}

func TestShellQuoteOnlyWhenNeeded(t *testing.T) {
	if shellQuote("/usr/local/bin/leap") != "/usr/local/bin/leap" || shellQuote("/Users/a b/leap") != "'/Users/a b/leap'" {
		t.Fatal("quoting")
	}
}

func TestTeamSnippetDoesNotPinAPersonalEnvironment(t *testing.T) {
	var buf bytes.Buffer
	printConfig(&buf, "claude-code", "/usr/local/bin/leap", []string{"mcp", "--env", "local"})
	out := buf.String()
	team := out[strings.Index(out, "{"):]
	if strings.Contains(team, "--env") || !strings.Contains(out, "mcp --env local") {
		t.Fatalf("personal command keeps --env, team file must not:\n%s", out)
	}
}

func TestUpsertJSONKeepsOtherServersByteForByte(t *testing.T) {
	path := filepath.Join(t.TempDir(), "mcp.json")
	_ = os.WriteFile(path, []byte(`{"theme":"dark","mcpServers":{"github":{"command":"gh-mcp","args":["--x"]}}}`), 0o644)
	entry := serverEntry{Command: "/usr/local/bin/leap", Args: []string{"mcp"}}
	backup, err := upsertJSON(path, "mcpServers", entry)
	if err != nil || backup != path+".bak" {
		t.Fatalf("first write: %v %q", err, backup)
	}
	b, _ := os.ReadFile(path)
	if !strings.HasPrefix(string(b), `{"theme":"dark","mcpServers":{"github":{"command":"gh-mcp","args":["--x"]}`) || !json.Valid(b) || !strings.Contains(string(b), `"tensorleap"`) {
		t.Fatalf("other servers and settings must survive byte for byte:\n%s", b)
	}
	if _, err := upsertJSON(path, "mcpServers", entry); !errors.Is(err, errUnchanged) {
		t.Fatalf("a second run must report no change, got %v", err)
	}
}

func TestUpsertJSONLeavesFilesWithCommentsAlone(t *testing.T) {
	path := filepath.Join(t.TempDir(), "mcp.json")
	original := "{\n  // my servers\n  \"servers\": {}\n}\n"
	_ = os.WriteFile(path, []byte(original), 0o644)
	if _, err := upsertJSON(path, "servers", serverEntry{Command: "leap", Args: []string{"mcp"}}); err == nil || !strings.Contains(err.Error(), "left as is") {
		t.Fatalf("got %v", err)
	}
	if b, _ := os.ReadFile(path); string(b) != original {
		t.Fatalf("file changed:\n%s", b)
	}
}

func TestUpsertCodexTOMLAddsOrReplacesOnlyOurTable(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.toml")
	_ = os.WriteFile(path, []byte("[model]\nname = \"o3\"\n\n[mcp_servers.tensorleap]\ncommand = \"old\"\nargs = []\n\n[mcp_servers.other]\ncommand = \"x\"\n"), 0o644)
	if _, err := upsertCodexTOML(path, serverEntry{Command: "/usr/local/bin/leap", Args: []string{"mcp", "--env", "prod"}}); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(path)
	want := "[model]\nname = \"o3\"\n\n[mcp_servers.tensorleap]\ncommand = \"/usr/local/bin/leap\"\nargs = [\"mcp\", \"--env\", \"prod\"]\n\n[mcp_servers.other]\ncommand = \"x\"\n"
	if string(b) != want {
		t.Fatalf("got:\n%s\nwant:\n%s", b, want)
	}
	if _, err := upsertCodexTOML(path, serverEntry{Command: "/usr/local/bin/leap", Args: []string{"mcp", "--env", "prod"}}); !errors.Is(err, errUnchanged) {
		t.Fatalf("a second run must report no change, got %v", err)
	}
	fresh := filepath.Join(t.TempDir(), "config.toml")
	if _, err := upsertCodexTOML(fresh, serverEntry{Command: "leap", Args: []string{"mcp"}}); err != nil {
		t.Fatal(err)
	}
	if b, _ := os.ReadFile(fresh); string(b) != "[mcp_servers.tensorleap]\ncommand = \"leap\"\nargs = [\"mcp\"]\n" {
		t.Fatalf("new file:\n%s", b)
	}
}

func TestUnknownAssistantNamesTheChoices(t *testing.T) {
	if _, err := pickAssistant(ui{w: io.Discard}, []string{"cursr"}); err == nil || !strings.Contains(err.Error(), "cursor") {
		t.Fatalf("got %v", err)
	}
}

func TestUpsertJSONTouchesNothingButOurEntry(t *testing.T) {
	entry := serverEntry{Command: "/usr/local/bin/leap", Args: []string{"mcp"}}
	cases := map[string]struct{ before, after string }{
		"append after other servers": {
			"{\n  \"mcpServers\": {\n    \"zeta\": {\"command\": \"z\", \"args\": [\"a\", \"b\"]},\n    \"alpha\": {\n      \"command\": \"a\"\n    }\n  },\n  \"theme\": 1\n}\n",
			"{\n  \"mcpServers\": {\n    \"zeta\": {\"command\": \"z\", \"args\": [\"a\", \"b\"]},\n    \"alpha\": {\n      \"command\": \"a\"\n    },\n    \"tensorleap\": {\n      \"command\": \"/usr/local/bin/leap\",\n      \"args\": [\n        \"mcp\"\n      ]\n    }\n  },\n  \"theme\": 1\n}\n",
		},
		"replace our own entry in place": {
			"{\"mcpServers\": {\"a\": {\"command\": \"a\"}, \"tensorleap\": {\"command\": \"old\"}, \"b\": {\"command\": \"b\"}}}",
			"{\"mcpServers\": {\"a\": {\"command\": \"a\"}, \"tensorleap\": {\n  \"command\": \"/usr/local/bin/leap\",\n  \"args\": [\n    \"mcp\"\n  ]\n}, \"b\": {\"command\": \"b\"}}}",
		},
		"empty servers object": {
			"{\n  \"mcpServers\": {}\n}\n",
			"{\n  \"mcpServers\": {\n    \"tensorleap\": {\n      \"command\": \"/usr/local/bin/leap\",\n      \"args\": [\n        \"mcp\"\n      ]\n    }\n  }\n}\n",
		},
		"no servers key yet": {
			"{\n  \"theme\": \"dark\"\n}\n",
			"{\n  \"theme\": \"dark\",\n  \"mcpServers\": {\n    \"tensorleap\": {\n      \"command\": \"/usr/local/bin/leap\",\n      \"args\": [\n        \"mcp\"\n      ]\n    }\n  }\n}\n",
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := patchJSON([]byte(c.before), "mcpServers", "tensorleap", entry)
			if err != nil {
				t.Fatal(err)
			}
			if string(got) != c.after {
				t.Fatalf("got:\n%s\nwant:\n%s", got, c.after)
			}
			if !json.Valid(got) {
				t.Fatal("result is not valid JSON")
			}
		})
	}
}
