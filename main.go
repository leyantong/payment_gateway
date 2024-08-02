package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"

	"payment_gateway/models"
	"payment_gateway/router"
	"payment_gateway/services"
)

func main() {
	log.Println("Starting main function")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bankSimulatorURL := os.Getenv("BANK_SIMULATOR_URL")
	if bankSimulatorURL == "" {
		log.Fatal("BANK_SIMULATOR_URL is not set")
	}

	// Initialize the database
	db, err := gorm.Open("sqlite3", "payments.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}
	defer db.Close()

	// Auto-migrate the Payment model
	db.AutoMigrate(&models.Payment{})

	// Initialize the service with the database
	services.InitializeService(db)

	// Initialize the router
	r := router.SetupRouter()

	log.Println("Starting server on port", os.Getenv("PORT"))
	r.Run(os.Getenv("PORT"))
}
