package clients

import (
	"github.com/amandavmanduca/fullcycle-golang-2-challenge/clients/http_clients"
	"github.com/amandavmanduca/fullcycle-golang-2-challenge/interfaces"
)

func NewClientsContainer() interfaces.ClientsContainer {
	brasilApi := NewHttpClient("https://brasilapi.com.br/api/cep/v1", nil)
	viaCepApi := NewHttpClient("http://viacep.com.br/ws", nil)

	return interfaces.ClientsContainer{
		BrasilApi: http_clients.NewBrasilApiClient(brasilApi),
		ViaCepApi: http_clients.NewViaCepApiClient(viaCepApi),
	}

}
