package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func UnassignTagByLabel(request TagAssignationByLabelRequest) error {
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

	unassignErr := unassignTag(TagAssignationRequest{TagId: tag.Id, BookmarkId: request.BookmarkId}, tx)
	if unassignErr == nil {
		tx.Commit(context.TODO())
	}
	return unassignErr
}

func UnassignTagById(request TagAssignationRequest) error {
	conn, connError := openConnection()
	if connError != nil {
		return connError
	}
	defer conn.Close(context.TODO())
	return unassignTag(request, conn)
}

func unassignTag(request TagAssignationRequest, session queryableSession) error {
	sqlInsertQuery := fmt.Sprintf("delete from assigned_tags where tagId = %d AND bookmarkId = %d returning *", request.TagId, request.BookmarkId)
	handler := func(dbError *pgconn.PgError) error {
		return handleDatabaseError(dbError)
	}
	_, err := execQueryAndReturn[AssignedTag](sqlInsertQuery, session, handler)
	return err
}
