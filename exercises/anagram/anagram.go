package anagram

import (
	"reflect"
	"slices"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var anagrams []string = []string{}

	rearrangement := func(word string) []string {
		w := strings.Split(word, "")
		slices.Sort(w)
		return w
	}

	subject = strings.ToLower(subject)
	for k, candidate := range candidates {
		candidate = strings.ToLower(candidate)
		if len(candidate) != len(subject) || subject == candidate {
			continue
		}

		if reflect.DeepEqual(rearrangement(subject), rearrangement(candidate)) {
			anagrams = append(anagrams, candidates[k])
		}
	}

	return anagrams
}
