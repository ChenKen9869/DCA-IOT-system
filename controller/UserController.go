package controller

import (
	"go-backend/middleware"
	"go-backend/service"

	"github.com/gin-gonic/gin"
)

// 注册路由
func UserController(r *gin.Engine) *gin.Engine {
	user := r.Group("/user")
	user.Use(middleware.CORSMiddleware())
	user.POST("/register", service.Register)
	user.POST("/login", service.Login)
	user.GET("/info", middleware.AuthMiddleware(), service.Info)

	return r
}

// 下面的现在没用
// 注册时需要发放 token， Engine是路由

func Register(r *gin.Engine) *gin.Engine {
	r.POST("/user/register", service.Register)
	return r
}

// 登录时需要验证用户名与密码，然后发放 token
func Login(r *gin.Engine) *gin.Engine {
	r.POST("/user/login", service.Login)
	return r
}

// 获取当前用户信息
func Info(r *gin.Engine) *gin.Engine {
	r.GET("/user/info", middleware.AuthMiddleware(), service.Info)
	return r
}