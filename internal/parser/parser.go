package parser

import (
	"fmt"
	"log"
	"strings"
)

// Parser — структура парсера
type Parser struct{}

// ParsePage — функция для парсинга страницы товара
func (p *Parser) ParsePage(url string) {
	// Загружаем страницу через обычный HTTP запрос
	doc, err := FetchHTML(url)
	if err != nil {
		log.Fatalf("Ошибка загрузки страницы: %v", err)
	}

	// Извлекаем данные с страницы
	name := doc.Find("h1.quick-view-title").Text() // Извлекаем название товара
	priceText := doc.Find("span.nowrap").Text()    // Извлекаем цену товара

	// Преобразуем и выводим данные
	priceText = strings.ReplaceAll(priceText, "₽", "")
	priceText = strings.ReplaceAll(priceText, " ", "")
	fmt.Println("Название товара:", name)
	fmt.Println("Цена товара:", priceText)
}
