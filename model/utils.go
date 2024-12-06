package model

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
)

func handleDatabaseError(pgErr *pgconn.PgError) error {
	return fmt.Errorf("DB error occurred, code: %s, message: '%s', details: '%s'", pgErr.Code, pgErr.Message, pgErr.Detail)
}