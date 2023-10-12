package blog

import "github.com/luyasr/simple-blog/pkg/e"

type Status int

const (
	StatusDraft Status = iota
	StatusPublished
)

var (
	NotFound = e.NewNotFound("blog not found")
)
