package config

import (
	"os"
	"runtime"
	"testing"
)

func createTestConfig(t *testing.T, content string) string {
	file, err := os.CreateTemp("", "config_*.json")
	if err != nil {
		t.Fatal("Ошибка создания временного файла:", err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		t.Fatal("Ошибка записи в файл:", err)
	}
	file.Close()

	return file.Name()
}

func TestLoadConfig(t *testing.T) {
	// Тест 1: worker_pool_size = 0 → должен замениться на runtime.NumCPU() * 2
	testConfig := `{"database_url": "test_db", "worker_pool_size": 0}`
	path := createTestConfig(t, testConfig)
	defer os.Remove(path)

	cfg := LoadConfig(path)

	if cfg.DatabaseURL != "test_db" {
		t.Errorf("Ожидалось database_url = test_db, получено: %s", cfg.DatabaseURL)
	}

	expectedWorkers := runtime.NumCPU() * 2
	if cfg.WorkerPoolSize != expectedWorkers {
		t.Errorf("Ожидалось worker_pool_size = %d, получено: %d", expectedWorkers, cfg.WorkerPoolSize)
	}
}

func TestLoadConfig_CustomWorkers(t *testing.T) {
	// Тест 2: worker_pool_size = 10 → не должен изменяться
	testConfig := `{"database_url": "test_db", "worker_pool_size": 10}`
	path := createTestConfig(t, testConfig)
	defer os.Remove(path)

	cfg := LoadConfig(path)

	if cfg.WorkerPoolSize != 10 {
		t.Errorf("Ожидалось worker_pool_size = 10, получено: %d", cfg.WorkerPoolSize)
	}
}
