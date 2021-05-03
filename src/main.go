package main

import (
	"github.com/gin-gonic/gin"
	"my-service/src/handlers"
	"my-service/src/utils"
)

func main() {
	app := gin.Default() // create gin instance with defaults (logger & recovery middleware)

	app.Use(utils.CORSMiddleware())

	/**
	 * @api {get} /myservice/version
	 * @apiName Version
	 * @apiDescription displays version
	 * @apiGroup misc
	 * @apiVersion 0.0.0
	 * @apiExample {curl} Curl Example:
	 * curl http://localhost:8080/myservice/version
	 * @apiSuccess {String} version current version
	 */
	app.GET("/myservice/version", handlers.Version)

	func() { _ = app.Run() }()
}
