package model

import (
	"context"
	"fmt"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

// GetLatestProjectVersion returns the most recently created version of the
// project (by CreatedAt), or nil if the project has no versions yet.
func GetLatestProjectVersion(ctx context.Context, projectId string) (*tensorleapapi.SlimVersion, error) {
	versions, err := GetVersions(ctx, projectId)
	if err != nil {
		return nil, err
	}
	var latest *tensorleapapi.SlimVersion
	for i := range versions {
		v := &versions[i]
		if latest == nil || v.CreatedAt.After(latest.CreatedAt) {
			latest = v
		}
	}
	return latest, nil
}

// GetInsightsSettingsOverrides fetches the per-version insight-settings
// overrides for a version. An error (including a missing endpoint) is returned
// to the caller rather than swallowed.
func GetInsightsSettingsOverrides(ctx context.Context, projectId, versionId string) (*tensorleapapi.InsightsSettingsOverrides, error) {
	resp, httpRes, err := api.ApiClient.GetInsightsSettings(ctx).
		GetInsightsSettingsParams(*tensorleapapi.NewGetInsightsSettingsParams(projectId, versionId)).
		Execute()
	if err := api.CheckRes(httpRes, err); err != nil {
		return nil, fmt.Errorf("failed to fetch insight settings: %w", err)
	}
	return &resp.Overrides, nil
}

// HasAnyInsightsOverride reports whether any insight setting has been
// customized away from its default.
func HasAnyInsightsOverride(o *tensorleapapi.InsightsSettingsOverrides) bool {
	if o == nil {
		return false
	}
	return len(o.MetricActiveOverrides) > 0 ||
		len(o.MetricDirectionOverrides) > 0 ||
		len(o.IgnoredDomains) > 0 ||
		len(o.AddedDomains) > 0 ||
		len(o.LatentSpaceIgnoreOverrides) > 0
}
