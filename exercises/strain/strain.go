package strain

// Keep returns a new collection containing
// those elements where the predicate is true.
func Keep[T any](list []T, f func(T) bool) []T {
	value := make([]T, 0, len(list))

	for _, item := range list {
		if !f(item) {
			continue
		}
		value = append(value, item)
	}
	return value
}

// Discard returns a new collection containing
// those elements where the predicate is false.
func Discard[T any](list []T, f func(T) bool) []T {
	return Keep(list, func(t T) bool { return !f(t) })
}
