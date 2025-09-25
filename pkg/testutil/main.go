// Package testutil helps in the tests in this projects.
package testutil

import (
	"regexp"
)

func Anonymize(input string) string {
	input = AnonymizeTimeStrings(input)
	input = AnonymizeDateStrings(input)
	input = AnonymizeShortDateStrings(input)
	input = AnonymizeFloatStrings(input)
	input = AnonymizeNumberStrings(input)
	input = AnonymizeWeekdayStrings(input)
	return input
}

func AnonymizeWeekdayStrings(input string) string {
	// Compile regex to match weekday abbreviations (Mon, Tue, Wed, Thu, Fri, Sat, Sun)
	pattern := `\b(?:Mon|Tue|Wed|Thu|Fri|Sat|Sun)\b`
	re := regexp.MustCompile(pattern)
	result := re.ReplaceAllString(input, "aaa")
	return result
}

func AnonymizeNumberStrings(input string) string {
	pattern := `[0-9]+`
	re := regexp.MustCompile(pattern)
	result := re.ReplaceAllString(input, "NNNN")
	return result
}

func AnonymizeFloatStrings(input string) string {
	// Compile regex for any floats with number before AND after the decimal point.
	pattern := `[0-9]+\.[0-9]+`
	re := regexp.MustCompile(pattern)
	result := re.ReplaceAllString(input, "ff.ff")
	return result
}

func AnonymizeDateStrings(input string) string {
	// Compile a regex pattern to match dates in format YYYY-MM-DD
	pattern := `\d{4}-\d{2}-\d{2}`
	re := regexp.MustCompile(pattern)
	result := re.ReplaceAllString(input, "yyyy-mm-dd")
	return result
}

func AnonymizeShortDateStrings(input string) string {
	// Compile a regex pattern to match dates in format YY-MM-DD
	pattern := `\d{2}-\d{2}-\d{2}`
	re := regexp.MustCompile(pattern)
	result := re.ReplaceAllString(input, "yy-mm-dd")
	return result
}

// AnonymizeTimeStrings takes a string and replaces times like 07:32 or 23:59 into xx:yy for time-independent string tests.
func AnonymizeTimeStrings(timestr string) string {
	pattern := `(?:[01][0-9]|2[0-3]):[0-5][0-9]` // more restrictive pattern.
	re := regexp.MustCompile(pattern)
	result := re.ReplaceAllString(timestr, "HH:MM")
	return result
}
