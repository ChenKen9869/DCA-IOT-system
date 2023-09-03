package service

import (
	"go-backend/api/rule/ruleparser"
	"go-backend/api/rule/ruleparser/preprosess"
	"go-backend/api/rule/scheduler"
	"go-backend/api/server/dao"
	"go-backend/api/server/entity"
	"go-backend/api/server/vo"
	"strconv"
	"time"
)

// @Summary API of golang gin backend
// @Tags Rule
// @description create a rule
// @version 1.0
// @accept mpfd
// @param Datasource formData string true "datasource"
// @param Condition formData string true "condition"
// @param Action formData string true "action"
// @param CompanyId formData int true "company id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /rule/create [post]
func CreateRuleService(companyId uint, datasource string, condition string, action string, owner uint) uint {
	// prePROCESS: auth devices
	if !preprosess.AuthDevices(datasource, companyId) {
		panic("rule auth error!")
	}
	ruleId := dao.CreateRule(entity.Rule{
		Datasource: datasource,
		Condition:  condition,
		Action:     action,
		Owner:      owner,
		ParentId:   companyId,
		Stat:       entity.RuleInactive,
	})
	return ruleId
}

// @Summary API of golang gin backend
// @Tags Rule
// @description start a rule
// @version 1.0
// @accept application/json
// @param RuleId query int true "rule id"
// @param ExecInternal query string true "exec internal"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /rule/start [get]
func StartRuleService(ruleId uint, internal string) {
	// 先查出来 rule
	rule := dao.GetRuleInfo(ruleId)
	// 判断规则状态是否是未启动
	if rule.Stat != entity.RuleInactive {
		panic("Error: rule has started or scheduled!")
	}
	// preprocess: 取出数据源索引，更新数据源管理器
	preprosess.AddDatasource(rule.Datasource)
	// 调用解释函数，解析文本, 获取模式匹配函数
	matcherFunc := ruleparser.ParseRule(strconv.Itoa(int(rule.ID)), rule.Datasource, rule.Condition, rule.Action)
	// 生成和注册模式匹配器
	cronId, err := scheduler.RuleCron.AddFunc(internal, matcherFunc)
	if err != nil {
		panic(err.Error())
	}
	scheduler.RuleMap[ruleId] = cronId
	// 修改规则状态
	dao.UpdateRuleStat(ruleId, entity.RuleActive)
}

// @Summary API of golang gin backend
// @Tags Rule
// @description end a rule
// @version 1.0
// @accept application/json
// @param RuleId query int true "rule id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /rule/end [get]
func EndRuleService(ruleId uint) {
	scheduler.RuleCron.Remove(scheduler.RuleMap[ruleId])
	preprosess.RemoveDatasource(dao.GetRuleInfo(ruleId).Datasource)
	dao.UpdateRuleStat(ruleId, entity.RuleInactive)
}

// @Summary API of golang gin backend
// @Tags Rule
// @description schedule a rule
// @version 1.0
// @accept application/json
// @param RuleId query int true "rule id"
// @param ExecInternal query string true "exec internal"
// @param FutureTime query string true "future start time"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /rule/schedule [get]
func ScheduleRuleService(ruleId uint, internal string, futureTime time.Time) {
	time.AfterFunc(time.Until(futureTime), func() {
		StartRuleService(ruleId, internal)
	})
}

func getRuleRecursive(companyId uint, ruleList *[]entity.Rule) {
	rules := dao.GetRuleListByCompany(companyId)
	*ruleList = append(*ruleList, rules...)
	childrenList := dao.GetCompanyListByParent(companyId)
	for _, subCompany := range childrenList {
		getRuleRecursive(subCompany.ID, ruleList)
	}
}

// @Summary API of golang gin backend
// @Tags Rule
// @description get rule list by company id
// @version 1.0
// @accept application/json
// @param CompanyId query int true "company id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /rule/get_company [get]
func GetCompanyRuleService(companyId uint) []vo.RuleInfo {
	var result []vo.RuleInfo
	var ruleList []entity.Rule
	getRuleRecursive(companyId, &ruleList)
	for _, rule := range ruleList {
		result = append(result, vo.RuleInfo{
			Id:         rule.ID,
			Datasource: rule.Datasource,
			Condition:  rule.Condition,
			Action:     rule.Action,
			Owner:      rule.Owner,
			ParentId:   rule.ParentId,
			Stat:       rule.Stat,
			CreateTime: rule.CreatedAt,
		})
	}
	return result
}

// @Summary API of golang gin backend
// @Tags Rule
// @description get rule list by user id
// @version 1.0
// @accept application/json
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /rule/get_user [get]
func GetUserRuleService(userId uint) []vo.RuleInfo {
	var result []vo.RuleInfo
	ruleList := dao.GetRuleListByUser(userId)
	for _, rule := range ruleList {
		result = append(result, vo.RuleInfo{
			Id:         rule.ID,
			Datasource: rule.Datasource,
			Condition:  rule.Condition,
			Action:     rule.Action,
			Owner:      rule.Owner,
			ParentId:   rule.ParentId,
			Stat:       rule.Stat,
			CreateTime: rule.CreatedAt,
		})
	}
	return result
}

// @Summary API of golang gin backend
// @Tags Rule
// @description delete an inactive rule
// @version 1.0
// @accept application/json
// @param RuleId query int true "rule id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /rule/delete [delete]
func DeleteRuleService(ruleId uint) vo.RuleInfo {
	rule := dao.DeleteRule(ruleId)
	return vo.RuleInfo{
		Id:         rule.ID,
		Datasource: rule.Datasource,
		Condition:  rule.Condition,
		Action:     rule.Action,
		Owner:      rule.Owner,
		ParentId:   rule.ParentId,
		Stat:       rule.Stat,
		CreateTime: rule.CreatedAt,
	}
}

// @Summary API of golang gin backend
// @Tags Rule
// @description update an inactive rule
// @version 1.0
// @accept application/json
// @param Datasource query string true "datasource"
// @param Condition query string true "condition"
// @param Action query string true "action"
// @param RuleId query int true "rule id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /rule/update [put]
func UpdateRuleDCAService(ruleId uint, datasource string, condition string, action string) {
	dao.UpdateRuleDCA(ruleId, datasource, condition, action)
}
