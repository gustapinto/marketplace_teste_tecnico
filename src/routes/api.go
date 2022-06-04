package routes

import (
	"marketplace_teste_tecnico/src/controller"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine, c controller.Product) {
	router.GET("/api/products", c.GetAll)
	router.POST("/api/products", c.Save)
}
