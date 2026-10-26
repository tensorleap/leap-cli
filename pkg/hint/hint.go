// Package hint walks a new user through Tensorleap with their AI assistant: install, integrate, analyze.
// Each step is suggested once, by the command that makes it possible.
package hint

import (
	"fmt"
	"io"
	"slices"

	"github.com/spf13/viper"
	"github.com/tensorleap/leap-cli/pkg/config"
)

type Step struct{ key, ask string }

var (
	Install   = Step{"install", "Install Tensorleap on this machine"}
	Integrate = Step{"integrate", "Write a Tensorleap integration for my model"}
	Report    = Step{"report", "Write a Tensorleap analysis report for project %s"}
)

const (
	assistantKey = "hints.assistant"
	shownKey     = "hints.shown"
)

// SetAssistant records the assistant leap mcp config set up; hints name it, and wait for one to exist
func SetAssistant(title string) {
	viper.Set(assistantKey, title)
	_ = config.Save()
}

// Show prints what to ask the assistant next, unless no assistant is set up or this step was already suggested
func Show(w io.Writer, lead string, s Step, args ...any) bool {
	assistant := viper.GetString(assistantKey)
	shown := viper.GetStringSlice(shownKey)
	if assistant == "" || slices.Contains(shown, s.key) {
		return false
	}
	fmt.Fprintf(w, "\n%s, ask %s:\n    %s\n", lead, assistant, fmt.Sprintf(s.ask, args...))
	viper.Set(shownKey, append(shown, s.key))
	_ = config.Save()
	return true
}
