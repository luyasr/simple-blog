package blog

import (
	"encoding/json"
	"github.com/luyasr/simple-blog/app/common"
	"github.com/luyasr/simple-blog/pkg/logger"
)

type Blog struct {
	*common.Meta
	// 文章状态
	Status Status `json:"status"`
	// 审核文章的请求
	AuditAt     int64       `json:"audit_at"`
	AuditStatus AuditStatus `json:"audit_status"`
	// 创建文章的请求
	*CreateBlogRequest
}

func (b *Blog) TableName() string {
	return "blogs"
}

func (b *Blog) String() string {
	bytes, err := json.Marshal(b)
	if err != nil {
		logger.Console.Error().Err(err).Stack().Send()
	}
	return string(bytes)
}

func NewBlog(req *CreateBlogRequest) *Blog {
	return &Blog{
		Status:            StatusDraft,
		CreateBlogRequest: req,
	}
}
