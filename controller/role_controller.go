package controller

import (
	"go-backend/entity"
	"go-backend/server"
	"go-backend/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateVisitorController(ctx *gin.Context) {
	companyIdString := ctx.PostForm("CompanyId")
	userIdString := ctx.PostForm("UserId")
	companyId, errAtoiComanyId := strconv.Atoi(companyIdString)
	userId, errAtoiUserId := strconv.Atoi(userIdString)
	if errAtoiComanyId != nil || errAtoiUserId != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "服务器内部错误")
	}
	userInfo, exists := ctx.Get("user")
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user infromation does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	// 权限验证
	if (!service.AuthCompanyUser(user.ID, uint(companyId))) && (!service.AuthVisitor(user.ID, uint(companyId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	service.CreateVisitorService(uint(companyId), uint(userId))

	server.ResponseSuccess(ctx,
		gin.H{"companyId": companyId, "userId": userId},
		server.Success)
}

func DeleteVisitorController(ctx *gin.Context) {
	companyIdString := ctx.Query("CompanyId")
	userIdString := ctx.Query("UserId")
	companyId, errAtoiComanyId := strconv.Atoi(companyIdString)
	userId, errAtoiUserId := strconv.Atoi(userIdString)
	if errAtoiComanyId != nil || errAtoiUserId != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "服务器内部错误")
	}
	userInfo, exists := ctx.Get("user")
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user infromation does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	// 权限验证
	if (!service.AuthCompanyUser(user.ID, uint(companyId))) && (!service.AuthVisitor(user.ID, uint(companyId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	service.DeleteVisitorService(uint(companyId), uint(userId))
	server.ResponseSuccess(ctx,
		gin.H{"companyId": companyId, "userId": userId},
		server.Success)
}

func GetVisitorListController(ctx *gin.Context) {
	companyIdString := ctx.Query("CompanyId")
	companyId, errAtoiComanyId := strconv.Atoi(companyIdString)
	if errAtoiComanyId != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "服务器内部错误")
	}
	userInfo, exists := ctx.Get("user")
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user infromation does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	// 权限验证
	if (!service.AuthCompanyUser(user.ID, uint(companyId))) && (!service.AuthVisitor(user.ID, uint(companyId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	visitorList := service.GetVisitorListService(uint(companyId))

	result := []gin.H{}
	for visitor := range visitorList {
		result = append(result, gin.H{
			"id":          visitor.ID,
			"name":        visitor.Name,
			"authCompany": visitorList[visitor],
		})
	}

	server.ResponseSuccess(ctx, gin.H{"visitorList": result}, server.Success)
}

func GetVisitorCompanyListController(ctx *gin.Context) {
	userInfo, exists := ctx.Get("user")
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user infromation does not exists in application context")
		return
	}
	user := userInfo.(entity.User)

	visitorCompanyList := service.GetVisitorCompanyListService(user.ID)

	server.ResponseSuccess(ctx, gin.H{"visitorCompanyList": visitorCompanyList}, server.Success)
}