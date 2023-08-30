package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return -1, errors.New("the number must be a positive")
	}

	var steps int
	for ; n > 1; steps++ {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
	}

	return steps, nil
}
