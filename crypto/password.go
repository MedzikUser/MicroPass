package crypto

import (
	"crypto/rand"

	"github.com/bytepass/server/config"
)

var saleSize = config.Config.Crypto.Salt
var iter = config.Config.Crypto.Iterations

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, saleSize)

	_, err := rand.Read(salt[:])

	return salt, err
}
