package db

import (
	"database/sql"
	"fmt"
	"log"
)

// InitializeSchema создаёт таблицы, если их нет
func InitializeSchema(db *sql.DB) error {
	log.Println("Выполняется создание схемы...")

	// SQL-схема для создания таблиц
	schema := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		unit_of_measurement TEXT NOT NULL,  -- Единица измерения (шт, комплект, м и т.д.)
		price NUMERIC(10, 2) NOT NULL,
		url TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL UNIQUE
	);

	CREATE TABLE IF NOT EXISTS product_categories (
		product_id INT REFERENCES products(id) ON DELETE CASCADE,
		category_id INT REFERENCES categories(id) ON DELETE CASCADE,
		PRIMARY KEY (product_id, category_id)
	);

	CREATE TABLE IF NOT EXISTS logs (
		id SERIAL PRIMARY KEY,
		level TEXT NOT NULL,
		message TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);
	`

	// Логируем SQL для отладки
	log.Println("SQL для выполнения:\n", schema)

	// Выполняем SQL-схему
	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("Ошибка создания схемы базы данных: %v", err)
	}

	log.Println("Схема успешно создана.")
	return nil
}
