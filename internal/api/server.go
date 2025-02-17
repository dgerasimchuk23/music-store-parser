package api

import (
	"database/sql"
	"log"
)

func StartServer(dbConn *sql.DB) {
	router := SetupRouter(dbConn)

	log.Println("Запуск API на порту 8080")
	router.Run(":8080")
}
