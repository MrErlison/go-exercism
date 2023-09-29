package bottlesong

import (
	"fmt"
	"strings"
)

var n map[int]string = map[int]string{
	0:  "no",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
}

func Recite(startBottles, takeDown int) []string {
	title := func(s string) string { return strings.ToUpper(s[:1]) + s[1:] }
	bottles := func(n int) string {
		if n == 1 {
			return "bottle"
		}
		return "bottles"
	}

	song := []string{}
	bridge := "And if one green bottle should accidentally fall,"
	for i := startBottles; i > startBottles-takeDown; i-- {
		refrain := fmt.Sprintf("%s green %s hanging on the wall,", title(n[i]), bottles(i))
		chorus := fmt.Sprintf("There'll be %s green %s hanging on the wall.", n[i-1], bottles(i-1))
		song = append(song, refrain, refrain, bridge, chorus, "")
	}

	return song[:len(song)-1]
}
