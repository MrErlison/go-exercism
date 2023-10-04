package meetup

import (
	"time"
)

type WeekSchedule int

const (
	First  WeekSchedule = 1
	Second WeekSchedule = 8
	Third  WeekSchedule = 15
	Fourth WeekSchedule = 22
	Teenth WeekSchedule = 13
	Last   WeekSchedule = 0
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	var countDay WeekSchedule = 1
	if wSched == Last {
		countDay = -1
		month++
	}
	t := time.Date(year, month, int(wSched), 0, 0, 0, 0, time.UTC)
	for ; t.Weekday() != wDay; wSched += countDay {
		t = time.Date(year, month, int(wSched), 0, 0, 0, 0, time.UTC)
	}
	return int(t.Day())
}
