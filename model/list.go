package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/fandreuz/yabm/model/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func quoteAndJoin(values []string) string {
	mapped := make([]string, len(values))
	for idx, v := range values {
		mapped[idx] = fmt.Sprintf("'%s'", v)
	}
	return strings.Join(mapped, ",")
}

func ListBookmarks(tagNames []string) ([]entity.Bookmark, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

	if len(tagNames) == 0 {
		selectQuery := fmt.Sprintf("select * from %s", bookmarksTable)
		return listEntities[entity.Bookmark](selectQuery, conn)
	}

	selectTagWhereRhs := quoteAndJoin(tagNames)
	selectQuery := fmt.Sprintf(`
select * from %s where id in (
	select bookmarkId from %s where tagId in (
		select distinct id from %s where label in (%s)
	) group by bookmarkId having count(*) = %d
)`, bookmarksTable, assignedTagsTable, tagsTable, selectTagWhereRhs, len(tagNames))
	return listEntities[entity.Bookmark](selectQuery, conn)
}

func ListTags() ([]entity.Tag, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

	return listEntities[entity.Tag]("select * from tags", conn)
}

func listEntities[E any](sqlQuery string, session queryableSession) ([]E, error) {
	rows, queryErr := session.Query(context.TODO(), sqlQuery)
	if queryErr != nil {
		if pgErr, ok := queryErr.(*pgconn.PgError); ok {
			return nil, handleDatabaseError(pgErr)
		}
		return nil, queryErr
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[E])
}
