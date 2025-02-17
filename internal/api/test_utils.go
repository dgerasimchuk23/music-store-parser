package api

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
