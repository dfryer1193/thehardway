package models

// User is a sample user model
type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	WebAuthnID   []byte `json:"webauthn_id"`   // ID for WebAuthn (YubiKey)
	WebAuthnCred []byte `json:"webauthn_cred"` // Credential for WebAuthn (YubiKey)
}
