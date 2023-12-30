package database

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func StartRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("BREVIORI_REDIS_HOST"),
		Username: os.Getenv("BREVIORI_REDIS_USERNAME"),
		Password: os.Getenv("BREVIORI_REDIS_PASSWORD"),
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		slog.Error(fmt.Sprintf("failed connecting to redis: %v", err))
		os.Exit(1)
	}

	slog.Info("successfully connected to redis")
}

func GetRedisInstance() *redis.Client {
	return rdb
}
