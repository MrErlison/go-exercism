package cipher

import (
	"strings"
	"unicode"
)

type shift int

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance <= -26 || distance >= 26 || distance == 0 {
		return nil
	}

	return shift(distance)
}

func (c shift) Encode(input string) string {
	var output string

	for _, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			output += shiftTo(r, c)
		}
	}

	return output
}

func (c shift) Decode(input string) string {
	return shift(-c).Encode(input)
}

type vigenere []rune

func NewVigenere(key string) Cipher {
	if len(key) < 4 {
		return nil
	}
	for _, r := range key {
		if !unicode.IsLetter(r) {
			return nil
		}
	}

	return vigenere(key)
}

func (v vigenere) Encode(input string) string {
	var cleaned string
	for _, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			cleaned += string(r)
		}
	}

	var encoded string
	var l int = len(v)
	for i, r := range cleaned {
		encoded += shiftTo(r, shift(v[i%l]-'a'))
	}

	return encoded
}

func (v vigenere) Decode(input string) string {
	var decoded string
	var l int = len(v)
	for i, r := range input {
		decoded += shiftTo(r, -shift(v[i%l]-'a'))
	}

	return decoded
}

func shiftTo(r rune, s shift) string {
	var shifted rune = r + rune(s)
	switch {
	case shifted < 'a':
		shifted += 26
	case shifted > 'z':
		shifted -= 26
	}

	return string(shifted)
}
