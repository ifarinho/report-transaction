package datetime

import "time"

func TimeInUtc(t time.Time) time.Time {
	return t.In(time.UTC)
}

func ReportFormat(t time.Time) string {
	return t.Format("20060102T150405")
}
