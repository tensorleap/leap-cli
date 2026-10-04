package mcp

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var clients = map[string]string{
	"claude-code":    "Claude Code: add to .mcp.json in your project (or run the command shown)",
	"claude-desktop": "Claude Desktop: merge into claude_desktop_config.json (Settings > Developer > Edit Config)",
	"cursor":         "Cursor: add to .cursor/mcp.json in your project (or ~/.cursor/mcp.json for all projects)",
	"vscode":         "VS Code / GitHub Copilot: add to .vscode/mcp.json in your project",
	"codex":          "OpenAI Codex: add to ~/.codex/config.toml",
	"windsurf":       "Windsurf: add to ~/.codeium/windsurf/mcp_config.json",
}

func newConfigCmd() *cobra.Command {
	var envName string
	cmd := &cobra.Command{
		Use:       "config <claude-code|claude-desktop|cursor|vscode|codex|windsurf>",
		Short:     "Print the configuration that connects an AI assistant to this Tensorleap server",
		Args:      cobra.ExactValidArgs(1),
		ValidArgs: []string{"claude-code", "claude-desktop", "cursor", "vscode", "codex", "windsurf"},
		RunE: func(cmd *cobra.Command, args []string) error {
			bin, err := os.Executable()
			if err != nil {
				return err
			}
			if resolved, err := filepath.EvalSymlinks(bin); err == nil {
				bin = resolved
			}
			serverArgs := []string{"mcp"}
			if envName != "" {
				serverArgs = append(serverArgs, "--env", envName)
			}
			fmt.Fprintln(cmd.OutOrStdout(), clients[args[0]]+":")
			fmt.Fprintln(cmd.OutOrStdout())
			fmt.Fprintln(cmd.OutOrStdout(), snippet(args[0], bin, serverArgs))
			fmt.Fprintln(cmd.OutOrStdout())
			fmt.Fprintln(cmd.OutOrStdout(), "Enterprise GitHub Copilot: the 'MCP servers in Copilot' policy is off by default for Business/Enterprise; an organization admin must enable it.")
			return nil
		},
	}
	cmd.Flags().StringVar(&envName, "env", "", "Pin a 'leap auth' environment instead of following the current one")
	return cmd
}

func snippet(client, bin string, args []string) string {
	entry := map[string]any{"command": bin, "args": args}
	switch client {
	case "vscode":
		entry["type"] = "stdio"
		return mustJSON(map[string]any{"servers": map[string]any{"tensorleap": entry}})
	case "codex":
		quoted := make([]string, len(args))
		for i, a := range args {
			quoted[i] = fmt.Sprintf("%q", a)
		}
		return fmt.Sprintf("[mcp_servers.tensorleap]\ncommand = %q\nargs = [%s]", bin, strings.Join(quoted, ", "))
	case "claude-code":
		return mustJSON(map[string]any{"mcpServers": map[string]any{"tensorleap": entry}}) +
			fmt.Sprintf("\n\nor: claude mcp add tensorleap -- %q %s", bin, strings.Join(args, " "))
	}
	return mustJSON(map[string]any{"mcpServers": map[string]any{"tensorleap": entry}})
}

func mustJSON(v any) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}
