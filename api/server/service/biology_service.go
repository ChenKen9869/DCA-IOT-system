package service

import (
	"go-backend/api/server/dao"
	"go-backend/api/server/entity"
	"go-backend/api/server/tools/util"
	"go-backend/api/server/vo"
	"os"
	"path/filepath"
	"time"
)

// @Summary API of golang gin backend
// @Tags Biology
// @description create biology : 创建一个生物 参数列表：[生物名称、生物类别、该生物所在的牧舍ID、出生日期、性别] 访问携带token
// @version 1.0
// @accept mpfd
// @param BiologyName formData string true "biology name"
// @param BiologyType formData string true "biology type"
// @param CompanyId formData string true "company id(farmhouse id)"
// @param Birthday formData string true "biology birthday"
// @param Gender formData string true "biology gender"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /biology/create [post]
func CreateBiologyService(biologyName string, farmhouseId uint, biologyTypeId string, birthday time.Time, gender string, owner uint) uint {
	biology := entity.Biology{
		Name: biologyName,
		FarmhouseID: farmhouseId,
		BiologyTypeID: biologyTypeId,
		Owner: owner,
		Birthday: birthday,
		Gender: gender,
	}
	id := dao.CreateBiology(biology)
	// 判断 type 是否在服务器中注册过
	if !dao.ExistBiologyType(biologyTypeId) {
		newBiologyType := entity.BiologyType{
			BiologyTypeID: biologyTypeId,
		}
		dao.CreateBiologyType(newBiologyType)
	}
	return id
}

// @Summary API of golang gin backend
// @Tags Biology
// @description delete biology : 删除一个生物 参数列表：[本次删除的操作人员姓名，操作人员的联系方式，生物的去处（病死，屠宰场。卖出 等），生物ID] 访问携带token
// @version 1.0
// @accept application/json
// @param Operator query string true "name of operator"
// @param TelephonNumber query string true "telephone number of operator"
// @param LeavePlace query string true "leave place"
// @param Id query string true "id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router  /biology/delete [delete]
func DeleteBiologyService(operator string, telephoneNumber string, leavePlace string, biologyId uint) {
	portableDeviceList := dao.GetPortableDeviceListByBiology(biologyId)
	for _, portableDevice := range portableDeviceList {	
		dao.DeletePortableDevice(portableDevice.ID)
	}
	dao.DeleteBiology(biologyId)
	biologyChangeRecord := entity.BiologyChange {
		BiologyId: biologyId,
		FromCompany: int(dao.GetBiologyInfoById(biologyId).FarmhouseID),
		ToCompany: -1,
		Operator: operator,
		TelephoneNumber: telephoneNumber,
		LeavePlace: leavePlace,
	}
	dao.CreateBiologyChangeRecord(biologyChangeRecord)
}

// @Summary API of golang gin backend
// @Tags Biology
// @description create biology type : 创建生物类型 参数列表：[生物类型名称] 
// @version 1.0
// @accept mpfd
// @param BiologyTypeId formData string true "type name"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /biology/create_type [post]
func CreateBiologyTypeService(biologyTypeId string) {
	biologyType := entity.BiologyType{
		BiologyTypeID: biologyTypeId,
	}
	dao.CreateBiologyType(biologyType)
}

// @Summary API of golang gin backend
// @Tags Biology
// @description delete biology type : 删除生物类型 参数列表：[生物类型名称] 访问携带token
// @version 1.0
// @accept application/json
// @param BiologyTypeId query string true "type name"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /biology/delete_type [delete]
func DeleteBiologyTypeService(biologyTypeId string) {
	dao.DeleteBiologyType(biologyTypeId)
}

// @Summary API of golang gin backend
// @Tags Biology
// @description get all biologies of farmhouse : 通过牧舍ID获取其中的所有生物组成的列表 参数列表：[牧舍ID] 访问携带token
// @version 1.0
// @accept application/json
// @param CompanyId query int true "company id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /biology/get_list [get]
func GetBiologyListService(companyId uint) []entity.Biology {
	var biologyList []entity.Biology
	GetBiologyRecursive(companyId, biologyList)
	return biologyList
}

