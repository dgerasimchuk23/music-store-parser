package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL драйвер
)

// ConnectDB - подключение к базе данных
func ConnectDB() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Подключаемся к postgres (без указания конкретной базы)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения к серверу PostgreSQL: %v", err)
	}

	// Проверяем, существует ли база данных
	exists, err := checkDatabaseExists(db, dbName)
	if err != nil {
		return nil, err
	}

	// Создаём базу данных, если её нет
	if !exists {
		err = createDatabase(db, dbName)
		if err != nil {
			return nil, err
		}
		log.Printf("База данных %s успешно создана", dbName)
	}

	// Подключаемся к целевой базе данных
	dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db.Close() // Закрываем текущее подключение к postgres
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе %s: %v", dbName, err)
	}

	log.Printf("Подключение к базе данных %s установлено", dbName)
	return db, nil
}

// Проверяет, существует ли база данных
func checkDatabaseExists(db *sql.DB, dbName string) (bool, error) {
	query := `SELECT 1 FROM pg_database WHERE datname = $1`
	var exists int
	err := db.QueryRow(query, dbName).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, fmt.Errorf("Ошибка проверки базы данных: %v", err)
	}
	return exists == 1, nil
}

// Создаёт базу данных
func createDatabase(db *sql.DB, dbName string) error {
	query := fmt.Sprintf("CREATE DATABASE %s", dbName)
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("Ошибка создания базы данных %s: %v", dbName, err)
	}
	return nil
}
