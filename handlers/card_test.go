package handlers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	model "word-card-app/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Init() {
	err := model.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// 创建数据库表
	// err = model.CreateTables()
	if err != nil {
		log.Fatal("Failed to create tables:", err)
	}
}

func TestGetWordCards(t *testing.T) {
	Init()
	// w := &responseWriter{}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	GetWordCards(c)

	fmt.Println("return body", w.Body.String(), w.Code)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateWordCard(t *testing.T) {
	Init()
	data := []byte(`{"front":"test front", "back":"test back"}`)
	req, _ := http.NewRequest("POST", "/api/v1/word-cards", bytes.NewBuffer(data))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	CreateWordCard(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	fmt.Println("return body", w.Body.String(), w.Code)

	expectedJSON := `{
		"front": "test front",
		"back": "test back"
	}`
	assert.JSONEq(t, expectedJSON, w.Body.String())

}
