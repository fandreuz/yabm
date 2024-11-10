package model

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func ListBookmarks() ([]Bookmark, error) {
	return listEntities[Bookmark]("select * from bookmarks")
}

func ListTags() ([]Tag, error) {
	return listEntities[Tag]("select * from tags")
}

func listEntities[E any](sqlQuery string) ([]E, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

	rows, queryErr := conn.Query(context.TODO(), sqlQuery)
	if queryErr != nil {
		return nil, queryErr
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[E])
}
