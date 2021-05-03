package main

import (
	"github.com/gin-gonic/gin"
	"sms-service/src/handlers"
	"sms-service/src/utils"
)

func main() {
	app := gin.Default() // create gin instance with defaults (logger & recovery middleware)

	app.Use(utils.CORSMiddleware())

	/**
	 * @api {get} /version
	 * @apiName Version
	 * @apiDescription displays version
	 * @apiGroup misc
	 * @apiVersion 0.0.0
	 * @apiExample {curl} Curl Example:
	 * curl http://localhost:8080/version
	 * @apiSuccess {String} version current version
	 */
	app.GET("/version", handlers.Version)

	func() { _ = app.Run() }()
}