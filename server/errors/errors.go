package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	// ErrInvalidRequest is returned when the request is invalid or malformed.
	ErrInvalidRequest = NewError(http.StatusBadRequest, "invalid_request", "The request is invalid or malformed.")
	// ErrGenerateToken is returned when the access or refresh token cannot be generated.
	ErrGenerateToken = NewError(http.StatusInternalServerError, "generate_token", "The access or refresh token cannot be generated.")
	// ErrInvalidToken is returned when the access or refresh token is invalid.
	ErrInvalidToken = NewError(http.StatusBadRequest, "invalid_token", "The access or refresh token is invalid.")
	// ErrDatabase is returned when the database error occured.
	ErrDatabase = NewError(http.StatusInternalServerError, "database", "The database error was occured.")
	// ErrUserNotFound is returned when the user doesn't found in the database.
	ErrUserNotFound = NewError(http.StatusBadRequest, "user_not_found", "The user doesn't found.")
	// ErrPasswordMismatch is returned when the password doesn't match.
	ErrPasswordMismatch = NewError(http.StatusBadRequest, "password_mismatch", "The password doesn't match.")
	// ErrPageNotFound is returned when the page doesn't found.
	ErrPageNotFound = NewError(http.StatusNotFound, "page_not_found", "The page doesn't found.")
	// ErrInvalidAuthorizationHeader is returned when the authorization header is invalid.
	ErrInvalidAuthorizationHeader = NewError(http.StatusBadRequest, "invalid_authorization_header", "The authorization header is invalid.")
	// ErrInvalidGrantType is returned when the grant type is invalid.
	ErrInvalidGrantType = NewError(http.StatusBadRequest, "invalid_grant_type", "The grant type is invalid.")
)

func NewError(code int, short string, long string) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(code, gin.H{"error": short, "error_description": long})
	}
}
