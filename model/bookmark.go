package model

import (
	"fmt"
	"time"
)

type Bookmark struct {
	Id    uint64
	Url   string
	Title string
	// Milliseconds since epoch (UTC)
	CreationDate time.Time
}

func (b Bookmark) String() string {
	return fmt.Sprintf("{id: '%d', url: '%s', title: '%v', creationDate: '%s'}", b.Id, b.Url, b.Title, b.CreationDate)
}

type BookmarkCreationRequest struct {
	Title string
	Url   string
}
