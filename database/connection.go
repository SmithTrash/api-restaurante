package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		return nil, fmt.Errorf("variável DATABASE_URL não configurada")
	}

	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar pool de conexões: %w", err)
	}

	// pgxpool.New cria o pool, mas não garante imediatamente
	// que o banco está acessível. Ping valida a conexão.
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("erro ao conectar ao PostgreSQL: %w", err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

	return pool, nil
}
