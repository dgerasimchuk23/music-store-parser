package parser

import (
	"fmt"

	"go.uber.org/zap"
)

type Parser struct {
	WorkerPool *WorkerPool
	Logger     *zap.Logger
}

func NewParser(workerPool *WorkerPool, logger *zap.Logger) *Parser {
	return &Parser{WorkerPool: workerPool, Logger: logger}
}

// Парсинг данных с нескольких страниц
func (p *Parser) Parse(urls []string) {
	p.Logger.Info("Начало парсинга")

	// Параллельный запуск задач
	for _, url := range urls {
		url := url // новая копия переменной (избегаем проблем с замыканием)
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
				fmt.Println("Найдено:", item)
			}
		})
	}

	// Ожидание завершения всех задач
	p.WorkerPool.Wait()
}
