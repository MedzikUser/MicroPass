package utils

import "errors"

var (
	ErrInvalidSigningMethod = errors.New("invalid signing method")
	ErrInvalidToken         = errors.New("invalid token")
	ErrInvalidTokenSubject  = errors.New("invalid token subject")
	ErrInvalidTokenType     = errors.New("invalid token type")
)
