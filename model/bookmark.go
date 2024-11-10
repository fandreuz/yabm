package model

import (
	"fmt"
)

type Bookmark struct {
	Id           *uint64
	Url          string
	Title        *string
	// Milliseconds since epoch (UTC)
	CreationDate uint64
}

func NewBookmark(url string, creationDate uint64) Bookmark {
	return Bookmark{Url: url, CreationDate: creationDate}
}

func (b Bookmark) String() string {
	if b.Id != nil {
		return fmt.Sprintf("{url: '%s', title: '%v', creationDate: '%d', id: '%d'}", b.Url, b.Title, b.CreationDate, *(b.Id))
	}
	return fmt.Sprintf("{url: '%s', title: '%v', creationDate: '%d'}", b.Url, b.Title, b.CreationDate)
}

func (b Bookmark) WithId(id uint64) Bookmark {
	return Bookmark{Id: &id, Url: b.Url, Title: b.Title, CreationDate: b.CreationDate}
}
