package date_util

import "time"

const (
	DATE_FORMAT = "2006-01-02T15:04:05Z"
)

func TimeNow() string {
	time := time.Now()

	return time.Format(DATE_FORMAT)
}
