package redis

import (
	"context"
	"fmt"

	"belajargo/internal/config"
	"belajargo/internal/services"

	redisv9 "github.com/redis/go-redis/v9"
)

func ConncetRedis(cfg *config.Config) (*redisv9.Client, error) {
	protocol := 2
	if cfg.RedisProtocol != 0 {
		protocol = cfg.RedisProtocol
	}

	rdb := redisv9.NewClient(&redisv9.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
		Protocol: protocol,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	services.Log(services.Logger{
		Name:    "REDIS",
		Message: "Connected to redis",
	})
	return rdb, nil
}
