package auth

import (
	"fmt"
	"strings"

	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

const LicenseContact = "support@tensorleap.ai"
const FreeTierMaxEvaluateSamples = 5000
const freeTierDefaultMaxProjects = 1

// LicenseActivated is false only for a colored (not yet activated) license; the
// generated client keeps the server's `activated` field in AdditionalProperties.
func LicenseActivated(lic *tensorleapapi.LicenseMetadata) bool {
	v, ok := lic.AdditionalProperties["activated"].(bool)
	return !ok || v
}

// LicenseLine renders the server's license for `leap auth whoami`; "" when the
// server reports no license.
func LicenseLine(info *tensorleapapi.GetEnvironmentInfoResponse) string {
	if info == nil {
		return ""
	}
	lic, ok := info.GetLicenseMetadataOk()
	if !ok || lic == nil {
		return ""
	}
	maxProjects := licenseMaxProjects(lic)
	notActivated := ""
	if !LicenseActivated(lic) {
		notActivated = ", not activated"
	}
	if lic.GetType() == tensorleapapi.LICENSETYPE_FREE {
		if maxProjects == 0 {
			maxProjects = freeTierDefaultMaxProjects
		}
		return fmt.Sprintf(
			"License: Free Tier%s, %s, evaluate up to %s samples, no dataset curation or sample analysis; contact %s to unlock the full platform",
			notActivated, projectsPhrase(maxProjects), formatThousands(FreeTierMaxEvaluateSamples), LicenseContact,
		)
	}
	parts := []string{string(lic.GetType()) + notActivated}
	if exp, ok := lic.GetExpirationDateOk(); ok && exp != nil {
		parts = append(parts, "expires "+exp.Format("2006-01-02"))
	}
	if maxProjects > 0 {
		parts = append(parts, projectsPhrase(maxProjects))
	}
	return "License: " + strings.Join(parts, ", ")
}

// maxProjects reaches the CLI through the generated model's additional
// properties until the Go client is regenerated from a server that has it.
func licenseMaxProjects(lic *tensorleapapi.LicenseMetadata) int {
	if v, ok := lic.AdditionalProperties["maxProjects"]; ok {
		if f, ok := v.(float64); ok && f > 0 {
			return int(f)
		}
	}
	return 0
}

func projectsPhrase(n int) string {
	if n == 1 {
		return "1 project"
	}
	return fmt.Sprintf("%d projects", n)
}

func formatThousands(n int) string {
	s := fmt.Sprintf("%d", n)
	for i := len(s) - 3; i > 0; i -= 3 {
		s = s[:i] + "," + s[i:]
	}
	return s
}
