package model

type Bookmark struct {
	id           uint64
	url          string
	title        string
	// Milliseconds since epoch (UTC)
	creationDate uint64
}

type Tag struct {
	id    uint64
	label string
	// Milliseconds since epoch (UTC)
	creationDate uint64
}
