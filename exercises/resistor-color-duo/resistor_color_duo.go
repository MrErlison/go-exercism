package resistorcolorduo

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	var color_values = map[string]int{
		"black": 0, "brown": 1, "red": 2,
		"orange": 3, "yellow": 4, "green": 5,
		"blue": 6, "violet": 7, "grey": 8,
		"white": 9,
	}

	return color_values[colors[0]]*10 + color_values[colors[1]]
}