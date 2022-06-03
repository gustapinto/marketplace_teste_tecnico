package routes

import (
	"marketplace_teste_tecnico/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Api(router *gin.Engine, db *gorm.DB) {
	router.GET("/api/products", func(ctx *gin.Context) {
		var newProduct models.Product
		if err := ctx.BindJSON(&newProduct); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
			return
		}

		db.Save(&newProduct)
		ctx.IndentedJSON(http.StatusCreated, newProduct)
	})
}
