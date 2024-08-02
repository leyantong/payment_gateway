package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	// Initialize the Gin router
	r := gin.Default()

	// Initialize the database
	var err error
	db, err = gorm.Open("sqlite3", "payments.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Auto migrate the Payment model
	db.AutoMigrate(&Payment{})

	// Set up routes
	r.POST("/process_payment", processPayment)
	r.GET("/retrieve_payment/:id", retrievePayment)

	// Run the server
	r.Run(":8083") // Default listens and serves on 0.0.0.0:8080
}
