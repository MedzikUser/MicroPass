package crypto

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"

	"github.com/bytepass/server/config"
	"golang.org/x/crypto/pbkdf2"
)

var (
	saleSize = config.Config.Crypto.Salt
	iter     = config.Config.Crypto.Iterations
)

// Generate a pseudo-random salt.
func GenerateSalt() ([]byte, error) {
	salt := make([]byte, saleSize)

	_, err := rand.Read(salt[:])

	return salt, err
}

// Compute PBKDF2-SHA512 hash of the given password.
func HashPassword(password string, salt []byte) string {
	// convert password string to byte slice
	passwordBytes := []byte(password)

	// compute password hash using PBKDF2-SHA512 algorithm
	dk := pbkdf2.Key(passwordBytes, salt, iter, 64, sha512.New)

	// convert the hashed password to a hex string
	hex := hex.EncodeToString(dk)

	return hex
}

// Check if two passwords match
//
// Example:
//	// generate random salt
//	salt := generateSalt()
//
//	// hash password (e.g. password returned from database)
//	hashedPassword := hashPassword("Super#secret#passphrase", salt)
//
//	PasswordMatch(hashedPassword, "Super#secret#passphrase", salt)
func PasswordMatch(hashedPassword string, currentPassword string, salt []byte) bool {
	currentPassword = HashPassword(currentPassword, salt)

	return hashedPassword == currentPassword
}
