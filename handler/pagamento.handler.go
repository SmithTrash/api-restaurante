package handler

import (
	"api-restaurante/models"
	"api-restaurante/repository"
	"encoding/json"
	"net/http"
)

func ProcessarPagamento(w http.ResponseWriter, r *http.Request) {
	var pagamento models.Pagamento

	err := json.NewDecoder(r.Body).Decode(&pagamento)
	if err != nil {
		http.Error(w, "Erro ao ler json", http.StatusBadRequest)
		return
	}

	if pagamento.Status == "RECUSADO" {
		http.Error(w, "Pagamento recusado", http.StatusBadRequest)
		return
	}

	pedido, ok := repository.AtualizarStatus(
		pagamento.PedidoID,
		"PAGO",
	)

	if !ok {
		http.Error(w, "Pedido não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pedido)
}
