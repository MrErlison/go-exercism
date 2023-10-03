package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	if pt == "" {
		return ""
	}

	var t []rune
	for _, v := range strings.ToLower(pt) {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			t = append(t, v)
		}
	}
	normalizedText := string(t)

	sqrt := math.Sqrt(float64(len(normalizedText)))
	cols := int(math.Round(sqrt))
	rows := cols
	if sqrt > math.Round(sqrt) {
		rows++
	}

	for i := rows*cols - len(normalizedText); i > 0; i-- {
		normalizedText += " "
	}

	var cipherText string
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cipherText += string(normalizedText[i+(j*rows)])
		}
		cipherText += " "
	}

	return cipherText[:len(cipherText)-1]
}
