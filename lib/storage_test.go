package lib

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

var testRedisService = &RedisService{}

func init() {
  testRedisService = InitializeRedis()
}

func TestRedisInit(t *testing.T) {
  assert.True(t, testRedisService.redisClient != nil)
}

func TestStoreAndGet(t *testing.T) {
  initialLink := "https://www.google.com"
  shortURL := "goog"
  
  StoreShortenedUrl(shortURL, initialLink)
  assert.Equal(t, initialLink, GetShortenedUrl(shortURL))
}
