package core

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var RedisContext = context.Background()

func RedisClient() *redis.Client {
	host := SharedConfig.Redis.Host
	port := SharedConfig.Redis.Port
	addr := fmt.Sprintf("%s:%d", host, port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	SharedLogger.Infof("addr: %s", addr)
	return client
}
