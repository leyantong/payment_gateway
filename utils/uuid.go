package utils

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func GeneratePaymentUUID(cardNumber string, amount float64) uuid.UUID {
	timestamp := time.Now().UnixNano()
	data := fmt.Sprintf("%s:%f:%d", cardNumber, amount, timestamp)
	hash := sha1.New()
	hash.Write([]byte(data))
	hashedData := hash.Sum(nil)
	return uuid.NewSHA1(uuid.NameSpaceOID, hashedData)
}
