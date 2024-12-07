package model

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type queryableSession interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

// TODO
const connectionUrl = "postgres://admin:pwd@localhost:5432/admin"

func openConnection() (*pgx.Conn, error) {
	return pgx.Connect(context.TODO(), connectionUrl)
}

func handleDatabaseError(pgErr *pgconn.PgError) error {
	panic(fmt.Errorf("DB error occurred, code: %s, message: '%s', details: '%s'", pgErr.Code, pgErr.Message, pgErr.Detail))
}
