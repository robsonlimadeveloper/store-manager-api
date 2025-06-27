package store

// Package store provides the service layer for managing stores in the store manager API.

import "errors"

var (
	ErrNotFound         = errors.New("loja não encontrada")
	ErrInvalidInput     = errors.New("dados inválidos")
	ErrCreateFailed   = errors.New("falha ao criar loja")
	ErrUpdateFailed   = errors.New("falha ao atualizar loja")
	ErrDeleteFailed   = errors.New("falha ao deletar loja")
	ErrInvalidID       = errors.New("ID inválido")
	ErrForeignKeyExists = errors.New("não é possível remover: a loja possui estabelecimentos vinculados")
	ErrListStoresFailed = errors.New("falha ao listar lojas")
	ErrInternalServer   = errors.New("erro interno do servidor")
)