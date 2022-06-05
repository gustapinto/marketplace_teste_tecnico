package controller

import (
	"marketplace_teste_tecnico/src/models"
	"marketplace_teste_tecnico/src/utils"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSaveMustReturnSuccessWithValidJson(t *testing.T) {
	mockedJsonContent := map[string]any{
		"codigo": utils.RandomStr(),
		"nome":   "foo",
		"estoque": map[string]any{
			"total": 100,
			"corte": 75,
		},
		"preco_de":  100,
		"preco_por": 75,
	}

	controller := NewProductController(utils.MockGormDB(&models.Product{}))
	ctx, recorder := utils.MockGinContext(http.MethodPost)
	_ = utils.MockJsonPost(ctx, mockedJsonContent)

	controller.Save(ctx)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expecting to create a new product with a valid json, got code %d", recorder.Code)
	}
}

func TestSaveMustReturnBadRequestOnInvalidBody(t *testing.T) {
	mockedJsonContent := map[string]any{
		"codigo": utils.RandomStr(),
		"nome":   "foo",
		"estoque": map[string]any{
			"total": 100,
			"corte": 75,
		},
		"preco_de":  100,
		"preco_por": 110,
	}

	controller := NewProductController(utils.MockGormDB(&models.Product{}))
	ctx, recorder := utils.MockGinContext(http.MethodPost)
	_ = utils.MockJsonPost(ctx, mockedJsonContent)

	controller.Save(ctx)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected to fail with a invalid json body, got code %d", recorder.Code)
		return
	}
}

func TestUpdateMustReturnBadRequestOnInvalidBody(t *testing.T) {
	mockedCode := utils.RandomStr()
	controller := NewProductController(utils.MockGormDB(&models.Product{}))
	ctx, recorder := utils.MockGinContext(http.MethodPost)

	mockedJsonContent := map[string]any{
		"codigo": mockedCode,
		"nome":   "foo",
		"estoque": map[string]any{
			"total": 100,
			"corte": 75,
		},
		"preco_de":  100,
		"preco_por": 90,
	}
	_ = utils.MockJsonPost(ctx, mockedJsonContent)
	controller.Save(ctx)
	if recorder.Code != http.StatusCreated {
		t.Errorf("Did not espect to fail on first insertion, got code %d", recorder.Code)
		return
	}

	ctx, recorder = utils.MockGinContext(http.MethodPost)
	codeParam := gin.Param{Key: "code", Value: mockedCode}
	ctx.Params = append(ctx.Params, codeParam)
	mockedInvalidUpdatedJsonContent := map[string]any{
		"codigo": mockedCode,
		"nome":   "foo",
		"estoque": map[string]any{
			"total": 100,
			"corte": 75,
		},
		"preco_de":  100,
		"preco_por": 120,
	}
	_ = utils.MockJsonPost(ctx, mockedInvalidUpdatedJsonContent)
	controller.Update(ctx)
	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected to fail with a invalid updated json body, got code %d", recorder.Code)
		return
	}
}
