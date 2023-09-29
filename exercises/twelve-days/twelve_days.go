package twelve

var gift = map[int][]string{
	1:  {"On the first day of Christmas my true love gave to me:", " a Partridge in a Pear Tree"},
	2:  {"On the second day of Christmas my true love gave to me:", " two Turtle Doves"},
	3:  {"On the third day of Christmas my true love gave to me:", " three French Hens"},
	4:  {"On the fourth day of Christmas my true love gave to me:", " four Calling Birds"},
	5:  {"On the fifth day of Christmas my true love gave to me:", " five Gold Rings"},
	6:  {"On the sixth day of Christmas my true love gave to me:", " six Geese-a-Laying"},
	7:  {"On the seventh day of Christmas my true love gave to me:", " seven Swans-a-Swimming"},
	8:  {"On the eighth day of Christmas my true love gave to me:", " eight Maids-a-Milking"},
	9:  {"On the ninth day of Christmas my true love gave to me:", " nine Ladies Dancing"},
	10: {"On the tenth day of Christmas my true love gave to me:", " ten Lords-a-Leaping"},
	11: {"On the eleventh day of Christmas my true love gave to me:", " eleven Pipers Piping"},
	12: {"On the twelfth day of Christmas my true love gave to me:", " twelve Drummers Drumming"},
}

// Verse returns one chosen verse from the song of Christmas
func Verse(i int) string {
	if i == 1 { // Only first day
		return gift[1][0] + gift[1][1] + "."
	}

	verse := gift[i][0]
	for ; i > 1; i-- {
		verse += gift[i][1] + ","
	}
	verse += " and" + gift[1][1] + "."

	return verse
}

// Song returns the song of Christmas with all gifts
func Song() string {
	var s string
	for i := 1; i < 12; i++ {
		s += Verse(i) + "\n"
	}
	s += Verse(12)

	return s
}
