package main

import (
	"fmt"
	"marketplace_teste_tecnico/src/routes"
	"marketplace_teste_tecnico/src/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	config := utils.GetConfig()
	router := gin.Default()
	address := fmt.Sprintf("%s:%s", config.Api.Host, config.Api.Port)

	db, err := utils.DB(config.Db)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate()
	routes.Api(router, db)

	if err := router.Run(address); err != nil {
		panic(err)
	}
}
