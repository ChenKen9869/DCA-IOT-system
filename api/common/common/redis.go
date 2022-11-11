package common

import (
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

// 连接 redis 服务器
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
        Password: RedisPassword, // no password set
        DB:       RedisDB,  // use default DB
    })

	RedisClient = rdb

}

func GetRedis() *redis.Client {
	return RedisClient
}