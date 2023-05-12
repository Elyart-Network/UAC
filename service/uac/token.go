package uac

import "github.com/gin-gonic/gin"

// TokenService godoc
// @Summary Token Endpoint
// @Description Generate and Maintain Tokens
// @Tags UAC
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Router /uac/token [post]
func TokenService(ctx *gin.Context) {

}

// RefreshService godoc
// @Summary Refresh Endpoint
// @Description Refresh Tokens
// @Tags UAC
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Router /uac/token [put]
func RefreshService(ctx *gin.Context) {

}
