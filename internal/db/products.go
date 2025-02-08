package db

import (
	"database/sql"
	"log"
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

// AddProduct — добавляет новый продукт в БД
func AddProduct(db *sql.DB, product Product) error {
	query := `INSERT INTO products (name, unit_of_measurement, price, url) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, product.Name, product.UnitOfMeasurement, product.Price, product.URL)
	if err != nil {
		log.Printf("Ошибка добавления продукта: %v", err)
		return err
	}
	return nil
}

// GetProducts — получает список всех продуктов
func GetProducts(db *sql.DB) ([]Product, error) {
	query := `SELECT id, name, unit_of_measurement, price, url, created_at FROM products`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Ошибка получения продуктов: %v", err)
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.UnitOfMeasurement, &p.Price, &p.URL, &p.CreatedAt)
		if err != nil {
			log.Printf("Ошибка чтения строки: %v", err)
			continue
		}
		products = append(products, p)
	}
	return products, nil
}

// UpdateProductPrice — обновляет цену продукта
func UpdateProductPrice(db *sql.DB, productID int, newPrice float64) error {
	query := `UPDATE products SET price = $1 WHERE id = $2`
	_, err := db.Exec(query, newPrice, productID)
	if err != nil {
		log.Printf("Ошибка обновления цены продукта: %v", err)
		return err
	}
	return nil
}

// DeleteProduct — удаляет продукт по ID
func DeleteProduct(db *sql.DB, productID int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := db.Exec(query, productID)
	if err != nil {
		log.Printf("Ошибка удаления продукта: %v", err)
		return err
	}
	return nil
}
