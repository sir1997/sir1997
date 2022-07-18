package dao

import (
	"github.com/go-redis/redis"
)

var Redis *redis.Client

func SetRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       8,
	})

}
