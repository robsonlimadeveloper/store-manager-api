package establishment

// Establishment represents the data structure for an establishment.
// It includes fields for the establishment's ID, number, name, corporate name, address, city, state, zip code, and street number.
// This struct is used to model the data in the application and is typically used in conjunction with a database.


type Establishment struct {
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
