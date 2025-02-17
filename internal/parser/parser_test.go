package parser

import (
	"database/sql"
	"os"
	"testing"

	"parser/internal/db"

	"github.com/stretchr/testify/assert"
)

// Настройка тестовой базы данных
func setupTestDB(t *testing.T) *sql.DB {
	t.Helper()

	// Устанавливаем переменные окружения для тестовой БД
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "newpassword")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "test_parser_db") // Теперь DB_NAME всегда задана

	// Подключаемся к БД
	database, err := db.ConnectDB()
	assert.Nil(t, err, "Ошибка подключения к тестовой базе")

	return database
}

// Тест сохранения продукта в БД
func TestSaveProduct(t *testing.T) {
	dbConn := setupTestDB(t)
	defer dbConn.Close()

	// Очищаем базу перед тестом
	_, err := dbConn.Exec("DELETE FROM products;")
	assert.Nil(t, err, "Ошибка очистки базы перед тестом")

	// Добавляем тестовый продукт
	err = db.SaveProduct(dbConn, "Тестовый продукт", "шт", "999.99", "https://example.com/test-product")
	assert.Nil(t, err, "Ошибка сохранения в БД")

	// Проверяем, что запись добавлена
	var count int
	err = dbConn.QueryRow("SELECT COUNT(*) FROM products WHERE name = 'Тестовый продукт'").Scan(&count)
	assert.Nil(t, err)
	assert.Equal(t, 1, count, "Ошибка: товар не добавлен в БД")
}
