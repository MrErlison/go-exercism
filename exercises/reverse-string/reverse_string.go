package reverse

// Reverse returns a reversed string
func Reverse(input string) string {
	var reverse string
	for _, r := range input {
		reverse = string(r) + reverse
	}
	return reverse
}
