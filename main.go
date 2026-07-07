package main

import (
	"fmt"
	"net/http"

	"api-restaurante/routes"
)
func main() {

	routes.ConfigurarRotas()

	fmt.Println("Servidor rodando em http://localhost:3000")

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println(err)
	}
}
