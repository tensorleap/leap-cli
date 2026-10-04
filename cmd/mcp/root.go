package mcp

import (
	"fmt"
	"os"
	"sort"
	"strings"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
	mcpPkg "github.com/tensorleap/leap-cli/pkg/mcp"
	"github.com/tensorleap/leap-cli/pkg/version"
	"golang.org/x/term"
)

func NewMcpCmd() *cobra.Command {
	var envName string
	var allowLegacy bool
	cmd := &cobra.Command{
		Use:   "mcp",
		Short: "Let an AI assistant read your Tensorleap analysis results (MCP server)",
		Long: `Runs a read-only Model Context Protocol server on stdin/stdout so an AI assistant
(Claude Code, Claude Desktop, Cursor, VS Code/Copilot, Codex, Windsurf) can read this
Tensorleap server's insights, metrics and samples. Your assistant starts it; you don't
run it yourself. To connect an assistant: leap mcp config <assistant>

The server is chosen from, in order: TL_API_URL/TL_API_KEY, --env <name>, the current 'leap auth' environment.`,
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			env, err := chooseEnv(envName)
			if err != nil {
				return err
			}
			if term.IsTerminal(int(os.Stdin.Fd())) {
				fmt.Fprintf(os.Stderr, "leap mcp is started by your AI assistant, not by hand.\n"+
					"To connect one, run: leap mcp config <%s>\n"+
					"Serving %s on stdin/stdout anyway; press Ctrl-C to exit.\n", strings.Join(clientNames, "|"), strings.TrimSuffix(strings.TrimRight(env.ApiUrl, "/"), "/api/v2"))
			}
			server := mcpPkg.NewServer(mcpPkg.NewClient(env.ApiUrl, env.ApiKey), version.CliVersion, allowLegacy)
			return server.Run(cmd.Context(), &sdk.StdioTransport{})
		},
	}
	cmd.Flags().StringVar(&envName, "env", "", "Name of the 'leap auth' environment to serve (default: the current one)")
	cmd.Flags().BoolVar(&allowLegacy, "allow-legacy-server", false, "Serve a Tensorleap server that predates AI access controls (nothing on the server limits what the assistant receives)")
	cmd.AddCommand(newConfigCmd())
	return cmd
}

func chooseEnv(name string) (*auth.Env, error) {
	if u := os.Getenv("TL_API_URL"); u != "" {
		return &auth.Env{ApiUrl: u, ApiKey: os.Getenv("TL_API_KEY")}, nil
	}
	if name != "" {
		env, err := auth.GetEnvAuth(name)
		if err == nil {
			return env, nil
		}
		known := []string{}
		for _, e := range auth.GetEnvs() {
			known = append(known, e.Name)
		}
		sort.Strings(known)
		if len(known) == 0 {
			return nil, fmt.Errorf("no 'leap auth' environment named %q; log in first with: leap auth login", name)
		}
		return nil, fmt.Errorf("no 'leap auth' environment named %q (known: %s)", name, strings.Join(known, ", "))
	}
	env := auth.GetCurrentEnv()
	if env.ApiUrl == "" {
		return nil, fmt.Errorf("%w to Tensorleap; run: leap auth login", auth.ErrNotLoggedIn)
	}
	return env, nil
}
