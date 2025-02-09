package entity

import (
	"time"
)

type Tag struct {
	Id    uint64
	Label string
	// Milliseconds since epoch (UTC)
	CreationDate time.Time
}

func (t Tag) String() string {
	return EntityToString(t)
}

type TagCreationRequest struct {
	Label string
}
