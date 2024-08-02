package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"payment_gateway/models"
	"payment_gateway/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Payment{})
	return db
}

func TestProcessPayment(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	services.InitializeService(db)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/process_payment", ProcessPayment)

	requestBody, _ := json.Marshal(models.ProcessPaymentRequest{
		CardNumber:  "4242424242424242",
		ExpiryMonth: "12",
		ExpiryYear:  "2024",
		CVV:         "123",
		Amount:      100.00,
		Currency:    "USD",
	})

	req, _ := http.NewRequest("POST", "/process_payment", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "approved", response["status"])
}

func TestRetrievePayment(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	services.InitializeService(db)

	// Add a sample payment to the database
	payment := models.Payment{
		CardNumber:  "4242",
		ExpiryMonth: "12",
		ExpiryYear:  "2024",
		Amount:      100.00,
		Currency:    "USD",
		Status:      "approved",
	}
	db.Create(&payment)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/retrieve_payment/:id", RetrievePayment)

	req, _ := http.NewRequest("GET", "/retrieve_payment/1", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, float64(100), response["amount"].(float64))
	assert.Equal(t, "**** **** **** 4242", response["card_number"].(string))
	assert.Equal(t, "USD", response["currency"].(string))
	assert.Equal(t, "12", response["expiry_month"].(string))
	assert.Equal(t, "2024", response["expiry_year"].(string))
	assert.Equal(t, "approved", response["status"].(string))
}
