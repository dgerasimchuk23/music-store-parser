package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Тестируем поиск категории товаров
func TestFindCategoryURL(t *testing.T) {
	baseURL := "https://doctorhead.ru/catalog/"
	categoryURL, err := FindCategoryURL(baseURL)

	assert.Nil(t, err, "Ошибка поиска категории")
	assert.Contains(t, categoryURL, "kontrolnye-studiynye-monitory", "Ошибка: URL категории неверный")
}