func GetBiologyRecursive(companyId uint, biologyList []entity.Biology) {
	biologies := dao.GetBiologyListByFarmhouse(companyId)
	biologyList = append(biologyList, biologies...)
	childrenList := dao.GetCompanyListByParent(companyId)
	for _, subCompany := range childrenList {
		GetBiologyRecursive(subCompany.ID, biologyList)
	}
}

// @Summary API of golang gin backend
// @Tags Biology
// @description get all biologies with devices of company : 根据农牧场ID获取其中所有携带有便携式设备的生物所组成的列表（包括每个生物对应的设备信息） 参数列表：[农牧场ID] 访问携带token
// @version 1.0
// @accept application/json
// @param CompanyId query int true "company id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /biology/get_with_device_list [get]
func GetBiologyWithDeviceListService(companyId uint) []vo.BiologyDevice {
	farmhouseList := dao.GetCompanyListByParent(companyId)
	var result []vo.BiologyDevice
	// 遍历farmhouse，查找所有farmhouse中有生物的设备
	for _, farmhouse := range farmhouseList {
		biologyList := dao.GetBiologyListByFarmhouse(farmhouse.ID)
		for _, biology := range biologyList {
			deviceList := dao.GetPortableDeviceListByBiology(biology.ID)
			for _, device := range deviceList {
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
		}
	}
	return result
}

// @Summary API of golang gin backend
// @Tags Biology
// @description update biology farmhouse : 更新生物所属的牧舍（转舍） 参数列表：[本次转舍的操作人员姓名、操作人员联系方式、生物ID、生物的目的牧舍ID] 访问携带token
// @version 1.0
// @accept application/json
// @param Operator formData string true "name of operator"
// @param TelephonNumber formData string true "telephone number of operator"
// @param BiologyId formData int true "biology id"
// @param FarmhouseId formData string true "farmhouse id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /biology/update_farmhouse [put]
func UpdateBiologyFarmhouseService(operator string, telephoneNumber string, biologyId uint, farmhouseId uint) {
	if biologyId == farmhouseId {
		return
	}
	dao.UpdateBiologyFarmhouse(biologyId, farmhouseId)
	biologyChangeRecord := entity.BiologyChange {
		BiologyId: biologyId,
		FromCompany: int(dao.GetBiologyInfoById(biologyId).FarmhouseID),
		ToCompany: int(farmhouseId),
		Operator: operator,
		TelephoneNumber: telephoneNumber,
		LeavePlace: "null",
	}
	dao.CreateBiologyChangeRecord(biologyChangeRecord)
}

// @Summary API of golang gin backend
// @Tags Biology
// @description create biology epidemic prevention record : 新增生物的防疫记录 参数列表：[生物ID、本次使用的疫苗信息记录（疫苗描述信息）、注射时间] 访问携带token
// @version 1.0
// @accept mpfd
// @param BiologyId formData int true "biology id"
// @param VaccineDescription formData string true "vaccine description"
// @param InoculationTime formData string true "inoculation time"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /biology/create_epidemic_prevention_record [post]
func CreateEpidemicPreventionRecordService(biologyId uint, vaccineDescription string, inoculationTime string) {
	dao.CreateEpidemicPreventionRecord(entity.EpidemicPrevention{
		BiologyId: biologyId,
		VaccineDescription: vaccineDescription,
		InoculationTime: util.ParseDate(inoculationTime),
	})
}

// @Summary API of golang gin backend
// @Tags Biology
// @description get epidemic prevention record list of biology : 获取生物的防疫信息记录列表 参数列表：[生物ID] 访问携带token
// @version 1.0
// @accept application/json
// @param BiologyId query int true "biology id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /biology/get_epidemic_prevention_record_list [get]
func GetEpidemicPreventionRecordListService(biologyId uint) []vo.EpidemicPreventionRecord {
	recordList := dao.GetEpidemicPreventionRecordListByBiology(biologyId)
	var resultList []vo.EpidemicPreventionRecord
	for _, record := range recordList {
		resultList = append(resultList, vo.EpidemicPreventionRecord{
			VaccineDescription: record.VaccineDescription,
			InoculationTime: record.InoculationTime,
		})
	}
	return resultList
}

// @Summary API of golang gin backend
// @Tags Biology
// @description create biology operation record : 新增生物的手术记录 参数列表：[生物ID、手术医生、手术时间、过程记录、手术结果] 访问携带token
// @version 1.0
// @accept mpfd
// @param BiologyId formData int true "biology id"
// @param Doctor formData string true "doctor"
// @param OperationTime formData string true "operation time"
// @param ProcessDescription formData string true "process description"
// @param Result formData string true "result"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /biology/create_operation_record [post]
func CreateOperationRecordService(biologyId uint, doctor string, operationTime string, processDescription string, result string) {
	dao.CreateOperationRecord(entity.OperationHistory{
		BiologyId: biologyId,
		Doctor: doctor,
		OperationTime: util.ParseDate(operationTime),
		ProcessDescription: processDescription,
		Result: result,
	})
}

// @Summary API of golang gin backend
// @Tags Biology
// @description get operation record list of biology : 获取生物的手术记录列表 参数列表：[生物ID] 访问携带token
// @version 1.0
// @accept application/json
// @param BiologyId query int true "biology id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /biology/get_operation_record_list [get]
func GetOperationRecordListService(biologyId uint) []vo.OperationRecord {
	recordList := dao.GetOperationRecordListByBiology(biologyId)
	var resultList []vo.OperationRecord
	for _, record := range recordList {
		resultList = append(resultList, vo.OperationRecord{
			Doctor: record.Doctor,
			OperationTime: record.OperationTime,
			ProcessDescription: record.ProcessDescription,
			Result: record.Result,
		})
	}
	return resultList
}

// @Summary API of golang gin backend
// @Tags Biology
// @description create biology medical record : 新增生物的用药记录（新增病历） 参数列表：[生物ID、疾病描述、患病时间、治疗方案] 访问携带token
// @version 1.0
// @accept mpfd
// @param BiologyId formData int true "biology id"
// @param Disease formData string true "disease"
// @param IllnessTime formData string true "illness time"
// @param TreatmentPlan formData string true "treatment plan"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /biology/create_medical_record [post]
func CreateMedicalRecordService(biologyId uint, disease string, illnessTime string, treatmentPlan string) {
	dao.CreateMedicalRecord(entity.MedicalHistory{
		BiologyId: biologyId,
		Disease: disease,
		IllnessTime: util.ParseDate(illnessTime),
		TreatmentPlan: treatmentPlan,
	})
}

// @Summary API of golang gin backend
// @Tags Biology
// @description get medical record list of biology : 获取生物的病历列表 参数列表：[生物ID] 访问携带token
// @version 1.0
// @accept application/json
// @param BiologyId query int true "biology id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /biology/get_medical_record_list [get]
func GetMedicalRecordListService(biologyId uint) []vo.MedicalRecord {
	recordList := dao.GetMedicalRecordListByBiology(biologyId)
	var resultList []vo.MedicalRecord
	for _, record := range recordList {
		resultList = append(resultList, vo.MedicalRecord{
			Disease: record.Disease,
			IllnessTime: record.IllnessTime,
			TreatmentPlan: record.TreatmentPlan,
		})
	}
	return resultList
}

// @Summary API of golang gin backend
// @Tags Biology
// @description update biology picture : 上传（更新）生物的照片 参数列表：[生物ID、照片文件] 访问携带token
// @version 1.0
// @accept application/json
// @param BiologyId formData int true "biology id"
// @param BiologyPicture formData file true "new picture"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /biology/update_picture [post]
func UpdateBiologyPicturePathService(biologyId uint, picturePath string) {
	dao.UpdateBiologyPicturePath(biologyId, picturePath)
}

// @Summary API of golang gin backend
// @Tags Biology
// @description get static picture path of biology : 获取生物的照片（获取生物照片在服务器中的静态资源地址） 参数列表：[生物ID] 访问携带token
// @version 1.0
// @accept application/json
// @param BiologyId query int true "biology id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /biology/get_picture_path [get]
func GetBiologyPicturePathService(biologyId uint) string {
	picturePath := dao.GetBiologyInfoById(biologyId).PicturePath
	_, fileName := filepath.Split(picturePath)
	filePath := "/biology_pictures/" + fileName
	return filePath
}

// @Summary API of golang gin backend
// @Tags Biology
// @description get picture of biology : 获取生物的照片（获取生物照片的 bytes 形式）参数列表：[生物ID] 访问携带token
// @version 1.0
// @accept application/json
// @param BiologyId query int true "biology id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /biology/get_picture [get]
func GetBiologyPictureService(biologyId uint) []byte {
	picturePath := dao.GetBiologyInfoById(biologyId).PicturePath
	picture, errRead := os.ReadFile(picturePath)
	if errRead != nil {
		panic(errRead.Error())
	}
	return picture
}

// @Summary API of golang gin backend
// @Tags Biology
// @description get picture of biology : 获取生物的详细信息 参数列表：[生物ID] 访问携带token
// @version 1.0
// @accept application/json
// @param BiologyId query int true "biology id"
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /biology/get_info [get]
func GetBiologyInfoService(biologyId uint) vo.BiologyInfo {
	biology := dao.GetBiologyInfoById(biologyId)
	if biology.ID == 0 {
		panic("biology does not exist")
	}
	biologyInfo := vo.BiologyInfo {
		Id: biology.ID,
		Name: biology.Name,
		Type: biology.BiologyTypeID,
		Gender: biology.Gender,
		Birthday: biology.Birthday,
		CreateTime: biology.CreatedAt,
		FarmhouseId: biology.FarmhouseID,
	}
	return biologyInfo
}

func getChildNodeRecursive(currentId uint, nodeList *[]uint) {
	chidrenList := dao.GetCompanyListByParent(currentId)
	if len(chidrenList) == 0 {
		*nodeList = append(*nodeList, currentId)
		return
	} else {
		for _, child := range chidrenList {
			getChildNodeRecursive(child.ID, nodeList)
		}
	}
}

// @Summary API of golang gin backend
// @Tags Biology
// @description get user's auth biology list : 获取当前用户有权限的所有生物信息 参数列表：[] 访问携带token
// @version 1.0
// @accept application/json
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /biology/get_auth_list [get]
func GetAuthBiologyListService(userId uint) []vo.AuthBology {
	companies := dao.GetCompanyListByUserID(userId)
	var childNodeList []uint
	for _, company := range companies {
		getChildNodeRecursive(company.CompanyID, &childNodeList)
	}
	var result []vo.AuthBology
	for _, node := range childNodeList {
		currList := dao.GetBiologyListByFarmhouse(node)
		for _, curr := range currList {
			result = append(result, vo.AuthBology{
				BiologyId: curr.ID,
				BiologyName: curr.Name,
				BiologyType: curr.BiologyTypeID,
				Gender: curr.Gender,
				FarmhouseId: curr.FarmhouseID,
				Birthday: curr.Birthday,
				CreateDate: curr.CreatedAt,
			})
		}
	}
	return result
}

// @Summary API of golang gin backend
// @Tags Biology
// @description get own biology list : 获取当前用户拥有的所有生物信息 参数列表：[] 访问携带token
// @version 1.0
// @accept application/json
// @param Authorization header string true "token"
// @Success 200 {object} server.SuccessResponse200 "成功"
// @router /biology/own_list [get]
func GetOwnBiologyListService(userId uint) []vo.OwnBiology {
	var biologyInfoList []vo.OwnBiology
	biologyList := dao.GetOwnBiologyList(userId)
	for _, biology := range biologyList {
		biologyInfoList = append(biologyInfoList, vo.OwnBiology{
			Id: biology.ID,
			Name: biology.Name,
			Type: biology.BiologyTypeID,
			Gender: biology.Gender,
			Birthday: biology.Birthday,
			CreateTime: biology.CreatedAt,
			FarmhouseId: biology.FarmhouseID,
		})
	}
	return biologyInfoList
}