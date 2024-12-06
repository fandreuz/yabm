package model

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
)

type AssignedTag struct {
	TagId      uint64
	BookmarkId uint64
}

func (t AssignedTag) String() string {
	return fmt.Sprintf("{tagId: '%d', bookmarkId: '%d'}", t.TagId, t.BookmarkId)
}

type TagAssignationRequest struct {
	TagId      uint64
	BookmarkId uint64
}

func AssignTag(request TagAssignationRequest) (*AssignedTag, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

	sqlInsertQuery := "insert into assigned_tags (tagId, bookmarkId) values ($1, $2)"

	_, dbInsertErr := conn.Query(context.TODO(), sqlInsertQuery, request.TagId, request.BookmarkId)
	if dbInsertErr != nil {
		if pgErr, ok := dbInsertErr.(*pgconn.PgError); ok {
			return nil, handleDatabaseError(pgErr)
		}
		return nil, dbInsertErr
	}

	return &AssignedTag{TagId: request.TagId, BookmarkId: request.BookmarkId}, nil
}
