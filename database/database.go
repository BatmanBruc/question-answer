package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	dsn := os.Getenv("DATABASE_URL")

	var err error
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("PostgreSQL connected successfully")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
