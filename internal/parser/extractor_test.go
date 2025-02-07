package parser

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

// Тест извлечения данных из HTML
func TestExtractData(t *testing.T) {
	// Пример HTML-контента
	htmlContent := `<html>
		<head><title>Test Page</title></head>
		<body>
			<div class="product-title">Товар 1</div>
			<div class="product-title">Товар 2</div>
			<div class="product-title">Товар 3</div>
		</body>
	</html>`

	// Создание объекта goquery.Document из HTML-строки
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		t.Fatalf("Ошибка при создании goquery.Document: %v", err)
	}

	// Запуск ExtractData
	results := ExtractData(doc)

	// Ожидание результата
	expectedResults := []string{"Товар 1", "Товар 2", "Товар 3"}

	// Проверка количество найденных элементов
	if len(results) != len(expectedResults) {
		t.Errorf("Ожидалось %d результатов, получено %d", len(expectedResults), len(results))
	}

	// Проверка содержимого
	for i, expected := range expectedResults {
		if results[i] != expected {
			t.Errorf("Ожидалось '%s', получено '%s'", expected, results[i])
		}
	}
}

// Тест на пустую страницу
func TestExtractData_EmptyPage(t *testing.T) {
	htmlContent := `<html><body></body></html>`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		t.Fatalf("Ошибка при создании goquery.Document: %v", err)
	}

	results := ExtractData(doc)

	// Ожидаем пустой массив
	if len(results) != 0 {
		t.Errorf("Ожидалось 0 элементов, получено %d", len(results))
	}
}
