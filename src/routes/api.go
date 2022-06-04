package routes

import (
	"marketplace_teste_tecnico/src/controller"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine, controller *controller.Product) {
	router.POST("/api/products", controller.Save)
}
