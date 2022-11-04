package database

import (
	"errors"
)

var (
	ErrPasswordEmpty    = errors.New("password is empty")
	ErrPasswordMismatch = errors.New("password mismatch")
)
