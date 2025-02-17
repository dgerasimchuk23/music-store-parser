package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"parser/internal/db"

	"github.com/stretchr/testify/assert"
)

// Проверяет запуск сервера API
func TestServer(t *testing.T) {
	dbConn := setupTestDB(t)
	defer dbConn.Close()

	// Добавляем тестовый продукт
	_ = db.SaveProduct(dbConn, "Тестовый товар", "шт", "999.99", "https://example.com/product")

	// Запускаем API через тестовый сервер
	server := httptest.NewServer(SetupRouter(dbConn))
	defer server.Close()

	// Запрашиваем список товаров
	req, _ := http.NewRequest("GET", server.URL+"/products", nil)
	rec := httptest.NewRecorder()
	SetupRouter(dbConn).ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
