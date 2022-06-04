package repository

import (
	"errors"
	"marketplace_teste_tecnico/src/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

var ErrPriceOfLowerThanPriceFor = errors.New("price Of cant be lower than Price For")

func (r *ProductRepository) Save(product *models.Product) error {
	if product.PriceOf < product.PriceFor {
		return ErrPriceOfLowerThanPriceFor
	}

	product.Inventory.Available = product.Inventory.Total - product.Inventory.Cut
	r.DB.Save(product)
	return nil
}
