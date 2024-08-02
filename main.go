package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"

	"payment_gateway/config"
	"payment_gateway/models"
	"payment_gateway/router"
	"payment_gateway/services"
)

func main() {
	log.Println("Starting application")

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load configuration
	config.LoadConfig()

	// Initialize the database
	db, err := gorm.Open("sqlite3", "payments.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Auto-migrate the Payment model
	if err := db.AutoMigrate(&models.Payment{}).Error; err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}

	// Initialize the service with the database
	services.InitializeService(db, config.AppConfig.BankSimulatorURL)

	// Initialize the router
	r := router.SetupRouter()

	log.Printf("Starting server on port %s", config.AppConfig.Port)
	if err := r.Run(":" + config.AppConfig.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
