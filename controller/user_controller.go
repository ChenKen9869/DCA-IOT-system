package controller

import (
	"go-backend/entity"
	"go-backend/server"
	"go-backend/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserRegisterController(ctx *gin.Context) {
	name := ctx.PostForm("Name")
	password := ctx.PostForm("Password")

	user := entity.User{
		Name: name,
		Password: password,
	}

	id, token, err := service.RegisterService(user)
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
	// body, _:= util.ReadAll(ctx.Request.Body)
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
	idString := ctx.Query("Id")

	id, err:= strconv.Atoi(idString)
	if err != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "server inter failed")
		return
	}
	userInfo, err := service.InfoService(uint(id))
	if err != nil {
		if msg := err.Error(); msg == server.UserNotExist {
			server.Response(ctx, http.StatusBadRequest, 400, nil, server.UserNotExist)
			return
		}
	}
	server.ResponseSuccess(ctx, *userInfo, server.Success)
}
