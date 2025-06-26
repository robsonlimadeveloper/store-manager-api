package establishment

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
