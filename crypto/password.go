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

// GenerateSalt returns a pseudo-random salt.
func GenerateSalt() ([]byte, error) {
	salt := make([]byte, saleSize)

	_, err := rand.Read(salt[:])

	return salt, err
}

// HashPassword returns a PBKDF2-SHA512 hash of the given password.
func HashPassword(password string, salt []byte) string {
	// convert password string to byte slice
	passwordBytes := []byte(password)

	// compute password hash using PBKDF2-SHA512 algorithm
	dk := pbkdf2.Key(passwordBytes, salt, iter, 64, sha512.New)

	// convert the hashed password to a hex string
	return hex.EncodeToString(dk)
}

// PasswordMatch validates the two passwords.
func PasswordMatch(hashedPassword string, unhashedPassword string, salt []byte) bool {
	// compute hash of the unhashed password
	hashedPasswordTwo := HashPassword(unhashedPassword, salt)

	// compare the two hashes
	return hashedPassword == hashedPasswordTwo
}
