package blog

import (
	"context"
	"encoding/json"
	"github.com/luyasr/simple-blog/pkg/logger"
	"github.com/luyasr/simple-blog/pkg/utils"
)

const (
	Name = "blog"
)

type Service interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	DeleteBlog(context.Context, *DeleteBlogRequest) error
	UpdateBlog(context.Context, *UpdateBlogRequest) error
	UpdateBlogStatus(context.Context, *UpdateBlogStatusRequest) error
	QueryBlog(context.Context, *QueryBlogRequest) (*Blogs, error)
	QueryBlogById(context.Context, *QueryBlogByIdRequest) (*Blog, error)
}

type CreateBlogRequest struct {
	Title    string            `json:"title" validate:"required" label:"标题"`
	Author   string            `json:"author" validate:"required" label:"作者"`
	CreateBy string            `json:"create_by" validate:"omitempty" label:"创建用户"`
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
	Id int64 `json:"id" validate:"required"`
}

func NewDeleteBlogRequest() *DeleteBlogRequest {
	return &DeleteBlogRequest{}
}

type UpdateBlogRequest struct {
	Id       int64             `json:"id" validate:"required" label:"博客id"`
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

type UpdateBlogStatusRequest struct {
	BlogId int64  `json:"blog_id"`
	Status Status `json:"status"`
}

func NewUpdateBlogStatusRequest() *UpdateBlogStatusRequest {
	return &UpdateBlogStatusRequest{}
}

type QueryBlogRequest struct {
	Status     *Status `json:"status" validate:"omitempty" label:"状态"`
	PageSize   int     `json:"page_size" validate:"omitempty" label:"分页大小"`
	PageNumber int     `json:"page_number" validate:"omitempty" label:"分页页数"`
}

func (r *QueryBlogRequest) Offset() int {
	return r.PageSize * (r.PageNumber - 1)
}

func (r *QueryBlogRequest) SetStatus(s Status) {
	r.Status = &s
}

func (r *QueryBlogRequest) ParsePageSize(pageSize string) {
	if i := utils.StringToInt(pageSize); i != 0 {
		r.PageSize = i
	}
}

func (r *QueryBlogRequest) ParsePageNumber(pageNumber string) {
	if i := utils.StringToInt(pageNumber); i != 0 {
		r.PageNumber = i
	}
}

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageSize:   10,
		PageNumber: 1,
	}
}

type QueryBlogByIdRequest struct {
	Id int64 `json:"id" validate:"required"`
}

func NewQueryBlogByIdRequest(id int64) *QueryBlogByIdRequest {
	return &QueryBlogByIdRequest{
		Id: id,
	}
}

type Blogs struct {
	Total int64   `json:"total"`
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

func NewBlogs() *Blogs {
	return &Blogs{}
}
