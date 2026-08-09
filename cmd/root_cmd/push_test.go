package root_cmd

import "testing"

func TestValidatePushInputs(t *testing.T) {
	cases := []struct {
		name    string
		in      pushInputs
		wantErr bool
	}{
		{"batch with eval", pushInputs{batch: "8", runEval: true}, false},
		// -u implies --eval; --batch must see the resolved state, not reject first
		{"batch with update", pushInputs{batch: "8", updateParts: []string{"model"}}, false},
		{"batch alone", pushInputs{batch: "8"}, true},
		{"novis with update", pushInputs{noVisualization: true, updateParts: []string{"model"}}, false},
		{"novis alone", pushInputs{noVisualization: true}, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if err := validatePushInputs(&c.in); (err != nil) != c.wantErr {
				t.Fatalf("wantErr=%v, got %v", c.wantErr, err)
			}
		})
	}
}
