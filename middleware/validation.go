package middleware

import (
	"log"
	"net/http"
	"payment_gateway/models"

	"github.com/gin-gonic/gin"
)

func ValidatePaymentRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Middleware: Start ValidatePaymentRequest")
		var request models.ProcessPaymentRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			log.Printf("Validation Error: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
			c.Abort()
			return
		}

		if request.CardNumber == "" || request.ExpiryMonth == "" || request.ExpiryYear == "" || request.CVV == "" || request.Amount <= 0 || request.Currency == "" {
			log.Println("Validation Error: Missing fields or invalid amount")
			c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required and amount should be greater than zero"})
			c.Abort()
			return
		}

		log.Printf("Validated Request: %+v\n", request)
		c.Set("payment_request", request)
		log.Println("Middleware: End ValidatePaymentRequest")
		c.Next()
	}
}
