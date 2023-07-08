package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	AppName  string
	Port     string
	MongoUri string
}

func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		return nil
	}
	return &Config{
		AppName:  os.Getenv("APP_NAME"),
		Port:     os.Getenv("PORT"),
		MongoUri: os.Getenv("MONGODB_URI"),
	}
}

var AppConfig = GetConfig()
