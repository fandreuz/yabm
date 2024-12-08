package model

import (
	"context"

	"github.com/fandreuz/yabm/model/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func GetBookmarkById(id uint64) (entity.Bookmark, error) {
	return getById[entity.Bookmark]("select * from bookmarks where id=$1", id)
}

func GetTagById(id uint64) (entity.Tag, error) {
	return getById[entity.Tag]("select * from tags where id=$1", id)
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
		if pgErr, ok := queryErr.(*pgconn.PgError); ok {
			return errorEntity, handleDatabaseError(pgErr)
		}
		return errorEntity, queryErr
	}
	defer rows.Close()

	return pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[E])
}
