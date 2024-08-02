package tests

import (
	"payment_gateway/models"
	"payment_gateway/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessPayment(t *testing.T) {
	request := models.ProcessPaymentRequest{
		CardNumber:  "4242424242424242",
		ExpiryMonth: "12",
		ExpiryYear:  "2024",
		CVV:         "123",
		Amount:      100.00,
		Currency:    "USD",
	}

	response, err := services.ProcessPayment(request)
	assert.NoError(t, err)
	assert.Equal(t, "approved", response.Status)
}

func TestRetrievePayment(t *testing.T) {
	payment, err := services.RetrievePayment("1")
	assert.NoError(t, err)
	assert.NotNil(t, payment)
	assert.Equal(t, "**** **** **** 4242", payment.CardNumber)
}
