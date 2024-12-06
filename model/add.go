package model

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
)

func AddBookmark(request BookmarkCreationRequest) (*Bookmark, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

	sqlInsertQuery := "insert into bookmarks (url, title, creationDate) values ($1, $2, now()) returning id, creationDate"

	var id uint64
	var creationDate time.Time
	dbInsertErr := conn.QueryRow(context.TODO(), sqlInsertQuery, request.Url, request.Title).Scan(&id, &creationDate)
	if dbInsertErr != nil {
		if pgErr, ok := dbInsertErr.(*pgconn.PgError); ok {
			return nil, handleDatabaseError(pgErr)
		}
		return nil, dbInsertErr
	}

	return &Bookmark{Url: request.Url, Title: request.Title, Id: id, CreationDate: creationDate}, nil
}

func AddTag(request TagCreationRequest) (*Tag, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

	sqlInsertQuery := "insert into tags (label, creationDate) values ($1, now()) returning id, creationDate"

	var id uint64
	var creationDate time.Time
	dbInsertErr := conn.QueryRow(context.TODO(), sqlInsertQuery, request.Label).Scan(&id, &creationDate)
	if dbInsertErr != nil {
		if pgErr, ok := dbInsertErr.(*pgconn.PgError); ok {
			return nil, handleDatabaseError(pgErr)
		}
		return nil, dbInsertErr
	}

	return &Tag{Label: request.Label, CreationDate: creationDate, Id: id}, nil
}

func handleDatabaseError(pgErr *pgconn.PgError) error {
	return fmt.Errorf("DB error occurred, code: %s, message: '%s', details: '%s'", pgErr.Code, pgErr.Message, pgErr.Detail)
}