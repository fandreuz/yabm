package model

import "context"


func GetBookmarkById(id uint64) (*Bookmark, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

    sqlQuery := "select url, title, creationDate from bookmarks where id=$1"

    var url string
	var title *string
	var creationDate uint64
	queryErr := conn.QueryRow(context.TODO(), sqlQuery, id).Scan(&url, &title, &creationDate)
	if queryErr != nil {
		return nil, queryErr
	}

	return &Bookmark{Url: url, Title: title, CreationDate: creationDate, Id: &id}, nil
}

func GetTagById(id uint64) (*Tag, error) {
	conn, connError := openConnection()
	if connError != nil {
		return nil, connError
	}
	defer conn.Close(context.TODO())

    sqlQuery := "select label, creationDate from tags where id=$1"

    var label string
	var creationDate uint64
	queryErr := conn.QueryRow(context.TODO(), sqlQuery, id).Scan(&label, &creationDate)
	if queryErr != nil {
		return nil, queryErr
	}

	return &Tag{Label: label, CreationDate: creationDate, Id: &id}, nil
}
