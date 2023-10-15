package blog

import "github.com/luyasr/simple-blog/pkg/e"

type Status int

const (
	StatusDraft Status = iota
	StatusPublished
)

type AuditStatus int

const (
	AuditStatusEditorInvited AuditStatus = iota
	AuditStatusUnderReview
	AuditStatusReject
	AuditStatusAccept
)

var (
	NotFound         = e.NewNotFound("blog not found")
	UpdateFailed     = e.NewUpdateFailed("blog update failed, affected 0")
	PermissionDenied = e.NewAccessDenied("Permission denied")
)
