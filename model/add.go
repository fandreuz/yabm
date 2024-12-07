package model

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
)

func CreateBookmark(request BookmarkCreationRequest) (*Bookmark, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

	sqlInsertQuery := fmt.Sprintf("insert into bookmarks (url, title, creationDate) values ('%s', '%s', now()) returning *", request.Url, request.Title)
	return execQueryAndReturn[Bookmark](sqlInsertQuery, conn, handleDatabaseError)
}

func findTagByLabel(label string, session queryableSession) (*Tag, error) {
	sqlSelectQuery := fmt.Sprintf("select * from tags where label = '%s'", label)
	return execQueryAndReturn[Tag](sqlSelectQuery, session, handleDatabaseError)

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

	tag, findTagByLabelErr := findTagByLabel(request.Label, tx)
	if tag != nil || findTagByLabelErr != nil {
		return tag, findTagByLabelErr
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
