package stringset

import (
	"fmt"
	"strings"
)

type Set map[string]bool

func New() Set {
	return Set{}
}

func NewFromSlice(l []string) Set {
	s := New()
	for _, text := range l {
		s[text] = true
	}
	return s
}

func (s Set) String() string {
	output := make([]string, 0, len(s)+1)
	for v := range s {
		output = append(output, fmt.Sprintf(`"%s"`, v))
	}

	return fmt.Sprintf("{%s}", strings.Join(output, ", "))
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s[elem]
	return ok
}

func (s Set) Add(elem string) {
	s[elem] = true
}

func Subset(s1, s2 Set) bool {
	if s1.IsEmpty() {
		return true
	}

	if len(s2) < len(s1) {
		return false
	}

	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	for k := range s1 {
		if s2.Has(k) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}

	if len(s1) == 0 {
		return true
	}

	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}

	return true
}

func Intersection(s1, s2 Set) Set {
	s := New()

	if len(s1) == 0 || len(s2) == 0 {
		return s
	}

	for k := range s1 {
		if s2.Has(k) {
			s[k] = true
		}
	}

	return s
}

func Difference(s1, s2 Set) Set {
	s := New()
	for k := range s1 {
		if !s2.Has(k) {
			s[k] = true
		}
	}
	return s
}

func Union(s1, s2 Set) Set {
	s := New()

	for k := range s1 {
		s[k] = true
	}

	for k := range s2 {
		s[k] = true
	}

	return s
}
