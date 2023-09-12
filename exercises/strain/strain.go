package strain

// Keep returns a new collection containing
// those elements where the predicate is true.
func Keep[T any](list []T, f func(T) bool) []T {
	var result []T = make([]T, 0, len(list))
	for _, v := range list {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

// Discard returns a new collection containing
// those elements where the predicate is false.
func Discard[T any](list []T, f func(T) bool) []T {
	return Keep(list, func(t T) bool { return !f(t) })
}
