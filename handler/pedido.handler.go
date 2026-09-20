package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"api-restaurante/models"
	"api-restaurante/repository"
	"api-restaurante/service"
)

func ListarPedido(
	pedidoRepo *repository.PedidoRepository,
	w http.ResponseWriter,
	r *http.Request,
) {

	pedidos, err := pedidoRepo.ListarPedidosDB(r.Context())

	if err != nil {
		http.Error(
			w,
			"Erro ao listar pedidos",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(pedidos)
}

func BuscarPedidoID(
	pedidoRepo *repository.PedidoRepository,
	w http.ResponseWriter,
	r *http.Request,
) {
	idTexto := strings.TrimPrefix(r.URL.Path, "/pedidos/")

	id, err := strconv.Atoi(idTexto)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	pedido, err := pedidoRepo.BuscarPedidoPorIDDB(
		r.Context(),
		id,
	)

	if err != nil {
		if errors.Is(err, repository.ErrpedidoNaoEncontrado) {
			http.Error(w, "Pedido não encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Erro Interno do Servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pedido)
}

// Impementação de atualização de Pedidos

func AtualizarPedido(
	pedidoRepo *repository.PedidoRepository,
	w http.ResponseWriter,
	r *http.Request,
) {
	idTexto := strings.TrimPrefix(r.URL.Path, "/pedidos/")

	id, err := strconv.Atoi(idTexto)
	if err != nil {
		http.Error(w, "ID Inválido", http.StatusBadRequest)
		return
	}

	var dados struct {
		Status string `json:"status"`
	}

	err = json.NewDecoder(r.Body).Decode(&dados)
	if err != nil {
		http.Error(w, "Erro ao ler JSON", http.StatusBadRequest)
		return
	}

	if dados.Status == "" {
		http.Error(w, "Status é Obrigatório", http.StatusBadRequest)
		return
	}

	if !service.StatusValido(dados.Status) {
		http.Error(w, "Status Inválido", http.StatusBadRequest)
		return
	}

	pedido, err := service.AtualizarStatusPedido(
		r.Context(),
		pedidoRepo,
		id,
		dados.Status,
	)

	if err != nil {
		if errors.Is(err, repository.ErrpedidoNaoEncontrado) {
			http.Error(w, "Pedido Não Encontrado", http.StatusNotFound)
			return
		}

		if errors.Is(err, service.ErrtransicaoStatusInvalida) {
			http.Error(w, "Transição de Status Inválida", http.StatusBadRequest)
			return
		}

		http.Error(w, "Erro Interno do Servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(pedido)
}

func CriarPedido(
	pedidoRepo *repository.PedidoRepository,
	w http.ResponseWriter,
	r *http.Request) {

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

	pedido, err := pedidoRepo.CriarPedidoDB(
		r.Context(),
		novoPedido,
	)

	if err != nil {
		http.Error(
			w,
			"Erro ao Criar Pedido",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pedido)
}
