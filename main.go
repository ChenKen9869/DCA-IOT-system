package main

import (
	"go-backend/common"
	"go-backend/controller"
	"go-backend/model"
	"os"
	docs "go-backend/docs"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
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

// @host localhost:8081
// @BasePath /
func main () {
	InitConfig()
	db := common.InitDB()
	defer db.Close()

	db.AutoMigrate(&model.User{})
	r := gin.Default()
	r = controller.UserController(r)
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
