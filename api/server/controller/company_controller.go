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

func CreateCompanyController(ctx *gin.Context) {
	name := ctx.PostForm("Name")
	parentId := ctx.PostForm("ParentId")
	location := ctx.PostForm("Location")

	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	parent, err := strconv.Atoi(parentId)
	if err != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, err.Error())
		return
	}
	if parent == 0 {
		id, errService := service.CreateCompanyService(uint(parent), name, user.ID, location)
		if errService != nil {
			server.Response(ctx, http.StatusUnprocessableEntity, 422, nil, errService.Error())
			return
		}
		service.CreateCompanyUserService(id, user.ID)
		server.ResponseSuccess(ctx, gin.H{"CompanyId": id}, server.Success)
		return
	}
	if user.ID != dao.GetCompanyInfoByID(uint(parent)).Owner {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "permission denied")
		return
	}
	id, errService := service.CreateCompanyService(uint(parent), name, user.ID, location)
	if errService != nil {
		server.Response(ctx, http.StatusUnprocessableEntity, 422, nil, errService.Error())
		return
	}
	server.ResponseSuccess(ctx, gin.H{"CompanyId": id}, server.Success)
}

func GetCompanyTreeListController(ctx *gin.Context) {
	userInfo, exists := ctx.Get("user")
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user infromation does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	treeList, companyList := service.GetCompanyTreeListService(user.ID)
	companyTreeList := gin.H{
		"mechanism":    treeList,
		"company_list": companyList,
	}
	server.ResponseSuccess(ctx, companyTreeList, server.Success)
}

func DeleteCompanyController(ctx *gin.Context) {
	companyIdString := ctx.Query("CompanyId")

	companyId, errAtoi := strconv.Atoi(companyIdString)
	if errAtoi != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, errAtoi.Error())
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
	err := service.DeleteCompanyService(uint(companyId), user)
	if err != nil {
		if msg := err.Error(); msg == server.CompanyNotExist {
			server.Response(ctx, http.StatusBadRequest, 400, nil, msg)
			return
		} else if msg == server.NodeHasSubcompany {
			server.Response(ctx, http.StatusBadRequest, 400, nil, msg)
			return
		}
	}
	if dao.GetCompanyInfoByID(uint(companyId)).ParentID == 0 {
		service.DeleteCompanyUserService(uint(companyId), user.ID)
	}
	server.ResponseSuccess(ctx, nil, server.Success)
}

func CreateCompanyUserController(ctx *gin.Context) {
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
	err := service.CreateCompanyUserService(uint(companyId), uint(userId))
	if err != nil {
		if msg := err.Error(); msg == server.CompanyNotExist {
			server.Response(ctx, http.StatusBadRequest, 400, nil, msg)
			return
		} else if msg == server.UserNotExist {
			server.Response(ctx, http.StatusBadRequest, 400, nil, msg)
			return
		} else {
			server.Response(ctx, http.StatusBadRequest, 400, nil, msg)
			return
		}
	}
	server.ResponseSuccess(ctx, gin.H{"companyId": companyId, "userId": userId}, server.Success)
}

func DeleteCompanyUserController(ctx *gin.Context) {
	companyIdString := ctx.Query("CompanyId")
	userIdString := ctx.Query("UserId")

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
	service.DeleteCompanyUserService(uint(companyId), uint(userId))
	server.ResponseSuccess(ctx,
		gin.H{"companyId": companyId, "userId": userId},
		server.Success)
}

func GetEmployeeListController(ctx *gin.Context) {
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
	employeeList := service.GetEmployeeListService(uint(companyId))
	result := []gin.H{}
	for employee := range employeeList {
		result = append(result, gin.H{
			"id":               employee.ID,
			"name":             employee.Name,
			"authCompany":      employeeList[employee],
			"telephone_number": employee.Telephone,
			"email":            employee.Email,
		})
	}
	server.ResponseSuccess(ctx, gin.H{"employeeList": result}, server.Success)
}

func GetCompanyInfoController(ctx *gin.Context) {
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
	if (!service.AuthCompanyUser(user.ID, uint(companyId))) && (!service.AuthVisitor(user.ID, uint(companyId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "permission denied")
		return
	}
	result := service.GetCompanyInfoService(uint(companyId))
	server.ResponseSuccess(ctx, gin.H{"companyInfo": result}, server.Success)
}

func GetOwnCompanyListController(ctx *gin.Context) {
	userInfo, exists := ctx.Get("user")
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user infromation does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	companyList := service.GetOwnCompanyListService(user.ID)
	server.ResponseSuccess(ctx, gin.H{"company_list": companyList}, server.Success)
}

func UpdateCompanyInfoController(ctx *gin.Context) {
	companyName := ctx.Query("Name")
	location := ctx.Query("Location")
	companyIdString := ctx.Query("CompanyId")

	companyId, errAtoi := strconv.Atoi(companyIdString)
	if errAtoi != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, errAtoi.Error())
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
	service.UpdateCompanyInfoService(uint(companyId), companyName, location)
	server.ResponseSuccess(ctx, nil, server.Success)
}
