package version

import (
	"context"
	"fmt"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/log"
)

// MinServerSchemaVersion is the lowest server schema version this CLI build is
// compatible with. Bump it together with any CLI change that depends on server
// behavior that isn't available on older servers. The value is compared to
// GetEnvironmentInfo.schemaVersion at runtime.
const MinServerSchemaVersion = 1

// CheckServerCompatibility probes the connected server and logs a warning if
// its schema version is older than what this CLI requires. It is intentionally
// best-effort: any failure to reach the server or interpret the response is
// swallowed so the user's actual command can proceed and surface its own error.
func CheckServerCompatibility(ctx context.Context) {
	envInfo, _, err := api.ApiClient.GetEnvironmentInfo(ctx).Execute()
	if err != nil || envInfo == nil {
		return
	}
	if msg := serverUpgradeWarning(int(envInfo.SchemaVersion), MinServerSchemaVersion); msg != "" {
		log.Warn(msg)
	}
}

// serverUpgradeWarning returns a user-facing warning when the server schema is
// below the required minimum. Returns "" when the server is compatible.
func serverUpgradeWarning(serverSchema, minRequired int) string {
	if serverSchema >= minRequired {
		return ""
	}
	return fmt.Sprintf(
		"Your Tensorleap server (schema v%d) is older than this CLI requires (>= v%d). "+
			"Commands may fail due to a version mismatch. "+
			"To upgrade the server, run 'leap server upgrade' on the server host.",
		serverSchema, minRequired,
	)
}
