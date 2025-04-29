package regex

import "regexp"

// / Returns a regex pattern that matches a version string
// / with an optional prefix (like "v") and a suffix (like "-beta").
// / The pattern is case-insensitive and allows for optional whitespace
// / around the version string.
func Version(suffixText string) *regexp.Regexp {
	return regexp.MustCompile(`(?i)(?:\s*|\s+|:)\s*(v?\d+\.\d+\.\d+)\s*` + suffixText)
}
