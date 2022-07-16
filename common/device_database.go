package common

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var DeivceDB *mongo.Database
var Ctx context.Context

var deviceDBHost string
var deviceDBPort string
var deviceDatabase string

func InitDeviceDB() *mongo.Database {
	deviceDBHost = viper.GetString("mongodb.host")
	deviceDBPort = viper.GetString("mongodb.port")
	deviceDatabase = viper.GetString("mongodb.database")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	opt := options.Client().ApplyURI("mongodb://" + deviceDBHost + ":" + deviceDBPort)
	// 自带连接池，默认值 100
	// opt.SetMaxPoolSize()
	
	Ctx = ctx
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		fmt.Println(err.Error())
	}
	DeivceDB = client.Database(deviceDatabase)
	return DeivceDB
}

func GetDeviceDB() *mongo.Database {
	return DeivceDB
}