package config

import (
	"coursework/helpers"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
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
	path, err := helpers.GetPathToFile("config1.yaml")
	if err != nil {
		return nil, err
	}
	// Итеративный поиск конфигурационного файла
	var cfg Config
	err = cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config from %s: %w", path, err)
	}

	return &cfg, nil
}
