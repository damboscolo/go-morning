package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"go-morning/api"
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

	router.POST("/morning", api.SendMorningMessage)
	router.POST("/evening", api.SendEveningMessage)
	router.POST("/dani", api.SaveDani)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":" + os.Getenv("PORT"))
}
