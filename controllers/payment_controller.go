package controllers

import (
	"log"
	"net/http"
	"payment_gateway/models"
	"payment_gateway/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ProcessPayment(c *gin.Context) {
	log.Println("Controller: Start ProcessPayment")

	request, exists := c.Get("payment_request")
	if !exists {
		log.Println("Controller Error: Payment request not found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	paymentRequest, ok := request.(models.ProcessPaymentRequest)
	if !ok {
		log.Println("Controller Error: Invalid payment request type")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	log.Printf("Controller: Processing Payment: %+v\n", paymentRequest)

	response, paymentID, err := services.ProcessPayment(paymentRequest)
	if err != nil {
		log.Printf("Controller Service Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Controller: Processed Payment Response: %+v\n", response)
	c.JSON(http.StatusOK, gin.H{
		"status": response.Status,
		"id":     paymentID,
	})
}

func RetrievePayment(c *gin.Context) {
	log.Println("Controller: Start RetrievePayment")

	idParam := c.Param("id")
	paymentID, err := uuid.Parse(idParam)
	if err != nil {
		log.Printf("Controller Error: Invalid UUID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	log.Printf("Controller: Retrieving Payment ID: %s\n", idParam)
	payment, err := services.RetrievePayment(paymentID)
	if err != nil {
		log.Printf("Controller RetrievePayment Error: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	response := gin.H{
		"payment_id":   payment.ID,
		"card_number":  "**** **** **** " + payment.CardNumber[len(payment.CardNumber)-4:], // Masked card number
		"expiry_month": payment.ExpiryMonth,
		"expiry_year":  payment.ExpiryYear,
		"amount":       payment.Amount,
		"currency":     payment.Currency,
		"status":       payment.Status,
	}

	log.Printf("Controller: Retrieved Payment Response: %+v\n", response)
	c.JSON(http.StatusOK, response)
}
