package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

type ProcessPaymentRequest struct {
	CardNumber  string  `json:"card_number"`
	ExpiryMonth string  `json:"expiry_month"`
	ExpiryYear  string  `json:"expiry_year"`
	CVV         string  `json:"cvv"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
}

type BankResponse struct {
	Status string `json:"status"`
}

func processPayment(c *gin.Context) {
	var request ProcessPaymentRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Simulate bank response using the mock bank simulator
	bankSimulatorURL := "http://localhost:8081/simulate_bank"
	jsonValue, _ := json.Marshal(request)
	bankResp, err := http.Post(bankSimulatorURL, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil || bankResp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Bank processing failed"})
		return
	}

	var bankResponse BankResponse
	if err := json.NewDecoder(bankResp.Body).Decode(&bankResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Bank response decoding failed"})
		return
	}

	payment := Payment{
		CardNumber:  request.CardNumber[len(request.CardNumber)-4:],
		ExpiryMonth: request.ExpiryMonth,
		ExpiryYear:  request.ExpiryYear,
		Amount:      request.Amount,
		Currency:    request.Currency,
		Status:      bankResponse.Status,
	}

	db.Create(&payment)

	c.JSON(http.StatusOK, gin.H{"payment_id": payment.ID, "status": payment.Status})
}

func retrievePayment(c *gin.Context) {
	id := c.Param("id")
	var payment Payment
	if db.First(&payment, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payment_id":   payment.ID,
		"card_number":  "**** **** **** " + payment.CardNumber,
		"expiry_month": payment.ExpiryMonth,
		"expiry_year":  payment.ExpiryYear,
		"amount":       payment.Amount,
		"currency":     payment.Currency,
		"status":       payment.Status,
	})
}
