package crypto

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/bytepass/server/config"
	"github.com/bytepass/server/log"
	"github.com/golang-jwt/jwt/v4"
)

var (
	publicKey    *rsa.PublicKey
	privateKey   *rsa.PrivateKey
	jwtAlgorithm = jwt.SigningMethodRS256
)

func init() {
	// parse RSA Public Key
	public, err := jwt.ParseRSAPublicKeyFromPEM([]byte(config.Config.Jwt.PublicKey))
	if err != nil {
		log.Fatal("Failed to parse RSA public key: %v", err)
	}

	// parse RSA Private Key
	private, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(config.Config.Jwt.PrivateKey))
	if err != nil {
		log.Fatal("Failed to parse RSA private key: %v", err)
	}

	// move a local variable to a global variables
	publicKey = public
	privateKey = private
}

// Generate and sign a JWT token.
func GenerateJWT(userId string) (string, error) {
	// create token
	token := jwt.NewWithClaims(jwtAlgorithm, jwt.MapClaims{
		"user": userId,
		"exp":  time.Now().Add(time.Hour * config.Config.Jwt.Expires).Unix(),
	})

	// sign token
	tokenString, err := token.SignedString(privateKey)

	return tokenString, err
}

// Validate JWT token.
func ValidateJWT(token string) (*string, error) {
	var claims jwt.MapClaims

	// parse token and get claims
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	// get user UUID from token
	userId := claims["user"].(string)

	return &userId, nil
}
