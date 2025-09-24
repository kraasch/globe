// Package testutil helps in the tests in this projects.
package testutil

import (
	"regexp"
)

func Anonymize(input string) string {
	input = AnonymizeTimeStrings(input)
	input = AnonymizeDateStrings(input)
	input = AnonymizeFloatStrings(input)
	input = AnonymizeWeekdayStrings(input)
	return input
}

func AnonymizeWeekdayStrings(input string) string {
	// Compile regex to match weekday abbreviations (Mon, Tue, Wed, Thu, Fri, Sat, Sun)
	pattern := `\b(?:Mon|Tue|Wed|Thu|Fri|Sat|Sun)\b`
	re := regexp.MustCompile(pattern)
	// Replace all occurrences with "weekday"
	result := re.ReplaceAllString(input, "aaa")
	return result
}

func AnonymizeFloatStrings(input string) string {
	// Compile regex for floats:
	// Pattern matches:
	// 1) [0-9]+\.[0-9]*  (e.g., 123.456, 0.5)
	// 2) [0-9]*\.[0-9]+  (e.g., .456, 0.456)
	pattern := `\b(?:[0-9]+\.[0-9]*|[0-9]*\.[0-9]+)\b`
	re := regexp.MustCompile(pattern)
	// Replace all float occurrences with "float"
	result := re.ReplaceAllString(input, "ff.ff")
	return result
}

func AnonymizeDateStrings(input string) string {
	// Compile a regex pattern to match dates in format YYYY-MM-DD
	pattern := `\d{4}-\d{2}-\d{2}`
	re := regexp.MustCompile(pattern)
	// Replace all occurrences with "yyyy-mm-dd"
	result := re.ReplaceAllString(input, "yyyy-mm-dd")
	return result
}

// AnonymizeTimeStrings takes a string and replaces times like 07:32 or 23:59 into xx:yy for time-independent string tests.
func AnonymizeTimeStrings(timestr string) string {
	pattern := `(?:[01][0-9]|2[0-3]):[0-5][0-9]` // more restrictive pattern.
	re := regexp.MustCompile(pattern)
	result := re.ReplaceAllString(timestr, "HH:MM")
	return result
}
