// middleware_test.go
package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestValidatePaymentRequest(t *testing.T) {
	router := gin.Default()
	router.POST("/test", ValidatePaymentRequest(), func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	jsonPayload := `{
		"card_number": "4242424242424242",
		"expiry_month": "12",
		"expiry_year": "2024",
		"cvv": "123",
		"amount": 100.00,
		"currency": "USD"
	}`

	req, _ := http.NewRequest("POST", "/test", bytes.NewBufferString(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
