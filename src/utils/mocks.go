package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MockGinContext(method string) (*gin.Context, *httptest.ResponseRecorder) {
	mockedRecorder := httptest.NewRecorder()
	mockedContext, _ := gin.CreateTestContext(mockedRecorder)
	mockedContext.Request = &http.Request{Header: make(http.Header)}
	mockedContext.Request.Method = method

	return mockedContext, mockedRecorder
}

func MockJsonPost(ctx *gin.Context, content any) []byte {
	ctx.Request.Header.Set("Content-Type", "application/json")

	jsonContent, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonContent))

	return jsonContent
}

func MockGormDB() *gorm.DB {
	sqlDB, _, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return gormDB
}
