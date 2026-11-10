package model

import (
	"strings"
	"testing"

	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

// The "samples" option must parse from -u, plan as an update-evaluate (never a
// full reset on its own), print a human label, and still yield to metric's reset.
func TestSamplesUpdateAction(t *testing.T) {
	cases := []struct {
		flags []string
		kind  EvaluatePlanKind
		first string
	}{
		{[]string{"samples"}, EvaluatePlanUpdate, "Evaluate newly added samples"},
		{[]string{"update_samples", "viz"}, EvaluatePlanUpdate, "Evaluate newly added samples"},
		{[]string{"samples", "metric"}, EvaluatePlanReset, "Re-evaluate (full)"},
	}
	for _, c := range cases {
		acts, err := ParseUpdateActionsFromFlags(c.flags)
		if err != nil {
			t.Fatalf("%v: %v", c.flags, err)
		}
		plan := PlanFromUpdateActions(acts)
		if plan.Kind != c.kind {
			t.Fatalf("%v: kind %v, want %v", c.flags, plan.Kind, c.kind)
		}
		if !PlanEvaluatesNewSamples(plan) {
			t.Fatalf("%v: a plan with samples must always run the evaluation", c.flags)
		}
		if got := FormatEvaluatePlan(plan)[0]; got != c.first {
			t.Fatalf("%v: first line %q, want %q", c.flags, got, c.first)
		}
	}
}

// "foundation" and its full name parse to the foundation update, plan as an
// update-evaluate (never a full reset on its own), and still yield to metric's reset.
func TestFoundationUpdateAction(t *testing.T) {
	cases := []struct {
		flags   []string
		kind    EvaluatePlanKind
		actions int
	}{
		{[]string{"foundation"}, EvaluatePlanUpdate, 1},
		{[]string{" Foundation "}, EvaluatePlanUpdate, 1},
		{[]string{"update_foundation"}, EvaluatePlanUpdate, 1},
		{[]string{"foundation", "update_foundation"}, EvaluatePlanUpdate, 1},
		{[]string{"foundation", "samples"}, EvaluatePlanUpdate, 2},
		{[]string{"foundation", "metric"}, EvaluatePlanReset, 2},
	}
	for _, c := range cases {
		acts, err := ParseUpdateActionsFromFlags(c.flags)
		if err != nil {
			t.Fatalf("%v: %v", c.flags, err)
		}
		if len(acts) != c.actions || acts[0] != tensorleapapi.UPDATEACTION_UPDATE_FOUNDATION {
			t.Fatalf("%v: actions %v, want %d starting with %q", c.flags, acts, c.actions, tensorleapapi.UPDATEACTION_UPDATE_FOUNDATION)
		}
		if plan := PlanFromUpdateActions(acts); plan.Kind != c.kind {
			t.Fatalf("%v: kind %v, want %v", c.flags, plan.Kind, c.kind)
		}
	}
}

// Unknown values keep the existing error text, and the hint lists foundation.
func TestInvalidUpdateActionHintListsFoundation(t *testing.T) {
	_, err := ParseUpdateActionsFromFlags([]string{"fundation"})
	if err == nil {
		t.Fatal("expected an error for an unknown --update value")
	}
	msg := err.Error()
	if !strings.HasPrefix(msg, `invalid --update value "fundation"`) {
		t.Fatalf("error text changed: %q", msg)
	}
	if !strings.HasSuffix(msg, "(allowed: "+updateActionAllowedHint+")") || !strings.Contains(updateActionAllowedHint, "foundation") {
		t.Fatalf("allowed-values hint does not list foundation: %q", msg)
	}
}

// Picking "Foundation embeddings" in the overwrite prompt plans an
// update-evaluate with the foundation action and prints a human label.
func TestFoundationChangeOption(t *testing.T) {
	plan := planUpdateEvaluate(map[ChangeKey]bool{ChangeFoundation: true})
	if plan.Kind != EvaluatePlanUpdate {
		t.Fatalf("kind %v, want update", plan.Kind)
	}
	if len(plan.UpdateActions) != 1 || plan.UpdateActions[0] != tensorleapapi.UPDATEACTION_UPDATE_FOUNDATION {
		t.Fatalf("actions %v, want [%s]", plan.UpdateActions, tensorleapapi.UPDATEACTION_UPDATE_FOUNDATION)
	}
	if got, want := FormatEvaluatePlan(plan), "Compute foundation-model embeddings for every sample missing them (GPU)"; len(got) != 1 || got[0] != want {
		t.Fatalf("plan lines %q, want [%q]", got, want)
	}
}
