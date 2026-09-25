package routes

import (
	"fmt"
	"net/http"

	"api-restaurante/handler"
	"api-restaurante/repository"
)

func ConfigurarRotas(pedidoRepo *repository.PedidoRepository) {

	//Página Inicial
	http.HandleFunc("/", handler.Home)

	//Pagamentos
	http.HandleFunc("/pagamentos", func(w http.ResponseWriter, r *http.Request) {
		handler.ProcessarPagamento(pedidoRepo, w, r)
	})

	// Pedidos
	http.HandleFunc("/pedidos", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Método recebido:", r.Method)

		switch r.Method {

		case http.MethodGet:
			handler.ListarPedido(pedidoRepo, w, r)

		case http.MethodPost:
			handler.CriarPedido(pedidoRepo, w, r)

		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/pedidos/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			handler.BuscarPedidoID(pedidoRepo, w, r)

		case http.MethodPatch:
			handler.AtualizarPedido(pedidoRepo, w, r)

		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})
}
