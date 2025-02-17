package main

import (
	"fmt"
	"log"
	"parser/internal/db"
	"parser/internal/parser"
	"sync"

	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные из .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	log.Println("Запуск парсера...")

	// Подключаемся к базе данных
	dbConn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer dbConn.Close()

	// Очищаем таблицу перед парсингом
	_, err = dbConn.Exec("DELETE FROM products;")
	if err != nil {
		log.Fatalf("Ошибка очистки таблицы: %v", err)
	}
	log.Println("Таблица `products` очищена перед парсингом")

	// Динамически получаем URL категории
	categoryURL, err := parser.FindCategoryURL("https://doctorhead.ru/")
	if err != nil {
		log.Fatalf("Ошибка поиска URL категории: %v", err)
	}

	// Извлекаем ссылки на товары
	productLinks := parser.ParseCategory(categoryURL)

	// Запускаем многопоточный парсинг товаров
	var wg sync.WaitGroup
	for _, link := range productLinks {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			product := parser.ParseProduct(url)
			if product != nil {
				err := db.SaveProduct(dbConn, product.Name, "шт", fmt.Sprintf("%.2f", product.Price), product.URL)
				if err != nil {
					log.Printf("Ошибка сохранения продукта в БД: %v", err)
				}
			}
		}(link)
	}

	wg.Wait()
	log.Println("Парсинг завершен!")
}
