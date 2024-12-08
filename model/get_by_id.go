package model

import (
	"context"
	"fmt"

	"github.com/fandreuz/yabm/model/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func GetBookmarkById(id uint64) (entity.Bookmark, error) {
	sqlQuery := fmt.Sprintf("select * from %s where id=%d", bookmarksTable, id)
	return getById[entity.Bookmark](sqlQuery)
}

func GetTagById(id uint64) (entity.Tag, error) {
	sqlQuery := fmt.Sprintf("select * from %s where id=%d", tagsTable, id)
	return getById[entity.Tag](sqlQuery)
}

func getById[E any](sqlQuery string) (E, error) {
	var errorEntity E

	conn, connError := openConnection()
	if connError != nil {
		return errorEntity, connError
	}
	defer conn.Close(context.TODO())

	rows, queryErr := conn.Query(context.TODO(), sqlQuery)
	if queryErr != nil {
		if pgErr, ok := queryErr.(*pgconn.PgError); ok {
			return errorEntity, handleDatabaseError(pgErr)
		}
		return errorEntity, queryErr
	}
	defer rows.Close()

	return pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[E])
}
