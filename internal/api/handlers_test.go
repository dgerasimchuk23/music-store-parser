package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"parser/internal/db"

	"github.com/stretchr/testify/assert"
)

// Проверяет добавление продукта
func TestAddProduct(t *testing.T) {
	dbConn := setupTestDB(t)
	defer dbConn.Close()
	router := SetupRouter(dbConn)

	product := db.Product{
		Name:              "Тестовый продукт",
		UnitOfMeasurement: "шт",
		Price:             999.99,
		URL:               "https://example.com/test",
	}

	jsonData, _ := json.Marshal(product)
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
}
