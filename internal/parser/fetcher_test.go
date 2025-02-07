package parser

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Тест успешной загрузки HTML-страницы
func TestFetchHTML_Success(t *testing.T) {
	// Создание тестового HTTP-сервера
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<html><head><title>Test Page</title></head><body><h1>Hello</h1></body></html>`))
	}))
	defer server.Close()

	// Запуск `FetchHTML` с URL тестового сервера
	doc, err := FetchHTML(server.URL)
	if err != nil {
		t.Fatalf("FetchHTML вернул ошибку: %v", err)
	}

	// Проверка, что заголовок страницы соответствует ожидаемому
	title := doc.Find("title").Text()
	expectedTitle := "Test Page"
	if title != expectedTitle {
		t.Errorf("Ожидалось '%s', получено '%s'", expectedTitle, title)
	}
}

// Тест ошибки при неверном URL
func TestFetchHTML_InvalidURL(t *testing.T) {
	_, err := FetchHTML("http://invalid-url")
	if err == nil {
		t.Error("Ожидалась ошибка при загрузке недоступного URL, но ошибка не была возвращена")
	}
}

// Тест ошибки при неуспешном HTTP-ответе (например, 500)
func TestFetchHTML_HttpError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	_, err := FetchHTML(server.URL)
	if err == nil {
		t.Error("Ожидалась ошибка при получении HTTP 500, но ошибка не была возвращена")
	}
}
