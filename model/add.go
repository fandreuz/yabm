package model

import (
	"context"
)


func AddBookmark(b Bookmark) (*Bookmark, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

	transaction, transactionErr := conn.Begin(context.TODO())
	if transactionErr != nil {
		return nil, transactionErr
	}

    sqlInsertQuery := "insert into bookmarks (url, creationDate) values ($1, $2)"
	_, dbInsertErr := transaction.Exec(context.TODO(), sqlInsertQuery, b.Url, b.CreationDate)
	if dbInsertErr != nil {
		return nil, dbInsertErr
	}

	sqlIdQuery := "select currval(pg_get_serial_sequence('bookmarks', 'id'))"
	var id uint64
	dbSelectErr := transaction.QueryRow(context.TODO(), sqlIdQuery).Scan(&id)
	if dbSelectErr != nil {
		return nil, dbSelectErr
	}

	commitErr := transaction.Commit(context.TODO())
	if commitErr != nil {
		return nil, commitErr
	}

	var newBookmark = b.WithId(id)
	return &newBookmark, nil
}

func AddTag(t Tag) (*Tag, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

	transaction, transactionErr := conn.Begin(context.TODO())
	if transactionErr != nil {
		return nil, transactionErr
	}

    sqlInsertQuery := "insert into tags (label, creationDate) values ($1, $2)"
	_, dbInsertErr := transaction.Exec(context.TODO(), sqlInsertQuery, t.Label, t.CreationDate)
	if dbInsertErr != nil {
		return nil, dbInsertErr
	}

	sqlIdQuery := "select currval(pg_get_serial_sequence('tags', 'id'))"
	var id uint64
	dbSelectErr := transaction.QueryRow(context.TODO(), sqlIdQuery).Scan(&id)
	if dbSelectErr != nil {
		return nil, dbSelectErr
	}

	commitErr := transaction.Commit(context.TODO())
	if commitErr != nil {
		return nil, commitErr
	}

	var newTag = t.WithId(id)
	return &newTag, nil
}
