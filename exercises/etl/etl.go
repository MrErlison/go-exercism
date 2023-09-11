package etl

import "strings"

// Transform returns individual letter with its score in a one-to-one
func Transform(in map[int][]string) map[string]int {
	var oneToOne map[string]int = map[string]int{}

	for point, letters := range in {
		for _, letter := range letters {
			oneToOne[strings.ToLower(letter)] = point
		}
	}
	return oneToOne
}
