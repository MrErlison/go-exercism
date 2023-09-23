package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	var sum float64
	power := float64(len(strconv.Itoa(n)))
	for x := n; x > 0; x /= 10 {
		sum += math.Pow(float64(x%10), power)
	}

	return int(sum) == n
}
