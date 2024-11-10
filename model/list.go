package model

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func ListBookmarks() ([]Bookmark, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

    sqlQuery := "select id, url, title, creationDate from bookmarks"

	rows, queryErr := conn.Query(context.TODO(), sqlQuery)
	if queryErr != nil {
		return nil, queryErr
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[Bookmark])
}

func ListTags() ([]Tag, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

    sqlQuery := "select id, label, creationDate from tags"

	rows, queryErr := conn.Query(context.TODO(), sqlQuery)
	if queryErr != nil {
		return nil, queryErr
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[Tag])
}
