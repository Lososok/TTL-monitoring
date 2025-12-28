package services

import "errors"

var (
	ErrInvalidURL    = errors.New("invalid url")
	ErrInvalidArgs   = errors.New("invalid args")
	ErrTLS           = errors.New("tls handshake failed")
	ErrAlreadyExists = errors.New("certificate already exists")
	ErrNotFound      = errors.New("certificate not found")
)
