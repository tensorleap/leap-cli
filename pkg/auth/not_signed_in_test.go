package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/spf13/viper"
)

func TestNotSignedInErrorNamesTheKnownServer(t *testing.T) {
	origTTY := stdinIsTerminal
	t.Cleanup(func() { stdinIsTerminal = origTTY })
	cases := []struct {
		name, apiUrl, installed, want string
	}{
		{name: "configured api url", apiUrl: "https://api.acme.tensorleap.ai/api/v2", installed: "http://localhost:4589", want: "Not signed in. Run: leap auth login https://acme.tensorleap.ai\nAI assistant: start `leap auth login https://acme.tensorleap.ai --non-interactive` in the background, send the user the sign-in link it prints in a message, and wait for that command to finish. Only the user can sign in: don't look for API keys or passwords."},
		{name: "local install", installed: "http://localhost:4589", want: "Not signed in. Run: leap auth login http://localhost:4589\nAI assistant: start `leap auth login http://localhost:4589 --non-interactive` in the background, send the user the sign-in link it prints in a message, and wait for that command to finish. Only the user can sign in: don't look for API keys or passwords."},
		{name: "nothing known", want: "Not signed in. Run: leap auth login <url>\nAI assistant: start `leap auth login <url> --non-interactive` in the background, send the user the sign-in link it prints in a message, and wait for that command to finish. Only the user can sign in: don't look for API keys or passwords."},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			initTestConfig(t)
			viper.Set(API_URL_CONFIG_PATH, c.apiUrl)
			orig := installedServerUrl
			t.Cleanup(func() { installedServerUrl = orig })
			installedServerUrl = func() string { return c.installed }

			stdinIsTerminal = func() bool { return false }
			if got := NotSignedInError().Error(); got != c.want {
				t.Fatalf("no terminal: got %q, want %q", got, c.want)
			}
			// a person at a terminal gets the first line only
			stdinIsTerminal = func() bool { return true }
			if got, want := NotSignedInError().Error(), strings.SplitN(c.want, "\n", 2)[0]; got != want {
				t.Fatalf("terminal: got %q, want %q", got, want)
			}
		})
	}
}

// Only a key the server rejects means "not signed in"; a server that can't be reached says so instead
func TestCheckSignInTellsARejectedKeyFromAnUnreachableServer(t *testing.T) {
	initTestConfig(t)
	orig := installedServerUrl
	t.Cleanup(func() { installedServerUrl = orig })
	installedServerUrl = func() string { return "" }

	reply := func(status int, body string) *httptest.Server {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)
			_, _ = w.Write([]byte(body))
		}))
		t.Cleanup(ts.Close)
		return ts
	}
	ok := reply(http.StatusOK, `{"cid":"u1","recordSession":false,"role":"user","teamName":"t","activated":true,"local":{"email":"a@b.c","name":"a","teamId":"t1"}}`)
	rejected := reply(http.StatusUnauthorized, `{"message":"Invalid token"}`)
	forbidden := reply(http.StatusForbidden, `{"message":"user not authorized"}`)
	down := reply(http.StatusOK, `{}`)
	down.Close() // nothing listens there any more

	cases := []struct{ name, url, wantPrefix string }{
		{"accepted", ok.URL, ""},
		{"rejected key", rejected.URL, "Not signed in. Run: leap auth login " + rejected.URL},
		{"forbidden", forbidden.URL, "Not signed in. Run: leap auth login " + forbidden.URL},
		{"server down", down.URL, "Can't reach the Tensorleap server at " + down.URL + ": dial tcp"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			apiUrl := c.url + "/api/v2"
			viper.Set(API_URL_CONFIG_PATH, apiUrl)
			err := CheckSignIn(context.Background(), &Env{ApiUrl: apiUrl, ApiKey: "k"})
			if c.wantPrefix == "" {
				if err != nil {
					t.Fatal(err)
				}
				return
			}
			if err == nil || !strings.HasPrefix(err.Error(), c.wantPrefix) {
				t.Fatalf("got %v, want %q…", err, c.wantPrefix)
			}
		})
	}
}

func TestCheckSignInGivesUpOnAServerThatNeverAnswers(t *testing.T) {
	block := make(chan struct{})
	ts := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) { <-block }))
	t.Cleanup(func() { close(block); ts.Close() })
	orig := signInCheckTimeout
	t.Cleanup(func() { signInCheckTimeout = orig })
	signInCheckTimeout = 100 * time.Millisecond

	err := CheckSignIn(context.Background(), &Env{ApiUrl: ts.URL + "/api/v2", ApiKey: "k"})
	if err == nil || err.Error() != "Can't reach the Tensorleap server at "+ts.URL+": no answer in time" {
		t.Fatalf("got %v", err)
	}
}
