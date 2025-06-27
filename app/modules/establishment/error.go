package establishment

// Package error defines custom error messages for the establishment module.
// These errors are used to provide meaningful feedback when operations fail, such as when an establishment is not found or when input validation fails.

import "errors"

var (
	ErrNotFound		 = errors.New("estabelecimento não encontrado")
	ErrInvalidInput  = errors.New("dados inválidos")
	ErrForeignKeyExists = errors.New("não é possível remover: o estabelecimento possui lojas vinculadas")
	ErrCreateFailed  = errors.New("falha ao criar estabelecimento")
	ErrUpdateFailed  = errors.New("falha ao atualizar estabelecimento")
	ErrDeleteFailed  = errors.New("falha ao deletar estabelecimento")
	ErrInvalidID     = errors.New("ID inválido")
	ErrListEstablishmentsFailed = errors.New("falha ao listar estabelecimentos")
)