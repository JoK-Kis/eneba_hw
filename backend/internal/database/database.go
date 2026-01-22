package database

import (
	"fmt"
	"log"
	"os"

	"github.com/JoK-Kis/eneba-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	sslmode := os.Getenv("DB_SSLMODE")

	var dsn string
	if password == "" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s", 
			host, port, user, dbname, sslmode)
	} else {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", 
			host, port, user, password, dbname, sslmode)
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

func Migrate() {
	if err := DB.Exec("CREATE EXTENSION IF NOT EXISTS pg_trgm").Error; err != nil {
		log.Fatal("Failed to create pg_trgm extension:", err)
	}

	err := DB.AutoMigrate(
		&models.Platform{},
		&models.Product{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
}