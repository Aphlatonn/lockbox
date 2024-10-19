package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes the given password using bcrypt
func HashMasterPassword(password string) (string, error) {
	// Generate a random salt and hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword checks if the provided password matches the hashed password
func VerifyMasterPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
