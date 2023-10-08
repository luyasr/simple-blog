package blog

type Status int

const (
	StatusUnKnown Status = iota
	StatusDraft
	StatusPublished
)
