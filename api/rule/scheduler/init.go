package scheduler

import (
	"time"

	"github.com/robfig/cron/v3"
)

func InitScheduler() {
	RuleCron = cron.New()
	RuleMap = make(map[uint]cron.EntryID)
	RuleCron.Start()
	ScheduledMap = make(map[uint]*time.Timer)
}