package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/polunzh/url-shortener/handler"
	"github.com/polunzh/url-shortener/store"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Go URL Shorter",
		})
	})

	r.POST("/", handler.CreateShortUrl)

	r.GET("/:url", handler.HandleShortUrlRedirect)

	store.InitializeStore()

	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
