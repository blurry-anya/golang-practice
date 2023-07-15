package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re, _ := regexp.Compile(`^(\[ERR\]|\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[FTL\])`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, _ := regexp.Compile(`<(\*|\=|\~|\-)*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re, _ := regexp.Compile(`(?i)".*password"`)
	for _, v := range lines {
		if re.MatchString(v) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re, _ := regexp.Compile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re, _ := regexp.Compile(`User\s+\w+`)
	whitespaces := regexp.MustCompile(`\s+`)
	for i, v := range lines {
		if re.MatchString(v) {
			userName := whitespaces.Split(re.FindString(v), -1)[1]
			lines[i] = fmt.Sprintf("[USR] %v %v", userName, v)
		}
	}
	return lines
}
