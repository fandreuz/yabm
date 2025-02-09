package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/fandreuz/yabm/model/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Errors of type ErrNoRows are not forwarded
func findTagByLabel(label string, session queryableSession) (*entity.Tag, error) {
	query := fmt.Sprintf("select * from %s where label = '@label'", tagsTable)
	tag, err := execQueryAndReturn[entity.Tag](session, handleDatabaseError, query, pgx.NamedArgs{"label": label})
	if tag != nil {
		return tag, nil
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return nil, err
}

func getOrCreateTag(request entity.TagCreationRequest, tx pgx.Tx) (*entity.Tag, error) {
	tag, findTagByLabelErr := findTagByLabel(request.Label, tx)
	if tag != nil || findTagByLabelErr != nil {
		return tag, findTagByLabelErr
	}

	query := fmt.Sprintf("insert into %s (label, creationDate) values ('@label', now()) returning *", tagsTable)
	return execQueryAndReturn[entity.Tag](tx, handleDatabaseError, query, pgx.NamedArgs{"label": request.Label})
}

func assignTagById(request entity.TagAssignationRequest, session queryableSession) (*entity.AssignedTag, error) {
	query := fmt.Sprintf("insert into %s (tagId, bookmarkId) values (@tagId, @bookmarkId) returning *", tagsTable)
	handler := func(dbError *pgconn.PgError) error {
		if dbError.Code == "23505" {
			return nil
		}
		return handleDatabaseError(dbError)
	}
	return execQueryAndReturn[entity.AssignedTag](session, handler, query, pgx.NamedArgs{"tagId": request.TagId, "bookmarkId": request.BookmarkId})
}

func CreateBookmark(request entity.BookmarkCreationRequest) (*entity.Bookmark, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

	query := fmt.Sprintf("insert into %s (url, title, creationDate) values ('@url', '@title', now()) returning *", bookmarksTable)
	return execQueryAndReturn[entity.Bookmark](conn, handleDatabaseError, query, pgx.NamedArgs{"url": request.Url, "title": request.Title})
}

func GetOrCreateTag(request entity.TagCreationRequest) (*entity.Tag, error) {
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

	tag, getOrCreateErr := getOrCreateTag(request, tx)
	if getOrCreateErr == nil {
		tx.Commit(context.TODO())
	}
	return tag, getOrCreateErr
}

func AssignTagById(request entity.TagAssignationRequest) (*entity.AssignedTag, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

	return assignTagById(request, conn)
}

func AssignTagByLabel(request entity.TagAssignationByLabelRequest) (*entity.AssignedTag, error) {
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

	tag, getOrCreateErr := getOrCreateTag(entity.TagCreationRequest{Label: request.TagLabel}, tx)
	if getOrCreateErr != nil {
		return nil, getOrCreateErr
	}

	assignedTag, assignErr := assignTagById(entity.TagAssignationRequest{TagId: tag.Id, BookmarkId: request.BookmarkId}, tx)
	if assignErr == nil {
		tx.Commit(context.TODO())
	}
	return assignedTag, assignErr
}
