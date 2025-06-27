package core

// Error definitions for common application errors
// These errors can be used throughout the application to provide consistent error handling.

import "errors"

var (
	ErrInternalServer = errors.New("erro interno do servidor")
	ErrUnauthorized   = errors.New("não autorizado")
	ErrTokenExpired   = errors.New("token expirado")
	ErrInvalidToken   = errors.New("token inválido")
)