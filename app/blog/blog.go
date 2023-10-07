package blog

import (
	"context"
	"github.com/luyasr/simple-blog/config"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

var (
	_ Service = (*ServiceImpl)(nil)
)

func init() {
	ioc.Controller().Registry(&ServiceImpl{})
}

func (s *ServiceImpl) Init() error {
	s.db = config.C.Mysql.GetConn()
	return nil
}

func (s *ServiceImpl) Name() string {
	return Name
}

type ServiceImpl struct {
	db  *gorm.DB
	log zerolog.Logger
}

func (s *ServiceImpl) CreateBlog(ctx context.Context, req *CreateBlogRequest) (*Blog, error) {
	return nil, nil
}
func (s *ServiceImpl) DeleteBlog(ctx context.Context, req *DeleteBlogRequest) error {
	return nil
}
func (s *ServiceImpl) UpdateBlog(ctx context.Context, req *UpdateBlogRequest) error {
	return nil
}
func (s *ServiceImpl) QueryBlog(ctx context.Context, req *QueryBlogRequest) (*Blogs, error) {
	return nil, nil
}
