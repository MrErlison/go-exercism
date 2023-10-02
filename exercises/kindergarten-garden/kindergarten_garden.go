package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

type Garden map[string][]string

func NewGarden(diagram string, children []string) (*Garden, error) {
	if len(diagram) <= 0 || diagram[0] != '\n' {
		return nil, errors.New("not a valid garden")
	}

	if children == nil {
		return nil, errors.New("no children")
	}

	sortedChildren := make([]string, len(children))
	copy(sortedChildren, children)
	sort.Strings(sortedChildren)
	for k, child := range sortedChildren[1:] {
		if sortedChildren[k] == child {
			return nil, errors.New("child is duplicated")
		}
	}

	plantsName := map[rune]string{
		'V': "violets",
		'R': "radishes",
		'C': "clover",
		'G': "grass",
	}

	garden := Garden{}
	for _, row := range strings.Split(diagram[1:], "\n") {

		if len(children) != len(row)/2 || len(children)*2 != len(row) {
			return nil, errors.New("not enough children")
		}

		for p, plantCode := range row {
			plantName, ok := plantsName[plantCode]
			if !ok {
				return nil, errors.New("invalid plant code")
			}

			child := sortedChildren[p/2]
			plants, ok := garden[child]
			if !ok {
				plants = []string{}
			}

			garden[child] = append(plants, plantName)
		}
	}

	return &garden, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	p, ok := (*g)[child]
	return p, ok
}
