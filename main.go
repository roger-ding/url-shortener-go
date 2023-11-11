package main

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/roger-ding/url-shortener-go/lib"
  "github.com/roger-ding/url-shortener-go/handler"
)

func main() {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "URL Shortner in Go!",
    })
  })

  r.POST("/create-shortened-url", func(c *gin.Context) {
    handler.CreateShortenedUrl(c)
  })
  
  r.GET("/:shortenedUrl", func(c *gin.Context) {
    handler.RedirectShortenedUrl(c)
  })

  // Redis initialization
  lib.InitializeRedis()

  err := r.Run()
  if err != nil {
    panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
  }
}
