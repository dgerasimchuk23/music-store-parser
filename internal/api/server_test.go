package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Проверяет запуск сервера API
func TestServer(t *testing.T) {
	dbConn := setupTestDB(t)
	defer dbConn.Close()

	// Запуск сервера в тестовом режиме
	go func() {
		StartServer(dbConn)
	}()

	// Тестовый запрос к серверу
	req, _ := http.NewRequest("GET", "http://localhost:8080/products", nil)
	rec := httptest.NewRecorder()
	router := SetupRouter(dbConn)
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
