package services

import (
	externalInterfaces "github.com/amandavmanduca/fullcycle-golang-2-challenge/interfaces"
	"github.com/amandavmanduca/fullcycle-golang-2-challenge/internal/interfaces"
)

type ServicesContainer struct {
	CepService interfaces.CepServiceInterface
}

func NewServicesContainer(clients *externalInterfaces.ClientsContainer) *ServicesContainer {
	return &ServicesContainer{
		CepService: NewCepService(clients),
	}
}
