package server

import (
	"github.com/Elyart-Network/UAC/internal/config"
	"github.com/gin-gonic/gin"
)

func Entry(e *gin.Engine) {
	api := e.Group("/")
	if !config.Docker() {
		api = e.Group("/api")
	}
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
}
