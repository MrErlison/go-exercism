package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	var frequency Frequency = Frequency{}

	phrase = strings.ToLower(phrase)
	regex := regexp.MustCompile(`[0-9a-z]+(\'[0-9a-z]+)?`)

	for _, w := range regex.FindAllString(phrase, -1) {
		frequency[w]++
	}

	return frequency
}
