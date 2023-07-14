package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// logLine represents a table of the characters in application
var logLine map[string]string = map[string]string{
	"U+2757":  "recommendation",
	"U+1F50D": "search",
	"U+2600":  "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, v := range log {
		line, exists := logLine[fmt.Sprintf("%U", v)]
		if exists {
			return line
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
