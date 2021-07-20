package handlers

import "github.com/gin-gonic/gin"

func Home(ctx *gin.Context) {
	ctx.JSON(200, "Hello, world!")
}
