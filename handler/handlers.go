package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/polunzh/url-shortener/shortener"
	"github.com/polunzh/url-shortener/store"
)

const host string = "http://127.0.0.1:8080/"

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	c.JSON(http.StatusOK, gin.H{"message": "created", "short_url": host + shortUrl})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("url")
	originalUrl := store.RetriveInitialUrl(shortUrl)
	if originalUrl == "" {
		c.Data(http.StatusNotFound, "text/html;charset=utf-8", []byte("Not Found"))
		return
	}

	c.Redirect(302, originalUrl)
}
