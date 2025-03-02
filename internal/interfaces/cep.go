package interfaces

import (
	"context"

	"github.com/amandavmanduca/fullcycle-golang-2-challenge/internal/structs"
)

type CepServiceInterface interface {
	GetAddress(ctx context.Context, cep string) (*structs.AddressResponse, error)
}
