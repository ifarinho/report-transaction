package datetime

import "time"

func TimeInUtc(t time.Time) time.Time {
	return t.In(time.UTC)
}

func ParseDateOnly(s string) (time.Time, error) {
	return time.Parse(time.DateOnly, s)
}
