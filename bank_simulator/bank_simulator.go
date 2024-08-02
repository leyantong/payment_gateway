package main

import (
	"log"
	"math/rand"
	"net/http"
	"payment_gateway/models"
	"time"

	"github.com/gin-gonic/gin"
)

func simulateBank(c *gin.Context) {
	var request models.ProcessPaymentRequest
	if err := c.BindJSON(&request); err != nil {
		log.Printf("SimulateBank Error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	rand.Seed(time.Now().UnixNano())
	success := rand.Float32() > 0.2 // 80% chance of success
	status := "DECLINED"
	if success {
		status = "APPROVED"
	}

	log.Printf("SimulateBank Processed Request: %+v\n", request)
	c.JSON(http.StatusOK, gin.H{
		"status":             status,
		"masked_card_number": "**** **** **** " + request.CardNumber[len(request.CardNumber)-4:],
	})
}

func main() {
	router := gin.Default()
	router.POST("/simulate_bank", simulateBank)
	log.Fatal(router.Run(":8081")) // Running on port 8081 for simulation
}
