package util

import "time"

func ParseDate(dateString string) time.Time {
	rfc3339MilliLayout := "2006-01-02T15:04:05.999Z07:00"
	parsedDate, err := time.Parse(rfc3339MilliLayout, dateString)
	if err != nil {
		panic(err.Error())
	}
	return parsedDate
}
