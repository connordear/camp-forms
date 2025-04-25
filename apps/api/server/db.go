package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"os"
)

func connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {

		return nil, err
	}

	return conn, nil
}
