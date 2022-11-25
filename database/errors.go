package database

import "errors"

var (
	ErrPasswordEmpty       = errors.New("password is empty")
	ErrPasswordMismatch    = errors.New("password mismatch")
	ErrUpdateCipherEmptyID = errors.New("trying to update cipher without id (all ciphers would be updated with the same data)")
)
