package thefarm

import (
	"errors"
	"fmt"
)

// DivideFood divides the food evenly between the cows.
func DivideFood(f FodderCalculator, cows int) (float64, error) {
	totalFodder, amountError := f.FodderAmount(cows)
	if amountError != nil {
		return 0, amountError
	}
	factor, factorError := f.FatteningFactor()
	if factorError != nil {
		return 0, factorError
	}
	food := totalFodder * factor / float64(cows)
	return food, nil
}

// ValidateInputAndDivideFood checks the input value before calling DivideFood.
func ValidateInputAndDivideFood(f FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(f, cows)
}

type InvalidCowsError struct {
	cowsNumber int
	msg        string
}

// Error returns custom error InvalidCowsError.
func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %v", e.cowsNumber, e.msg)
}

// ValidateNumberOfCows accepts the number of cows as an integer and returns an error (or nil).
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows, "there are no negative cows"}
	} else if cows == 0 {
		return &InvalidCowsError{0, "no cows don't need food"}
	}
	return nil
}
