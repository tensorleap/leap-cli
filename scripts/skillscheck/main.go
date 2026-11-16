// skillscheck checks the Tensorleap skills against this CLI: every leap command and flag
// they call must exist, and no field they read from leap's JSON output may be removed,
// renamed or retyped.
//
//	go run ./scripts/skillscheck fields                              # this CLI's JSON output fields
//	go run ./scripts/skillscheck check <skills dir> [base-fields.json]
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/tensorleap/leap-cli/cmd"
	"github.com/tensorleap/leap-cli/pkg/model"
	"github.com/tensorleap/leap-cli/pkg/version"
)

// jsonOutputs are leap's machine-readable outputs, keyed by the command that prints them;
// a new JSON output the skills read belongs here.
var jsonOutputs = map[string]any{
	"run info": model.RunInfo{},
	"version":  version.SkillsInfo{},
}

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "fields" {
		out, _ := json.MarshalIndent(outputFields(), "", "  ")
		fmt.Println(string(out))
		return
	}
	if len(os.Args) < 3 || os.Args[1] != "check" {
		fmt.Fprintln(os.Stderr, "usage: skillscheck fields | skillscheck check <skills dir> [base-fields.json]")
		os.Exit(2)
	}
	files, err := skillFiles(os.Args[2])
	if err != nil || len(files) == 0 {
		fmt.Fprintf(os.Stderr, "no skill files under %s: %v\n", os.Args[2], err)
		os.Exit(2)
	}
	problems, calls := checkCommands(commandTree(cmd.RootCommand), files)
	if len(os.Args) > 3 {
		base := map[string]map[string]string{}
		if data, err := os.ReadFile(os.Args[3]); err == nil && len(data) > 0 {
			if err := json.Unmarshal(data, &base); err != nil {
				fmt.Fprintf(os.Stderr, "unreadable %s: %v\n", os.Args[3], err)
				os.Exit(2)
			}
		}
		problems = append(problems, checkFields(base, outputFields(), calls)...)
	}
	sort.Strings(problems)
	for _, p := range problems {
		fmt.Println(p)
	}
	if len(problems) > 0 {
		os.Exit(1)
	}
	fmt.Printf("OK: %d skill files use only commands, flags and output fields this CLI has.\n", len(files))
}

