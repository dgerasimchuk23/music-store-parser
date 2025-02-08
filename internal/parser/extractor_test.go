package parser

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestExtractData(t *testing.T) {
	htmlContent := `<html>
		<div class="product-item">
			<h2 class="product-title">Тестовый продукт</h2>
			<span class="product-unit">шт</span>
			<span class="product-price">999.99 ₽</span>
		</div>
	</html>`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		t.Fatalf("Ошибка при создании goquery.Document: %v", err)
	}

	results := ExtractData(doc)

	expectedResults := []ProductData{
		{Name: "Тестовый продукт", UnitOfMeasurement: "шт", Price: 999.99},
	}

	if len(results) != len(expectedResults) {
		t.Errorf("Ожидалось %d результатов, получено %d", len(expectedResults), len(results))
	}

	for i, expected := range expectedResults {
		if results[i].Name != expected.Name {
			t.Errorf("Ожидалось '%s', получено '%s'", expected.Name, results[i].Name)
		}
	}
}

func TestExtractData_EmptyPage(t *testing.T) {
	htmlContent := `<html><body></body></html>`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		t.Fatalf("Ошибка при создании goquery.Document: %v", err)
	}

	results := ExtractData(doc)

	if len(results) != 0 {
		t.Errorf("Ожидалось 0 элементов, получено %d", len(results))
	}
}
