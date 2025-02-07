package parser

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Загружает страницу по URL и возвращает объект goquery.Document
func FetchHTML(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Ошибка загрузки страницы: %v", err)
	}
	defer resp.Body.Close()

	// Проверка кода ответа
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Ошибка: статус-код %d", resp.StatusCode)
	}

	// Создание goquery.Document из тела ответа
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Ошибка парсинга HTML: %v", err)
	}

	return doc, nil
}
