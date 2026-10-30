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
	"github.com/spf13/viper"
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
		Long: `Runs the Tensorleap MCP server for an AI assistant (Claude Code, Claude Desktop, Cursor,
VS Code/Copilot, Codex, Windsurf), so it can read this Tensorleap server's insights, metrics
and samples. Your assistant starts it; to connect one, run: leap mcp config
Run in a terminal, it checks the connection instead: who you are signed in as, whether the
MCP server answers, and what AI access withholds.

The server is chosen from, in order: --env <name>, TL_API_URL/TL_API_KEY, the current 'leap auth' environment.`,
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			interactive := term.IsTerminal(int(os.Stdin.Fd()))
			if !interactive {
				// an assistant only shows "failed to start", so the server starts even when it can't connect yet and tells why
				return mcpPkg.NewSignedInServer(signin(envName), version.CliVersion).Run(cmd.Context(), &sdk.StdioTransport{})
			}
			env, err := chooseEnv(envName)
			if err != nil {
				return err
			}
			// a person ran it: there is nothing to serve on a terminal, so show the state and how to connect
			out := ui{w: cmd.OutOrStdout()}
			out.title("Tensorleap MCP server")
			out.text("Your AI assistant starts this command itself. To connect one, run: %s", codeStyle.Render("leap mcp config"))
			out.text("")
			st, err := showStatus(cmd.Context(), out, mcpPkg.NewClient(env.ApiUrl, env.ApiKey))
			showAccess(out, st)
			out.text("")
			if err != nil {
				return ErrShown
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&envName, "env", "", "Name of the 'leap auth' environment to serve (default: the current one)")
	cmd.AddCommand(newConfigCmd())
	return cmd
}

// signin reads the sign-in from the config file on every call, so a leap auth login made after the
// assistant started leap mcp reaches it without a restart
func signin(envName string) func() (*mcpPkg.Client, error) {
	return func() (*mcpPkg.Client, error) {
		_ = viper.ReadInConfig()
		env, err := chooseEnv(envName)
		if err != nil {
			return nil, fmt.Errorf("%v. Fix it in a terminal (leap auth login, or the --env in your assistant's MCP config)", err)
		}
		return mcpPkg.NewClient(env.ApiUrl, env.ApiKey), nil
	}
}

// showStatus checks the server the way setup does: sign-in, the MCP endpoint, and what AI access withholds
// A rejected key is fatal for setup; an unreachable server is not, since the assistant connects later.
func showStatus(ctx context.Context, out ui, client *mcpPkg.Client) (*mcpPkg.ServerStatus, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	var st *mcpPkg.ServerStatus
	var err error
	tools, probeErr := 0, error(nil)
	spin(out.w, "Checking "+client.UIBase()+"…", func() {
		if st, err = mcpPkg.Status(ctx, client); err == nil {
			tools, probeErr = mcpPkg.Probe(ctx, client)
		}
	})
	if err != nil && strings.Contains(err.Error(), "API key") {
		out.fail("%s", err)
		return nil, err
	}
	if err != nil {
		out.warn("Can't reach %s right now", client.UIBase())
		out.detail("%s", unreachableReason(err))
		return nil, nil
	}
	out.ok("Signed in to %s as %s", st.UI, st.Email)
	if probeErr != nil {
		out.warn("%s", probeErr)
	} else {
		out.ok("Tensorleap MCP server answers (%d tools)", tools)
	}
	return st, nil
}

// unreachableReason turns a transport error into the likely cause a person can act on
func unreachableReason(err error) string {
	msg := err.Error()
	switch {
	case strings.Contains(msg, "no such host"):
		return "its address doesn't resolve (is the URL right, or does it need a VPN?)"
	case strings.Contains(msg, "connection refused"):
		return "nothing is listening there (is the server running, or is the tunnel down?)"
	case strings.Contains(msg, "deadline exceeded"), strings.Contains(msg, "timeout"):
		return "it didn't answer in time (VPN or tunnel down?)"
	}
	return truncate(msg, 160)
}

func truncate(s string, n int) string {
	if len([]rune(s)) <= n {
		return s
	}
	return string([]rune(s)[:n]) + "…"
}

// showAccess says what the assistant will not be able to read, and who can change it
func showAccess(out ui, st *mcpPkg.ServerStatus) {
	if st == nil {
		return
	}
	if len(st.Off) == 0 {
		out.ok("AI access: everything is on")
	} else {
		out.warn("AI access: %s %s turned off", strings.Join(st.Off, " and "), map[bool]string{true: "is", false: "are"}[len(st.Off) == 1])
		out.detail("The assistant says when something is withheld. An admin can turn it on:\ngear icon (top right) > AI ACCESS")
	}
	switch st.CustomProjects {
	case 0:
	case 1:
		out.detail("1 project has its own AI access settings")
	default:
		out.detail("%d projects have their own AI access settings", st.CustomProjects)
	}
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
			return nil, fmt.Errorf("%w to Tensorleap; run: leap auth login", auth.ErrNotLoggedIn)
		}
		hint := ""
		if c := closest(name, known); c != "" {
			hint = fmt.Sprintf(" Did you mean %s?", c)
		}
		return nil, fmt.Errorf("No 'leap auth' environment is named %q.%s Yours: %s", name, hint, strings.Join(known, ", "))
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
