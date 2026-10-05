package mcp

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
	mcpPkg "github.com/tensorleap/leap-cli/pkg/mcp"
	"github.com/tensorleap/leap-cli/pkg/version"
	"golang.org/x/term"
)

func NewMcpCmd() *cobra.Command {
	var envName string
	cmd := &cobra.Command{
		Use:   "mcp",
		Short: "Let an AI assistant read your Tensorleap analysis results (MCP server)",
		Long: `Runs a read-only Model Context Protocol server on stdin/stdout so an AI assistant
(Claude Code, Claude Desktop, Cursor, VS Code/Copilot, Codex, Windsurf) can read this
Tensorleap server's insights, metrics and samples. Your assistant starts it; you don't
run it yourself. To connect an assistant: leap mcp config <assistant>

The server is chosen from, in order: --env <name>, TL_API_URL/TL_API_KEY, the current 'leap auth' environment.`,
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			interactive := term.IsTerminal(int(os.Stdin.Fd()))
			env, err := chooseEnv(envName)
			if err != nil && interactive {
				return err
			}
			if err != nil {
				// an assistant only shows "failed to start"; serving the reason lets it tell the user
				failed := fmt.Errorf("leap mcp can't reach Tensorleap: %v. Fix it in a terminal (leap auth login, or the --env in your assistant's MCP config), then restart the assistant", err)
				return mcpPkg.NewServer(mcpPkg.NewFailedClient(failed), version.CliVersion).Run(cmd.Context(), &sdk.StdioTransport{})
			}
			client := mcpPkg.NewClient(env.ApiUrl, env.ApiKey)
			if interactive {
				fmt.Fprintf(os.Stderr, "leap mcp is started by your AI assistant; you don't need to run it.\n"+
					"Connect one: leap mcp config <%s>\n", strings.Join(clientNames, "|"))
				for _, l := range liveStatus(cmd.Context(), client) {
					fmt.Fprintln(os.Stderr, l)
				}
				fmt.Fprintln(os.Stderr, "Waiting for an assistant on stdin/stdout. Press Ctrl-C to quit.")
			}
			server := mcpPkg.NewServer(client, version.CliVersion)
			return server.Run(cmd.Context(), &sdk.StdioTransport{})
		},
	}
	cmd.Flags().StringVar(&envName, "env", "", "Name of the 'leap auth' environment to serve (default: the current one)")
	cmd.AddCommand(newConfigCmd())
	return cmd
}

func liveStatus(ctx context.Context, client *mcpPkg.Client) []string {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return mcpPkg.StatusLines(ctx, client)
}

func chooseEnv(name string) (*auth.Env, error) {
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
	if u := os.Getenv("TL_API_URL"); u != "" {
		return &auth.Env{ApiUrl: u, ApiKey: os.Getenv("TL_API_KEY")}, nil
	}
	env := auth.GetCurrentEnv()
	if env.ApiUrl == "" {
		return nil, fmt.Errorf("%w to Tensorleap; run: leap auth login", auth.ErrNotLoggedIn)
	}
	return env, nil
}
