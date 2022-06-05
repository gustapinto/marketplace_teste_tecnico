package repository

import (
	"errors"
	"marketplace_teste_tecnico/src/models"
	"marketplace_teste_tecnico/src/utils"
	"testing"
)

func TestSaveMustErrorWithInvalidPrices(t *testing.T) {
	repo := ProductRepository{DB: utils.MockGormDB(&models.Product{})}

	mockedCode := utils.RandomStr()
	mockedName := utils.RandomStr()
	mockedProduct := &models.Product{
		Code:     &mockedCode,
		Name:     &mockedName,
		PriceOf:  75,
		PriceFor: 100,
	}

	if err := repo.Save(mockedProduct); !errors.Is(err, ErrPriceOfLowerThanPriceFor) {
		t.Errorf("Expected to fail on invalid product prices, got error %+v", err)
	}
}

func TestGet(t *testing.T) {
	repo := ProductRepository{DB: utils.MockGormDB(&models.Product{})}

	mockedCode := utils.RandomStr()
	mockedName := utils.RandomStr()
	mockedProduct := &models.Product{
		Code: &mockedCode,
		Name: &mockedName,
	}

	if err := repo.Save(mockedProduct); err != nil {
		t.Errorf("Did not expect to fail on product save, got error %+v", err)
		return
	}

	_, err := repo.Get(mockedCode)
	if err != nil {
		t.Errorf("Did not expect to fail on get, got error %+v", err)
		return
	}
}
