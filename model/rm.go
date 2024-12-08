package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/fandreuz/yabm/model/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func UnassignTagByLabel(request entity.TagAssignationByLabelRequest) error {
	conn, connError := openConnection()
	if connError != nil {
		return connError
	}
	defer conn.Close(context.TODO())

	tx, transactionErr := conn.Begin(context.TODO())
	if transactionErr != nil {
		return transactionErr
	}
	defer tx.Rollback(context.TODO())

	tag, findTagByLabelErr := findTagByLabel(request.TagLabel, tx)
	if findTagByLabelErr != nil {
		if errors.Is(findTagByLabelErr, pgx.ErrNoRows) {
			// Tag not found, nothing to do
			return nil
		}
		return findTagByLabelErr
	}

	if unassignErr := unassignTag(entity.TagAssignationRequest{TagId: tag.Id, BookmarkId: request.BookmarkId}, tx); unassignErr != nil {
		return unassignErr
	}

	tx.Commit(context.TODO())
	return nil
}

func UnassignTagById(request entity.TagAssignationRequest) error {
	conn, connError := openConnection()
	if connError != nil {
		return connError
	}
	defer conn.Close(context.TODO())
	return unassignTag(request, conn)
}

func deleteAssignedTags(id uint64, columnName string, session queryableSession) error {
	whereClause := fmt.Sprintf("where %s = %d", columnName, id)
	return deleteEntity[entity.AssignedTag](assignedTagsTable, whereClause, session)
}

func DeleteBookmarkById(id uint64) error {
	conn, connError := openConnection()
	if connError != nil {
		return connError
	}
	defer conn.Close(context.TODO())

	tx, transactionErr := conn.Begin(context.TODO())
	if transactionErr != nil {
		return transactionErr
	}
	defer tx.Rollback(context.TODO())

	if err := deleteAssignedTags(id, "bookmarkId", tx); err != nil {
		return err
	}

	whereClause := fmt.Sprintf("where id = %d", id)
	if err := deleteEntity[entity.Bookmark](bookmarksTable, whereClause, tx); err != nil {
		return err
	}

	tx.Commit(context.TODO())
	return nil
}

func unassignTag(request entity.TagAssignationRequest, session queryableSession) error {
	whereClause := fmt.Sprintf("where tagId = %d AND bookmarkId = %d", request.TagId, request.BookmarkId)
	return deleteEntity[entity.AssignedTag](assignedTagsTable, whereClause, session)
}

func deleteEntity[T any](table string, whereClause string, session queryableSession) error {
	sqlQuery := fmt.Sprintf("delete from %s %s", table, whereClause)
	if err := execQuery(sqlQuery, session); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			return handleDatabaseError(pgErr)
		}
		return err
	}
	return nil
}
