package config

import (
	"api-gateway/internal/model"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Routes []model.Route `yaml:"routes"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
