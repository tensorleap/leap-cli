package root_cmd

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/tensorleap/leap-cli/pkg/config"
	"github.com/tensorleap/leap-cli/pkg/hint"
)

func TestValidatePushInputs(t *testing.T) {
	const interactive, headless = true, false

	cases := []struct {
		name        string
		in          pushInputs
		interactive bool
		wantErr     bool
	}{
		{"batch with eval", pushInputs{batch: "8", runEval: true}, interactive, false},
		// -u implies --eval; --batch must see the resolved state, not reject first
		{"batch with update", pushInputs{batch: "8", updateParts: []string{"model"}}, interactive, false},
		{"batch alone", pushInputs{batch: "8"}, interactive, true},
		{"novis with update", pushInputs{noVisualization: true, updateParts: []string{"model"}}, interactive, false},
		{"novis alone", pushInputs{noVisualization: true}, interactive, true},

		// --project-name is padded-safe but must not be blank when passed.
		{"project name", pushInputs{projectName: "My Project"}, interactive, false},
		{"project name padded", pushInputs{projectName: "  My Project  "}, interactive, false},
		{"project name blank", pushInputs{projectName: "   "}, interactive, true},

		// Everything a push needs is collected before the job is created, while
		// the user is still watching — so --no-wait prompts like any other push.
		{"no-wait eval, prompt for everything", pushInputs{noWait: true, runEval: true}, interactive, false},
		{"no-wait eval, prompt for batch only", pushInputs{noWait: true, runEval: true, modelPath: "m.h5"}, interactive, false},
		{"no-wait eval, prompt for model only", pushInputs{noWait: true, runEval: true, batch: "8"}, interactive, false},

		// Headless there is nobody to answer, so the same inputs must be explicit.
		{"headless no-wait eval, fully specified", pushInputs{noWait: true, runEval: true, batch: "8", modelPath: "m.h5"}, headless, false},
		{"headless no-wait eval, overwrite instead of model", pushInputs{noWait: true, runEval: true, batch: "latest", overwriteVersionRef: "v1"}, headless, false},
		{"headless no-wait eval, no batch", pushInputs{noWait: true, runEval: true, modelPath: "m.h5"}, headless, true},
		{"headless no-wait eval, no model or overwrite", pushInputs{noWait: true, runEval: true, batch: "8"}, headless, true},

		// -u implies --eval, so this is a --no-wait chain request for an
		// update-evaluate. Not chainable server-side, so no prompt would help:
		// rejected whether or not there is a terminal.
		{"no-wait with update", pushInputs{noWait: true, batch: "8", modelPath: "m.h5", updateParts: []string{"viz"}}, interactive, true},
		{"headless no-wait with update", pushInputs{noWait: true, batch: "8", modelPath: "m.h5", updateParts: []string{"viz"}}, headless, true},

		// --no-wait on its own still needs nothing extra.
		{"no-wait without eval", pushInputs{noWait: true}, interactive, false},
		{"headless no-wait without eval", pushInputs{noWait: true}, headless, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if err := validatePushInputs(&c.in, c.interactive); (err != nil) != c.wantErr {
				t.Fatalf("wantErr=%v, got %v", c.wantErr, err)
			}
		})
	}
}

func TestWantsNewVersion(t *testing.T) {
	cases := []struct {
		name string
		in   pushInputs
		want bool
	}{
		{"name only", pushInputs{modelVersionName: "v"}, true},
		{"name and model path", pushInputs{modelVersionName: "v", modelPath: "m.h5"}, true},
		{"name but overwrite wins", pushInputs{modelVersionName: "v", overwriteVersionRef: "v1"}, false},
		{"no name", pushInputs{modelPath: "m.h5"}, false},
		{"nothing", pushInputs{}, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := wantsNewVersion(&c.in); got != c.want {
				t.Fatalf("wantsNewVersion=%v, want %v", got, c.want)
			}
		})
	}
}

func TestValidatePushInputsTrimsProjectName(t *testing.T) {
	in := pushInputs{projectName: "  My Project  "}
	if err := validatePushInputs(&in, true); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if in.projectName != "My Project" {
		t.Fatalf("projectName not trimmed, got %q", in.projectName)
	}
}

func TestReportHintIsUsedUpOnlyByAPushThatWaited(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)
	if err := config.InitConfig(filepath.Join(t.TempDir(), "config.yaml")); err != nil {
		t.Fatal(err)
	}
	hint.SetAssistant("Claude Code")
	var out bytes.Buffer
	showReportHint(&out, true, "http://localhost:4589", "mnist")
	if !strings.Contains(out.String(), "Write a Tensorleap analysis report for project mnist") {
		t.Fatalf("--no-wait still shows the hint:\n%s", out.String())
	}
	// the --no-wait push failed on the server; the re-push shows the hint again
	out.Reset()
	showReportHint(&out, true, "http://localhost:4589", "mnist")
	if out.Len() == 0 {
		t.Fatal("a --no-wait push must not use the hint up")
	}
	out.Reset()
	showReportHint(&out, false, "http://localhost:4589", "mnist")
	if out.Len() == 0 {
		t.Fatal("still due")
	}
	out.Reset()
	showReportHint(&out, true, "http://localhost:4589", "mnist")
	if out.Len() > 0 {
		t.Fatalf("a push that waited and succeeded uses it up:\n%s", out.String())
	}
}

func TestReportHintWithoutAnAssistantPointsToTheResults(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)
	if err := config.InitConfig(filepath.Join(t.TempDir(), "config.yaml")); err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	showReportHint(&out, false, "http://localhost:4589", "mnist")
	if out.String() != "\nWhen the evaluation finishes, see the results at http://localhost:4589\nPrefer an AI assistant to do this? Run leap mcp config.\n" {
		t.Fatalf("got %q", out.String())
	}
}
