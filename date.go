package newlogger

import "time"

func CurrentDate() string {
	t := time.Now()
	return t.Format("2006-01-02")
}
