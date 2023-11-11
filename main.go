package main

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "URL Shortner in Go!",
    })
  })

  err := r.Run()
  if err != nil {
    panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
  }
}
