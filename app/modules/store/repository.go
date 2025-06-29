package store

// Package store provides the repository interface and implementation for managing stores in the store manager API.

import (
	"database/sql"
	"store-manager-api/app/core"
)



type StoreRepository interface {
	core.Repository[Store]
	FindByEstablishmentID(establishmentID int) ([]Store, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) StoreRepository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]Store, error) {
	rows, err := r.db.Query(`
		SELECT id, number, name, corporate_name, address, city, state, zip_code, street_number, establishment_id
		FROM stores
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var stores []Store
	for rows.Next() {
		var s Store
		if err := rows.Scan(&s.ID, &s.Number, &s.Name, &s.CorporateName, &s.Address, &s.City, &s.State, &s.ZipCode, &s.StreetNumber, &s.EstablishmentID); err != nil {
			return nil, err
		}
		stores = append(stores, s)
	}
	return stores, nil
}

func (r *repository) FindByID(id int) (*Store, error) {
	row := r.db.QueryRow(`
		SELECT id, number, name, corporate_name, address, city, state, zip_code, street_number, establishment_id
		FROM stores WHERE id = $1
	`, id)

	var s Store
	if err := row.Scan(&s.ID, &s.Number, &s.Name, &s.CorporateName, &s.Address, &s.City, &s.State, &s.ZipCode, &s.StreetNumber, &s.EstablishmentID); err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *repository) Create(s *Store) error {
	_, err := r.db.Exec(`
		INSERT INTO stores (number, name, corporate_name, address, city, state, zip_code, street_number, establishment_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`, s.Number, s.Name, s.CorporateName, s.Address, s.City, s.State, s.ZipCode, s.StreetNumber, s.EstablishmentID)
	return err
}

func (r *repository) Update(s *Store) error {
	_, err := r.db.Exec(`
		UPDATE stores SET
			number = $1,
			name = $2,
			corporate_name = $3,
			address = $4,
			city = $5,
			state = $6,
			zip_code = $7,
			street_number = $8,
			establishment_id = $9
		WHERE id = $10
	`, s.Number, s.Name, s.CorporateName, s.Address, s.City, s.State, s.ZipCode, s.StreetNumber, s.EstablishmentID, s.ID)
	return err
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM stores WHERE id = $1`, id)
	return err
}

func (r *repository) FindByEstablishmentID(establishmentID int) ([]Store, error) {
	rows, err := r.db.Query(`
		SELECT id, number, name, corporate_name, address, city, state, zip_code, street_number, establishment_id
		FROM stores WHERE establishment_id = $1`, establishmentID)
	if err != nil {
		return []Store{}, err
	}
	defer rows.Close()

	stores := make([]Store, 0)
	for rows.Next() {
		var s Store
		if err := rows.Scan(
			&s.ID, &s.Number, &s.Name, &s.CorporateName,
			&s.Address, &s.City, &s.State, &s.ZipCode,
			&s.StreetNumber, &s.EstablishmentID,
		); err != nil {
			return []Store{}, err
		}
		stores = append(stores, s)
	}

	return stores, nil
}
