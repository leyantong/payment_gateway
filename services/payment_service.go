package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"payment_gateway/models"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitializeService(database *gorm.DB) {
	db = database
}

func ProcessPayment(request models.ProcessPaymentRequest) (models.BankResponse, error) {
	log.Println("Service: Start ProcessPayment")

	bankSimulatorURL := os.Getenv("BANK_SIMULATOR_URL")
	jsonValue, _ := json.Marshal(request)
	log.Printf("Service: Sending request to Bank Simulator: %s\n", jsonValue)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", bankSimulatorURL, bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Printf("Service Error: creating new HTTP request: %v\n", err)
		return models.BankResponse{}, errors.New("failed to create HTTP request: " + err.Error())
	}

	req.Close = true
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	bankResp, err := client.Do(req)
	if err != nil {
		log.Printf("Service Error: sending request to Bank Simulator: %v\n", err)
		return models.BankResponse{}, errors.New("bank processing failed: " + err.Error())
	}
	defer bankResp.Body.Close()

	log.Printf("Service: Bank Simulator HTTP status: %s\n", bankResp.Status)
	if bankResp.StatusCode != http.StatusOK {
		log.Printf("Service Error: Bank Simulator returned non-OK status: %s\n", bankResp.Status)
		return models.BankResponse{}, errors.New("bank processing failed with status: " + bankResp.Status)
	}

	var bankResponse models.BankResponse
	if err := json.NewDecoder(bankResp.Body).Decode(&bankResponse); err != nil {
		log.Printf("Service Error: decoding Bank Simulator response: %v\n", err)
		return models.BankResponse{}, errors.New("bank response decoding failed: " + err.Error())
	}

	payment := models.Payment{
		CardNumber:  request.CardNumber[len(request.CardNumber)-4:],
		ExpiryMonth: request.ExpiryMonth,
		ExpiryYear:  request.ExpiryYear,
		Amount:      request.Amount,
		Currency:    request.Currency,
		Status:      bankResponse.Status,
	}

	if err := db.Create(&payment).Error; err != nil {
		log.Printf("Service Error: saving payment to database: %v\n", err)
		return models.BankResponse{}, errors.New("database error: " + err.Error())
	}

	log.Printf("Service: Processed Payment: %+v\n", payment)
	log.Println("Service: End ProcessPayment")
	return bankResponse, nil
}

func RetrievePayment(id string) (*models.Payment, error) {
	log.Println("Service: Start RetrievePayment")
	var payment models.Payment
	if err := db.First(&payment, id).Error; err != nil {
		log.Printf("Service Error: retrieving payment by ID: %v\n", err)
		return nil, err
	}
	log.Printf("Service: Retrieved Payment: %+v\n", payment)
	log.Println("Service: End RetrievePayment")
	return &payment, nil
}
