package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/dfryer1193/thehardway/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
)

// TODO: Figure out what I want to use for this.
// In-memory store for active challenges (could be in Redis or a database in a production setup)
var challengeStore = map[string]string{}

// LoginChallenge handles the first step of login by generating a nonce
func LoginChallenge(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Generate the nonce (challenge)
	nonce, err := utils.GenerateNonce()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate nonce"})
		return
	}

	// Store nonce with the email (for simplicity; in production use Redis, etc.)
	challengeStore[req.Email] = nonce

	// Return the nonce to the client
	c.JSON(http.StatusOK, gin.H{"nonce": nonce})
}

// Login handles the second step of login by verifying the password hash and nonce
func Login(c *gin.Context) {
	var req struct {
		Email        string `json:"email"`
		PasswordHash string `json:"password_hash"`
		Nonce        string `json:"nonce"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Get stored default email and password hash (or from a DB in production)
	defaultEmail := os.Getenv("DEFAULT_EMAIL")
	defaultPasswordHash := os.Getenv("DEFAULT_PASSWORD_HASH")

	// Verify email
	if req.Email != defaultEmail {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	// Retrieve the nonce from the challengeStore (or Redis in production)
	storedNonce, exists := challengeStore[req.Email]
	if !exists || storedNonce != req.Nonce {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired nonce"})
		return
	}

	// Generate the expected hash (stored password hash + nonce)
	expectedHash := sha256.New()
	expectedHash.Write([]byte(defaultPasswordHash + req.Nonce))
	expectedPasswordHash := hex.EncodeToString(expectedHash.Sum(nil))

	// Compare the expected hash with the received password hash
	if expectedPasswordHash != req.PasswordHash {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// At this point, the password is verified; proceed to YubiKey 2FA
	// (This is the next stage in the flow)
	c.JSON(http.StatusOK, gin.H{"message": "Password verified, proceed to YubiKey 2FA"})
}

// YubiKeyVerification handles the YubiKey 2FA verification
func YubiKeyVerification(c *gin.Context) {
	var req struct {
		Email      string `json:"email"`
		YubiKeyOTP string `json:"yubikey_otp"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate the YubiKey OTP (for example, using the Yubico Web API)
	// In production, this would involve an external call to validate the YubiKey OTP
	isValid, err := utils.ValidateYubiKeyOTP(req.YubiKeyOTP)
	if err != nil || !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid YubiKey OTP"})
		return
	}

	// If OTP is valid, create and return a JWT for session management
	token, err := utils.GenerateJWT(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// ChangePassword allows the site owner to change their password after re-authentication
func ChangePassword(c *gin.Context) {
	var req struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Retrieve current password hash from environment (or DB)
	currentPasswordHash := os.Getenv("DEFAULT_PASSWORD_HASH")

	// Verify current password
	if err := bcrypt.CompareHashAndPassword([]byte(currentPasswordHash), []byte(req.CurrentPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid current password"})
		return
	}

	// Hash the new password
	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash new password"})
		return
	}

	// Update the password hash in environment (or save to DB)
	os.Setenv("DEFAULT_PASSWORD_HASH", string(newPasswordHash))

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}
