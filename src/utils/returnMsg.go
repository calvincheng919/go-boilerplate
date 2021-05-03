package utils

import (
	"github.com/gin-gonic/gin"
)

func ReturnMsg(ctx *gin.Context, msg string, responseCode int) {
	ctx.JSON(responseCode, map[string]string{
		"message": msg,
	})
}
