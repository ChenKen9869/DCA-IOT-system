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
	Msg  string      `json:"msg" example:"操作成功"`
}

type FailureResponse400 struct {
	Code int    `json:"code" example:"400"`
	Data string `json:"data" example:"null"`
	Msg  string `json:"msg" example:"语法无效"`
}

type FailureResponse401 struct {
	Code int    `json:"code" example:"401"`
	Data string `json:"data" example:"null"`
	Msg  string `json:"msg" example:"权限不足"`
}

type FailureResponse422 struct {
	Code int    `json:"code" example:"422"`
	Data string `json:"data" example:"null"`
	Msg  string `json:"msg" example:"无法处理"`
}

type FailureResponse500 struct {
	Code int    `json:"code" example:"500"`
	Data string `json:"data" example:"null"`
	Msg  string `json:"msg" example:"服务器内部错误"`
}
