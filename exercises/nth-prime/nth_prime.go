package prime

import "errors"

func isPrime(n int) bool {
	for i := 2; i <= n; i++ {
		if n%i == 0 && n != i {
			return false
		}
	}

	return true
}

// Nth returns the nth prime number.
// An error must be returned if the nth prime number
// can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	var count, prime int = 1, 0

	if n <= 0 {
		return 0, errors.New("invalid number")
	}

	for i := 2; count <= n; i++ {
		if isPrime(i) {
			prime = i
			count++
		}
	}

	return prime, nil
}
