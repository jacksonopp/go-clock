package helpers

import "time"

func SimpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(1337, time.January, 1, hours, minutes, seconds, 0, time.UTC)
}

func TestName(t time.Time) string {
	return t.Format("15:05:04")
}
