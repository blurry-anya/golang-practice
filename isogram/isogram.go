package isogram

import (
	"regexp"
	"strings"
)

func IsIsogram(word string) bool {

	word = strings.ToLower(word)
	justWordSymbolsRegex := regexp.MustCompile(`[^a-z]+`)

	word = justWordSymbolsRegex.ReplaceAllString(word, "")

	for i, v := range word {
		for j := i + 1; j < len(word); j++ {
			if v == rune(word[j]) {
				return false
			}
		}
	}

	return true
}
