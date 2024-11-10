package model

import "fmt"

type Tag struct {
	Id    *uint64
	Label string
	// Milliseconds since epoch (UTC)
	CreationDate uint64	
}

func NewTag(label string, creationDate uint64) Tag {
	return Tag{Label: label, CreationDate: creationDate}
}

func (t Tag) String() string {
	if t.Id != nil {
		return fmt.Sprintf("{label: '%s', creationDate: '%d', id: '%d'}", t.Label, t.CreationDate, *(t.Id))
	}
	return fmt.Sprintf("{label: '%s', creationDate: '%d'}", t.Label, t.CreationDate)
}

func (t Tag) WithId(id uint64) Tag {
	return Tag{Id: &id, Label: t.Label, CreationDate: t.CreationDate}
}
