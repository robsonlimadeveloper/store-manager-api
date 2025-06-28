package auth

// Package auth provides error definitions for authentication-related operations.
// These errors can be used throughout the authentication module to provide consistent error handling.

import "errors"

var (
	ErrInvalidCredentials = errors.New("Invalid credentials")
	ErrTokenExpired       = errors.New("Token expired")
	ErrUnauthorized       = errors.New("Not authorized")
	ErrInternalServer     = errors.New("Internal server error")
	ErrInvalidInput       = errors.New("Invalid input data")
)