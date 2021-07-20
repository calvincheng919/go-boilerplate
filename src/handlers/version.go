package handlers

import (
	"github.com/gin-gonic/gin"
)

/**
 * @api {get} /api/version
 * @apiName Version
 * @apiDescription displays version
 * @apiGroup misc
 * @apiVersion 0.0.0
 * @apiExample {curl} Curl Example:
 * curl http://localhost:8080/api/version
 * @apiSuccess {String} version current version
 */
func Version(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{
		"version": "0.0.0",
	})
}
