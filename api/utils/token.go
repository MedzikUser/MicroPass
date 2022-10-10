package utils

import (
	"fmt"
	"strings"

	"github.com/bytepass/server/crypto"
	"github.com/bytepass/server/database"
	"github.com/gin-gonic/gin"
)

type Token struct {
	Token  string
	UserId string
	User   database.User
}

// Get token from http request.
func GetToken(c *gin.Context) (*Token, error) {
	// get `Authorization` header
	authHeader := c.GetHeader("Authorization")

	// split bearer token
	parts := strings.Split(authHeader, "Bearer")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid authorization header")
	}

	// get bearer token
	token := strings.TrimSpace(parts[1])
	if len(token) == 0 {
		return nil, fmt.Errorf("missing token")
	}

	// validate token
	userId, err := crypto.ValidateJwt(token)
	if err != nil {
		return nil, err
	}

	// take user from database
	var user database.User
	user.Id = *userId

	user, err = user.Get()
	if err != nil {
		return nil, err
	}

	// check if user have verified email
	if !user.EmailVerified {
		return nil, fmt.Errorf("user doesn't have a verified email address")
	}

	return &Token{token, *userId, user}, nil
}
