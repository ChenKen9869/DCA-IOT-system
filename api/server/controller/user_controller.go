package controller

import (
	"go-backend/api/server/entity"
	"go-backend/api/server/service"
	"go-backend/api/server/tools/server"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserRegisterController(ctx *gin.Context) {
	name := ctx.PostForm("Name")
	password := ctx.PostForm("Password")
	telephone := ctx.PostForm("Telephone")
	email := ctx.PostForm("Email")

	id, token, err := service.RegisterService(name, password, telephone, email)
	if err != nil {
		if msg := err.Error(); msg == server.NameTooShort ||
			msg == server.PasswordTooShort ||
			msg == server.UsernameAlreadyExist {
			server.Response(ctx, http.StatusUnprocessableEntity, 422, nil, msg)
			return
		} else if msg == server.PasswordEncryptionFailed ||
			msg == server.TokenGenerateFailed {
			server.Response(ctx, http.StatusInternalServerError, 500, nil, msg)
			return
		}
	}
	server.ResponseSuccess(ctx, gin.H{"token": token, "id": id}, server.Success)
}

func UserLoginController(ctx *gin.Context) {
	name := ctx.PostForm("Name")
	password := ctx.PostForm("Password")

	id, token, defaultCompany, err := service.LoginService(name, password)
	if err != nil {
		if msg := err.Error(); msg == server.PasswordTooShort ||
			msg == server.NameTooShort {
			server.Response(ctx, http.StatusUnprocessableEntity, 422, nil, msg)
			return
		} else if msg == server.UserNotExist ||
			msg == server.WrongPassword {
			server.Response(ctx, http.StatusBadRequest, 400, nil, msg)
			return
		} else if msg == server.TokenGenerateFailed {
			server.Response(ctx, http.StatusInternalServerError, 500, nil, msg)
			return
		}
	}
	server.ResponseSuccess(ctx, gin.H{"token": token, "id": id, "default_company": defaultCompany}, server.Success)
}

func GetUserInfoController(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user_info := user.(entity.User)
	userInfo, err := service.InfoService(user_info.ID)
	if err != nil {
		if msg := err.Error(); msg == server.UserNotExist {
			server.Response(ctx, http.StatusBadRequest, 400, nil, server.UserNotExist)
			return
		}
	}
	server.ResponseSuccess(ctx, *userInfo, server.Success)
}

func UpdateUserInfoController(ctx *gin.Context) {
	name := ctx.Query("Name")
	password := ctx.Query("Password")
	telephone := ctx.Query("Telephone")
	email := ctx.Query("Email")
	
	user, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user_info := user.(entity.User)
	service.UpdateUserInfoService(user_info.ID, name, password, telephone, email)
	server.ResponseSuccess(ctx, nil, server.Success)
}

func UpdateUserDefaultCompanyController(ctx *gin.Context) {
	companyIdString := ctx.Query("CompanyId")
	
	companyId, _ := strconv.Atoi(companyIdString)
	user, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user_info := user.(entity.User)
	if (!service.AuthCompanyUser(user_info.ID, uint(companyId))) && (!service.AuthVisitor(user_info.ID, uint(companyId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	service.UpdateUserDefaultCompanyService(user_info.ID, uint(companyId))
	server.ResponseSuccess(ctx, nil, server.Success)
}