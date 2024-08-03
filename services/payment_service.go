package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"payment_gateway/models"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

var (
	db               *gorm.DB
	bankSimulatorURL string
	cache            = make(map[string]time.Time)
	cacheMutex       sync.Mutex
	cacheDuration    = 1 * time.Hour
)

func InitializeService(database *gorm.DB, bankURL string) {
	db = database
	bankSimulatorURL = bankURL
}

func isDuplicatePayment(request models.ProcessPaymentRequest) bool {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	amountStr := strconv.FormatFloat(request.Amount, 'f', 2, 64)
	key := request.CardNumber + request.ExpiryMonth + request.ExpiryYear + request.Currency + amountStr
	if timestamp, found := cache[key]; found {
		if time.Since(timestamp) < cacheDuration {
			return true
		}
	}

	cache[key] = time.Now()
	return false
}

func ProcessPayment(request models.ProcessPaymentRequest) (models.BankResponse, uuid.UUID, error) {
	if bankSimulatorURL == "" {
		return models.BankResponse{}, uuid.Nil, errors.New("BANK_SIMULATOR_URL is not set")
	}

	if isDuplicatePayment(request) {
		return models.BankResponse{}, uuid.Nil, errors.New("duplicate payment detected")
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

	paymentID := uuid.New()
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
