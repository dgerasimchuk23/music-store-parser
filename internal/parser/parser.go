package parser

import (
	"database/sql"
	"fmt"
	"parser/internal/db"

	"go.uber.org/zap"
)

type Parser struct {
	WorkerPool *WorkerPool
	Logger     *zap.Logger
	DB         *sql.DB
}

func NewParser(workerPool *WorkerPool, logger *zap.Logger, dbConn *sql.DB) *Parser {
	return &Parser{
		WorkerPool: workerPool,
		Logger:     logger,
		DB:         dbConn,
	}
}

// Парсинг данных с нескольких страниц
func (p *Parser) Parse(urls []string) {
	p.Logger.Info("Начало парсинга")

	for _, url := range urls {
		url := url
		p.WorkerPool.Submit(func() {
			p.Logger.Info("Парсинг страницы", zap.String("URL", url))

			// Загрузка HTML страницы
			doc, err := FetchHTML(url)
			if err != nil {
				p.Logger.Error("Ошибка загрузки страницы", zap.String("URL", url), zap.Error(err))
				return
			}

			// Извлечение данных
			items := ExtractData(doc)
			for _, item := range items {
				fmt.Println("Найдено:", item.Name, item.UnitOfMeasurement, item.Price)

				// Добавление продуктов в БД
				err := db.AddProduct(p.DB, db.Product{
					Name:              item.Name,
					UnitOfMeasurement: item.UnitOfMeasurement,
					Price:             item.Price,
					URL:               url,
				})
				if err != nil {
					p.Logger.Error("Ошибка сохранения продукта в БД", zap.Error(err))
				}
			}
		})
	}

	// Ожидание завершения всех задач
	p.WorkerPool.Wait()
}
