package date_utils

import "time"

const apiDateLayout = "2006-01-02T15:04:05Z"

func GetNow(isUTC bool) string {
	now := time.Now()
	if isUTC {
		now = now.UTC()
	}
	return now.Format(apiDateLayout)
}