package isogram

import (
	"strings"
)

// IsIsogram returns if a word or phrase is an isogram
// Spaces and hyphens are ignored
func IsIsogram(word string) bool {
	word = strings.ToUpper(word)
	for k, char := range word {
		if strings.ContainsRune(word[k+1:], char) && char != ' ' && char != '-' {
			return false
		}
	}

	return true
}
