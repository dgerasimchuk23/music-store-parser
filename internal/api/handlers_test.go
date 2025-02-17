package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Тест API `/products` (Добавление товара)
func TestAddProduct(t *testing.T) {
	dbConn := setupTestDB(t)
	defer dbConn.Close()
	router := SetupRouter(dbConn)

	// Очищаем базу перед тестом
	_, err := dbConn.Exec("DELETE FROM products;")
	assert.Nil(t, err, "Ошибка очистки БД перед тестом")

	// Данные в формате `db.Product`
	product := map[string]interface{}{
		"name":                "Тестовый продукт",
		"unit_of_measurement": "шт",
		"price":               999.99, // Проверяем, принимает ли API float
		"url":                 "https://example.com/test",
	}

	jsonData, _ := json.Marshal(product)
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Логируем тело ответа, если API вернул ошибку
	if rec.Code != http.StatusCreated {
		t.Logf("Ответ сервера: %s", rec.Body.String())
	}

	assert.Equal(t, http.StatusCreated, rec.Code, "Ошибка: API вернул 400, ожидался 201")

	// Проверяем, что продукт реально добавился в БД
	var count int
	err = dbConn.QueryRow("SELECT COUNT(*) FROM products WHERE name = $1", "Тестовый продукт").Scan(&count)
	assert.Nil(t, err)
	assert.Equal(t, 1, count, "Ошибка: Тестовый продукт должен быть в БД")
}
