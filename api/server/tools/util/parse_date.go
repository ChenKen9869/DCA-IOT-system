package util

import "time"

func ParseDate(dateString string) time.Time {
	parsedDate, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		panic(err.Error())
	}
	return parsedDate
}
