package atbash

func Atbash(s string) string {
	var cipher string
	var count int

	for _, v := range s {
		switch {
		case 'a' <= v && v <= 'z':
			v = 'z' - v + 'a'
		case 'A' <= v && v <= 'Z':
			v = 'Z' - v + 'a'
		case '0' <= v && v <= '9':
			v = rune(v)
		default:
			continue
		}
		if count == 5 {
			cipher += " "
			count = 0
		}
		cipher += string(v)
		count++
	}
	return cipher
}
