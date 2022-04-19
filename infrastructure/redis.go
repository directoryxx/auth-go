package infrastructure

import (
	"github.com/go-redis/redis/v8"
	"os"
)

func OpenRedis() *redis.Client  {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
