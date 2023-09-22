package clock

import "fmt"

type Clock struct {
	Hour    int
	Minutes int
}

func New(h, m int) Clock {
	h += m / 60

	m %= 60
	if m < 0 {
		m += 60
		h--
	}

	h %= 24
	if h < 0 {
		h += 24
	}
	return Clock{Hour: h, Minutes: m}
}

func (c Clock) Add(m int) Clock {
	return New(c.Hour, c.Minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.Hour, c.Minutes-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minutes)
}
