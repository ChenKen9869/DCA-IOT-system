package router

import (
	"go-backend/controller"
	"go-backend/middleware"

	"github.com/gin-gonic/gin"
)

// 注册路由
func UserRouter(r *gin.Engine) *gin.Engine {
	user := r.Group("/user")
	user.POST("/register", controller.UserRegisterController)
	user.POST("/login", controller.UserLoginController)
	user.GET("/info", middleware.AuthMiddleware(), controller.UserInfoController)

	return r
}