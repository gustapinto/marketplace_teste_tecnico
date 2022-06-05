package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

func MockGormDB(model any) *gorm.DB {
	gormDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	gormDB.Logger.LogMode(logger.Silent)
	gormDB.AutoMigrate(model)

	return gormDB
}

func RandomStr() string {
	rand.Seed(time.Now().UnixMicro())
	str := fmt.Sprintf("%d", rand.Int63())

	return str
}
