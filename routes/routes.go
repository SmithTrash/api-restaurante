package routes

import (
	"fmt"
	"net/http"

	"api-restaurante/handler"
)

func ConfigurarRotas() {

	//Página Inicial
	http.HandleFunc("/", handler.Home)

	//Pagamentos
	http.HandleFunc("/pagamentos", handler.ProcessarPagamento)

	// Pedidos
	http.HandleFunc("/pedidos", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Método recebido:", r.Method)

		switch r.Method {

		case http.MethodGet:
			handler.ListarPedidos(w, r)

		case http.MethodPost:
			handler.CriarPedido(w, r)

		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})
}
