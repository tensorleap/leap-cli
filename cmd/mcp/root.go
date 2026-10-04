package mcp

import (
	"os"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
	mcpPkg "github.com/tensorleap/leap-cli/pkg/mcp"
	"github.com/tensorleap/leap-cli/pkg/version"
)

func NewMcpCmd() *cobra.Command {
	var envName string
	cmd := &cobra.Command{
		Use:   "mcp",
		Short: "Serve Tensorleap analysis results to an AI assistant over MCP (stdio)",
		Long: `Runs a read-only Model Context Protocol server on stdin/stdout so an AI assistant
(Copilot, Cursor, Claude Code, Codex, ...) can query this Tensorleap server's analysis results.

The server is chosen from, in order: TL_API_URL/TL_API_KEY, --env <name>, the current 'leap auth' environment.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			env := auth.GetCurrentEnv()
			if envName != "" {
				named, err := auth.GetEnvAuth(envName)
				if err != nil {
					return err
				}
				env = named
			}
			if u := os.Getenv("TL_API_URL"); u != "" {
				env = &auth.Env{ApiUrl: u, ApiKey: os.Getenv("TL_API_KEY")}
			}
			if env.ApiUrl == "" {
				return auth.ErrNotLoggedIn
			}
			server := mcpPkg.NewServer(mcpPkg.NewClient(env.ApiUrl, env.ApiKey), version.CliVersion)
			return server.Run(cmd.Context(), &sdk.StdioTransport{})
		},
	}
	cmd.Flags().StringVar(&envName, "env", "", "Name of the 'leap auth' environment to serve (default: the current one)")
	cmd.AddCommand(newConfigCmd())
	return cmd
}
