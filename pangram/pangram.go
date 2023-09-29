package pangram

import (
	"regexp"
	"strings"
)

const abc = "abcdefghijklmnopqrstuvwxyz"

func IsPangram(input string) bool {
	notWordSymbolsRegex := regexp.MustCompile(`\W+`)
	input = notWordSymbolsRegex.ReplaceAllString(input, "")

	letterMap := make(map[rune]bool)

	input = strings.ToLower(input)

	for _, char := range input {
		if 'a' <= char && char <= 'z' {
			letterMap[char] = true
		}
	}

	return len(letterMap) == len(abc)

}
