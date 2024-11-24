package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
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
	// Получаем путь к текущему файлу исходного кода
	_, filename, _, _ := runtime.Caller(0)
	// Поднимаемся на уровень папки проекта
	projectRoot := filepath.Join(filepath.Join(filepath.Dir(filename), ".."), "..")

	configPath := filepath.Join(projectRoot, "cmd", "config1.yaml")

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	// Затем переменные окружения переопределят значения из файла
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
