package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {

	if span <= 0 || span > len(digits) {
		return 0, errors.New("invalid span")
	}

	if len(digits) == 0 {
		return 0, errors.New("invalid digits")
	}

	var bigProduct int64
	for i := span; i <= len(digits); i++ {
		var product int64 = 1
		for _, d := range digits[i-span : i] {
			digit, err := strconv.Atoi(string(d))
			if err != nil {
				return 0, errors.New("invalid digits")
			}
			product *= int64(digit)
		}
		if bigProduct < product {
			bigProduct = product
		}
	}

	return bigProduct, nil
}
