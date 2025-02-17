package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"parser/internal/db"

	"github.com/stretchr/testify/assert"
)

// Тест API `/products`
func TestRouter(t *testing.T) {
	dbConn := setupTestDB(t)
	defer dbConn.Close()
	router := SetupRouter(dbConn)

	// Очищаем базу перед тестом
	_, err := dbConn.Exec("DELETE FROM products;")
	assert.Nil(t, err, "Ошибка очистки базы перед тестом")

	// Добавляем тестовый продукт в БД
	err = db.SaveProduct(dbConn, "Тестовый товар", "шт", "999.99", "https://example.com/product")
	assert.Nil(t, err, "Ошибка добавления тестового товара в БД")

	// Ожидание, чтобы БД обработала вставку
	time.Sleep(200 * time.Millisecond)

	// Запрос к API `/products`
	req, _ := http.NewRequest("GET", "/products", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Проверяем, что API вернул 200 OK
	assert.Equal(t, http.StatusOK, rec.Code, "Ошибка: API `/products` вернул %d, ожидался 200", rec.Code)

	// Декодируем JSON-ответ API
	var response []map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Nil(t, err, "Ошибка парсинга JSON-ответа от API")

	// Логируем ответ API, если тест не проходит
	t.Logf("Ответ API: %s", rec.Body.String())

	// Проверяем, что тестовый товар есть в ответе
	found := false
	for _, p := range response {
		if name, ok := p["Name"].(string); ok && name == "Тестовый товар" {
			found = true
			break
		}
	}
	assert.True(t, found, "Ошибка: Тестовый товар должен быть в ответе API")
}
