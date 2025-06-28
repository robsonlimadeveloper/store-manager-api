package establishment

// Package establishment provides the service layer for managing establishments in the store manager API.

import (
	"errors"
	"store-manager-api/app/core"
)

type Service interface {
	core.Service[Establishment]
}

type service struct {
	repo EstablishmentRepository
}

func NewService(r EstablishmentRepository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() ([]Establishment, error) {
	return s.repo.FindAll()
}

func (s *service) GetByID(id int) (*Establishment, error) {
	return s.repo.FindByID(id)
}

func (s *service) Create(e *Establishment) error {
	return s.repo.Create(e)
}

func (s *service) Update(e *Establishment) error {
	return s.repo.Update(e)
}

func (s *service) Delete(id int) error {
	hasStores, err := s.repo.HasStores(id)
	if err != nil {
		return err
	}
	if hasStores {
		return errors.New(ErrForeignKeyExists.Error())
	}
	return s.repo.Delete(id)
}
