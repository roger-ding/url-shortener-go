package handler

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/roger-ding/url-shortener-go/lib"
)

type UrlCreationRequest struct {
  SourceUrl string `json:"source_url" binding:"required"`
}

func CreateShortenedUrl(c *gin.Context) {
  var creationRequest UrlCreationRequest
  if err := c.ShouldBindJSON(&creationRequest); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  
  shortenedUrl := lib.GenerateShortenedUrl(creationRequest.SourceUrl)
  lib.StoreShortenedUrl(shortenedUrl, creationRequest.SourceUrl)
  
  host := "http://localhost:8080/"
  c.JSON(http.StatusOK, gin.H{
    "message":   "Shortened url created successfully!",
    "shortened_url": host + shortenedUrl,
  })
}

func RedirectShortenedUrl(c *gin.Context) {
  shortenedUrl := c.Param("shortenedUrl")
  sourceUrl := lib.GetShortenedUrl(shortenedUrl)
  c.Redirect(302, sourceUrl)
}
