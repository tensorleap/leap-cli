package local

import (
	"strings"
)

func SanitizeDirectoryName(name string) string {
	// Replace characters that are not allowed in directory names
	replacer := strings.NewReplacer(
		"/", "_",
		"\\", "_",
		":", "_",
		"*", "_",
		"?", "_",
		"\"", "_",
		"<", "_",
		">", "_",
		"|", "_",
		" ", "_",
	)
	return replacer.Replace(name)
}
