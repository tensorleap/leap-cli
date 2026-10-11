package mcp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
	mcpPkg "github.com/tensorleap/leap-cli/pkg/mcp"
)

const (
	skillsInstaller   = "https://raw.githubusercontent.com/tensorleap/skills/main/install.sh"
	skillsMarketplace = "tensorleap/skills"
	analysisPlugin    = "analysis@tensorleap"
)

// ErrShown is returned after a failure was already printed, so the CLI exits 1 without a second message
var ErrShown = errors.New("")

var errShown = ErrShown

func newConfigCmd() *cobra.Command {
	var envName string
	var printOnly, noSkills bool
	cmd := &cobra.Command{
		Use:   "config [" + strings.Join(clientNames, "|") + "]",
		Short: "Connect an AI assistant to Tensorleap (MCP server and skills)",
		Long: `Sets up an AI assistant to read your Tensorleap analysis: registers the Tensorleap MCP
server in the assistant's configuration, installs the Tensorleap skills where the assistant
supports them, and checks that the server answers. Run it without arguments to be asked.

Claude Code gets the 'analysis' plugin, which carries both the MCP server and the skills.
Use --print to see the configuration without changing anything.`,
		Example: "  leap mcp config\n" +
			"  leap mcp config cursor\n" +
			"  leap mcp config vscode --env prod\n" +
			"  leap mcp config codex --print",
		ValidArgs:     clientNames,
		Args:          cobra.MaximumNArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			out := ui{w: cmd.OutOrStdout()}
			asking := len(args) == 0 && interactive()
			if asking {
				out.title("Connect an AI assistant to Tensorleap")
			}
			a, err := pickAssistant(out, args)
			if errors.Is(err, errCancelled) {
				out.text("Cancelled; nothing was changed.")
				return errShown
			}
			if err != nil {
				out.text("%s", err.Error())
				return errShown
			}
			if envName == "" && asking {
				if envName, err = pickEnv(out); errors.Is(err, errCancelled) {
					out.text("Cancelled; nothing was changed.")
					return errShown
				}
			}
			bin, err := leapPath()
			if err != nil {
				return err
			}
			serverArgs := []string{"mcp"}
			if envName != "" {
				serverArgs = append(serverArgs, "--env", envName)
			}
			if printOnly {
				printConfig(cmd.OutOrStdout(), a.id, bin, serverArgs)
				return nil
			}
			if !asking {
				out.title("Connecting " + a.title + " to Tensorleap")
			}
			return setup(cmd.Context(), out, a, bin, serverArgs, envName, !noSkills)
		},
	}
	cmd.Flags().StringVar(&envName, "env", "", "Pin a 'leap auth' environment instead of following the current one")
	cmd.Flags().BoolVar(&printOnly, "print", false, "Print the configuration to add by hand; change nothing")
	cmd.Flags().BoolVar(&noSkills, "no-skills", false, "Register the MCP server only, without the Tensorleap skills")
	cmd.SetFlagErrorFunc(func(c *cobra.Command, err error) error {
		fmt.Fprintf(c.OutOrStdout(), "%s\nRun 'leap mcp config --help' for usage.\n", err)
		return errShown
	})
	return cmd
}

func pickAssistant(out ui, args []string) (*assistant, error) {
	if len(args) == 1 {
		if a := findAssistant(args[0]); a != nil {
			return a, nil
		}
		hint := ""
		if c := closest(args[0], clientNames); c != "" {
			hint = fmt.Sprintf(" Did you mean %s?", c)
		}
		return nil, fmt.Errorf("%q is not an assistant leap can set up.%s\nChoose one of: %s", args[0], hint, strings.Join(clientNames, ", "))
	}
	if !interactive() {
		return nil, fmt.Errorf("Which assistant? Run: leap mcp config <%s>", strings.Join(clientNames, "|"))
	}
	choices := make([]choice, len(assistants))
	def := -1
	for i, a := range assistants {
		choices[i] = choice{label: a.title, detail: "MCP server"}
		if a.skillsTool != "" || a.id == "claude-code" {
			choices[i].detail = "MCP server + skills"
		}
		if a.detect() {
			choices[i].badge = "● installed"
			if def < 0 {
				def = i
			}
		}
	}
	i, err := choose(out, "Which AI assistant should read your Tensorleap analysis?", "Assistant", choices, max(def, 0))
	if err != nil {
		return nil, err
	}
	return &assistants[i], nil
}

