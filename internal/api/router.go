package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// Создаёт маршруты API
func SetupRouter(dbConn *sql.DB) *gin.Engine {
	router := gin.Default()

	productRoutes := router.Group("/products")
	{
		productRoutes.GET("", GetProducts(dbConn))
		productRoutes.POST("", AddProduct(dbConn))
		productRoutes.PUT(":id", UpdateProductPrice(dbConn))
		productRoutes.DELETE(":id", DeleteProduct(dbConn))
	}

	return router
}
