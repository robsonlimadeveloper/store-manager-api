package core

// Service defines a generic interface for CRUD operations.
// It can be used for any type T, allowing for flexible data handling.

type Service[T any] interface {
	GetAll() ([]T, error)
	GetByID(id int) (*T, error)
	Create(t *T) error
	Update(t *T) error
	Delete(id int) error
}