package blog

import (
	"encoding/json"
	"github.com/luyasr/simple-blog/pkg/logger"
)

type Blog struct {
	// 文章主键id
	Id int64 `json:"id"`
	// 文章创建时间
	CreateAt int64 `json:"create_at" gorm:"autoCreateTime"`
	// 文章更新时间
	UpdateAt int64 `json:"update_at" gorm:"autoUpdateTime"`
	// 文章状态
	Status Status `json:"status"`
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
