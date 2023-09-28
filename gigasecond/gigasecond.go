package gigasecond

import "time"

// AddGigasecond adds one gigasecond to the provided date.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
