package database

import (
	"fmt"
	"log"
	"os"

	"github.com/biggaji/ggsays/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbClient struct {
	*gorm.DB
}

var Client DbClient

func Connect() {
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSslMode := os.Getenv("DB_SSL_MODE")
	dbTimeZone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", dbHost, dbUser, dbPassword, dbName, dbPort, dbSslMode, dbTimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	db.Logger.LogMode(logger.Info)
	db.Logger.LogMode(logger.Error)
	log.Println("Running migrations")

	db.AutoMigrate(&models.User{}, &models.Post{})

	// Assigns the database client
	Client = DbClient{db}
}
