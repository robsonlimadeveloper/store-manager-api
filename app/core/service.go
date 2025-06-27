package core

type Service[T any] interface {
	GetAll() ([]T, error)
	GetByID(id int) (*T, error)
	Create(t *T) error
	Update(t *T) error
	Delete(id int) error
}