package router

import (
	"go-backend/api/common/middleware"
	"go-backend/api/server/controller"

	"github.com/gin-gonic/gin"
)

func RuleRouter(r *gin.Engine) *gin.Engine {
	rule := r.Group("/rule")
	rule.Use(middleware.AuthMiddleware())

	rule.POST("/create", controller.CreateRuleController)
	rule.PUT("/update", controller.UpdateRuleController)
	rule.GET("/start", controller.StartRuleController)
	rule.GET("/schedule", controller.ScheduleRuleController)
	rule.GET("/end", controller.EndRuleController)
	rule.DELETE("/delete", controller.DeleteRuleController)

	rule.GET("/get_user", controller.GetUserRuleController)
	rule.GET("/get_company", controller.GetCompanyRuleController)

	return r
}
