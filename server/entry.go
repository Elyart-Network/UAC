package server

import (
	"github.com/gin-gonic/gin"
)

func Entry(server *gin.Engine) {
	server.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
}
