package establishment

// Package error defines custom error messages for the establishment module.
// These errors are used to provide meaningful feedback when operations fail, such as when an establishment is not found or when input validation fails.

import "errors"

var (
	ErrNotFound		 = errors.New("establishment not found")
	ErrInvalidInput  = errors.New("invalid input data")
	ErrForeignKeyExists = errors.New("cannot delete establishment with associated stores")
	ErrCreateFailed  = errors.New("failed to create establishment")
	ErrUpdateFailed  = errors.New("failed to update establishment")
	ErrDeleteFailed  = errors.New("failed to delete establishment")
	ErrInvalidID     = errors.New("invalid ID")
	ErrListEstablishmentsFailed = errors.New("failed to list establishments")
)