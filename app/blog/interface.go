package blog

import (
	"context"
	"encoding/json"
	"github.com/luyasr/simple-blog/pkg/logger"
)

const (
	Name = "blog"
)

type Service interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	DeleteBlog(context.Context, *DeleteBlogRequest) error
	UpdateBlog(context.Context, *UpdateBlogRequest) error
	QueryBlog(context.Context, *QueryBlogRequest) (*Blogs, error)
}

type CreateBlogRequest struct {
	Title    string            `json:"title" validate:"required" label:"标题"`
	Author   string            `json:"author" validate:"required" label:"作者"`
	CreateBy string            `json:"create_by" validate:"required" label:"用户"`
	Summary  string            `json:"summary" validate:"required" label:"概要"`
	Content  string            `json:"content" validate:"required" label:"内容"`
	Tags     map[string]string `json:"tags" gorm:"serializer:json" validate:"omitempty" label:"标签"`
}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{
		Tags: map[string]string{},
	}
}

type DeleteBlogRequest struct {
	Id string `json:"id" validate:"required"`
}

func NewDeleteBlogRequest() *DeleteBlogRequest {
	return &DeleteBlogRequest{}
}

type UpdateBlogRequest struct {
	Title    string            `json:"title" validate:"omitempty" label:"标题"`
	Author   string            `json:"author" validate:"omitempty" label:"作者"`
	CreateBy string            `json:"create_by" validate:"omitempty" label:"用户"`
	Summary  string            `json:"summary" validate:"omitempty" label:"概要"`
	Content  string            `json:"content" validate:"omitempty" label:"内容"`
	Tags     map[string]string `json:"tags" validate:"omitempty" label:"标签"`
}

func NewUpdateBlogRequest() *UpdateBlogRequest {
	return &UpdateBlogRequest{}
}

type QueryBlogRequest struct {
	Username   string  `json:"username" validate:"omitempty" label:"用户名"`
	Status     *Status `json:"status" validate:"omitempty" label:"状态"`
	PageSize   int     `json:"page_size" validate:"omitempty" label:"分页大小"`
	PageNumber int     `json:"page_number" validate:"omitempty" label:"分页数量"`
}

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{}
}

type Blogs struct {
	Total int     `json:"total"`
	Items []*Blog `json:"items"`
}

func (b *Blogs) String() string {
	bytes, err := json.Marshal(b)
	if err != nil {
		logger.Console.Error().Err(err).Stack().Send()
	}
	return string(bytes)
}

func (b *Blogs) Add(items ...*Blog) {
	b.Items = append(b.Items, items...)
}
