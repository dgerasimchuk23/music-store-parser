package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Проверка инициализации маршрутов
func TestRouter(t *testing.T) {
	dbConn := setupTestDB(t)
	defer dbConn.Close()
	router := SetupRouter(dbConn)

	req, _ := http.NewRequest("GET", "/products", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
