package establishment

import "errors"

type Service interface {
	GetAll() ([]Establishment, error)
	GetByID(id int) (*Establishment, error)
	Create(e Establishment) error
	Update(e Establishment) error
	Delete(id int) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() ([]Establishment, error) {
	return s.repo.FindAll()
}

func (s *service) GetByID(id int) (*Establishment, error) {
	return s.repo.FindByID(id)
}

func (s *service) Create(e Establishment) error {
	return s.repo.Create(e)
}

func (s *service) Update(e Establishment) error {
	return s.repo.Update(e)
}

func (s *service) Delete(id int) error {
	hasStores, err := s.repo.HasStores(id)
	if err != nil {
		return err
	}
	if hasStores {
		return errors.New("Não é possível remover: o estabelecimento possui lojas vinculadas")
	}
	return s.repo.Delete(id)
}
