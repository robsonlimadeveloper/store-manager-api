package core

// Error definitions for common application errors
// These errors can be used throughout the application to provide consistent error handling.

import "errors"

var (
	ErrInternalServer = errors.New("Internal server error")
	ErrUnauthorized   = errors.New("Unauthorized access")
	ErrTokenExpired   = errors.New("Token expired")
	ErrInvalidToken   = errors.New("Invalid token")
)