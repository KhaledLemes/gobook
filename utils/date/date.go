package date

import "time"

func DateToString(date time.Time) string {
	return date.Format("2006-01-02")
}

func StringToDate(date string) time.Time {
	return time.Now()
}
