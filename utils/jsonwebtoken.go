package utils

import (
	"crypto/rsa"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	jwtPublicKey        *rsa.PublicKey
	jwtPrivateKey       *rsa.PrivateKey
	jwtAlgorithm        = jwt.SigningMethodRS256
	jwtAccessTokenType  = "access"
	jwtRefreshTokenType = "refresh"
)

func init() {
	publicKeyContent, err := os.ReadFile(Config.Jwt.PublicKey)
	if err != nil {
		log.Fatal("Failed to open RSA public key file: ", err)
	}

	privateKeyContent, err := os.ReadFile(Config.Jwt.PrivateKey)
	if err != nil {
		log.Fatal("Failed to open RSA public key file: ", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyContent)
	if err != nil {
		log.Fatal("Failed to parse RSA public key: ", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyContent)
	if err != nil {
		log.Fatal("Failed to parse RSA private key: ", err)
	}

	// move a local variable to a global variables
	jwtPublicKey = publicKey
	jwtPrivateKey = privateKey
}

// GenerateAccessToken returns a signed access token.
func GenerateAccessToken(userId string) (string, error) {
	return generateToken(userId, jwtAccessTokenType, time.Minute*time.Duration(Config.Jwt.AccessTokenExpires))
}

// GenerateRefreshToken returns a signed refresh token.
func GenerateRefreshToken(userId string) (string, error) {
	return generateToken(userId, jwtRefreshTokenType, time.Minute*time.Duration(Config.Jwt.AccessTokenExpires))
}

// ValidateAccessToken validates the access token.
func ValidateAccessToken(token string) (*string, error) {
	return validateToken(token, jwtAccessTokenType)
}

// ValidateRefreshToken validates the refresh token.
func ValidateRefreshToken(token string) (*string, error) {
	return validateToken(token, jwtRefreshTokenType)
}

func generateToken(userId string, tokenType string, expireTime time.Duration) (string, error) {
	// create token claims
	token := jwt.NewWithClaims(jwtAlgorithm, jwt.MapClaims{
		"iss": Config.Jwt.Issuer,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(expireTime).Unix(),
		"sub": userId,
		"typ": tokenType,
	})

	// sign token
	tokenString, err := token.SignedString(jwtPrivateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateToken(token string, tokenType string) (*string, error) {
	claims := jwt.MapClaims{}

	// parse token and get claims
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		// if token isn't an rsa token
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, ErrInvalidSigningMethod
		}

		return jwtPublicKey, nil
	})
	if err != nil {
		return nil, ErrInvalidToken
	}

	// get user id from the token
	userId, exists := claims["sub"].(string)
	if !exists {
		return nil, ErrInvalidTokenSubject
	}

	// check if the token type is corrected
	tokenTypeClaim, exists := claims["typ"].(string)
	if !exists && tokenTypeClaim != tokenType {
		return nil, ErrInvalidTokenType
	}

	return &userId, nil
}
