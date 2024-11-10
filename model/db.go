package model

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// TODO
const connectionUrl = "postgres://admin:pwd@localhost:5432/admin"

func openConnection() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), connectionUrl)
}

func GetBookmarkById(id uint64) (*Bookmark, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}

    sqlQuery := "select url, title, creationDate from bookmarks where id=$1"

    var url, title string
	var creationDate uint64
	queryErr := conn.QueryRow(context.Background(), sqlQuery, id).Scan(&url, &title, &creationDate)
	if queryErr != nil {
		return nil, queryErr
	}

	return &Bookmark{url: url, title: title, creationDate: creationDate, id: id}, nil
}

func GetTagById(id uint64) (*Tag, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}

    sqlQuery := "select label, creationDate from bookmarks where id=$1"

    var label string
	var creationDate uint64
	queryErr := conn.QueryRow(context.Background(), sqlQuery, id).Scan(&label, &creationDate)
	if queryErr != nil {
		return nil, queryErr
	}

	return &Tag{label: label, creationDate: creationDate, id: id}, nil
}
