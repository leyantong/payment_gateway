package utils

import (
	"crypto/sha1"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

var mu sync.Mutex

func GeneratePaymentUUID(cardNumber string, amount float64) uuid.UUID {
	mu.Lock()
	defer mu.Unlock()

	timestamp := time.Now().UnixNano()
	data := fmt.Sprintf("%s:%f:%d", cardNumber, amount, timestamp)
	hash := sha1.New()
	hash.Write([]byte(data))
	hashedData := hash.Sum(nil)
	return uuid.NewSHA1(uuid.NameSpaceOID, hashedData)
}
