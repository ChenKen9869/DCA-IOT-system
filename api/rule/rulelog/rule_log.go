package rulelog

import (
	"log"
	"os"
	"time"
)

var RuleLog *log.Logger

func InitRuleLogger() {
	fileDir := "./logs/rule"
	os.MkdirAll(fileDir, os.ModePerm)

	file := "./logs/rule/" + time.Now().Format(time.RFC3339) + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	RuleLog = log.New(logFile, "[Rule Model]", log.LstdFlags|log.Lshortfile|log.LUTC)
}
