package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProcessPaymentRequest struct {
	CardNumber  string  `json:"card_number"`
	ExpiryMonth string  `json:"expiry_month"`
	ExpiryYear  string  `json:"expiry_year"`
	CVV         string  `json:"cvv"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
}

func simulateBank(c *gin.Context) {
	var request ProcessPaymentRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Simple validation for simulation
	status := "declined"
	if request.CVV == "123" {
		status = "approved"
	}

	c.JSON(http.StatusOK, gin.H{"status": status})
}

func main() {
	r := gin.Default()
	r.POST("/simulate_bank", simulateBank)
	r.Run(":8081")
}
