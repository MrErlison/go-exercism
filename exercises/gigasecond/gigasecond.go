package gigasecond

import "time"

// AddGigasecond returns the date and time one gigasecond after a certain date.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
