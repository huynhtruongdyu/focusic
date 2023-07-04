package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	AppName string
	Port    string
}

func GetConfig() *Config {
	err := loadEnv()
	if err != nil {
		panic(err)
	}
	return &Config{
		AppName: os.Getenv("APP_NAME"),
		Port:    os.Getenv("PORT"),
	}
}

func loadEnv() error {
	return godotenv.Load()
}
