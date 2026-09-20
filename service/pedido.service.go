package service

import (
	"api-restaurante/models"
	"api-restaurante/repository"
	"context"
	"errors"
)

var ErrtransicaoStatusInvalida = errors.New("Transição de Status Inválida")

func StatusValido(status string) bool {
	switch status {

	case "CRIADO":
		return true

	case "PAGO":
		return true

	case "EM_PREPARO":
		return true

	case "PRONTO":
		return true

	case "ENTREGUE":
		return true

	default:
		return false

	}
}
func TransicaoStatusValida(statusAtual string, novoStatus string) bool {
	switch statusAtual {

	case "CRIADO":
		return novoStatus == "PAGO"

	case "PAGO":
		return novoStatus == "EM_PREPARO"

	case "EM_PREPARO":
		return novoStatus == "PRONTO"

	case "PRONTO":
		return novoStatus == "ENTREGUE"

	default:
		return false
	}
}

func AtualizarStatusPedido(
	ctx context.Context,
	pedidoRepo *repository.PedidoRepository,
	id int,
	novoStatus string,
) (models.Pedido, error) {

	pedidoAtual, err := pedidoRepo.BuscarPedidoPorIDDB(
		ctx,
		id,
	)
	if err != nil {
		return models.Pedido{}, err
	}

	if !TransicaoStatusValida(pedidoAtual.Status, novoStatus) {

		return models.Pedido{}, ErrtransicaoStatusInvalida
	}

	pedidoAtualizado, err := pedidoRepo.AtualizarStatusDB(
		ctx,
		id,
		novoStatus,
	)

	if err != nil {
		return models.Pedido{}, err
	}

	return pedidoAtualizado, nil
}
