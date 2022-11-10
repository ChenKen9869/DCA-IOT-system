package service

import (
	"go-backend/geoalgorithm"
	"go-backend/geocontainer"
	"go-backend/dao"
	"go-backend/entity"
	"go-backend/sensor"
	"go-backend/util"
	"go-backend/vo"
	"strconv"
	"time"
)

// 结束围栏任务:中止围栏任务
func AbortFenceService(fenceId uint) {
	dao.ModifyFenceStat(fenceId, entity.FenceAbort)
	dao.UpdateFenceEndTime(fenceId, time.Now())
}

// 将围栏状态更新为完成:正常完成围栏任务
func UpdateFenceToFinishedStat(fenceId uint) {
	currentStat := dao.GetFenceRecordById(fenceId).Stat
	if currentStat == entity.FenceFinished {
		return
	}
	dao.ModifyFenceStat(fenceId, entity.FenceFinished)
}
// 新建围栏
func CreateFenceService(position string, deviceList string, duration int, parentId uint, name string, coordinate string) uint {
	owner := dao.GetCompanyInfoByID(parentId).Owner
	fenceRecord := entity.FenceRecord{
		Position: position,
		DeviceList: deviceList,
		StartTime: time.Now(),
		EndTime: time.Now().Add(time.Duration(duration) * time.Minute),
		AlarmTime: 0,
		ParentId: uint(parentId),
		Owner: owner,
		Name: name,
		Coordinate: coordinate,
		Stat: entity.FenceRunning,
	}
	fenceRecordId := dao.CreateFenceRecord(fenceRecord)
	return fenceRecordId
}
// 获取围栏的当前详细情况
// @Summary API of golang gin backend
// @Tags Fence
// @description get fence status : 获取围栏任务的执行状态 参数列表：[围栏ID] 访问携带token
// @version 1.0
// @accept application/json
// @param FenceId query string true "fence id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /fence/get_status [get]
func GetFenceStatService(fenceId uint) (string, string, string, uint, string) {
	fenceRecord := dao.GetFenceRecordById(fenceId)
	coordinate := fenceRecord.Coordinate
	position := fenceRecord.Position
	deviceList := fenceRecord.DeviceList
	alarmTime := fenceRecord.AlarmTime
	endtime := fenceRecord.EndTime
	remainingTime := time.Until(endtime)
	hours := remainingTime.Hours()
	minutes := remainingTime.Minutes()
	seconds := remainingTime.Seconds()
	remain := strconv.Itoa(int(hours)) + "hours, " + strconv.Itoa(int(minutes)) + "minutes, " + strconv.Itoa(int(seconds)) + "seconds, " 
	return coordinate, position, deviceList, alarmTime, remain
}

// 记录报警事件
func AddAlarmTimeService(fenceId uint, vitalAbnormalList string, positionAbnormalList string) {
	dao.AddAlarmTime(fenceId)
	// 将异常设备记录在另外一张表中
}

// 验证围栏是否与设备列表中任意一个设备绑定的生物所在的农舍处于同一个农牧场
func AuthFenceDeviceList(parentId uint, deviceList string) bool {
	devices := util.String2ListUint(deviceList)
	for _, deviceId := range devices {
		deviceInfo := dao.GetPortableDeviceInfoById(deviceId)
		biologyInfo := dao.GetBiologyInfoById(deviceInfo.BiologyID)
		companyInfo := dao.GetCompanyInfoByID(biologyInfo.FarmhouseID)
		if parentId != companyInfo.ParentID {
			return false
		}
	}
	return true
}

// 监控围栏中的设备是否正常（定期监控）
func MonitorFenceService(fenceId uint) (string, string) {
	fenceRecord := dao.GetFenceRecordById(fenceId)
	deviceList := util.String2ListUint(fenceRecord.DeviceList)
	coordinate := fenceRecord.Coordinate

	warningVitList := monitorVitalSigns(deviceList)
	warningPosList := monitorPositions2d(fenceRecord.Position, deviceList, coordinate)

	return util.ListUint2String(warningVitList), util.ListUint2String(warningPosList)
}

// 功能函数 1： 判断围栏是否包含设备, 返回异常的设备列表
func monitorPositions2d(position string, deviceList []uint, coordinate string) []uint {
	var warningPositionList []uint
	var polygon geocontainer.Polygon
	(&polygon).InitFromString(position)
	
	for _, deviceId := range deviceList {
		deviceInfo := dao.GetPortableDeviceInfoById(deviceId)
		sensorMessage := sensor.GetLatestDataCollar(deviceInfo.DeviceID)
			point := geocontainer.Point{
			Longitude: sensorMessage.Longitude,
			Latitude: sensorMessage.Latitude,
		}
		if coordinate == geocontainer.GCJ_02 {
			if !geoalgorithm.PolygonContainsPoint(polygon, point) {
				warningPositionList = append(warningPositionList, deviceId)
			}
		}
	}
	return warningPositionList
}

// 功能函数 2： 判断生命体征是否异常, 返回异常的设备列表
func monitorVitalSigns(deviceList []uint) []uint {
	var warningVitalList []uint
	for _, deviceId := range deviceList {
		deviceInfo := dao.GetPortableDeviceInfoById(deviceId)
		sensorMessage := sensor.GetLatestDataCollar(deviceInfo.DeviceID)
		if sensorMessage.Temperature < 38.0 || sensorMessage.Temperature > 39.5 {
			warningVitalList = append(warningVitalList, deviceId)
		}
	}
	return warningVitalList
}

// 通过companyId获取活跃的围栏列表
// @Summary API of golang gin backend
// @Tags Fence
// @description get active fence list by company id : 获取牧场中所有处于活跃状态的围栏 参数列表：[牧场ID] 访问携带token
// @version 1.0
// @accept application/json
// @param CompanyId query string true "company id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /fence/get_active_list [get]
func GetActiveFenceListByCompanyService(companyId uint) []vo.ActiveFence {
	fenceRecordList := dao.GetActiveFenceListByCompanyId(companyId)
	var activeList []vo.ActiveFence
	for _, fenceRecord := range fenceRecordList {
		activeList = append(activeList, vo.ActiveFence{
			Name: fenceRecord.Name,
			Id: fenceRecord.ID,
			ExpireTime: fenceRecord.EndTime,
			Position: fenceRecord.Position,
		})
	}
	return activeList
}
