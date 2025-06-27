package establishment

type CreateEstablishmentDTO struct {
	Number        string `json:"number" validate:"required"`
	Name          string `json:"name" validate:"required"`
	CorporateName string `json:"corporate_name"`
	Address       string `json:"address"`
	City          string `json:"city"`
	State         string `json:"state"`
	ZipCode       string `json:"zip_code"`
	StreetNumber  string `json:"street_number"`
}

type UpdateEstablishmentDTO struct {
	Name         string `json:"name"`
	Address      string `json:"address"`
	StreetNumber string `json:"street_number"`
}

type EstablishmentResponseDTO struct {
	ID            int    `json:"id"`
	Number        string `json:"number"`
	Name          string `json:"name"`
	CorporateName string `json:"corporate_name"`
	Address       string `json:"address"`
	City          string `json:"city"`
	State         string `json:"state"`
	ZipCode       string `json:"zip_code"`
	StreetNumber  string `json:"street_number"`
}
