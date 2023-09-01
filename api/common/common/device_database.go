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
var deviceDBUserName string
var deviceDBPassword string

func InitDeviceDB() *mongo.Database {
	deviceDBHost = viper.GetString("mongodb.host")
	deviceDBPort = viper.GetString("mongodb.port")
	deviceDatabase = viper.GetString("mongodb.database")
	deviceDBUserName = viper.GetString("mongodb.username")
	deviceDBPassword = viper.GetString("mongodb.password")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opt := options.Client().ApplyURI("mongodb://" + deviceDBUserName + ":" + deviceDBPassword + "@" + deviceDBHost + ":" + deviceDBPort)
	Ctx = ctx
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		fmt.Println(err.Error())
	}
	DeivceDB = client.Database(deviceDatabase)

	fmt.Println("[INITIAL SUCCESS] The message database module is initialized successfully!")
	return DeivceDB
}

func GetDeviceDB() *mongo.Database {
	return DeivceDB
}
