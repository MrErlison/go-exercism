package rotationalcipher

import "unicode"

// RotationalCipher returns a simple shift cipher that relies on transposing
// all the letters in the alphabet using an integer key between 0 and 26.
func RotationalCipher(plain string, shiftKey int) string {
	var result string
	for _, r := range plain {
		switch {
		case unicode.IsLower(r):
			result += string(((r-'a')+rune(shiftKey))%26 + 'a')
		case unicode.IsUpper(r):
			result += string(((r-'A')+rune(shiftKey))%26 + 'A')
		default: // !unicode.IsLetter(r)
			result += string(r)
		}
	}
	return result
}
