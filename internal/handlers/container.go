package handlers

import "github.com/amandavmanduca/fullcycle-golang-2-challenge/internal/services"

type HandlerContainer struct {
	CepHandler cepHandler
}

func NewHandlerContainers(services *services.ServicesContainer) *HandlerContainer {
	return &HandlerContainer{
		CepHandler: NewCepHandler(services.CepService),
	}
}
