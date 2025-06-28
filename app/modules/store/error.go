package store

// Package store provides the service layer for managing stores in the store manager API.

import "errors"

var (
	ErrNotFound         = errors.New("Store not found")
	ErrInvalidInput     = errors.New("Invalid input data")
	ErrCreateFailed   = errors.New("Failed to create store")
	ErrUpdateFailed   = errors.New("Failed to update store")
	ErrDeleteFailed   = errors.New("Failed to delete store")
	ErrInvalidID       = errors.New("Invalid ID")
	ErrForeignKeyExists = errors.New("Cannot delete store with existing foreign key references")
	ErrListStoresFailed = errors.New("Failed to list stores")
	ErrInternalServer   = errors.New("Internal server error")
)