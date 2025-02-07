package parser

import (
	"sync"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	var wg sync.WaitGroup
	pool := NewWorkerPool(5) // Создание пула из 5 воркеров

	taskCount := 10
	counter := 0
	mu := sync.Mutex{} // Мьютекс для безопасного инкремента

	// Запуск 10 задач
	for i := 0; i < taskCount; i++ {
		wg.Add(1)
		pool.Submit(func() {
			defer wg.Done()
			mu.Lock()
			counter++ // Увеличиваем счётчик, чтобы убедиться, что все задачи выполнились
			mu.Unlock()
			time.Sleep(100 * time.Millisecond) // Симуляция работы
		})
	}

	// Ожидание завершения всех задач
	wg.Wait()
	pool.Wait()

	// Проверка, что все 10 задач выполнены
	if counter != taskCount {
		t.Errorf("Ожидалось %d задач, но выполнено %d", taskCount, counter)
	}
}
