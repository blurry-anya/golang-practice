package bob

import (
	"regexp"
	"strings"
	"unicode"
)

// Hey produces an answer to remark.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	if strings.HasSuffix(remark, "?") {
		if !isAllUpperCase(remark) {
			return "Sure."
		}
	}

	if isAllUpperCase(remark) {
		if strings.HasSuffix(remark, "?") {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}

	return "Whatever."

}

// isAllUpperCase checks if all the letters in string are upper cased.
func isAllUpperCase(remark string) bool {

	remark = removeSpecialChars(remark)

	if remark == "" {
		return false
	}

	for _, char := range remark {
		if !unicode.IsUpper(char) {
			return false
		}
	}

	return true
}

// removeSpecialChars removes special characters from a string.
func removeSpecialChars(remark string) string {
	specialCharacters := regexp.MustCompile(`[@#$%^'&*?!,()0-9 ]`)

	return specialCharacters.ReplaceAllString(remark, "")
}
