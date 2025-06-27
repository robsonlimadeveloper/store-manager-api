package dto

// Package dto provides data transfer objects for the store module.
// These DTOs are used to transfer data between the client and server, ensuring that the data structure is consistent and validated.

type CreateStoreDTO struct {
	Number          string `json:"number" validate:"required"`
	Name            string `json:"name" validate:"required"`
	CorporateName   string `json:"corporate_name"`
	Address         string `json:"address"`
	City            string `json:"city"`
	State           string `json:"state"`
	ZipCode         string `json:"zip_code"`
	StreetNumber    string `json:"street_number"`
	EstablishmentID int    `json:"establishment_id" validate:"required"`
}

type UpdateStoreDTO struct {
	Name          string `json:"name"`
	Address       string `json:"address"`
	StreetNumber  string `json:"street_number"`
}

type StoreResponseDTO struct {
	ID              int    `json:"id"`
	Number          string `json:"number"`
	Name            string `json:"name"`
	CorporateName   string `json:"corporate_name"`
	Address         string `json:"address"`
	City            string `json:"city"`
	State           string `json:"state"`
	ZipCode         string `json:"zip_code"`
	StreetNumber    string `json:"street_number"`
	EstablishmentID int    `json:"establishment_id"`
}
type StoreListResponseDTO struct {
	Stores []StoreResponseDTO `json:"stores"`
}

type StoreIDParam struct {
	ID int `param:"id" validate:"required"`
}