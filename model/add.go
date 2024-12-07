package model

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type databaseErrorHandler func(dbError *pgconn.PgError) error

func execQueryAndReturn[T any](sqlInsertQuery string, session queryableSession, handler databaseErrorHandler) (*T, error) {
	rows, dbInsertErr := session.Query(context.TODO(), sqlInsertQuery)
	if dbInsertErr != nil {
		if pgErr, ok := dbInsertErr.(*pgconn.PgError); ok {
			return nil, handler(pgErr)
		}
		return nil, dbInsertErr
	}
	defer rows.Close()

	mappedRows, rowsCollectionError := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[T])
	if rowsCollectionError != nil {
		if pgErr, ok := rowsCollectionError.(*pgconn.PgError); ok {
			return nil, handler(pgErr)
		}
		return nil, fmt.Errorf("Error occurred during row collection to struct of type '%T': '%s'", *new(T), rowsCollectionError.Error())
	}
	return &mappedRows, nil
}

func CreateBookmark(request BookmarkCreationRequest) (*Bookmark, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

	sqlInsertQuery := fmt.Sprintf("insert into bookmarks (url, title, creationDate) values ('%s', '%s', now()) returning *", request.Url, request.Title)
	return execQueryAndReturn[Bookmark](sqlInsertQuery, conn, handleDatabaseError)
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

	sqlSelectQuery := fmt.Sprintf("select * from tags where label = '%s'", request.Label)
	tag, selectErr := execQueryAndReturn[Tag](sqlSelectQuery, tx, handleDatabaseError)
	if tag != nil || selectErr != nil {
		return tag, selectErr
	}

	sqlInsertQuery := fmt.Sprintf("insert into tags (label, creationDate) values ('%s', now()) returning *", request.Label)
	tag, insertErr := execQueryAndReturn[Tag](sqlInsertQuery, tx, handleDatabaseError)
	if insertErr != nil {
		return nil, insertErr
	}

	tx.Commit(context.TODO())
	return tag, nil
}

func AssignTag(request TagAssignationRequest) (*AssignedTag, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

	sqlInsertQuery := fmt.Sprintf("insert into assigned_tags (tagId, bookmarkId) values (%d, %d) returning *", request.TagId, request.BookmarkId)
	handler := func(dbError *pgconn.PgError) error {
		if dbError.Code == "23505" {
			return nil
		}
		return handleDatabaseError(dbError)
	}
	return execQueryAndReturn[AssignedTag](sqlInsertQuery, conn, handler)
}
