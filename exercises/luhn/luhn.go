package luhn

import "strings"

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	var total, duplicate int
	for i := len(id) - 1; i >= 0; i-- {
		num := int(id[i] - '0')

		if num < 0 || num > 9 {
			return false
		}

		if duplicate%2 != 0 {
			num <<= 1
			if num > 9 {
				num -= 9
			}
		}

		total += num
		duplicate++
	}

	return total%10 == 0 && len(id) > 1
}
