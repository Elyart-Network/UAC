package uac

import "github.com/gin-gonic/gin"

// AuthorizeService godoc
// @Summary Authorize Endpoint
// @Description Authorize Users under Datasource Mode
// @Tags UAC
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Router /uac/auth [post]
func AuthorizeService(ctx *gin.Context) {

}

// AuthProviderService godoc
// @Summary AuthProvider Endpoint
// @Description Authorize Users under Provider Mode
// @Tags UAC
// @Success 302 "<callback.uri>?code=[auth_code]"
// @Router /uac/auth [get]
func AuthProviderService(ctx *gin.Context) {

}
