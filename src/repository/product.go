package repository

import (
	"errors"
	"marketplace_teste_tecnico/src/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

var ErrPriceOfLowerThanPriceFor = errors.New("price of preco_de cant be lower than preco_por")

func (r *ProductRepository) Save(product *models.Product) error {
	if product.PriceOf < product.PriceFor {
		return ErrPriceOfLowerThanPriceFor
	}

	product.Inventory.Available = product.Inventory.Total - product.Inventory.Cut
	if result := r.DB.Save(product); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ProductRepository) Get(productCode string) (*models.Product, error) {
	var product models.Product
	if result := r.DB.First(&product, "code = ?", productCode); result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}
