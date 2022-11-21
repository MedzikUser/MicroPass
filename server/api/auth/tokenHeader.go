package auth

import (
	"fmt"
	"strings"

	"github.com/MedzikUser/MicroPass/database"
	"github.com/MedzikUser/MicroPass/server/errors"
	"github.com/MedzikUser/MicroPass/utils"
	"github.com/gin-gonic/gin"
)

type Token struct {
	Token string
	User  database.User
}

func GetAccessToken(ctx *gin.Context) *Token {
	// get token from header
	token, err := getToken(ctx)
	if err != nil {
		errors.ErrInvalidAuthorizationHeader(ctx)
		return nil
	}

	// validate the token
	userId, err := utils.ValidateAccessToken(*token)
	if err != nil {
		errors.ErrInvalidToken(ctx)
		return nil
	}

	// find user in the database
	user := database.User{
		Id: *userId,
	}
	err = user.Take()
	if err != nil {
		errors.ErrDatabase(ctx)
		return nil
	}

	// return token and user
	return &Token{
		Token: *token,
		User:  user,
	}
}

func getToken(ctx *gin.Context) (*string, error) {
	// get `Authorization` header
	authHeader := ctx.GetHeader("Authorization")

	// split bearer token
	parts := strings.Split(authHeader, "Bearer")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid authorization header")
	}

	// get token
	token := strings.TrimSpace(parts[1])
	if len(token) == 0 {
		return nil, fmt.Errorf("missing token")
	}

	return &token, nil
}
