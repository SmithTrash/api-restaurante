package main

import (
	"context"
	"fmt"
	"net/http"

	"api-restaurante/database"
	"api-restaurante/repository"
	"api-restaurante/routes"
)

func main() {
	ctx := context.Background()

	db, err := database.Connect(ctx)
	if err != nil {
		fmt.Println("Erro ao iniciar banco de dados:", err)
		return
	}
	defer db.Close()

	pedidoRepo := repository.NewPedidoRepository(db)
	routes.ConfigurarRotas(pedidoRepo)

	fmt.Println("Servidor rodando em http://localhost:3000")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}
}
