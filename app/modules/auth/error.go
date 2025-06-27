package auth

// Package auth provides error definitions for authentication-related operations.
// These errors can be used throughout the authentication module to provide consistent error handling.

import "errors"

var (
	ErrInvalidCredentials = errors.New("credenciais inválidas")
	ErrTokenExpired       = errors.New("erro ao gerar o token")
	ErrUnauthorized       = errors.New("não autorizado")
	ErrInternalServer     = errors.New("erro interno do servidor")
	ErrInvalidInput       = errors.New("dados inválidos")
)