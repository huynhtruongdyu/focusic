package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	AppName string
	Port    string
}

func GetConfig() (*Config, error) {
	err := loadEnv()
	if err != nil {
		return nil, err
	}
	return &Config{
		AppName: os.Getenv("APP_NAME"),
		Port:    os.Getenv("PORT"),
	}, nil
}

func loadEnv() error {
	return godotenv.Load()
}
