package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	AppName      string
	Port         string
	SqliteDbName string
}

func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		return nil
	}
	return &Config{
		AppName:      os.Getenv("APP_NAME"),
		Port:         os.Getenv("PORT"),
		SqliteDbName: os.Getenv("SQLITE_DB_NAME"),
	}
}

var AppConfig = GetConfig()
