package service

import (
	"go-backend/common"
	"go-backend/model"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

func CompanyRegisterService(ctx *gin.Context) {
	DB := common.GetDB()
	// get request parameters
	var requestCompany = model.Dept{}
	ctx.Bind(&requestCompany)

	// get parameters
	name := requestCompany.Name

	// data validation
	if len(name) < 2 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "公司名不可少于两个字符")
	}

	if isNameExists(DB, name) {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "公司已存在")
	}

	// create company

	newCompany := model.Dept{
		Name : name,
	}

	DB.Create(&newCompany)

}

func isNameExists(db *gorm.DB, name string) bool {
	var company model.Dept
	db.Where("name = ?", name).First(&company)
	return company.ID != 0
}