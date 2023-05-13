package health

import (
	"github.com/Elyart-Network/UAC/utils"
	"github.com/gin-gonic/gin"
)

// Service godoc
// @Summary Health Check
// @Description Check if the server is up and running
// @Tags Core
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Router /health [get]
func Service(ctx *gin.Context) {
	utils.BaseResponse(ctx, 200, "OK", nil)
}
