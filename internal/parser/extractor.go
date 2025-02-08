package parser

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Структура данных продукта
type ProductData struct {
	Name              string
	UnitOfMeasurement string
	Price             float64
}

// Функция для извлечения данных из HTML-страницы
func ExtractData(doc *goquery.Document) []ProductData {
	var results []ProductData

	doc.Find(".product-item").Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Find(".product-title").Text())
		unit := strings.TrimSpace(s.Find(".product-unit").Text()) // Парсим единицы измерения
		priceText := strings.TrimSpace(s.Find(".product-price").Text())

		// Конвертация цены в число
		price, err := strconv.ParseFloat(strings.Replace(priceText, "₽", "", -1), 64)
		if err != nil {
			price = 0.0
		}

		results = append(results, ProductData{
			Name:              name,
			UnitOfMeasurement: unit,
			Price:             price,
		})
	})

	return results
}
