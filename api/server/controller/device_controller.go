package controller

import (
	"go-backend/api/server/dao"
	"go-backend/api/server/entity"
	"go-backend/api/server/service"
	"go-backend/api/server/tools/server"
	"go-backend/api/server/tools/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateFixedDeviceController(ctx *gin.Context) {
	CompanyIdString := ctx.PostForm("CompanyId")
	deviceId := ctx.PostForm("DeviceId")
	typeId := ctx.PostForm("TypeId")
	installTime := ctx.PostForm("InstallTime")
	boughtTime := ctx.PostForm("BoughtTime")
	
	userInfo, exists:= ctx.Get("user")
	user := userInfo.(entity.User)
	if !exists {
		panic("error: user information does not exists in application context")
	}
	companyId, err := strconv.Atoi(CompanyIdString)
	if err != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "server inter failed")
		return
	}	
	if !service.AuthCompanyUser(user.ID, uint(companyId)) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	company := dao.GetCompanyInfoByID(uint(companyId))
	owner := company.Owner
	installDate := util.ParseDate(installTime)
	boughtDate := util.ParseDate(boughtTime)
	if !dao.ExistFixedDeviceType(typeId) {
		server.Response(ctx, http.StatusBadRequest, 400, nil, "不支持的固定设备类型")
		return
	}
	id := service.CreateFixedDeviceService(deviceId, uint(companyId), typeId, owner, installDate, boughtDate)
	server.ResponseSuccess(ctx, gin.H{"Id": id}, server.Success)
}

func DeleteFixedDeviceController(ctx *gin.Context) {
	fixedDeviceIdString := ctx.Query("Id")

	fixedDeviceId, _ := strconv.Atoi(fixedDeviceIdString)
	userInfo, exists:= ctx.Get("user")
	user := userInfo.(entity.User)
	if !exists {
		panic("error: user information does not exists in application context")
	}
	companyId := dao.GetFixedDeviceInfoById(uint(fixedDeviceId)).FarmhouseID
	if !service.AuthCompanyUser(user.ID, uint(companyId)) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	service.DeleteFixedDeviceService(uint(fixedDeviceId))
	server.ResponseSuccess(ctx, nil, server.Success)
}

func CreatePortableDeviceController(ctx *gin.Context) {
	biologyIdString := ctx.PostForm("BiologyId")
	portableDeviceId := ctx.PostForm("DeviceId")
	typeId := ctx.PostForm("TypeId")
	installTime := ctx.PostForm("InstallTime")
	boughtTime := ctx.PostForm("BoughtTime")

	installDate := util.ParseDate(installTime)
	boughtDate := util.ParseDate(boughtTime)
	biologyId, err := strconv.Atoi(biologyIdString)
	if err != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "server inter failed")
		return
	}
	userInfo, exists := ctx.Get("user") 
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user information does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	companyId := dao.GetBiologyInfoById(uint(biologyId)).FarmhouseID
	if !service.AuthCompanyUser(user.ID, uint(companyId)) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	if !dao.ExistPortableDeviceType(typeId) {
		server.Response(ctx, http.StatusBadRequest, 400, nil, "不支持的携带设备类型")
		return
	}
	id := service.CreatePortableDeviceService(uint(biologyId), portableDeviceId, typeId, installDate, boughtDate)
	server.ResponseSuccess(ctx, gin.H{"Id": id}, server.Success)
}

func DeletePortableDeviceController(ctx *gin.Context) {
	portableDeviceIdString := ctx.Query("Id")

	portableDeviceId, _ := strconv.Atoi(portableDeviceIdString) 
	biologyId := dao.GetPortableDeviceInfoById(uint(portableDeviceId)).BiologyID
	companyId := dao.GetBiologyInfoById(biologyId).FarmhouseID
	userInfo, exists := ctx.Get("user") 
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user information does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	if !service.AuthCompanyUser(user.ID, uint(companyId)) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	service.DeletePortableDeviceService(uint(portableDeviceId))
	server.ResponseSuccess(ctx, nil, server.Success)
}

func CreateFixedDeviceTypeController(ctx *gin.Context) {
	fixedDeviceTypeId := ctx.PostForm("FixedDeviceTypeId")

	service.CreateFixedDeviceTypeService(fixedDeviceTypeId)
	server.ResponseSuccess(ctx, nil, server.Success)
}

func DeleteFixedDeviceTypeController(ctx *gin.Context) {
	fixedDeviceTypeId := ctx.Query("FixedDeviceTypeId")

	service.DeleteFixedDeviceTypeService(fixedDeviceTypeId)
	server.ResponseSuccess(ctx, nil, server.Success)
}

