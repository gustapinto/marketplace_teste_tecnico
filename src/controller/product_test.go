package controller

import (
	"marketplace_teste_tecnico/src/utils"
	"net/http"
	"testing"
)

func TestSaveMustReturnSuccessWithValidJson(t *testing.T) {
	mockedJsonContent := map[string]any{
		"codigo": "foobar",
		"nome":   "foo",
		"estoque": map[string]any{
			"total": 100,
			"corte": 75,
		},
		"preco_de":  100,
		"preco_por": 75,
	}

	controller := NewProductController(utils.MockGormDB())
	ctx, recorder := utils.MockGinContext(http.MethodPost)
	_ = utils.MockJsonPost(ctx, mockedJsonContent)

	controller.Save(ctx)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expecting to create a new product with a valid json, got code %d", recorder.Code)
	}
}

func TestSaveMustReturnBadRequestOnInvalidBody(t *testing.T) {
	mockedJsonContent := map[string]any{
		"codigo": "foobar",
		"nome":   "foo",
		"estoque": map[string]any{
			"total": 100,
			"corte": 75,
		},
		"preco_de":  100,
		"preco_por": 110,
	}

	controller := NewProductController(utils.MockGormDB())
	ctx, recorder := utils.MockGinContext(http.MethodPost)
	_ = utils.MockJsonPost(ctx, mockedJsonContent)

	controller.Save(ctx)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected to fail with a invalid json body, got code %d", recorder.Code)
	}
}
