package service

import (
	"go-backend/api/server/dao"
	"go-backend/api/server/vo"
	"go-backend/api/sys/iot/sensor"
)

func GetFioHistoryService(fioId uint, startTime string, endTime string) []vo.FioData {
	deviceId := dao.GetFixedDeviceInfoById(fioId).DeviceID
	messageList := sensor.GetRecordListByTimeFio(deviceId, startTime, endTime)
	dataList := []vo.FioData{}
	for _, message := range messageList {
		dataList = append(dataList, vo.FioData{
			Id:          fioId,
			Humidity:    message.Humidity,
			Temperature: message.Temperature,
			Methane:     message.Methane,
			Ammonia:     message.Ammonia,
			Hydrogen:    message.Hydrogen,
			Time:        message.Time,
		})
	}
	return dataList
}

// @Summary API of golang gin backend
// @Tags Device-fixed
// @description get latest five-in-one device information
// @version 1.0
// @accept application/json
// @param Id query string true "Id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "success"
// @router /device/fixed/get_fio_latest [get]
func GetLatestFioService(fioId uint) vo.FioData {
	deviceId := dao.GetFixedDeviceInfoById(fioId).DeviceID
	message := sensor.GetLatestDataFio(deviceId)
	data := vo.FioData{
		Id:          fioId,
		Humidity:    message.Humidity,
		Temperature: message.Temperature,
		Methane:     message.Methane,
		Ammonia:     message.Ammonia,
		Hydrogen:    message.Hydrogen,
		Time:        message.Time,
	}
	return data
}

// @Summary API of golang gin backend
// @Tags Device-fixed
// @description get five-in-one information within time period
// @version 1.0
// @accept application/json
// @param Id query string true "id"
// @param StartTime query string true "start time"
// @param EndTime query string true "end time"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "success"
// @router  /device/fixed/get_fio_list_by_time [get]
func GetFioRecordListByTimeService(fioId uint, startTime string, endTime string) []vo.FioData {
	var result []vo.FioData
	deviceId := dao.GetFixedDeviceInfoById(fioId).DeviceID
	messageList := sensor.GetRecordListByTimeFio(deviceId, startTime, endTime)
	for _, message := range messageList {
		data := vo.FioData{
			Id:          fioId,
			Humidity:    message.Humidity,
			Temperature: message.Temperature,
			Methane:     message.Methane,
			Ammonia:     message.Ammonia,
			Hydrogen:    message.Hydrogen,
			Time:        message.Time,
		}
		result = append(result, data)
	}

	return result
}

// @Summary API of golang gin backend
// @Tags Device-portable
// @description get latest position collar information
// @version 1.0
// @accept application/json
// @param Id query string true "Id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "success"
// @router /device/portable/get_position_collar_latest [get]
func GetLatestPosCollarService(posCollarId uint) vo.PosCollarData {
	deviceId := dao.GetPortableDeviceInfoById(posCollarId).DeviceID
	message := sensor.GetLatestDataPosCollar(deviceId)
	data := vo.PosCollarData{
		Id:        posCollarId,
		Latitude:  message.Latitude,
		Longitude: message.Longitude,
		Altitude:  message.Altitude,
		Time:      message.Time,
	}
	return data
}
