package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// ValidateYubiKeyOTP validates the provided YubiKey OTP using the Yubico Web API
func ValidateYubiKeyOTP(otp string) (bool, error) {
	clientID := os.Getenv("YUBICO_CLIENT_ID")
	secretKey := os.Getenv("YUBICO_SECRET_KEY")

	if clientID == "" || secretKey == "" {
		return false, fmt.Errorf("Yubico API credentials not configured")
	}

	nonce, err := GenerateNonce()
	if err != nil {
		log.Err(err).Msg("Error generating nonce")
	}

	// Prepare the API request
	apiURL := "https://api.yubico.com/wsapi/2.0/verify"
	query := url.Values{}
	query.Set("id", clientID)
	query.Set("otp", otp)
	query.Set("nonce", nonce)

	// Sign the request with HMAC-SHA1 using the secret key
	signature := signYubiKeyRequest(query, secretKey)
	query.Set("h", signature)

	// Make the request to the Yubico API
	response, err := http.Get(apiURL + "?" + query.Encode())
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	// Check if the API response contains "status=OK"
	if strings.Contains(string(body), "status=OK") {
		return true, nil
	}

	return false, fmt.Errorf("YubiKey OTP validation failed")
}

// signYubiKeyRequest signs the request parameters using HMAC-SHA1 with the secret key
func signYubiKeyRequest(params url.Values, secret string) string {
	// Create the string to sign
	data := params.Encode()

	// Decode the secret key (it's base64 encoded)
	secretKey, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return ""
	}

	// Create HMAC-SHA1 signature
	h := hmac.New(sha1.New, secretKey)
	h.Write([]byte(data))

	// Encode the signature to base64
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature
}
