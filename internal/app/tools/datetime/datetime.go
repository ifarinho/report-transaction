package datetime

import "time"

func TimeInUtc(t time.Time) time.Time {
	return t.In(time.UTC)
}
