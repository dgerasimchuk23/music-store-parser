package parser

import (
	"database/sql"
	"os"
	"strings"
	"testing"

	"parser/internal/db"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func setupTestDB(t *testing.T) *sql.DB {
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "newpassword")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "test_parser_db")

	testDB, err := db.ConnectDB()
	if err != nil {
		t.Fatalf("Ошибка подключения к тестовой базе данных: %v", err)
	}

	if err := db.InitializeSchema(testDB); err != nil {
		t.Fatalf("Ошибка инициализации схемы в тестовой базе: %v", err)
	}

	return testDB
}

func TestParserIntegration(t *testing.T) {
	testDB := setupTestDB(t)
	defer testDB.Close()

	logger := zap.NewExample()
	workerPool := NewWorkerPool(1)

	p := NewParser(workerPool, logger, testDB)

	p.Parse([]string{"https://example.com/test"})

	html := `<html>
		<div class="product-item">
			<h2 class="product-title">Тестовый продукт</h2>
			<span class="product-unit">шт</span>
			<span class="product-price">999.99 ₽</span>
		</div>
	</html>`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatalf("Ошибка создания goquery.Document: %v", err)
	}

	products := ExtractData(doc)
	if len(products) == 0 {
		t.Fatal("Ошибка: не извлечены продукты из тестового HTML")
	}

	err = db.AddProduct(testDB, db.Product{
		Name:              products[0].Name,
		UnitOfMeasurement: products[0].UnitOfMeasurement,
		Price:             products[0].Price,
		URL:               "https://example.com/test",
	})
	if err != nil {
		t.Fatalf("Ошибка добавления продукта в базу: %v", err)
	}

	storedProducts, err := db.GetProducts(testDB)
	if err != nil {
		t.Fatalf("Ошибка получения списка продуктов: %v", err)
	}

	if len(storedProducts) == 0 {
		t.Fatal("Ошибка: продукт не добавлен в базу")
	}
}
