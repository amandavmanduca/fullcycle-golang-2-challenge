package services

import (
	"context"
	"time"

	externalInterfaces "github.com/amandavmanduca/fullcycle-golang-2-challenge/interfaces"
	"github.com/amandavmanduca/fullcycle-golang-2-challenge/internal/interfaces"
	"github.com/amandavmanduca/fullcycle-golang-2-challenge/internal/structs"
)

type cepService struct {
	clients *externalInterfaces.ClientsContainer
}

func NewCepService(clients *externalInterfaces.ClientsContainer) interfaces.CepServiceInterface {
	return &cepService{
		clients: clients,
	}
}

func (s *cepService) GetAddress(ctx context.Context, cep string) (*structs.AddressResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	brasilApi := make(chan structs.AddressResponse, 1)
	go s.getAddressFromBrasilApi(ctx, cep, brasilApi)

	viaCepApi := make(chan structs.AddressResponse, 1)
	go s.getAddressFromViaCepApi(ctx, cep, viaCepApi)

	select {
	case res := <-brasilApi:
		return &res, nil
	case res := <-viaCepApi:
		return &res, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (s *cepService) getAddressFromBrasilApi(ctx context.Context, cep string, c chan<- structs.AddressResponse) {
	res, err := s.clients.BrasilApi.GetAddress(ctx, cep)
	if err == nil {
		select {
		case c <- *res:
		case <-ctx.Done():
		}
	}
}

func (s *cepService) getAddressFromViaCepApi(ctx context.Context, cep string, c chan<- structs.AddressResponse) {
	res, err := s.clients.ViaCepApi.GetAddress(ctx, cep)
	if err == nil {
		select {
		case c <- *res:
		case <-ctx.Done():
		}
	}
}
