package model

import (
	"context"
	"fmt"

	"github.com/fandreuz/yabm/model/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func GetBookmarkById(id uint64) (entity.Bookmark, error) {
	sqlQuery := fmt.Sprintf("select * from %s where id=@id", bookmarksTable)
	return getById[entity.Bookmark](sqlQuery, pgx.NamedArgs{"id": id})
}

func GetTagById(id uint64) (entity.Tag, error) {
	sqlQuery := fmt.Sprintf("select * from %s where id=@id", tagsTable)
	return getById[entity.Tag](sqlQuery, pgx.NamedArgs{"id": id})
}

func getById[E any](sqlQuery string, namedArgs pgx.NamedArgs) (E, error) {
	var errorEntity E

	conn, connError := openConnection()
	if connError != nil {
		return errorEntity, connError
	}
	defer conn.Close(context.TODO())

	rows, queryErr := conn.Query(context.TODO(), sqlQuery, namedArgs)
	if queryErr != nil {
		if pgErr, ok := queryErr.(*pgconn.PgError); ok {
			return errorEntity, handleDatabaseError(pgErr)
		}
		return errorEntity, queryErr
	}
	defer rows.Close()

	return pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[E])
}
