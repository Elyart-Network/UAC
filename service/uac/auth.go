package uac

import "github.com/gin-gonic/gin"

// AuthorizeService godoc
// @Summary Authorize Endpoint
// @Description Authorize Users
// @Tags UAC
// @Success 302 "<callback.uri>?code=[auth_code]"
// @Router /uac/auth [get]
func AuthorizeService(ctx *gin.Context) {

}
