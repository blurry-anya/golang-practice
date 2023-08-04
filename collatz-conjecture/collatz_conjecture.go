package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {
	steps := 0

	if n <= 0 {
		return 0, errors.New("n must be positive")
	}

	for {
		if n == 1 {
			break
		}
		n = CalculateN(n)
		steps++
	}
	return steps, nil
}

func CalculateN(n int) int {
	if n%2 == 0 {
		n = n / 2
	} else {
		n = n*3 + 1
	}
	return n
}
