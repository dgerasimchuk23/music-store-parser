package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"parser/internal/db"

	"github.com/gin-gonic/gin"
)

// Обработка запроса на получение списка товаров
func GetProducts(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := dbConn.Query("SELECT id, name, unit_of_measurement, price, url, created_at FROM products")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения списка товаров"})
			return
		}
		defer rows.Close()

		var products []db.Product
		for rows.Next() {
			var product db.Product
			err := rows.Scan(&product.ID, &product.Name, &product.UnitOfMeasurement, &product.Price, &product.URL, &product.CreatedAt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обработки данных товаров"})
				return
			}
			products = append(products, product)
		}

		c.JSON(http.StatusOK, products)
	}
}

// Обработка запроса на добавление нового товара
func AddProduct(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product db.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
			return
		}

		err := db.SaveProduct(dbConn, product.Name, product.UnitOfMeasurement, strconv.FormatFloat(product.Price, 'f', 2, 64), product.URL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка добавления продукта"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Продукт добавлен"})
	}
}

// Обновление цены
func UpdateProductPrice(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Цена обновлена (заглушка)"})
	}
}

// Удаление продукта
func DeleteProduct(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Продукт удалён (заглушка)"})
	}
}
