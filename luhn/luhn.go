package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

func Valid(id string) bool {

	nonNumbersAndSpacesRegex := regexp.MustCompile(`[^0-9 ]+`)

	if hasNonNumbersAndSpaces := nonNumbersAndSpacesRegex.MatchString(id); hasNonNumbersAndSpaces {
		return false
	}

	id = strings.ReplaceAll(id, " ", "")

	if len(id) <= 1 {
		return false
	}

	idArray := getIntFromString(id)

	coef := 9

	for i := len(idArray) - 2; i >= 0; i -= 2 {
		if doubled := idArray[i] * 2; doubled > coef {
			idArray[i] = doubled - coef
		} else {
			idArray[i] = doubled
		}
	}

	return sumOfInts(idArray...)%10 == 0
}

func getIntFromString(input string) []int {
	res := []int{}

	for _, v := range input {
		intValue, _ := strconv.Atoi(string(v))
		res = append(res, intValue)
	}

	return res
}

func sumOfInts(n ...int) int {
	count := 0

	for _, v := range n {
		count += v
	}

	return count
}
