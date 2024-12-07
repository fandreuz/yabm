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

type databaseErrorHandler func(dbError *pgconn.PgError) error

// TODO
const connectionUrl = "postgres://admin:pwd@localhost:5432/admin"

func openConnection() (*pgx.Conn, error) {
	return pgx.Connect(context.TODO(), connectionUrl)
}

func handleDatabaseError(pgErr *pgconn.PgError) error {
	panic(fmt.Errorf("DB error occurred, code: %s, message: '%s', details: '%s'", pgErr.Code, pgErr.Message, pgErr.Detail))
}

func execQueryAndReturn[T any](sqlInsertQuery string, session queryableSession, handler databaseErrorHandler) (*T, error) {
	rows, dbInsertErr := session.Query(context.TODO(), sqlInsertQuery)
	if dbInsertErr != nil {
		if pgErr, ok := dbInsertErr.(*pgconn.PgError); ok {
			return nil, handler(pgErr)
		}
		return nil, dbInsertErr
	}
	defer rows.Close()

	mappedRows, rowsCollectionError := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[T])
	if rowsCollectionError != nil {
		if pgErr, ok := rowsCollectionError.(*pgconn.PgError); ok {
			return nil, handler(pgErr)
		}
		return nil, rowsCollectionError
	}
	return &mappedRows, nil
}
