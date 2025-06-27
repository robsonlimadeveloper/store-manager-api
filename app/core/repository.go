package core

type Repository[T any] interface {
	FindAll() ([]T, error)
	FindByID(id int) (*T, error)
	Create(t *T) error
	Update(t *T) error
	Delete(id int) error
}