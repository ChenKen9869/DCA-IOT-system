package middleware

import (
	"go-backend/api/common/common"
	"go-backend/api/server/dao"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get Authorization header
		tokenString := ctx.GetHeader("Authorization")

		// fmt.Print("请求token", tokenString)

		// validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:] //截取字符

		token, claims, err := common.ParseToken(tokenString)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		// token通过验证, 获取claims中的UserID
		userId := claims.UserId
		user := dao.GetUserInfoById(userId)
		// 验证用户是否存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		//用户存在 将user信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}
