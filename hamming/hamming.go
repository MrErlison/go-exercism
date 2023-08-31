package hamming

import "errors"

// Distance returns the Hamming distance between two values
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("samples must have the same length")
	}

	var distance int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
