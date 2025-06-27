package establishment

import (
	"database/sql"
	"store-manager-api/app/core"
)

type Repository interface {
	core.Repository[Establishment]
	HasStores(establishmentID int) (bool, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]Establishment, error) {
	rows, err := r.db.Query(`SELECT id, number, name, corporate_name, address, city, state, zip_code, street_number FROM establishments`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var establishments []Establishment
	for rows.Next() {
		var e Establishment
		err := rows.Scan(&e.ID, &e.Number, &e.Name, &e.CorporateName, &e.Address, &e.City, &e.State, &e.ZipCode, &e.StreetNumber)
		if err != nil {
			return nil, err
		}
		establishments = append(establishments, e)
	}
	return establishments, nil
}

func (r *repository) FindByID(id int) (*Establishment, error) {
	row := r.db.QueryRow(`SELECT id, number, name, corporate_name, address, city, state, zip_code, street_number FROM establishments WHERE id=$1`, id)
	var e Establishment
	if err := row.Scan(&e.ID, &e.Number, &e.Name, &e.CorporateName, &e.Address, &e.City, &e.State, &e.ZipCode, &e.StreetNumber); err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *repository) Create(e *Establishment) error {
	_, err := r.db.Exec(`
		INSERT INTO establishments (number, name, corporate_name, address, city, state, zip_code, street_number)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		e.Number, e.Name, e.CorporateName, e.Address, e.City, e.State, e.ZipCode, e.StreetNumber,
	)
	return err
}

func (r *repository) Update(e *Establishment) error {
	_, err := r.db.Exec(`
		UPDATE establishments
		SET number=$1, name=$2, corporate_name=$3, address=$4, city=$5, state=$6, zip_code=$7, street_number=$8
		WHERE id=$9`,
		e.Number, e.Name, e.CorporateName, e.Address, e.City, e.State, e.ZipCode, e.StreetNumber, e.ID,
	)
	return err
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM establishments WHERE id=$1`, id)
	return err
}

func (r *repository) HasStores(establishmentID int) (bool, error) {
	var count int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM stores WHERE establishment_id = $1`, establishmentID).Scan(&count)
	return count > 0, err
}
