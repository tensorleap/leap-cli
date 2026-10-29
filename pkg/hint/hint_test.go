package hint

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/tensorleap/leap-cli/pkg/config"
)

func freshConfig(t *testing.T) string {
	t.Helper()
	viper.Reset()
	t.Cleanup(viper.Reset)
	path := filepath.Join(t.TempDir(), "config.yaml")
	if err := config.InitConfig(path); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestEachStepIsSuggestedOnceAndNamesTheAssistant(t *testing.T) {
	path := freshConfig(t)
	var out bytes.Buffer
	if Show(&out, "Next", Install, Manual{}) || out.Len() > 0 {
		t.Fatalf("no assistant and nothing to do by hand, so no hint:\n%s", out.String())
	}

	SetAssistant("Cursor")
	if !Show(&out, "Next", Report, Results("http://localhost:4589"), "mnist") || !strings.Contains(out.String(), "Next, ask Cursor:\n    Write a Tensorleap analysis report for project mnist") {
		t.Fatalf("got:\n%s", out.String())
	}
	out.Reset()
	if Show(&out, "Next", Report, Results("http://localhost:4589"), "mnist") || out.Len() > 0 {
		t.Fatalf("shown twice:\n%s", out.String())
	}
	if !Show(&out, "Next", Install, Manual{}) {
		t.Fatal("another step is still due")
	}

	viper.Reset()
	if err := config.InitConfig(path); err != nil {
		t.Fatal(err)
	}
	out.Reset()
	if Show(&out, "Next", Install, Manual{}) || Show(&out, "Next", Report, Manual{}, "mnist") {
		t.Fatalf("a later leap run repeated a hint:\n%s", out.String())
	}
	if !Show(&out, "Next", Integrate, Pushed) || !strings.Contains(out.String(), "Next, open Cursor in the repo you want to integrate and ask:\n    Write a Tensorleap integration") {
		t.Fatalf("the assistant should persist too:\n%s", out.String())
	}
}

func TestWithoutAnAssistantEachCommandNamesTheNextStepByHand(t *testing.T) {
	freshConfig(t)
	var out bytes.Buffer
	if !Show(&out, "Next", Integrate, Registered("http://localhost:4589")) {
		t.Fatal("manual mode: install should say what to do next")
	}
	want := "\nNext, open http://localhost:4589 and register, then sign in:\n    leap auth login http://localhost:4589\nPrefer an AI assistant to do this? Run leap mcp config.\n"
	if out.String() != want {
		t.Fatalf("got:\n%q\nwant:\n%q", out.String(), want)
	}
	out.Reset()
	if !Show(&out, "Next", Integrate, Pushed) || !strings.Contains(out.String(), "Next, push your integration:\n    leap push\nPrefer an AI assistant") {
		t.Fatalf("login is a different manual step than install, though both lead to Integrate:\n%s", out.String())
	}
	for _, s := range []string{"folder", "directory", "repo", "local", "remote"} {
		if strings.Contains(strings.ToLower(out.String()), s) {
			t.Errorf("manual hints say nothing about %s:\n%s", s, out.String())
		}
	}
	out.Reset()
	if Show(&out, "Next", Integrate, Pushed) || out.Len() > 0 {
		t.Fatalf("manual hints are once-only too:\n%s", out.String())
	}

	// leap mcp config later switches to helper mode; the manual hints shown so far don't use up the helper ones
	SetAssistant("Claude Code")
	if !Show(&out, "Next", Integrate, Pushed) || !strings.Contains(out.String(), "ask:\n    Write a Tensorleap integration for my model\n") || strings.Contains(out.String(), "leap mcp config") {
		t.Fatalf("helper mode after leap mcp config:\n%s", out.String())
	}
}

func TestOfferLeavesTheStepDueUntilItSucceeds(t *testing.T) {
	freshConfig(t)
	SetAssistant("Claude Code")
	var out bytes.Buffer
	if Offer(&out, "When the evaluation finishes", Report, Manual{}, "mnist") != "report" || !strings.Contains(out.String(), "Write a Tensorleap analysis report for project mnist") {
		t.Fatalf("got:\n%s", out.String())
	}
	out.Reset()
	if !Show(&out, "When the evaluation finishes", Report, Manual{}, "mnist") {
		t.Fatal("an offered hint must still be due")
	}
	out.Reset()
	if Offer(&out, "When the evaluation finishes", Report, Manual{}, "mnist") != "" || out.Len() > 0 {
		t.Fatalf("shown after success, so never again:\n%s", out.String())
	}
}
