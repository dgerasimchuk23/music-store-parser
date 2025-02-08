package db

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func TestConnectDB_Success(t *testing.T) {
	// Устанавливаем переменные окружения
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "newpassword")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "parser_db")

	// Подключаемся к базе данных
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Проверяем соединение
	err = db.Ping()
	if err != nil {
		t.Fatalf("Ошибка проверки соединения с базой данных: %v", err)
	}
}
