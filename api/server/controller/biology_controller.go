package controller

import (
	"fmt"
	"go-backend/api/server/dao"
	"go-backend/api/server/entity"
	"go-backend/api/server/service"
	"go-backend/api/server/tools/server"
	"go-backend/api/server/tools/util"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func CreateBiologyController(ctx *gin.Context) {
	// 创建 biology 需要：从前端传入的 biology name，该生物的类型，该生物的牧舍（存放在上下文中）
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	biologyName := ctx.PostForm("BiologyName")
	biologyType := ctx.PostForm("BiologyType")
	farmhouseIdString := ctx.PostForm("CompanyId")
	birth := ctx.PostForm("Birthday")
	gender := ctx.PostForm("Gender")

	birthday := util.ParseDate(birth)
	companyId, err := strconv.Atoi(farmhouseIdString)
	if err != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "server inter failed")
		return
	}
	// 验证company与user权限
	if !service.AuthCompanyUser(user.ID, uint(companyId)) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	companyInfo := dao.GetCompanyInfoByID(uint(companyId))

	owner := companyInfo.Owner
	if len(biologyName) < 1 {
		server.Response(ctx, http.StatusBadRequest, 400, nil, "name too short")
		return
	}
	if len(biologyType) < 1 {
		server.Response(ctx, http.StatusBadRequest, 400, nil, "typename too short")
		return
	}
	id := service.CreateBiologyService(biologyName, uint(companyId), biologyType, birthday, gender, owner)

	server.ResponseSuccess(ctx, gin.H{"Id": id}, server.Success)
}

func DeleteBiologyController(ctx *gin.Context) {
	// 从前端传入要删除的生物 id
	operator := ctx.Query("Operator")
	telephoneNumber := ctx.Query("TelephoneNumber")
	leavePlace := ctx.Query("LeavePlace")
	biologyIdString := ctx.Query("Id")
	biologyId, err := strconv.Atoi(biologyIdString)
	if err != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "server inter failed")
		return
	}

	// 验证权限
	userInfo, exists := ctx.Get("user")
	if !exists {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "user info does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	companyId := dao.GetBiologyInfoById(uint(biologyId)).FarmhouseID
	if !service.AuthCompanyUser(user.ID, uint(companyId)) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}

	service.DeleteBiologyService(operator, telephoneNumber, leavePlace, uint(biologyId))

	server.ResponseSuccess(ctx, nil, server.Success)
}

