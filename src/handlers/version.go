package handlers

import (
	"github.com/gin-gonic/gin"
)

func Version(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{
		"version": "0.0.0",
	})
}
