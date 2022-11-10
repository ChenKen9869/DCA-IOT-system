package router

import (
	"go-backend/controller"
	"go-backend/middleware"

	"github.com/gin-gonic/gin"
)

func RoleRouter(r *gin.Engine) *gin.Engine {
	role := r.Group("/role")
	role.Use(middleware.AuthMiddleware())

	visitor := role.Group("/visitor")
	visitor.POST("/create", controller.CreateVisitorController)
	visitor.DELETE("/delete", controller.DeleteVisitorController)
	visitor.GET("/get_list", controller.GetVisitorListController)
	visitor.GET("/get_company_list", controller.GetVisitorCompanyListController)
	return r
}