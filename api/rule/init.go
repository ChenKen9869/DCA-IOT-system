package rule

import (
	"fmt"
	"go-backend/api/rule/accepter"
	"go-backend/api/rule/actions"
	"go-backend/api/rule/rulelog"
	"go-backend/api/rule/ruleparser"
	"go-backend/api/rule/ruleparser/matcher"
	"go-backend/api/rule/scheduler"
)

func InitRule() {
	rulelog.InitRuleLogger()

	accepter.InitAccepter()

	ruleparser.InitRuleparser()
	matcher.InitMatcher()

	scheduler.InitScheduler()

	actions.InitAction()

	fmt.Println("[INITIAL SUCCESS] The rule module is initialized successfully!")
}
