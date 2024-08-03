package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"payment_gateway/models"
	"payment_gateway/utils"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var bankSimulatorURL string

func InitializeService(database *gorm.DB, bankURL string) {
	db = database
	bankSimulatorURL = bankURL
}

func ProcessPayment(request models.ProcessPaymentRequest) (models.BankResponse, uuid.UUID, error) {
	if bankSimulatorURL == "" {
		return models.BankResponse{}, uuid.Nil, errors.New("BANK_SIMULATOR_URL is not set")
	}

	jsonValue, _ := json.Marshal(request)
	req, err := http.NewRequest("POST", bankSimulatorURL, bytes.NewBuffer(jsonValue))
	if err != nil {
		return models.BankResponse{}, uuid.Nil, errors.New("failed to create HTTP request: " + err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	bankResp, err := client.Do(req)
	if err != nil {
		return models.BankResponse{}, uuid.Nil, errors.New("bank processing failed: " + err.Error())
	}
	defer bankResp.Body.Close()

	if bankResp.StatusCode != http.StatusOK {
		return models.BankResponse{}, uuid.Nil, errors.New("bank processing failed with status: " + bankResp.Status)
	}

	var bankResponse models.BankResponse
	if err := json.NewDecoder(bankResp.Body).Decode(&bankResponse); err != nil {
		return models.BankResponse{}, uuid.Nil, errors.New("bank response decoding failed: " + err.Error())
	}

	// Generate a payment UUID using the improved method
	paymentID := utils.GeneratePaymentUUID(request.CardNumber, request.Amount)
	payment := models.Payment{
		ID:          paymentID,
		CardNumber:  request.CardNumber[len(request.CardNumber)-4:], // Mask card number
		ExpiryMonth: request.ExpiryMonth,
		ExpiryYear:  request.ExpiryYear,
		Amount:      request.Amount,
		Currency:    request.Currency,
		Status:      bankResponse.Status,
	}

	if err := db.Create(&payment).Error; err != nil {
		return models.BankResponse{}, uuid.Nil, errors.New("database error: " + err.Error())
	}

	return bankResponse, paymentID, nil
}

func RetrievePayment(id uuid.UUID) (*models.Payment, error) {
	var payment models.Payment
	if err := db.Where("id = ?", id).First(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}
