package handler

import (
	"api-restaurante/models"
	"api-restaurante/repository"
	"encoding/json"
	"errors"
	"net/http"
)

func ProcessarPagamento(
	pedidoRepo *repository.PedidoRepository,
	w http.ResponseWriter,
	r *http.Request,
) {
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

	pedido, err := pedidoRepo.AtualizarStatusDB(
		r.Context(),
		pagamento.PedidoID,
		"PAGO",
	)

	if err != nil {
		if errors.Is(err, repository.ErrpedidoNaoEncontrado) {
			http.Error(w, "Pedido não encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Erro ao atualizar status do pedido", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pedido)
}
