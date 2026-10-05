package mcp

import (
	"bytes"
	"encoding/json"
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
