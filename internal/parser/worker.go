package parser

import (
	"sync"
)

type WorkerPool struct {
	wg sync.WaitGroup
}

func NewWorkerPool(size int) *WorkerPool {
	return &WorkerPool{}
}

// Добавление задачи в пул
func (wp *WorkerPool) Submit(task func()) {
	wp.wg.Add(1)
	go func() {
		defer wp.wg.Done()
		task()
	}()
}

// Ожидаем завершения всех задач
func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}
