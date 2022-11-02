package store

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	log.Println("initialized store...")

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient

	return storeService
}

func SaveUrlMapping(shortUrl string, originalUrl string, urlId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration)
	if err.Err() != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

func RetriveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		log.Println(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
		return ""
	}

	return result
}
