package prime

// Factors returns all prime factors of a number
func Factors(n int64) []int64 {
	var factors []int64 = []int64{}
	var d int64 = 2

	for n > 1 {
		if n%d == 0 {
			factors = append(factors, d)
			n /= d
		} else {
			d++
		}
	}

	return factors
}