func GetBiologyListController(ctx *gin.Context) {
	companyIdString := ctx.Query("CompanyId")
	companyId, errAtoi := strconv.Atoi(companyIdString)
	if errAtoi != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "atoi error")
		return
	}
	// 验证权限
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	if (!service.AuthCompanyUser(user.ID, uint(companyId))) && (!service.AuthVisitor(user.ID, uint(companyId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}

	// 查询
	biologyList := service.GetBiologyListService(uint(companyId))
	// 构造返回结构
	var result []gin.H
	// 查找每个生物绑定的设备数量
	for _, biologyInfo := range biologyList {
		devices := dao.GetPortableDeviceListByBiology(biologyInfo.ID)
		result = append(result, gin.H{
			"id": biologyInfo.ID,
			"name": biologyInfo.Name,
			"type": biologyInfo.BiologyTypeID,
			"farmhouse_id": biologyInfo.FarmhouseID,
			"device_nums": len(devices),
			"gender": biologyInfo.Gender,
			"birthday": biologyInfo.Birthday,
			"create_date": biologyInfo.CreatedAt,
		})
	}
	resultList := gin.H{
		"biologyList": result,
	}
	server.ResponseSuccess(ctx, resultList, server.Success)
}

func GetBiologyInfoController(ctx *gin.Context) {
	biologyIdString := ctx.Query("BiologyId")
	biologyId, errAtoi := strconv.Atoi(biologyIdString)
	if errAtoi != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "atoi error")
		return
	}
	companyId := dao.GetBiologyInfoById(uint(biologyId)).FarmhouseID
	// 验证权限
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	if (!service.AuthCompanyUser(user.ID, uint(companyId))) && (!service.AuthVisitor(user.ID, uint(companyId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}

	// 查询
	biologyInfo := service.GetBiologyInfoService(uint(biologyId))
	// 返回
	server.ResponseSuccess(ctx, gin.H{"biology_info": biologyInfo}, server.Success)
}

func CreateBiologyTypeController(ctx *gin.Context) {
	biologyTypeId := ctx.PostForm("BiologyTypeId")
	service.CreateBiologyTypeService(biologyTypeId)

	server.ResponseSuccess(ctx, nil, server.Success)
}

func DeleteBiologyTypeController(ctx *gin.Context) {
	biologyTypeId := ctx.PostForm("BiologyTypeId")
	service.DeleteBiologyTypeService(biologyTypeId)

	server.ResponseSuccess(ctx, nil, server.Success)
}

func GetBiologyWithDeviceListController(ctx *gin.Context) {
	companyIdString := ctx.Query("CompanyId")
	companyId, errAtoi := strconv.Atoi(companyIdString)
	if errAtoi != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "atoi error")
		return
	}
	// 验证权限
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	if (!service.AuthCompanyUser(user.ID, uint(companyId))) && (!service.AuthVisitor(user.ID, uint(companyId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	biologyWithDeviceList := service.GetBiologyWithDeviceListService(uint(companyId))
	server.ResponseSuccess(ctx, gin.H{"biologyWithDevice": biologyWithDeviceList}, server.Success)
}

func UpdateBiologyFarmhouseController(ctx *gin.Context) {
	operator := ctx.Query("Operator")
	telephoneNumber := ctx.Query("TelephoneNumber")
	biologyIdString := ctx.Query("BiologyId")
	farmhouseIdString := ctx.Query("FarmhouseId")
	biologyId, errBAtoi := strconv.Atoi(biologyIdString)
	farmhouseId, errFAtoi := strconv.Atoi(farmhouseIdString)
	if errBAtoi != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "biologyId atoi error")
		return
	}
	if errFAtoi != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "farmhouseId atoi error")
		return
	}
	currentBiology := dao.GetBiologyInfoById(uint(biologyId))
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	if !(service.AuthCompanyUser(user.ID, currentBiology.FarmhouseID) && service.AuthCompanyUser(user.ID, uint(farmhouseId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	if int(currentBiology.FarmhouseID) == farmhouseId {
		server.Response(ctx, http.StatusBadRequest, 400, nil, "转入转出地 id 相同")
		return
	}
	if dao.GetCompanyInfoByID(uint(farmhouseId)).Owner != currentBiology.Owner {
		server.Response(ctx, http.StatusBadRequest, 400, nil, "不允许跨集团转舍")
		return	
	}
	service.UpdateBiologyFarmhouseService(operator, telephoneNumber, uint(biologyId), uint(farmhouseId))

	server.ResponseSuccess(ctx, nil, server.Success)
}

func CreateEpidemicPreventionRecordController(ctx *gin.Context) {
	biologyIdString := ctx.PostForm("BiologyId")
	vaccineDescription := ctx.PostForm("VaccineDescription")
	inoculationTime := ctx.PostForm("InoculationTime")
	biologyId, errAtoi := strconv.Atoi(biologyIdString)
	if errAtoi != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "biology id atoi err")
		return
	}
	// auth
	farmhouseId := dao.GetBiologyInfoById(uint(biologyId)).FarmhouseID
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	if !service.AuthCompanyUser(user.ID, farmhouseId) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}

	service.CreateEpidemicPreventionRecordService(uint(biologyId), vaccineDescription, inoculationTime)
	server.ResponseSuccess(ctx, nil, server.Success)
}

func GetEpidemicPreventRecordListController(ctx *gin.Context) {
	biologyIdString := ctx.Query("BiologyId")
	biologyId, errAtoi := strconv.Atoi(biologyIdString)
	if errAtoi != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "atoi err")
		return
	}
	// auth
	farmhouseId := dao.GetBiologyInfoById(uint(biologyId)).FarmhouseID
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	if (!service.AuthCompanyUser(user.ID, uint(farmhouseId))) && (!service.AuthVisitor(user.ID, uint(farmhouseId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}

	resultList := service.GetEpidemicPreventionRecordListService(uint(biologyId))
	server.ResponseSuccess(ctx, gin.H{
		"result_list": resultList,
	}, server.Success)
}

func CreateOperationRecordController(ctx *gin.Context) {
	biologyIdString := ctx.PostForm("BiologyId")
	doctor := ctx.PostForm("Doctor")
	operationTime := ctx.PostForm("OperationTime")
	processDescription := ctx.PostForm("ProcessDescription")
	result := ctx.PostForm("Result")
	biologyId, errAtoi := strconv.Atoi(biologyIdString)
	if errAtoi != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "biology id atoi err")
		return
	}
	// auth
	farmhouseId := dao.GetBiologyInfoById(uint(biologyId)).FarmhouseID
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	if (!service.AuthCompanyUser(user.ID, uint(farmhouseId))) && (!service.AuthVisitor(user.ID, uint(farmhouseId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}

	service.CreateOperationRecordService(uint(biologyId), doctor, operationTime, processDescription, result)
	server.ResponseSuccess(ctx, nil, server.Success)
}

func GetOperationRecordListController(ctx *gin.Context) {
	biologyIdString := ctx.Query("BiologyId")
	biologyId, errAtoi := strconv.Atoi(biologyIdString)
	if errAtoi != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "atoi err")
		return
	}
	// auth
	farmhouseId := dao.GetBiologyInfoById(uint(biologyId)).FarmhouseID
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	if (!service.AuthCompanyUser(user.ID, uint(farmhouseId))) && (!service.AuthVisitor(user.ID, uint(farmhouseId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}

	resultList := service.GetOperationRecordListService(uint(biologyId))
	server.ResponseSuccess(ctx, gin.H{
		"result_list": resultList,
	}, server.Success)
}

func CreateMedicalRecordController(ctx *gin.Context) {
	biologyIdString := ctx.PostForm("BiologyId")
	disease := ctx.PostForm("Disease")
	illnessTime := ctx.PostForm("IllnessTime")
	treatmentPlan := ctx.PostForm("TreatmentPlan")
	biologyId, errAtoi := strconv.Atoi(biologyIdString)
	if errAtoi != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "biology id atoi err")
		return
	}
	// auth
	farmhouseId := dao.GetBiologyInfoById(uint(biologyId)).FarmhouseID
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	if !service.AuthCompanyUser(user.ID, farmhouseId) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}

	service.CreateMedicalRecordService(uint(biologyId), disease, illnessTime, treatmentPlan)
	server.ResponseSuccess(ctx, nil, server.Success)
}

