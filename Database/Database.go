package Database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:123456@localhost:5432/GameBackend")
	if err != nil {
		return nil, err
	}
	return conn, nil
}
