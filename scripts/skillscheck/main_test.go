package main

import (
	"strings"
	"testing"
)

func TestCheckCall(t *testing.T) {
	tree := map[string]command{
		"run":                      {canonical: "run", hasSubs: true},
		"run info":                 {canonical: "run info", acceptsArg: true, flags: map[string]bool{"output": true, "o": true}},
		"server tools kubectl":     {canonical: "server tools kubectl", hasSubs: true, acceptsArg: true, flags: map[string]bool{"n": true}},
		"server tools kubectl get": {canonical: "server tools kubectl get", acceptsArg: true, flags: map[string]bool{"n": true, "l": true}},
		"server":                   {canonical: "server", hasSubs: true},
		"server tools":             {canonical: "server tools", hasSubs: true},
	}
	cases := []struct {
		rest    string
		pyList  bool
		strict  bool
		wantErr string
		path    string
	}{
		{` run info $ID -o json 2>/dev/null`, false, true, "", "run info"},
		{` "run", "info", run_id], capture_output=True)`, true, true, "", "run info"},
		{` run infox $ID`, false, true, "no such command", ""},
		{` run info $ID --nope`, false, true, "no such flag", "run info"},
		{` server tools kubectl -n tensorleap get pod -l app=mongodb`, false, false, "", "server tools kubectl get"},
		{` answers yes when asked.`, false, false, "", ""},
		{` deploy now`, false, true, "no such command", ""},
	}
	for _, c := range cases {
		msg, path := checkCall(tree, tokens(c.rest, c.pyList), c.strict)
		if (c.wantErr == "") != (msg == "") || (c.wantErr != "" && !strings.Contains(msg, c.wantErr)) || path != c.path {
			t.Errorf("%q: got (%q, %q), want error %q and path %q", c.rest, msg, path, c.wantErr, c.path)
		}
	}
}
