package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"my-service/src/handlers"
	"my-service/src/utils"
)

func main() {
	app := gin.Default() // create gin instance with defaults (logger & recovery middleware)
	app.Use(utils.CORSMiddleware())

	api := app.Group("/api")

	api.GET("/version", handlers.Version)

	// displays api docs (docs/index.html) on index page
	app.Use(static.Serve("/", static.LocalFile("docs", true)))

	func() { _ = app.Run() }()
}
