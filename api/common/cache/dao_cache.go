package cache

import (
	"context"
	"go-backend/api/common/common"
	"math/rand"
	"time"
)

func getRandomExpireTime() time.Duration {
	rand.Seed(time.Now().UnixNano())
	return time.Duration(rand.Intn(100)*int(time.Hour))
}

type CacheUnit struct {
	Key string
	Value interface{}
}

var ctx = context.Background()

func (cacheUnit *CacheUnit)Set(key string, value interface{}) {
	// set redis的 key value
	common.GetRedis().Set(ctx, key, value, getRandomExpireTime())
}

func (CacheUnit *CacheUnit)Get(key string) interface{} {
	value, err := common.GetRedis().Get(ctx, key).Result()
	if err != nil {
		panic(err.Error())
	}
	return value
}

