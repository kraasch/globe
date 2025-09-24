// Package testutil helps in the tests in this projects.
package testutil

import (
	"regexp"
)

// AnonymizeTimeStrings takes a string and replaces times like 07:32 or 24:20 into xx:yy for time-independent string tests.
func AnonymizeTimeStrings(timestr string) string {
	pattern := `[0-9]{2}:[0-9]{2}`
	re := regexp.MustCompile(pattern)
	result := re.ReplaceAllString(timestr, "xx:yy")
	return result
}
