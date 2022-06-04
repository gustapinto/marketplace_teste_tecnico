package repository

import (
	"errors"
	"marketplace_teste_tecnico/src/models"
	"marketplace_teste_tecnico/src/utils"
	"testing"
)

func TestSaveMustErrorWithInvalidPrices(t *testing.T) {
	repo := ProductRepository{DB: utils.MockGormDB()}

	mockedInvalidProduct := &models.Product{}
	mockedInvalidProduct.PriceOf = 75
	mockedInvalidProduct.PriceFor = 100

	if err := repo.Save(mockedInvalidProduct); !errors.Is(err, ErrPriceOfLowerThanPriceFor) {
		t.Error("Expected to fail on invalid product prices, got a success")
	}
}
