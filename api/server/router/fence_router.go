package router

import (
	"go-backend/api/server/controller"
	"go-backend/api/common/middleware"
	"github.com/gin-gonic/gin"
)

func FenceRouter(r *gin.Engine) *gin.Engine {
	fence := r.Group("/fence")
	fence.Use(middleware.AuthMiddleware())

	fence.POST("/create", controller.CreateFenceController)
	
	fence.GET("/get_active_list", controller.GetActiveFenceByCompanyIdController)
	fence.GET("/get_status", controller.GetActiveFenceStat)

	fence.DELETE("/stop", controller.StopFenceController)

	return r
}