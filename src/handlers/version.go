package handlers

import (
	"github.com/gin-gonic/gin"
	T "sms-service/src/types"
)

func Version(ctx *gin.Context) {
	ctx.JSON(200, T.Map{
		"version": "0.0.0",
	})
}