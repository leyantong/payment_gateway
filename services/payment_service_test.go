package services

import (
	"testing"

	"payment_gateway/models"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	db.AutoMigrate(&models.Payment{})
	return db
}

func TestProcessPayment(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	InitializeService(db, "http://localhost:8081/simulate_bank")

	request := models.ProcessPaymentRequest{
		CardNumber:  "4242424242424242",
		ExpiryMonth: "12",
		ExpiryYear:  "2024",
		CVV:         "123",
		Amount:      100.00,
		Currency:    "USD",
	}

	response, paymentID, err := ProcessPayment(request)
	assert.Nil(t, err)
	assert.NotEqual(t, uuid.Nil, paymentID)
	assert.Equal(t, "APPROVED", response.Status)
}

func TestRetrievePayment(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	InitializeService(db, "http://localhost:8081/simulate_bank")

	paymentID := uuid.New()
	payment := models.Payment{
		ID:          paymentID,
		CardNumber:  "4242",
		ExpiryMonth: "12",
		ExpiryYear:  "2024",
		Amount:      100.00,
		Currency:    "USD",
		Status:      "APPROVED",
	}
	db.Create(&payment)

	retrievedPayment, err := RetrievePayment(paymentID)
	assert.Nil(t, err)
	assert.Equal(t, paymentID, retrievedPayment.ID)
	assert.Equal(t, payment.Amount, retrievedPayment.Amount)
	assert.Equal(t, payment.Currency, retrievedPayment.Currency)
	assert.Equal(t, payment.Status, retrievedPayment.Status)
}
