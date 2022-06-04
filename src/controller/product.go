package controller

import (
	"marketplace_teste_tecnico/src/models"
	"marketplace_teste_tecnico/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Product struct {
	repo repository.ProductRepository
}

func NewProductController(db *gorm.DB) *Product {
	return &Product{
		repo: repository.ProductRepository{DB: db},
	}
}

func (c *Product) Save(ctx *gin.Context) {
	var newProduct models.Product
	if err := ctx.BindJSON(&newProduct); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	err := c.repo.Save(&newProduct)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
	}

	ctx.IndentedJSON(http.StatusCreated, newProduct)
	return
}
