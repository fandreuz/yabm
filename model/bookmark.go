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
	content := fmt.Sprintf("url: '%s', title: '%v', creationDate: '%d'", b.Url, b.Title, b.CreationDate)
	if b.Id != nil {
		return fmt.Sprintf("{%s, id: '%d'}", content, *(b.Id))
	}
	return fmt.Sprintf("{%s}", content)
}

func (b Bookmark) WithId(id uint64) Bookmark {
	return Bookmark{Id: &id, Url: b.Url, Title: b.Title, CreationDate: b.CreationDate}
}
