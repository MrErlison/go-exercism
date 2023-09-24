package series

// All returns a list of all substrings of s with length n
func All(n int, s string) []string {
	var result []string = []string{}
	for i := 0; i < len(s)-n+1; i++ {
		result = append(result, s[i:i+n])
	}

	return result
}

// UnsafeFirst returns the first substring of s with length n
func UnsafeFirst(n int, s string) string {
	if f, ok := First(n, s); ok {
		return f
	}

	panic(nil)
}

func First(n int, s string) (first string, ok bool) {
	if len(s) < n || n <= 0 {
		return "", false
	}

	return s[:n], true
}
