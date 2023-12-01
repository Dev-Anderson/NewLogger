package util

import (
	"time"
)

func TimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func CurrentDate() string {
	t := time.Now()
	date := t.Format("2006-01-02")
	return date
}
