package parser

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Product структура данных о товаре
type Product struct {
	Name  string
	Price float64
	URL   string
}

// ParseCategory получает ссылки на товары из каталога
func ParseCategory(url string) []string {
	doc, err := FetchHTML(url)
	if err != nil {
		log.Printf("Ошибка загрузки каталога: %v\n", err)
		return nil
	}

	var productLinks []string
	doc.Find("div.product-image.catalog-product-image-slider > div.swiper-wrapper > div:nth-child(1) > a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists {
			fullLink := "https://doctorhead.ru" + link
			productLinks = append(productLinks, fullLink)
		}
	})

	log.Printf("Найдено %d товаров\n", len(productLinks))
	return productLinks
}

// ParseProduct парсит страницу товара и возвращает объект Product
func ParseProduct(url string) *Product {
	doc, err := FetchHTML(url)
	if err != nil {
		log.Printf("Ошибка загрузки страницы товара: %v\n", err)
		return nil
	}

	// Перебираем возможные селекторы для названия
	var name string
	nameSelectors := []string{
		"h1.product-main-info__title", // Новый селектор
		"body > div.page-container > div.box.box--overlay > div > div:nth-child(8) > h1",
		"h1.product-title",
		"div.product-header h1",
		"h1",
	}

	for _, selector := range nameSelectors {
		name = strings.TrimSpace(doc.Find(selector).First().Text())
		if name != "" {
			break // Если нашли название, выходим
		}
	}

	if name == "" {
		log.Printf("Ошибка: пустое название у товара (URL: %s)", url)
		return nil
	}

	// Попробуем несколько вариантов селекторов для цены
	var priceText string
	priceSelectors := []string{
		"div.product-price.sale-percent > span", // Если скидка
		"div.product-price > span",              // Основная цена
		"span.price",                            // Альтернативный вариант
	}

	for _, selector := range priceSelectors {
		priceText = strings.TrimSpace(doc.Find(selector).First().Text())
		if priceText != "" {
			break // Если нашли цену, выходим из цикла
		}
	}

	if priceText == "" {
		log.Printf("Ошибка: пустая цена у товара %s (URL: %s)", name, url)
		return nil
	}

	// Очищаем цену от пробелов и символов
	priceText = strings.ReplaceAll(priceText, "\u00a0", "") // Убираем неразрывные пробелы
	priceText = strings.ReplaceAll(priceText, "руб.", "")   // Убираем символ рубля
	priceText = strings.ReplaceAll(priceText, " ", "")      // Убираем пробелы между разрядами (чтобы 79 990 не превращалось в 79)
	priceText = strings.ReplaceAll(priceText, ",", ".")     // Меняем запятую на точку
	priceText = strings.TrimSpace(priceText)

	price, err := strconv.ParseFloat(priceText, 64)
	if err != nil {
		log.Printf("Ошибка преобразования цены '%s' у товара %s (URL: %s): %v", priceText, name, url, err)
		return nil
	}

	fmt.Printf("Название: %s\n", name)
	fmt.Printf("Цена: %.2f руб.\n", price)

	return &Product{Name: name, Price: price, URL: url}
}
