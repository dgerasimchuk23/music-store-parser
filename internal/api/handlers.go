package api

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"parser/internal/db"

	"github.com/gin-gonic/gin"
)

// Обработка запроса на получение списка товаров
func GetProducts(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := db.GetProducts(dbConn)
		if err != nil {
			log.Printf("Ошибка получения продуктов: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения данных"})
			return
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

		if err := db.AddProduct(dbConn, product); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка добавления продукта"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Продукт добавлен"})
	}
}

// Обрабатывает обновление цены продукта
func UpdateProductPrice(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
			return
		}

		var req struct {
			Price float64 `json:"price"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
			return
		}

		err = db.UpdateProductPrice(dbConn, id, req.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления цены"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Цена обновлена"})
	}
}

// Обрабатывает удаление продукта
func DeleteProduct(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
			return
		}

		err = db.DeleteProduct(dbConn, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления продукта"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Продукт удалён"})
	}
}
