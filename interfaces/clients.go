package interfaces

import (
	"context"
	"net/http"

	"github.com/amandavmanduca/fullcycle-golang-2-challenge/internal/structs"
)

type HttpClientInterface interface {
	Get(ctx context.Context, path string) (*http.Response, error)
}

type ClientsContainer struct {
	BrasilApi BrasilApiInterface
	ViaCepApi ViaCepApiInterface
}

type BrasilApiInterface interface {
	GetAddress(ctx context.Context, cep string) (*structs.AddressResponse, error)
}

type ViaCepApiInterface interface {
	GetAddress(ctx context.Context, cep string) (*structs.AddressResponse, error)
}
