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

// GenerateJwt returns a signed access token.
func GenerateJwt(userId string) (string, error) {
	return generateJwt(userId, "access", time.Minute * config.Config.Jwt.ExpiresAccessToken)
}

// GenerateActivationJwt returns a signed account activation token.
func GenerateActivationJwt(userId string) (string, error) {
	return generateJwt(userId, "activation", time.Hour * config.Config.Jwt.ExpiresActivationToken)
}

// GenerateRefreshJwt returns a signed refresh token.
func GenerateRefreshJwt(userId string) (string, error) {
	return generateJwt(userId, "refresh", time.Minute * config.Config.Jwt.ExpiresRefreshToken)
}

// ValidateJwt validates the access token.
func ValidateJwt(token string) (*string, error) {
	return validateJwt(token, "access")
}

// ValidateActivationJwt validates the account activation token.
func ValidateActivationJwt(token string) (*string, error) {
	return validateJwt(token, "activation")
}

// ValidateRefreshJwt validates the refresh token.
func ValidateRefreshJwt(token string) (*string, error) {
	return validateJwt(token, "refresh")
}

func generateJwt(userId string, tokenType string, expireTime time.Duration) (string, error) {
	// create token
	token := jwt.NewWithClaims(jwtAlgorithm, jwt.MapClaims{
		"iss": config.Config.Jwt.Issuer,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(expireTime).Unix(),
		"sub": userId,
		"typ": tokenType,
	})

	// sign token
	tokenString, err := token.SignedString(privateKey)

	return tokenString, err
}

func validateJwt(token string, tokenType string) (*string, error) {
	claims := jwt.MapClaims{}

	// parse token and get claims
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		// if token isn't an rsa token
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return publicKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}

	// get user UUID from token
	userId, exists := claims["sub"].(string)
	if !exists {
		return nil, fmt.Errorf("token doesn't contain a subject")
	}

	// get user UUID from token
	tokenTypeClaim, exists := claims["typ"].(string)
	if !exists && tokenTypeClaim != tokenType {
		return nil, fmt.Errorf("token isn't an '%s' token", tokenType)
	}

	return &userId, nil
}
