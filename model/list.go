package model

import (
	"context"
	"fmt"
	"strings"

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

func ListBookmarks(tagNames []string) ([]Bookmark, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

	if len(tagNames) == 0 {
		selectQuery := fmt.Sprintf("select * from bookmarks")
		return listEntities[Bookmark](selectQuery, conn)
	}

	selectTagWhereRhs := quoteAndJoin(tagNames)
	selectQuery := fmt.Sprintf(`
select * from bookmarks where id in (
	select bookmarkId from assigned_tags where tagId in (
		select distinct id from tags where label in (%s)
	) group by bookmarkId having count(*) = %d
)`, selectTagWhereRhs, len(tagNames))
	return listEntities[Bookmark](selectQuery, conn)
}

func ListTags() ([]Tag, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

	return listEntities[Tag]("select * from tags", conn)
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
