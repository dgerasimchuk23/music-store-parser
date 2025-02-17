package db

import (
	"database/sql"
	"log"
	"time"
)

// Product — структура продукта
type Product struct {
	ID                int
	Name              string
	UnitOfMeasurement string
	Price             float64
	URL               string
	CreatedAt         string
}

// SaveProduct сохраняет товар в базу данных
func SaveProduct(database *sql.DB, name, unit, price, url string) error {
	if database == nil {
		database = DB // Если передана nil-ссылка, используем глобальную БД
	}

	if database == nil {
		log.Println("Ошибка: база данных не инициализирована!")
		return sql.ErrConnDone
	}

	_, err := database.Exec(`
		INSERT INTO products (name, unit_of_measurement, price, url, created_at)
		VALUES ($1, $2, $3, $4, $5)`,
		name, unit, price, url, time.Now(),
	)
	if err != nil {
		log.Println("Ошибка при сохранении товара:", err)
		return err
	}
	log.Println("Товар успешно сохранен в БД:", name)
	return nil
}
