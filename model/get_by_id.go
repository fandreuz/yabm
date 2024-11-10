package model

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func GetBookmarkById(id uint64) (Bookmark, error) {
	return getById[Bookmark]("select * from bookmarks where id=$1", id)
}

func GetTagById(id uint64) (Tag, error) {
	return getById[Tag]("select * from tags where id=$1", id)
}

func getById[E any](sqlQuery string, id uint64) (E, error) {
	var errorEntity E

	conn, connError := openConnection()
	if connError != nil {
		return errorEntity, connError
	}
	defer conn.Close(context.TODO())

	rows, queryErr := conn.Query(context.TODO(), sqlQuery, id)
	if queryErr != nil {
		return errorEntity, queryErr
	}

	return pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[E])
}
