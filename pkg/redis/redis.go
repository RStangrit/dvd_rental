package redisClient

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	log.Println("successfully connected to Redis")
	return redisClient
}
