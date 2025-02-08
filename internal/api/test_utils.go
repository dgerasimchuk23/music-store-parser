package api

import (
	"database/sql"
	"os"
	"testing"

	"parser/internal/db"
)

// Подключение тестовой db (используется во всех тестах)
func setupTestDB(t *testing.T) *sql.DB {
	t.Helper() // Вспомогательная функция для тестов

	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "newpassword")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "test_parser_db")

	testDB, err := db.ConnectDB()
	if err != nil {
		t.Fatalf("Ошибка подключения к тестовой базе: %v", err)
	}

	if err := db.InitializeSchema(testDB); err != nil {
		t.Fatalf("Ошибка инициализации схемы: %v", err)
	}

	return testDB
}
