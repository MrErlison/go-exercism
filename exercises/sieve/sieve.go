package sieve

func Sieve(limit int) []int {
	var primes []int = make([]int, limit+1)
	for i := 2; i*i <= limit; i++ {
		if primes[i] == -1 {
			continue
		}

		for j := i + 1; j <= limit; j++ {
			if j%i == 0 {
				primes[j] = -1
			}
		}
	}

	var result []int = []int{}
	for k, v := range primes {
		if v >= 0 && k >= 2 {
			result = append(result, k)
		}
	}

	return result
}
