package blog

import (
	"context"
	"fmt"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"testing"
)

var (
	svc Service
	ctx = context.Background()
)

func init() {
	_ = ioc.Controller().Init()
	svc = ioc.Controller().Get(Name).(Service)
}

func TestServiceImpl_CreateBlog(t *testing.T) {
	req := NewCreateBlogRequest()
	req.Title = "golang"
	req.Summary = "golang技术栈"
	req.Content = "err!=nil"
	req.Author = "admin"
	blog, err := svc.CreateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(blog)
}

func TestServiceImpl_DeleteBlog(t *testing.T) {
	req := NewDeleteBlogRequest()
	req.Id = 2
	err := svc.DeleteBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestServiceImpl_QueryBlog(t *testing.T) {
	req := NewQueryBlogRequest()

	blogs, err := svc.QueryBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(blogs)
	t.Log(blogs)
}

func TestServiceImpl_QueryBlogById(t *testing.T) {
	req := NewQueryBlogByIdRequest(6)
	blogById, err := svc.QueryBlogById(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(blogById)
}

func TestServiceImpl_UpdateBlog(t *testing.T) {
	req := NewUpdateBlogRequest()
	req.Id = 6
	req.Title = "test golang"

	err := svc.UpdateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestHandler_AuditBlog(t *testing.T) {
	req := NewAuditBlogRequest()
	req.Id = 12
	req.AuditStatus = AuditStatusAccept

	err := svc.AuditBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}
