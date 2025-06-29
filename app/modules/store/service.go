package store

// Package store provides the service layer for managing stores in the store manager API.

import (
	"store-manager-api/app/core"
)

type StoreService interface {
	core.Service[Store]
	GetByEstablishmentID(estID int) ([]Store, error)
}

type service struct {
	repo StoreRepository
}

func NewService(r StoreRepository) StoreService {
	return &service{repo: r}
}

func (s *service) GetAll() ([]Store, error) {
	return s.repo.FindAll()
}

func (s *service) GetByID(id int) (*Store, error) {
	est, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if est == nil {
		return nil, ErrNotFound
	}
	return est, nil
}

func (s *service) Create(store *Store) error {
	return s.repo.Create(store)
}

func (s *service) Update(store *Store) error {
	return s.repo.Update(store)
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *service) GetByEstablishmentID(estID int) ([]Store, error) {
	return s.repo.FindByEstablishmentID(estID)
}
