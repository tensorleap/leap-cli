package run

import (
	"fmt"
	"strings"
	"time"
)

func ParseAndFormatDate(raw string) (string, error) {
	// Trim the (Coordinated Universal Time) part
	clean := strings.SplitN(raw, " (", 2)[0]

	// JS-style format: Mon Sep 08 2025 14:46:56 GMT+0000
	const layoutIn = "Mon Jan 02 2006 15:04:05 GMT-0700"

	t, err := time.Parse(layoutIn, clean)
	if err != nil {
		return "", fmt.Errorf("failed to parse time %q: %w", clean, err)
	}

	// Convert to local time zone
	localTime := t.Local()

	// Desired output: Mon, 08 Sep 2025 14:46 (local)
	return localTime.Format("Mon, 02 Jan 2006 15:04"), nil
}
