package thefarm

import (
	"errors"
	"fmt"
)

// DivideFood function
func DivideFood(fc FodderCalculator, NumberOfCows int) (float64, error) {
	amount, err := fc.FodderAmount(NumberOfCows)
	if err != nil {
		return 0, err
	}

	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return amount * factor / float64(NumberOfCows), nil
}

// ValidateInputAndDivideFood function
func ValidateInputAndDivideFood(fc FodderCalculator, NumberOfCows int) (float64, error) {
	if NumberOfCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(fc, NumberOfCows)
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

// ValidateNumberOfCows function
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows, "there are no negative cows"}
	}

	if cows == 0 {
		return &InvalidCowsError{cows, "no cows don't need food"}
	}

	return nil
}
