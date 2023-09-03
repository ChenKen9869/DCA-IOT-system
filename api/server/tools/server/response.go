package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func ResponseSuccess(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func ResponseFail(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, http.StatusOK, 400, data, msg)
}

type SuccessResponse200 struct {
	Code int         `json:"code" example:"200"`
	Data interface{} `json:"data" `
	Msg  string      `json:"msg" example:"success"`
}

type FailureResponse400 struct {
	Code int    `json:"code" example:"400"`
	Data string `json:"data" example:"null"`
	Msg  string `json:"msg" example:"syntax error"`
}

type FailureResponse401 struct {
	Code int    `json:"code" example:"401"`
	Data string `json:"data" example:"null"`
	Msg  string `json:"msg" example:"permission denied"`
}

type FailureResponse422 struct {
	Code int    `json:"code" example:"422"`
	Data string `json:"data" example:"null"`
	Msg  string `json:"msg" example:"unprocessable entity"`
}

type FailureResponse500 struct {
	Code int    `json:"code" example:"500"`
	Data string `json:"data" example:"null"`
	Msg  string `json:"msg" example:"internal server error"`
}
