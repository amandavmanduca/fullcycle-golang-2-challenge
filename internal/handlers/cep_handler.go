package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/amandavmanduca/fullcycle-golang-2-challenge/internal/interfaces"
)

type cepHandler struct {
	cepService interfaces.CepServiceInterface
}

func NewCepHandler(cepService interfaces.CepServiceInterface) cepHandler {
	return cepHandler{
		cepService: cepService,
	}
}

func (h *cepHandler) GetAddress(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "CEP não informado", http.StatusBadRequest)
		return
	}
	address, err := h.cepService.GetAddress(ctx, "96085000")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	response := fmt.Sprintf("%s/%s - Origem: %s", address.Address.Street, address.Address.State, address.Origin)
	fmt.Println(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
