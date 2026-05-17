package version

import (
	"strings"
	"testing"
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
