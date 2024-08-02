package controllers

import (
	"log"
	"net/http"
	"payment_gateway/models"
	"payment_gateway/services"

	"github.com/gin-gonic/gin"
)

func ProcessPayment(c *gin.Context) {
	log.Println("Controller: Start ProcessPayment")

	var request models.ProcessPaymentRequest
	if err := c.BindJSON(&request); err != nil {
		log.Printf("Controller Error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	log.Printf("Controller: Processing Payment: %+v\n", request)

	response, err := services.ProcessPayment(request)
	if err != nil {
		log.Printf("Controller Service Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Controller: Processed Payment Response: %+v\n", response)
	c.JSON(http.StatusOK, response)
}

func RetrievePayment(c *gin.Context) {
	log.Println("Controller: Start RetrievePayment")

	id := c.Param("id")
	log.Printf("Controller: Retrieving Payment ID: %s\n", id)
	payment, err := services.RetrievePayment(id)
	if err != nil {
		log.Printf("Controller RetrievePayment Error: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	// Mask card number
	maskedCardNumber := "**** **** **** " + payment.CardNumber[len(payment.CardNumber)-4:]
	response := gin.H{
		"payment_id":   payment.ID,
		"card_number":  maskedCardNumber,
		"expiry_month": payment.ExpiryMonth,
		"expiry_year":  payment.ExpiryYear,
		"amount":       payment.Amount,
		"currency":     payment.Currency,
		"status":       payment.Status,
	}

	log.Printf("Controller: Retrieved Payment Response: %+v\n", response)
	c.JSON(http.StatusOK, response)
}
