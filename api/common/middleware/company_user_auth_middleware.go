package middleware

import (
	"go-backend/api/server/dao"
	"go-backend/api/server/entity"
	"go-backend/api/server/tools/server"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)
// 没有使用
func CompanyUserAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user, exists = ctx.Get("user")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足, token 未写入用户信息 ",
			})
			ctx.Abort()
			return	
		}
		companyIdString := ctx.Query("companyId")
		company_id, err := strconv.Atoi(companyIdString)
		companyId := uint(company_id)
		if err != nil {
			server.Response(ctx, http.StatusInternalServerError, 500, nil, "服务器内部错误")
			ctx.Abort()
			return
		}
		userId := user.(*entity.User).ID
		// 验证用户是否有公司的操作权限
		companyList := dao.GetCompanyListByUserID(userId)
		for _, company := range companyList {
			if companyId == company.CompanyID {
				company := dao.GetCompanyInfoByID(companyId)
				ctx.Set("company", company)
				ctx.Next()
				return
			}
		}
		ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足, 用户不具备操作企业的权限",
			})
		ctx.Abort()
	}
}
