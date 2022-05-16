package controller

import (
	"github.com/gin-gonic/gin"
	"go-backend/service"
)

/*
1.注册公司 register(name)
2.注销公司 delete(name)
3.注册农牧场 register(name)
4.注销农牧场 delete(name)
5.注册厂房 register(name)
6.注销厂房 delete(name)
7.查看公司信息 find(name)
8.查看农牧场信息 find(name)
9.查看厂房信息 find(name)

*/
func CompanyRegisterController(r *gin.Engine) *gin.Engine {
	r.POST("/api/company/auth/register", service.CompanyRegisterService)
	return r
}