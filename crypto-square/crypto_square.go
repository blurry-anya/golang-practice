package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

func Encode(pt string) string {
	pt = normalizeText(pt)

	if len(pt) <= 1 {
		return pt
	}

	rows, cols := getDimensions(pt)

	rectangle := getRectangle(pt, rows, cols)

	encodedRectangle := [][]rune{}

	for j := 0; j < len(rectangle[0]); j++ {
		row := []rune{}

		for i := 0; i < len(rectangle); i++ {
			row = append(row, rectangle[i][j])
		}

		encodedRectangle = append(encodedRectangle, row)
	}

	encodedMessage := ""
	for i, row := range encodedRectangle {
		encodedMessage += string(row)
		if i < len(encodedRectangle)-1 {
			encodedMessage += " "
		}
	}

	return encodedMessage
}

func getRectangle(text string, rows, cols int) [][]rune {
	rectangle := [][]rune{}

	textLen := len([]rune(text))

	for i := 0; i < textLen; i += cols {
		start := i
		end := i + cols - 1

		if end >= textLen {
			end = textLen - 1
		}

		chunk := text[start : end+1]
		populateWithSpaces(&chunk, cols)

		rectangle = append(rectangle, []rune(chunk))
	}

	return rectangle
}

func populateWithSpaces(word *string, mustLen int) {
	if len(*word) < mustLen {
		dif := mustLen - len(*word)
		for i := 0; i < dif; i++ {
			*word += " "
		}
	}
}

func getDimensions(text string) (rows, cols int) {
	textLength := len([]rune(text))

	r := int(math.Floor(math.Sqrt(float64(textLength))))
	c := r

	if c*r < textLength {
		c++
	}

	return r, c
}

func normalizeText(text string) string {
	text = strings.ToLower(text)

	justWordSymbolsRegex := regexp.MustCompile(`[^a-z1-9]+`)

	text = justWordSymbolsRegex.ReplaceAllString(text, "")

	return text
}
