package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

// парс данных из yml

type Config struct {
	DB struct {
		Url string `json:"url"`
	} `yaml:"db"`

	HTTP struct {
		Port string `yaml:"port"`
	} `yaml:"http"`

	Redis struct {
		Address  string `yaml:"address"`
		Password string `yaml:"password"`
		DBName   int    `yaml:"db"`
	} `yaml:"redis"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	fmt.Println(cfg)

	return cfg, nil
}