// pickEnv asks which server when there is more than one leap environment; "" keeps the current one
func pickEnv(out ui) (string, error) {
	envs := auth.GetEnvs()
	if len(envs) < 2 {
		return "", nil
	}
	current := currentEnvName()
	// the current server first, then by name, so the list never reorders between runs
	sort.SliceStable(envs, func(i, j int) bool {
		if (envs[i].Name == current) != (envs[j].Name == current) {
			return envs[i].Name == current
		}
		return envs[i].Name < envs[j].Name
	})
	choices := make([]choice, len(envs))
	def := 0
	for i, e := range envs {
		choices[i] = choice{label: e.Name, detail: mcpPkg.NewClient(e.ApiUrl, "").UIBase()}
		if e.Name == current {
			choices[i].badge = "● current"
			def = i
		}
	}
	i, err := choose(out, "Which Tensorleap server?", "Server", choices, def)
	if err != nil {
		return "", err
	}
	if envs[i].Name == current {
		return "", nil
	}
	return envs[i].Name, nil
}

func setup(ctx context.Context, out ui, a *assistant, bin string, serverArgs []string, envName string, withSkills bool) error {
	env, err := chooseEnv(envName)
	if errors.Is(err, auth.ErrNotLoggedIn) {
		out.fail("You're not signed in to Tensorleap")
		out.detail("Sign in, then run this again:")
		out.command("leap auth login")
		return errShown
	}
	if err != nil {
		out.fail("%s", err)
		return errShown
	}
	client := mcpPkg.NewClient(env.ApiUrl, env.ApiKey)
	st, err := showStatus(ctx, out, client)
	if err != nil {
		return errShown
	}
	if st == nil {
		out.detail("The assistant is still set up; it connects once the server is reachable.")
	}

	entry := serverEntry{Command: bin, Args: serverArgs}
	skills := false
	if a.id == "claude-code" {
		if skills, err = setupClaudeCode(out, entry, envName, withSkills); err != nil {
			return err
		}
	} else {
		configureFile(out, a, entry)
		if withSkills {
			skills = installSkills(out, a)
		} else {
			out.skip("Skills: skipped (--no-skills)")
		}
	}

	showAccess(out, st)

	project := "<project>"
	if st != nil && len(st.Projects) > 0 {
		project = st.Projects[0]
	}
	out.text("\n%s %s, then ask:", okStyle.Render("Done."), a.restart)
	out.command(fmt.Sprintf("What is my model's biggest weakness in project %s?", project))
	if skills {
		out.command(fmt.Sprintf("Write a Tensorleap analysis report for project %s", project))
	}
	out.text("")
	return nil
}

func configureFile(out ui, a *assistant, entry serverEntry) {
	path, key := configFile(a.id)
	if a.id == "vscode" {
		entry.Type = "stdio"
	}
	var backup string
	var err error
	if a.id == "codex" {
		backup, err = upsertCodexTOML(path, entry)
	} else {
		backup, err = upsertJSON(path, key, entry)
	}
	switch {
	case errors.Is(err, errUnchanged):
		out.ok("%s already has the Tensorleap MCP server (%s)", a.title, tilde(path))
	case err != nil:
		out.fail("Could not update %s: %v", tilde(path), err)
		out.detail("Add this by hand (leap mcp config %s --print shows it in full):", a.id)
		var buf bytes.Buffer
		printConfig(&buf, a.id, entry.Command, entry.Args)
		out.detail("%s", strings.TrimSpace(buf.String()))
	default:
		out.ok("Added the Tensorleap MCP server to %s", tilde(path))
		if backup != "" {
			out.detail("previous version saved as %s", tilde(backup))
		}
	}
	if a.id == "vscode" {
		out.detail("Copilot Business/Enterprise: an organization admin must allow 'MCP servers in Copilot'.")
	}
}

