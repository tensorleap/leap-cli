package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tensorleap/leap-cli/pkg/api"
)

func TestSetLicenseSavesTheRegisteredUserForAnalytics(t *testing.T) {
	cases := map[string]struct{ reply, want string }{
		"server names the user": {`{"success":true,"registeredUser":{"email":"a+b@c.d","id":"cid-1"}}`, "cid-1"},
		"older server":          {`{"success":true,"registeredUser":{"email":"a+b@c.d"}}`, ""},
		"no user registered":    {`{"success":true}`, ""},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/api/v2/auth/extendTrial" {
					t.Errorf("unexpected request %s", r.URL.Path)
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(c.reply))
			}))
			defer server.Close()
			var saved []string
			old := saveAnalyticsUserID
			t.Cleanup(func() { saveAnalyticsUserID = old })
			saveAnalyticsUserID = func(id string) error {
				saved = append(saved, id)
				return nil
			}

			ctx := api.CreateAuthenticatedContext(context.Background(), "", server.URL+"/api/v2")
			if err := SetLicense(ctx, "h.e30.s"); err != nil {
				t.Fatal(err)
			}
			if (c.want == "" && len(saved) != 0) || (c.want != "" && (len(saved) != 1 || saved[0] != c.want)) {
				t.Fatalf("saved %v, want %q", saved, c.want)
			}
		})
	}
}
