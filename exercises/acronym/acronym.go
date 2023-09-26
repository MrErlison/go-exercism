package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate converts a long name to acronym
func Abbreviate(s string) string {
	s = strings.ToUpper(s)
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '\''
	}

	var acronym string
	for _, v := range strings.FieldsFunc(s, f) {
		acronym += string(v[0])
	}

	return acronym
}
