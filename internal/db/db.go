package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL драйвер
)

// Глобальная переменная для хранения подключения к БД
var DB *sql.DB

// ConnectDB - подключение к базе данных
func ConnectDB() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Проверка, что имя базы данных не пустое
	if dbName == "" {
		log.Fatal("Ошибка: переменная окружения DB_NAME не задана!")
	}

	// Подключаемся к базе данных
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе `%s`: %v", dbName, err)
		return nil, err
	}

	// Проверяем соединение
	if err := db.Ping(); err != nil {
		log.Fatalf("Ошибка при проверке соединения с базой `%s`: %v", dbName, err)
		return nil, err
	}

	DB = db
	log.Printf("Подключение к базе `%s` установлено", dbName)

	// Создаём таблицы, если их нет
	err = InitializeSchema(db)
	if err != nil {
		log.Fatalf("Ошибка при создании таблиц в `%s`: %v", dbName, err)
		return nil, err
	}

	return db, nil
}
