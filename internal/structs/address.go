package structs

type CepOrigin string

var (
	BRASIL_API CepOrigin = "Brasil API"
	VIA_CEP    CepOrigin = "Via Cep API"
)

type BrasilAddressResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func (a *BrasilAddressResponse) ToAddressResponse() *AddressResponse {
	return &AddressResponse{
		Address: Address{
			Cep:          a.Cep,
			State:        a.State,
			City:         a.City,
			Neighborhood: a.Neighborhood,
			Street:       a.Street,
		},
		Origin: BRASIL_API,
	}
}

type ViaCepAddressResponse struct {
	Cep          string `json:"cep"`
	Street       string `json:"logradouro"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	State        string `json:"uf"`
	Estado       string `json:"estado"`
}

func (a *ViaCepAddressResponse) ToAddressResponse() *AddressResponse {
	return &AddressResponse{
		Address: Address{
			Cep:          a.Cep,
			State:        a.State,
			City:         a.City,
			Neighborhood: a.Neighborhood,
			Street:       a.Street,
		},
		Origin: VIA_CEP,
	}
}

type Address struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

type AddressResponse struct {
	Address Address
	Origin  CepOrigin `json:"origin"`
}
