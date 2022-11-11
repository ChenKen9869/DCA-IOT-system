package monitor

import (
	"strconv"
	"time"
)

const (
	FenceJob = "fenceJob"
)

func MakeMessage(recordType string, recordId uint, body string) string {
	id := strconv.Itoa(int(recordId))
	time := time.Now()
	return time.String() + "#" + recordType + "#" + id + "#" + body
}