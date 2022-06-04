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

func (c *Product) GetAll(ctx *gin.Context) {
	var products []models.Product
	if result := c.repo.DB.Find(&products); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, products)
	return
}

func (c *Product) Save(ctx *gin.Context) {
	var newProduct models.Product
	if err := ctx.BindJSON(&newProduct); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.repo.Save(&newProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, newProduct)
	return
}

func (c *Product) GetByCode(ctx *gin.Context) {
	productCode := ctx.Param("code")

	var product models.Product
	if result := c.repo.DB.First(&product, "code = ?", productCode); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, product)
	return
}
