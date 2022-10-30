package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// Force log's color
	gin.ForceConsoleColor()
	router.Use(gin.Logger())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// dani := router.Group("/good")
	// {
	// 	dani.POST("/morning", sendGoodMorningMessage)
	// 	// 	dani.POST("/night", nightEndpoint)
	// }

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
}
