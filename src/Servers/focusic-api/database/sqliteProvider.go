package database

import (
	"focusic-api/config"
	"focusic-api/models/domains"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectSqlite() {
	dbName := config.AppConfig.SqliteDbName
	database, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Starting migration...")
	err = database.AutoMigrate(&domains.User{})
	if err != nil {
		log.Fatalln("Migration failed", err)
	}
	log.Println("Migration completed!")
	DB = database
}
