package router

import (
	"go-backend/middleware"
	"go-backend/monitor"

	"github.com/gin-gonic/gin"
)

func MonitorRouter(r *gin.Engine) *gin.Engine {
	monitorCentor := r.Group("/monitorCentor")
	monitorCentor.Use(middleware.AuthMiddleware())
	
	monitorCentor.GET("/connect", monitor.ConnectToMonitorCentor)
	monitorCentor.DELETE("/disconnect", monitor.DisconnectMonitorCentor)

	return r
}
