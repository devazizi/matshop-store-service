package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouteV1(router *gin.Engine) {

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotAcceptable, map[string]any{
			"message": "not acceptable request",
		})
	})
}
