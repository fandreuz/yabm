package model

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// TODO
const connectionUrl = "postgres://admin:pwd@localhost:5432/admin"

func openConnection() (*pgx.Conn, error) {
	return pgx.Connect(context.TODO(), connectionUrl)
}