func skillFiles(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		switch filepath.Ext(path) {
		case ".md", ".sh", ".py":
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// --- commands and flags ---

type command struct {
	canonical  string          // "run info" for every alias path of it
	flags      map[string]bool // name or shorthand -> takes a value
	hasSubs    bool
	acceptsArg bool
}

// commandTree maps every command path ("run info", and alias paths like "runs get") to its flags.
func commandTree(root *cobra.Command) map[string]command {
	tree := map[string]command{}
	var walk func(c *cobra.Command, paths []string)
	walk = func(c *cobra.Command, paths []string) {
		flags := map[string]bool{"help": false, "h": false}
		add := func(f *pflag.Flag) {
			takesValue := f.NoOptDefVal == ""
			flags[f.Name] = takesValue
			if f.Shorthand != "" {
				flags[f.Shorthand] = takesValue
			}
		}
		c.LocalFlags().VisitAll(add)
		c.InheritedFlags().VisitAll(add)
		for _, p := range paths {
			tree[p] = command{canonical: strings.TrimSpace(strings.TrimPrefix(c.CommandPath(), root.Name())),
				flags: flags, hasSubs: c.HasSubCommands(), acceptsArg: c.Runnable()}
		}
		for _, sub := range c.Commands() {
			var subPaths []string
			for _, p := range paths {
				for _, name := range append([]string{sub.Name()}, sub.Aliases...) {
					subPaths = append(subPaths, strings.TrimSpace(p+" "+name))
				}
			}
			walk(sub, subPaths)
		}
	}
	walk(root, []string{""})
	return tree
}

// A leap call: through a CLI variable in shell, a Python argument list, `leap …` in backticks
// (strict), or a plain "leap …" in text, which counts only when a real command follows.
var anchor = regexp.MustCompile(`("?\$\{?(?:TL_CLI|CLI|TL_LEAP)\}?"?|\[\s*(?:CLI|leap_cmd\(\))\s*,|_leap\(\s*\[|\bleap\b)`)

// checkCommands also returns, per file, the commands it calls.
func checkCommands(tree map[string]command, files []string) ([]string, map[string]map[string]bool) {
	var problems []string
	calls := map[string]map[string]bool{}
	for _, file := range files {
		eachLine(file, func(n int, line string) {
			for _, loc := range anchor.FindAllStringIndex(line, -1) {
				if loc[0] > 0 && line[loc[0]-1] == '=' {
					continue // TL_CLI="$TL_CLI": an assignment, not a call
				}
				if loc[1] < len(line) && line[loc[1]] == '\'' {
					continue // '$TL_CLI' names the command in a message
				}
				a := line[loc[0]:loc[1]]
				strict := a != "leap" || (loc[0] > 0 && line[loc[0]-1] == '`')
				msg, path := checkCall(tree, tokens(line[loc[1]:], strings.Contains(a, "[")), strict)
				if msg != "" {
					problems = append(problems, fmt.Sprintf("%s:%d: %s", file, n, msg))
				}
				if path != "" {
					if calls[file] == nil {
						calls[file] = map[string]bool{}
					}
					calls[file][tree[path].canonical] = true
				}
			}
		})
	}
	return problems, calls
}

// tokens reads the words of one call: shell words, or the items of a Python list.
func tokens(rest string, pyList bool) []string {
	var out []string
	sep := " \t"
	if pyList {
		sep = ", \t"
	}
	for _, raw := range strings.FieldsFunc(rest, func(r rune) bool { return strings.ContainsRune(sep, r) }) {
		if i := strings.IndexAny(raw, "`|;&>)]"); i >= 0 {
			if t := strings.Trim(raw[:i], `"'`); t != "" {
				out = append(out, t)
			}
			break
		}
		t := strings.Trim(raw, `"'`)
		if t == "" {
			continue
		}
		if !pyList && strings.ContainsAny(t[len(t)-1:], ".,:") {
			out = append(out, strings.TrimRight(t, ".,:"))
			break
		}
		out = append(out, t)
	}
	return out
}

var word = regexp.MustCompile(`^[a-z][a-z-]*$`)

func checkCall(tree map[string]command, toks []string, strict bool) (string, string) {
	if len(toks) == 0 || !word.MatchString(toks[0]) {
		return "", ""
	}
	path, positional, skipValue := "", false, false
	var flags []string
	for _, t := range toks {
		if skipValue {
			skipValue = false
			continue
		}
		if strings.HasPrefix(t, "-") && t != "-" && t != "--" {
			flags = append(flags, t)
			name := strings.SplitN(strings.TrimLeft(t, "-"), "=", 2)[0]
			skipValue = !strings.Contains(t, "=") && tree[path].flags[name]
			continue
		}
		if !positional && word.MatchString(t) {
			if _, ok := tree[strings.TrimSpace(path+" "+t)]; ok {
				path = strings.TrimSpace(path + " " + t)
				continue
			}
			if path == "" {
				if strict {
					return fmt.Sprintf("`leap %s`: no such command", t), ""
				}
				return "", ""
			}
			if c := tree[path]; c.hasSubs && !c.acceptsArg {
				return fmt.Sprintf("`leap %s %s`: no such command", path, t), ""
			}
		}
		if path == "" {
			return "", ""
		}
		positional = true
	}
	if path == "" {
		return "", ""
	}
	for _, t := range flags {
		name := strings.SplitN(strings.TrimLeft(t, "-"), "=", 2)[0]
		if _, ok := tree[path].flags[name]; name != "" && !ok {
			return fmt.Sprintf("`leap %s %s`: no such flag", path, t), path
		}
	}
	return "", path
}

// --- JSON output fields ---

func outputFields() map[string]map[string]string {
	out := map[string]map[string]string{}
	for name, v := range jsonOutputs {
		fields := map[string]string{}
		collect("", reflect.TypeOf(v), fields)
		out[name] = fields
	}
	return out
}

func collect(prefix string, t reflect.Type, fields map[string]string) {
	for t.Kind() == reflect.Ptr || t.Kind() == reflect.Slice {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		key := strings.Split(f.Tag.Get("json"), ",")[0]
		if !f.IsExported() || key == "-" {
			continue
		}
		if key == "" {
			key = f.Name
		}
		path := strings.TrimPrefix(prefix+"."+key, ".")
		fields[path] = f.Type.String()
		collect(path, f.Type, fields)
	}
}

// checkFields reports every field present on the base branch that is gone or retyped here
// and that a skill calling that command reads by name.
func checkFields(base, head map[string]map[string]string, calls map[string]map[string]bool) []string {
	var problems []string
	for output, fields := range base {
		var gone []string
		for path, typ := range fields {
			if head[output][path] != typ {
				gone = append(gone, path)
			}
		}
		sort.Strings(gone)
		for _, path := range gone {
			key := path[strings.LastIndex(path, ".")+1:]
			used := regexp.MustCompile(`["']` + regexp.QuoteMeta(key) + `["']`)
			for file, called := range calls {
				if !called[output] {
					continue
				}
				first, count := 0, 0
				eachLine(file, func(n int, line string) {
					if used.MatchString(line) {
						if count == 0 {
							first = n
						}
						count++
					}
				})
				if count > 0 {
					more := ""
					if count > 1 {
						more = fmt.Sprintf(" (and %d more lines)", count-1)
					}
					problems = append(problems, fmt.Sprintf("%s:%d: reads %q%s, which `leap %s` no longer has (removed, renamed or retyped)", file, first, key, more, output))
				}
			}
		}
	}
	return problems
}

func eachLine(file string, fn func(n int, line string)) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer func() { _ = f.Close() }()
	s := bufio.NewScanner(f)
	s.Buffer(make([]byte, 1024*1024), 1024*1024)
	for n := 1; s.Scan(); n++ {
		fn(n, s.Text())
	}
}
