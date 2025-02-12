package parser

import (
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ExtractData извлекает ссылки на страницы товаров
func ExtractData(doc *goquery.Document) []string {
	var productLinks []string

	// Ищем все товары на странице
	doc.Find("div.catalog-list_item.product.check-show").Each(func(i int, s *goquery.Selection) {
		// Находим ссылку на товар
		link, exists := s.Find("a").Attr("href")
		if exists {
			fullLink := "https://doctorhead.ru" + link // Добавляем базовый URL
			productLinks = append(productLinks, fullLink)
		}
	})

	log.Printf("Найдено %d ссылок на товары\n", len(productLinks))
	return productLinks
}

// ExtractProductDetails извлекает детали товара с его страницы
func ExtractProductDetails(doc *goquery.Document) (string, float64) {
	// Извлекаем название товара
	name := strings.TrimSpace(doc.Find("h1.quick-view-title").Text())

	// Извлекаем цену товара
	priceText := strings.TrimSpace(doc.Find("span.nowrap").Text())
	priceText = strings.ReplaceAll(priceText, "₽", "") // Убираем символ рубля
	priceText = strings.ReplaceAll(priceText, " ", "") // Убираем пробелы
	price, err := strconv.ParseFloat(priceText, 64)
	if err != nil {
		log.Printf("Ошибка преобразования цены: %v\n", err)
		price = 0
	}

	return name, price
}
