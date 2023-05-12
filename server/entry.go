package server

import (
	"github.com/Elyart-Network/UAC/docs"
	"github.com/Elyart-Network/UAC/service/health"
	"github.com/Elyart-Network/UAC/service/uac"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Entry(e *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	e.GET("/health", health.Service)
	uacRoute := e.Group("/uac")
	{
		uacRoute.POST("/token", uac.TokenService)
		uacRoute.PUT("/token", uac.RefreshService)
		uacRoute.POST("/cert", uac.CertService)
		uacRoute.GET("/auth", uac.AuthProviderService)
		uacRoute.POST("/auth", uac.AuthorizeService)
	}
}
