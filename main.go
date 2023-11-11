package main

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/roger-ding/url-shortener-go/lib"
)

func main() {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "URL Shortner in Go!",
    })
  })

  lib.InitializeRedis()

  testKey := "test"
  lib.StoreShortenedUrl(testKey, "hello world!")
  fmt.Printf("Retrieved value for key abc - %s\n", lib.GetShortenedUrl(testKey))

  err := r.Run()
  if err != nil {
    panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
  }
}
