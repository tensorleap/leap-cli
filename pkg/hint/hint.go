// Package hint walks a new user through Tensorleap: install, integrate, analyze.
// Each step is suggested once, by the command that makes it possible. Once leap mcp config has set up an
// AI assistant (helper mode), a hint is what to ask the assistant; before that (manual mode), it is the
// next thing to do by hand, plus a pointer to leap mcp config.
package hint

import (
	"fmt"
	"io"
	"slices"

	"github.com/spf13/viper"
	"github.com/tensorleap/leap-cli/pkg/config"
)

// Step is what to ask the assistant (helper mode); where, if set, is where to open the assistant before asking
type Step struct{ key, ask, where string }

var (
	Install   = Step{"install", "Install Tensorleap on this machine", ""}
	Integrate = Step{"integrate", "Write a Tensorleap integration for my model", "the repo you want to integrate"}
	Report    = Step{"report", "Write a Tensorleap analysis report for project %s", ""}
)

// Manual is the next step to do by hand (manual mode); the zero value says nothing
type Manual struct{ key, do string }

// Registered follows leap server install: create the first account, then sign the CLI in
func Registered(url string) Manual {
	return Manual{"manual-register", fmt.Sprintf("open %[1]s and register, then sign in:\n    leap auth login %[1]s", url)}
}

// Pushed follows leap auth login
var Pushed = Manual{"manual-push", "push your integration:\n    leap push"}

// Results follows leap push --eval
func Results(url string) Manual {
	return Manual{"manual-results", "see the results at " + url}
}

const mcpConfigLine = "Prefer an AI assistant to do this? Run leap mcp config."

const (
	assistantKey = "hints.assistant"
	shownKey     = "hints.shown"
)

// SetAssistant records the assistant leap mcp config set up; from then on hints name it
func SetAssistant(title string) {
	viper.Set(assistantKey, title)
	_ = config.Save()
}

// Show prints the next step and marks it shown, unless it was already suggested. Call it once the step it
// follows has succeeded.
func Show(w io.Writer, lead string, s Step, m Manual, args ...any) bool {
	key := Offer(w, lead, s, m, args...)
	if key != "" {
		markShown(key)
	}
	return key != ""
}

// Offer prints the next step like Show but leaves it due, for when the step it follows may still fail
// out of sight (leap push --no-wait). It returns the key it printed, or "".
func Offer(w io.Writer, lead string, s Step, m Manual, args ...any) string {
	assistant := viper.GetString(assistantKey)
	shown := viper.GetStringSlice(shownKey)
	switch {
	case assistant != "":
		if slices.Contains(shown, s.key) {
			return ""
		}
		if s.where != "" {
			fmt.Fprintf(w, "\n%s, open %s in %s and ask:\n    %s\n", lead, assistant, s.where, fmt.Sprintf(s.ask, args...))
		} else {
			fmt.Fprintf(w, "\n%s, ask %s:\n    %s\n", lead, assistant, fmt.Sprintf(s.ask, args...))
		}
		return s.key
	case m.key != "" && !slices.Contains(shown, m.key):
		fmt.Fprintf(w, "\n%s, %s\n%s\n", lead, m.do, mcpConfigLine)
		return m.key
	}
	return ""
}

func markShown(key string) {
	viper.Set(shownKey, append(viper.GetStringSlice(shownKey), key))
	_ = config.Save()
}
