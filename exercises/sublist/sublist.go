package sublist

import (
	"strconv"
	"strings"
)

func Sublist(l1, l2 []int) Relation {
	l := func(list []int) (l string) {
		for _, v := range list {
			l += strconv.Itoa(v) + " "
		}
		return l
	}

	list1, list2 := l(l1), l(l2)
	diff := len(list1) - len(list2)
	if diff == 0 && list1 == list2 {
		return RelationEqual
	} else if diff > 0 && strings.Contains(list1, list2) {
		return RelationSuperlist
	} else if diff < 0 && strings.Contains(list2, list1) {
		return RelationSublist
	}

	return RelationUnequal
}
