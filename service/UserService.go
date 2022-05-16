package service

import (
	"go-backend/common"
	"go-backend/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)
// 工具方法: 判断用户名是否存在
func isUserNameExist(db *gorm.DB, name string) bool {
	var user model.User
	db.Where("name = ?", name).First(&user)
	return user.ID != 0
}
// Service: 用户注册
// @Summary API of golang gin backend
// @Tags User
// @description user register
// @version 1.0
// @accept application/json
// @param user body string true "userinfo"
// @Success 200 {object} swagResponse.SuccessResponse200 "注册成功"
// @Failure 422 {object} swagResponse.FailureResponse422 "输入参数错误"
// @Failure 500 {object} swagResponse.FailureResponse500 "系统异常"
// @router /user/register [post]
func Register(ctx *gin.Context) {
	DB := common.GetDB()

	var requestUser = model.User{}
	ctx.Bind(&requestUser)
	
	name := requestUser.Name
	password := requestUser.Password

	// 数据验证
	if len(name) < 2 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名不能小于2")
		return
	}
	if len(password) < 6 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}
	// 判断用户是否存在
	if isUserNameExist(DB, name) {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名已存在")
		return
	}

	// 创建用户
	// 密码加密
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "密码加密失败")
		return
	}
	// 写入数据库
	newUser := model.User{
		Name:	name,
		Password: string(hasePassword),
	}
	DB.Create(&newUser)

	// 发放 token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "系统异常, token生成失败")
		log.Printf("token generate error: %v", err)
		return
	}

	// 返回结果
	common.ResponseSuccess(ctx, gin.H{"token": token}, "注册成功")
}


// Service: 用户登录
// @Summary API of golang gin backend
// @Tags User
// @description user login
// @version 1.0
// @accept application/json
// @param user body string true "userinfo"
// @Success 200 {object} swagResponse.SuccessResponse200 "登录成功"
// @Failure 422 {object} swagResponse.FailureResponse422 "输入参数错误"
// @Failure 500 {object} swagResponse.FailureResponse500 "系统异常"
// @router /user/login [post]
func Login(ctx *gin.Context) {
	DB := common.GetDB()

	var requestUser = model.User{}
	ctx.Bind(&requestUser)

	name := requestUser.Name
	password := requestUser.Password
	// 数据验证
	if len(name) < 2 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名不能小于2")
		return
	}
	if len(password) < 6 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}
	// 判断用户名与密码是否正确
	var user model.User
	DB.Where("name = ?", name).First(&user)
	if user.ID == 0 {
		common.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "用户不存在")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		common.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}	
	
	// 发放 token
	token, err := common.ReleaseToken(user)
	if err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "系统异常, token生成失败")
		log.Printf("token generate error: %v", err)
		return
	}

	// 返回结果
	common.ResponseSuccess(ctx, gin.H{"token": token}, "登录成功")
}

// Service: 获取当前用户信息
// @Summary API of golang gin backend
// @Tags User
// @description get user informations
// @version 1.0
// @accept application/json
// @param name query string true "username"
// @param Authorization header string true "token"
// @Success 200 {object} swagResponse.SuccessResponse200 "查询成功"
// @Failure 400 {object} swagResponse.FailureResponse400 "用户信息不存在"
// @Failure 401 {object} swagResponse.FailureResponse401 "权限不足"
// @router /user/info [get]
func Info(ctx *gin.Context) {
	DB := common.GetDB()

	name := ctx.Query("name")

	// 执行查询
	var user model.User
	DB.Where("name = ?", name).Find(&user)
	if user.ID == 0 {
		common.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "用户信息不存在")
		return
	}
	info_map := gin.H{
		"name": name,
		"id": user.ID,
		"create_time": user.CreatedAt,
		"update_time": user.UpdatedAt,
	}

	// 返回结果
	common.ResponseSuccess(ctx, info_map, "查询成功")
}
