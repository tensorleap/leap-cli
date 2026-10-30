package server

import (
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// The first install adds the user to the tensorleap group and re-runs itself under sg; that second run
// must not count another start, but its result is the only one sent.
func TestReexecUnderSharedGroupSendsTheStartOnce(t *testing.T) {
	commands := map[string]func() *cobra.Command{"install": NewInstallCmd, "reinstall": NewReinstallCmd}
	for name, newCmd := range commands {
		for _, failAt := range []string{"", "init_data_dir", "run"} {
			t.Run(fmt.Sprintf("%s/fail=%q", name, failAt), func(t *testing.T) {
				first, _ := runWithStubs(t, newCmd, failAt)
				t.Setenv(sharedGroupReexecEnv, "1")
				again, _ := runWithStubs(t, newCmd, failAt)

				if len(first) != 2 || !strings.HasSuffix(string(first[0].Event), "_started") {
					t.Fatalf("a normal run sends started and a result: %v", first)
				}
				if len(again) != 1 || again[0].Event != first[1].Event {
					t.Fatalf("the re-executed run sends only the result %s, got %v", first[1].Event, again)
				}
			})
		}
	}
}
