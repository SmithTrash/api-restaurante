package repository

import "api-restaurante/models"

var Pedidos = []models.Pedido{
	{
		ID:      1,
		Cliente: "João",
		Status:  "CRIADO",
	},
}

func ListarPedidos() []models.Pedido {
	return Pedidos
}

func CriarPedido(pedido models.Pedido) models.Pedido {

	pedido.ID = len(Pedidos) + 1
	pedido.Status = "CRIADO"

	Pedidos = append(Pedidos, pedido)

	return pedido
}

func BuscarPedidoPorID(id int) (*models.Pedido, int) {

	for i := range Pedidos {
		if Pedidos[i].ID == id {
			return &Pedidos[i], i
		}
	}

	return nil, -1
}

func AtualizarStatus(id int, status string) (*models.Pedido, bool) {

	pedido, _ := BuscarPedidoPorID(id)

	if pedido == nil {
		return nil, false
	}

	pedido.Status = status

	return pedido, true
}
