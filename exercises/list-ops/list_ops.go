package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

// // given a function, a list, and initial accumulator,
// fold (reduce) each item into the accumulator from the left
func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
		initial = fn(initial, v)
	}

	return initial
}

// given a function, a list, and an initial accumulator,
// fold (reduce) each item into the accumulator from the right
func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := s.Length() - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}

	return initial
}

// given a predicate and a list, return the list of all items for which predicate(item) is True)
func (s IntList) Filter(fn func(int) bool) IntList {
	var r IntList = IntList{}
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}

	return r
}

// Length returns the total number of items within it
func (s IntList) Length() int {
	return len(s)
}

// Map returns the list of the results of applying function(item) on all items)
func (s IntList) Map(fn func(int) int) IntList {
	var r IntList = IntList{}
	for _, v := range s {
		r = append(r, fn(v))
	}

	return r
}

// Reverse returns a list with all the original items, but in reversed order
func (s IntList) Reverse() IntList {
	var r []int = []int{}
	for i := s.Length() - 1; i >= 0; i-- {
		r = append(r, s[i])
	}

	return r
}

// given two lists, add all items in the second list to the end of the first list
func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

// given a series of lists, combine all items in all lists into one flattened list
func (s IntList) Concat(lists []IntList) IntList {
	for _, v := range lists {
		s = s.Append(v)
	}

	return s
}
