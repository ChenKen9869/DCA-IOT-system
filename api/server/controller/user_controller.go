package controller

import (
	"go-backend/api/server/entity"
	"go-backend/api/server/tools/server"
	"go-backend/api/server/service"
	"net/http"
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

	id, token, err := service.LoginService(name, password)
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
	server.ResponseSuccess(ctx, gin.H{"token": token, "id": id}, server.Success)
}

func UserInfoController(ctx *gin.Context) {
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
