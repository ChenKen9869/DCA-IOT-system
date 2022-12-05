package router

import (
	"go-backend/api/server/controller"
	"go-backend/api/common/middleware"
	"github.com/gin-gonic/gin"
)

// 注册路由
func UserRouter(r *gin.Engine) *gin.Engine {
	user := r.Group("/user")
	
	user.POST("/register", controller.UserRegisterController)
	user.POST("/login", controller.UserLoginController)
	
	user.GET("/info", middleware.AuthMiddleware(), controller.GetUserInfoController)

	user.PUT("/update", middleware.AuthMiddleware(), controller.UpdateUserInfoController)

	return r
}