func GetMedicalRecordListController(ctx *gin.Context) {
	biologyIdString := ctx.Query("BiologyId")
	biologyId, errAtoi := strconv.Atoi(biologyIdString)
	if errAtoi != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "atoi err")
		return
	}
	// auth
	farmhouseId := dao.GetBiologyInfoById(uint(biologyId)).FarmhouseID
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	if (!service.AuthCompanyUser(user.ID, uint(farmhouseId))) && (!service.AuthVisitor(user.ID, uint(farmhouseId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}

	resultList := service.GetMedicalRecordListService(uint(biologyId))
	server.ResponseSuccess(ctx, gin.H{
		"result_list": resultList,
	}, server.Success)
}

func UpdateBiologyPictureController(ctx *gin.Context) {
	biologyIdString := ctx.PostForm("BiologyId")
	biologyId, errAtoi := strconv.Atoi(biologyIdString)
	if errAtoi != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "atoi err")
		return
	}
	// auth
	farmhouseId := dao.GetBiologyInfoById(uint(biologyId)).FarmhouseID
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	if !service.AuthCompanyUser(user.ID, farmhouseId) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}

	// 将已保存的图片在服务器本地删除
	oldPicturePath := service.GetBiologyPicturePathService(uint(biologyId))
	go os.Remove(oldPicturePath)

	// 将图片保存到本地，并将图片路径传给service层写入数据库
	file, _ := ctx.FormFile("BiologyPicture")
	fileExt := strings.ToLower(path.Ext(file.Filename))
	if fileExt != ".png" && fileExt != ".jpg" {
		server.Response(ctx, http.StatusBadRequest, 400, nil, "上传文件的格式错误")
		return
	}
	pictureDir := viper.GetString("biology.picturedir")
	picturePath := pictureDir + file.Filename
	errUpload := ctx.SaveUploadedFile(file, picturePath)
	if errUpload != nil {
		panic(errUpload.Error())
	}
	service.UpdateBiologyPicturePathService(uint(biologyId), picturePath)
	server.ResponseSuccess(ctx, nil, server.Success)
}

func GetBiologyPictureController(ctx *gin.Context) {
	biologyIdString := ctx.Query("BiologyId")
	biologyId, errAtoi := strconv.Atoi(biologyIdString)
	if errAtoi != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "atoi err")
		return
	}
	// auth
	farmhouseId := dao.GetBiologyInfoById(uint(biologyId)).FarmhouseID
	fmt.Println(farmhouseId)
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	if (!service.AuthCompanyUser(user.ID, uint(farmhouseId))) && (!service.AuthVisitor(user.ID, uint(farmhouseId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	picture := service.GetBiologyPictureService(uint(biologyId))
	server.ResponseSuccess(ctx, gin.H{"picture": picture}, server.Success)
}

func GetBiologyPicturePathController(ctx *gin.Context) {
	biologyIdString := ctx.Query("BiologyId")
	biologyId, errAtoi := strconv.Atoi(biologyIdString)
	if errAtoi != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "atoi err")
		return
	}
	// auth
	farmhouseId := dao.GetBiologyInfoById(uint(biologyId)).FarmhouseID
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	if (!service.AuthCompanyUser(user.ID, uint(farmhouseId))) && (!service.AuthVisitor(user.ID, uint(farmhouseId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}

	picturePath := service.GetBiologyPicturePathService(uint(biologyId))
	server.ResponseSuccess(ctx, gin.H{"picture_path": picturePath}, server.Success)
}

func GetBiologyAuthListController(ctx *gin.Context) {
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	result := service.GetAuthBiologyListService(user.ID)
	server.ResponseSuccess(ctx, gin.H{"biology_list": result}, server.Success)
}

func GetOwnBiologyListController(ctx *gin.Context) {
	userInfo, exists := ctx.Get("user")
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user infromation does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	biologyList := service.GetOwnBiologyListService(user.ID)
	server.ResponseSuccess(ctx, gin.H{"biology_list": biologyList}, server.Success)
}