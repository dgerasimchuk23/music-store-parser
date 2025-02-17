package parser

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// FindCategoryURL парсит страницу каталога и находит ссылку на "Студийные мониторы"
func FindCategoryURL(baseURL string) (string, error) {
	resp, err := http.Get(baseURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("получен статус-код: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	var categoryURL string

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists && strings.Contains(link, "/catalog/pro-audio/studiynye_monitory_i_sabvufery/kontrolnye-studiynye-monitory/") {
			categoryURL = link
		}
	})

	if categoryURL == "" {
		return "", fmt.Errorf("Не удалось найти категорию в каталоге")
	}

	fullURL := "https://doctorhead.ru" + categoryURL
	log.Println("Найден URL категории:", fullURL)

	return fullURL, nil
}
