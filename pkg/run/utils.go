package run

import (
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
		return "", err
	}

	// Desired output: Mon, 08 Sep 2025 14:46
	return t.Format("Mon, 02 Jan 2006 15:04"), nil
}
