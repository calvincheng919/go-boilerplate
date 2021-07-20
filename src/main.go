package main

import (
	"github.com/gin-gonic/gin"
	"my-service/src/handlers"
	"my-service/src/utils"
)

func main() {
	app := gin.Default() // create gin instance with defaults (logger & recovery middleware)

	app.Use(utils.CORSMiddleware())

	app.GET("/version", handlers.Version)

	app.GET("/", handlers.Home)

	func() { _ = app.Run() }()
}
