package anagram

import (
	"reflect"
	"sort"
	"strings"
)

type letters []string

func (s letters) Len() int           { return len(s) }
func (s letters) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s letters) Less(i, j int) bool { return s[i] < s[j] }

func Detect(subject string, candidates []string) []string {
	var anagrams []string = []string{}

	rearrangement := func(word string) []string {
		w := letters(strings.Split(word, ""))
		sort.Sort(w)
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
