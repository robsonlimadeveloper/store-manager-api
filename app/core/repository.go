package core

// Repository defines a generic interface for CRUD operations.
// It can be used for any type T, allowing for flexible data handling.

type Repository[T any] interface {
	FindAll() ([]T, error)
	FindByID(id int) (*T, error)
	Create(t *T) error
	Update(t *T) error
	Delete(id int) error
}