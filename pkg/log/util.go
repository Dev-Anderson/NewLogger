package log

import (
	"time"
)

func timeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func currentDate() string {
	t := time.Now()
	date := t.Format("2006-01-02")
	return date
}
