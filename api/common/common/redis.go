package common

import (
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

var RedisAddr string
var RedisPassword string
var RedisDB int
var RedisClient *redis.Client

func InitRedis() {
	RedisAddr = viper.GetString("")
	RedisPassword = viper.GetString("")
	RedisDB = viper.GetInt("")
	rdb := redis.NewClient(&redis.Options{
        Addr:     RedisAddr,
        Password: RedisPassword, 
        DB:       RedisDB, 
    })
	RedisClient = rdb
}

func GetRedis() *redis.Client {
	return RedisClient
}