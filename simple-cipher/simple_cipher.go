package cipher

import (
	"math"
	"math/rand"
	"strings"
	"unicode"
)

type shift struct {
	distance int
}

type vigenere struct {
	key string
}

var alphabet = [26]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func NewCaesar() Cipher {
	c := shift{
		distance: 3,
	}
	return c
}

func NewShift(distance int) Cipher {
	if math.Abs(float64(distance)) > 25 || math.Abs(float64(distance)) < 1 {
		return nil
	}
	s := shift{
		distance: distance,
	}
	return s
}

func (c shift) Encode(input string) string {
	input = strings.ToLower(input)

	cleanInput := []rune{}
	for _, value := range input {
		if unicode.IsLetter(value) {
			cleanInput = append(cleanInput, value)
		}
	}

	encoded := []rune{}
	maxIndex := len(alphabet) - 1
	for _, value := range cleanInput {
		for index, letter := range alphabet {
			if letter == value {
				shiftIndex := index + c.distance

				if shiftIndex != maxIndex {
					shiftIndex %= len(alphabet)
				}

				if shiftIndex < 0 {
					shiftIndex += len(alphabet)
				}

				encoded = append(encoded, alphabet[shiftIndex])
			}
		}
	}

	return string(encoded)
}

func (c shift) Decode(input string) string {
	input = strings.ToLower(input)

	cleanInput := []rune{}
	for _, value := range input {
		if unicode.IsLetter(value) {
			cleanInput = append(cleanInput, value)
		}
	}

	decoded := []rune{}
	maxIndex := len(alphabet) - 1
	for _, value := range cleanInput {
		for index, letter := range alphabet {
			if letter == value {
				shiftIndex := index - c.distance

				if shiftIndex != maxIndex {
					shiftIndex %= len(alphabet)
				}

				if shiftIndex < 0 {
					shiftIndex += len(alphabet)
				}

				decoded = append(decoded, alphabet[shiftIndex])
			}
		}
	}

	return string(decoded)
}

func getRandomString(length int) string {
	b := []rune{}

	i := 0
	for i < length {
		b = append(b, alphabet[rand.Intn(len(alphabet))])
		i++
	}

	return string(b)
}

func NewVigenere(key string) Cipher {

	if len(key) == 0 {
		return nil
	}

	letterACount := 0
	digitFlag := false
	upperFlag := false
	for _, value := range key {
		if !unicode.IsLetter(value) {
			digitFlag = true
		}
		if unicode.IsUpper(value) {
			upperFlag = true
		}
		if value == 'a' {
			letterACount++
		}
	}
	if letterACount == len(key) || digitFlag || upperFlag {
		return nil
	}

	v := vigenere{
		key: key,
	}
	return v
}

func getIndexInSlice(search rune, slice []rune) int {
	for index, value := range slice {
		if value == search {
			return index
		}
	}
	return -1
}

func (v vigenere) Encode(input string) string {
	input = strings.ToLower(input)

	cleanInput := []rune{}
	for _, value := range input {
		if unicode.IsLetter(value) {
			cleanInput = append(cleanInput, value)
		}
	}

	encodingKey := v.key
	if len(encodingKey) == 0 {
		encodingKey = getRandomString(100)
	}
	if len(encodingKey) < len(cleanInput) {
		isTooShort := true
		for isTooShort {
			if len(encodingKey) >= len(cleanInput) {
				isTooShort = false
			}
			encodingKey += encodingKey
		}
	}

	encoded := []rune{}
	for i, value := range cleanInput {
		currentShift := getIndexInSlice(rune(encodingKey[i]), alphabet[:])

		currentLetter := getIndexInSlice(value, alphabet[:])

		shiftIndex := currentLetter + currentShift

		if shiftIndex != len(alphabet)-1 {
			shiftIndex %= len(alphabet)
		}

		if shiftIndex < 0 {
			shiftIndex += len(alphabet)
		}

		encoded = append(encoded, alphabet[shiftIndex])
	}

	return string(encoded)
}

func (v vigenere) Decode(input string) string {
	input = strings.ToLower(input)

	cleanInput := []rune{}
	for _, value := range input {
		if unicode.IsLetter(value) {
			cleanInput = append(cleanInput, value)
		}
	}

	decodingKey := v.key
	if len(decodingKey) < len(cleanInput) {
		isTooShort := true
		for isTooShort {
			if len(decodingKey) >= len(cleanInput) {
				isTooShort = false
			}
			decodingKey += decodingKey
		}
	}

	decoded := []rune{}
	for i, value := range input {
		if i >= len(input) {
			break
		}
		currentShift := getIndexInSlice(rune(decodingKey[i]), alphabet[:])

		currentLetter := getIndexInSlice(value, alphabet[:])

		shiftIndex := currentLetter - currentShift

		if shiftIndex != len(alphabet)-1 {
			shiftIndex %= len(alphabet)
		}

		if shiftIndex < 0 {
			shiftIndex += len(alphabet)
		}

		decoded = append(decoded, alphabet[shiftIndex])
	}

	return string(decoded)
}