func installSkills(out ui, a *assistant) bool {
	if a.skillsTool == "" {
		out.skip("Skills: %s doesn't load skills; the MCP tools work on their own", a.title)
		return false
	}
	cmdline := fmt.Sprintf("curl -fsSL %s | sh -s -- --tool %s --global", skillsInstaller, a.skillsTool)
	if !onPath("curl") || !onPath("sh") {
		out.warn("Skills not installed: this needs curl and sh. Install them with:")
		out.command(cmdline)
		return false
	}
	var res []byte
	var err error
	spin(out.w, "Installing the Tensorleap skills…", func() { res, err = exec.Command("sh", "-c", cmdline).CombinedOutput() })
	if err != nil {
		out.warn("Skills not installed: %v", err)
		out.detail("%s", strings.TrimSpace(string(res)))
		out.detail("Retry with:")
		out.command(cmdline)
		return false
	}
	out.ok("Installed the Tensorleap skills in %s (analysis, integration, install)", a.skillsDir)
	return true
}

// Claude Code gets the analysis plugin: it carries the MCP server and the skills, and updates through the
// marketplace. The plugin follows the current leap environment; another one is registered on its own.
func setupClaudeCode(out ui, entry serverEntry, envName string, withSkills bool) (skills bool, err error) {
	current := currentEnvName()
	pinned := envName != "" && envName != current
	if pinned && withSkills && interactive() {
		i, cerr := choose(out, fmt.Sprintf("The analysis plugin always uses your current leap environment (%s). How should Claude Code reach %s?", current, envName), "",
			[]choice{
				{label: "Make " + envName + " my current environment", detail: "MCP server + skills; same as leap auth select " + envName},
				{label: "Keep " + current + " current, add " + envName + " on its own", detail: "MCP server only, without the skills"},
			}, 0)
		if cerr != nil {
			out.text("Cancelled; nothing was changed.")
			return false, errShown
		}
		if i == 0 {
			if err := auth.UseEnv(envName); err != nil {
				out.fail("%s", err)
				return false, errShown
			}
			out.ok("Current leap environment is now %s", envName)
			pinned = false
		}
	}
	if !onPath("claude") {
		out.warn("The claude command is not on your PATH, so finish the setup inside Claude Code:")
		if withSkills && !pinned {
			out.command("/plugin marketplace add " + skillsMarketplace)
			out.command("/plugin install " + analysisPlugin)
		} else {
			out.command(claudeAddLine(entry))
		}
		return false, nil
	}
	if withSkills && !pinned {
		if hasClaudePlugin(analysisPlugin) {
			out.ok("Claude Code already has the Tensorleap analysis plugin (MCP server and skills)")
		} else {
			var res []byte
			var err error
			spin(out.w, "Installing the Tensorleap plugin in Claude Code…", func() {
				res, err = exec.Command("claude", "plugin", "marketplace", "add", skillsMarketplace).CombinedOutput()
				if err != nil && bytes.Contains(bytes.ToLower(res), []byte("already")) {
					err = nil
				}
				if err == nil {
					res, err = exec.Command("claude", "plugin", "install", analysisPlugin).CombinedOutput()
				}
			})
			if err != nil {
				out.fail("Could not install the analysis plugin: %s", strings.TrimSpace(string(res)))
				out.detail("Inside Claude Code, run instead:")
				out.command("/plugin marketplace add " + skillsMarketplace)
				out.command("/plugin install " + analysisPlugin)
				return false, errShown
			}
			out.ok("Installed the Tensorleap analysis plugin in Claude Code (MCP server and skills)")
		}
		// an older stand-alone registration would show every tool twice
		if exec.Command("claude", "mcp", "get", "tensorleap").Run() == nil {
			_ = exec.Command("claude", "mcp", "remove", "tensorleap", "--scope", "user").Run()
			out.detail("removed your earlier stand-alone 'tensorleap' entry; the plugin replaces it")
		}
		out.detail("follows your current leap environment (%s); change it with: leap auth select", currentEnvName())
		return true, nil
	}
	_ = exec.Command("claude", "mcp", "remove", "tensorleap", "--scope", "user").Run()
	args := append([]string{"mcp", "add", "--scope", "user", "tensorleap", "--", entry.Command}, entry.Args...)
	if res, err := exec.Command("claude", args...).CombinedOutput(); err != nil {
		out.fail("Could not register the MCP server: %s", strings.TrimSpace(string(res)))
		return false, errShown
	}
	out.ok("Registered the Tensorleap MCP server in Claude Code (all your projects)")
	switch {
	case !withSkills:
		out.skip("Skills: skipped (--no-skills)")
	case pinned:
		out.skip("Skills: not installed; the analysis plugin follows your current environment (%s)", current)
	}
	return false, nil
}

