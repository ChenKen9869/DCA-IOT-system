package router

import (
	"go-backend/controller"
	"go-backend/middleware"

	"github.com/gin-gonic/gin"
)

func CompanyRouter(r *gin.Engine) *gin.Engine {
	company := r.Group("/company")
	company.Use(middleware.AuthMiddleware())
	company.POST("/create", controller.CreateCompanyController)
	company.DELETE("/delete", controller.DeleteCompanyController)
	
	companyGet := company.Group("/get")
	companyGet.GET("/treelist", controller.GetCompanyTreeListController)
	companyGet.GET("/employeelist", controller.GetEmployeeListController)
	companyGet.GET("/info", controller.GetCompanyInfoController)

	companyUser := company.Group("/company_user")
	companyUser.POST("/create", controller.CreateCompanyUserController)
	companyUser.DELETE("/delete", controller.DeleteCompanyUserController)
	return r
}

