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

type PositionCollarMessage struct {
	DeviceId  string
	Latitude  float64
	Longitude float64
	Altitude  float64
	Time      time.Time `bson:"time"`
}

func GetLatestDataListPosCollar(deviceId string, nums int64) []PositionCollarMessage {
	var results []PositionCollarMessage
	if nums <= 0 {
		panic("param wrong, got nums: 0")
	}
	filter := bson.D{primitive.E{Key: "deviceid", Value: deviceId}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.Find().SetSort(bson.D{primitive.E{Key: "_id", Value: -1}}).SetLimit(nums)
	cursor, err := db.GetDeviceDB().Collection(posCollarCollection).Find(ctx, filter, opts)
	if err != nil {
		panic(err.Error())
	}
	cursor.All(context.TODO(), &results)
	return results
}

func GetRecordListByTimePosCollar(deviceId string, startTime string, endTime string) []PositionCollarMessage {
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
	var results []PositionCollarMessage
	cursor, err := db.GetDeviceDB().Collection(posCollarCollection).Find(ctx, filter, opts)
	if err != nil {
		panic(err.Error())
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err.Error())
	}
	return results
}

func GetLatestDataPosCollar(deviceId string) PositionCollarMessage {
	var result PositionCollarMessage

	filter := bson.D{primitive.E{Key: "deviceid", Value: deviceId}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "_id", Value: -1}})
	err := db.GetDeviceDB().Collection(posCollarCollection).FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		panic(err.Error())
	}
	return result
}
