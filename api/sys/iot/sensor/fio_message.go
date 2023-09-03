package sensor

import (
	"context"
	"go-backend/api/common/db"
	"go-backend/api/server/tools/util"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FiveInOneMessage struct {
	Head          string
	Edition       string
	DeviceType    string
	DeviceId      string
	Session       string
	Commond       string
	MessageLength string
	Humidity      float32
	Temperature   float32
	Methane       float32
	Ammonia       float32
	Hydrogen      float32
	End           string
	Time          time.Time `bson:"time"`
}

func GetLatestDataListFio(deviceId string, nums int64) []FiveInOneMessage {
	var results []FiveInOneMessage
	if nums <= 0 {
		panic("param wrong, got nums: 0")
	}
	filter := bson.D{primitive.E{Key: "deviceid", Value: deviceId}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.Find().SetSort(bson.D{primitive.E{Key: "_id", Value: -1}}).SetLimit(nums)
	cursor, err := db.GetDeviceDB().Collection(fioCollection).Find(ctx, filter, opts)
	if err != nil {
		panic(err.Error())
	}
	cursor.All(context.TODO(), &results)
	return results
}

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
	cursor, err := db.GetDeviceDB().Collection(fioCollection).Find(ctx, filter, opts)
	if err != nil {
		panic(err.Error())
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err.Error())
	}
	return results
}

func GetLatestDataFio(deviceId string) FiveInOneMessage {
	var result FiveInOneMessage

	filter := bson.D{primitive.E{Key: "deviceid", Value: deviceId}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "session", Value: -1}})
	err := db.GetDeviceDB().Collection(fioCollection).FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		panic(err.Error())
	}
	return result
}
