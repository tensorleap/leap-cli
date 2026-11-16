package version

import (
	"context"
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/log"
)

//go:embed skills-ref
var skillsRefFile string

// SkillsRef is the tensorleap/skills git tag this CLI ships with: skills install and update to exactly
// this ref, so a skill never runs ahead of the CLI and server it was released with.
var SkillsRef = strings.TrimSpace(skillsRefFile)

type SkillsInfo struct {
	Cli         string `json:"cli"`
	ServerLevel *int   `json:"serverLevel"`
	SkillsRef   string `json:"skillsRef"`
}

// GetSkillsInfo leaves ServerLevel nil when no server answers (not signed in, unreachable),
// so the skills skip the server check. A server that answers without skillsLevel predates
// it and counts as level 0.
func GetSkillsInfo(ctx context.Context) SkillsInfo {
	info := SkillsInfo{Cli: CliVersion, SkillsRef: SkillsRef}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	envInfo, _, err := api.ApiClient.GetEnvironmentInfo(ctx).Execute()
	if err != nil || envInfo == nil {
		return info
	}
	level := int(envInfo.GetSkillsLevel())
	info.ServerLevel = &level
	return info
}

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
