package config

import (
	"dco_mart/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Construct the DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=require password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
	if os.Getenv("DB_SYNC") == "true" {
		migrateDatabase(DB)
	}
	log.Println("Database connection established.")
}

func migrateDatabase(db *gorm.DB) {
	db.Migrator().DropTable(
		&models.User{},
		&models.Order{},
		&models.OrderDetail{},
		&models.Item{},
		&models.Category{},
	)
	db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Item{},
		&models.Order{},
		&models.OrderDetail{},
	)
}
