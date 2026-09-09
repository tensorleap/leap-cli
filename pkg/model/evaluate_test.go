package model

import "testing"

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
		if got := FormatEvaluatePlan(plan)[0]; got != c.first {
			t.Fatalf("%v: first line %q, want %q", c.flags, got, c.first)
		}
	}
}
