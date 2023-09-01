package router

import (
	"go-backend/api/common/middleware"
	"go-backend/api/sys/iot/monitor"

	"github.com/gin-gonic/gin"
)

func MonitorRouter(r *gin.Engine) *gin.Engine {
	monitorCenter := r.Group("/monitorCenter")
	monitorCenter.Use(middleware.AuthMiddleware())

	monitorCenter.GET("/connect", monitor.ConnectToMonitorCenter)

	monitorCenter.DELETE("/disconnect", monitor.DisconnectMonitorCenter)

	return r
}
