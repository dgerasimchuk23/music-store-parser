package main

import (
	"log"
	"parser/internal/db"
	"parser/internal/parser"

	"github.com/joho/godotenv"
)

func main() {
	// Загрузка переменных окружения
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла:", err)
	}

	// Подключение к базе данных
	dbConn, err := db.ConnectDB() // Используем функцию ConnectDB
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	defer dbConn.Close()

	// Проверка текущей базы данных
	log.Println("Текущая база данных:")
	row := dbConn.QueryRow("SELECT current_database();")
	var currentDB string
	if err := row.Scan(&currentDB); err != nil {
		log.Fatalf("Ошибка получения текущей базы: %v", err)
	}
	log.Println("Подключен к базе:", currentDB)

	// Инициализация схемы базы данных
	log.Println("Запуск инициализации схемы...")
	if err := db.InitializeSchema(dbConn); err != nil {
		log.Fatalf("Ошибка инициализации схемы базы данных: %v", err)
	}
	log.Println("Схема успешно создана!")

	// URL страницы товара для парсинга
	url := "https://doctorhead.ru/product/amphion_one25a/"

	// Создаем экземпляр парсера
	p := parser.Parser{}

	// Запускаем парсинг страницы
	p.ParsePage(url)

	// Примерный вывод:
	// Название товара: Amphon One25A Left
	// Цена товара: 630 000 ₽
}
