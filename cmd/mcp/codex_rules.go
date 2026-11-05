package mcp

import (
	"os"
	"path/filepath"
)

// codexRules lets Codex run leap commands outside its sandbox without asking. The sandbox has no
// network (not even localhost), so otherwise every leap command waits for the user's approval.
// Codex applies a rule only when the whole command matches, so leap must run on its own, not in a pipe.
const codexRules = `# Written by leap mcp config: lets Codex run leap commands without asking each time
prefix_rule(pattern=["leap"], decision="allow")
`

// codexRulesPath is leap's own rules file; the user's other rules files are never touched
func codexRulesPath() string {
	return home(".codex", "rules", "tensorleap.rules")
}

func codexRulesInPlace() bool {
	b, err := os.ReadFile(codexRulesPath())
	return err == nil && string(b) == codexRules
}

func writeCodexRules() error {
	path := codexRulesPath()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(codexRules), 0o644)
}

// allowLeapInCodex writes the rule when allow is set (--codex-allow-leap), asks first in a terminal,
// and otherwise only says how to turn it on
func allowLeapInCodex(out ui, allow bool) {
	path := tilde(codexRulesPath())
	if codexRulesInPlace() {
		out.ok("Codex already runs leap commands without asking (%s)", path)
		return
	}
	if !allow && interactive() {
		i, err := choose(out, "Let Codex run leap commands without asking each time? (writes "+path+")", "",
			[]choice{
				{label: "Yes, run leap commands without asking", detail: "Codex's sandbox has no network, so each leap command would need your approval"},
				{label: "No, ask me each time"},
			}, 0)
		allow = err == nil && i == 0
	}
	if !allow {
		out.skip("Codex asks before each leap command; to stop that, run: leap mcp config codex --codex-allow-leap")
		return
	}
	if err := writeCodexRules(); err != nil {
		out.warn("Could not write %s: %v", path, err)
		return
	}
	out.ok("Codex runs leap commands without asking (%s)", path)
}
