package config

import (
	"dco_mart/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
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
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
	// migrateDatabase(DB)
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
