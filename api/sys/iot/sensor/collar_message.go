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
type CollarMessage struct {
	Head          string
	Edition       string
	DeviceId      string
	Session       string
	Command       string
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

func GetLatestDataListCollar(deviceId string, nums int64) []CollarMessage {
	var results []CollarMessage
	if nums <= 0 {
		panic("param wrong, got nums: 0")
	}
	filter := bson.D{primitive.E{Key: "deviceid", Value: deviceId}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.Find().SetSort(bson.D{primitive.E{Key: "session", Value: -1}}).SetLimit(nums)
	cursor, err := db.GetDeviceDB().Collection(collarCollection).Find(ctx, filter, opts)
	if err != nil {
		return []CollarMessage{}
	}
	cursor.All(context.TODO(), &results)
	return results
}

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
	cursor, err := db.GetDeviceDB().Collection(collarCollection).Find(ctx, filter, opts)
	if err != nil {
		return []CollarMessage{}
	}
	cursor.All(context.TODO(), &results)
	return results
}

func GetLatestDataCollar(deviceId string) CollarMessage {
	var result CollarMessage

	filter := bson.D{primitive.E{Key: "deviceid", Value: deviceId}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "session", Value: -1}})
	err := db.GetDeviceDB().Collection(collarCollection).FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		return CollarMessage{}
	}
	return result
}
