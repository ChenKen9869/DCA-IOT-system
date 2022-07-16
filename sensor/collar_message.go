package sensor

import (
	"context"
	"go-backend/common"
	"go-backend/util"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type CollarMessage struct {
	Head          string
	Edition       string
	DeviceId      string
	Session       string
	Commond       string
	MessageLength string
	Behavior      string
	Temperature   float32
	Longitude     float64
	Latitude      float64
	Signal        float32
	Voltage       float32
	End           string
	Time          time.Time `bson:"time"`
}

// 根据 设备id 和 消息数量 N，获得最新的 N 条数据
func GetLatestDataListCollar(deviceId string, nums int64) []CollarMessage {
	var results []CollarMessage
	if nums <= 0 {
		panic("param wrong, got nums: 0")
	}
	filter := bson.D{primitive.E{Key: "deviceid", Value: deviceId}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.Find().SetSort(bson.D{primitive.E{Key: "session", Value: -1}}).SetLimit(nums)
	cursor, err := common.GetDeviceDB().Collection(collarCollection).Find(ctx, filter, opts)
	if err != nil {
		return []CollarMessage{}
	}
	cursor.All(context.TODO(), &results)
	return results
}

// 根据 设备id 和 接收时间段， 获取全部记录， 按照 session 排序，从小到大
func GetRecordListByTimeCollar(deviceId string, startTime string, endTime string) []CollarMessage {
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
	opts := options.Find().SetSort(bson.D{primitive.E{Key: "session", Value: 1}})
	var results []CollarMessage
	cursor, err := common.GetDeviceDB().Collection(collarCollection).Find(ctx, filter, opts)
	if err != nil {
		return []CollarMessage{}
	}
	cursor.All(context.TODO(), &results)
	return results
}

// 根据 设备id 获取最新记录
func GetLatestDataCollar(deviceId string) CollarMessage {
	var result CollarMessage

	filter := bson.D{primitive.E{Key: "deviceid", Value: deviceId}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "session", Value: -1}})
	err := common.GetDeviceDB().Collection(collarCollection).FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		return CollarMessage{}
	}
	return result
}
