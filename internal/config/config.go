package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"path/filepath"
	"runtime"
)

type DbConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Port     string `yaml:"port"`
}

type Config struct {
	Database DbConfig `yaml:"database"`
}

func LoadConfig() (*Config, error) {
	// Список потенциальных путей к конфигурационному файлу
	possiblePaths := []string{}

	// 1. Путь из переменной окружения
	if envPath := os.Getenv("CONFIG_PATH"); envPath != "" {
		possiblePaths = append(possiblePaths, envPath)
	}

	// 2. Текущая рабочая директория
	if cwd, err := os.Getwd(); err == nil {
		possiblePaths = append(possiblePaths, filepath.Join(cwd, "config1.yaml"))
	}

	// 3. Директория исполняемого файла
	if execPath, err := os.Executable(); err == nil {
		execDir := filepath.Dir(execPath)
		possiblePaths = append(possiblePaths, filepath.Join(execDir, "config1.yaml"))
	}

	// 4. Поиск относительно исходного кода (для dev-среды)
	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(filename), "..", "..")
	possiblePaths = append(possiblePaths, filepath.Join(projectRoot, "cmd", "config1.yaml"))

	// Итеративный поиск конфигурационного файла
	var cfg Config
	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			// Файл найден
			err := cleanenv.ReadConfig(path, &cfg)
			if err != nil {
				return nil, fmt.Errorf("error reading config from %s: %w", path, err)
			}
			break
		}
	}

	// Проверяем, был ли загружен конфиг
	if (Config{}) == cfg {
		return nil, fmt.Errorf("config file not found in paths: %v", possiblePaths)
	}

	return &cfg, nil
}
