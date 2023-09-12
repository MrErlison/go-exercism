package romannumerals

import (
	"errors"
)

// ToRomanNumeral convert from normal numbers to Roman Numerals
func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", errors.New("out of range")
	}

	romanNumerals := map[int][]string{
		0: {"", "", "", ""},
		1: {"I", "X", "C", "M"},
		2: {"II", "XX", "CC", "MM"},
		3: {"III", "XXX", "CCC", "MMM"},
		4: {"IV", "XL", "CD"},
		5: {"V", "L", "D"},
		6: {"VI", "LX", "DC"},
		7: {"VII", "LXX", "DCC"},
		8: {"VIII", "LXXX", "DCCC"},
		9: {"IX", "XC", "CM"}}

	var result string
	for i := 0; i < 4; i++ {
		digit := input % 10
		result = romanNumerals[digit][i] + result
		input /= 10
	}

	return result, nil
}
