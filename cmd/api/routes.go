package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func routes(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hi",
		})
	})
}