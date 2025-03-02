package main

import (
	"net/http"

	"github.com/amandavmanduca/fullcycle-golang-2-challenge/clients"
	"github.com/amandavmanduca/fullcycle-golang-2-challenge/internal/handlers"
	"github.com/amandavmanduca/fullcycle-golang-2-challenge/internal/services"
)

func main() {

	clients := clients.NewClientsContainer()
	serv := services.NewServicesContainer(&clients)
	handler := handlers.NewHandlerContainers(serv)

	http.HandleFunc("/address", handler.CepHandler.GetAddress)
	http.ListenAndServe(":8080", nil)

}
