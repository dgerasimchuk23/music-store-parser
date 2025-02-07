package config

import (
	"encoding/json"
	"log"
	"os"
	"runtime"
)

type Config struct {
	DatabaseURL    string `json:"database_url"`
	WorkerPoolSize int    `json:"worker_pool_size"`
}

func LoadConfig(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Ошибка загрузки конфига:", err)
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		log.Fatal("Ошибка декодирования JSON:", err)
	}

	// Авто-определение пула воркеров, если в конфиге 0
	if config.WorkerPoolSize == 0 {
		config.WorkerPoolSize = runtime.NumCPU() * 2
	}

	return &config
}
