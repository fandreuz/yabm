package entity

type AssignedTag struct {
	TagId      uint64
	BookmarkId uint64
}

func (t AssignedTag) String() string {
	return EntityToString(t)
}

type TagAssignationRequest struct {
	TagId      uint64
	BookmarkId uint64
}

type TagAssignationByLabelRequest struct {
	TagLabel   string
	BookmarkId uint64
}
