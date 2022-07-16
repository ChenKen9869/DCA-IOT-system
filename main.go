package main

import (
	"go-backend/common"
	docs "go-backend/docs"
	"go-backend/geocontainer"
	"go-backend/middleware"
	"go-backend/monitor"
	"go-backend/router"
	"go-backend/sensor"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Intelligent Pasture Backend APIs
// @version version(1.0)
// @description golang-backend interface
// @Precautions when using termsOfService specifications

// @contact.name aken
// @contact.url https://github.com/ChenKen9869
// @contact.email 972576519@qq.com

// @license.name license(Mandatory)
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 8.142.115.160:5930
// @BasePath /
func main () {
	InitConfig()
	db := common.InitDB()
	deviceDb := common.InitDeviceDB()
	sensor.InitCollections()
	geocontainer.InitContainer()
	monitor.InitMonitor()
	defer db.Close()
	defer deviceDb.Client().Disconnect(common.Ctx)
	
	r := gin.Default()
	r.Use(middleware.LoggerToFile())
	r.Use(middleware.CORSMiddleware())
	r = router.UserRouter(r)
	r = router.CompanyRouter(r)
	r = router.DeviceRouter(r)
	r = router.BiologyRouter(r)
	r = router.FenceRouter(r)
	r = router.MonitorRouter(r)
	r.StaticFS("/biology_pictures", http.Dir("./pictures"))
	port := viper.GetString("server.port")

	docs.SwaggerInfo.BasePath = ""

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}
