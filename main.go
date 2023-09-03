package main

import (
	"fmt"
	"go-backend/api/common/db"
	"go-backend/api/common/middleware"
	"go-backend/api/rule"
	"go-backend/api/server/router"
	"go-backend/api/sys/gis/geo/geocontainer"
	"go-backend/api/sys/iot/monitor"
	"go-backend/api/sys/iot/sensor"
	docs "go-backend/docs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @host localhost:5930
// @BasePath /
func main() {
	InitConfig()
	database := db.InitDB()
	deviceDb := db.InitDeviceDB()
	sensor.InitCollections()
	geocontainer.InitContainer()
	monitor.InitMonitor()
	rule.InitRule()
	defer database.Close()
	defer deviceDb.Client().Disconnect(db.Ctx)

	r := gin.Default()
	r.Use(middleware.LoggerToFile())
	r.Use(middleware.CORSMiddleware())
	r = router.UserRouter(r)
	r = router.CompanyRouter(r)
	r = router.DeviceRouter(r)
	r = router.BiologyRouter(r)
	r = router.MonitorRouter(r)
	r = router.RoleRouter(r)
	r = router.RuleRouter(r)
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
	fmt.Println("[INITIAL SUCCESS] System config is initialized successfully!")
}
