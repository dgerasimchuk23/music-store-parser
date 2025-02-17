package parser

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// FetchHTML загружает HTML-страницу по URL и возвращает goquery документ
func FetchHTML(url string) (*goquery.Document, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	// Делаем GET-запрос
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Ошибка HTTP-запроса: %w", err)
	}
	defer resp.Body.Close()

	// Проверяем статус код ответа
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Ошибка: сервер вернул статус %d", resp.StatusCode)
	}

	// Загружаем HTML в goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Ошибка парсинга HTML: %w", err)
	}

	return doc, nil
}
