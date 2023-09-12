package controller

import (
	"go-backend/api/server/dao"
	"go-backend/api/server/entity"
	"go-backend/api/server/service"
	"go-backend/api/server/tools/server"
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
		server.Response(ctx, http.StatusInternalServerError, 500, nil, errAtoiComanyId.Error())
	}
	userInfo, exists := ctx.Get("user")
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user infromation does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	if user.ID != dao.GetCompanyInfoByID(uint(companyId)).Owner {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "permission denied")
		return
	}
	service.CreateVisitorService(uint(companyId), uint(userId))
	server.ResponseSuccess(ctx, gin.H{"companyId": companyId, "userId": userId}, server.Success)
}

func DeleteVisitorController(ctx *gin.Context) {
	companyIdString := ctx.Query("CompanyId")
	userIdString := ctx.Query("UserId")

	companyId, errAtoiComanyId := strconv.Atoi(companyIdString)
	userId, errAtoiUserId := strconv.Atoi(userIdString)
	if errAtoiComanyId != nil || errAtoiUserId != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, errAtoiComanyId.Error())
		return
	}
	userInfo, exists := ctx.Get("user")
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user infromation does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	if user.ID != dao.GetCompanyInfoByID(uint(companyId)).Owner {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "permission denied")
		return
	}
	if !dao.GetVisitorInfoExists(uint(companyId), uint(userId)) {
		server.Response(ctx, http.StatusBadRequest, 400, nil, "permission information does not exist")
		return
	}
	service.DeleteVisitorService(uint(companyId), uint(userId))
	server.ResponseSuccess(ctx, gin.H{"companyId": companyId, "userId": userId}, server.Success)
}

func GetVisitorListController(ctx *gin.Context) {
	companyIdString := ctx.Query("CompanyId")

	companyId, errAtoiComanyId := strconv.Atoi(companyIdString)
	if errAtoiComanyId != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, errAtoiComanyId.Error())
	}
	userInfo, exists := ctx.Get("user")
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user infromation does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	if user.ID != dao.GetCompanyInfoByID(uint(companyId)).Owner {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "permission denied")
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
