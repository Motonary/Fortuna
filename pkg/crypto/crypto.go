package crypto

import (
	"encoding/base64"
	
	"github.com/google/uuid"
)

func SecureRandom() string {
	return uuid.New().String()
}

func SecureRandomBase64() string {
	return base64.StdEncoding.EncodeToString(uuid.New().NodeID())
}

func LongSecureRandomBase64() string {
	return SecureRandomBase64() + SecureRandomBase64()
}

func MultipleSecureRandomBase64(n int) string {
	if n <= 1 {
			return SecureRandomBase64()
	}
	return SecureRandomBase64() + MultipleSecureRandomBase64(n - 1)
}
