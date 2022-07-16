package monitor

import (
	"go-backend/service"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

type ActiveFence struct {
	Job *cron.Cron
	Timeout *time.Timer
}

var ActiveFenceList map[uint]ActiveFence

// @Summary API of golang gin backend
// @Tags Fence
// @description fence create
// @version 1.0
// @accept mpfd
// @param Position formData string true "position"
// @param DeviceList formData string true "device list"
// @param Duration formData string true "duration"
// @param Coordinate formData string true "coordinate"
// @param Name formData string true "name of fence"
// @param ParentId formData string true "parent id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /fence/create [post]
func StartFenceJob(userId uint, position string, deviceList string, duration int, parentId uint, name string, coordinate string) uint {
	fenceId := service.CreateFenceService(position, deviceList, duration, parentId, name, coordinate)
	spec := viper.GetString("fence.spec")
	c := cron.New()
	c.AddFunc("@every " + spec, func() {
		vitalAbnormalList, positionAbnormalList := service.MonitorFenceService(fenceId)
		if len(vitalAbnormalList) == 0 && len(positionAbnormalList) == 0 {
			return
		} else {
			service.AddAlarmTimeService(fenceId, vitalAbnormalList, positionAbnormalList)
			// 如果存在监控连接，则推送错误消息
			if _, ok := MonitorCentor[userId]; ok {
				message := "01" + vitalAbnormalList + "#" + positionAbnormalList
				msg := MakeMessage(FenceJob, fenceId, message)
				MonitorCentor[userId].MessageChan <- msg
			}
		}
	})
	timeout := time.AfterFunc(time.Duration(duration) * time.Minute, func ()  {
		// 停止定时任务
		c.Stop()
		delete(ActiveFenceList, fenceId)
		service.UpdateFenceToFinishedStat(fenceId)
	})
	ActiveFenceList[fenceId] = ActiveFence {
		Job: c,
		Timeout: timeout,
	}
	c.Start()
	return fenceId
}

// @Summary API of golang gin backend
// @Tags Fence
// @description fence stop
// @version 1.0
// @accept application/json
// @param FenceId query int true "fence id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /fence/stop [delete]
func StopFenceJob(fenceId uint) {
	if _, ok := ActiveFenceList[fenceId]; !ok {
		return
	}
	job := ActiveFenceList[fenceId].Job
	timeout := ActiveFenceList[fenceId].Timeout
	job.Stop()
	timeout.Stop()
	delete(ActiveFenceList, fenceId)
	service.AbortFenceService(fenceId)
}