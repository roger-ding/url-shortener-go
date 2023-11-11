package lib

import (
  "context"
  "fmt"
  "time"
  "github.com/redis/go-redis/v9"
)

type RedisService struct {
  redisClient *redis.Client
}

var (
  redisService = &RedisService{}
  ctx = context.Background()
)

// Allow for automatic cleanup after specified time
const CacheDuration = 1 * time.Hour

// Initialize the redis storage service and return reference 
func InitializeRedis() *RedisService {
  redisClient := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "",
    DB:       0,
  })

  pong, err := redisClient.Ping(ctx).Result()
  if err != nil {
    panic(fmt.Sprintf("Redis client failed to initialize due to: %v", err))
  }
  
  fmt.Printf("\nRedis client started successfully")
  fmt.Printf("\nPong message = {%s}", pong)
  redisService.redisClient = redisClient
  return redisService
}

func StoreShortenedUrl(shortenedUrl string, sourceUrl string) {
  err := redisService.redisClient.Set(ctx, shortenedUrl, sourceUrl, CacheDuration).Err()
  if err != nil {
    panic(fmt.Sprintf("Failed to save shortened url %s due to: %v\n", shortenedUrl, err))
  }
  fmt.Printf("Successfully saved shortened URL!\nshortened URL: %s\nsource URL: %s\n", shortenedUrl, sourceUrl)
}

func GetShortenedUrl(shortenedUrl string) string {
  result, err := redisService.redisClient.Get(ctx, shortenedUrl).Result()
  if err != nil {
    panic(fmt.Sprintf("Failed to get shortened URL %s due to: %v\n", shortenedUrl, err))
  }
  return result
}
