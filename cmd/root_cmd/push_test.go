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

		// --eval --no-wait hands the evaluate to the server, so every parameter
		// must be on the command line: nothing can be prompted for after the
		// command returns.
		{"no-wait eval with batch and model", pushInputs{noWait: true, runEval: true, batch: "8", modelPath: "m.h5"}, false},
		{"no-wait eval with batch and overwrite", pushInputs{noWait: true, runEval: true, batch: "latest", overwriteVersionRef: "v1"}, false},
		{"no-wait eval without batch", pushInputs{noWait: true, runEval: true, modelPath: "m.h5"}, true},
		{"no-wait eval without model or overwrite", pushInputs{noWait: true, runEval: true, batch: "8"}, true},
		// -u implies --eval, so this is a --no-wait chain request for an
		// update-evaluate, which is not chainable server-side.
		{"no-wait with update", pushInputs{noWait: true, batch: "8", modelPath: "m.h5", updateParts: []string{"viz"}}, true},
		// --no-wait on its own still needs nothing extra.
		{"no-wait without eval", pushInputs{noWait: true}, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if err := validatePushInputs(&c.in); (err != nil) != c.wantErr {
				t.Fatalf("wantErr=%v, got %v", c.wantErr, err)
			}
		})
	}
}
