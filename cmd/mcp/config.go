package mcp

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	mcpPkg "github.com/tensorleap/leap-cli/pkg/mcp"
)

var clientNames = []string{"claude-code", "claude-desktop", "cursor", "vscode", "codex", "windsurf"}

var clientTitles = map[string]string{
	"claude-code":    "Claude Code",
	"claude-desktop": "Claude Desktop",
	"cursor":         "Cursor",
	"vscode":         "VS Code",
	"codex":          "OpenAI Codex",
	"windsurf":       "Windsurf",
}

func newConfigCmd() *cobra.Command {
	var envName string
	cmd := &cobra.Command{
		Use:   "config <" + strings.Join(clientNames, "|") + ">",
		Short: "Print how to connect an AI assistant to Tensorleap",
		Example: "  leap mcp config claude-code\n" +
			"  leap mcp config cursor --env prod",
		ValidArgs:     clientNames,
		SilenceUsage:  true,
		SilenceErrors: true,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 && clientTitles[args[0]] != "" {
				return nil
			}
			return fmt.Errorf("which assistant? choose one of: %s (for example: leap mcp config claude-code)", strings.Join(clientNames, ", "))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			bin, err := leapPath()
			if err != nil {
				return err
			}
			serverArgs := []string{"mcp"}
			if envName != "" {
				serverArgs = append(serverArgs, "--env", envName)
			}
			printConfig(cmd.OutOrStdout(), args[0], bin, serverArgs)
			if env, err := chooseEnv(envName); err != nil {
				fmt.Fprintln(cmd.OutOrStdout(), err.Error())
			} else {
				for _, l := range liveStatus(cmd.Context(), mcpPkg.NewClient(env.ApiUrl, env.ApiKey)) {
					fmt.Fprintln(cmd.OutOrStdout(), l)
				}
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&envName, "env", "", "Pin a 'leap auth' environment instead of following the current one")
	return cmd
}

// leapPath prefers the PATH entry (e.g. /usr/local/bin/leap) over the resolved binary, so the
// snippet keeps working after an upgrade replaces a versioned install directory.
// GUI assistants don't inherit the shell PATH, hence an absolute path at all.
func leapPath() (string, error) {
	self, err := os.Executable()
	if err != nil {
		return "", err
	}
	if onPath, err := exec.LookPath("leap"); err == nil {
		if abs, err := filepath.Abs(onPath); err == nil && sameFile(abs, self) {
			return abs, nil
		}
	}
	return self, nil
}

func sameFile(a, b string) bool {
	ra, errA := filepath.EvalSymlinks(a)
	rb, errB := filepath.EvalSymlinks(b)
	return errA == nil && errB == nil && ra == rb
}

type serverEntry struct {
	Type    string   `json:"type,omitempty"`
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

func printConfig(w io.Writer, client, bin string, args []string) {
	entry := serverEntry{Command: bin, Args: args}
	servers := map[string]any{"mcpServers": map[string]any{"tensorleap": entry}}
	p := func(format string, a ...any) { fmt.Fprintf(w, format+"\n", a...) }
	switch client {
	case "claude-code":
		p("Claude Code: run this once to add Tensorleap to all your projects:\n")
		p("  claude mcp add --scope user tensorleap -- %s %s\n", shellQuote(bin), strings.Join(args, " "))
		// --env names are personal leap auth settings, so the shared file follows each user's current one
		p("To share it with your team instead, commit this as .mcp.json in the repository (it expects leap on everyone's PATH and uses each person's current leap auth environment):\n")
		p("%s", mustJSON(map[string]any{"mcpServers": map[string]any{"tensorleap": serverEntry{Command: "leap", Args: []string{"mcp"}}}}))
	case "claude-desktop":
		p("Claude Desktop: open Settings > Developer > Edit Config (%s) and merge in:\n", desktopConfigPath())
		p("%s", mustJSON(servers))
	case "cursor":
		p("Cursor: add to ~/.cursor/mcp.json (all projects) or .cursor/mcp.json (one project):\n")
		p("%s", mustJSON(servers))
	case "vscode":
		entry.Type = "stdio"
		p("VS Code (GitHub Copilot): add to .vscode/mcp.json in your project:\n")
		p("%s", mustJSON(map[string]any{"servers": map[string]any{"tensorleap": entry}}))
		p("\nCopilot Business/Enterprise: the 'MCP servers in Copilot' policy is off by default; an organization admin must enable it.")
	case "codex":
		quoted := make([]string, len(args))
		for i, a := range args {
			quoted[i] = fmt.Sprintf("%q", a)
		}
		p("OpenAI Codex: add to ~/.codex/config.toml:\n")
		p("[mcp_servers.tensorleap]\ncommand = %q\nargs = [%s]", bin, strings.Join(quoted, ", "))
	case "windsurf":
		p("Windsurf: add to ~/.codeium/windsurf/mcp_config.json:\n")
		p("%s", mustJSON(servers))
	}
	p("\nThen restart %s and ask, for example: \"What is my model's biggest weakness in project <name>?\"\n", clientTitles[client])
}

func desktopConfigPath() string {
	if runtime.GOOS == "windows" {
		return `%APPDATA%\Claude\claude_desktop_config.json`
	}
	if runtime.GOOS == "darwin" {
		return "~/Library/Application Support/Claude/claude_desktop_config.json"
	}
	return "~/.config/Claude/claude_desktop_config.json"
}

func shellQuote(s string) string {
	if strings.ContainsAny(s, " '\"$`\\") {
		return "'" + strings.ReplaceAll(s, "'", `'\''`) + "'"
	}
	return s
}

func mustJSON(v any) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}