func CreatePortableDeviceTypeController(ctx *gin.Context) {
	portableDeviceTypeId := ctx.PostForm("PortableDeviceTypeId")

	service.CreatePortableDeviceTypeService(portableDeviceTypeId)
	server.ResponseSuccess(ctx, nil, server.Success)
}

func DeletePortableDeviceTypeController(ctx *gin.Context) {
	portableDeviceTypeId := ctx.Query("PortableDeviceTypeId")

	service.DeletePortableDeviceTypeService(portableDeviceTypeId)
	server.ResponseSuccess(ctx, nil, server.Success)
}

func GetMonitorStreamController(ctx *gin.Context) {
	deviceId := ctx.Query("Id")

	deviceIdInt, _ := strconv.Atoi(deviceId)
	if len(deviceId) < 1 {
		server.Response(ctx, http.StatusBadRequest, 400, nil, "monitorId is empty")
		return
	}
	userInfo, exists := ctx.Get("user") 
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user information does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	companyId := dao.GetFixedDeviceInfoById(uint(deviceIdInt)).FarmhouseID
	if (!service.AuthCompanyUser(user.ID, uint(companyId))) && (!service.AuthVisitor(user.ID, uint(companyId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	deviceType := dao.GetFixedDeviceInfoById(uint(deviceIdInt)).FixedDeviceTypeID
	if deviceType != "摄像头" {
		server.Response(ctx, http.StatusUnauthorized, 403, nil, "设备类型错误")
		return	
	}
	resultUrl, resultExpireTime, id, msg, accessToken:= service.GetMonitorStreamByDeviceIdService(uint(deviceIdInt))
	payload := gin.H{
		"id" : id,
		"url" : resultUrl,
		"expireTime" : resultExpireTime,
		"msg" : msg, 
		"accessToken": accessToken,
	}
	server.ResponseSuccess(ctx, payload, server.Success)
}

func GetNewCollarRealtimeController(ctx *gin.Context) {
	deviceId := ctx.Query("Id")

	deviceIdInt, _ := strconv.Atoi(deviceId)
	if len(deviceId) < 1 {
		server.Response(ctx, http.StatusBadRequest, 400, nil, "monitorId is empty")
		return
	}
	userInfo, exists := ctx.Get("user") 
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user information does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	deviceInfo := dao.GetPortableDeviceInfoById(uint(deviceIdInt))
	biologyId := deviceInfo.BiologyID
	companyId := dao.GetBiologyInfoById(biologyId).FarmhouseID
	if (!service.AuthCompanyUser(user.ID, uint(companyId))) && (!service.AuthVisitor(user.ID, uint(companyId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	deviceType := deviceInfo.PortableDeviceTypeID
	if deviceType != "中农智联耳标" {
		server.Response(ctx, http.StatusUnauthorized, 403, nil, "设备类型错误")
		return	
	}
	data, msg:= service.GetNewCollarRealtimeByDeviceIdService(uint(deviceIdInt))
	payload := gin.H{
		"data": data,
		"msg": msg,
	}
	server.ResponseSuccess(ctx, payload, server.Success)
}

func GetLatestFioController(ctx *gin.Context) {
	fioIdString := ctx.Query("Id")

	fioId, _ := strconv.Atoi(fioIdString)
	userInfo, exists := ctx.Get("user") 
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user information does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	companyId := dao.GetFixedDeviceInfoById(uint(fioId)).FarmhouseID
	if (!service.AuthCompanyUser(user.ID, uint(companyId))) && (!service.AuthVisitor(user.ID, uint(companyId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	data := service.GetLatestFioService(uint(fioId))
	server.ResponseSuccess(ctx, gin.H{"latest": data}, server.Success)
}

func GetFioListByTime(ctx *gin.Context) {
	fioIdString := ctx.Query("Id")
	startTime := ctx.Query("StartTime")
	endTime := ctx.Query("EndTime")

	fioId, _ := strconv.Atoi(fioIdString)
	userInfo, exists := ctx.Get("user") 
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user information does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	companyId := dao.GetFixedDeviceInfoById(uint(fioId)).FarmhouseID
	if (!service.AuthCompanyUser(user.ID, uint(companyId))) && (!service.AuthVisitor(user.ID, uint(companyId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	result := service.GetFioRecordListByTimeService(uint(fioId), startTime, endTime)
	server.ResponseSuccess(ctx, gin.H{"recordList": result}, server.Success)
}

func GetFixedDeviceListByFarmhouseController(ctx *gin.Context) {
	companyIdString := ctx.Query("CompanyId")

	companyId, errAtoiComanyId := strconv.Atoi(companyIdString)
	if errAtoiComanyId != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "服务器内部错误")
	}
	userInfo, exists := ctx.Get("user")
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user infromation does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	if (!service.AuthCompanyUser(user.ID, uint(companyId))) && (!service.AuthVisitor(user.ID, uint(companyId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	fixedDeviceList := service.GetFixedDeviceListByFarmhouseService(uint(companyId))
	var result []gin.H
	for _, deviceInfo := range fixedDeviceList {
		result = append(result, gin.H{
			"id": deviceInfo.ID,
			"type": deviceInfo.FixedDeviceTypeID,
			"farmhouse_id": deviceInfo.FarmhouseID,
			"create_date": deviceInfo.CreatedAt,
			"device_id": deviceInfo.DeviceID,
			"bought_time": deviceInfo.BoughtTime,
			"install_time": deviceInfo.InstallTime,
			"stat": deviceInfo.Stat,
		})
	}
	server.ResponseSuccess(ctx, gin.H{"fixedDeviceList": result}, server.Success)
}

func GetPortableDeviceListByFarmhouseController(ctx *gin.Context) {
	companyIdString := ctx.Query("CompanyId")
	
	companyId, errAtoiComanyId := strconv.Atoi(companyIdString)
	if errAtoiComanyId != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, errAtoiComanyId.Error())
		return
	}
	userInfo, exists := ctx.Get("user")
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user infromation does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	if (!service.AuthCompanyUser(user.ID, uint(companyId))) && (!service.AuthVisitor(user.ID, uint(companyId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	result := service.GetPortableDeviceListByFarmhouseService(uint(companyId))
	server.ResponseSuccess(ctx, gin.H{"portable_device_list": result}, server.Success)
}

func GetPortableDeviceListByBiologyController(ctx *gin.Context) {
	biologyIdString := ctx.Query("BiologyId")

	biologyId, errAtoi := strconv.Atoi(biologyIdString)
	if errAtoi != nil {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, errAtoi.Error())
		return
	}
	userInfo, exists := ctx.Get("user")
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user infromation does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	companyId := dao.GetBiologyInfoById(uint(biologyId)).FarmhouseID
	if (!service.AuthCompanyUser(user.ID, uint(companyId))) && (!service.AuthVisitor(user.ID, uint(companyId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	result := service.GetPortableDeviceListByBiologyService(uint(biologyId))
	server.ResponseSuccess(ctx, gin.H{"portable_device_list": result}, server.Success)
}

func GetFixedDeviceAuthListController(ctx *gin.Context) {
	userInfo, exists := ctx.Get("user")
	if !exists {
		panic("error: user information does not exists in application context")
	}
	user := userInfo.(entity.User)
	result := service.GetAuthFixedDeviceListService(user.ID)
	server.ResponseSuccess(ctx, gin.H{"fixed_device_list": result}, server.Success)
}

func GetOwnFixedDeviceListController(ctx *gin.Context) {
	userInfo, exists := ctx.Get("user")
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user infromation does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	fixedDeviceList := service.GetOwnFixedDeviceListService(user.ID)
	server.ResponseSuccess(ctx, gin.H{"fixed_device_list": fixedDeviceList}, server.Success)
}

func GetOwnPortableListController(ctx *gin.Context) {
	userInfo, exists := ctx.Get("user")
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user infromation does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	portableDeviceList := service.GetOwnPortableDeviceListService(user.ID)
	server.ResponseSuccess(ctx, gin.H{"portable_device_list": portableDeviceList}, server.Success)
}

func GetLatestPosCollarController(ctx *gin.Context) {
	fioIdString := ctx.Query("Id")

	fioId, _ := strconv.Atoi(fioIdString)
	userInfo, exists := ctx.Get("user") 
	if !exists {
		server.Response(ctx, http.StatusInternalServerError, 500, nil, "user information does not exists in application context")
		return
	}
	user := userInfo.(entity.User)
	companyId := dao.GetFixedDeviceInfoById(uint(fioId)).FarmhouseID
	if (!service.AuthCompanyUser(user.ID, uint(companyId))) && (!service.AuthVisitor(user.ID, uint(companyId))) {
		server.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
		return
	}
	data := service.GetLatestPosCollarService(uint(fioId))
	server.ResponseSuccess(ctx, gin.H{"latest": data}, server.Success)
}