package grains

import (
	"errors"
)

func Square(number int) (uint64, error) {
	var result uint64 = 1

	if number <= 0 || number > 64 {
		return 0, errors.New("invalid number")
	}

	if number == 1 {
		return result, nil
	}

	for i := 2; i <= number; i++ {
		result = uint64(result * 2)
	}

	return result, nil
}

func Total() uint64 {
	squaresAmount := 64
	var total uint64 = 0

	for i := 1; i <= squaresAmount; i++ {
		currentSquare, _ := Square(i)
		total += currentSquare
	}

	return total
}
