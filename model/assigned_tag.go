package model

import (
	"fmt"
)

type AssignedTag struct {
	TagId      uint64
	BookmarkId uint64
}

func (t AssignedTag) String() string {
	return fmt.Sprintf("{tagId: '%d', bookmarkId: '%d'}", t.TagId, t.BookmarkId)
}

type TagAssignationRequest struct {
	TagId      uint64
	BookmarkId uint64
}
