package version

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/tensorleap/leap-cli/pkg/api"
)

func TestServerUpgradeWarning(t *testing.T) {
	tests := []struct {
		name          string
		serverSchema  int
		minRequired   int
		wantEmpty     bool
		wantSubstring string
	}{
		{
			name:         "compatible (equal)",
			serverSchema: 5,
			minRequired:  5,
			wantEmpty:    true,
		},
		{
			name:         "compatible (newer)",
			serverSchema: 7,
			minRequired:  5,
			wantEmpty:    true,
		},
		{
			name:          "incompatible (older)",
			serverSchema:  3,
			minRequired:   5,
			wantSubstring: "leap server upgrade",
		},
		{
			name:          "incompatible mentions both versions",
			serverSchema:  2,
			minRequired:   10,
			wantSubstring: "schema v2",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := serverUpgradeWarning(tc.serverSchema, tc.minRequired)
			if tc.wantEmpty {
				if got != "" {
					t.Fatalf("expected empty warning, got %q", got)
				}
				return
			}
			if got == "" {
				t.Fatalf("expected non-empty warning for server=%d min=%d", tc.serverSchema, tc.minRequired)
			}
			if tc.wantSubstring != "" && !strings.Contains(got, tc.wantSubstring) {
				t.Fatalf("warning %q missing expected substring %q", got, tc.wantSubstring)
			}
		})
	}
}

func TestGetSkillsInfo(t *testing.T) {
	base := `"clientStoragePrefixUrl":"","schemaVersion":1,"disableDatadogMetrics":true,"isCloud":false,"gatewayUrl":"","proxyUrl":""`
	serve := func(body string) string {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(body))
		}))
		t.Cleanup(srv.Close)
		return srv.URL
	}
	level := func(url string) string {
		info := GetSkillsInfo(api.CreateAuthenticatedContext(context.Background(), "key", url))
		if info.SkillsRef == "" || info.SkillsRef != SkillsRef {
			t.Fatalf("skillsRef %q, want %q", info.SkillsRef, SkillsRef)
		}
		if info.ServerLevel == nil {
			return "nil"
		}
		return strconv.Itoa(*info.ServerLevel)
	}

	if got := level(serve("{" + base + `,"skillsLevel":3}`)); got != "3" {
		t.Errorf("server with skillsLevel 3: got %s", got)
	}
	if got := level(serve("{" + base + "}")); got != "0" {
		t.Errorf("server predating skillsLevel: got %s, want 0", got)
	}
	if got := level("http://127.0.0.1:1"); got != "nil" {
		t.Errorf("unreachable server: got %s, want nil", got)
	}
}
