package entity

import (
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
	return EntityToString(b)
}

type BookmarkCreationRequest struct {
	Title string
	Url   string
}
