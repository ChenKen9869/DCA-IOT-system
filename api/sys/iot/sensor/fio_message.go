package sensor

import (
	"time"
	"context"
	"fmt"
	"go-backend/api/common/common"
	"go-backend/api/server/tools/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FiveInOneMessage struct {
	Head string
	Edition string
	DeviceType string
	DeviceId string
	Session string
	Commond string
	MessageLength string
	Humidity float32
	Temperature float32
	Methane float32
	Ammonia float32
	Hydrogen float32
	End string
	Time time.Time `bson:"time"`
}

// 根据 设备id 和 消息数量 N，获得最新的 N 条数据
func GetLatestDataListFio(deviceId string, nums int64) []FiveInOneMessage {
	var results []FiveInOneMessage
	if nums <= 0 {
		panic("param wrong, got nums: 0")
	}
	filter := bson.D{primitive.E{Key: "deviceid", Value: deviceId}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.Find().SetSort(bson.D{primitive.E{Key: "_id", Value: -1}}).SetLimit(nums)
	cursor, err := common.GetDeviceDB().Collection(fioCollection).Find(ctx, filter, opts)
	if err != nil {
		fmt.Println(err.Error())
	}
	cursor.All(context.TODO(), &results)
	return results
}

// 根据 设备id 和 接收时间段， 获取全部记录， 按照 session 排序，从小到大
func GetRecordListByTimeFio(deviceId string, startTime string, endTime string) []FiveInOneMessage {
	startAt := util.ParseDate(startTime)
	endAt := util.ParseDate(endTime)
	filter := bson.M{
		"deviceid": deviceId,
		"time": bson.M{
			"$gt": startAt,
			"$lt": endAt,
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.Find().SetSort(bson.D{primitive.E{Key: "_id", Value: -1}})
	var results []FiveInOneMessage
	cursor, err := common.GetDeviceDB().Collection(fioCollection).Find(ctx, filter, opts)
	if err != nil {
		fmt.Println(err.Error())
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err.Error())
	}
	return results
}

// 根据 设备id 获取最新记录
func GetLatestDataFio(deviceId string) FiveInOneMessage {
	var result FiveInOneMessage

	filter := bson.D{primitive.E{Key: "deviceid", Value: deviceId}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "session", Value: -1}})
	err := common.GetDeviceDB().Collection(fioCollection).FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		panic(err.Error())
	}
	return result
}
