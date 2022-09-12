package cache

import (
	"app/configs"
	"github.com/go-redis/redis/v9"
)

func Connection() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     configs.EnvString("REDIS_HOST") + ":" + configs.EnvString("REDIS_PORT"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
