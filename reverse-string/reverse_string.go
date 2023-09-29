package reverse

func Reverse(input string) string {
	inputRunes := []rune(input)
	length := len(inputRunes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		inputRunes[i], inputRunes[j] = inputRunes[j], inputRunes[i]
	}

	return string(inputRunes)
}
