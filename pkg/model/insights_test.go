package model

import (
	"testing"

	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func TestHasAnyInsightsOverride(t *testing.T) {
	cases := []struct {
		name string
		o    *tensorleapapi.InsightsSettingsOverrides
		want bool
	}{
		{"nil", nil, false},
		{"empty", &tensorleapapi.InsightsSettingsOverrides{}, false},
		{
			"metric active",
			&tensorleapapi.InsightsSettingsOverrides{MetricActiveOverrides: map[string]bool{"m": false}},
			true,
		},
		{
			"metric direction",
			&tensorleapapi.InsightsSettingsOverrides{MetricDirectionOverrides: map[string]string{"m": "Upward"}},
			true,
		},
		{
			"ignored domains",
			&tensorleapapi.InsightsSettingsOverrides{IgnoredDomains: []string{"metadata.x"}},
			true,
		},
		{
			"added domains",
			&tensorleapapi.InsightsSettingsOverrides{AddedDomains: []string{"metadata.y"}},
			true,
		},
		{
			"latent space ignore",
			&tensorleapapi.InsightsSettingsOverrides{LatentSpaceIgnoreOverrides: map[string]bool{"balanced": false}},
			true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := HasAnyInsightsOverride(c.o); got != c.want {
				t.Fatalf("HasAnyInsightsOverride(%s) = %v, want %v", c.name, got, c.want)
			}
		})
	}
}
