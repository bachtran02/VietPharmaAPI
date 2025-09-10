package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

func ReadConfig(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("failed to open config: %w", err)
	}
	defer file.Close()

	cfg := Config{
		Server: ServerConfig{
			Port: 8080, // default port
		},
	}
	if err = yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("failed to decode config: %w", err)
	}
	return cfg, nil
}
