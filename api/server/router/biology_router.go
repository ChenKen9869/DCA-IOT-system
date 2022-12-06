package router

import (
	"go-backend/api/server/controller"
	"go-backend/api/common/middleware"
	"github.com/gin-gonic/gin"
)

func BiologyRouter(r *gin.Engine) *gin.Engine {
	biology := r.Group("/biology")
	biology.Use(middleware.AuthMiddleware())
	
	biology.POST("/create", controller.CreateBiologyController)
	biology.POST("/create_type", controller.CreateBiologyTypeController)
	biology.POST("/update_picture", controller.UpdateBiologyPictureController)
	biology.POST("/create_epidemic_prevention_record", controller.CreateEpidemicPreventionRecordController)
	biology.POST("/create_operation_record", controller.CreateOperationRecordController)
	biology.POST("/create_medical_record", controller.CreateMedicalRecordController)

	biology.GET("/get_list", controller.GetBiologyListController)
	biology.GET("/get_with_device_list", controller.GetBiologyWithDeviceListController)
	biology.GET("/get_picture", controller.GetBiologyPictureController)
	biology.GET("/get_picture_path", controller.GetBiologyPicturePathController)
	biology.GET("/get_info", controller.GetBiologyInfoController)
	biology.GET("/get_auth_list", controller.GetBiologyAuthListController)
	biology.GET("/own_list", controller.GetOwnBiologyListController)
	biology.GET("/get_epidemic_prevention_record_list", controller.GetEpidemicPreventRecordListController)
	biology.GET("/get_operation_record_list", controller.GetOperationRecordListController)
	biology.GET("/get_medical_record_list", controller.GetMedicalRecordListController)
	biology.GET("/get_statistic", controller.GetBiologyStatisticController)

	biology.PUT("/update_farmhouse", controller.UpdateBiologyFarmhouseController)

	biology.DELETE("/delete", controller.DeleteBiologyController)
	biology.DELETE("/delete_type", controller.DeleteBiologyTypeController)

	return r
}