package config

import (
	"context"
	"ecommerce/helper"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

func InitRedis(redisConfig RedisConfig) {
	ctx := context.Background() // Konteks untuk operasi-operasi Redis

	if redisConfig.Password == "null" {
		redisConfig.Password = ""
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		Password: redisConfig.Password,
		DB:       0,
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		helper.LogError(fmt.Errorf("Error connecting to Redis: \n%s", err))
	}
}
