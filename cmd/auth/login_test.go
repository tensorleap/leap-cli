package auth

import (
	"strings"
	"testing"
)

func TestNonInteractiveLoginAsksNothing(t *testing.T) {
	cases := []struct {
		name                     string
		hasUrl                   bool
		apiKey, userName, passwd string
		want, wantErr            string
	}{
		{name: "browser link by default", hasUrl: true, want: "browser_oauth"},
		{name: "api key flag", hasUrl: true, apiKey: "k", want: "api_key_flag"},
		{name: "api key wins over username", hasUrl: true, apiKey: "k", userName: "u", want: "api_key_flag"},
		{name: "username and password", hasUrl: true, userName: "u", passwd: "p", want: "username_password"},
		{name: "username alone would prompt", hasUrl: true, userName: "u", wantErr: "both --username and --password"},
		{name: "password alone would prompt", hasUrl: true, passwd: "p", wantErr: "both --username and --password"},
		{name: "no url would prompt", wantErr: "needs the server URL"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := nonInteractiveMethod(c.hasUrl, c.apiKey, c.userName, c.passwd)
			if c.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), c.wantErr) {
					t.Fatalf("want error %q, got %q, %v", c.wantErr, got, err)
				}
				return
			}
			if err != nil || got != c.want {
				t.Fatalf("got %q, %v; want %q", got, err, c.want)
			}
		})
	}
}

func TestNonInteractiveFlagIsParsedAndFailsWithoutPrompting(t *testing.T) {
	cmd := NewLoginCmd()
	if f := cmd.Flags().Lookup("non-interactive"); f == nil || f.DefValue != "false" {
		t.Fatalf("flag: %+v", f)
	}
	var out strings.Builder
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetIn(strings.NewReader("")) // a prompt would hit EOF, as it does under Claude Code's `!`
	cmd.SetArgs([]string{"--non-interactive", "--username", "u"})
	err := cmd.Execute()
	if err == nil || !strings.Contains(err.Error(), "needs the server URL") {
		t.Fatalf("got %v", err)
	}
}
