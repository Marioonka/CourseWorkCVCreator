package config

import "github.com/ilyakaznacheev/cleanenv"

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
	var cfg Config
	err := cleanenv.ReadConfig("/home/marina/GolandProjects/coursework/cmd/config1.yaml", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
