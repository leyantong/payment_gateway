package main

import (
	"log"
	"net/http"
	"payment_gateway/models"

	"github.com/gin-gonic/gin"
)

func simulateBank(c *gin.Context) {
	var request models.ProcessPaymentRequest
	if err := c.BindJSON(&request); err != nil {
		log.Printf("SimulateBank Error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	status := "declined"
	if request.CVV == "123" {
		status = "approved"
	}

	log.Printf("SimulateBank Processed Request: %+v\n", request)
	c.JSON(http.StatusOK, gin.H{"status": status})
}

func main() {
	router := gin.Default()
	router.POST("/simulate_bank", simulateBank)
	log.Fatal(router.Run(":8081")) // Running on port 8081 for simulation
}
