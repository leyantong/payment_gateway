package utils

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// GeneratePaymentUUID generates a UUID based on the masked card number, date, and amount.
func GeneratePaymentUUID(cardNumber string, amount float64) uuid.UUID {
	currentDate := time.Now().Format("2006-01-02") // Get current date in YYYY-MM-DD format
	data := fmt.Sprintf("%s:%f:%s", cardNumber[len(cardNumber)-4:], amount, currentDate)
	hash := sha1.New()
	hash.Write([]byte(data))
	hashedData := hash.Sum(nil)
	return uuid.NewSHA1(uuid.NameSpaceOID, hashedData)
}
