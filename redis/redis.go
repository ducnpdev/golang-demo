package redis

import (
	"github.com/go-redis/redis/v9"
)

var RedisIn *redis.Client

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	RedisIn = rdb
}