func claudeAddLine(entry serverEntry) string {
	return fmt.Sprintf("claude mcp add --scope user tensorleap -- %s %s", shellQuote(entry.Command), strings.Join(entry.Args, " "))
}

func currentEnvName() string {
	return auth.GetCurrentEnv().Name
}

// hasClaudePlugin reads Claude Code's plugin registry; CLAUDE_CONFIG_DIR moves it
func hasClaudePlugin(id string) bool {
	dir := os.Getenv("CLAUDE_CONFIG_DIR")
	if dir == "" {
		dir = home(".claude")
	}
	b, err := os.ReadFile(filepath.Join(dir, "plugins", "installed_plugins.json"))
	if err != nil {
		return false
	}
	var reg struct {
		Plugins map[string]any `json:"plugins"`
	}
	return json.Unmarshal(b, &reg) == nil && reg.Plugins[id] != nil
}

// leapPath prefers the PATH entry (e.g. /usr/local/bin/leap) over the resolved binary, so the
// configuration keeps working after an upgrade replaces a versioned install directory.
// GUI assistants don't inherit the shell PATH, hence an absolute path at all.
func leapPath() (string, error) {
	self, err := os.Executable()
	if err != nil {
		return "", err
	}
	if onPathBin, err := exec.LookPath("leap"); err == nil {
		if abs, err := filepath.Abs(onPathBin); err == nil && sameFile(abs, self) {
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

// printConfig is --print: the configuration to add by hand, changing nothing
func printConfig(w io.Writer, client, bin string, args []string) {
	entry := serverEntry{Command: bin, Args: args}
	path, key := configFile(client)
	p := func(format string, a ...any) { fmt.Fprintf(w, format+"\n", a...) }
	switch client {
	case "claude-code":
		p("Claude Code: run this once to add Tensorleap to all your projects:\n")
		p("  %s\n", claudeAddLine(entry))
		// --env names are personal leap auth settings, so the shared file follows each user's current one
		p("To share it with your team instead, commit this as .mcp.json in the repository (it expects leap on everyone's PATH and uses each person's current leap auth environment):\n")
		p("%s", mustJSON(map[string]any{"mcpServers": map[string]any{"tensorleap": serverEntry{Command: "leap", Args: []string{"mcp"}}}}))
	case "codex":
		quoted := make([]string, len(args))
		for i, a := range args {
			quoted[i] = fmt.Sprintf("%q", a)
		}
		p("OpenAI Codex: add to %s:\n", tilde(path))
		p("[mcp_servers.tensorleap]\ncommand = %q\nargs = [%s]", bin, strings.Join(quoted, ", "))
	default:
		if client == "vscode" {
			entry.Type = "stdio"
		}
		p("%s: add to %s:\n", findAssistant(client).title, tilde(path))
		p("%s", mustJSON(map[string]any{key: map[string]any{"tensorleap": entry}}))
		if client == "vscode" {
			p("\nCopilot Business/Enterprise: the 'MCP servers in Copilot' policy is off by default; an organization admin must enable it.")
		}
	}
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
