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

	ruleId := service.CreateRuleService(uint(companyId), dString, cString, aString, user.ID)
	server.ResponseSuccess(ctx, gin.H{"ruleId": ruleId}, server.Success)
}

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
	rule := dao.GetRuleInfo(uint(ruleId))
	if !service.AuthCompanyUser(user.ID, rule.ParentId) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	service.StartRuleService(rule.ID, execInternal)
	server.ResponseSuccess(ctx, nil, server.Success)
}

func EndRuleController(ctx *gin.Context) {
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
	rule := dao.GetRuleInfo(uint(ruleId))
	if !service.AuthCompanyUser(user.ID, rule.ParentId) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	if rule.Stat != entity.RuleInactive {
		server.Response(ctx, http.StatusForbidden, 403, nil, "error: rule is active or scheduled!")
		return
	}
	service.UpdateRuleDCAService(rule.ID, dString, cString, aString)
	server.ResponseSuccess(ctx, nil, server.Success)
}

func ScheduleRuleController(ctx *gin.Context) {
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
	rule := dao.GetRuleInfo(uint(ruleId))
	if !service.AuthCompanyUser(user.ID, rule.ParentId) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	service.ScheduleRuleService(rule.ID, execInternal, util.ParseDate(futureTime))
	server.ResponseSuccess(ctx, nil, server.Success)
}

func GetCompanyRuleController(ctx *gin.Context) {
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
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	ruleList := service.GetUserRuleService(user.ID)
	server.ResponseSuccess(ctx, gin.H{"rule_list": ruleList}, server.Success)
}
