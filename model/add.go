package model

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/fandreuz/yabm/model/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func findTagByLabel(label string, session queryableSession) (*entity.Tag, error) {
	sqlSelectQuery := fmt.Sprintf("select * from %s where label = '%s'", tagsTable, label)
	tag, err := execQueryAndReturn[entity.Tag](sqlSelectQuery, session, handleDatabaseError)
	if tag != nil || errors.Is(err, pgx.ErrNoRows) {
		return tag, nil
	}
	return nil, err
}

func getOrCreateTag(request entity.TagCreationRequest, tx pgx.Tx) (*entity.Tag, error) {
	tag, findTagByLabelErr := findTagByLabel(request.Label, tx)
	if tag != nil || findTagByLabelErr != nil {
		return tag, findTagByLabelErr
	}

	sqlInsertQuery := fmt.Sprintf("insert into %s (label, creationDate) values ('%s', now()) returning *", tagsTable, request.Label)
	return execQueryAndReturn[entity.Tag](sqlInsertQuery, tx, handleDatabaseError)
}

func assignTagById(request entity.TagAssignationRequest, session queryableSession) (*entity.AssignedTag, error) {
	sqlInsertQuery := fmt.Sprintf("insert into %s (tagId, bookmarkId) values (%d, %d) returning *", assignedTagsTable, request.TagId, request.BookmarkId)
	handler := func(dbError *pgconn.PgError) error {
		if dbError.Code == "23505" {
			return nil
		}
		return handleDatabaseError(dbError)
	}
	return execQueryAndReturn[entity.AssignedTag](sqlInsertQuery, session, handler)
}

func CreateBookmark(request entity.BookmarkCreationRequest) (*entity.Bookmark, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

	sanitizedTitle := strings.Replace(request.Title, "'", "''", -1)
	sqlInsertQuery := fmt.Sprintf("insert into %s (url, title, creationDate) values ('%s', '%s', now()) returning *", bookmarksTable, request.Url, sanitizedTitle)
	return execQueryAndReturn[entity.Bookmark](sqlInsertQuery, conn, handleDatabaseError)
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
