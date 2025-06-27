package core

import "errors"

var (
	ErrNotFound         = errors.New("registro não encontrado")
	ErrInvalidInput     = errors.New("dados inválidos")
	ErrForeignKeyExists = errors.New("registro referenciado por outra entidade")
)