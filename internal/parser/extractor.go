package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Функция для извлечения данных из HTML-страницы
func ExtractData(doc *goquery.Document) []string {
	var results []string

	// Поиск заголовков товаров (заменить селектор на нужный)
	doc.Find(".product-title").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Text()) // Убираем лишние пробелы
		results = append(results, title)
	})

	return results
}
