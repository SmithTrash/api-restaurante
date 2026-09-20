package repository

import (
	"api-restaurante/models"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PedidoRepository struct {
	db *pgxpool.Pool
}

var ErrpedidoNaoEncontrado = errors.New("Pedido Não Encontrado")

func NewPedidoRepository(db *pgxpool.Pool) *PedidoRepository {
	return &PedidoRepository{db: db}
}

func (r *PedidoRepository) CriarPedidoDB(
	ctx context.Context,
	pedido models.Pedido) (models.Pedido, error) {

	query := `
		INSERT INTO pedidos (cliente, status)
		VALUES ($1, $2)
		RETURNING id, cliente, status
	`

	err := r.db.QueryRow(
		ctx,
		query,
		pedido.Cliente,
		"CRIADO",
	).Scan(
		&pedido.ID,
		&pedido.Cliente,
		&pedido.Status,
	)

	if err != nil {
		return models.Pedido{}, fmt.Errorf(
			"erro ao criar pedido no banco: %w",
			err,
		)
	}

	return pedido, nil
}

var Pedidos = []models.Pedido{
	{
		ID:      1,
		Cliente: "João",
		Status:  "CRIADO",
	},
}

func (r *PedidoRepository) ListarPedidosDB(
	ctx context.Context,
) ([]models.Pedido, error) {

	query := `
        SELECT id, cliente, status
        FROM pedidos
        ORDER BY id
    `

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf(
			"erro ao listar pedidos no banco: %w",
			err,
		)
	}
	defer rows.Close()

	pedidos := []models.Pedido{}

	for rows.Next() {
		var pedido models.Pedido

		err := rows.Scan(
			&pedido.ID,
			&pedido.Cliente,
			&pedido.Status,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"erro ao ler pedido do banco: %w",
				err,
			)
		}

		pedidos = append(pedidos, pedido)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"erro ao percorrer pedidos: %w",
			err,
		)
	}

	return pedidos, nil
}
func (r *PedidoRepository) BuscarPedidoPorIDDB(
	ctx context.Context,
	id int,
) (models.Pedido, error) {
	query := `
		SELECT id, cliente, status
		From pedidos
		WHERE id = $1
	`
	var pedido models.Pedido

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(

		&pedido.ID,
		&pedido.Cliente,
		&pedido.Status,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Pedido{}, ErrpedidoNaoEncontrado
		}
		return models.Pedido{}, err
	}

	return pedido, nil
}

func ListarPedido() []models.Pedido {

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

func (r *PedidoRepository) AtualizarStatusDB(
	ctx context.Context,
	id int,
	status string,
) (models.Pedido, error) {

	query := `
		UPDATE pedidos
		SET status = $1
		WHERE id = $2
		RETURNING id, cliente, status
		`

	var pedido models.Pedido

	err := r.db.QueryRow(
		ctx,
		query,
		status,
		id,
	).Scan(
		&pedido.ID,
		&pedido.Cliente,
		&pedido.Status,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Pedido{}, ErrpedidoNaoEncontrado
		}
		return models.Pedido{}, err
	}

	return pedido, nil
}
