package service

import (
	"go-backend/dao"
	"go-backend/entity"
	"go-backend/server"
	"go-backend/vo"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

// @Summary API of golang gin backend
// @Tags Device-fixed
// @description create fixed device
// @version 1.0
// @accept mpfd
// @param CompanyId formData string true "company id"
// @param DeviceId formData string true "device id"
// @param TypeId formData string true "type name"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /device/fixed/create [post]
func CreateFixedDeviceService(deviceId string, farmhouseId uint, fixedDeviceTypeId string, owner uint) uint {
	fixedDevice := entity.FixedDevice{
		DeviceID: deviceId,
		FarmhouseID: farmhouseId,
		FixedDeviceTypeID: fixedDeviceTypeId,
		Owner: owner,
	}
	id := dao.CreateFixedDevice(fixedDevice)
	return id
}

// @Summary API of golang gin backend
// @Tags Device-fixed
// @description delete fixed device
// @version 1.0
// @accept application/json
// @param Id query int true "id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /device/fixed/delete [delete]
func DeleteFixedDeviceService(fixedDeviceId uint) {
	dao.DeleteFixedDevice(fixedDeviceId)
}

func UpdateFixedDeviceService(deviceId string){

}

// @Summary API of golang gin backend
// @Tags Device-portable
// @description create portable device
// @version 1.0
// @accept mpfd
// @param BiologyId formData string true "biology id"
// @param DeviceId formData string true "device id"
// @param TypeId formData string true "type name"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /device/portable/create [post]
func CreatePortableDeviceService(portableDeviceId string, biologyId uint, portableDeviceTypeId string) uint {
	biology := dao.GetBiologyInfoById(biologyId)
	if biology.ID == 0 {
		panic("biology id does not exists")
	}
	owner := biology.Owner
	portableDevice := entity.PortableDevice{
		DeviceID: portableDeviceId,
		BiologyID: biologyId,
		PortableDeviceTypeID: portableDeviceTypeId,
		Owner: owner,
	}
	id := dao.CreatePortableDevice(portableDevice)
	return id
}

// @Summary API of golang gin backend
// @Tags Device-portable
// @description delete portable device
// @version 1.0
// @accept application/json
// @param Id query int true "id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /device/portable/delete [delete]
func DeletePortableDeviceService(portableDeviceId uint) {
	dao.DeletePortableDevice(portableDeviceId)
}

// @Summary API of golang gin backend
// @Tags Device-fixed
// @description create fixed device type
// @version 1.0
// @accept mpfd
// @param FixedDeviceTypeId formData string true "type name"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /device/fixed/create_type [post]
func CreateFixedDeviceTypeService(fixedDeviceTypeId string) {
	fixedDeviceType := entity.FixedDeviceType{
		FixedDeviceTypeID: fixedDeviceTypeId,
	}
	dao.CreateFixedDeviceType(fixedDeviceType)
}

// @Summary API of golang gin backend
// @Tags Device-fixed
// @description delete fixed device type
// @version 1.0
// @accept application/json
// @param FixedDeviceTypeId query int true "type name"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /device/fixed/delete_type [delete]
func DeleteFixedDeviceTypeService(fixedDeviceTypeId string) {
	dao.DeleteFixedDeviceType(fixedDeviceTypeId)
}

// @Summary API of golang gin backend
// @Tags Device-portable
// @description create portable device type
// @version 1.0
// @accept mpfd
// @param PortableDeviceTypeId formData string true "type name"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router  /device/portable/create_type [post]
func CreatePortableDeviceTypeService(portableDeviceTypeId string) {
	portableDeviceType := entity.PortableDeviceType{
		PortableDeviceTypeID: portableDeviceTypeId,
	}
	dao.CreatePortableDeviceType(portableDeviceType)
}

// @Summary API of golang gin backend
// @Tags Device-portable
// @description delete portable device type
// @version 1.0
// @accept application/json
// @param PortableDeviceTypeId query int true "type name"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /device/portable/delete_type [delete]
func DeletePortableDeviceTypeService(fixedDeviceTypeId string) {
	dao.DeleteFixedDeviceType(fixedDeviceTypeId)
}

func updateMonitorStreamToken(streamAccessToken map[string]interface{}) {
	appKey := viper.GetString("monitor.AppKey")
	secret := viper.GetString("monitor.Secret")
	url := accessMonitorTokenUrl
	contentType := "application/x-www-form-urlencoded"
	payload := strings.NewReader("appKey=" + appKey + "&appSecret=" + secret)
	response, err := http.Post(url, contentType, payload)
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()
	responseBody := server.GetResponseBody(response)
	code := responseBody["code"]
	if code != "200" {
		panic("getStreamFailed")
	}
	data := responseBody["data"].(map[string]interface{})
	streamAccessToken["accessToken"] = data["accessToken"].(string)
	streamAccessToken["expireTime"] = data["expireTime"].(int64)

}

var StreamAccessToken = map[string]interface{} {
	"accessToken" : "null",
	"expireTime" : int64(0),
}

// @Summary API of golang gin backend
// @Tags Device-fixed
// @description get monitor streaming address
// @version 1.0
// @accept application/json
// @param Id query string true "id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /device/fixed/get_monitor [get]
func GetMonitorStreamByDeviceIdService(deviceId uint) (string, string, string){
	accessToken := StreamAccessToken["accessToken"].(string)
	expireTime := StreamAccessToken["expireTime"].(int64)
	if accessToken == "null" || 
		expireTime <= int64(99999999) {
			updateMonitorStreamToken(StreamAccessToken)
		}
	monitorDeviceId := dao.GetFixedDeviceInfoById(deviceId).DeviceID
	serverUrl := getStreamAddressUrl
	contentType := "application/x-www-form-urlencoded"
	payload := strings.NewReader("accessToken=" + accessToken + "&deviceSerial=" + monitorDeviceId)
	response, err := http.Post(serverUrl, contentType, payload)
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()
	responseBody := server.GetResponseBody(response)
	code := responseBody["code"].(string)
	if code != "200" {
		panic("error occurs during get streaming address from server by deviceID ")
	}
	data := responseBody["data"].(map[string]interface{})
	resultUrl := data["url"].(string)
	resultExpireTime := data["expireTime"].(string)
	id := data["id"].(string)
	return resultUrl, resultExpireTime, id
}

var accessMonitorTokenUrl string
var getStreamAddressUrl string

func init() {
	accessMonitorTokenUrl = viper.GetString("monitor.AccessMonitorTokenUrl")
	getStreamAddressUrl = viper.GetString("monitor.GetStreamAddressUrl")
}

// @Summary API of golang gin backend
// @Tags Device-fixed
// @description get all fixed devices by farmhouse
// @version 1.0
// @accept application/json
// @param CompanyId query string true "company id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /device/fixed/get_by_farmhouse [get]
func GetFixedDeviceListByFarmhouseService(farmhouseId uint) []vo.FixedDeviceVo {
	var fixedDeviceVoList []vo.FixedDeviceVo
	fixedDeviceList := dao.GetFixedDeviceListByFarmhouse(farmhouseId)
	for _, fixedDevice := range fixedDeviceList {
		fixedDeviceVoList = append(fixedDeviceVoList, vo.FixedDeviceVo{
			Id: fixedDevice.ID,
			Type: fixedDevice.FixedDeviceTypeID,
		})
	}
	return fixedDeviceVoList
}

// @Summary API of golang gin backend
// @Tags Device-portable
// @description get all portable devices by farmhouse
// @version 1.0
// @accept application/json
// @param CompanyId query string true "company id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /device/portable/get_by_farmhouse [get]
func GetPortableDeviceListByFarmhouseService(farmhouseId uint) []vo.BiologyDevice {
	var result []vo.BiologyDevice
	biologyList := dao.GetBiologyListByFarmhouse(farmhouseId)
	for _, biology := range biologyList {
		device := dao.GetPortableDeviceInfoByBiology(biology.ID)
		if device.ID != 0 {
			result = append(result, vo.BiologyDevice{
				BiologyId: biology.ID,
				BiologyName: biology.Name,
				BiologyType: biology.BiologyTypeID,
				DeviceId: device.ID,
				DeviceType: device.PortableDeviceTypeID,
			})
		}
	}
	return result
}