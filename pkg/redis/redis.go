package redisClient

import (
	"context"
	"log"
	"main/config"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func InitRedis() *RedisClient {
	params := config.LoadConfig()

	client := redis.NewClient(&redis.Options{
		Addr:     params.RedisAddr,
		Password: params.RedisPass,
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("error connecting to Redis: %v", err)
	}

	log.Println("successfully connected to Redis")

	return &RedisClient{Client: client}
}

func (r *RedisClient) GetKey(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *RedisClient) SetKey(ctx context.Context, key string, value string, expiration time.Duration) error {
	return r.Client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisClient) DeleteKey(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}
