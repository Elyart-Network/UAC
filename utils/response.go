package utils

import (
	"github.com/Elyart-Network/UAC/model"
	"github.com/gin-gonic/gin"
)

func BaseResponse(ctx *gin.Context, code int, msg string, data any) {
	ctx.JSON(code, model.BaseResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func CommonResponse(ctx *gin.Context, code int) {
	switch code {
	case 200:
		BaseResponse(ctx, 200, "ok", nil)
	case 400:
		BaseResponse(ctx, 400, "Bad Request", nil)
	case 401:
		BaseResponse(ctx, 401, "Unauthorized", nil)
	case 403:
		BaseResponse(ctx, 403, "Forbidden", nil)
	case 404:
		BaseResponse(ctx, 404, "Not Found", nil)
	case 500:
		BaseResponse(ctx, 500, "Internal Server Error", nil)
	}
}
