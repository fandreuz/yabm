package entity

import (
	"fmt"
	"time"
)

type Tag struct {
	Id    uint64
	Label string
	// Milliseconds since epoch (UTC)
	CreationDate time.Time
}

func (t Tag) String() string {
	return fmt.Sprintf("{id: '%d', label: '%s', creationDate: '%s'}", t.Id, t.Label, t.CreationDate)
}

type TagCreationRequest struct {
	Label string
}
