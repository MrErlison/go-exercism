package isbn

// IsValidISBN check if the provided string is a valid ISBN-10
func IsValidISBN(isbn string) bool {

	if len(isbn) < 10 {
		return false
	}

	var digits, check int
	for _, d := range isbn {
		switch {
		case d == '-':
			continue
		case d == 'X' && digits == 9:
			check += 10 * (10 - digits)
		case d >= '0' && d <= '9':
			check += int(d-'0') * (10 - digits)
		default:
			return false
		}
		digits++
	}

	return digits == 10 && check%11 == 0
}
