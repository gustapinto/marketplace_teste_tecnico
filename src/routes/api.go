package routes

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Api(router *gin.Engine) {
	router.GET("/api/products", func(ctx *gin.Context) {
		log.Printf("Pong")
	})
}
