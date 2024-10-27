package config

import (
	"log"
	"os"
	"testku/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitializeDB ConnectDB initializes the database connection and performs auto migration
func InitializeDB() (*gorm.DB, error) {
	var err error

	DB, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil, err
	}

	// Perform auto migration for all entities
	err = DB.AutoMigrate(
		&entities.User{},
		&entities.Wallet{},
		&entities.Product{},
		&entities.Transaction{},
		&entities.Order{},
	)
	if err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
		return nil, err
	}

	return DB, nil
}
