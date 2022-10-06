package utils

import (
	"fmt"
	"strings"

	"github.com/bytepass/server/crypto"
	"github.com/gin-gonic/gin"
)

type Token struct {
	Token  string
	UserId string
}

func GetToken(c *gin.Context) (*Token, error) {
	authHeader := c.GetHeader("Authorization")

	parts := strings.Split(authHeader, "Bearer")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid authorization header")
	}

	token := strings.TrimSpace(parts[1])
	if len(token) == 0 {
		return nil, fmt.Errorf("missing token")
	}

	userId, err := crypto.ValidateJWT(token)
	if err != nil {
		return nil, err
	}

	return &Token{token, *userId}, nil
}
