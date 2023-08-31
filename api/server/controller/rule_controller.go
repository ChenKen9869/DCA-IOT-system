package controller

import (
	"go-backend/api/server/dao"
	"go-backend/api/server/entity"
	"go-backend/api/server/service"
	"go-backend/api/server/tools/server"
	"go-backend/api/server/tools/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRuleController(ctx *gin.Context) {
	// Rule Upload
	dString := ctx.PostForm("Datasource")
	cString := ctx.PostForm("Condition")
	aString := ctx.PostForm("Action")
	companyIdString := ctx.PostForm("CompanyId")

	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)

	companyId, err := strconv.Atoi(companyIdString)
	if err != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "server inter failed")
		return
	}
	if !service.AuthCompanyUser(user.ID, uint(companyId)) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}

	// Save Rule
	ruleId := service.CreateRuleService(uint(companyId), dString, cString, aString, user.ID)
	server.ResponseSuccess(ctx, gin.H{"ruleId": ruleId}, server.Success)
}

// 用 rule id 做索引找到退出管道，以无限循环的子进程方式实现模式匹配器，利用管道将数据推送至模式匹配器
func StartRuleController(ctx *gin.Context) {
	ruleIdString := ctx.Query("RuleId")
	execInternal := ctx.Query("ExecInternal")
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)

	ruleId, err := strconv.Atoi(ruleIdString)
	if err != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, err.Error())
		return
	}
	// auth
	rule := dao.GetRuleInfo(uint(ruleId))
	if !service.AuthCompanyUser(user.ID, rule.ParentId) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	service.StartRuleService(rule.ID, execInternal)
	server.ResponseSuccess(ctx, nil, server.Success)
}

func EndRuleController(ctx *gin.Context) {
	// 根据 rule id 查找 cron id，然后 remove 掉这个任务即可
	ruleIdString := ctx.Query("RuleId")
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)

	ruleId, err := strconv.Atoi(ruleIdString)
	if err != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, err.Error())
		return
	}
	// auth
	rule := dao.GetRuleInfo(uint(ruleId))
	if !service.AuthCompanyUser(user.ID, rule.ParentId) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	if rule.Stat == entity.RuleInactive {
		server.Response(ctx, http.StatusForbidden, 403, nil, "error: rule is inactive!")
		return
	}
	service.EndRuleService(rule.ID)
	server.ResponseSuccess(ctx, nil, server.Success)
}

func DeleteRuleController(ctx *gin.Context) {
	// 只有非激活状态的规则可以删除
	ruleIdString := ctx.Query("RuleId")
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)

	ruleId, err := strconv.Atoi(ruleIdString)
	if err != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, err.Error())
		return
	}
	// auth
	rule := dao.GetRuleInfo(uint(ruleId))
	if !service.AuthCompanyUser(user.ID, rule.ParentId) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	if rule.Stat != entity.RuleInactive {
		server.Response(ctx, http.StatusForbidden, 403, nil, "error: rule is active or scheduled!")
		return
	}
	deletedRule := service.DeleteRuleService(rule.ID)
	server.ResponseSuccess(ctx, gin.H{"deleted_rule": deletedRule}, server.Success)
}

func UpdateRuleController(ctx *gin.Context) {
	dString := ctx.Query("Datasource")
	cString := ctx.Query("Condition")
	aString := ctx.Query("Action")
	ruleIdString := ctx.Query("RuleId")

	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)

	ruleId, err := strconv.Atoi(ruleIdString)
	if err != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, err.Error())
		return
	}
	// auth
	rule := dao.GetRuleInfo(uint(ruleId))
	if !service.AuthCompanyUser(user.ID, rule.ParentId) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	if rule.Stat != entity.RuleInactive {
		server.Response(ctx, http.StatusForbidden, 403, nil, "error: rule is active or scheduled!")
		return
	}
	// Update Rule
	service.UpdateRuleDCAService(rule.ID, dString, cString, aString)
	server.ResponseSuccess(ctx, nil, server.Success)
}

func ScheduleRuleController(ctx *gin.Context) {
	// 利用 cron 的 schedule 实现这个 schedule
	ruleIdString := ctx.Query("RuleId")
	execInternal := ctx.Query("ExecInternal")
	futureTime := ctx.Query("FutureTime")

	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)

	ruleId, err := strconv.Atoi(ruleIdString)
	if err != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, err.Error())
		return
	}
	// auth
	rule := dao.GetRuleInfo(uint(ruleId))
	if !service.AuthCompanyUser(user.ID, rule.ParentId) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	service.ScheduleRuleService(rule.ID, execInternal, util.ParseDate(futureTime))
	server.ResponseSuccess(ctx, nil, server.Success)
}

func GetCompanyRuleController(ctx *gin.Context) {
	// 查询一个公司的所有 rule（包括其子公司）
	companyIdString := ctx.Query("CompanyId")
	companyId, errAtoi := strconv.Atoi(companyIdString)
	if errAtoi != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "atoi error")
		return
	}
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	if (!service.AuthCompanyUser(user.ID, uint(companyId))) && (!service.AuthVisitor(user.ID, uint(companyId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	ruleList := service.GetCompanyRuleService(uint(companyId))
	server.ResponseSuccess(ctx, gin.H{"rule_list": ruleList}, server.Success)
}

func GetUserRuleController(ctx *gin.Context) {
	// 查询一个用户创建的所有 rule
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	ruleList := service.GetUserRuleService(user.ID)
	server.ResponseSuccess(ctx, gin.H{"rule_list": ruleList}, server.Success)
}
