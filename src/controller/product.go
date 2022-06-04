package controller

import (
	"marketplace_teste_tecnico/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Product struct {
	Db *gorm.DB
}

func (c *Product) Save(ctx *gin.Context) {
	var newProduct models.Product
	if err := ctx.BindJSON(&newProduct); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Db.Save(&newProduct)
	ctx.IndentedJSON(http.StatusCreated, newProduct)
	return
}
