package main

import (
	"fmt"
	"marketplace_teste_tecnico/src/controller"
	"marketplace_teste_tecnico/src/models"
	"marketplace_teste_tecnico/src/routes"
	"marketplace_teste_tecnico/src/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	config := utils.GetConfig()
	db := utils.DB(config.Db)
	if err := db.AutoMigrate(&models.Product{}); err != nil {
		panic(err)
	}

	router := gin.Default()
	routes.ProductRoutes(router, *controller.NewProductController(db))

	address := fmt.Sprintf("%s:%s", config.Api.Host, config.Api.Port)
	if err := router.Run(address); err != nil {
		panic(err)
	}
}
