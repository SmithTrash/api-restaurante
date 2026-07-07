package handler

import (
	"encoding/json"
	"net/http"

	"api-restaurante/models"
	"api-restaurante/repository"
)

func ListarPedidos(w http.ResponseWriter, r *http.Request) {

	pedidos := repository.ListarPedidos()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(pedidos)
}

func CriarPedido(w http.ResponseWriter, r *http.Request) {

	var novoPedido models.Pedido

	err := json.NewDecoder(r.Body).Decode(&novoPedido)

	if err != nil {
		http.Error(w, "Erro ao ler JSON", http.StatusBadRequest)
		return
	}

	if novoPedido.Cliente == "" {
		http.Error(w, "Cliente é obrigatório", http.StatusBadRequest)
		return
	}

	pedido := repository.CriarPedido(novoPedido)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pedido)
}
