package model

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type queryableSession interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

func execQueryAndReturn[T any](sqlInsertQuery string, session queryableSession) (*T, error) {
	rows, dbInsertErr := session.Query(context.TODO(), sqlInsertQuery)
	if dbInsertErr != nil {
		if pgErr, ok := dbInsertErr.(*pgconn.PgError); ok {
			return nil, handleDatabaseError(pgErr)
		}
		return nil, dbInsertErr
	}

	mappedRows, rowsCollectionError := pgx.CollectRows(rows, pgx.RowToStructByName[T])
	if rowsCollectionError != nil {
		if pgErr, ok := rowsCollectionError.(*pgconn.PgError); ok {
			return nil, handleDatabaseError(pgErr)
		}
		return nil, rowsCollectionError
	}

	if len(mappedRows) == 0 {
		return nil, nil
	}
	return &mappedRows[0], nil
}

func CreateBookmark(request BookmarkCreationRequest) (*Bookmark, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

	sqlInsertQuery := fmt.Sprintf("insert into bookmarks (url, title, creationDate) values (%s, %s, now()) returning id, creationDate", request.Url, request.Title)
	return execQueryAndReturn[Bookmark](sqlInsertQuery, conn)
}

func GetOrCreateTag(request TagCreationRequest) (*Tag, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

	tx, transactionErr := conn.Begin(context.TODO())
	if transactionErr != nil {
		return nil, transactionErr
	}
	defer tx.Rollback(context.TODO())

	sqlSelectQuery := fmt.Sprintf("select (id, label, creationDate) from tags where label = '%s'", request.Label)
	tag, err := execQueryAndReturn[Tag](sqlSelectQuery, tx)
	if tag != nil || err != nil {
		return tag, err
	}

	sqlInsertQuery := fmt.Sprintf("insert into tags (label, creationDate) values (%s, now()) returning id, creationDate", request.Label)
	return execQueryAndReturn[Tag](sqlInsertQuery, tx)
}

func AssignTag(request TagAssignationRequest) (*AssignedTag, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

	sqlInsertQuery := fmt.Sprintf("insert into assigned_tags (tagId, bookmarkId) values (%d, %d)", request.TagId, request.BookmarkId)
	return execQueryAndReturn[AssignedTag](sqlInsertQuery, conn)
}
