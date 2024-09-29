package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateNonce creates a random nonce for the challenge-response flow
func GenerateNonce() (string, error) {
	nonce := make([]byte, 16)
	_, err := rand.Read(nonce)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(nonce), nil
}
