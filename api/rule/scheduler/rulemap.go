package scheduler

import "github.com/robfig/cron/v3"

var RuleMap map[uint]cron.EntryID = make(map[uint]cron.EntryID)

var RuleCron *cron.Cron = cron.New()

func InitRuleScheduler() {
	RuleCron.Start()
}
